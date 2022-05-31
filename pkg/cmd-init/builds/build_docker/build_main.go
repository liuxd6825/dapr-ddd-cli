package build_docker

import (
	"github.com/liuxd6825/dapr-ddd-cli/pkg/cmd-init/builds"
	"github.com/liuxd6825/dapr-ddd-cli/pkg/config"
)

type BuildDockerLayer struct {
	builds.BaseBuild
	aggregate *config.Aggregate
	outDir    string
}

func NewBuildDockerLayer(cfg *config.Config, outDir string) *BuildDockerLayer {
	res := &BuildDockerLayer{
		BaseBuild: builds.BaseBuild{
			Config: cfg,
		},
		outDir: outDir,
	}
	return res
}

func (b *BuildDockerLayer) Build() error {
	var list []builds.Build
	values := map[string]interface{}{}
	outDir := b.outDir
	list = append(list, b.NewFileBuild("/docker/docker.mk.tpl", outDir+"/docker.mk", values))
	list = append(list, b.NewFileBuild("/docker/cmd/Dockerfile.tpl", outDir+"/cmd/Dockerfile", values))
	list = append(list, b.NewFileBuild("/docker/query/Dockerfile.tpl", outDir+"/query/Dockerfile", values))
	return b.DoBuild(list...)
}
