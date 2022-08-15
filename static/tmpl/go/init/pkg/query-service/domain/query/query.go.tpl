package query

import (
	"{{.Namespace}}/pkg/query-service/domain/{{.aggregate_name}}/view"
)

type {{.Name}}FindByIdQuery struct {
	TenantId string `json:"tenantId"`
	Id       string `json:"id"`
}

func New{{.Name}}FindByIdQuery(tenantId, id string) *{{.Name}}FindByIdQuery  {
    return &{{.Name}}FindByIdQuery{
        TenantId: tenantId,
        Id: id,
    }
}

type {{.Name}}FindByIdsQuery struct {
	TenantId  string `json:"tenantId"`
	Ids       []string `json:"ids"`
}

func New{{.Name}}FindByIdsQuery(tenantId string, ids []string) *{{.Name}}FindByIdsQuery  {
    return &{{.Name}}FindByIdsQuery{
        TenantId: tenantId,
        Ids: ids,
    }
}

type {{.Name}}FindAllQuery struct {
	TenantId string `json:"tenantId"`
}

func New{{.Name}}FindAllQuery(tenantId string) *{{.Name}}FindAllQuery  {
    return &{{.Name}}FindAllQuery{
        TenantId: tenantId,
    }
}

type {{.Name}}FindPagingQuery struct {
	TenantId string `json:"tenantId"`
	Fields   string `json:"fields"`
	Filter   string `json:"filter"`
	Sort     string `json:"sort"`
	PageNum  int64  `json:"pageNum"`
	PageSize int64  `json:"pageSize"`
}

func New{{.Name}}FindPagingQuery(tenantId string, fields string, filter string, sort string, pageNum int64, pageSize int64) *{{.Name}}FindPagingQuery  {
    return &{{.Name}}FindPagingQuery{
        TenantId : tenantId,
        Fields   : fields,
        Filter   : filter,
        Sort     : sort,
        PageNum  : pageNum,
        PageSize : pageSize,
    }
}

func (q *{{.Name}}FindPagingQuery) GetTenantId() string {
	return q.TenantId
}

func (q *{{.Name}}FindPagingQuery) GetFields() string {
	return q.Fields
}

func (q *{{.Name}}FindPagingQuery) GetFilter() string {
	return q.Filter
}

func (q *{{.Name}}FindPagingQuery) GetSort() string {
	return q.Sort
}

func (q *{{.Name}}FindPagingQuery) GetPageNum() int64 {
	return q.PageNum
}

func (q *{{.Name}}FindPagingQuery) GetPageSize() int64 {
	return q.PageSize
}

{{- if .IsEntity }}
type {{.Name}}FindBy{{.AggregateName}}IdQuery struct {
    TenantId string `json:"tenantId"`
    {{.AggregateName}}Id  string `json:"{{.aggregateName}}Id"`
}

func New{{.Name}}FindBy{{.AggregateName}}IdQuery(tenantId string, {{.aggregateName}}Id string) *{{.Name}}FindBy{{.AggregateName}}IdQuery {
    return &{{.Name}}FindBy{{.AggregateName}}IdQuery{
        TenantId: tenantId,
        {{.AggregateName}}Id: {{.aggregateName}}Id,
    }
}
{{- end }}

type {{.Name}}FindPagingResult struct {
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

func New{{.Name}}FindPagingResult() *{{.Name}}FindPagingResult {
	return &{{.Name}}FindPagingResult{}
}
