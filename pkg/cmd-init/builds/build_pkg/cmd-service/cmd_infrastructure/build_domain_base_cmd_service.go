package cmd_infrastructure

import (
	"github.com/liuxd6825/dapr-ddd-cli/pkg/cmd-init/builds"
)

type BuildDomainBaseCmdService struct {
	builds.BaseBuild
}

func NewBuildDomainBaseCmdService(base builds.BaseBuild, outFile string) *BuildDomainBaseCmdService {
	res := &BuildDomainBaseCmdService{
		BaseBuild: base,
	}
	res.ValuesFunc = res.Values
	res.TmplFile = "static/tmpl/go/init/pkg/cmd-service/infrastructure/domain/service/base_command_service.go.tpl"
	res.OutFile = outFile
	return res
}

func (b *BuildDomainBaseCmdService) Values() map[string]interface{} {
	res := b.BaseBuild.Values()
	return res
}
