package build_makefile

import (
	"github.com/dapr/dapr-ddd-cli/pkg/cmd-init/builds"
	"github.com/dapr/dapr-ddd-cli/pkg/config"
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
	values := map[string]interface{}{}
	values["ServiceName"] = b.Config.Configuration.ServiceName

	outDir := b.outDir
	list = append(list, b.NewFileBuild("/Makefile.tpl", outDir+"/Makefile", values))

	return b.DoBuild(list...)
}
