package gizmo

import (
	"fmt"
)

// The regionMenu function displays a list of Provinces to choose from,
func regionMenu() {
	fmt.Println("\n    0 - Alberta")
	fmt.Println("    1 - BC")
	fmt.Println("    2 - Manitoba")
	fmt.Println("    3 - New Brunswick")
	fmt.Println("    4 - NCR")
	fmt.Println("    5 - Newfoundland")
	fmt.Println("    6 - Ontario")
	fmt.Println("    7 - PEI")
	fmt.Println("    8 - Quebec")
	fmt.Println("    9 - Saskatchewan")
}

// The welcome function displays the program name and author information.
func welcome(lg int) {
	clear()
	fmt.Print(" F - Français")
	fmt.Println("\tE - English")
	fmt.Println("\n"+bgGreen, fgWhite, "                                                                           ")
	fmt.Println(`     ___     ___     ___       o-o              o      o-O-o        o        `)
	fmt.Println(`    (_-<    / _ \   (_-<      o   o      o      | /      |          |        `)
	fmt.Println(`    /__/_   \___/   /__/_     |   | o  o    o-o OO       |  o-o o-o | o-o    `)
	fmt.Println(`   |"""""|_|"""""|_|"""""|    o   O |  | | |    | \      |  | | | | |  \     `)
	fmt.Println(`    -0-0-   -0-0-   -0-0-      o-O\ o--o |  o-o o  o     o  o-o o-o o o-o    `)
	fmt.Println("                                                                             ")
	fmt.Println("                       " + language[0][lg] + "                     ")
	fmt.Println("                                                                            ", colorReset)
	fmt.Println("\n"+fgBrightBlue, language[173][lg]+fgWhite, "Marc-Antoine Beord (marc-antoine.beord@ssc-spc.gc.ca)")
	fmt.Println(fgBrightCyan, language[1][lg]+fgWhite, "Byron Stuike (byron.stuike@inspection.gc.ca)")
	fmt.Println(fgBrightYellow, language[2][lg]+fgWhite, "Byron Stuike (byron.stuike@inspection.gc.ca)")
	fmt.Println("\n", language[4][lg]+fgGreen, ldapUser)
	fmt.Println(colorReset)
}

// The mainMenu function displays the list of initial options.
func mainMenu(lg int) {
	mainTitle()
	fmt.Println("\n    0 -", language[159][lg]) // ORCA Status Verification
	fmt.Println("    1 -", language[6][lg])     // RESET Password
	fmt.Println("    2 -", language[7][lg])     // LOCKED OUT account
	fmt.Println("\n    3 -", language[8][lg])   // USER Information
	fmt.Println("    4 -", language[9][lg])     // COMPUTER Information
	fmt.Println("    5 -", language[10][lg])    // PRINTER Information
	fmt.Println("    6 -", language[11][lg])    // GROUP Information
	fmt.Println("\n    7 -", language[12][lg])  // Advanced Computer Tools
	fmt.Println("\n    9 -", language[14][lg])  // Exit
}

// The advancedMenu function displays the Advanced Tools menu.
func advancedMenu(lg int) {
	advancedTitle()
	fmt.Println("\n    1 -", language[19][lg]) // Force Logoff
	fmt.Println("    2 -", language[20][lg])   // Restart Computer
	fmt.Println("    3 -", language[21][lg])   // Test Network Connection
	fmt.Println("\n    4 -", language[22][lg]) // Disable Network Card
	fmt.Println("    5 -", language[23][lg])   // Process Tools
	fmt.Println("    6 -", language[24][lg])   // Service Tools
	fmt.Println("\n    8 -", language[13][lg]) // Back
	fmt.Println("    9 -", language[14][lg])   // Exit
}

// The serviceMenu function displays the Service Tools menu.
func serviceMenu(lg int) {
	subTitle(24)
	fmt.Println("\n    1 -", language[39][lg]) // Get service list
	fmt.Println("    2 -", language[40][lg])   // Start service(s)
	fmt.Println("    3 -", language[41][lg])   // Restart service(s)
	fmt.Println("    4 -", language[42][lg])   // Stop Service(s)
	fmt.Println("\n    8 -", language[13][lg]) // Back
	fmt.Println("    9 -", language[14][lg])   // Exit
}

