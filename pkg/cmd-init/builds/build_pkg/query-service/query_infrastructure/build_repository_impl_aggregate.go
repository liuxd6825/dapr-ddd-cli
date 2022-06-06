package query_infrastructure

import (
	"fmt"
	"github.com/liuxd6825/dapr-ddd-cli/pkg/cmd-init/builds"
	"github.com/liuxd6825/dapr-ddd-cli/pkg/config"
	"github.com/liuxd6825/dapr-ddd-cli/pkg/utils"
)

type BuildRepositoryImplAggregate struct {
	builds.BaseBuild
	aggregate *config.Aggregate
}

func NewBuildRepositoryImplAggregate(base builds.BaseBuild, aggregate *config.Aggregate, outFile string) *BuildRepositoryImplAggregate {
	res := &BuildRepositoryImplAggregate{
		BaseBuild: base,
		aggregate: aggregate,
	}
	res.ValuesFunc = res.Values
	res.TmplFile = "static/tmpl/go/init/pkg/query-service/infrastructure/domain/repository_impl/mongodb/repository_impl_aggregate.go.tpl"
	res.OutFile = outFile
	return res
}

func (b *BuildRepositoryImplAggregate) Values() map[string]interface{} {
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
	res["name"] = utils.SnakeString(b.aggregate.Name)
	return res
}

func (b *BuildRepositoryImplAggregate) ClassName() string {
	return utils.FirstUpper(b.aggregate.Name + "RepositoryImpl")
}

func (b *BuildRepositoryImplAggregate) AggregateType() string {
	return utils.FirstUpper(fmt.Sprintf("%s.%s", b.Namespace(), b.ClassName()))
}
