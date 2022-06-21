package config

import "fmt"

type Any map[string]interface{}

type Metadata map[string]interface{}

type MetadataUtil struct {
	mapValue Metadata
}

func NewMetadataUtil(mapValue Metadata) *MetadataUtil {
	return &MetadataUtil{
		mapValue: mapValue,
	}
}

func (u *MetadataUtil) GetValue2(key string) (string, bool) {
	if u == nil {
		var s string
		return s, false
	}
	value, ok := u.mapValue[key]
	return fmt.Sprintf("%s", value), ok
}

//
// GetValue
// @Description: 获取值
// @receiver u
// @param key 关键字
// @param defaultValue 默认
// @return string
//
func (u *MetadataUtil) GetValue(key string, defaultValue string) string {
	v, ok := u.GetValue2(key)
	if ok {
		return v
	}
	return defaultValue
}

func (u *MetadataUtil) Namespace() string {
	return u.GetValue("namespace", "{{.namespace}}")
}
