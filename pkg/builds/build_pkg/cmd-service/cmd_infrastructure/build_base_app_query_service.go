package cmd_infrastructure

import (
	"github.com/liuxd6825/dapr-ddd-cli/pkg/builds"
)

type BuildAppBaseQueryService struct {
	builds.BaseBuild
}

func NewBuildAppBaseQueryService(base builds.BaseBuild, outFile string) *BuildAppBaseQueryService {
	res := &BuildAppBaseQueryService{
		BaseBuild: base,
	}
	res.ValuesFunc = res.Values
	res.TmplFile = "static/tmpl/go/init/pkg/cmd-service/infrastructure/base/application/service/base_query_appservice.go.tpl"
	res.OutFile = outFile

	return res
}

func (b *BuildAppBaseQueryService) Values() map[string]interface{} {
	res := b.BaseBuild.Values()
	return res
}
