package cmd_domain

import (
	"fmt"
	"github.com/liuxd6825/dapr-ddd-cli/pkg/builds"
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
	values := b.BaseBuild.ValuesOfEvent(b.event)
	values["Namespace"] = b.Namespace()
	values["ClassName"] = b.ClassName()
	values["Name"] = b.name
	if b.event.DataProperty != nil {
		values["FieldName"] = b.event.DataProperty.Type
	}
	b.AddTimePackageValue(values, &b.event.Properties)
	return values
}

func (b *BuildEvent) ClassName() string {
	version := strings.ToLower(b.event.Version)
	version = strings.ReplaceAll(version, ".", "s")
	if version[0] == 'v' {
		version = version[1:]
	}
	return fmt.Sprintf("%s", utils.FirstUpper(b.event.Name))
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
