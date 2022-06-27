package action

import (
	"fmt"
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

func ServicesAction(c *cli.Context) error {
	fmt.Println("services Action added task: ", c.Args())
	fmt.Println("flagName: " + strings.Join(c.FlagNames(), ","))
	flag, err := NewCommonFlag(c)
	if err != nil {
		return err
	}
	services, err := getNamesFlag(c, "names")
	if err != nil {
		return err
	}
	return aggregateBuild(flag.modelPath, flag.lang, flag.outPath, services)
}

func serviceBuild(modelPath string, lang string, outPath string, services []string) error {

	cfg, err := config.NewConfigWithDir(modelPath, lang)
	if err != nil {
		return err
	}

	cmdDir := fmt.Sprintf("%s/pkg/cmd-service", outPath)
	for _, agg := range cfg.Aggregates {
		if hasName(services, agg.Name) {
			buildDomain := cmd_domain.NewBuildDomainLayer(cfg, agg, cmdDir+"/domain")
			if err := buildDomain.Build(); err != nil {
				panic(err)
			}

			buildInfra := cmd_infrastructure.NewBuildInfrastructureLayer(cfg, agg, cmdDir+"/infrastructure")
			if err := buildInfra.Build(); err != nil {
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

	queryDir := fmt.Sprintf("%s/pkg/query-service", outPath)
	for _, agg := range cfg.Aggregates {
		if hasName(services, agg.Name) {
			buildDomain := query_domain.NewBuildDomainLayer(cfg, agg, queryDir+"/domain")
			if err := buildDomain.Build(); err != nil {
				panic(err)
			}

			buildInfra := query_infrastructure.NewBuildInfrastructureLayer(cfg, agg, queryDir+"/infrastructure")
			if err := buildInfra.Build(); err != nil {
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

	println("init project build success.")
	return nil
}
