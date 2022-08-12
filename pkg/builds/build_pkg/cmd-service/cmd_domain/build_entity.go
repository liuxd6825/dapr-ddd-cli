package cmd_domain

import (
	"fmt"
	"github.com/liuxd6825/dapr-ddd-cli/pkg/builds"
	"github.com/liuxd6825/dapr-ddd-cli/pkg/config"
	"github.com/liuxd6825/dapr-ddd-cli/pkg/utils"
)

type BuildEntityObject struct {
	builds.BaseBuild
	name    string
	entity  *config.Entity
	outFile string
}

func NewBuildEntityObject(base builds.BaseBuild, Entity *config.Entity, outFile string) *BuildEntityObject {
	res := &BuildEntityObject{
		BaseBuild: base,
		entity:    Entity,
		outFile:   outFile,
	}
	res.ValuesFunc = res.Values
	res.TmplFile = "static/tmpl/go/init/pkg/cmd-service/domain/model/aggregate/entity.go.tpl"
	res.OutFile = outFile
	return res
}

func (b *BuildEntityObject) Values() map[string]interface{} {
	values := b.BaseBuild.Values()
	values["name"] = utils.FirstLower(b.entity.Name)
	values["Name"] = utils.FirstUpper(b.entity.Name)
	values["Package"] = fmt.Sprintf("%s_model", b.entity.Aggregate.SnakeName())
	values["ClassName"] = fmt.Sprintf("%s", b.entity.Name)
	values["Properties"] = b.entity.Properties
	values["Description"] = b.entity.Description
	values["Fields"] = b.entity
	b.AddTimePackageValue(values, &b.entity.Properties)
	return values
}
