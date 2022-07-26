package executor

import "github.com/liuxd6825/dapr-go-ddd-sdk/utils/singleutils"

{{- range $i, $cmd := .Commands}}

func Get{{.Name}}Executor() {{$cmd.Name}}Executor {
    return singleutils.GetObject[*{{$cmd.FirstLowerName}}Executor]()
}
{{- end }}

func GetFindAggregateByIdExecutor() FindAggregateByIdExecutor {
    return singleutils.GetObject[*findAggregateByIdExecutor]()
}

func init(){
    {{- range $i, $cmd := .Commands}}
    if err := singleutils.Set[*{{$cmd.FirstLowerName}}Executor](func()*{{$cmd.FirstLowerName}}Executor{ return new{{$cmd.Name}}Executor()}); err!=nil {
        panic(err)
    }
    {{- end }}
    if err := singleutils.Set[*findAggregateByIdExecutor](func()*findAggregateByIdExecutor{ return newFindAggregateByIdExecutor()}); err!=nil {
        panic(err)
    }
}

