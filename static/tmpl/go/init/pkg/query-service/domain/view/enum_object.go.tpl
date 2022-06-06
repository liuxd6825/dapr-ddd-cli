package view


type {{.Name}} int
{{- $Name:=.Name}}

const (
{{- range  $valueName, $value := .Values}}
    {{$Name}}{{$valueName}} {{$Name}} = {{$value.Value}}
{{- end }}
)

func (e {{.Name}}) String() string {
    res := "UNKNOWN"
    switch e {
    {{- range  $valueName, $value := .Values}}
    case {{$Name}}{{$value.Name}}:
        res = "{{$valueName}}"
    {{- end }}
    default:
        res = "UNKNOWN"
    }
    return res
}

func (e {{.Name}}) Title() string {
    res := "UNKNOWN"
    switch e {
    {{- range  $valueName, $value := .Values}}
    case {{$Name}}{{$value.Name}}:
        res = {{if $value.HasTitle}} "{{$value.Title}}" {{else}} "{{$value.Name}}" {{end}}
    {{- end }}
    default:
        return "UNKNOWN"
    }
    return res
}