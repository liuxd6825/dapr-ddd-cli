package  {{.aggregate_name}}_model
{{$namespace:=.Namespace}}
import (
{{- range $name, $agg := .Aggregates}}
    "{{$namespace}}/pkg/cmd-service/domain/{{$agg.LowerName}}_model"
{{- end}}
	"github.com/dapr/dapr-go-ddd-sdk/ddd"
)

func init() {
{{- range $name, $agg := .Aggregates}}
	ddd.RegisterAggregateType({{$agg.LowerName}}_model.AggregateType, {{$agg.LowerName}}_model.NewAggregate)
{{- end}}
}
