package build_cmd

import (
	"github.com/liuxd6825/dapr-ddd-cli/pkg/builds"
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
	res.init()
	return res
}

func (b *BuildMainLayer) init() {
	values := map[string]interface{}{}
	values["ServiceName"] = b.Config.Configuration.ServiceName
	values["Namespace"] = b.Config.Configuration.GetNamespace()

	build1 := b.NewFileBuild("/cmd/cmd-service/main.go.tpl", b.outDir+"/cmd-service/main.go", values)
	b.AddBuild(build1)

	build2 := b.NewFileBuild("/cmd/query-service/main.go.tpl", b.outDir+"/query-service/main.go", values)
	b.AddBuild(build2)
}
