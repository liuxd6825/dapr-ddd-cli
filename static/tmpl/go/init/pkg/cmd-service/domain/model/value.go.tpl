package model

import (
	"fmt"
	"{{.Namespace}}/pkg/cmd-service/domain/model/{{.aggregateName}}_models/user_fields"
)

type AddressValue struct {
	Id       string `json:"id"`
	Province string `json:"province"`
	City     string `json:"city"`
	Area     string `json:"area"`
	Location string `json:"location"`
}

func NewAddressValue(fields *user_fields.AddressFields) *AddressValue {
	return &AddressValue{
		Id:       fields.Id,
		Province: fields.Province,
		City:     fields.City,
		Area:     fields.Area,
		Location: fields.Location,
	}
}

func (a *AddressValue) ToString() string {
	return fmt.Sprintf("%s省 %s市 %s区 详细：%s", a.Province, a.City, a.Area, a.Location)
}
