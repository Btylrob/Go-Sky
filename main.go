package main

import (
	"fmt"
	"os"
	"os/exec"
	"runtime"

	"github.com/fatih/color"
	"github.com/nexidian/gocliselect"
)

func main() {

	bsLogo()

	menu := gocliselect.NewMenu("Welcome to GoSky")
	menu.AddItem("Post", "red")
	menu.AddItem("Reply", "blue")
	menu.AddItem("Quote", "green")
	menu.AddItem("Help", "Help")

	choice := menu.Display()

	fmt.Println("Choice: %s\n", choice)

	if choice == "Help" {
		CallClear()
		help()

	}

}

func bsLogo() string {
	b, err := os.ReadFile("asciiGsky.txt")
	if err != nil {
		fmt.Println(err)
		return ""
	}

	fmt.Println(string(b))

	return string(b)
}

func help() {

	Bold := color.New(color.Bold)
	blue := color.New(color.FgBlue, color.Bold)

	blue.Println("General FAQ: ")
	Bold.Print("Usage: ")
	fmt.Print("Hello")
	blue.Println("Options:")

}

var clear map[string]func()

func init() {

	clear = make(map[string]func())

	clear["linux"] = func() {
		cmd := exec.Command("clear") // Linux and Unix Clear Func
		cmd.Stdout = os.Stdout
		cmd.Run()
	}
	clear["windows"] = func() {
		cmd := exec.Command("cls") // Windows clear func (use 'cls' instead of 'clear')
		cmd.Stdout = os.Stdout
		cmd.Run()
	}
}

func CallClear() {
	value, ok := clear[runtime.GOOS]
	if ok {
		value()
	} else {
		panic("Platform is unsupported!! Terminal screen will not be cleared!!")
	}
}
