package cmd_init

import (
	"fmt"
	"github.com/dapr/dapr-ddd-cli/pkg/cmd-init/builds/query-service/query_domain"
	"github.com/dapr/dapr-ddd-cli/pkg/cmd-init/builds/query-service/query_infrastructure"
	"github.com/dapr/dapr-ddd-cli/pkg/config"
	"github.com/dapr/dapr-ddd-cli/pkg/resource"
	"github.com/dapr/dapr-ddd-cli/pkg/utils"
	"github.com/urfave/cli/v2"
	"strings"
)

func Acton(c *cli.Context) error {
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

	err := initProject(modelPath, lang, outPath)
	return err
}

func initProject(modelPath string, lang string, out string) error {
	cfg, err := config.NewConfigWithDir(modelPath, lang)
	if err != nil {
		return err
	}

	/*	for _, agg := range cfg.Aggregates {
			buildDomain := cmd_domain.NewBuildInfrastructureLayer(cfg, agg, out+"/cmd-service/domain")
			if err := buildDomain.Build(); err != nil {
				panic(err)
			}

			buildApplication := cmd_application.NewBuildApplicationLayer(cfg, agg, out+"/cmd-service/application")
			if err := buildApplication.Build(); err != nil {
				panic(err)
			}

			buildUserInterface := cmd_userinterface.NewBuildRestControllerLayer(cfg, agg, out+"/cmd-service/userinterface")
			if err := buildUserInterface.Build(); err != nil {
				panic(err)
			}
		}
	*/

	for _, agg := range cfg.Aggregates {
		buildDomain := query_domain.NewBuildDomainLayer(cfg, agg, out+"/query-service/domain")
		if err := buildDomain.Build(); err != nil {
			panic(err)
		}

		buildInfrastructure := query_infrastructure.NewBuildInfrastructureLayer(cfg, agg, out+"/query-service/infrastructure")
		if err := buildInfrastructure.Build(); err != nil {
			panic(err)
		}
	}
	println("build success.")
	return nil
}

func run(tmplDir string, out string, cfg *config.Config) error {
	dirs, err := resource.Local().ReadDir(tmplDir)
	if err != nil {
		return err
	}

	for _, dir := range dirs {
		if dir.Name() == "./" {
			continue
		}
		t := newTemplateFile(dir, tmplDir, out)
		if err := t.action(); err != nil {
			println(err)
			return err
		}
		if t.isDir() {
			outPath := out + "/" + dir.Name()
			tmplPath := tmplDir + "/" + t.getName()
			if err := run(tmplPath, outPath, cfg); err != nil {
				return err
			}
		}
	}
	return nil
}
