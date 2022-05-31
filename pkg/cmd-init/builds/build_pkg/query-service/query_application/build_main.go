package query_application

import (
	"fmt"
	"github.com/dapr/dapr-ddd-cli/pkg/cmd-init/builds"
	"github.com/dapr/dapr-ddd-cli/pkg/config"
	"github.com/dapr/dapr-ddd-cli/pkg/utils"
)

type BuildApplicationLayer struct {
	builds.BaseBuild
	aggregate *config.Aggregate
	outDir    string

	buildAppServiceAggregate *BuildAppServiceAggregate
	buildAppServiceEntities  []*BuildAppServiceEntity
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
	return res
}

func (b *BuildApplicationLayer) Build() error {
	var list []builds.Build

	// aggregate
	list = append(list, b.buildAppServiceAggregate)

	// entityObject
	buildAppServiceEntities := func() []builds.Build {
		var res []builds.Build
		for _, item := range b.buildAppServiceEntities {
			res = append(res, item)
		}
		return res
	}
	list = append(list, buildAppServiceEntities()...)

	return b.DoBuild(list...)
}

func (b *BuildApplicationLayer) initAppServiceAggregate() {
	outFile := fmt.Sprintf("%s/internales/service/%s_service/%s_query_appservice.go", b.outDir, b.aggregate.FileName(), b.aggregate.FileName())
	b.buildAppServiceAggregate = NewBuildAppServiceAggregate(b.BaseBuild, b.aggregate, utils.ToLower(outFile))
}

func (b *BuildApplicationLayer) initAppServiceEntities() {
	b.buildAppServiceEntities = []*BuildAppServiceEntity{}
	for _, item := range b.aggregate.Entities {
		outFile := fmt.Sprintf("%s/internales/service/%s_service/%s_query_appservice.go", b.outDir, b.aggregate.FileName(), item.FileName())
		build := NewBuildRestControllerEntity(b.BaseBuild, item, utils.ToLower(outFile))
		b.buildAppServiceEntities = append(b.buildAppServiceEntities, build)
	}
}
