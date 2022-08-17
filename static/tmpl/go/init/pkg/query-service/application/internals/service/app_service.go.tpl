package service
{{$AggregateName := .AggregateName}}
import (
    "sync"
	"context"
    "{{.Namespace}}/pkg/query-service/domain/{{.aggregate_name}}/view"
    "{{.Namespace}}/pkg/query-service/domain/{{.aggregate_name}}/query"
    "{{.Namespace}}/pkg/query-service/application/internals/{{.aggregate_name}}/appquery"
    "{{.Namespace}}/pkg/query-service/application/internals/{{.aggregate_name}}/assembler"
    "{{.Namespace}}/pkg/query-service/application/internals/{{.aggregate_name}}/executor"
    "{{.Namespace}}/pkg/query-service/application/internals/{{.aggregate_name}}/executor/{{.snake_name}}_impl"
)

//
// {{.Name}}QueryAppService
// @Description: {{.Description}}查询应用服务类
//
type {{.Name}}QueryAppService struct {
	{{.name}}CreateExecutor     executor.{{.Name}}CreateExecutor
	{{.name}}CreateManyExecutor executor.{{.Name}}CreateManyExecutor

	{{.name}}UpdateExecutor     executor.{{.Name}}UpdateExecutor
	{{.name}}UpdateManyExecutor executor.{{.Name}}UpdateManyExecutor

	{{.name}}DeleteByIdExecutor executor.{{.Name}}DeleteByIdExecutor
	{{.name}}DeleteManyExecutor executor.{{.Name}}DeleteManyExecutor
	{{.name}}DeleteAllExecutor  executor.{{.Name}}DeleteAllExecutor
	{{- if .IsEntity }}
	{{.name}}DeleteBy{{.AggregateName}}IdExecutor  executor.{{.Name}}DeleteBy{{.AggregateName}}IdExecutor
	{{- end}}

    {{.name}}FindAllExecutor    executor.{{.Name}}FindAllExecutor
    {{.name}}FindByIdExecutor   executor.{{.Name}}FindByIdExecutor
    {{.name}}FindByIdsExecutor  executor.{{.Name}}FindByIdsExecutor
    {{.name}}FindPagingExecutor executor.{{.Name}}FindPagingExecutor
    {{- if .IsEntity }}
    {{.name}}FindBy{{.AggregateName}}IdExecutor  executor.{{.Name}}FindBy{{.AggregateName}}IdExecutor
    {{- end}}
}

// 单例应用服务
var {{.name}}QueryAppService *{{.Name}}QueryAppService

// 并发安全
var once{{.Name}} sync.Once

//
// Get{{.Name}}QueryAppService
// @Description: 获取单例应用服务
// @return *{{.Name}}QueryAppService
//
func Get{{.Name}}QueryAppService() *{{.Name}}QueryAppService {
	once{{.Name}}.Do(func() {
		{{.name}}QueryAppService = new{{.Name}}QueryAppService()
	})
	return {{.name}}QueryAppService
}

