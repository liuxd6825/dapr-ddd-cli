package query_infrastructure

import (
	"github.com/liuxd6825/dapr-ddd-cli/pkg/builds"
	"github.com/liuxd6825/dapr-ddd-cli/pkg/utils"
)

type BuildRegisterRestApi struct {
	builds.BaseBuild
}

func NewBuildRegisterRestApi(base builds.BaseBuild, outFile string) *BuildRegisterRestApi {
	res := &BuildRegisterRestApi{
		BaseBuild: base,
	}
	res.ValuesFunc = res.Values
	res.TmplFile = "static/tmpl/go/init/pkg/query-service/infrastructure/register/register_rest_api.go.tpl"
	res.OutFile = outFile
	return res
}

func (b *BuildRegisterRestApi) Values() map[string]interface{} {
	res := b.BaseBuild.Values()
	return res
}

func (b *BuildRegisterRestApi) ClassName() string {
	return utils.FirstUpper(b.AggregateName() + "RestController")
}
