package build_cmd

import (
	"github.com/liuxd6825/dapr-ddd-cli/pkg/cmd-init/builds"
	"github.com/liuxd6825/dapr-ddd-cli/pkg/config"
)

type BuildMainLayer struct {
	builds.BaseBuild
	aggregate *config.Aggregate
	outDir    string
}

func NewMainLayer(cfg *config.Config, outDir string) *BuildMainLayer {
	res := &BuildMainLayer{
		BaseBuild: builds.BaseBuild{
			Config: cfg,
		},
		outDir: outDir,
	}
	return res
}

func (b *BuildMainLayer) Build() error {
	var list []builds.Build
	values := map[string]interface{}{}
	values["ServiceName"] = b.Config.Configuration.ServiceName
	values["GetNamespace"] = b.Config.Configuration.GetNamespace()
	list = append(list, b.NewFileBuild("/cmd/cmd-service/main.go.tpl", b.outDir+"/cmd-service/main.go", values))
	list = append(list, b.NewFileBuild("/cmd/query-service/main.go.tpl", b.outDir+"/query-service/main.go", values))
	return b.DoBuild(list...)
}
