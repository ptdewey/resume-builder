package main

import (
	"fmt"
	"resume/internal"
	"strings"
)

func main() {
	// TODO: take in input file as cli arg

	// contentsPath := "./contents.lua"
	contentsPath := "./templates/template.lua"
	contents, err := internal.ParseLuaResumeContents(contentsPath)
	if err != nil {
		panic(err)
	}

	templatePath := "./template.typ"
	outputPath := strings.ReplaceAll(contents.Personal.Name, " ", "") + "-Resume.typ"
	if err := internal.PopulateTemplate(contents, templatePath, outputPath); err != nil {
		panic(err)
	}

	fmt.Println(contents)
}
