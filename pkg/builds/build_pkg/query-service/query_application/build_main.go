package query_application

import (
	"fmt"
	"github.com/liuxd6825/dapr-ddd-cli/pkg/builds"
	"github.com/liuxd6825/dapr-ddd-cli/pkg/config"
	"github.com/liuxd6825/dapr-ddd-cli/pkg/utils"
	"strings"
)

type BuildApplicationLayer struct {
	builds.BaseBuild
	aggregate *config.Aggregate
	outDir    string
}

func NewBuildApplicationLayer(cfg *config.Config, aggregate *config.Aggregate, outDir string) *BuildApplicationLayer {
	res := &BuildApplicationLayer{
		BaseBuild: builds.BaseBuild{
			Config:    cfg,
			Aggregate: aggregate,
		},
		aggregate: aggregate,
		outDir:    outDir,
	}
	res.initQueryAppService()
	res.initQueryHandlerAggregate()
	res.initQueryHandlerEntities()
	res.initExecutor()
	res.initExecutorImpls()
	res.initAppQuery()
	res.initAssembler()
	return res
}

func (b *BuildApplicationLayer) initAssembler() {
	outFile := fmt.Sprintf("%s/internals/%s/assembler/%s_assembler.go", b.outDir, b.aggregate.FileName(), b.aggregate.FileName())
	build := NewBuildAssembler(b.BaseBuild, nil, utils.ToLower(outFile))
	b.AddBuild(build)
	for _, item := range b.aggregate.Entities {
		outFile := fmt.Sprintf("%s/internals/%s/assembler/%s_assembler.go", b.outDir, b.aggregate.FileName(), item.FileName())
		build := NewBuildAssembler(b.BaseBuild, item, utils.ToLower(outFile))
		b.AddBuild(build)
	}
}

func (b *BuildApplicationLayer) initAppQuery() {
	outFile := fmt.Sprintf("%s/internals/%s/appquery/%s_appquery.go", b.outDir, b.aggregate.FileName(), b.aggregate.FileName())
	build := NewBuildAppQuery(b.BaseBuild, nil, utils.ToLower(outFile))
	b.AddBuild(build)
	for _, item := range b.aggregate.Entities {
		outFile := fmt.Sprintf("%s/internals/%s/appquery/%s_appquery.go", b.outDir, b.aggregate.FileName(), item.FileName())
		build := NewBuildAppQuery(b.BaseBuild, item, utils.ToLower(outFile))
		b.AddBuild(build)
	}
}

func (b *BuildApplicationLayer) initQueryAppService() {
	outFile := fmt.Sprintf("%s/internals/%s/service/%s_query_appservice.go", b.outDir, b.aggregate.FileName(), b.aggregate.FileName())
	build := NewBuildAppService(b.BaseBuild, nil, utils.ToLower(outFile))
	b.AddBuild(build)
	for _, item := range b.aggregate.Entities {
		outFile := fmt.Sprintf("%s/internals/%s/service/%s_query_appservice.go", b.outDir, b.aggregate.FileName(), item.FileName())
		build := NewBuildAppService(b.BaseBuild, item, utils.ToLower(outFile))
		b.AddBuild(build)
	}
}

func (b *BuildApplicationLayer) initExecutorImplItem(entity *config.Entity) {
	entityFileName := b.aggregate.FileName()
	if entity != nil {
		entityFileName = entity.FileName()
	}

	tpls := []string{
		"find_all_executor.go.tpl",
		"find_by_id_executor.go.tpl",
		"find_by_ids_executor.go.tpl",
		"find_paging_executor.go.tpl",
		"create_executor.go.tpl",
		"create_many_executor.go.tpl",
		"delete_all_executor.go.tpl",
		"delete_by_id_executor.go.tpl",
		"delete_many_executor.go.tpl",
		"update_all_executor.go.tpl",
		"update_executor.go.tpl",
	}

	for _, tpl := range tpls {
		executorName := strings.Replace(tpl, ".go.tpl", "", 1)
		outFile := fmt.Sprintf("%s/internals/%s/executor/%s_impl/%s.go", b.outDir, b.aggregate.FileName(), entityFileName, executorName)
		build := NewBuildExecutorImpl(b.BaseBuild, entity, utils.ToLower(outFile), tpl)
		b.AddBuild(build)
	}

	if entity != nil {
		tpl := "find_by_aggregate_id_executor.go.tpl"
		executorName := fmt.Sprintf("find_by_%s_id_executor", b.aggregate.SnakeName())
		outFile := fmt.Sprintf("%s/internals/%s/executor/%s_impl/%s.go", b.outDir, b.aggregate.FileName(), entityFileName, executorName)
		build := NewBuildExecutorImpl(b.BaseBuild, entity, utils.ToLower(outFile), tpl)
		b.AddBuild(build)

		tpl = "delete_by_aggregate_id_executor.go.tpl"
		executorName = fmt.Sprintf("delete_by_%s_id_executor", b.aggregate.SnakeName())
		outFile = fmt.Sprintf("%s/internals/%s/executor/%s_impl/%s.go", b.outDir, b.aggregate.FileName(), entityFileName, executorName)
		build = NewBuildExecutorImpl(b.BaseBuild, entity, utils.ToLower(outFile), tpl)
		b.AddBuild(build)
	}

	tpl := "x_init.go.tpl"
	outFile := fmt.Sprintf("%s/internals/%s/executor/%s_impl/x_init.go", b.outDir, b.aggregate.FileName(), entityFileName)
	build := NewBuildExecutorImpl(b.BaseBuild, entity, utils.ToLower(outFile), tpl)
	b.AddBuild(build)

}
func (b *BuildApplicationLayer) initExecutorImpls() {
	b.initExecutorImplItem(nil)
	for _, entity := range b.aggregate.Entities {
		b.initExecutorImplItem(entity)
	}
}

func (b *BuildApplicationLayer) initExecutor() {
	outFile := fmt.Sprintf("%s/internals/%s/executor/%s_executor.go", b.outDir, b.aggregate.FileName(), b.aggregate.FileName())
	build := NewBuildExecutor(b.BaseBuild, nil, utils.ToLower(outFile))
	b.AddBuild(build)
	for _, item := range b.aggregate.Entities {
		outFile := fmt.Sprintf("%s/internals/%s/executor/%s_executor.go", b.outDir, b.aggregate.FileName(), item.FileName())
		build := NewBuildExecutor(b.BaseBuild, item, utils.ToLower(outFile))
		b.AddBuild(build)
	}
}

func (b *BuildApplicationLayer) initQueryHandlerEntities() {
	for _, item := range b.aggregate.Entities {
		outFile := fmt.Sprintf("%s/internals/%s/handler/%s_query_handler.go", b.outDir, b.Aggregate.FileName(), item.FileName())
		build := NewBuildQueryHandlerEntity(b.BaseBuild, b.aggregate, item, utils.ToLower(outFile))
		b.AddBuild(build)
	}
}

func (b *BuildApplicationLayer) initQueryHandlerAggregate() {
	outFile := fmt.Sprintf("%s/internals/%s/handler/%s_query_handler.go", b.outDir, b.Aggregate.FileName(), b.Aggregate.FileName())
	build := NewBuildQueryHandler(b.BaseBuild, b.aggregate, utils.ToLower(outFile))
	b.AddBuild(build)
}
