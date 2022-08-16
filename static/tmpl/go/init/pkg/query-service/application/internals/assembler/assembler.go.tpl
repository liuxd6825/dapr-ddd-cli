package assembler

import (
	"{{.Namespace}}/pkg/query-service/application/internals/{{.aggregate_name}}/appquery"
	"{{.Namespace}}/pkg/query-service/domain/{{.aggregate_name}}/query"
)

func Ass{{.Name}}FindByIdQuery(tenantId, id string) *appquery.{{.Name}}FindByIdAppQuery {
	res := appquery.New{{.Name}}FindByIdAppQuery()
	res.TenantId = tenantId
	res.Id = id
	return res
}

func Ass{{.Name}}FindByIdsQuery(tenantId string, ids []string) *appquery.{{.Name}}FindByIdsAppQuery {
	res := appquery.New{{.Name}}FindByIdsAppQuery()
	res.TenantId = tenantId
	res.Ids = ids
	return res
}

func Ass{{.Name}}FindAllQuery(tenantId string) *appquery.{{.Name}}FindAllAppQuery {
	return &appquery.{{.Name}}FindAllAppQuery{TenantId: tenantId}
}

func Ass{{.Name}}FindPagingResult(fpr *query.{{.Name}}FindPagingResult) *appquery.{{.Name}}FindPagingResult {
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
