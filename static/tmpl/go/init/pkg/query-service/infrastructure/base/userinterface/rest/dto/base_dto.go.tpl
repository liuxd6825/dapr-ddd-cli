package dto

import (
	"github.com/liuxd6825/dapr-go-ddd-sdk/types"
)

type FindByIdRequest struct {
	TenantId string
	Id       string
}

type FindByIdsRequest struct {
	TenantId string
	Ids      []string
}

type FindAllRequest struct {
	TenantId string
}

type FindPagingRequest struct {
	TenantId string
	PageNum  int64
	PageSize int64
	Filter   string
	Sort     string
	Fields   string
}

func (q *FindPagingRequest) GetTenantId() string { return q.TenantId }
func (q *FindPagingRequest) GetPageNum() int64   { return q.PageNum }
func (q *FindPagingRequest) GetPageSize() int64  { return q.PageSize }
func (q *FindPagingRequest) GetFilter() string   { return q.Filter }
func (q *FindPagingRequest) GetSort() string     { return q.Sort }
func (q *FindPagingRequest) GetFields() string   { return q.Fields }

func (q *FindPagingRequest) SetTenantId(value string) { q.TenantId = value }
func (q *FindPagingRequest) SetPageNum(value int64)   { q.PageNum = value }
func (q *FindPagingRequest) SetPageSize(value int64)  { q.PageSize = value }
func (q *FindPagingRequest) SetFilter(value string)   { q.Filter = value }
func (q *FindPagingRequest) SetSort(value string)     { q.Sort = value }
func (q *FindPagingRequest) SetFields(value string)   { q.Fields = value }

//
// FindPagingResponse
// @Description: FindPagingResponse
//
type FindPagingResponse[T any] struct {
	Data       []T   `json:"data"`
	TotalRows  int64  `json:"totalRows"`
	TotalPages int64  `json:"totalPages"`
	PageNum    int64  `json:"pageNum"`
	PageSize   int64  `json:"pageSize"`
	Filter     string `json:"filter"`
	Sort       string `json:"sort"`
}

func NewFindPagingResponse[T any]() *FindPagingResponse[T] {
	res := FindPagingResponse[T]{}
	res.InitData()
	return &res
}

func (r *FindPagingResponse[T]) Init() {
	r.InitData()
}

func (r *FindPagingResponse[T]) InitData() {
	data := make([]T, 0)
	r.Data = data
}

type BaseDto struct {
{{- range $name, $property := .Properties}}
    {{- if $property.IsArrayEntityType }}
    {{$property.UpperName}} []*{{$property.LanType}}Dto `json:"{{$property.LowerName}},omitempty"{{if $property.HasValidate}}  validate:"{{$property.Validate}}"{{- end}}`  // {{$property.Description}}
    {{- else if $property.IsEntityType}}
    {{$property.UpperName}} *{{$property.LanType}}Dto `json:"{{$property.LowerName}},omitempty"{{if $property.HasValidate}}  validate:"{{$property.Validate}}"{{- end}}`  // {{$property.Description}}
    {{- else if $property.IsDates }}
    {{$property.UpperName}} *types.JSONDate `json:"{{$property.LowerName}},omitempty"{{if $property.HasValidate}}  validate:"{{$property.Validate}}"{{- end}}`  // {{$property.Description}}
    {{- else if $property.IsTimes }}
    {{$property.UpperName}} *types.JSONTime `json:"{{$property.LowerName}},omitempty"{{if $property.HasValidate}}  validate:"{{$property.Validate}}"{{- end}}`  // {{$property.Description}}
    {{- else }}
    {{$property.UpperName}} {{$property.LanType}}`json:"{{$property.LowerName}},omitempty"{{if $property.HasValidate}}  validate:"{{$property.Validate}}"{{- end}}`  // {{$property.Description}}
    {{- end}}
{{- end}}
}