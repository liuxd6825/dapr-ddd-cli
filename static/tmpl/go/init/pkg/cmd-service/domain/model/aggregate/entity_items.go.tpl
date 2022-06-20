package model

import (
	"github.com/liuxd6825/dapr-go-ddd-sdk/types"
)

type {{.ClassName}} struct {
    types.Items[*{{.Entity.Name}}]
}

func New{{.ClassName}}() *{{.ClassName}}{
    res := &{{.ClassName}}{}
    res.Init(func() interface{} {
        return &{{.Entity.Name}}{}
    })
	return res
}

