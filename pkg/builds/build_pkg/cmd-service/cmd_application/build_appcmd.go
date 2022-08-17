package cmd_application

import (
	"github.com/liuxd6825/dapr-ddd-cli/pkg/builds"
	"github.com/liuxd6825/dapr-ddd-cli/pkg/config"
)

type BuildAppCmd struct {
	builds.BaseBuild
	entity *config.Entity
}

func NewBuildAppCmd(base builds.BaseBuild, entity *config.Entity, outFile string) *BuildAppCmd {
	res := &BuildAppCmd{
		BaseBuild: base,
		entity:    entity,
	}
	res.ValuesFunc = res.Values
	res.TmplFile = "static/tmpl/go/init/pkg/cmd-service/application/internals/appcmd/appcmd.go.tpl"
	res.OutFile = outFile
	return res
}

func (b *BuildAppCmd) Values() map[string]interface{} {
	values := b.BaseBuild.ValuesOfEntity(b.entity)
	if b.entity != nil {
		values["Commands"] = b.entity.GetCommands()
	} else {
		values["Commands"] = b.Aggregate.AggregateCommands()
	}
	return values
}
