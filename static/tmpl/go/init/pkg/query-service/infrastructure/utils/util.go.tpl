package utils

import (
    "context"
    {{- if .HasDateTimeType }}
    "time"
    {{- end}}
    "github.com/liuxd6825/dapr-go-ddd-sdk/mapper"
)

type ViewDefaultFields interface {
{{- range $name, $property := .DefaultViewProperties}}
{{- if not $property.IsArray}}
    Get{{$property.UpperName}}(){{$property.LanType}}   // 获取 {{$property.Description}}
    Set{{$property.UpperName}}({{$property.LanType}})   // 设置 {{$property.Description}}
{{- end}}
{{- end}}
}

type SetViewType int

const (
	SetViewCreated SetViewType = iota // 开始生成枚举值, 默认为0
	SetViewUpdated
	SetViewDeleted
	SetViewOther
)

const StringEmpty = ""

//
// ViewMapper
// @Description: 视图属性自动复制
// @param ctx 上下文
// @param toView 视图对象
// @param fromData Event.Data 事件数据对象
// @return error 错误
//
func ViewMapper(ctx context.Context, toView ViewObject, event ddd.DomainEvent, setType SetViewType) error {
	err := mapper.Mapper(event.GetData(), toView)
	if err != nil {
		return err
	}
	err = SetViewDefaultFields(ctx, toView, event.GetTime(), setType)
	if err != nil {
		return err
	}
	return nil
}

//
// SetViewDefaultFields
// @Description:      通过ctx上下文，设置view视图对象属性， 如从ctx中的Token信息服务
// @param ctx         上下文
// @param viewFields  view视图对象
// @return error      错误
//
func SetViewDefaultFields(ctx context.Context, viewObj ViewObject, setTime time.Time, setType SetViewType) error {
	if viewObj == nil {
		return nil
	}
	userName := "userName"
	userId := "userId"
	nowTime := &setTime
	if nowTime.IsZero() {
		t := time.Now()
		nowTime = &t
	}

	switch setType {
	case SetViewCreated:
		viewObj.SetCreatedName(userName)
		viewObj.SetCreatedId(userId)
		viewObj.SetCreatedTime(&setTime)

		viewObj.SetUpdatedName(userName)
		viewObj.SetUpdatedId(userId)
		viewObj.SetUpdatedTime(nowTime)

		viewObj.SetDeletedName(StringEmpty)
		viewObj.SetDeletedId(StringEmpty)
		viewObj.SetDeletedTime(nil)
		viewObj.SetIsDeleted(true)
		break
	case SetViewUpdated:
		viewObj.SetUpdatedName(userName)
		viewObj.SetUpdatedId(userId)
		viewObj.SetUpdatedTime(nowTime)
		break
	case SetViewDeleted:
		viewObj.SetDeletedName(userName)
		viewObj.SetDeletedId(userId)
		viewObj.SetDeletedTime(nowTime)
		viewObj.SetIsDeleted(true)
		break
	default:
		break
	}
	return nil
}
