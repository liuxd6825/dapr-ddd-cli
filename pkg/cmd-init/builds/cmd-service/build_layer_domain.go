package cmd_service

import (
	"fmt"
	"github.com/dapr/dapr-ddd-cli/pkg/cmd-init/builds"
	"github.com/dapr/dapr-ddd-cli/pkg/config"
)

type BuildLayerDomain struct {
	builds.BaseBuild
	aggregate      *config.Aggregate
	dir            string
	saveFile       string
	buildFields    []*BuildFields
	buildCommands  []*BuildCommand
	buildEvents    []*BuildEvent
	buildAggregate *BuildAggregate
}

func NewBuildLayerDomain(cfg *config.Config, aggregate *config.Aggregate, dir string) *BuildLayerDomain {
	res := &BuildLayerDomain{
		BaseBuild: builds.BaseBuild{
			Config:    cfg,
			Aggregate: aggregate,
		},
		aggregate: aggregate,
		dir:       dir,
	}

	res.initFields()
	res.initCommands()
	res.initEvents()
	res.initModel()
	return res
}

func (b *BuildLayerDomain) Build() error {
	/*
		for _, build := range b.buildFields {
			build.Build()
		}
		for _, build := range b.buildCommands {
			build.Build()
		}
		for _, build := range b.buildEvents {
			build.Build()
		}
		if b.buildAggregate != nil {
			b.buildAggregate.Build()
		}
	*/
	for _, build := range b.buildCommands {
		build.Build()
	}
	return nil
}

func (b *BuildLayerDomain) initCommands() {
	// commandDir := fmt.Sprintf("%s/command/%s_commands", b.dir, b.name)
	for _, command := range b.aggregate.Commands {
		buildCommand := NewBuildCommand(b.BaseBuild, command)
		b.buildCommands = append(b.buildCommands, buildCommand)
	}
}

func (b *BuildLayerDomain) initEvents() {
	dir := fmt.Sprintf("%s/event/%s_events", b.dir, b.aggregate.Name)
	for name, event := range b.aggregate.Events {
		item := NewBuildEvent(b.BaseBuild, name, event, dir)
		b.buildEvents = append(b.buildEvents, item)
	}
}

func (b *BuildLayerDomain) initFields() {
	dir := fmt.Sprintf("%s/fields/%s_fields", b.dir, b.aggregate.Name)
	for name, fields := range b.aggregate.FieldsObjects {
		item := NewBuildFields(b.BaseBuild, name, fields, dir)
		b.buildFields = append(b.buildFields, item)
	}
}

func (b *BuildLayerDomain) initModel() {
	b.buildAggregate = NewBuildAggregate(b.BaseBuild, b.aggregate)
}
