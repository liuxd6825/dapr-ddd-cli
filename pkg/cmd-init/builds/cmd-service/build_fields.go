package cmd_service

import (
	"github.com/dapr/dapr-ddd-cli/pkg/cmd-init/builds"
	"github.com/dapr/dapr-ddd-cli/pkg/config"
	"github.com/dapr/dapr-ddd-cli/pkg/utils"
)

type BuildFields struct {
	builds.BaseBuild
	name   string
	fields *config.Fields
	dir    string
	values interface{}
}

func NewBuildFields(base builds.BaseBuild, name string, field *config.Fields, dir string) *BuildFields {
	res := &BuildFields{
		BaseBuild: base,
		name:      name,
		fields:    field,
		dir:       dir,
	}
	res.ValuesFunc = res.Values
	res.TmplFile = "static/tmpl/go/init/pkg/cmd-service/domain/fields/fields/fields.go.tpl"
	res.OutFile = ""
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
