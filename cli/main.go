package main

import (
	"fmt"
	"os"
	osUser "os/user"

	"github.com/vimiomori/habitira/cli/usr"

	"github.com/AlecAivazis/survey/v2"
	"github.com/common-nighthawk/go-figure"
)

func habitiraMessage() {
  myFigure := figure.NewColorFigure("HABITIRA", "puffy", "cyan", true)
  myFigure.Print()
}

func main() {
	habitiraMessage()
	homeDir, _ := os.UserHomeDir()
	osName, _ := osUser.Current()
	user := usr.NewUser(homeDir, osName.Name)
	if user.HasConfig() {
		// display options for view/edit
		fmt.Println("welcome back")
	} else {
		// begin setup
		usr, _ := osUser.Current()
		fmt.Printf("Welcome %s!!", usr.Name)
	}

	var test string
	initBank := &survey.Input{
		Message: "Input the path to store your boba bank configurations and data:",
		Default: "testtest",
	}

	err := survey.AskOne(initBank, &test)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
}
