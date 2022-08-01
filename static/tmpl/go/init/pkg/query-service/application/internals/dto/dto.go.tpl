package dto

type FindByIdQueryDto struct {
	TenantId string `json:"tenantId"`
	Id       string `json:"id"`
}

type FindAllQueryDto struct {
	TenantId string `json:"tenantId"`
}

type FindPagingQueryDto struct {
	TenantId string
	Fields   string
	Filter   string
	Sort     string
	PageNum  int64
	PageSize int64
}
