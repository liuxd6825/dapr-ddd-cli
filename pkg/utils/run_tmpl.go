package utils

import (
	"fmt"
	"github.com/dapr/dapr-ddd-cli/pkg/resource"
	"os"
	"text/template"
)

func RunTemplate(tmplFile string, data interface{}, outFile string) error {
	if len(outFile) > 0 {
		if err := os.MkdirAll(outFile, os.ModePerm); err != nil {
			return err
		}
	}

	bytes, err := resource.Local().ReadFile(tmplFile)
	if err != nil {
		return err
	}
	tmpl := template.New(tmplFile).Funcs(template.FuncMap{
		"runTemplate": RunTemplate,
		"firstUpper":  FirstUpper,
		"firstLower":  FirstLower,
	})
	tmpl, err = tmpl.Parse(string(bytes))
	if err != nil {
		fmt.Println("template.Parse():", err)
		return err
	}

	w := os.Stdout
	defer func() {
		if w != os.Stdout {
			_ = w.Close()
		}
	}()

	if len(outFile) > 0 {
		w, err = os.Create(outFile)
		if err != nil {
			return err
		}
	}
	if err := tmpl.Execute(w, data); err != nil {
		println("")
		return err
	}
	return nil
}
