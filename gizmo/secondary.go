package gizmo

import (
	"fmt"
	"os"
	"os/exec"
	"os/user"
	"strconv"
	"strings"

	"github.com/go-ldap/ldap/v3"
)

type ADObject struct {
	sam, name, display, description, path, lastLogon string
	// Embedded structures
	employee Employee
	machine  Machine
}

type Employee struct {
	email, department, title, description, office, province, path, directory string
}

type Machine struct {
	fqdn string
	os   string
	osv  string
}

type Values struct {
	accountCode, accountCreated, badPasswordCount, accountExpires, passwordExpires, passwordLastSet string
	// Embedded structure
	derived Derived
}

type Derived struct {
	accountLocked, accountDisabled, accountExpired, passwordExpired, passwordNeverExpires string
}

// Constants for searching Active Directory.
const (
	adInfinity = 9223372036854775807
	filterSAM  = "(sAMAccountName=%s)"
	filterName = "(name=%s)"
	isTrue     = "True"
	isNever    = "Never"
)

// Open declarations.
var result *ldap.SearchResult
var item *ldap.Entry
var index, redIndex, yellowIndex int
var controllers, redDCs, yellowDCs []string
var good, adm bool

// Valued declarations.
var now = strings.TrimSpace(powerShellRVS("(Get-date).TofileTime()"))
var ps, _ = exec.LookPath("powershell")

//var redDCs []string
var ado = ADObject{}
var uav = Values{}

// The getInput function takes a string prompt and asks the user for input.
func getInput(prompt string) string {
	fmt.Print("\n ", prompt)
	userInput, _ := reader.ReadString('\n')
	userInput = strings.TrimSpace(userInput)
	return userInput
}

// The checkError function executes the builtin panic function if an error is detected.
func checkError(err error) bool {
	var check bool
	if err != nil {
		check = false
	} else {
		check = true
	}
	return check
}

// The enterKey function pauses the transition to the nest screen until the enter key is pressed.
func enterKey() {
	fmt.Print("\n ", colorReset, language[27][lg])
	fmt.Scanln()
}

// The clear function clears the terminal or screen.
func clear() {
	powerShellEXE("cls")
	fmt.Println(colorReset)
}

// The localPC function gets the name of the local computer.
func localPC() string {
	pc, _ := os.Hostname()
	good = checkError(err)
	return pc
}

// The cliu function gets the name of the local user.
func cliu() string {
	person, err := user.Current()
	good = checkError(err)
	loginName = person.Username
	disn := person.Name
	adm = true

	if !strings.Contains(loginName, "adm-") {
		adm = false
		s := strings.Split(disn, " ")
		disnf := strings.TrimSpace(s[1])
		disnl := strings.TrimSpace(strings.TrimSuffix(s[0], ","))
		disn = disnf + " " + disnl
		loginName = "adm-" + loginName
	}
	return disn
}

// The admWarning function prints a warning if the logged in account does not have ADM privledges.
func admWarning(lg int) {
	if !adm {
		fmt.Println(fgRed, language[3][lg])
		fmt.Println()
	}
}

// The query function searches AD though an established LDAP connection.
func query(link *ldap.Conn, filter string, search []string, object string) {
	// Filters must start and finish with ()!
	filterDN := fmt.Sprintf(filter, ldap.EscapeFilter(object))
	searchReq := ldap.NewSearchRequest(baseDN, ldap.ScopeWholeSubtree, ldap.NeverDerefAliases, 0, 0, false, filterDN, search, []ldap.Control{})

	result, err = link.Search(searchReq)
	if err != nil {
		fmt.Println("Failed to query LDAP: ", err)
		return
	}
}

// The searchBadPassword function scans a specified DC to determine how many bad password attempts have occured.
func searchBadPassword(drip *ldap.SearchResult, controller string) {
	var bpcs string
	var bpc int

	for _, item = range drip.Entries {
		bpcs = item.GetAttributeValue("badPwdCount")
		bpc = intFromString(bpcs)
		if bpc > 2 {
			controllers = append(controllers[:index], fgRed+controller+"           "+bpcs+"            "+language[176][lg])
			redDCs = append(redDCs[:redIndex], controller)
			redIndex++
		} else if bpc < 3 && bpc > 0 {
			controllers = append(controllers[:index], fgYellow+controller+"           "+bpcs+"            "+language[177][lg])
			yellowDCs = append(yellowDCs[:yellowIndex], controller)
			yellowIndex++
		} else {
			controllers = append(controllers[:index], fgGreen+controller+"           0             "+language[178][lg])
		}
	}
}

