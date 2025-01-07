package main

import (
	"fmt"
	"os"
	"os/exec"
	"resume-builder/internal"
	"strings"
)

func main() {
	var contentsPath string
	if len(os.Args) > 1 {
		contentsPath = os.Args[1]
	} else {
		contentsPath = "./examples/contents.lua"
	}

	fmt.Println("Parsing resume contents...")
	contents, err := internal.ParseLuaResumeContents(contentsPath)
	if err != nil {
		panic(err)
	}

	fmt.Println("Populating template file...")
	templatePath := "./templates/template.typ"
	outputPath := strings.ReplaceAll(contents.Personal.Name, " ", "") + "-Resume.typ"
	if err := internal.PopulateTemplate(contents, templatePath, outputPath); err != nil {
		panic(err)
	}

	fmt.Println("Compiling with Typst...")
	cmd := exec.Command("typst", "compile", outputPath)
	if err := cmd.Run(); err != nil {
		panic(err)
	}

	// TODO: cleanup?

	fmt.Println("Done!")
}
