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

type BuildType int

const (
	ProjectBuildType BuildType = iota
	AggregateBuildType
	ServiceBuildType
	LayerBuildType
)

func BuildProject(modelPath string, lang string, out string, aggregates []string, layers []string, services []string, buildType BuildType) error {

	cfg, err := config.NewConfigWithDir(modelPath, lang)
	if err != nil {
		return err
	}

	if isLayers("other", layers, buildType) {
		buildMain := build_cmd.NewMainLayer(cfg, out+"/cmd")
		if err := buildMain.Builds(); err != nil {
			panic(err)
		}

		buildConfig := build_config.NewBuildConfigLayer(cfg, out+"/config")
		if err := buildConfig.Builds(); err != nil {
			panic(err)
		}

		buildDocker := build_docker.NewBuildDockerLayer(cfg, out+"/docker")
		if err := buildDocker.Builds(); err != nil {
			panic(err)
		}

		buildK8s := build_k8s.NewBuildK8sLayer(cfg, out+"/k8s")
		if err := buildK8s.Builds(); err != nil {
			panic(err)
		}

		buildMakefile := build_other.NewBuildMakefile(cfg, out)
		if err := buildMakefile.Builds(); err != nil {
			panic(err)
		}

	}

	if isServices("cmd", services, buildType) {
		cmdDir := fmt.Sprintf("%s/pkg/cmd-service", out)
		for _, agg := range cfg.Aggregates {
			if isUpdateAggregate(agg.Name, aggregates, buildType) {

				if isLayers("domain", layers, buildType) {
					buildDomain := cmd_domain.NewBuildDomainLayer(cfg, agg, cmdDir+"/domain")
					if err := buildDomain.Builds(); err != nil {
						panic(err)
					}
				}

				if isLayers("infra", layers, buildType) {
					buildInfra := cmd_infrastructure.NewBuildInfrastructureLayer(cfg, agg, cmdDir+"/infrastructure")
					if err := buildInfra.Builds(); err != nil {
						panic(err)
					}
				}

				if isLayers("app", layers, buildType) {
					buildApplication := cmd_application.NewBuildApplicationLayer(cfg, agg, cmdDir+"/application")
					if err := buildApplication.Builds(); err != nil {
						panic(err)
					}
				}

				if isLayers("ui", layers, buildType) {
					buildUserInterface := cmd_userinterface.NewBuildRestControllerLayer(cfg, agg, cmdDir+"/userinterface")
					if err := buildUserInterface.Builds(); err != nil {
						panic(err)
					}
				}
			}
		}
	}

	if isServices("query", services, buildType) {
		queryDir := fmt.Sprintf("%s/pkg/query-service", out)
		for _, agg := range cfg.Aggregates {
			if isUpdateAggregate(agg.Name, aggregates, buildType) {
				if isLayers("domain", layers, buildType) {
					buildDomain := query_domain.NewBuildDomainLayer(cfg, agg, queryDir+"/domain")
					if err := buildDomain.Builds(); err != nil {
						panic(err)
					}
				}

				if isLayers("infra", layers, buildType) {
					buildInfra := query_infrastructure.NewBuildInfrastructureLayer(cfg, agg, queryDir+"/infrastructure")
					if err := buildInfra.Builds(); err != nil {
						panic(err)
					}
				}

				if isLayers("app", layers, buildType) {
					buildApplication := query_application.NewBuildApplicationLayer(cfg, agg, queryDir+"/application")
					if err := buildApplication.Builds(); err != nil {
						panic(err)
					}
				}

				if isLayers("ui", layers, buildType) {
					buildUserInterface := query_userinterface.NewBuildUserInterfaceLayer(cfg, agg, queryDir+"/userinterface")
					if err := buildUserInterface.Builds(); err != nil {
						panic(err)
					}
				}
			}
		}
	}
	println("build success.")
	return nil
}

func isUpdateAggregate(aggName string, aggNames []string, buildType BuildType) bool {
	switch buildType {
	case ProjectBuildType:
		return true
	case AggregateBuildType:
		if len(aggNames) == 0 {
			return true
		}
		name := strings.ToLower(aggName)
		for _, item := range aggNames {
			if strings.ToLower(item) == name {
				return true
			}
		}
	}
	return false
}

func isLayers(layer string, layers []string, buildType BuildType) bool {
	switch buildType {
	case ProjectBuildType:
		return true
	case LayerBuildType:
		if len(layers) == 0 {
			return true
		}
		name := strings.ToLower(layer)
		for _, item := range layers {
			if strings.ToLower(item) == name {
				return true
			}
		}
	}
	return false
}

func isServices(service string, services []string, buildType BuildType) bool {
	switch buildType {
	case ProjectBuildType:
		return true
	case AggregateBuildType:
		return true
	case ServiceBuildType:
		return true
	}
	if len(services) == 0 {
		return false
	}
	name := strings.ToLower(service)
	for _, item := range services {
		if strings.ToLower(item) == name {
			return true
		}
	}
	return false
}
