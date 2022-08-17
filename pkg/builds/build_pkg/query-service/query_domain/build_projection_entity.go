package query_domain

import (
	"fmt"
	"github.com/liuxd6825/dapr-ddd-cli/pkg/builds"
	"github.com/liuxd6825/dapr-ddd-cli/pkg/config"
	"github.com/liuxd6825/dapr-ddd-cli/pkg/utils"
)

type BuildProjectionEntity struct {
	builds.BaseBuild
	name   string
	entity *config.Entity
	values interface{}
}

func NewBuildProjectionEntity(base builds.BaseBuild, entity *config.Entity, outFile string) *BuildProjectionEntity {
	res := &BuildProjectionEntity{
		BaseBuild: base,
		entity:    entity,
	}
	res.ValuesFunc = res.Values
	res.TmplFile = "static/tmpl/go/init/pkg/query-service/domain/view/entity_view.go.tpl"
	res.OutFile = outFile
	return res
}

func (b *BuildProjectionEntity) Values() map[string]interface{} {
	values := b.BaseBuild.Values()
	props := config.NewProperties(b.Aggregate, &b.entity.Properties, b.Config.GetDefaultViewProperties())
	values["name"] = utils.FirstLower(b.entity.Name)
	values["Name"] = utils.FirstUpper(b.entity.Name)
	values["ClassName"] = fmt.Sprintf("%sView", utils.FirstUpper(b.entity.Name))
	values["Properties"] = props
	values["Description"] = b.entity.Description
	values["Aggregate"] = b.Aggregate
	values["HasTimeType"] = b.HasTimeType()
	b.AddTimePackageValue(values, props)
	b.AddTimePackageValue(values, &b.entity.Properties)
	return values
}

func (b *BuildProjectionEntity) HasTimeType() bool {
	hasTimeType := b.entity.Properties.HasTimeType()
	if !hasTimeType {
		defaultProperties := config.NewProperties(b.Aggregate, b.Config.GetDefaultViewProperties(), &b.entity.Properties)
		hasTimeType = defaultProperties.HasTimeType()
	}
	return hasTimeType
}
