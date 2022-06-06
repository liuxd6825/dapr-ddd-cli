package query_infrastructure

import (
	"github.com/liuxd6825/dapr-ddd-cli/pkg/cmd-init/builds"
)

type BuildUtils struct {
	builds.BaseBuild
}

func NewBuildUtils(base builds.BaseBuild, outFile string) *BuildUtils {
	res := &BuildUtils{
		BaseBuild: base,
	}
	res.ValuesFunc = res.Values
	res.TmplFile = "static/tmpl/go/init/pkg/query-service/infrastructure/utils/util.go.tpl"
	res.OutFile = outFile
	return res
}

func (b *BuildUtils) Values() map[string]interface{} {
	res := b.BaseBuild.Values()
	res["HasTimeType"] = b.Config.Configuration.DefaultReservedProperties.ViewProperties.HasTimeType()
	return res
}
