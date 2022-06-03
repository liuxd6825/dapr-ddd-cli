package dto

import (
)

//
// GetRequest
// @Description:
//
type GetRequest struct {
    TenantId   string `json:"tenantId"  validate:"gt=0"`
}


//
// GetByIdRequest
// @Description:
//
type GetByIdRequest struct {
    GetRequest
    Id string   `json:"id"`
}

//
// GetByIdResponse
// @Description:
//
type GetByIdResponse struct {

}


//
// GetAllRequest
// @Description: GetPagingResponse
//
type GetAllRequest struct {
    GetRequest
}


//
// GetAllResponse
// @Description: GetPagingResponse
//
type GetAllResponse[T any]  List[T]


//
// GetPagingRequest
// @Description: GetPagingRequest
//
type GetPagingRequest[T any] struct {
    GetRequest
    PageNum    int64  `json:"pageNum"`
    PageSize   int64  `json:"pageSize"`
    Filter     string `json:"filter"`
    Sort       string `json:"sort"`
}


//
// GetPagingResponse
// @Description: GetPagingResponse
//
type GetPagingResponse[T any] struct {
    Data       *[]any   `json:"data"`
    TotalRows  int64  `json:"totalRows"`
    TotalPages int64  `json:"totalPages"`
    PageNum    int64  `json:"pageNum"`
    PageSize   int64  `json:"pageSize"`
    Filter     string `json:"filter"`
    Sort       string `json:"sort"`
}


//
// List
// @Description: List
//
type List[T any]  []T







