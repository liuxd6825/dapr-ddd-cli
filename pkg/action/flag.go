package action

import (
	"github.com/liuxd6825/dapr-ddd-cli/pkg/utils"
	"github.com/urfave/cli/v2"
	"strings"
)

type CommonFlag struct {
	outPath   string
	modelPath string
	lang      string
}

func NewCommonFlag(c *cli.Context) (*CommonFlag, error) {
	var (
		modelPath string
		lang      string
		outPath   string
	)

	if s, err := getModelFlag(c); err != nil {
		return nil, err
	} else {
		modelPath = s
	}
	if s, err := getLangFlag(c); err != nil {
		return nil, err
	} else {
		lang = s
	}
	if s, err := getOutPath(c); err != nil {
		return nil, err
	} else {
		outPath = s
	}
	return &CommonFlag{
		modelPath: modelPath,
		lang:      lang,
		outPath:   outPath,
	}, nil
}

func getModelFlag(c *cli.Context) (string, error) {
	modelPath := c.String("model")
	if err := utils.ValidEmptyStr(modelPath, "-model 不能为空"); err != nil {
		return "", err
	}
	return modelPath, nil
}

func getLangFlag(c *cli.Context) (string, error) {
	lang := c.String("lang")
	if err := utils.ValidEmptyStr(lang, "-lang 不能为空"); err != nil {
		return "", err
	}
	return lang, nil
}

func getOutPath(c *cli.Context) (string, error) {
	lang := c.String("out")
	if err := utils.ValidEmptyStr(lang, "-out 不能为空"); err != nil {
		return "", err
	}
	return lang, nil
}

func getNamesFlag(c *cli.Context, flagName string) ([]string, error) {
	s := c.String(flagName)
	if err := utils.ValidEmptyStr(s, "-"+flagName+" 不能为空"); err != nil {
		return []string{}, err
	}
	sl := split(s, ",")
	return sl, nil
}

func split(s string, sep string) []string {
	value := strings.Replace(s, " ", "", 0)
	value = strings.ToLower(value)
	if value == "" {
		return []string{}
	}
	return strings.Split(s, sep)
}
