package cmd_application

import (
	"fmt"
	"github.com/dapr/dapr-ddd-cli/pkg/cmd-init/builds"
	"github.com/dapr/dapr-ddd-cli/pkg/config"
	"github.com/dapr/dapr-ddd-cli/pkg/utils"
)

type BuildApplicationLayer struct {
	builds.BaseBuild
	aggregate                        *config.Aggregate
	outDir                           string
	buildCmdApplicationService       *BuildCmdApplicationService
	buildQueryApplicationService     *BuildQueryApplicationService
	buildBaseQueryApplicationService *builds.BuildAnyFile
}

func NewBuildApplicationLayer(cfg *config.Config, aggregate *config.Aggregate, outDir string) *BuildApplicationLayer {
	res := &BuildApplicationLayer{
		BaseBuild: builds.BaseBuild{
			Config:    cfg,
			Aggregate: aggregate,
		},
		aggregate: aggregate,
		outDir:    outDir,
	}
	res.init()
	return res
}

func (b *BuildApplicationLayer) init() {
	outFile := fmt.Sprintf("%s/internales/cmdappservice/%s_cmd_appservice.go", b.outDir, b.aggregate.Name)
	b.buildCmdApplicationService = NewBuildCmdApplicationService(b.BaseBuild, b.aggregate, utils.ToLower(outFile))

	outFile = fmt.Sprintf("%s/internales/queryappservice/%s_query_service.go", b.outDir, b.aggregate.Name)
	b.buildQueryApplicationService = NewBuildQueryApplicationService(b.BaseBuild, b.aggregate, utils.ToLower(outFile))

	tempFile := "static/tmpl/go/init/pkg/cmd-service/application/internals/queryappservice/base_query_appservice.go.tpl"
	outFile = fmt.Sprintf("%s/internales/queryappservice/base_query_appservice.go", b.outDir)
	b.buildBaseQueryApplicationService = builds.NewBuildAnyFile(b.BaseBuild, map[string]interface{}{}, tempFile, utils.ToLower(outFile))
}

func (b *BuildApplicationLayer) Build() error {
	var list []builds.Build
	list = append(list, b.buildCmdApplicationService)
	list = append(list, b.buildQueryApplicationService)
	list = append(list, b.buildBaseQueryApplicationService)
	return b.DoBuild(list...)
}
