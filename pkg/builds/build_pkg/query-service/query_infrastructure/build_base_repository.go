package query_infrastructure

import (
	"github.com/liuxd6825/dapr-ddd-cli/pkg/builds"
	"github.com/liuxd6825/dapr-ddd-cli/pkg/config"
)

type BuildBaseRepository struct {
	builds.BaseBuild
	aggregate *config.Aggregate
}

func NewBuildRepositoryBase(base builds.BaseBuild, aggregate *config.Aggregate, outFile string) *BuildBaseRepository {
	res := &BuildBaseRepository{
		BaseBuild: base,
		aggregate: aggregate,
	}
	res.ValuesFunc = res.Values
	res.TmplFile = "static/tmpl/go/init/pkg/query-service/infrastructure/base/domain/repository/mongodb_base/base_repository.go.tpl"
	res.OutFile = outFile
	return res
}

func (b *BuildBaseRepository) Values() map[string]interface{} {
	res := b.BaseBuild.Values()
	return res
}