// The assignObjectValues function builds two objects and uses the contained information to determine user statuses such as password or account expiration.
func assignObjectValues(nond *ldap.SearchResult) {
	for _, item = range nond.Entries {
		ado = ADObject{item.GetAttributeValue("sAMAccountName"), item.GetAttributeValue("name"), item.GetAttributeValue("displayName"), item.GetAttributeValue("description"), item.GetAttributeValue("canonicalName"), strings.TrimSpace(convertLargeInt(item.GetAttributeValue("lastLogon"))), Employee{item.GetAttributeValue("mail"), item.GetAttributeValue("department"), item.GetAttributeValue("title"), item.GetAttributeValue("description"), item.GetAttributeValue("physicalDeliveryOfficeName"), item.GetAttributeValue("st"), item.GetAttributeValue("canonicalName"), item.GetAttributeValue("homeDirectory")}, Machine{item.GetAttributeValue("dNSHostName"), item.GetAttributeValue("operatingSystem"), item.GetAttributeValue("operatingSystemVersion")}}
		uav = Values{item.GetAttributeValue("userAccountControl"), splitOZ(item.GetAttributeValue("whenCreated")), item.GetAttributeValue("badPwdCount"), item.GetAttributeValue("accountExpires"), item.GetAttributeValue("msDS-UserPasswordExpiryTimeComputed"), strings.TrimSpace(convertLargeInt(item.GetAttributeValue("pwdLastSet"))), Derived{"False", "False", "False", "False", "False"}}

		if intFromString(uav.badPasswordCount) > 2 || uav.accountCode == "2" {
			uav.derived.accountLocked = isTrue
		}

		if intFromString(uav.accountExpires) == 0 || uav.accountExpires == "" || intFromString(uav.accountExpires) == adInfinity {
			uav.accountExpires = isNever
		} else {
			if now > uav.accountExpires {
				uav.derived.accountExpired = isTrue
			}
			uav.accountExpires = strings.TrimSpace(convertLargeInt(item.GetAttributeValue("accountExpires")))
		}

		if intFromString(uav.passwordExpires) == 0 || uav.passwordExpires == isNever || intFromString(uav.passwordExpires) == adInfinity {
			uav.passwordExpires = isNever
			uav.derived.passwordNeverExpires = isTrue
		} else {
			if now > uav.passwordExpires {
				uav.derived.passwordExpired = isTrue
			}
			uav.passwordExpires = strings.TrimSpace(convertLargeInt(item.GetAttributeValue("msDS-UserPasswordExpiryTimeComputed")))
		}

		if ado.employee.email != "" {
			printUserValues()
		} else {
			printComputerValues()
		}
	}
}

// The printUserValues function prints all the previously collected user information to the terminal.
func printUserValues() {
	clear()
	fmt.Print(" ", fgGreen, language[130][lg])
	fmt.Println(fgBrightYellow, ado.display, colorReset)
	fmt.Println("\n", language[107][lg], "                   :", ado.sam)
	fmt.Println("", language[108][lg], "                  :", ado.name)
	fmt.Println(" Display Name                :", ado.display)
	fmt.Println(" Email Address               :", ado.employee.email)
	fmt.Println(" Department                  :", ado.employee.department)
	fmt.Println(" Title                       :", ado.employee.title)
	fmt.Println(" Description                 :", ado.employee.description)
	fmt.Println(" Office                      :", ado.employee.office)
	fmt.Println(" Province                    :", ado.employee.province)
	fmt.Println(" AD Path                     :", ado.employee.path)
	fmt.Println(" Home Directory              :", ado.employee.directory)
	fmt.Println(" Last Logon                  :", ado.lastLogon)
	fmt.Println(" Account Created             :", uav.accountCreated)
	fmt.Println(" Account Disabled            :", uav.derived.accountDisabled)
	fmt.Println(" Account Expired             :", uav.derived.accountExpired)
	fmt.Println(" Account Expires             :", uav.accountExpires)
	fmt.Println(" Password Last Set           :", uav.passwordLastSet)
	fmt.Println(" Password Expires            :", uav.passwordExpires)
	fmt.Println(" Password Expired            :", uav.derived.passwordExpired)
	fmt.Println(" Password Never Expires      :", uav.derived.passwordNeverExpires)
}

