package cmd_service

import (
	"fmt"
	"github.com/dapr/dapr-ddd-cli/pkg/cmd-init/builds"
	"github.com/dapr/dapr-ddd-cli/pkg/config"
	"github.com/dapr/dapr-ddd-cli/pkg/utils"
	"strings"
)

type BuildEvent struct {
	builds.BaseBuild
	name  string
	event *config.Event
	dir   string
}

func NewBuildEvent(base builds.BaseBuild, name string, event *config.Event, dir string) *BuildEvent {
	res := &BuildEvent{
		BaseBuild: base,
		name:      name,
		event:     event,
		dir:       dir,
	}
	res.TmplFile = "static/tmpl/go/init/pkg/cmd-service/domain/event/events/event.go.tpl"
	res.OutFile = ""
	res.ValuesFunc = res.Values
	return res
}

func (b *BuildEvent) Values() map[string]interface{} {
	res := b.BaseBuild.Values()
	res["Namespace"] = b.Namespace()
	res["ClassName"] = b.ClassName()
	res["Version"] = b.Version()
	res["Properties"] = b.event.Properties
	return res
}

func (b *BuildEvent) Package() string {
	return fmt.Sprintf("%s_events", utils.FirstLower(b.AggregateName()))
}

func (b *BuildEvent) ClassName() string {
	version := strings.ToLower(b.event.Version)
	version = strings.ReplaceAll(version, ".", "s")
	if version[0] == 'v' {
		version = version[1:]
	}
	return fmt.Sprintf("%sV%s", utils.FirstUpper(b.event.Name), version)
}

func (b *BuildEvent) Version() string {
	return strings.ToLower(b.event.Version)
}
