package query_application

import (
	"fmt"
	"github.com/liuxd6825/dapr-ddd-cli/pkg/cmd-init/builds"
	"github.com/liuxd6825/dapr-ddd-cli/pkg/config"
	"github.com/liuxd6825/dapr-ddd-cli/pkg/utils"
)

type BuildQueryHandlerEntity struct {
	builds.BaseBuild
	aggregate *config.Aggregate
	entity    *config.Entity
}

func NewBuildQueryHandlerEntity(base builds.BaseBuild, aggregate *config.Aggregate, entity *config.Entity, outFile string) *BuildQueryHandlerEntity {
	res := &BuildQueryHandlerEntity{
		BaseBuild: base,
		aggregate: aggregate,
		entity:    entity,
	}
	res.TmplFile = "static/tmpl/go/init/pkg/query-service/application/internals/handler/aggregate/entity_query_handler.go.tpl"
	res.OutFile = outFile
	res.ValuesFunc = res.Values
	return res
}

func (b *BuildQueryHandlerEntity) Values() map[string]interface{} {
	res := b.BaseBuild.Values()
	res["ClassName"] = fmt.Sprintf("%sQueryHandler", utils.FirstUpper(b.entity.Name))
	res["AggregateName"] = utils.FirstUpper(b.Aggregate.Name)
	res["aggregateName"] = utils.FirstLower(b.Aggregate.Name)
	res["Entities"] = b.Aggregate.Entities
	res["Events"] = b.Aggregate.Events.GetEntityEvents(b.entity.Name)
	res["Commands"] = b.aggregate.Commands
	res["Entity"] = b.entity
	res["EntityName"] = b.entity.Name
	res["entityName"] = b.entity.FirstLowerName()
	res["Properties"] = b.entity.Properties
	res["Package"] = fmt.Sprintf("%s_queryhandler", b.aggregate.Name)
	res["ServiceName"] = b.Config.Configuration.ServiceName
	res["Name"] = b.entity.FirstUpperName()
	return res
}
