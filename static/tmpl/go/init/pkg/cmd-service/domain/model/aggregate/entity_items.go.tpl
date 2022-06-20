package model

import (
	"github.com/liuxd6825/dapr-go-ddd-sdk/types"
)

type {{.Name}}Items struct {
    types.Items[*{{.Entity.Name}}]
}

func New{{.Name}}Items() *{{.Name}}Items{
    res := &{{.Name}}Items{}
    res.Init(func() interface{} {
        return &{{.Entity.Name}}{}
    })
	return res
}

