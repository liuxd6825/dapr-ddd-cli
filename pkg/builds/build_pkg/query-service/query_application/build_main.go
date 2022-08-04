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

	res.initAppServiceAggregate()
	res.initAppServiceEntities()

	res.initQueryHandlerAggregate()
	res.initQueryHandlerEntities()

	return res
}

func (b *BuildApplicationLayer) initAppServiceAggregate() {
	outFile := fmt.Sprintf("%s/internals/%s/service/%s_query_appservice.go", b.outDir, b.aggregate.FileName(), b.aggregate.FileName())
	build := NewBuildAppServiceAggregate(b.BaseBuild, b.aggregate, utils.ToLower(outFile))
	b.AddBuild(build)
}

func (b *BuildApplicationLayer) initAppServiceEntities() {
	for _, item := range b.aggregate.Entities {
		outFile := fmt.Sprintf("%s/internals/%s/service/%s_query_appservice.go", b.outDir, b.aggregate.FileName(), item.FileName())
		build := NewBuildRestControllerEntity(b.BaseBuild, item, utils.ToLower(outFile))
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
