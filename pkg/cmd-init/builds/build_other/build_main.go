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
	values := map[string]interface{}{}
	values["ServiceName"] = b.Config.Configuration.ServiceName
	values["Description"] = b.Config.Configuration.Description
	values["Namespace"] = b.Config.Configuration.Namespace.Go
	values["Metadata"] = b.Config.Configuration.Metadata
	values["GoMetadata"] = b.Config.Configuration.Go
	values["JavaMetadata"] = b.Config.Configuration.Java
	values["C#Metadata"] = b.Config.Configuration.CSharp
	outDir := b.outDir
	list = append(list, b.NewFileBuild("/Makefile.tpl", outDir+"/Makefile", values))
	list = append(list, b.NewFileBuild("/go.mod.tpl", outDir+"/go.mod", values))
	list = append(list, b.NewFileBuild("/README.md.tpl", outDir+"/README.md", values))

	return b.DoBuild(list...)
}
