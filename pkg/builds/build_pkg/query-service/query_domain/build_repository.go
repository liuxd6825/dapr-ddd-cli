package query_domain

import (
	"github.com/liuxd6825/dapr-ddd-cli/pkg/builds"
	"github.com/liuxd6825/dapr-ddd-cli/pkg/config"
)

type BuildRepository struct {
	builds.BaseBuild
	entity *config.Entity
}

func NewBuildRepository(base builds.BaseBuild, entity *config.Entity, outFile string) *BuildRepository {
	res := &BuildRepository{
		BaseBuild: base,
		entity:    entity,
	}
	res.TmplFile = "static/tmpl/go/init/pkg/query-service/domain/repository/repository.go.tpl"
	res.OutFile = outFile
	res.ValuesFunc = res.Values
	return res
}

func (b *BuildRepository) Values() map[string]interface{} {
	return b.BaseBuild.ValuesOfEntity(b.entity)
}
