package main

import (
	_ "embed"
	"flag"
	"fmt"
	"os/exec"
	"strings"

	"github.com/ptdewey/resume-builder/internal"
)

// TODO: maybe make the template file a runtime dependency (passed in with a flag)
//
//go:embed templates/template.typ
var templateFile []byte

func main() {
	contentsPath := flag.String("i", "examples/contents.lua", "-i=contents.lua")
	noCompileFlag := flag.Bool("no-compile", false, "Do not compile result typ file.")
	flag.Parse()

	if *contentsPath == "examples/contents.lua" {
		fmt.Println("No input file was provided, defaulting to './examples/contents.lua'. Use 'resume-builder -i <your-resume.lua>'")
	}

	fmt.Println("Parsing resume contents...")
	contents, err := internal.ParseLuaResumeContents(*contentsPath)
	if err != nil {
		panic(err)
	}

	fmt.Println("Populating template file...")
	outputPath := strings.ReplaceAll(contents.Personal.Name, " ", "") + "-Resume.typ"
	if err := internal.PopulateTemplate(contents, templateFile, outputPath); err != nil {
		panic(err)
	}

	if !*noCompileFlag {
		fmt.Println("Compiling with Typst...")
		cmd := exec.Command("typst", "compile", outputPath)
		if err := cmd.Run(); err != nil {
			panic(err)
		}
	} else {
		fmt.Println("Skipping typst compilation.")
	}

	fmt.Println("Done!")
}
