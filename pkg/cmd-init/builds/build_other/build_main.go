package build_other

import (
	"github.com/liuxd6825/dapr-ddd-cli/pkg/cmd-init/builds"
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
	return res
}

func (b *BuildMakefile) Build() error {
	var list []builds.Build
	values := b.BaseBuild.Values()
	outDir := b.outDir
	list = append(list, b.NewFileBuild("/Makefile.tpl", outDir+"/Makefile", values))
	list = append(list, b.NewFileBuild("/go.mod.tpl", outDir+"/go.mod", values))
	list = append(list, b.NewFileBuild("/go.sum.tpl", outDir+"/go.sum", values))
	list = append(list, b.NewFileBuild("/README.md.tpl", outDir+"/README.md", values))

	return b.DoBuild(list...)
}
