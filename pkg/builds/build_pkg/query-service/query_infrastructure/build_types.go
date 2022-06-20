package query_infrastructure

import (
	"github.com/liuxd6825/dapr-ddd-cli/pkg/builds"
	"github.com/liuxd6825/dapr-ddd-cli/pkg/config"
)

type BuildTypes struct {
	builds.BaseBuild
	aggregate *config.Aggregate
}

func NewBuildTypes(base builds.BaseBuild, outFile string) *BuildTypes {
	res := &BuildTypes{
		BaseBuild: base,
	}
	res.ValuesFunc = res.Values
	res.TmplFile = "static/tmpl/go/init/pkg/query-service/infrastructure/types/types.go.tpl"
	res.OutFile = outFile
	return res
}

func (b *BuildTypes) Values() map[string]interface{} {
	res := b.BaseBuild.Values()
	return res
}
