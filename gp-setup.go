package main

import (
	"fmt"
	"github.com/manifoldco/promptui"
	"os"
	"strings"
)

func main() {
	if len(os.Args) == 1 {
		fmt.Println("Help:\n    init: initialize gitpod")
		return
	}
	if os.Args[1] == "init" {
		if len(os.Args) >= 3 {
			switch os.Args[2] {
			case "julia":
				juliaInit()
				return
			case "nim":
				nimInit()
				return
			case "hy":
				hyInit()
				return
			case "clojure":
				clojureInit()
				return
			case "haskell":
				haskellInit()
				return
			case "dotnet":
				dotNetInit()
				return
			case "zsh":
				zshInit()
				return
			default:
				fmt.Println("Invalid argument '" + strings.Join(os.Args[2:], " ") + "'")
				return
			}
		}
		initInteractive()
		return
	}
}

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
			Items: []string{"ZSH", "Back", "Never Mind"},
		}
		_, result1, err1 := shellPrompt.Run()
		if err1 != nil {
			exit()
		}
		switch result1 {
		case "ZSH":
			zshInit()
			return
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
			Items: []string{"Julia", "Nim", "Hy", "Clojure", "Haskell", ".NET", "Back", "Never Mind"},
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
		case "Haskell":
			haskellInit()
		case ".NET":
			dotNetInit()
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
func juliaInit() {
	initBase(juliaDockerFile, juliaYaml)
	fmt.Println("Julia Setup Complete!")
}
func nimInit() {
	initBase(nimDockerFile, nimYaml)
	fmt.Println("Nim Setup Complete!")
}
func hyInit() {
	initBase(hyDockerfile, hyYaml)
	fmt.Println("Hy Setup Complete!")
}
func clojureInit() {
	initBase(clojureDockerfile, clojureYaml)
	fmt.Println("Clojure Setup Complete!")
}
func haskellInit() {
	initBase(haskellDockerfile, haskellYaml)
	fmt.Println("Haskell Setup Complete!")
}
func dotNetInit() {
	initBase(dotNetDockerfile, dotNetYaml)
	fmt.Println(".NET Setup Complete!")
}
func zshInit() {
	initBase(zshDockerfile, zshYaml)
	fmt.Println("ZSh Setup Complete!")
}
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
