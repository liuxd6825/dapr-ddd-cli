package query_application

import (
	"fmt"
	"github.com/liuxd6825/dapr-ddd-cli/pkg/builds"
	"github.com/liuxd6825/dapr-ddd-cli/pkg/config"
	"github.com/liuxd6825/dapr-ddd-cli/pkg/utils"
)

type BuildApplicationLayer struct {
	builds.BaseBuild
	aggregate *config.Aggregate
	outDir    string
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
	res.initQueryAppService()
	res.initQueryHandlerAggregate()
	res.initQueryHandlerEntities()
	return res
}

func (b *BuildApplicationLayer) initQueryAppService() {
	outFile := fmt.Sprintf("%s/internals/%s/service/%s_query_appservice.go", b.outDir, b.aggregate.FileName(), b.aggregate.FileName())
	build := NewBuildAppService(b.BaseBuild, nil, utils.ToLower(outFile))
	b.AddBuild(build)
	for _, item := range b.aggregate.Entities {
		outFile := fmt.Sprintf("%s/internals/%s/service/%s_query_appservice.go", b.outDir, b.aggregate.FileName(), item.FileName())
		build := NewBuildAppService(b.BaseBuild, item, utils.ToLower(outFile))
		b.AddBuild(build)
	}
}

func (b *BuildApplicationLayer) initQueryHandlerEntities() {
	for _, item := range b.aggregate.Entities {
		outFile := fmt.Sprintf("%s/internals/%s/handler/%s_query_handler.go", b.outDir, b.Aggregate.FileName(), item.FileName())
		build := NewBuildQueryHandlerEntity(b.BaseBuild, b.aggregate, item, utils.ToLower(outFile))
		b.AddBuild(build)
	}
}

func (b *BuildApplicationLayer) initQueryHandlerAggregate() {
	outFile := fmt.Sprintf("%s/internals/%s/handler/%s_query_handler.go", b.outDir, b.Aggregate.FileName(), b.Aggregate.FileName())
	build := NewBuildQueryHandler(b.BaseBuild, b.aggregate, utils.ToLower(outFile))
	b.AddBuild(build)
}
