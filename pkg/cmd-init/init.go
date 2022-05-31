package cmd_init

import (
	"fmt"
	"github.com/liuxd6825/dapr-ddd-cli/pkg/cmd-init/builds/build_cmd"
	"github.com/liuxd6825/dapr-ddd-cli/pkg/cmd-init/builds/build_config"
	"github.com/liuxd6825/dapr-ddd-cli/pkg/cmd-init/builds/build_docker"
	"github.com/liuxd6825/dapr-ddd-cli/pkg/cmd-init/builds/build_k8s"
	"github.com/liuxd6825/dapr-ddd-cli/pkg/cmd-init/builds/build_other"
	"github.com/liuxd6825/dapr-ddd-cli/pkg/cmd-init/builds/build_pkg/cmd-service/cmd_application"
	"github.com/liuxd6825/dapr-ddd-cli/pkg/cmd-init/builds/build_pkg/cmd-service/cmd_domain"
	"github.com/liuxd6825/dapr-ddd-cli/pkg/cmd-init/builds/build_pkg/cmd-service/cmd_infrastructure"
	"github.com/liuxd6825/dapr-ddd-cli/pkg/cmd-init/builds/build_pkg/cmd-service/cmd_userinterface"
	"github.com/liuxd6825/dapr-ddd-cli/pkg/cmd-init/builds/build_pkg/query-service/query_application"
	"github.com/liuxd6825/dapr-ddd-cli/pkg/cmd-init/builds/build_pkg/query-service/query_domain"
	"github.com/liuxd6825/dapr-ddd-cli/pkg/cmd-init/builds/build_pkg/query-service/query_infrastructure"
	"github.com/liuxd6825/dapr-ddd-cli/pkg/cmd-init/builds/build_pkg/query-service/query_userinterface"
	"github.com/liuxd6825/dapr-ddd-cli/pkg/config"
	"github.com/liuxd6825/dapr-ddd-cli/pkg/utils"
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

	buildMain := build_cmd.NewMainLayer(cfg, out+"/cmd")
	if err := buildMain.Build(); err != nil {
		panic(err)
	}

	buildConfig := build_config.NewBuildConfigLayer(cfg, out+"/config")
	if err := buildConfig.Build(); err != nil {
		panic(err)
	}

	buildDocker := build_docker.NewBuildDockerLayer(cfg, out+"/docker")
	if err := buildDocker.Build(); err != nil {
		panic(err)
	}

	buildK8s := build_k8s.NewBuildK8sLayer(cfg, out+"/k8s")
	if err := buildK8s.Build(); err != nil {
		panic(err)
	}

	buildMakefile := build_other.NewBuildMakefile(cfg, out)
	if err := buildMakefile.Build(); err != nil {
		panic(err)
	}

	isBuild := true
	if isBuild {
		cmdDir := fmt.Sprintf("%s/pkg/cmd-service", out)
		for _, agg := range cfg.Aggregates {
			buildDomain := cmd_domain.NewBuildDomainLayer(cfg, agg, cmdDir+"/domain")
			if err := buildDomain.Build(); err != nil {
				panic(err)
			}

			buildInfr := cmd_infrastructure.NewBuildInfrastructureLayer(cfg, agg, cmdDir+"/infrastructure")
			if err := buildInfr.Build(); err != nil {
				panic(err)
			}

			buildApplication := cmd_application.NewBuildApplicationLayer(cfg, agg, cmdDir+"/application")
			if err := buildApplication.Build(); err != nil {
				panic(err)
			}

			buildUserInterface := cmd_userinterface.NewBuildRestControllerLayer(cfg, agg, cmdDir+"/userinterface")
			if err := buildUserInterface.Build(); err != nil {
				panic(err)
			}
		}
	}
	if isBuild {
		queryDir := fmt.Sprintf("%s/pkg/query-service", out)
		for _, agg := range cfg.Aggregates {
			buildDomain := query_domain.NewBuildDomainLayer(cfg, agg, queryDir+"/domain")
			if err := buildDomain.Build(); err != nil {
				panic(err)
			}

			buildInfrastructure := query_infrastructure.NewBuildInfrastructureLayer(cfg, agg, queryDir+"/infrastructure")
			if err := buildInfrastructure.Build(); err != nil {
				panic(err)
			}

			buildApplication := query_application.NewBuildApplicationLayer(cfg, agg, queryDir+"/application")
			if err := buildApplication.Build(); err != nil {
				panic(err)
			}

			buildUserInterface := query_userinterface.NewBuildUserInterfaceLayer(cfg, agg, queryDir+"/userinterface")
			if err := buildUserInterface.Build(); err != nil {
				panic(err)
			}
		}
	}
	println("build success.")
	return nil
}
