package cmd_application

import (
	"fmt"
	"github.com/liuxd6825/dapr-ddd-cli/pkg/cmd-init/builds"
	"github.com/liuxd6825/dapr-ddd-cli/pkg/config"
	"github.com/liuxd6825/dapr-ddd-cli/pkg/utils"
)

type BuildApplicationLayer struct {
	builds.BaseBuild
	aggregate                     *config.Aggregate
	outDir                        string
	buildCmdApplicationService    *BuildCmdApplicationService
	buildQueryAppServiceEntities  []*BuildQueryAppServiceEntity
	buildQueryAppServiceAggregate *BuildQueryAppServiceAggregate
}

func NewBuildApplicationLayer(cfg *config.Config, aggregate *config.Aggregate, outDir string) *BuildApplicationLayer {
	res := &BuildApplicationLayer{
		BaseBuild: builds.BaseBuild{
			Config:    cfg,
			Aggregate: aggregate,
		},
		aggregate: aggregate,
		outDir:    outDir,
	}
	res.init()
	return res
}

func (b *BuildApplicationLayer) init() {
	aggregateName := b.aggregate.SnakeName()

	outFile := fmt.Sprintf("%s/internals/service/%s_service/%s_command_appservice.go", b.outDir, aggregateName, aggregateName)
	b.buildCmdApplicationService = NewBuildCmdApplicationService(b.BaseBuild, b.aggregate, utils.ToLower(outFile))

	outFile = fmt.Sprintf("%s/internals/service/%s_service/%s_query_appservice.go", b.outDir, aggregateName, aggregateName)
	b.buildQueryAppServiceAggregate = NewBuildQueryAppServiceAggregate(b.BaseBuild, b.aggregate, utils.ToLower(outFile))

	b.buildQueryAppServiceEntities = []*BuildQueryAppServiceEntity{}
	for _, entity := range b.aggregate.Entities {
		outFile = fmt.Sprintf("%s/internals/service/%s_service/%s_query_appservice.go", b.outDir, aggregateName, entity.SnakeName())
		build := NewBuildQueryAppServiceEntity(b.BaseBuild, entity, outFile)
		b.buildQueryAppServiceEntities = append(b.buildQueryAppServiceEntities, build)
	}

}

func (b *BuildApplicationLayer) Build() error {
	var list []builds.Build
	list = append(list, b.buildCmdApplicationService)
	list = append(list, b.buildQueryAppServiceAggregate)

	entities := func() []builds.Build {
		var res []builds.Build
		for _, item := range b.buildQueryAppServiceEntities {
			res = append(res, item)
		}
		return res
	}
	list = append(list, entities()...)

	return b.DoBuild(list...)
}
