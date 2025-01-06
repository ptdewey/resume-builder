package main

import (
	"fmt"
	"resume/internal"
)

func main() {
	// contents, err := internal.ParseResumeContents("./contents.toml")
	contents, err := internal.ParseLuaResumeContents("./contents.lua")
	if err != nil {
		panic(err)
	}

	fmt.Println(contents)
	fmt.Println(contents.Work)
}
