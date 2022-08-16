package {{.snake_name}}_impl
{{ $AggregateName := .AggregateName }}
import (
	"github.com/liuxd6825/dapr-go-ddd-sdk/utils/singleutils"
	"{{.Namespace}}/pkg/query-service/application/internals/{{.aggregate_name}}/executor"
)

func Get{{.Name}}FindByIdExecutor() executor.{{.Name}}FindByIdExecutor {
	return singleutils.GetObject[*{{.name}}FindByIdExecutor]()
}

func Get{{.Name}}FindByIdsExecutor() executor.{{.Name}}FindByIdsExecutor {
	return singleutils.GetObject[*{{.name}}FindByIdsExecutor]()
}

func Get{{.Name}}FindPagingExecutor() executor.{{.Name}}FindPagingExecutor {
	return singleutils.GetObject[*{{.name}}FindPagingExecutor]()
}

func Get{{.Name}}FindAllExecutor() executor.{{.Name}}FindAllExecutor {
	return singleutils.GetObject[*{{.name}}FindAllExecutor]()
}

func Get{{.Name}}CreateExecutor() executor.{{.Name}}CreateExecutor {
	return singleutils.GetObject[*{{.name}}CreateExecutor]()
}

func Get{{.Name}}CreateManyExecutor() executor.{{.Name}}CreateManyExecutor {
	return singleutils.GetObject[*{{.name}}CreateManyExecutor]()
}

func Get{{.Name}}UpdateExecutor() executor.{{.Name}}UpdateExecutor {
	return singleutils.GetObject[*{{.name}}UpdateExecutor]()
}

func Get{{.Name}}UpdateManyExecutor() executor.{{.Name}}UpdateManyExecutor {
	return singleutils.GetObject[*{{.name}}UpdateManyExecutor]()
}

func Get{{.Name}}DeleteByIdExecutor() executor.{{.Name}}DeleteByIdExecutor {
	return singleutils.GetObject[*{{.name}}DeleteByIdExecutor]()
}

func Get{{.Name}}DeleteManyExecutor() executor.{{.Name}}DeleteManyExecutor {
	return singleutils.GetObject[*{{.name}}DeleteManyExecutor]()
}

func Get{{.Name}}DeleteAllExecutor() executor.{{.Name}}DeleteAllExecutor {
	return singleutils.GetObject[*{{.name}}DeleteAllExecutor]()
}

func init() {
	if err := singleutils.Set[*{{.name}}FindByIdExecutor](func() *{{.name}}FindByIdExecutor { return new{{.Name}}FindByIdExecutor() }); err != nil {
		panic(err)
	}
	if err := singleutils.Set[*{{.name}}FindByIdsExecutor](func() *{{.name}}FindByIdsExecutor { return new{{.Name}}FindByIdsExecutor() }); err != nil {
		panic(err)
	}
	if err := singleutils.Set[*{{.name}}FindPagingExecutor](func() *{{.name}}FindPagingExecutor { return new{{.Name}}FindPagingExecutor() }); err != nil {
		panic(err)
	}
	if err := singleutils.Set[*{{.name}}FindAllExecutor](func() *{{.name}}FindAllExecutor { return new{{.Name}}FindAllExecutor() }); err != nil {
		panic(err)
	}
	if err := singleutils.Set[*{{.name}}CreateExecutor](func() *{{.name}}CreateExecutor { return new{{.Name}}CreateExecutor() }); err != nil {
		panic(err)
	}
	if err := singleutils.Set[*{{.name}}CreateManyExecutor](func() *{{.name}}CreateManyExecutor { return new{{.Name}}CreateManyExecutor() }); err != nil {
		panic(err)
	}
	if err := singleutils.Set[*{{.name}}UpdateExecutor](func() *{{.name}}UpdateExecutor { return new{{.Name}}UpdateExecutor() }); err != nil {
		panic(err)
	}
	if err := singleutils.Set[*{{.name}}UpdateManyExecutor](func() *{{.name}}UpdateManyExecutor { return new{{.Name}}UpdateManyExecutor() }); err != nil {
		panic(err)
	}
	if err := singleutils.Set[*{{.name}}DeleteByIdExecutor](func() *{{.name}}DeleteByIdExecutor { return new{{.Name}}DeleteByIdExecutor() }); err != nil {
		panic(err)
	}
	if err := singleutils.Set[*{{.name}}DeleteManyExecutor](func() *{{.name}}DeleteManyExecutor { return new{{.Name}}DeleteManyExecutor() }); err != nil {
		panic(err)
	}
	if err := singleutils.Set[*{{.name}}DeleteAllExecutor](func() *{{.name}}DeleteAllExecutor { return new{{.Name}}DeleteAllExecutor() }); err != nil {
		panic(err)
	}
}