//
// New{{.Name}}QueryAppService
// @Description: 创建{{.Name}}查询应用服务
// @return *{{.Name}}QueryAppService
//
func new{{.Name}}QueryAppService() *{{.Name}}QueryAppService {
	return &{{.Name}}QueryAppService{
        {{.name}}CreateExecutor:     {{.snake_name}}_impl.Get{{.Name}}CreateExecutor(),
        {{.name}}CreateManyExecutor: {{.snake_name}}_impl.Get{{.Name}}CreateManyExecutor(),

        {{.name}}UpdateExecutor:     {{.snake_name}}_impl.Get{{.Name}}UpdateExecutor(),
        {{.name}}UpdateManyExecutor: {{.snake_name}}_impl.Get{{.Name}}UpdateManyExecutor(),

        {{.name}}DeleteByIdExecutor: {{.snake_name}}_impl.Get{{.Name}}DeleteByIdExecutor(),
        {{.name}}DeleteManyExecutor: {{.snake_name}}_impl.Get{{.Name}}DeleteManyExecutor(),
        {{.name}}DeleteAllExecutor:  {{.snake_name}}_impl.Get{{.Name}}DeleteAllExecutor(),
        {{- if .IsEntity }}
        {{.name}}DeleteBy{{.AggregateName}}IdExecutor:  {{.snake_name}}_impl.Get{{.Name}}DeleteBy{{.AggregateName}}IdExecutor(),
        {{- end}}

		{{.name}}FindAllExecutor:    {{.snake_name}}_impl.Get{{.Name}}FindAllExecutor(),
		{{.name}}FindByIdExecutor:   {{.snake_name}}_impl.Get{{.Name}}FindByIdExecutor(),
		{{.name}}FindByIdsExecutor:  {{.snake_name}}_impl.Get{{.Name}}FindByIdsExecutor(),
		{{.name}}FindPagingExecutor: {{.snake_name}}_impl.Get{{.Name}}FindPagingExecutor(),
        {{- if .IsEntity }}
        {{.name}}FindBy{{.AggregateName}}IdExecutor:  {{.snake_name}}_impl.Get{{.Name}}FindBy{{.AggregateName}}IdExecutor(),
        {{- end}}
	}
}

//
// Create
// @Description: 创建{{.Name}}View
// @param ctx 上下文
// @param *view.{{.Name}}View {{.Name}}实体对象
// @return error 错误
//
func (a *{{.Name}}QueryAppService) Create(ctx context.Context, v *view.{{.Name}}View) error {
	return a.{{.name}}CreateExecutor.Execute(ctx, v)
}

//
// CreateMany
// @Description: 创建{{.Name}}View
// @param ctx
// @return []*view.{{.Name}}View  {{.Name}}实体对象切片
// @return error 错误
//
func (a *{{.Name}}QueryAppService) CreateMany(ctx context.Context, vList []*view.{{.Name}}View) error {
	return a.{{.name}}CreateManyExecutor.Execute(ctx, vList)
}

//
// Update
// @Description: 按ID更新{{.Name}}View
// @receiver a
// @param ctx
// @param v  *view.{{.Name}}View
// @return error 错误
//
func (a *{{.Name}}QueryAppService) Update(ctx context.Context, v *view.{{.Name}}View) error {
	return a.{{.name}}UpdateExecutor.Execute(ctx, v)
}

//
// UpdateMany
// @Description:  创建{{.Name}}View
// @param ctx
// @return []*view.{{.Name}}View  {{.Name}}实体对象切片
// @return error 错误
//
func (a *{{.Name}}QueryAppService) UpdateMany(ctx context.Context, vList []*view.{{.Name}}View) error {
	return a.{{.name}}UpdateManyExecutor.Execute(ctx, vList)
}

//
// DeleteById
// @Description: 按ID删除{{.Name}}View
// @param ctx
// @param tenantId 租户ID
// @param id 视图ID
// @return error 错误
//
func (a *{{.Name}}QueryAppService) DeleteById(ctx context.Context, tenantId, id string) error {
	return a.{{.name}}DeleteByIdExecutor.Execute(ctx, tenantId, id)
}

//
// DeleteMany
// @Description: 删除多个{{.Name}}View
// @param ctx
// @param tenantId 租户ID
// @param []*view.{{.Name}}View  {{.Name}}实体对象切片
// @return error 错误
//
func (a *{{.Name}}QueryAppService) DeleteMany(ctx context.Context, tenantId string, vList []*view.{{.Name}}View) error {
	return a.{{.name}}DeleteManyExecutor.Execute(ctx, tenantId, vList)
}

