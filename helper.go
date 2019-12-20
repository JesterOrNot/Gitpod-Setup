package main

import (
	"fmt"
	"os"
)

func juliaInit() {
	initBase(juliaDockerFile, juliaYaml)
}
func nimInit() {
	initBase(nimDockerFile, nimYaml)
}
func hyInit() {
	initBase(hyDockerfile, hyYaml)
}
func clojureInit() {
	initBase(clojureDockerfile, clojureYaml)
}
func haskellInit() {
	initBase(haskellDockerfile, haskellYaml)
}
func dotNetInit() {
	initBase(dotNetDockerfile, dotNetYaml)
}
func zshInit() {
	initBase(zshDockerfile, zshYaml)
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
