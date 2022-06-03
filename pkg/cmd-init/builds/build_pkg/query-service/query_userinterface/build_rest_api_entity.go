package query_userinterface

import (
	"fmt"
	"github.com/liuxd6825/dapr-ddd-cli/pkg/cmd-init/builds"
	"github.com/liuxd6825/dapr-ddd-cli/pkg/config"
	"github.com/liuxd6825/dapr-ddd-cli/pkg/utils"
)

type BuildRestControllerEntity struct {
	builds.BaseBuild
	entity    *config.Entity
	aggregate *config.Aggregate
}

func NewBuildRestControllerEntity(base builds.BaseBuild, aggregate *config.Aggregate, entity *config.Entity, outFile string) *BuildRestControllerEntity {
	res := &BuildRestControllerEntity{
		BaseBuild: base,
		entity:    entity,
		aggregate: aggregate,
	}
	res.ValuesFunc = res.Values
	res.TmplFile = "static/tmpl/go/init/pkg/query-service/userinterface/rest/facade/api_entity.go.tpl"
	res.OutFile = outFile
	return res
}

func (b *BuildRestControllerEntity) Values() map[string]interface{} {
	res := b.BaseBuild.Values()
	res["Name"] = b.entity.Name
	res["EntityPluralName"] = b.entity.PluralName()
	res["AggregatePluralName"] = b.Aggregate.PluralName()
	res["ServiceName"] = b.entity.FirstUpperName() + "AppService"
	res["ParentId"] = fmt.Sprintf("{%sId}", utils.FirstLower(b.Aggregate.Name))
	return res
}
