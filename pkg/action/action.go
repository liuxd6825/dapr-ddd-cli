package action

import (
	"fmt"
	"github.com/liuxd6825/dapr-ddd-cli/pkg/utils"
	"github.com/urfave/cli/v2"
	"strings"
)

func Action(c *cli.Context) error {
	fmt.Println("init project added task: ", c.Args())
	fmt.Println("flagName: " + strings.Join(c.FlagNames(), ","))

	modelPath := c.String("model")
	if err := utils.ValidEmptyStr(modelPath, "-model 不能为空"); err != nil {
		return err
	}

	lang := c.String("lang")
	if err := utils.ValidEmptyStr(lang, "-lang 不能为空"); err != nil {
		return err
	}

	outPath := c.String("out")
	if err := utils.ValidEmptyStr(outPath, "-out 不能为空"); err != nil {
		return err
	}

	s := c.String("aggregate")
	aggregates := split(s, ",")

	s = c.String("layer")
	layers := split(s, ",")

	s = c.String("service")
	services := split(s, ",")

	err := BuildProject(modelPath, lang, outPath, aggregates, layers, services)
	return err
}

func split(s string, sep string) []string {
	value := strings.Replace(s, " ", "", 0)
	value = strings.ToLower(value)
	if value == "" {
		return []string{}
	}
	return strings.Split(s, sep)
}