// The printUserValues function prints all the previously collected computer information to the terminal.
func printComputerValues() {
	clear()
	fmt.Print(" ", fgGreen, language[105][lg])
	fmt.Println(fgBrightYellow, ado.name, colorReset)
	fmt.Println("\n Full Name                 :", ado.name)
	fmt.Println(" AD Path                   :", ado.employee.path)
	fmt.Println(" FQDN                      :", ado.machine.fqdn)
	fmt.Println(" Description               :", ado.employee.description)
	fmt.Println(" Last Logon                :", ado.lastLogon)
	fmt.Println(" Computer Created          :", uav.accountCreated)
	fmt.Println(" Operating System          :", ado.machine.os)
	fmt.Println(" OS Version                :", ado.machine.osv)
}

// The printControllerValues function prints all the previously collected bad password information accross all DC's to the terminal.
func printControllerValues() {
	clear()
	fmt.Println("     DC Name            Attempts        Status")
	fmt.Println(" ---------------        --------        ------")
	for _, s := range controllers {
		fmt.Println(" " + s)
	}
}

// The callForUnlock function asks the user if an unlock should be attempted and responds accordingly.
func callForUnlock() {
	if (len(redDCs)) > 0 {
		fmt.Println("\n"+fgCyan, userName+colorReset, language[93][lg])
		answer := strings.ToUpper(getInput(language[94][lg]))
		if answer == "Y" {
			fmt.Print("\n ", language[95][lg], fgCyan, userName)
			fmt.Println(colorReset)
			unlock()
			fmt.Print("\n ", language[97][lg], fgCyan, userName)
			fmt.Println(colorReset)
		}
	} else {
		fmt.Println("\n", colorReset+language[91][lg]+fgCyan, userName, colorReset+language[92][lg])
	}
}

// The unlock function attempts to unlock the user using the Unlock-ADAccount PowerShell function.
func unlock() {
	for _, s := range redDCs {
		ldapMultiConnect(s)
		powerShellEXE("Unlock-ADAccount -Server " + s + fqdn + " -Identity " + userName)
	}
	for _, s := range yellowDCs {
		ldapMultiConnect(s)
		powerShellEXE("Unlock-ADAccount -Server " + s + fqdn + " -Identity " + userName)
	}
}

// The testConnection function will test the connection to a computer.
func testConnection() {
	var pingArray = [3]int{}
	var avgSpeed int

	atConfirm()
	fmt.Println(fgYellow, language[149][lg], fgWhite+computerName+"\n") // Checking connection to

	for index = 0; index < 3; index++ {
		pingArray[index] = pingComputer()
		if pingArray[index] > 0 {
			fmt.Println(" "+language[179][lg], pingArray[index], "ms")
		} else {
			fmt.Println(" Request timed out")
		}
	}
	avgSpeed = (pingArray[0] + pingArray[1] + pingArray[2]) / 3

	if avgSpeed > 0 {
		fmt.Println(" "+language[180][lg], avgSpeed, "ms")
		fmt.Println("\n"+fgGreen, language[150][lg]) // Connection succeeded!
	} else {
		fmt.Println("\n"+fgRed, language[151][lg]) // Connection failed!
	}
	enterKey()
}

func pingComputer() int {
	var pingTest string
	var pingResult int

	pingTest = strings.TrimSpace(powerShellRVS("Test-Connection -ComputerName " + computerName + " -Count 1 -ErrorAction SilentlyContinue | Select -exp ResponseTime"))
	pingResult = intFromString(pingTest)

	return pingResult
}

