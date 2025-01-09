package internal

import (
	"os"
	"text/template"
)

func PopulateTemplate(contents ResumeContents, bytes []byte, outputPath string) error {
	tmpl, err := template.New("resume").Parse(string(bytes))
	if err != nil {
		return err
	}

	f, err := os.Create(outputPath)
	if err != nil {
		return err
	}

	if err := tmpl.Execute(f, contents); err != nil {
		return err
	}

	return nil
}
