package build_config

import (
	"github.com/liuxd6825/dapr-ddd-cli/pkg/builds"
	"github.com/liuxd6825/dapr-ddd-cli/pkg/config"
)

type BuildConfigLayer struct {
	builds.BaseBuild
	aggregate *config.Aggregate
	outDir    string
}

func NewBuildConfigLayer(cfg *config.Config, outDir string) *BuildConfigLayer {
	res := &BuildConfigLayer{
		BaseBuild: builds.BaseBuild{
			Config: cfg,
		},
		outDir: outDir,
	}
	return res
}

func (b *BuildConfigLayer) Build() error {
	var list []builds.Build
	values := b.Values()
	values["CommandServiceName"] = b.Config.Configuration.CommandServiceName()
	values["QueryServiceName"] = b.Config.Configuration.QueryServiceName()
	values["Metadata"] = b.Config.Configuration.GoMetadata
	outDir := b.outDir

	list = append(list, b.NewFileBuild("/config/cmd-config.yaml.tpl", outDir+"/cmd-config.yaml", values))
	list = append(list, b.NewFileBuild("/config/query-config.yaml.tpl", outDir+"/query-config.yaml", values))

	daprDir := b.outDir + "/dapr"
	list = append(list, b.NewFileBuild("/config/dapr/config.yaml.tpl", daprDir+"/config.yaml", values))
	list = append(list, b.NewFileBuild("/config/dapr/components/applogger-mongo.yaml.tpl", daprDir+"/components/applogger-mongo.yaml", values))
	list = append(list, b.NewFileBuild("/config/dapr/components/eventstorage-mongo.yaml.tpl", daprDir+"/components/eventstorage-mongo.yaml", values))
	list = append(list, b.NewFileBuild("/config/dapr/components/pubsub.yaml.tpl", daprDir+"/components/pubsub.yaml", values))
	list = append(list, b.NewFileBuild("/config/dapr/components/pubsub-rabbitmq.yaml.tpl", daprDir+"/components/pubsub-rabbitmq.yaml", values))
	list = append(list, b.NewFileBuild("/config/dapr/components/statestore.yaml.tpl", daprDir+"/components/statestore.yaml", values))

	return b.DoBuild(list...)
}
