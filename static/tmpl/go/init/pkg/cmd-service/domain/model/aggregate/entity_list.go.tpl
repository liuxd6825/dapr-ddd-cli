package model

import (
	"context"
	"errors"
	"fmt"
	"{{.Namespace}}/pkg/cmd-service/domain/{{.aggregate_name}}/field"
	"github.com/liuxd6825/dapr-go-ddd-sdk/mapper"
)

type {{.ClassName}} map[string]*{{.Entity.Name}}

//
// AddByField
// @Description: 添加
// @receiver s
// @param ctx    上下文
// @param data   更新数据
// @return error 错误
//
func (s *{{.ClassName}}) AddByField(ctx context.Context, data *field.{{.Entity.Name}}Fields) error {
	m := *s
	item, ok := m[data.Id]
	if ok {
		return errors.New(fmt.Sprintf("新建 {{.Entity.Name}}.Id \"%s\" 已经存在", data.Id))
	}
	item = &{{.Entity.Name}}{}
	err := mapper.Mapper(data, item)
	if err!=nil {
	    m[data.Id] = item
	}
	return err
}

//
// UpdateByField
// @Description:     更新
// @receiver s
// @param ctx        上下文
// @param data       更新数据
// @param updateMask 更新字段项
// @return error
//
func (s *{{.ClassName}}) UpdateByField(ctx context.Context, data *field.{{.Entity.Name}}Fields, updateMask []string) error {
	m := *s
	item, ok := m[data.Id]
	if !ok {
		return nil
	}
	return mapper.MaskMapper(data, item, updateMask)
}

//
// DeleteById
// @Description:
// @receiver s
// @param ctx     上下文
// @param id      Id主键
// @return error  错误
//
func (s *{{.ClassName}}) DeleteById(ctx context.Context, id string) error {
	m := *s
	delete(m, id)
	return nil
}

//
// DeleteByIds
// @Description:  按id删除多个
// @receiver s
// @param ctx     上下文
// @param id      Id主键
// @return error  错误
//
func (s *{{.ClassName}}) DeleteByIds(ctx context.Context, ids ...string) error {
	m := *s
	if len(ids) > 0 {
		for _, id := range ids {
			delete(m, id)
		}
	}
	return nil
}
