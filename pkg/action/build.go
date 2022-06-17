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
	"strings"
)

func BuildProject(modelPath string, lang string, out string, aggregates []string, layers []string, services []string) error {

	cfg, err := config.NewConfigWithDir(modelPath, lang)
	if err != nil {
		return err
	}

	if isLayers("other", layers) {
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

	}

	if isServices("cmd", services) {
		cmdDir := fmt.Sprintf("%s/pkg/cmd-service", out)
		for _, agg := range cfg.Aggregates {
			if isUpdateAggregate(agg.Name, aggregates) {
				if isLayers("domain", layers) {
					buildDomain := cmd_domain.NewBuildDomainLayer(cfg, agg, cmdDir+"/domain")
					if err := buildDomain.Build(); err != nil {
						panic(err)
					}
				}
				if isLayers("infra", layers) {
					buildInfra := cmd_infrastructure.NewBuildInfrastructureLayer(cfg, agg, cmdDir+"/infrastructure")
					if err := buildInfra.Build(); err != nil {
						panic(err)
					}
				}

				if isLayers("app", layers) {
					buildApplication := cmd_application.NewBuildApplicationLayer(cfg, agg, cmdDir+"/application")
					if err := buildApplication.Build(); err != nil {
						panic(err)
					}
				}

				if isLayers("ui", layers) {
					buildUserInterface := cmd_userinterface.NewBuildRestControllerLayer(cfg, agg, cmdDir+"/userinterface")
					if err := buildUserInterface.Build(); err != nil {
						panic(err)
					}
				}
			}
		}
	}

	if isServices("query", services) {
		queryDir := fmt.Sprintf("%s/pkg/query-service", out)
		for _, agg := range cfg.Aggregates {
			if isUpdateAggregate(agg.Name, aggregates) {
				if isLayers("domain", layers) {
					buildDomain := query_domain.NewBuildDomainLayer(cfg, agg, queryDir+"/domain")
					if err := buildDomain.Build(); err != nil {
						panic(err)
					}
				}

				if isLayers("infra", layers) {
					buildInfra := query_infrastructure.NewBuildInfrastructureLayer(cfg, agg, queryDir+"/infrastructure")
					if err := buildInfra.Build(); err != nil {
						panic(err)
					}
				}

				if isLayers("app", layers) {
					buildApplication := query_application.NewBuildApplicationLayer(cfg, agg, queryDir+"/application")
					if err := buildApplication.Build(); err != nil {
						panic(err)
					}
				}

				if isLayers("ui", layers) {
					buildUserInterface := query_userinterface.NewBuildUserInterfaceLayer(cfg, agg, queryDir+"/userinterface")
					if err := buildUserInterface.Build(); err != nil {
						panic(err)
					}
				}
			}
		}
	}
	println("build success.")
	return nil
}

func isUpdateAggregate(aggName string, aggNames []string) bool {
	if len(aggNames) == 0 {
		return true
	}
	name := strings.ToLower(aggName)
	for _, item := range aggNames {
		if strings.ToLower(item) == name {
			return true
		}
	}
	return false
}

func isLayers(layer string, layers []string) bool {
	if len(layers) == 0 {
		return true
	}
	name := strings.ToLower(layer)
	for _, item := range layers {
		if strings.ToLower(item) == name {
			return true
		}
	}
	return false
}

func isServices(service string, services []string) bool {
	if len(services) == 0 {
		return true
	}
	name := strings.ToLower(service)
	for _, item := range services {
		if strings.ToLower(item) == name {
			return true
		}
	}
	return false
}
