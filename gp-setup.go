package main

import (
	"fmt"
	"os"

	"github.com/manifoldco/promptui"
)

func initInteractive() {
	START:
		{
			startPrompt := promptui.Select{
				Label: "What are you configuring",
				Items: []string{"Language", "Shell", "Never Mind"},
			}
			_, result, err := startPrompt.Run()
			if err != nil {
				println("Error, exiting...")
				os.Exit(1)
			}
			switch result {
			case "Never Mind":
				fmt.Println("Ok, bye!")
				os.Exit(0)
			case "Shell":
				goto SHELL
			case "Language":
				goto Language
			}
		}
	SHELL:
		{
			shellPrompt := promptui.Select{
				Label: "What shell are you configuring",
				Items: []string{"Fish", "ZSH", "Back", "Never Mind"},
			}
			_, result1, err1 := shellPrompt.Run()
			if err1 != nil {
				println("Error, exiting...")
				os.Exit(1)
			}
			switch result1 {
			case "Never Mind":
				fmt.Println("Ok, bye!")
				os.Exit(0)
			case "Back":
				goto START
			}
		}
	Language:
		{
			langPrompt := promptui.Select{
				Label: "What Language are you configuring",
				Items: []string{"Julia", "Nim", "Back", "Never Mind"},
			}
			_, result2, err2 := langPrompt.Run()
			if err2 != nil {
				println("Error, exiting...")
				os.Exit(1)
			}
			switch result2 {
			case "Never Mind":
				fmt.Println("Ok, bye!")
				os.Exit(0)
			case "Back":
				goto START
			}
		}
}

func main() {
	initInteractive()
}
