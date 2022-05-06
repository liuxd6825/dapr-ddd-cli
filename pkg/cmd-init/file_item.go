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

type fileItem struct {
	sourceDir  string
	sourceFile string
	outDir     string
	outFile    string
	dirEntry   fs.DirEntry
}

func newTemplate(dirEntry fs.DirEntry, sourceDirName, outDirName string) *fileItem {
	sourceDir := strings.ReplaceAll(sourceDirName, "//", "/")
	outDir := strings.ReplaceAll(outDirName, "//", "/")

	outFile := outDir + "/" + dirEntry.Name()
	if strings.HasSuffix(outFile, ".tpl") {
		i := strings.LastIndex(outFile, ".tpl")
		outFile = outFile[:i]
	}
	outFile = strings.ReplaceAll(outFile, "//", "/")

	return &fileItem{
		sourceDir:  sourceDir,
		sourceFile: sourceDirName + "/" + dirEntry.Name(),
		outDir:     outDir,
		outFile:    outFile,
		dirEntry:   dirEntry,
	}
}

func (f *fileItem) getFileName() string {
	return f.sourceFile
}

func (f *fileItem) getName() string {
	return f.dirEntry.Name()
}

func (f *fileItem) action() error {
	if f.isDir() {
		return f.createDir()
	}
	return f.createFile()
}

func (f *fileItem) isDir() bool {
	return f.dirEntry.IsDir()
}

func (f *fileItem) createDir() error {
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

func (f *fileItem) createFile() error {
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
	tmpl, err := template.New(f.getName()).Parse(string(bytes))
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
