package main

import (
	"fmt"
	"resume/internal"
)

func main() {
	// contents, err := internal.ParseTomlResumeContents("./contents.toml")
	contents, err := internal.ParseLuaResumeContents("./contents.lua")
	if err != nil {
		panic(err)
	}

	templatePath := "./templates/template.typ"
	outputPath := "./resume.typ" // TODO: name resume output file from parsed name in contents
	if err := internal.PopulateTemplate(contents, templatePath, outputPath); err != nil {
		panic(err)
	}

	fmt.Println(contents)
}
