package main

import (
	"fmt"
	"os"

	"github.com/fatih/color"
	"github.com/nexidian/gocliselect"
)

func main() {

	bsLogo()

	menu := gocliselect.NewMenu("Welcome to GoSky")
	menu.AddItem("Post", "red")
	menu.AddItem("Reply", "blue")
	menu.AddItem("Quote", "green")
	menu.AddItem("Help", "cyan")

	choice := menu.Display()

	fmt.Println("Choice: %s\n", choice)

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

	color.Blue("General FAQ: ")
	Bold.Printf("Usage: BlueSky Terminal Posting  Applicaction using Go and BlueSky API \nGithubID: Btylrob\nRepository\n")
	color.Blue("Options:")

}
