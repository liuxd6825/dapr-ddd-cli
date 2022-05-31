package query_infrastructure

import (
	"github.com/dapr/dapr-ddd-cli/pkg/cmd-init/builds"
	"github.com/dapr/dapr-ddd-cli/pkg/config"
)

type BuildRepositoryBase struct {
	builds.BaseBuild
	aggregate *config.Aggregate
}

func NewBuildRepositoryBase(base builds.BaseBuild, aggregate *config.Aggregate, outFile string) *BuildRepositoryBase {
	res := &BuildRepositoryBase{
		BaseBuild: base,
		aggregate: aggregate,
	}
	res.ValuesFunc = res.Values
	res.TmplFile = "static/tmpl/go/init/pkg/query-service/infrastructure/domain/repository/mongodb/base_repository.go.tpl"
	res.OutFile = outFile
	return res
}

func (b *BuildRepositoryBase) Values() map[string]interface{} {
	res := b.BaseBuild.Values()
	return res
}
