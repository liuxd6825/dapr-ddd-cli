package cmd_domain

import (
	"fmt"
	"github.com/liuxd6825/dapr-ddd-cli/pkg/cmd-init/builds"
	"github.com/liuxd6825/dapr-ddd-cli/pkg/config"
	"github.com/liuxd6825/dapr-ddd-cli/pkg/utils"
)

type BuildEnumObject struct {
	builds.BaseBuild
	name       string
	enumObject *config.EnumObject
	outFile    string
}

func NewBuildEnumObject(base builds.BaseBuild, enumObject *config.EnumObject, outFile string) *BuildEnumObject {
	res := &BuildEnumObject{
		BaseBuild:  base,
		enumObject: enumObject,
		outFile:    outFile,
	}
	res.ValuesFunc = res.Values
	res.TmplFile = "static/tmpl/go/init/pkg/cmd-service/domain/model/aggregate/enum_object.go.tpl"
	res.OutFile = outFile
	return res
}

func (b *BuildEnumObject) Values() map[string]interface{} {
	res := b.BaseBuild.Values()
	res["name"] = utils.FirstLower(b.enumObject.Name)
	res["Name"] = utils.FirstUpper(b.enumObject.Name)
	res["Package"] = fmt.Sprintf("%s_model", b.enumObject.Aggregate.SnakeName())
	res["ClassName"] = fmt.Sprintf("%s", b.enumObject.Name)
	res["Values"] = b.enumObject.EnumValues
	res["Description"] = b.enumObject.Description
	return res
}
