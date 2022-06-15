package cmd_domain

import (
	"fmt"
	"github.com/liuxd6825/dapr-ddd-cli/pkg/cmd-init/builds"
	"github.com/liuxd6825/dapr-ddd-cli/pkg/config"
	"github.com/liuxd6825/dapr-ddd-cli/pkg/utils"
	"strings"
)

type BuildEvent struct {
	builds.BaseBuild
	name  string
	event *config.Event
}

func NewBuildEvent(base builds.BaseBuild, name string, event *config.Event, outFile string) *BuildEvent {
	res := &BuildEvent{
		BaseBuild: base,
		name:      name,
		event:     event,
	}
	res.TmplFile = "static/tmpl/go/init/pkg/cmd-service/domain/event/event/event.go.tpl"
	res.OutFile = outFile
	res.ValuesFunc = res.Values
	return res
}

func (b *BuildEvent) Values() map[string]interface{} {
	res := b.BaseBuild.Values()
	res["Namespace"] = b.Namespace()
	res["ClassName"] = b.ClassName()
	res["Name"] = b.name
	res["Version"] = b.Version()
	res["Properties"] = b.event.Properties
	res["Package"] = fmt.Sprintf("%s_event", b.Aggregate.SnakeName())
	res["FieldPackage"] = fmt.Sprintf("%s_field", b.Aggregate.SnakeName())
	res["Aggregate"] = b.Aggregate
	res["AggregateName"] = b.Aggregate.Name
	res["ServiceName"] = b.Config.Configuration.ServiceName
	res["EventType"] = b.event.EventType
	res["FieldName"] = b.FieldName()
	res["EventName"] = b.event.Name
	res["Event"] = b.event
	res["HasDataProperty"] = b.event.HasDataProperty()
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
	return fmt.Sprintf("%s", utils.FirstUpper(b.event.Name))
}

func (b *BuildEvent) Version() string {
	return strings.ToLower(b.event.Version)
}

func (b *BuildEvent) FieldName() string {
	if b.event.HasDataProperty() {
		return b.event.DataProperty.Type
	}
	name := "Fields"
	if b.event.Action == "delete" {
		name = "DeleteFields"
	}
	if b.event.IsAggregate() {
		return b.Aggregate.Name + name
	}
	return b.event.To + name
}
