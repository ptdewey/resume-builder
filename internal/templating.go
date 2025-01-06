package internal

import (
	"os"
	"text/template"
)

func FillTemplate(templatePath string, outputPath string) error {
	bytes, err := os.ReadFile(templatePath)
	if err != nil {
		return err
	}

	tmpl, err := template.New("resume").Parse(string(bytes))
	if err != nil {
		return err
	}

	f, err := os.Create(outputPath)
	if err != nil {
		return err
	}

	// TODO: fill with resume contents
	if err := tmpl.Execute(f, []string{"text"}); err != nil {
		return err
	}

	return nil
}
