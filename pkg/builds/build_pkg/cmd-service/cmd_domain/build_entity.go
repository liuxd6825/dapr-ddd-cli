package cmd_domain

import (
	"fmt"
	"github.com/liuxd6825/dapr-ddd-cli/pkg/builds"
	"github.com/liuxd6825/dapr-ddd-cli/pkg/config"
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
	values := b.BaseBuild.ValuesOfEntity(b.entity)
	values["Package"] = fmt.Sprintf("%s_model", b.entity.Aggregate.SnakeName())
	values["ClassName"] = fmt.Sprintf("%s", b.entity.Name)
	values["Commands"] = b.entity.EntityCommands()
	values["Events"] = b.entity.EntityEvents()
	b.AddTimePackageValue(values, &b.entity.Properties)
	return values
}
