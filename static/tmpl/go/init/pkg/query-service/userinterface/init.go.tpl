package userinterface

import (
	"github.com/kataras/iris"
	"github.com/kataras/iris/v12/mvc"
	"github.com/liuxd6825/dapr-go-ddd-example/pkg/query-service/userinterface/rest/controller"
)

func RegisterMvcController(app *iris.Application) {
	mvc.Configure(app.Party("/api/v1.0"), registerMvcController)
}

func registerMvcController(app *mvc.Application) {
	app.Handle(controller.NewUserController())
}
