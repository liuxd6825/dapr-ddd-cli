package cmd_domain

import (
	"fmt"
	"github.com/liuxd6825/dapr-ddd-cli/pkg/builds"
	"github.com/liuxd6825/dapr-ddd-cli/pkg/config"
	"github.com/liuxd6825/dapr-ddd-cli/pkg/utils"
)

type BuildEntityList struct {
	builds.BaseBuild
	name    string
	entity  *config.Entity
	outFile string
}

func NewBuildEntityList(base builds.BaseBuild, Entity *config.Entity, outFile string) *BuildEntityList {
	res := &BuildEntityList{
		BaseBuild: base,
		entity:    Entity,
		outFile:   outFile,
	}
	res.ValuesFunc = res.Values
	res.TmplFile = "static/tmpl/go/init/pkg/cmd-service/domain/model/aggregate/entity_list.go.tpl"
	res.OutFile = outFile
	return res
}

func (b *BuildEntityList) Values() map[string]interface{} {
	res := b.BaseBuild.Values()
	res["name"] = utils.FirstLower(b.entity.Name)
	res["Name"] = utils.FirstUpper(b.entity.Name)
	res["Entity"] = b.entity
	res["Package"] = fmt.Sprintf("%s_model", b.entity.Aggregate.SnakeName())
	res["ClassName"] = utils.FirstUpper(utils.Plural(b.entity.Name))
	res["Properties"] = b.entity.Properties
	res["Description"] = b.entity.Description
	return res
}
