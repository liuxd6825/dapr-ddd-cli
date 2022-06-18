package cmd_domain

import (
	"fmt"
	"github.com/liuxd6825/dapr-ddd-cli/pkg/builds"
	"github.com/liuxd6825/dapr-ddd-cli/pkg/config"
	"github.com/liuxd6825/dapr-ddd-cli/pkg/utils"
)

type BuildEntityItems struct {
	builds.BaseBuild
	name    string
	entity  *config.Entity
	outFile string
}

func NewBuildEntityItems(base builds.BaseBuild, Entity *config.Entity, outFile string) *BuildEntityItems {
	res := &BuildEntityItems{
		BaseBuild: base,
		entity:    Entity,
		outFile:   outFile,
	}
	res.ValuesFunc = res.Values
	res.TmplFile = "static/tmpl/go/init/pkg/cmd-service/domain/model/aggregate/entity_items.go.tpl"
	res.OutFile = outFile
	return res
}

func (b *BuildEntityItems) Values() map[string]interface{} {
	res := b.BaseBuild.Values()
	res["name"] = utils.FirstLower(b.entity.Name)
	res["Name"] = utils.FirstUpper(b.entity.Name)
	res["Entity"] = b.entity
	res["Package"] = fmt.Sprintf("%s_model", b.entity.Aggregate.SnakeName())
	res["ClassName"] = utils.FirstUpper(b.entity.Name) + "Items"
	res["Properties"] = b.entity.Properties
	res["Description"] = b.entity.Description
	return res
}
