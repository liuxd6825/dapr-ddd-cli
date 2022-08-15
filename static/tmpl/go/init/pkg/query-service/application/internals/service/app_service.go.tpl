package service
{{$AggregateName := .AggregateName}}
import (
    "sync"
	"context"
	"{{.Namespace}}/pkg/query-service/domain/{{.aggregate_name}}/service"
    "{{.Namespace}}/pkg/query-service/domain/{{.aggregate_name}}/view"
    "{{.Namespace}}/pkg/query-service/domain/{{.aggregate_name}}/query"
    "{{.Namespace}}/pkg/query-service/infrastructure/db/session"
    "{{.Namespace}}/pkg/query-service/infrastructure/domain_impl/{{.aggregate_name}}/service_impl"
)

//
// {{.Name}}QueryAppService
// @Description: {{.Description}}查询应用服务类
//
type {{.Name}}QueryAppService struct {
	{{.name}}DomainService service.{{.Name}}QueryDomainService
	{{- if .IsAggregate }}
    {{- range $entityName, $entity := .Aggregate.Entities}}
    {{$entity.FirstLowerName}}DomainService service.{{$entity.Name}}QueryDomainService
    {{- end }}
    {{- end }}
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
		{{.name}}DomainService: service_impl.Get{{.Name}}QueryDomainService(),
		{{- if .IsAggregate }}
        {{- range $entityName, $entity := .Aggregate.Entities}}
        {{$entity.FirstLowerName}}DomainService : service_impl.Get{{$entity.Name}}QueryDomainService(),
        {{- end }}
        {{- end }}
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
	return session.StartSession(ctx, func(ctx context.Context) error {
	    return a.{{.name}}DomainService.Create(ctx, v)
	})
}

//
// CreateMany
// @Description: 创建{{.Name}}View
// @param ctx
// @return []*view.{{.Name}}View  {{.Name}}实体对象切片
// @return error 错误
//
func (a *{{.Name}}QueryAppService) CreateMany(ctx context.Context, vList []*view.{{.Name}}View) error {
	return session.StartSession(ctx, func(ctx context.Context) error {
	    return a.{{.name}}DomainService.CreateMany(ctx, vList)
	})

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
	return session.StartSession(ctx, func(ctx context.Context) error {
	    return a.{{.name}}DomainService.Update(ctx, v)
	})
}


//
// UpdateMany
// @Description:  创建{{.Name}}View
// @param ctx
// @return []*view.{{.Name}}View  {{.Name}}实体对象切片
// @return error 错误
//
func (a *{{.Name}}QueryAppService) UpdateMany(ctx context.Context, vList []*view.{{.Name}}View) error {
	return session.StartSession(ctx, func(ctx context.Context) error {
	    return a.{{.name}}DomainService.UpdateMany(ctx, vList)
	})
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
	return session.StartSession(ctx, func(ctx context.Context) error {
        {{- if .IsAggregate }}
        {{- range $entityName, $entity := .Aggregate.Entities}}
        if err:= a.{{$entity.FirstLowerName}}DomainService.DeleteBy{{$AggregateName}}Id(ctx, tenantId, id); err!=nil {
            return err
        }
        {{- end }}
        {{- end }}
        return a.{{.name}}DomainService.DeleteById(ctx, tenantId, id)
	})
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
	return session.StartSession(ctx, func(ctx context.Context) error {
        {{- if .IsAggregate }}
        {{- if not .Aggregate.Entities.Empty }}
        for _, item := range vList {
            {{- range $entityName, $entity := .Aggregate.Entities}}
            if err:= a.{{$entity.FirstLowerName}}DomainService.DeleteBy{{$AggregateName}}Id(ctx, tenantId, item.Id); err!=nil {
                return err
            }
            {{- end }}
        }
        {{- end }}
        {{- end }}
        return a.{{.name}}DomainService.DeleteMany(ctx, tenantId, v)
	})
}

//
// DeleteAll
// @Description:  删除所有
// @receiver a
// @param ctx
// @param tenantId 租户ID
// @return error
//
func (a *{{.Name}}QueryAppService) DeleteAll(ctx context.Context, tenantId string) error {
	return session.StartSession(ctx, func(ctx context.Context) error {
        {{- if .IsAggregate }}
        {{- range $entityName, $entity := .Aggregate.Entities}}
        if err:= a.{{$entity.FirstLowerName}}DomainService.DeleteAll(ctx, tenantId); err!=nil {
            return err
        }
        {{- end }}
        {{- end }}
        return a.{{.name}}DomainService.DeleteAll(ctx, tenantId)
	})
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
	qry := query.New{{.Name}}FindByIdQuery(tenantId, id)
	return a.{{.name}}DomainService.FindById(ctx, qry)
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
	qry := query.New{{.Name}}FindAllQuery(tenantId)
	return a.{{.name}}DomainService.FindAll(ctx, qry)
}


//
// FindPaging
// @Description: 分页查询
// @receiver a
// @param ctx 上下文
// @param qry 分页查询条件
// @return *query.FindPagingResult 分页数据
// @return bool 是否查询到数据
// @return error 错误
//
func (a *{{.Name}}QueryAppService) FindPaging(ctx context.Context, qry *query.{{.Name}}FindPagingQuery) (*query.{{.Name}}FindPagingResult, bool, error) {
	return a.{{.name}}DomainService.FindPaging(ctx, qry)
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
	qry := query.New{{.Name}}FindBy{{.AggregateName}}IdQuery(tenantId, {{.aggregateName}}Id)
	return a.{{.name}}DomainService.FindBy{{.AggregateName}}Id(ctx, qry)
}
{{- end }}