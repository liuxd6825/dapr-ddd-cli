package executor

import (
	"context"
	"{{.Namespace}}/pkg/query-service/application/internals/{{.aggregate_name}}/appquery"
	"{{.Namespace}}/pkg/query-service/domain/{{.aggregate_name}}/view"
)

//
// {{.Name}}CreateExecutor
// @Description: 新建
//
type {{.Name}}CreateExecutor interface {
	Execute(context.Context, *view.{{.Name}}View) error
}

//
// {{.Name}}CreateManyExecutor
// @Description: 新建多个
//
type {{.Name}}CreateManyExecutor interface {
	Execute(context.Context, []*view.{{.Name}}View) error
}

//
// {{.Name}}DeleteByIdExecutor
// @Description: 按Id删除
//
type {{.Name}}DeleteByIdExecutor interface {
	Execute(ctx context.Context, tenantId string, id string) error
}

{{- if .IsEntity }}
//
// {{.Name}}DeleteBy{{.AggregateName}}IdExecutor
// @Description: 按聚合根Id删除
//
type {{.Name}}DeleteBy{{.AggregateName}}IdExecutor interface {
	Execute(ctx context.Context, tenantId string, {{.aggregateName}}Id string) error
}
{{- end }}

//
// {{.Name}}DeleteManyExecutor
// @Description: 删除多个
//
type {{.Name}}DeleteManyExecutor interface {
	Execute(context.Context, string, []*view.{{.Name}}View) error
}

//
// {{.Name}}DeleteAllExecutor
// @Description: 删除所有
//
type {{.Name}}DeleteAllExecutor interface {
	Execute(ctx context.Context, tenantId string) error
}

//
// {{.Name}}UpdateExecutor
// @Description: 更新
//
type {{.Name}}UpdateExecutor interface {
	Execute(context.Context, *view.{{.Name}}View) error
}

//
// {{.Name}}UpdateManyExecutor
// @Description: 更新多个
//
type {{.Name}}UpdateManyExecutor interface {
	Execute(context.Context, []*view.{{.Name}}View) error
}

//
// {{.Name}}FindAllExecutor
// @Description: 查询所有
//
type {{.Name}}FindAllExecutor interface {
	Execute(context.Context, *appquery.{{.Name}}FindAllAppQuery) ([]*view.{{.Name}}View, bool, error)
}

//
// {{.Name}}FindByIdExecutor
// @Description: 按Id查询
//
type {{.Name}}FindByIdExecutor interface {
	Execute(context.Context, *appquery.{{.Name}}FindByIdAppQuery) (*view.{{.Name}}View, bool, error)
}

//
// {{.Name}}FindByIdsExecutor
// @Description: 按Id列表查询多个
//
type {{.Name}}FindByIdsExecutor interface {
	Execute(context.Context, *appquery.{{.Name}}FindByIdsAppQuery) ([]*view.{{.Name}}View, bool, error)
}

//
// {{.Name}}FindPagingExecutor
// @Description: 分页查询
//
type {{.Name}}FindPagingExecutor interface {
	Execute(context.Context, *appquery.{{.Name}}FindPagingAppQuery) (*appquery.{{.Name}}FindPagingResult, bool, error)
}

{{- if .IsEntity }}
//
// {{.Name}}FindBy{{.AggregateName}}IdExecutor
// @Description: 按聚合Id查询
//
type {{.Name}}FindBy{{.AggregateName}}IdExecutor interface {
    Execute(context.Context, *appquery.{{.Name}}FindBy{{.AggregateName}}IdAppQuery) ([]*view.{{.Name}}View, bool, error)
}
{{- end }}