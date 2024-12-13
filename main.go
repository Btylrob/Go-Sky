package main

import (
	"fmt"
	"os"

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
