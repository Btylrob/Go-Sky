package main

import (
	"fmt"

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
	fmt.Print(


		██████╗  ██████╗ ███████╗██╗  ██╗██╗   ██╗
		██╔════╝ ██╔═══██╗██╔════╝██║ ██╔╝╚██╗ ██╔╝
		██║  ███╗██║   ██║███████╗█████╔╝  ╚████╔╝ 
		██║   ██║██║   ██║╚════██║██╔═██╗   ╚██╔╝  
		╚██████╔╝╚██████╔╝███████║██║  ██╗   ██║   
		 ╚═════╝  ╚═════╝ ╚══════╝╚═╝  ╚═╝   ╚═╝   
		
		

	)
}