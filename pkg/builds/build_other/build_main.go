package build_other

import (
	"github.com/liuxd6825/dapr-ddd-cli/pkg/builds"
	"github.com/liuxd6825/dapr-ddd-cli/pkg/config"
)

type BuildMakefile struct {
	builds.BaseBuild
	outDir string
}

func NewBuildMakefile(cfg *config.Config, outDir string) *BuildMakefile {
	res := &BuildMakefile{
		BaseBuild: builds.BaseBuild{
			Config: cfg,
		},
		outDir: outDir,
	}
	res.init()
	return res
}

func (b *BuildMakefile) init() {
	values := b.BaseBuild.Values()
	outDir := b.outDir
	b.AddBuild(b.NewFileBuild("/Makefile.tpl", outDir+"/Makefile", values))
	b.AddBuild(b.NewFileBuild("/go.mod.tpl", outDir+"/go.mod", values))
	b.AddBuild(b.NewFileBuild("/go.sum.tpl", outDir+"/go.sum", values))
	b.AddBuild(b.NewFileBuild("/README.md.tpl", outDir+"/README.md", values))
}
