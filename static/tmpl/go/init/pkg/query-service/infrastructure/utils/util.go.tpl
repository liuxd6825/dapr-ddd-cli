package utils

import (
    "context"
    "time"
    "github.com/liuxd6825/dapr-go-ddd-sdk/types"
    "github.com/liuxd6825/dapr-go-ddd-sdk/mapper"
)

type ViewObject interface {
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
	err := Mapper(event.GetData(), toView)
	if err != nil {
		return err
	}
	err = SetViewDefaultFields(ctx, toView, event.GetCreatedTime(), setType)
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
		viewObj.SetCreatorName(userName)
		viewObj.SetCreatorId(userId)
		viewObj.SetCreatedTime(&setTime)

		viewObj.SetUpdaterName(userName)
		viewObj.SetUpdaterId(userId)
		viewObj.SetUpdatedTime(nowTime)

		viewObj.SetDeleterName(StringEmpty)
		viewObj.SetDeleterId(StringEmpty)
		viewObj.SetDeletedTime(nil)
		viewObj.SetIsDeleted(true)
		break
	case SetViewUpdated:
		viewObj.SetUpdaterName(userName)
		viewObj.SetUpdaterId(userId)
		viewObj.SetUpdatedTime(nowTime)
		break
	case SetViewDeleted:
		viewObj.SetDeleterName(userName)
		viewObj.SetDeleterId(userId)
		viewObj.SetDeletedTime(nowTime)
		viewObj.SetIsDeleted(true)
		break
	default:
		break
	}
	return nil
}


//
// Mapper
// @Description: 进行struct属性复制，支持深度复制
// @param fromObj 来源
// @param toObj 目标
// @return error
//
func Mapper(fromObj, toObj interface{}) error {
	return types.Mapper(fromObj, toObj)
}

//
// MaskMapper
// @Description: 根据指定进行属性复制，不支持深度复制
// @param fromObj 来源
// @param toObj 目标
// @param mask 要复制属性列表
// @return error
//
func MaskMapper(fromObj, toObj interface{}, mask []string) error {
	return types.MaskMapper(fromObj, toObj, mask)
}
