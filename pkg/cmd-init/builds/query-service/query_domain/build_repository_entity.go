package query_domain

import (
	"fmt"
	"github.com/dapr/dapr-ddd-cli/pkg/cmd-init/builds"
	"github.com/dapr/dapr-ddd-cli/pkg/config"
	"github.com/dapr/dapr-ddd-cli/pkg/utils"
)

type BuildRepositoryEntity struct {
	builds.BaseBuild
	entity *config.Entity
}

func NewBuildRepositoryEntity(base builds.BaseBuild, entity *config.Entity, outFile string) *BuildRepositoryEntity {
	res := &BuildRepositoryEntity{
		BaseBuild: base,
		entity:    entity,
	}
	res.TmplFile = "static/tmpl/go/init/pkg/query-service/domain/repository/repository_entity.go.tpl"
	res.OutFile = outFile
	res.ValuesFunc = res.Values
	return res
}

func (b *BuildRepositoryEntity) Values() map[string]interface{} {
	res := b.BaseBuild.Values()
	res["ClassName"] = fmt.Sprintf("%sDomainService", utils.FirstUpper(b.AggregateName()))
	res["AggregateName"] = b.Aggregate.Name
	res["EntityName"] = b.entity.Name
	res["Package"] = fmt.Sprintf("%s_model", utils.ToLower(b.AggregateName()))
	return res
}
