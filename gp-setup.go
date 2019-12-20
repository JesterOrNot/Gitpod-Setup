package main

import (
	"fmt"
	"os"

	"github.com/manifoldco/promptui"
)
var juliaDockerFile string =`FROM gitpod/workspace-full\n\n

USER gitpod

# Install Julia
RUN sudo apt-get update \
	&& sudo apt-get install -y \
		build-essential \
		libatomic1 \
		python \
		gfortran \
		perl \
		wget \
		m4 \
		cmake \
		pkg-config \
		julia \
	&& sudo rm -rf /var/lib/apt/lists/*\n\n

# Give control back to Gitpod Layer\n
USER root`
var juliaYaml string = `image:
  file: .gitpod.Dockerfile

tasks:
- command: julia --version

vscode:
  extensions:
    - julialang.language-julia@0.12.3:lgRyBd8rjwUpMGG0C5GAig==
`
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
				Items: []string{"Julia", "Nim", "Back", "Never Mind"},
			}
			_, result2, err2 := langPrompt.Run()
			if err2 != nil {
				exit()
			}
			switch result2 {
				case "Julia":
					juliaInit()
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
	gitpodDockerfile, gitpodDockerfileErr := os.Create(".gitpod.Dockerfile")
	gitpodYaml, gitpodYamlErr := os.Create(".gitpod.yml")
	if gitpodDockerfileErr != nil {
		fmt.Println("\033[1;31mAlready exists\033[0m")
	}
	if gitpodYamlErr != nil {
		fmt.Println("\033[1;31mAlready exists\033[0m")
	}
	gitpodDockerfile.WriteString(juliaDockerFile)
	gitpodYaml.WriteString(juliaYaml)

}
func exit() {
	fmt.Println("Ok, bye!")
	os.Exit(0)
}
func main() {
	initInteractive()
}
