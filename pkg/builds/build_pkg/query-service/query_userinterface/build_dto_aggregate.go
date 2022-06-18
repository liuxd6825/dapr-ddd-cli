package query_userinterface

import (
	"github.com/liuxd6825/dapr-ddd-cli/pkg/builds"
	"github.com/liuxd6825/dapr-ddd-cli/pkg/config"
)

type BuildDtoAggregate struct {
	builds.BaseBuild
	aggregate *config.Aggregate
}

func NewBuildDtoAggregate(base builds.BaseBuild, aggregate *config.Aggregate, outFile string) *BuildDtoAggregate {
	res := &BuildDtoAggregate{
		BaseBuild: base,
		aggregate: aggregate,
	}
	res.ValuesFunc = res.Values
	res.TmplFile = "static/tmpl/go/init/pkg/query-service/userinterface/rest/dto/dto_aggregate.go.tpl"
	res.OutFile = outFile
	return res
}

func (b *BuildDtoAggregate) Values() map[string]interface{} {
	res := b.BaseBuild.Values()
	props := config.NewProperties(b.aggregate, &b.aggregate.Properties, b.Config.GetDefaultViewProperties())
	res["Name"] = b.aggregate.Name
	res["Properties"] = props
	res["Description"] = b.aggregate.Description
	res["name"] = b.aggregate.LowerName()
	return res
}
