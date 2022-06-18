package query_infrastructure

import (
	"github.com/liuxd6825/dapr-ddd-cli/pkg/builds"
	"github.com/liuxd6825/dapr-ddd-cli/pkg/config"
)

type BuildBaseDto struct {
	builds.BaseBuild
	aggregate *config.Aggregate
}

func NewBuildBaseDto(base builds.BaseBuild, aggregate *config.Aggregate, outFile string) *BuildBaseDto {
	res := &BuildBaseDto{
		BaseBuild: base,
		aggregate: aggregate,
	}
	res.ValuesFunc = res.Values
	res.TmplFile = "static/tmpl/go/init/pkg/query-service/infrastructure/base/userinterface/rest/dto/base_dto.go.tpl"
	res.OutFile = outFile
	return res
}

func (b *BuildBaseDto) Values() map[string]interface{} {
	res := b.BaseBuild.Values()
	res["Properties"] = b.Config.Configuration.DefaultReservedProperties.ViewProperties
	return res
}
