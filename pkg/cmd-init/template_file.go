package cmd_init

import (
	"fmt"
	"github.com/dapr/dapr-ddd-cli/pkg/resource"
	"io/fs"
	"os"
	"strings"
	"text/template"
)

var isDebug = false

type TemplateFile struct {
	sourceDir  string
	sourceFile string
	outDir     string
	outFile    string
	dirEntry   fs.DirEntry
}

func newTemplateFile(dirEntry fs.DirEntry, sourceDirName, outDirName string) *TemplateFile {
	sourceDir := strings.ReplaceAll(sourceDirName, "//", "/")
	outDir := strings.ReplaceAll(outDirName, "//", "/")

	outFile := outDir + "/" + dirEntry.Name()
	if strings.HasSuffix(outFile, ".tpl") {
		i := strings.LastIndex(outFile, ".tpl")
		outFile = outFile[:i]
	}
	outFile = strings.ReplaceAll(outFile, "//", "/")

	return &TemplateFile{
		sourceDir:  sourceDir,
		sourceFile: sourceDirName + "/" + dirEntry.Name(),
		outDir:     outDir,
		outFile:    outFile,
		dirEntry:   dirEntry,
	}
}

func (f *TemplateFile) getFileName() string {
	return f.sourceFile
}

func (f *TemplateFile) getName() string {
	return f.dirEntry.Name()
}

func (f *TemplateFile) action() error {
	if f.isDir() {
		return f.createDir()
	}
	return f.createFile()
}

func (f *TemplateFile) isDir() bool {
	return f.dirEntry.IsDir()
}

func (f *TemplateFile) createDir() error {
	fmt.Printf("DIR  %s %s \r\n", f.dirEntry.Name(), f.outDir)
	if !f.isDir() {
		return nil
	}
	if !isDebug && f.outDir != "./" {
		if err := os.MkdirAll(f.outDir, os.ModePerm); err != nil {
			return err
		}
	}
	return nil
}

func (f *TemplateFile) createFile() error {
	if f.isDir() {
		return nil
	}
	if err := os.MkdirAll(f.outDir, os.ModePerm); err != nil {
		return err
	}

	bytes, err := resource.Local().ReadFile(f.sourceFile)
	if err != nil {
		return err
	}
	
	funcs := map[string]any{
		"ToLower": ToLower,
		"ToUpper": ToUpper,
	}

	tmpl, err := template.New(f.getName()).Funcs(funcs).Parse(string(bytes))

	if err != nil {
		fmt.Println("template.Parse():", err)
		return err
	}

	w := os.Stdout
	defer w.Close()

	println("OutFile: " + f.outFile)
	if !isDebug {
		w, err = os.Create(f.outFile)
		if err != nil {
			return err
		}
	}
	if err := tmpl.Execute(w, "小明"); err != nil {
		return err
	}
	return nil
}

func ToLower(str string) string {
	return strings.ToLower(str)
}

func ToUpper(str string) string {
	return strings.ToUpper(str)
}
