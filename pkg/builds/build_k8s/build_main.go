package build_k8s

import (
	"github.com/liuxd6825/dapr-ddd-cli/pkg/builds"
	"github.com/liuxd6825/dapr-ddd-cli/pkg/config"
)

type BuildK8sLayer struct {
	builds.BaseBuild
	aggregate *config.Aggregate
	outDir    string
}

func NewBuildK8sLayer(cfg *config.Config, outDir string) *BuildK8sLayer {
	res := &BuildK8sLayer{
		BaseBuild: builds.BaseBuild{
			Config: cfg,
		},
		outDir: outDir,
	}
	res.init()
	return res
}

func (b *BuildK8sLayer) init() {
	values := b.Values()
	outDir := b.outDir
	b.AddBuild(b.NewFileBuild("/k8s/cmd-service.yaml.tpl", outDir+"/cmd-service.yaml", values))
	b.AddBuild(b.NewFileBuild("/k8s/query-service.yaml.tpl", outDir+"/query-service.yaml", values))
}
