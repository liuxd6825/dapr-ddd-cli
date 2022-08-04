package query_infrastructure

import (
	"github.com/liuxd6825/dapr-ddd-cli/pkg/builds"
	"github.com/liuxd6825/dapr-ddd-cli/pkg/config"
)

type BuildBaseDao struct {
	builds.BaseBuild
	aggregate *config.Aggregate
}

func NewBuildBaseDao(base builds.BaseBuild, aggregate *config.Aggregate, outFile string, tmplType string) *BuildBaseDao {
	res := &BuildBaseDao{
		BaseBuild: base,
		aggregate: aggregate,
	}
	res.ValuesFunc = res.Values
	res.TmplFile = "static/tmpl/go/init/pkg/query-service/infrastructure/base/domain/dao/" + tmplType + "/dao.go.tpl"
	res.OutFile = outFile
	return res
}

func (b *BuildBaseDao) Values() map[string]interface{} {
	res := b.BaseBuild.Values()
	return res
}
