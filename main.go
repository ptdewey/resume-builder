package main

import (
	_ "embed"
	"flag"
	"fmt"
	"os/exec"
	"strings"

	"github.com/ptdewey/rooibos/internal"
)

// TODO: maybe make the template file a runtime dependency (passed in with a flag)
//
//go:embed templates/template.typ
var templateFile []byte

func main() {
	contentsPath := flag.String("i", "examples/contents.lua", "-i=contents.lua")
	noCompileFlag := flag.Bool("no-compile", false, "--no-compile=true (default \"false\")")
	contentTags := flag.String("tags", "", "--tags=\"tag1 tag2 tag3\"")
	flag.Parse()

	if *contentsPath == "examples/contents.lua" {
		fmt.Println("No input file was provided, defaulting to './examples/contents.lua'. Use 'rooibos -i <your-resume.lua>'")
	}

	fmt.Println("Parsing resume contents...")
	contents, defaultTags, err := internal.ParseLuaResumeContents(*contentsPath)
	if err != nil {
		panic(err)
	}

	if *contentTags != "" {
		tags := strings.Split(*contentTags, " ")
		internal.SelectTags(&contents, tags)
	} else if defaultTags != nil && len(defaultTags) != 0 {
		internal.SelectTags(&contents, defaultTags)
	}

	fmt.Println("Populating template file...")
	// FIX: change to LastnameFirstName-Resume (multiple inputs in contents file)
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
