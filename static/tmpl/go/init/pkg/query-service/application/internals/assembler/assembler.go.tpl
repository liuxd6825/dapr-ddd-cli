package assembler

import (
	"{{.Namespace}}/pkg/query-service/application/internals/{{.aggregate_name}}/appquery"
	"{{.Namespace}}/pkg/query-service/domain/{{.aggregate_name}}/query"
)

type {{.name}}Assembler struct {
}

var {{.Name}} = {{.name}}Assembler{}

func (a {{.name}}Assembler) AssFindByIdAppQuery(tenantId, id string) *appquery.{{.Name}}FindByIdAppQuery {
	res := appquery.New{{.Name}}FindByIdAppQuery()
	res.TenantId = tenantId
	res.Id = id
	return res
}

func (a {{.name}}Assembler) AssFindByIdsAppQuery(tenantId string, ids []string) *appquery.{{.Name}}FindByIdsAppQuery {
	res := appquery.New{{.Name}}FindByIdsAppQuery()
	res.TenantId = tenantId
	res.Ids = ids
	return res
}

{{- if .IsEntity }}
func (a {{.name}}Assembler) AssFindBy{{.AggregateName}}IdAppQuery(tenantId string, {{.aggregateName}}Id string) *appquery.{{.Name}}FindBy{{.AggregateName}}IdAppQuery {
	res := appquery.New{{.Name}}FindBy{{.AggregateName}}IdAppQuery()
	res.TenantId = tenantId
	res.{{.AggregateName}}Id = {{.aggregateName}}Id
	return res
}
{{- end }}

func (a {{.name}}Assembler) AssFindAllAppQuery(tenantId string) *appquery.{{.Name}}FindAllAppQuery {
	return &appquery.{{.Name}}FindAllAppQuery{TenantId: tenantId}
}

func (a {{.name}}Assembler) AssFindPagingResult(fpr *query.{{.Name}}FindPagingResult) *appquery.{{.Name}}FindPagingResult {
	res := &appquery.{{.Name}}FindPagingResult{}
	res.Sort = fpr.Sort
	res.PageNum = fpr.PageNum
	res.PageSize = fpr.PageSize
	res.Filter = fpr.Filter
	res.Sort = fpr.Sort
	res.Data = fpr.Data
	res.Error = fpr.Error
	res.IsFound = fpr.IsFound
	res.TotalPages = fpr.TotalPages
	res.TotalRows = fpr.TotalRows
	return res
}
