package model

import (
	"context"
	"encoding/json"
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

func (i *{{.Name}}Items) MarshalJSON() ([]byte, error) {
    return json.Marshal(i.Items.MapData())
}

func (i *{{.Name}}Items) UnmarshalJSON(b []byte) error {
	data := i.Items.MapData()
	return json.Unmarshal(b, &data)
}

func (i *{{.Name}}Items) AddItem(ctx context.Context, row *{{.Name}}) error {
	return i.Items.AddItem(ctx, row)
}

func (i *{{.Name}}Items) DeleteItem(ctx context.Context, row *{{.Name}}) error {
	return i.Items.Delete(ctx, row)
}

func (i *{{.Name}}Items) UpdateItem(ctx context.Context, row *{{.Name}}) error {
	return i.Items.UpdateItem(ctx, row)
}

func (i *{{.Name}}Items) MapData() map[string]*{{.Name}} {
	return i.Items.MapData()
}

func (i *{{.Name}}Items) Length() int {
	m := i.MapData()
	return len(m)
}
