package cmd_domain

import (
	"fmt"
	"github.com/liuxd6825/dapr-ddd-cli/pkg/builds"
	"github.com/liuxd6825/dapr-ddd-cli/pkg/config"
	"github.com/liuxd6825/dapr-ddd-cli/pkg/utils"
)

type BuildValueObject struct {
	builds.BaseBuild
	name        string
	valueObject *config.ValueObject
	outFile     string
}

func NewBuildValueObject(base builds.BaseBuild, valueObject *config.ValueObject, outFile string) *BuildValueObject {
	res := &BuildValueObject{
		BaseBuild:   base,
		valueObject: valueObject,
		outFile:     outFile,
	}
	res.ValuesFunc = res.Values
	res.TmplFile = "static/tmpl/go/init/pkg/cmd-service/domain/model/aggregate/value_object.go.tpl"
	res.OutFile = outFile
	return res
}

func (b *BuildValueObject) Values() map[string]interface{} {
	res := b.BaseBuild.Values()
	res["name"] = utils.FirstLower(b.valueObject.Name)
	res["Name"] = utils.FirstUpper(b.valueObject.Name)
	res["Package"] = fmt.Sprintf("%s_model", b.valueObject.Aggregate.SnakeName())
	res["ClassName"] = fmt.Sprintf("%s", b.valueObject.Name)
	res["Properties"] = b.valueObject.Properties
	res["Description"] = b.valueObject.Description
	res["Fields"] = b.valueObject
	b.AddTimePackageValue(res, &b.valueObject.Properties)
	return res
}
