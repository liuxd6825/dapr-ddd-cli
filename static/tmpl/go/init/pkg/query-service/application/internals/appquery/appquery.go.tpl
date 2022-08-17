package appquery

import "{{.Namespace}}/pkg/query-service/domain/{{.aggregate_name}}/view"

type {{.Name}}FindByIdAppQuery struct {
	TenantId string
	Id       string
}

func New{{.Name}}FindByIdAppQuery() *{{.Name}}FindByIdAppQuery {
	return &{{.Name}}FindByIdAppQuery{}
}

type {{.Name}}FindByIdsAppQuery struct {
	TenantId string
	Ids      []string
}

func New{{.Name}}FindByIdsAppQuery() *{{.Name}}FindByIdsAppQuery {
	return &{{.Name}}FindByIdsAppQuery{}
}

type {{.Name}}FindByAggregateIdAppQuery struct {
	TenantId    string
	AggregateId string
}

func New{{.Name}}FindByAggregateIdAppQuery() *{{.Name}}FindByAggregateIdAppQuery {
	return &{{.Name}}FindByAggregateIdAppQuery{}
}

type {{.Name}}FindPagingAppQuery struct {
	TenantId string
	PageNum  int64
	PageSize int64
	Filter   string
	Sort     string
	Fields   string
}

func New{{.Name}}FindPagingAppQuery() *{{.Name}}FindPagingAppQuery {
	return &{{.Name}}FindPagingAppQuery{}
}

type {{.Name}}FindPagingResult struct {
	Data       []*view.{{.Name}}View `json:"data"`
	TotalRows  int64             `json:"totalRows"`
	TotalPages int64             `json:"totalPages"`
	PageNum    int64             `json:"pageNum"`
	PageSize   int64             `json:"pageSize"`
	Filter     string            `json:"filter"`
	Sort       string            `json:"sort"`
	Error      error             `json:"-"`
	IsFound    bool              `json:"-"`
}

func New{{.Name}}FindPagingResult() *{{.Name}}FindPagingResult {
	return &{{.Name}}FindPagingResult{}
}

type {{.Name}}FindAllAppQuery struct {
	TenantId string
}

func New{{.Name}}FindAllAppQuery() *{{.Name}}FindAllAppQuery {
	return &{{.Name}}FindAllAppQuery{}
}

type {{.Name}}FindBy{{.AggregateName}}IdAppQuery struct {
	TenantId string
	{{.AggregateName}}Id string
}

func New{{.Name}}FindBy{{.AggregateName}}IdAppQuery() *{{.Name}}FindBy{{.AggregateName}}IdAppQuery {
	return &{{.Name}}FindBy{{.AggregateName}}IdAppQuery{}
}
