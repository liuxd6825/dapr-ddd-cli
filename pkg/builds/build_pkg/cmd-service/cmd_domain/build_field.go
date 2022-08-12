package cmd_domain

import (
	"github.com/liuxd6825/dapr-ddd-cli/pkg/builds"
	"github.com/liuxd6825/dapr-ddd-cli/pkg/config"
	"github.com/liuxd6825/dapr-ddd-cli/pkg/utils"
	"strings"
)

type BuildField struct {
	builds.BaseBuild
	name   string
	fields *config.Fields
	values interface{}
}

func NewBuildField(base builds.BaseBuild, name string, field *config.Fields, outFile string) *BuildField {
	res := &BuildField{
		BaseBuild: base,
		name:      name,
		fields:    field,
	}
	res.ValuesFunc = res.Values
	res.TmplFile = "static/tmpl/go/init/pkg/cmd-service/domain/field/field.go.tpl"
	res.OutFile = outFile
	return res
}

func (b *BuildField) Values() map[string]interface{} {
	defaultProperties := config.NewProperties(b.Aggregate, b.Config.GetDefaultFieldProperties(), &b.fields.Properties)

	values := b.BaseBuild.Values()
	values["DefaultProperties"] = defaultProperties
	values["name"] = utils.FirstLower(b.fields.Name)
	values["Name"] = utils.FirstUpper(b.fields.Name)
	values["ClassName"] = utils.FirstUpper(b.fields.Name)
	values["Properties"] = b.fields.Properties
	values["Description"] = b.fields.Description
	values["Fields"] = b.fields
	values["IsItems"] = b.fields.Properties.IsItems()
	values["IsEntity"] = !strings.Contains(b.fields.Name, b.AggregateName())

	b.AddTimePackageValue(values, defaultProperties)
	b.AddTimePackageValue(values, &b.fields.Properties)
	return values
}
