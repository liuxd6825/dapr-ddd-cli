package utils

import (
	"fmt"
	"github.com/liuxd6825/dapr-ddd-cli/pkg/resource"
	"os"
	"strings"
	"text/template"
)

func RunTemplate(tmplFile string, data interface{}, outFile string) error {
	println("building: " + tmplFile)
	bytes, err := resource.Local().ReadFile(tmplFile)
	if err != nil {
		return err
	}
	tmpl := template.New(tmplFile).Funcs(template.FuncMap{
		"runTemplate": RunTemplate,
		"firstUpper":  FirstUpper,
		"firstLower":  FirstLower,
		"getData":     GetData,
		"add":         Add,
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
		dirName := GetDirName(outFile)
		if err = os.MkdirAll(dirName, os.ModePerm); err != nil {
			return err
		}

		if w, err = os.Create(outFile); err != nil {
			return err
		}
	}

	if err = tmpl.Execute(w, data); err != nil {
		return err
	} else {
		println("success: " + outFile)
	}
	return nil
}

func GetData(dataMap map[string]interface{}) interface{} {
	if res, ok := dataMap["data"]; ok {
		return res
	}
	return nil
}

func GetDirName(fileName string) string {
	last := strings.LastIndex(fileName, "/")
	if last > -1 {
		return fileName[0:last]
	}
	last = strings.LastIndex(fileName, "\\")
	if last > -1 {
		return fileName[0:last]
	}
	return ""
}