// The disableCard function will disable a network card on a remote computer.
func disableCard() {
	adapters := powerShellRVS("Get-NetAdapter -Name * -IncludeHidden | Format-List -Property deviceid,name")
	fmt.Print(adapters)
	fmt.Print(" " + language[153][lg] + ": ")
	deviceID, _ := reader.ReadString('\n')
	deviceID = strings.Title(strings.Replace(deviceID, "\r\n", "", -1))

	if deviceID == "10" {
		fmt.Println(deviceID)
	} else {
		fmt.Println(deviceID)
	}

	fmt.Println(fgRed, language[158][lg])
	enterKey()
}

// The restartService function will restart a named service or services.
func restartService(serviceName string) {
	powerShellEXE("Restart-Service -Name " + serviceName)
}

// The stopService function will stop a named service or services.
func stopService(serviceName string) {
	powerShellEXE("Stop-Service -Name " + serviceName)
}

// The reboot function will reboot a remote computer.
func reboot() {
	powerShellEXE("Restart-Computer -ComputerName " + "'" + computerName + "'" + " -Force")
}

// The logoff function will force a logoff.
func logoff() {
	powerShellEXE("shutdown /l /m \\" + computerName + " /t 0")
}

// The admLocationPrompt function returns a number corresponding to the chosen province.
func admLocationPrompt() int {
	regionTitle()
	regionMenu()
	return intFromString(getInput("Select your location: "))
}

// The atConfirm function displays the computer connected to the application.
func atConfirm() {
	clear()
	fmt.Print(" " + language[15][lg] + ":") // Advanced tools are currently linked to
	fmt.Println(fgGreen, computerName)
	fmt.Println()
}

// The usPrompt function loads the User Search dialog and prompts for a name to search.
func usPrompt() string {
	clear()
	usTitle()
	fmt.Println("\n [ex. beo, beordma, beo*dma...]")
	return getInput(language[65][lg])
}

// The csPrompt function loads the Computer Search dialog and prompts for a name to search.
func csPrompt() string {
	clear()
	csTitle()
	fmt.Println("\n [ex. ncdec652445, 652445, 65244*, 10.141.12.58, ncdec652445.cfia-acia.inspection.gc.ca...]")
	return getInput(language[71][lg])
}

// The psPrompt function loads the Printer Search dialog and prompts for a name to search.
func psPrompt() string {
	clear()
	psTitle()
	fmt.Println("\n [ex. P141022202, 10.136.52.188...]")
	return getInput(language[75][lg])
}

// The gsPrompt function loads the Group Search dialog and prompts for a name to search.
func gsPrompt() string {
	clear()
	gsTitle()
	return getInput(language[78][lg])
}

// The timezone function returns the current difference from UTC.
func timezone() string {
	return strings.TrimSpace(powerShellRVS("Get-TimeZone | Select -exp BaseUtcOffset | Select -exp Hours"))
}

// The intFromString function converts a String value to an Integer.
func intFromString(convertable string) int {
	freshInt, _ := strconv.Atoi(convertable)
	return freshInt
}

func splitOZ(oz string) string {
	strings.TrimSuffix(oz, ".0Z")
	year, month, day, hour, min, sec := oz[0:4], oz[4:6], oz[6:8], oz[8:10], oz[10:12], oz[12:14]
	utc := year + "-" + month + "-" + day + " " + hour + ":" + min + ":" + sec
	return utc
}

// The convertLargeInt function converts a Large Integer to a String value.
func convertLargeInt(value string) string {
	return powerShellRVS("(Get-Date 1/1/1601).AddDays(" + value + "/864000000000).AddHours(" + timezone() + ")")
}

// The powershellEXE function executes a PowerShell command directly.
func powerShellEXE(task string) {
	psCmd := exec.Command(ps, task)
	psCmd.Stdout = os.Stdout
	psCmd.Stderr = os.Stderr

	err := psCmd.Run()
	good = checkError(err)
}

// The powershellRVS function runs a PowerShell command and returns the output as a String.
func powerShellRVS(task string) string {
	psCmd := exec.Command(ps, task)
	psOut, _ := psCmd.CombinedOutput()
	return string(psOut)
}
