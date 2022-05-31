package cmd_domain

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

func NewBuildDomainService(base builds.BaseBuild, aggregate *config.Aggregate, outFile string) *BuildDomainService {
	res := &BuildDomainService{
		BaseBuild: base,
		aggregate: aggregate,
	}
	res.TmplFile = "static/tmpl/go/init/pkg/cmd-service/domain/service/domain_service.go.tpl"
	res.OutFile = outFile
	res.ValuesFunc = res.Values
	return res
}

func (b *BuildDomainService) Values() map[string]interface{} {
	res := b.BaseBuild.Values()
	res["ClassName"] = fmt.Sprintf("%sCommandDomainService", utils.FirstUpper(b.AggregateName()))
	res["Commands"] = b.aggregate.Commands
	res["AggregateName"] = b.Aggregate.Name
	res["Package"] = fmt.Sprintf("%s_service", b.aggregate.SnakeName())
	res["CommandPackage"] = fmt.Sprintf("%s_command", b.aggregate.SnakeName())
	res["ModelPackage"] = fmt.Sprintf("%s_model", b.aggregate.SnakeName())
	return res
}
