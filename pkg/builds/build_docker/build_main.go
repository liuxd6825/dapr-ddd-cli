package build_docker

import (
	"github.com/liuxd6825/dapr-ddd-cli/pkg/builds"
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
	res.init()
	return res
}

func (b *BuildDockerLayer) init() {
	values := b.Values()
	outDir := b.outDir
	b.AddBuild(b.NewFileBuild("/docker/docker.mk.tpl", outDir+"/docker.mk", values))
	b.AddBuild(b.NewFileBuild("/docker/cmd/Dockerfile.tpl", outDir+"/cmd/Dockerfile", values))
	b.AddBuild(b.NewFileBuild("/docker/query/Dockerfile.tpl", outDir+"/query/Dockerfile", values))
}
