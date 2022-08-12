package query_application

import (
	"fmt"
	"github.com/liuxd6825/dapr-ddd-cli/pkg/builds"
	"github.com/liuxd6825/dapr-ddd-cli/pkg/config"
	"github.com/liuxd6825/dapr-ddd-cli/pkg/utils"
)

type BuildAppServiceAggregate struct {
	builds.BaseBuild
	aggregate *config.Aggregate
}

func NewBuildAppServiceAggregate(base builds.BaseBuild, entity *config.Entity, outFile string) *BuildAppServiceAggregate {
	res := &BuildAppServiceAggregate{
		BaseBuild: base,
	}
	res.ValuesFunc = res.Values
	res.TmplFile = "static/tmpl/go/init/pkg/query-service/application/internals/service/app_service_aggregate.go.tpl"
	res.OutFile = outFile
	return res
}

func (b *BuildAppServiceAggregate) Values() map[string]interface{} {
	res := b.BaseBuild.Values()
	res["ClassName"] = b.ClassName()
	res["AggregateType"] = b.AggregateType()
	res["Properties"] = b.aggregate.Properties
	res["Events"] = b.aggregate.Events
	res["Commands"] = b.aggregate.Commands
	res["Description"] = b.aggregate.Description
	res["EnumObjects"] = b.aggregate.EnumObjects
	res["Id"] = b.aggregate.Id
	res["FieldsObjects"] = b.aggregate.FieldsObjects
	res["Aggregate"] = b.aggregate
	res["CommandPackage"] = fmt.Sprintf("%s_commands", utils.ToLower(b.aggregate.Name))
	res["ModelPackage"] = fmt.Sprintf("%s_model", utils.ToLower(b.aggregate.Name))
	res["Package"] = "repository_impl"
	res["ResourceName"] = fmt.Sprintf("%ss", utils.ToLower(b.aggregate.Name))
	res["Name"] = b.aggregate.Name
	res["name"] = b.aggregate.FirstLowerName()
	return res
}

func (b *BuildAppServiceAggregate) ClassName() string {
	return utils.FirstUpper(b.aggregate.Name + "RepositoryImpl")
}

func (b *BuildAppServiceAggregate) AggregateType() string {
	return utils.FirstUpper(fmt.Sprintf("%s.%s", b.Namespace(), b.ClassName()))
}
