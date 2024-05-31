package repl

import (
	"fmt"
	"strings"

	"github.com/otaleghani/sbes/internal/database"
	"github.com/otaleghani/sbes/internal/terminalinput"
)

func cmdUpdateDomainOAuth() {
	divider()
	domain := strings.TrimSpace(
		terminalinput.ReadInput("Enter new OAuth domain \n\r-> "))
	err := database.UpdateDomainOAuth(domain)
	if err != nil {
		fmt.Println("ERROR: ", err)
		return
	}
	divider()
	fmt.Println("SUCCESS: OAuth domain updated")
}

func cmdUpdateDomain() {
	divider()
	domain := strings.TrimSpace(
		terminalinput.ReadInput("Enter new tracker domain \n\r-> "))
	err := database.UpdateDomain(domain)
	if err != nil {
		fmt.Println("ERROR: ", err)
		return
	}
	divider()
	fmt.Println("SUCCESS: Tracker domain updated")
}
