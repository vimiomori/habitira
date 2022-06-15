package main

import (
	"github.com/vimiomori/habitira/cli/usr"

	"github.com/common-nighthawk/go-figure"
)

func habitiraMessage() {
  myFigure := figure.NewColorFigure("HABITIRA", "puffy", "cyan", true)
  myFigure.Print()
}

func main() {
	habitiraMessage()
	user := usr.NewUser()
	user.SetupConfig()

	if user.HasConfig() {
		user.WelcomeBack()
	} else {
		user.Init()
	}
	// var test string
	// initBank := &survey.Input{
	// 	Message: "Input the path to store your boba bank configurations and data:",
	// 	Default: "testtest",
	// }

	// err := survey.AskOne(initBank, &test)
	// if err != nil {
	// 	fmt.Println(err.Error())
	// 	return
	// }
}
