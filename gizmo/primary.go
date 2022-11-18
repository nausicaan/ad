package gizmo

import (
	"fmt"
	"log"
	"strings"

	"github.com/go-ldap/ldap/v3"
)

const (
	// Colour palette.
	colorReset = "\033[0m"
	fgRed      = "\033[31m"
	fgGreen    = "\033[32m"
	fgYellow   = "\033[33m"
	fgBlue     = "\033[34m"
	fgPurple   = "\033[35m"
	fgCyan     = "\033[36m"
	fgWhite    = "\033[37m"

	bgRed    = "\033[41m"
	bgGreen  = "\033[42m"
	bgYellow = "\033[43m"
	bgBlue   = "\033[44m"

	fgBrightRed     = "\033[91m"
	fgBrightGreen   = "\033[92m"
	fgBrightYellow  = "\033[93m"
	fgBrightBlue    = "\033[94m"
	fgBrightMagenta = "\033[95m"
	fgBrightCyan    = "\033[96m"
	fgBrightWhite   = "\033[97m"

	// Constants for binding and searching the LDAP server.
	many   = 200
	fqdn   = ".CFIA-ACIA.inspection.gc.ca"
	baseDN = "DC=CFIA-ACIA,DC=inspection,DC=gc,DC=ca"
)

// Open declarations.
var computerName, groupName, printerName, userName, ldapBind string
var links *ldap.Conn
var prov int

// Valued declarations.
var ldapUser = cliu()

var ldapURL = "ldaps://" + testDomain() + fqdn

//var ldapURL = "ldaps://CFONK1AWVDCP007" + fqdn
var link, err = ldap.DialURL(ldapURL)
var badPasswordSP = []string{"badPwdCount", "badPasswordTime"}
var computerSP = []string{"name", "description", "canonicalName", "dNSHostName", "lastLogon", "userAccountControl", "operatingSystem", "operatingSystemVersion", "whenCreated"}
var userSP = []string{"physicalDeliveryOfficeName", "sAMAccountName", "name", "displayName", "title", "mail", "department", "description", "st", "lastLogon", "pwdLastSet", "homeDirectory", "userAccountControl", "canonicalName", "accountExpires", "badPwdCount", "msDS-UserPasswordExpiryTimeComputed", "whenCreated"}

// The ldapConnect function connects to the best available Domain Controller.
func ldapConnect() {
	welcome(lg)
	prov = admLocationPrompt()
	ldapBind, _ = "CN="+ldapUser+",OU="+province[prov]+",OU=Administrative Objects,"+baseDN, err
	err = link.Bind(ldapBind, loginPassword)
	good = checkError(err)
	for !good {
		fmt.Print("\n ", language[86][lg])
		enterKey()
		welcome(lg)
		ldapBind, _ = "CN="+ldapUser+",OU="+province[admLocationPrompt()]+",OU=Administrative Objects,"+baseDN, err
		err = link.Bind(ldapBind, loginPassword)
		good = checkError(err)
	}
}

// The ldapConnect function reconnects to the targeted Domain Controller in order to search AD.
func ldapReconnect() {
	link, err = ldap.DialURL(ldapURL)
	ldapBind, _ = "CN="+ldapUser+",OU="+province[prov]+",OU=Administrative Objects,"+baseDN, err
	err = link.Bind(ldapBind, loginPassword)
}

// The ldapMultiConnect function connects systematically to all the Domain Controllers.
func ldapMultiConnect(e string) {
	ldapSearchURL := "ldaps://" + e + fqdn
	links, err = ldap.DialURL(ldapSearchURL)
	err = links.Bind(ldapBind, loginPassword)
	good = checkError(err)
}

// The testDomain function finds the connection speeds of the available Domain Controllers.
func testDomain() string {
	clear()
	var bestDC, pingDCstring string
	var pingDCint int
	fastestTime := 9999

	fmt.Println(fgBrightGreen, "Finding fastest Domain Controllers...")
	fmt.Println(fgBrightYellow, "Testing Domain Controller speed...")

	for _, s := range cfia {
		fmt.Println(fgBrightMagenta, s+fqdn)
		pingDCstring = strings.TrimSpace(powerShellRVS("Test-Connection -ComputerName " + s + fqdn + " -Count 1 -ErrorAction SilentlyContinue | Select -exp ResponseTime"))
		pingDCint = intFromString(pingDCstring)

		if pingDCint <= fastestTime {
			fastestTime = pingDCint
			bestDC = s
		}
	}
	return bestDC
}

// The searchDomainControllers function searches all known DC's for bad password counts.
func searchDomainControllers() {
	index = 0
	redIndex = 0
	yellowIndex = 0
	redDCs = redDCs[:0]
	yellowDCs = yellowDCs[:0]

	for _, s := range cfia {
		ldapMultiConnect(s)
		query(links, filterSAM, badPasswordSP, userName)
		searchBadPassword(result, s)
		links.Close()
		index++
	}
}

// The lcid function determines the base language of the operating system.
func lcid() int {
	oslang := 0
	display := strings.TrimSpace(powerShellRVS("Get-Culture | Select -exp LCID"))
	fre := intFromString(display)

	if fre == 3084 {
		oslang = 1
	}
	return oslang
}

// The orca function will verify if the specified user is an ORCA member or not.
func orca() {
	fmt.Println("\nYou chose 0")
	enterKey()
}

// The password function is used to reset a user password in AD. It asks for a new password, if the user must change password at next logon, for a confirmation and if the user wants to check if the account is locked out.
func changePassword(link *ldap.Conn) {
	userName = usPrompt()
	ldapReconnect()
	fmt.Println(language[80][lg])
	oldPassword := getInput("\n Enter your current password: ")
	fmt.Println(language[87][lg])
	newPassword := getInput(language[81][lg])
	confirmPassword := getInput(language[82][lg])

	if confirmPassword == newPassword {
		passwdModReq := ldap.NewPasswordModifyRequest(userName, oldPassword, newPassword)

		if _, err = link.PasswordModify(passwdModReq); err != nil {
			log.Fatalf("failed to modify password: %v", err)
		}
		fmt.Println(language[88][lg])
	} else {
		fmt.Println(language[83][lg])
	}
	link.Close()
	enterKey()
}

// The locked function will verify if an account is locked out. If yes, it will propose to unlock it.
func locked() {
	userName = usPrompt()
	fmt.Print("\n ", language[90][lg])
	searchDomainControllers()
	printControllerValues()
	callForUnlock()
	enterKey()
}

// The entity function asks the user for a username and pulls the account information from Active Directory. It also gives quick hints & warnings about the account (ex. if expired, disabled, etc.).
func entity() {
	userName = usPrompt()
	ldapReconnect()
	query(link, filterSAM, userSP, userName)
	assignObjectValues(result)
	link.Close()
	enterKey()
}

// The computer function asks the user for a computer name and pulls the machine information from Active Directory.
func computer() {
	computerName = csPrompt()
	fmt.Println("\n Checking for...", computerName)
	ldapReconnect()
	query(link, filterName, computerSP, computerName)
	assignObjectValues(result)
	link.Close()
	enterKey()
}

// The group function asks for a group name and then searches Active Directory.
func group() {
	groupName = gsPrompt()
	fmt.Println("\n Checking for...", groupName)
	enterKey()
}

// The printer function will ask for printer name, will retrieve the information from AD and test it. Optionally, you can retrieve the full list of CFIA printers.
func printer() {
	printerName = psPrompt()
	fmt.Println("\n Checking for...", printerName)
	enterKey()
}
