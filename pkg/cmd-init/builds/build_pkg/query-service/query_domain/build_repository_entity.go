package query_domain

import (
	"github.com/dapr/dapr-ddd-cli/pkg/cmd-init/builds"
	"github.com/dapr/dapr-ddd-cli/pkg/config"
)

type BuildRepositoryEntity struct {
	builds.BaseBuild
	Entity *config.Entity
}

func NewBuildRepositoryEntity(base builds.BaseBuild, entity *config.Entity, outFile string) *BuildRepositoryEntity {
	res := &BuildRepositoryEntity{
		BaseBuild: base,
		Entity:    entity,
	}
	res.TmplFile = "static/tmpl/go/init/pkg/query-service/domain/repository/repository_entity.go.tpl"
	res.OutFile = outFile
	res.ValuesFunc = res.Values
	res.Entity = entity
	return res
}

func (b *BuildRepositoryEntity) Values() map[string]interface{} {
	res := b.BaseBuild.Values()
	res["Name"] = b.Entity.FirstUpperName()
	return res
}
