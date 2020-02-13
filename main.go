package main

import (
	"fmt"
	"os"

	"strings"

	"github.com/gookit/color"
)

func initConfig() {
	config := Config{}
	color.New(color.FgLightGreen, color.Bold).Println("Initializing...")

	color.New(color.FgLightWhite, color.Bold).Print("Username to connect: ")
	fmt.Scanln(&config.Username)

	color.New(color.FgLightWhite, color.Bold).Print("Enter your password: ")
	fmt.Scanln(&config.Password)

	color.New(color.FgLightWhite, color.Bold).Print("Hostname or IP to connect: ")
	fmt.Scanln(&config.Host)

	color.New(color.FgLightWhite, color.Bold).Print("Port to conenct: ")
	fmt.Scanln(&config.Port)

	color.New(color.FgLightWhite, color.Bold).Print("Remote Path: ")
	fmt.Scanln(&config.RemoteDir)

	config.Save()
}

func main() {
	if len(os.Args) == 2 && strings.Compare(os.Args[1], "init") == 0 {
		initConfig()
	} else if len(os.Args) == 1 {
		config := LoadConfig()
		fmt.Println(config)
		WatcherTest()
	} else {
		color.New(color.FgLightWhite, color.Bold).Println("Usage: ")
		color.Green.Print("\tsyncgo init ")
		fmt.Println("(to initialize the config)")

		color.Green.Print("\tsyncgo ")
		fmt.Println("(to sync your files)")
	}
}
