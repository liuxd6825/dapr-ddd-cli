package register
{{$namespace:=.Namespace}}
import (
{{- range $name, $agg := .Aggregates}}
    {{$agg.SnakeName}}_model "{{$namespace}}/pkg/cmd-service/domain/{{$agg.SnakeName}}/model"
{{- end}}
	"github.com/liuxd6825/dapr-go-ddd-sdk/ddd"
)

func init() {
{{- range $name, $agg := .Aggregates}}
	ddd.RegisterAggregateType({{$agg.SnakeName}}_model.AggregateType, {{$agg.SnakeName}}_model.NewAggregate)
{{- end}}
}
