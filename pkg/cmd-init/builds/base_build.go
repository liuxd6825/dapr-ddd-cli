package builds

import (
	"github.com/dapr/dapr-ddd-cli/pkg/config"
	"github.com/dapr/dapr-ddd-cli/pkg/utils"
)

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
	return utils.RunTemplate(b.TmplFile, b.ValuesFunc(), b.OutFile)
}

func (b *BaseBuild) Values() map[string]interface{} {
	res := make(map[string]interface{})
	res["Config"] = b.Config
	res["Aggregate"] = b.Aggregate
	res["AggregateName"] = b.AggregateName()
	res["aggregateName"] = b.aggregateName()
	res["Namespace"] = b.Namespace()
	res["Aggregates"] = b.Config.Aggregates
	return res
}
