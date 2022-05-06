package cmd_init

import (
	"fmt"
	"github.com/dapr/dapr-ddd-cli/pkg/resource"
	"github.com/dapr/dapr-ddd-cli/pkg/utils"
	"github.com/urfave/cli/v2"
	"strings"
)

var templatePath string
var outPath string

func Acton(c *cli.Context) error {
	fmt.Println("init project added task: ", c.Args())
	fmt.Println("flagName: " + strings.Join(c.FlagNames(), ","))
	fmt.Printf("-name : %s\r\n", c.String("name"))

	name := c.String("name")
	if err := utils.ValidEmptyStr(name, "-name 不能为空"); err != nil {
		return err
	}

	service := c.String("service")
	if err := utils.ValidEmptyStr(service, "-service 不能为空"); err != nil {
		return err
	}

	lang := c.String("lang")
	if err := utils.ValidEmptyStr(service, "-lang 不能为空"); err != nil {
		return err
	}

	out := c.String("out")
	if err := utils.ValidEmptyStr(service, "-out 不能为空"); err != nil {
		return err
	}

	err := initProject(lang, out)
	return err
}

func initProject(lang string, out string) error {
	tmplDir := fmt.Sprintf("static/tmpl/%s/init", lang)
	return run(tmplDir, out)
}

func run(tmplDir string, out string) error {
	dirs, err := resource.Local().ReadDir(tmplDir)
	if err != nil {
		return err
	}
	for _, dir := range dirs {
		if dir.Name() == "./" {
			continue
		}
		t := newTemplate(dir, tmplDir, out)
		if err := t.action(); err != nil {
			println(err)
			return err
		}
		if t.isDir() {
			outPath := out + "/" + dir.Name()
			tmplPath := tmplDir + "/" + t.getName()
			if err := run(tmplPath, outPath); err != nil {
				return err
			}
		}
	}
	return nil
}
