package query_userinterface

import (
	"fmt"
	"github.com/liuxd6825/dapr-ddd-cli/pkg/cmd-init/builds"
	"github.com/liuxd6825/dapr-ddd-cli/pkg/config"
	"github.com/liuxd6825/dapr-ddd-cli/pkg/utils"
)

type BuildRestControllerEntity struct {
	builds.BaseBuild
	entity *config.Entity
}

func NewBuildRestControllerEntity(base builds.BaseBuild, entity *config.Entity, outFile string) *BuildRestControllerEntity {
	res := &BuildRestControllerEntity{
		BaseBuild: base,
		entity:    entity,
	}
	res.ValuesFunc = res.Values
	res.TmplFile = "static/tmpl/go/init/pkg/query-service/userinterface/rest/controller/controller_entity.go.tpl"
	res.OutFile = outFile
	return res
}

func (b *BuildRestControllerEntity) Values() map[string]interface{} {
	res := b.BaseBuild.Values()
	res["Name"] = b.entity.Name
	res["ResourceName"] = utils.MidlineString(b.entity.Name)
	res["ServiceName"] = b.entity.FirstUpperName() + "AppService"
	res["ParentResourceName"] = utils.MidlineString(b.Aggregate.Name)
	res["ParentId"] = fmt.Sprintf("{%sId}", utils.FirstLower(b.Aggregate.Name))
	return res
}
