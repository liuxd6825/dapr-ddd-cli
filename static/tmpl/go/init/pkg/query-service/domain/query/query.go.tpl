package query

import (
	"{{.Namespace}}/pkg/query-service/domain/{{.aggregate_name}}/view"
)

type FindByIdQuery struct {
	TenantId string `json:"tenantId"`
	Id       string `json:"id"`
}

func NewFindByIdQuery(tenantId, id string) *FindByIdQuery  {
    return &FindByIdQuery{
        TenantId: tenantId,
        Id: id,
    }
}

type FindByIdsQuery struct {
	TenantId  string `json:"tenantId"`
	Ids       []string `json:"ids"`
}

func NewFindByIdsQuery(tenantId string, ids []string) *FindByIdsQuery  {
    return &FindByIdsQuery{
        TenantId: tenantId,
        Ids: ids,
    }
}

type FindAllQuery struct {
	TenantId string `json:"tenantId"`
}

func NewFindAllQuery(tenantId string) *FindAllQuery  {
    return &FindAllQuery{
        TenantId: tenantId,
    }
}

type FindPagingQuery struct {
	TenantId string `json:"tenantId"`
	Fields   string `json:"fields"`
	Filter   string `json:"filter"`
	Sort     string `json:"sort"`
	PageNum  int64  `json:"pageNum"`
	PageSize int64  `json:"pageSize"`
}

func NewFindPagingQuery(tenantId string, fields string, filter string, sort string, pageNum int64, pageSize int64) *FindPagingQuery  {
    return &FindPagingQuery{
        TenantId : tenantId,
        Fields   : fields,
        Filter   : filter,
        Sort     : sort,
        PageNum  : pageNum,
        PageSize : pageSize,
    }
}

func (q *FindPagingQuery) GetTenantId() string {
	return q.TenantId
}

func (q *FindPagingQuery) GetFields() string {
	return q.Fields
}

func (q *FindPagingQuery) GetFilter() string {
	return q.Filter
}

func (q *FindPagingQuery) GetSort() string {
	return q.Sort
}

func (q *FindPagingQuery) GetPageNum() int64 {
	return q.PageNum
}

func (q *FindPagingQuery) GetPageSize() int64 {
	return q.PageSize
}

{{- if .IsEntity }}
type FindBy{{.AggregateName}}IdQuery struct {
    TenantId string `json:"tenantId"`
    {{.AggregateName}}Id  string `json:"{{.aggregateName}}Id"`
}

func NewFindBy{{.AggregateName}}IdQuery(tenantId string, {{.aggregateName}}Id string) *FindBy{{.AggregateName}}IdQuery {
    return &FindBy{{.AggregateName}}IdQuery{
        TenantId: tenantId,
        {{.AggregateName}}Id: {{.aggregateName}}Id,
    }
}
{{- end }}

type FindPagingResult struct {
	Data       []*view.{{.Name}}View `json:"data"`
	TotalRows  int64                `json:"totalRows"`
	TotalPages int64                `json:"totalPages"`
	PageNum    int64                `json:"pageNum"`
	PageSize   int64                `json:"pageSize"`
	Filter     string               `json:"filter"`
	Sort       string               `json:"sort"`
	Error      error                `json:"-"`
	IsFound    bool                 `json:"-"`
}

func NewFindPagingResult() *FindPagingResult {
	return &FindPagingResult{}
}