// The processes function displays the Process Tools menu.
func processMenu(lg int) {
	subTitle(23)
	fmt.Println("\n    1 -", language[28][lg]) // Get process list
	fmt.Println("    2 -", language[29][lg])   // Terminate process
	fmt.Println("\n    8 -", language[13][lg]) // Back
	fmt.Println("    9 -", language[14][lg])   // Exit
}

// The regionTitle surrounds the Region Menu title with yellow stars.
func regionTitle() {
	if lg == 0 {
		fmt.Println(fgYellow, "**************")
		fmt.Println(" **", fgWhite, language[175][lg], fgYellow, "**")
		fmt.Println(" **************", colorReset)
	} else {
		fmt.Println(fgYellow, "*******************")
		fmt.Println(" **", fgWhite, language[175][lg], fgYellow, "**")
		fmt.Println(" *******************", colorReset)
	}
}

// The advancedTitle surrounds the Advanced Tools title with yellow stars.
func advancedTitle() {
	fmt.Println(fgYellow, "**********************")
	fmt.Println(" **", fgWhite, language[18][lg], fgYellow, "**")
	fmt.Println(" **********************", colorReset)
}

// The mainTitle surrounds the Main Menu title with yellow stars.
func mainTitle() {
	if lg == 0 {
		fmt.Println(fgYellow, "*****************")
		fmt.Println(" **", fgWhite, language[174][lg], fgYellow, "**")
		fmt.Println(" *****************", colorReset)
	} else {
		fmt.Println(fgYellow, "**********************")
		fmt.Println(" **", fgWhite, language[174][lg], fgYellow, "**")
		fmt.Println(" **********************", colorReset)
	}
}

// The subTitle surrounds the Service or Process Tools title with yellow stars.
func subTitle(num int) {
	if lg == 0 {
		fmt.Println(fgYellow, "*********************")
		fmt.Println(" **", fgWhite, language[num][lg], fgYellow, "**")
		fmt.Println(" *********************", colorReset)
	} else if lg == 1 && num == 23 {
		fmt.Println(fgYellow, "*********************************")
		fmt.Println(" **", fgWhite, language[num][lg], fgYellow, "**")
		fmt.Println(" *********************************", colorReset)
	} else {
		fmt.Println(fgYellow, "********************************")
		fmt.Println(" **", fgWhite, language[num][lg], fgYellow, "**")
		fmt.Println(" ********************************", colorReset)
	}
}

// The csTitle surrounds the Computer Search title with yellow stars.
func csTitle() {
	if lg == 0 {
		fmt.Println(fgYellow, "***********************")
		fmt.Println(" **", fgWhite, language[70][lg], fgYellow, "**")
		fmt.Println(" ***********************", colorReset)
	} else {
		fmt.Println(fgYellow, "******************************")
		fmt.Println(" **", fgWhite, language[70][lg], fgYellow, "**")
		fmt.Println(" ******************************", colorReset)
	}
}

// The usTitle surrounds the User Search title with yellow stars.
func usTitle() {
	if lg == 0 {
		fmt.Println(fgYellow, "*******************")
		fmt.Println(" **", fgWhite, language[64][lg], fgYellow, "**")
		fmt.Println(" *******************", colorReset)
	} else {
		fmt.Println(fgYellow, "**************************")
		fmt.Println(" **", fgWhite, language[64][lg], fgYellow, "**")
		fmt.Println(" **************************", colorReset)
	}
}

// The psTitle surrounds the Printer Search title with yellow stars.
func psTitle() {
	if lg == 0 {
		fmt.Println(fgYellow, "**********************")
		fmt.Println(" **", fgWhite, language[74][lg], fgYellow, "**")
		fmt.Println(" **********************", colorReset)
	} else {
		fmt.Println(fgYellow, "******************************")
		fmt.Println(" **", fgWhite, language[74][lg], fgYellow, "**")
		fmt.Println(" ******************************", colorReset)
	}
}

// The gsTitle surrounds the Group Search title with yellow stars.
func gsTitle() {
	if lg == 0 {
		fmt.Println(fgYellow, "********************")
		fmt.Println(" **", fgWhite, language[77][lg], fgYellow, "**")
		fmt.Println(" ********************", colorReset)
	} else {
		fmt.Println(fgYellow, "***************************")
		fmt.Println(" **", fgWhite, language[77][lg], fgYellow, "**")
		fmt.Println(" ***************************", colorReset)
	}
}
