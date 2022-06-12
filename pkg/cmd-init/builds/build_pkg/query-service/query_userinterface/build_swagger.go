package query_userinterface

import (
	"github.com/liuxd6825/dapr-ddd-cli/pkg/cmd-init/builds"
	"github.com/liuxd6825/dapr-ddd-cli/pkg/config"
)

type BuildSwagger struct {
	builds.BaseBuild
	entity    *config.Entity
	aggregate *config.Aggregate
}

func NewBuildSwagger(base builds.BaseBuild, outFile string) *BuildSwagger {
	res := &BuildSwagger{
		BaseBuild: base,
	}
	res.ValuesFunc = res.Values
	res.TmplFile = "static/tmpl/go/init/pkg/query-service/userinterface/rest/swagger_main.go.tpl"
	res.OutFile = outFile
	return res
}

func (b *BuildSwagger) Values() map[string]interface{} {
	res := b.BaseBuild.Values()
	return res
}
