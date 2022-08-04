package command

type {{.Name}}FindByIdQuery struct {
	TenantId string `json:"tenantId"`
	Id       string `json:"id"`
}

func New{{.Name}}FindByIdQuery() *{{.Name}}FindByIdQuery  {
    return &{{.Name}}FindByIdQuery{}
}

type {{.Name}}FindByIdsQuery struct {
	TenantId  string `json:"tenantId"`
	Ids       []string `json:"ids"`
}

func New{{.Name}}FindByIdsQuery() *{{.Name}}FindByIdsQuery  {
    return &{{.Name}}FindByIdsQuery{}
}

type {{.Name}}FindAllQuery struct {
	TenantId string `json:"tenantId"`
}

func New{{.Name}}FindAllQuery() *{{.Name}}FindAllQuery  {
    return &{{.Name}}FindAllQuery{}
}

type {{.Name}}FindPagingQuery struct {
	TenantId string `json:"tenantId"`
	Fields   string `json:"fields"`
	Filter   string `json:"filter"`
	Sort     string `json:"sort"`
	PageNum  int64  `json:"pageNum"`
	PageSize int64  `json:"pageSize"`
}

func New{{.Name}}FindPagingQuery() *{{.Name}}FindPagingQuery  {
    return &{{.Name}}FindPagingQuery{}
}

{{- if .IsEntity }}
type {{.Name}}FindBy{{.AggregateName}}IdQuery struct {
    TenantId string `json:"tenantId"`
    {{.AggregateName}}Id  string `json:"{{.aggregateName}}Id"`
}

func New{{.Name}}FindBy{{.AggregateName}}IdQuery() *{{.Name}}1111FindBy{{.AggregateName}}IdQuery {
    return &{{.Name}}FindBy{{.AggregateName}}IdQuery{}
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