{{- if .IsEntity }}
//
// DeleteBy{{.AggregateName}}Id
// @Description: 删除多个{{.Name}}View
// @param ctx
// @param tenantId 租户ID
// @param []*view.{{.Name}}View  {{.Name}}实体对象切片
// @return error 错误
//
func (a *{{.Name}}QueryAppService) DeleteBy{{.AggregateName}}Id(ctx context.Context, tenantId string, {{.aggregateName}}Id string) error {
	return a.{{.name}}DeleteBy{{.AggregateName}}IdExecutor.Execute(ctx, tenantId, {{.aggregateName}}Id)
}
{{- end}}

//
// DeleteAll
// @Description:  删除所有
// @receiver a
// @param ctx
// @param tenantId 租户ID
// @return error
//
func (a *{{.Name}}QueryAppService) DeleteAll(ctx context.Context, tenantId string) error {
	return a.{{.name}}DeleteAllExecutor.Execute(ctx, tenantId)

}

//
// FindById
// @Description:  按ID查询{{.Name}}View
// @receiver a
// @param ctx
// @param qry 查询命令
// @return *view.{{.Name}}View
// @return bool 是否查询到数据
// @return error
//
func (a *{{.Name}}QueryAppService) FindById(ctx context.Context, tenantId string, id string) (*view.{{.Name}}View, bool, error) {
	qry := assembler.Ass{{.Name}}FindByIdQuery(tenantId, id)
	return a.{{.name}}FindByIdExecutor.Execute(ctx, qry)
}

//
// FindByIds
// @Description:  按多个ID查询{{.Name}}View
// @receiver a
// @param ctx
// @param qry 查询命令
// @return *view.{{.Name}}View
// @return bool 是否查询到数据
// @return error
//
func (a *{{.Name}}QueryAppService) FindByIds(ctx context.Context, tenantId string, ids []string) ([]*view.{{.Name}}View, bool, error) {
	qry := assembler.Ass{{.Name}}FindByIdsQuery(tenantId, ids)
	return a.{{.name}}FindByIdsExecutor.Execute(ctx, qry)
}

//
// FindAll
// @Description: 查询所有view.{{.Name}}View
// @receiver a
// @param ctx
// @param qry 查询命令
// @return []*view.{{.Name}}View
// @return bool 是否查询到数据
// @return error 错误
//
func (a *{{.Name}}QueryAppService) FindAll(ctx context.Context, tenantId string) ([]*view.{{.Name}}View, bool, error) {
	qry := assembler.Ass{{.Name}}FindAllQuery(tenantId)
	return a.{{.name}}FindAllExecutor.Execute(ctx, qry)
}

//
// FindPaging
// @Description: 分页查询
// @receiver a
// @param ctx 上下文
// @param qry 分页查询条件
// @return *appquery.{{.Name}}FindPagingResult 分页数据
// @return bool 是否查询到数据
// @return error 错误
//
func (a *{{.Name}}QueryAppService) FindPaging(ctx context.Context, aq *appquery.{{.Name}}FindPagingAppQuery) (*appquery.{{.Name}}FindPagingResult, bool, error) {
	return a.{{.name}}FindPagingExecutor.Execute(ctx, aq)
}

{{- if .IsEntity }}
//
// FindBy{{.AggregateName}}Id
// @Description: 根据{{.AggregateName}}Id查询
// @receiver a
// @param ctx 上下文
// @param tenantId string 租户ID
// @param {{.aggregateName}}Id  {{.AggregateName}}ID
// @return []*view.{{.Name}}View 实体切片
// @return bool 是否查询到数据
// @return error 错误
//
func (a *{{.Name}}QueryAppService) FindBy{{.AggregateName}}Id(ctx context.Context, tenantId string, {{.aggregateName}}Id string) ([]*view.{{.Name}}View, bool, error) {
	aq := assembler.Ass{{.Name}}FindBy{{.AggregateName}}IdQuery(tenantId, {{.aggregateName}}Id)
	return a.{{.name}}FindBy{{.AggregateName}}IdExecutor.Execute(ctx, aq)
}
{{- end }}