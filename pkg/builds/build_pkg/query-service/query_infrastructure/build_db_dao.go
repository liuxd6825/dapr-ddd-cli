package query_infrastructure

import (
	"github.com/liuxd6825/dapr-ddd-cli/pkg/builds"
	"github.com/liuxd6825/dapr-ddd-cli/pkg/config"
)

type BuildDbDao struct {
	builds.BaseBuild
	aggregate *config.Aggregate
}

func NewBuildDbDao(base builds.BaseBuild, aggregate *config.Aggregate, outFile string, tmplType string) *BuildDbDao {
	res := &BuildDbDao{
		BaseBuild: base,
		aggregate: aggregate,
	}
	res.ValuesFunc = res.Values
	res.TmplFile = "static/tmpl/go/init/pkg/query-service/infrastructure/db/dao/" + tmplType + "/dao.go.tpl"
	res.OutFile = outFile
	return res
}

func (b *BuildDbDao) Values() map[string]interface{} {
	res := b.BaseBuild.Values()
	return res
}
