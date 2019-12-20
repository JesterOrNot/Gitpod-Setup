package main

import (
	"fmt"
	"os"

	"github.com/manifoldco/promptui"
)

func initInteractive() {
	isError := false
	START:
		if isError {
			fmt.Println("\033[1;31mError Reading Option\033[0m")
			isError = false
		}
		{
			startPrompt := promptui.Select{
				Label: "What are you configuring",
				Items: []string{"Language", "Shell", "Never Mind"},
			}
			_, result, err := startPrompt.Run()
			if err != nil {
				exit()
			}
			switch result {
				case "Never Mind":
					exit()
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
				exit()
			}
			switch result1 {
				case "Never Mind":
					exit()
				case "Back":
					goto START
				default:
					isError = true
					goto START
			}
		}
	Language:
		{
			langPrompt := promptui.Select{
				Label: "What Language are you configuring",
				Items: []string{"Julia", "Nim", "Hy", "Clojure", "Haskell", "Back", "Never Mind"},
			}
			_, result2, err2 := langPrompt.Run()
			if err2 != nil {
				exit()
			}
			switch result2 {
				case "Julia":
					juliaInit()
				case "Nim":
					nimInit()
				case "Hy":
					hyInit()
				case "Clojure":
					clojureInit()
				case "Never Mind":
					exit()
				case "Back":
					goto START
				default:
					isError = true
					goto START
			}
		}
}
func juliaInit()  {
	initBase(juliaDockerFile, juliaYaml)
}
func nimInit() {
	initBase(nimDockerFile, nimYaml)
}
func hyInit() {
	initBase(hyDockerfile, hyYaml)
}
func clojureInit()  {
	initBase(clojureDockerfile, clojureYaml)
}
// func haskellInit() {
// 	initBase(haskellDockerfile, haskellYaml)
// }
func initBase(dockerFile, Yaml string) {
	gitpodDockerfile, _ := os.Create(".gitpod.Dockerfile")
	gitpodYaml, _ := os.Create(".gitpod.yml")
	gitpodDockerfile.WriteString(dockerFile)
	gitpodYaml.WriteString(Yaml)
}
func exit() {
	fmt.Println("Ok, bye!")
	os.Exit(0)
}
func main() {
	initInteractive()
}
