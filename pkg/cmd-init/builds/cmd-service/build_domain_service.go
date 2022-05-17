package cmd_service

import (
	"fmt"
	"github.com/dapr/dapr-ddd-cli/pkg/cmd-init/builds"
	"github.com/dapr/dapr-ddd-cli/pkg/config"
	"github.com/dapr/dapr-ddd-cli/pkg/utils"
)

type BuildDomainService struct {
	builds.BaseBuild
	aggregate *config.Aggregate
}

func NewBuildDomainService(base builds.BaseBuild, aggregate *config.Aggregate) *BuildDomainService {
	res := &BuildDomainService{
		BaseBuild: base,
		aggregate: aggregate,
	}
	res.TmplFile = "static/tmpl/go/init/pkg/cmd-service/domain/service/domain_service.go.tpl"
	res.OutFile = ""
	res.ValuesFunc = res.Values
	return res
}

func (b *BuildDomainService) Values() map[string]interface{} {
	res := b.BaseBuild.Values()
	res["ClassName"] = fmt.Sprintf("%sDomainService", utils.FirstUpper(b.AggregateName()))
	res["Commands"] = b.aggregate.Commands
	res["AggregateName"] = b.Aggregate.Name
	res["Package"] = fmt.Sprintf("%s_model", utils.ToLower(b.AggregateName()))
	res["CommandPackage"] = fmt.Sprintf("%s_commands", utils.ToLower(b.aggregate.Name))
	res["ModelPackage"] = fmt.Sprintf("%s_model", utils.ToLower(b.aggregate.Name))
	return res
}
