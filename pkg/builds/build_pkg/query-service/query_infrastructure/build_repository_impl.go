package query_infrastructure

import (
	"github.com/liuxd6825/dapr-ddd-cli/pkg/builds"
	"github.com/liuxd6825/dapr-ddd-cli/pkg/config"
)

type BuildRepositoryImpl struct {
	builds.BaseBuild
	entity *config.Entity
}

func NewBuildRepositoryImpl(base builds.BaseBuild, entity *config.Entity, dbType, outFile string) *BuildRepositoryImpl {
	res := &BuildRepositoryImpl{
		BaseBuild: base,
		entity:    entity,
	}
	res.ValuesFunc = res.Values
	res.TmplFile = "static/tmpl/go/init/pkg/query-service/infrastructure/domain/repository_impl/" + dbType + "/repository_impl.go.tpl"
	res.OutFile = outFile
	return res
}

func (b *BuildRepositoryImpl) Values() map[string]interface{} {
	return b.BaseBuild.ValuesOfEntity(b.entity)
}
