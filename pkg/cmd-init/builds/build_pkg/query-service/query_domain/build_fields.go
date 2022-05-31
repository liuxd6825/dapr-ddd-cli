package query_domain

import (
	"github.com/liuxd6825/dapr-ddd-cli/pkg/cmd-init/builds"
	"github.com/liuxd6825/dapr-ddd-cli/pkg/config"
	"github.com/liuxd6825/dapr-ddd-cli/pkg/utils"
)

type BuildFields struct {
	builds.BaseBuild
	name   string
	fields *config.Fields
	values interface{}
}

func NewBuildFields(base builds.BaseBuild, name string, field *config.Fields, outFile string) *BuildFields {
	res := &BuildFields{
		BaseBuild: base,
		name:      name,
		fields:    field,
	}
	res.ValuesFunc = res.Values
	res.TmplFile = "static/tmpl/go/init/pkg/query-service/domain/field/field/field.go.tpl"
	res.OutFile = outFile
	return res
}

func (b *BuildFields) Values() map[string]interface{} {
	res := b.BaseBuild.Values()
	res["name"] = utils.FirstLower(b.fields.Name)
	res["Name"] = utils.FirstUpper(b.fields.Name)
	res["ClassName"] = utils.FirstUpper(b.fields.Name)
	res["Properties"] = b.fields.Properties
	res["Description"] = b.fields.Description
	res["Fields"] = b.fields
	return res
}
