package builds

import (
	"errors"
	"fmt"
	"github.com/liuxd6825/dapr-ddd-cli/pkg/config"
	"github.com/liuxd6825/dapr-ddd-cli/pkg/utils"
	"os"
	"strings"
)

type Build interface {
	Build() error
	GetOutFile() string
}

type BaseBuild struct {
	Config     *config.Config
	Aggregate  *config.Aggregate
	TmplFile   string
	OutFile    string
	ValuesFunc func() map[string]interface{}
	builds     []Build
}

func NewBaseBuild(config *config.Config, aggregate *config.Aggregate) *BaseBuild {
	res := &BaseBuild{
		Config:    config,
		Aggregate: aggregate,
	}
	res.ValuesFunc = res.Values
	return res
}

func (b *BaseBuild) AddBuild(build Build) {
	b.builds = append(b.builds, build)
}

func (b *BaseBuild) AggregateName() string {
	return utils.FirstUpper(b.Aggregate.Name)
}

func (b *BaseBuild) UpperAggregateName() string {
	return utils.FirstUpper(b.Aggregate.Name)
}

func (b *BaseBuild) LowerAggregateName() string {
	return utils.FirstUpper(b.Aggregate.Name)
}

func (b *BaseBuild) aggregateName() string {
	return utils.FirstLower(b.AggregateName())
}

func (b *BaseBuild) Namespace() string {
	return b.Config.Configuration.GetNamespace()
}

func (b *BaseBuild) Build() (resErr error) {
	defer func() {
		if rec := recover(); rec != nil {
			if err, ok := rec.(error); ok {
				resErr = err
			} else if msg, ok := rec.(string); ok {
				resErr = errors.New(msg)
			}
		}
	}()
	if b.ValuesFunc != nil {
		return utils.RunTemplate(b.TmplFile, b.ValuesFunc(), b.OutFile)
	}
	return utils.RunTemplate(b.TmplFile, nil, b.OutFile)
}

func (b *BaseBuild) Builds() (resErr error) {
	defer func() {
		if rec := recover(); rec != nil {
			if err, ok := rec.(error); ok {
				resErr = err
			} else if msg, ok := rec.(string); ok {
				resErr = errors.New(msg)
			}
		}
	}()
	for _, build := range b.builds {
		if err := build.Build(); err != nil {
			return err
		}
	}
	return nil
}

func (b *BaseBuild) Values() map[string]interface{} {
	res := make(map[string]interface{})
	res["Config"] = b.Config
	res["Configuration"] = b.Config.Configuration
	res["Namespace"] = b.Namespace()
	res["namespace"] = strings.ToLower(b.Namespace())
	res["Aggregates"] = b.Config.Aggregates
	res["ServiceName"] = b.Config.Configuration.ServiceName
	res["CommandServiceName"] = b.Config.Configuration.CommandServiceName()
	res["QueryServiceName"] = b.Config.Configuration.QueryServiceName()
	res["ApiVersion"] = b.Config.Configuration.ApiVersion
	res["K8sNamespace"] = b.Config.Configuration.GetK8sNamespace()
	res["K8sQueryImage"] = b.Config.Configuration.GetK8sQueryImage()
	res["K8sCommandImage"] = b.Config.Configuration.GetK8sCommandImage()
	res["GoMetadata"] = b.Config.Configuration.GoMetadata
	res["JavaMetadata"] = b.Config.Configuration.JavaMetadata
	res["C#Metadata"] = b.Config.Configuration.CSharpMetadata

	if b.Config != nil && b.Config.Configuration != nil {
		res["DefaultViewProperties"] = b.Config.Configuration.DefaultReservedProperties.ViewProperties
		res["DefaultEntityProperties"] = b.Config.Configuration.DefaultReservedProperties.EntityProperties
		res["DefaultValueProperties"] = b.Config.Configuration.DefaultReservedProperties.ValueProperties
		res["DefaultAggregateProperties"] = b.Config.Configuration.DefaultReservedProperties.AggregateProperties
		res["DefaultFieldProperties"] = b.Config.Configuration.DefaultReservedProperties.FieldProperties
	}

	if b.Aggregate != nil {
		aggregateName := utils.SnakeString(b.Aggregate.Name)
		res["Aggregate"] = b.Aggregate
		res["AggregateName"] = b.AggregateName()
		res["aggregateName"] = b.aggregateName()
		res["aggregateMidlineName"] = b.Aggregate.MidlineName()
		res["aggregate_name"] = aggregateName
		res["AggregateCommandPackage"] = fmt.Sprintf("%s/command", aggregateName)
		res["AggregateEventPackage"] = fmt.Sprintf("%s/event", aggregateName)
		res["AggregateFieldPackage"] = fmt.Sprintf("%s/field", aggregateName)
		res["AggregateModelPackage"] = fmt.Sprintf("%s/model", aggregateName)
		res["AggregateFactoryPackage"] = fmt.Sprintf("%s/factory", aggregateName)
		res["AggregateServicePackage"] = fmt.Sprintf("%s/service", aggregateName)

		res["Name"] = b.Aggregate.Name
		res["name"] = b.Aggregate.FirstLowerName()
	}

	return res
}

func (b *BaseBuild) ValuesOfEntity(entity *config.Entity) map[string]interface{} {
	res := b.Values()
	if entity != nil {
		res["IsEntity"] = true
		res["name"] = utils.FirstLower(entity.Name)
		res["Name"] = utils.FirstUpper(entity.Name)
		res["Description"] = entity.Description
		res["Properties"] = entity.Properties
	}
	return res
}

func (b *BaseBuild) NewFileBuild(tmplFile, outFile string, values map[string]interface{}) *BuildAnyFile {
	return NewBuildAnyFile(*b, values, "static/tmpl/go/init"+tmplFile, outFile)
}

func (b *BaseBuild) DoBuild() error {
	if b.builds == nil {
		return nil
	}
	for _, build := range b.builds {
		println("building: " + build.GetOutFile())
		if err := build.Build(); err != nil {
			return err
		}
	}
	return nil
}

func (b *BaseBuild) GetOutFile() string {
	return b.OutFile
}

func (b *BaseBuild) Mkdir(dirs ...string) {
	for _, dirName := range dirs {
		_ = os.MkdirAll(dirName, os.ModePerm)
	}
}
