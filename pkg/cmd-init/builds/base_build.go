package builds

import (
	"github.com/dapr/dapr-ddd-cli/pkg/config"
	"github.com/dapr/dapr-ddd-cli/pkg/utils"
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
}

func NewBaseBuild(config *config.Config, aggregate *config.Aggregate) *BaseBuild {
	res := &BaseBuild{
		Config:    config,
		Aggregate: aggregate,
	}
	res.ValuesFunc = res.Values
	return res
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
	return b.Config.Configuration.Namespace()
}

func (b *BaseBuild) Build() error {
	defer func() {
		if rec := recover(); rec != nil {
			if err, ok := rec.(error); ok {
				println(err)
			} else {
				println(rec)
			}
		}
	}()
	if b.ValuesFunc != nil {
		return utils.RunTemplate(b.TmplFile, b.ValuesFunc(), b.OutFile)
	}
	return utils.RunTemplate(b.TmplFile, nil, b.OutFile)
}

func (b *BaseBuild) Values() map[string]interface{} {
	res := make(map[string]interface{})
	res["Config"] = b.Config
	res["Namespace"] = b.Namespace()
	res["namespace"] = strings.ToLower(b.Namespace())
	res["Aggregates"] = b.Config.Aggregates
	if b.Aggregate != nil {
		res["Aggregate"] = b.Aggregate
		res["AggregateName"] = b.AggregateName()
		res["aggregateName"] = b.aggregateName()
	}
	return res
}

func (b *BaseBuild) NewFileBuild(tmplFile, outFile string, values map[string]interface{}) *BuildAnyFile {
	return NewBuildAnyFile(*b, values, "static/tmpl/go/init"+tmplFile, outFile)
}

func (b *BaseBuild) DoBuild(builds ...Build) error {

	if builds == nil {
		return nil
	}
	for _, build := range builds {
		if err := build.Build(); err != nil {
			return err
		}
	}
	return nil
}

func (b *BaseBuild) GetOutFile() string {
	return b.OutFile
}
