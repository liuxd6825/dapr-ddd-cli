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
	res := b.BaseBuild.Values()
	props := config.NewProperties(b.Aggregate, &b.entity.Properties, b.Config.GetDefaultViewProperties())
	res["name"] = utils.FirstLower(b.entity.Name)
	res["Name"] = utils.FirstUpper(b.entity.Name)
	res["ClassName"] = fmt.Sprintf("%sView", utils.FirstUpper(b.entity.Name))
	res["Properties"] = props
	res["Description"] = b.entity.Description
	res["Aggregate"] = b.Aggregate
	res["HasTimeType"] = b.HasTimeType()
	return res
}

func (b *BuildProjectionEntity) HasTimeType() bool {
	hasTimeType := b.entity.Properties.HasTimeType()
	if !hasTimeType {
		defaultProperties := config.NewProperties(b.Aggregate, b.Config.GetDefaultViewProperties(), &b.entity.Properties)
		hasTimeType = defaultProperties.HasTimeType()
	}
	return hasTimeType
}
