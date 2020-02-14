package main

import (
	"fmt"
	"os"

	"strings"

	"github.com/gookit/color"
)

func initConfig() {
	config := Config{}
	fmt.Print("⫸ ")
	color.New(color.FgLightGreen, color.Bold).Println("Initializing...")

	color.Green.Print("? ")
	color.New(color.FgLightWhite, color.Bold).Print("Username to connect: ")
	fmt.Scanln(&config.Username)

	color.Green.Print("? ")
	color.New(color.FgLightWhite, color.Bold).Print("Enter your password: ")
	fmt.Scanln(&config.Password)

	color.Green.Print("? ")
	color.New(color.FgLightWhite, color.Bold).Print("Hostname or IP to connect: ")
	fmt.Scanln(&config.Host)

	color.Green.Print("? ")
	color.New(color.FgLightWhite, color.Bold).Print("Port to conenct: ")
	fmt.Scanln(&config.Port)

	color.Green.Print("? ")
	color.New(color.FgLightWhite, color.Bold).Print("Remote Path: ")
	fmt.Scanln(&config.RemoteDir)

	config.Save()
}

func main() {
	if len(os.Args) == 2 && strings.Compare(os.Args[1], "init") == 0 {
		initConfig()
	} else if len(os.Args) == 1 {
		config := LoadConfig()
		RunWatcher(config)
	} else {
		color.New(color.FgLightWhite, color.Bold).Println("Usage: ")
		color.Green.Print("\tsyncgo init ")
		fmt.Println("(to initialize the config)")

		color.Green.Print("\tsyncgo ")
		fmt.Println("(to sync your files)")
	}
}
