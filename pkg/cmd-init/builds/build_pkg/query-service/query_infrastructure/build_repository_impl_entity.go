package query_infrastructure

import (
	"fmt"
	"github.com/dapr/dapr-ddd-cli/pkg/cmd-init/builds"
	"github.com/dapr/dapr-ddd-cli/pkg/config"
	"github.com/dapr/dapr-ddd-cli/pkg/utils"
)

type BuildRepositoryImplEntity struct {
	builds.BaseBuild
	entity *config.Entity
}

func NewBuildRepositoryImplEntity(base builds.BaseBuild, entity *config.Entity, outFile string) *BuildRepositoryImplEntity {
	res := &BuildRepositoryImplEntity{
		BaseBuild: base,
		entity:    entity,
	}
	res.ValuesFunc = res.Values
	res.TmplFile = "static/tmpl/go/init/pkg/query-service/infrastructure/domain/repository/mongodb/repository_impl_entity.go.tpl"
	res.OutFile = outFile
	return res
}

func (b *BuildRepositoryImplEntity) Values() map[string]interface{} {
	res := b.BaseBuild.Values()
	res["ClassName"] = b.ClassName()
	res["AggregateType"] = b.AggregateType()
	res["Properties"] = b.entity.Properties
	res["Description"] = b.entity.Description
	res["Package"] = "repository_impl"
	res["Name"] = b.entity.Name
	return res
}

func (b *BuildRepositoryImplEntity) ClassName() string {
	return utils.FirstUpper(b.entity.Name + "RepositoryImpl")
}

func (b *BuildRepositoryImplEntity) AggregateType() string {
	return utils.FirstUpper(fmt.Sprintf("%s.%s", b.Namespace(), b.ClassName()))
}
