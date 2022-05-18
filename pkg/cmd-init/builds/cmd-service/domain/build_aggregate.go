package domain

import (
	"fmt"
	"github.com/dapr/dapr-ddd-cli/pkg/cmd-init/builds"
	"github.com/dapr/dapr-ddd-cli/pkg/config"
	"github.com/dapr/dapr-ddd-cli/pkg/utils"
)

type BuildAggregate struct {
	builds.BaseBuild
	aggregate *config.Aggregate
}

func NewBuildAggregate(base builds.BaseBuild, aggregate *config.Aggregate, outFile string) *BuildAggregate {
	res := &BuildAggregate{
		BaseBuild: base,
		aggregate: aggregate,
	}
	res.ValuesFunc = res.Values
	res.TmplFile = "static/tmpl/go/init/pkg/cmd-service/domain/model/aggregate.go.tpl"
	res.OutFile = outFile
	return res
}

func (b *BuildAggregate) Values() map[string]interface{} {
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
	res["EventPackage"] = fmt.Sprintf("%s_events", utils.ToLower(b.aggregate.Name))
	res["Package"] = fmt.Sprintf("%s_model", utils.ToLower(b.aggregate.Name))
	res["Version"] = b.aggregate.Version
	return res
}

func (b *BuildAggregate) ClassName() string {
	return utils.FirstUpper(b.AggregateName() + "Aggregate")
}

func (b *BuildAggregate) AggregateType() string {
	return utils.FirstUpper(fmt.Sprintf("%s.%s", b.Namespace(), b.ClassName()))
}