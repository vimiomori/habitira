package main

import (
	"fmt"

	"github.com/vimiomori/habitira/cli/usr"

	"github.com/AlecAivazis/survey/v2"
	"github.com/common-nighthawk/go-figure"
)

const (
	CONFIG_DIR = "./config/habitira"
	CONFIG_NAME = "config"
	CONFIG_TYPE = "yaml"
)

func habitiraMessage() {
  myFigure := figure.NewColorFigure("HABITIRA", "puffy", "cyan", true)
  myFigure.Print()
}

func main() {
	habitiraMessage()
	user := usr.NewUser()
	if user.HasConfig() {
		// display options for view/edit
		fmt.Println("welcome back")
	} else {
		// begin setup
		fmt.Println("welcome")
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
