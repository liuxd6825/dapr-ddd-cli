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
	res.TmplFile = "static/tmpl/go/init/pkg/cmd-service/domain/field/field/field.go.tpl"
	res.OutFile = outFile
	return res
}

func (b *BuildField) Values() map[string]interface{} {
	res := b.BaseBuild.Values()
	defaultProperties := config.NewProperties(b.Aggregate, b.Config.GetDefaultFieldProperties(), &b.fields.Properties)
	res["DefaultProperties"] = defaultProperties
	res["name"] = utils.FirstLower(b.fields.Name)
	res["Name"] = utils.FirstUpper(b.fields.Name)
	res["ClassName"] = utils.FirstUpper(b.fields.Name)
	res["Properties"] = b.fields.Properties
	res["Description"] = b.fields.Description
	res["Fields"] = b.fields
	res["IsEntity"] = !strings.Contains(b.fields.Name, b.AggregateName())
	return res
}
