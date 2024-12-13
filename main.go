package main

import (
	"fmt"
	"os"

	"github.com/nexidian/gocliselect"
)

func main() {

	bsLogo()

	menu := gocliselect.NewMenu("Chose a colour")
	menu.AddItem("Start", "red")
	menu.AddItem("Help", "blue")
	menu.AddItem("---", "green")
	menu.AddItem("", "yellow")
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
