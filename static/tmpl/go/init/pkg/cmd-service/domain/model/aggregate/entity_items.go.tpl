package model

import (
	"context"
	"errors"
	"fmt"
	"{{.Namespace}}/pkg/cmd-service/domain/{{.aggregate_name}}/field"
	"github.com/liuxd6825/dapr-go-ddd-sdk/mapper"
	"github.com/liuxd6825/dapr-go-ddd-sdk/types"
)

type {{.ClassName}} struct {
    types.Items[*{{.Entity.Name}}]
}

//
// Add
// @Description: 添加
// @param ctx    上下文
// @param data   更新数据
// @return error 错误
//
func (t *{{.ClassName}}) Add(ctx context.Context, id string, data *field.{{.Entity.Name}}Fields) error {
	return t.Items.Add(ctx, id, data)
}

//
// Update
// @Description:     更新
// @param ctx        上下文
// @param data       更新数据
// @param updateMask 更新字段项
// @return error
//
func (t {{.ClassName}}) Update(ctx context.Context, id string, data *field.{{.Entity.Name}}Fields, updateMask []string) error {
	return t.Items.Update(ctx, id, data,updateMask)
}

//
// Delete
// @Description: 删除明细
// @param ctx    上下文
// @param item   明细对象
// @return error 错误
//
func (t {{.ClassName}}) Delete(ctx context.Context, item *{{.Entity.Name}}) error {
	return t.Items.Delete(ctx, item)
}

//
// DeleteById
// @Description: 按Id删除
// @param ctx    上下文
// @param id     Id主键
// @return error 错误
//
func (t {{.ClassName}}) DeleteById(ctx context.Context, id string) error {
	return t.Items.DeleteById(ctx, id)
}

//
// DeleteByIds
// @Description:  按id删除多个
// @receiver s
// @param ctx     上下文
// @param id      Id主键
// @return error  错误
//
func (t {{.ClassName}}) DeleteByIds(ctx context.Context, ids ...string) error {
    return t.Items.DeleteByIds(ctx, ids...)
}
