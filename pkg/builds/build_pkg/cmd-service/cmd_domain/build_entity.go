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
	res := b.BaseBuild.Values()
	res["name"] = utils.FirstLower(b.entity.Name)
	res["Name"] = utils.FirstUpper(b.entity.Name)
	res["Package"] = fmt.Sprintf("%s_model", b.entity.Aggregate.SnakeName())
	res["ClassName"] = fmt.Sprintf("%s", b.entity.Name)
	res["Properties"] = b.entity.Properties
	res["Description"] = b.entity.Description
	res["Fields"] = b.entity
	return res
}
