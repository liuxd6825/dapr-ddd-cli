package query_infrastructure

import (
	"github.com/liuxd6825/dapr-ddd-cli/pkg/builds"
	"github.com/liuxd6825/dapr-ddd-cli/pkg/config"
)

type BuildTypesDateTime struct {
	builds.BaseBuild
	aggregate *config.Aggregate
}

func NewBuildTypesDateTime(base builds.BaseBuild, outFile string) *BuildTypesDateTime {
	res := &BuildTypesDateTime{
		BaseBuild: base,
	}
	res.ValuesFunc = res.Values
	res.TmplFile = "static/tmpl/go/init/pkg/query-service/infrastructure/types/date_time.go.tpl"
	res.OutFile = outFile
	return res
}

func (b *BuildTypesDateTime) Values() map[string]interface{} {
	res := b.BaseBuild.Values()
	return res
}
