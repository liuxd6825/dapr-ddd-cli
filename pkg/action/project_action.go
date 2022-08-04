package action

import (
	"fmt"
	"github.com/liuxd6825/dapr-ddd-cli/pkg/builds/build_cmd"
	"github.com/liuxd6825/dapr-ddd-cli/pkg/builds/build_config"
	"github.com/liuxd6825/dapr-ddd-cli/pkg/builds/build_docker"
	"github.com/liuxd6825/dapr-ddd-cli/pkg/builds/build_k8s"
	"github.com/liuxd6825/dapr-ddd-cli/pkg/builds/build_other"
	"github.com/liuxd6825/dapr-ddd-cli/pkg/builds/build_pkg/cmd-service/cmd_application"
	"github.com/liuxd6825/dapr-ddd-cli/pkg/builds/build_pkg/cmd-service/cmd_domain"
	"github.com/liuxd6825/dapr-ddd-cli/pkg/builds/build_pkg/cmd-service/cmd_infrastructure"
	"github.com/liuxd6825/dapr-ddd-cli/pkg/builds/build_pkg/cmd-service/cmd_userinterface"
	"github.com/liuxd6825/dapr-ddd-cli/pkg/builds/build_pkg/query-service/query_application"
	"github.com/liuxd6825/dapr-ddd-cli/pkg/builds/build_pkg/query-service/query_domain"
	"github.com/liuxd6825/dapr-ddd-cli/pkg/builds/build_pkg/query-service/query_infrastructure"
	"github.com/liuxd6825/dapr-ddd-cli/pkg/builds/build_pkg/query-service/query_userinterface"
	"github.com/liuxd6825/dapr-ddd-cli/pkg/config"
	"github.com/urfave/cli/v2"
	"strings"
)

func ProjectAction(c *cli.Context) error {
	fmt.Println("init project added task: ", c.Args())
	fmt.Println("flagName: " + strings.Join(c.FlagNames(), ","))

	flag, err := NewCommonFlag(c)
	if err != nil {
		return err
	}
	return projectBuild(flag.modelPath, flag.lang, flag.outPath)
}

func projectBuild(modelPath string, lang string, outPath string) error {

	cfg, err := config.NewConfigWithDir(modelPath, lang)
	if err != nil {
		return err
	}

	buildMain := build_cmd.NewMainLayer(cfg, outPath+"/cmd")
	if err := buildMain.Builds(); err != nil {
		panic(err)
	}

	buildConfig := build_config.NewBuildConfigLayer(cfg, outPath+"/config")
	if err := buildConfig.Builds(); err != nil {
		panic(err)
	}

	buildDocker := build_docker.NewBuildDockerLayer(cfg, outPath+"/docker")
	if err := buildDocker.Builds(); err != nil {
		panic(err)
	}

	buildK8s := build_k8s.NewBuildK8sLayer(cfg, outPath+"/k8s")
	if err := buildK8s.Builds(); err != nil {
		panic(err)
	}

	buildMakefile := build_other.NewBuildMakefile(cfg, outPath)
	if err := buildMakefile.Builds(); err != nil {
		panic(err)
	}

	cmdDir := fmt.Sprintf("%s/pkg/cmd-service", outPath)
	for _, agg := range cfg.Aggregates {

		buildDomain := cmd_domain.NewBuildDomainLayer(cfg, agg, cmdDir+"/domain")
		if err := buildDomain.Builds(); err != nil {
			panic(err)
		}

		buildInfra := cmd_infrastructure.NewBuildInfrastructureLayer(cfg, agg, cmdDir+"/infrastructure")
		if err := buildInfra.Builds(); err != nil {
			panic(err)
		}

		buildApplication := cmd_application.NewBuildApplicationLayer(cfg, agg, cmdDir+"/application")
		if err := buildApplication.Builds(); err != nil {
			panic(err)
		}

		buildUserInterface := cmd_userinterface.NewBuildRestControllerLayer(cfg, agg, cmdDir+"/userinterface")
		if err := buildUserInterface.Builds(); err != nil {
			panic(err)
		}

	}

	queryDir := fmt.Sprintf("%s/pkg/query-service", outPath)
	for _, agg := range cfg.Aggregates {

		buildDomain := query_domain.NewBuildDomainLayer(cfg, agg, queryDir+"/domain")
		if err := buildDomain.Builds(); err != nil {
			panic(err)
		}

		buildInfra := query_infrastructure.NewBuildInfrastructureLayer(cfg, agg, queryDir+"/infrastructure")
		if err := buildInfra.Builds(); err != nil {
			panic(err)
		}

		buildApplication := query_application.NewBuildApplicationLayer(cfg, agg, queryDir+"/application")
		if err := buildApplication.Builds(); err != nil {
			panic(err)
		}

		buildUserInterface := query_userinterface.NewBuildUserInterfaceLayer(cfg, agg, queryDir+"/userinterface")
		if err := buildUserInterface.Builds(); err != nil {
			panic(err)
		}

	}
	println("init project build success.")
	return nil
}
