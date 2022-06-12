package utils

import (
    "context"
    {{- if .HasTimeType }}
    "time"
    {{- end}}
)

type ViewDefaultFields interface {
{{- range $name, $property := .DefaultViewProperties}}
{{- if not $property.IsArray}}
    Get{{$property.UpperName}}(){{$property.LanType}}   // 获取 {{$property.Description}}
    Set{{$property.UpperName}}({{$property.LanType}})   // 设置 {{$property.Description}}
{{- end}}
{{- end}}
}


func SetViewDefaultFields(ctx context.Context, viewFields ViewDefaultFields) {

}

/*
type FindRequest interface {
    SetTenantId(value string)
}

func SetFindRequest( ctx iris.Context, r *FindRequest) error {
    tenantId := ctx.Params().Get("tenantId")
	r.SetTenantId(tenantId)
	return assert.NotEmpty(tenantId, assert.NewOptions("url \"{tenantId}\" cannot be empty"))
}

type FindByIdRequest interface {
    FindRequest
    SetId(value string)
}

func SetFindByIdRequest(ctx iris.Context, r *FindByIdRequest) error {
    if err!=SetFindRequest(ctx, r); err!=nil{
        return err
    }
    id := ctx.Params().Get("id")
    r.SetId(id)
    return assert.NotEmpty(id, assert.NewOptions("url \"{id}\" cannot be empty"))
}


type FindAllRequest interface {
    FindRequest
}

func SetFindAllRequest(ctx iris.Context, r *FindAllRequest) error {
    SetFindRequest(ctx, r)
}


type FindPagingRequest interface {
    FindRequest
    
    GetPageNum() int64
    SetPageNum(value int64)
    
    GetPageSize()
    SetPageSize(value int64)
    
    GetFilter() string
    SetFilter(value string)
    
    GetSort() string
    SetSort(value string)
    
    GetFields() string
    SetFields(value string)
}

func SetFindPagingRequest(ctx iris.Context, r *FindPagingRequest) error {
    if err!=SetFindRequest(ctx, r); err!=nil{
        return err
    }
	r.SetPageNum(ctx.URLParamInt64Default("pageNum", 0))
	r.SetPageSize(ctx.URLParamInt64Default("pageSize", 20))
	r.SetFilter(ctx.URLParamDefault("filter", ""))
	r.SetSort(ctx.URLParamDefault("sort", ""))
}

*/