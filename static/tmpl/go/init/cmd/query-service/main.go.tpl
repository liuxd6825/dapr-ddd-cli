package main

import (
	"flag"
	"{{.Namespace}}/pkg/cmd-service/domain/event"
	"{{.Namespace}}/pkg/query-service/domain/handler"
	"{{.Namespace}}/pkg/query-service/userinterface/rest/controller"
	"github.com/liuxd6825/dapr-go-ddd-sdk/restapp"
)

func main() {
	help := flag.Bool("help", false, "参数提示。")
	envType := flag.String("envType", "", "替换配置文件中的envType值。")
	config := flag.String("config", "./config/query-config.yaml", "配置文件。")
	flag.Parse()

	if *help {
		return
	}

	if _, err := restapp.RunWithConfig(*envType, *config, subscribes, controllers, events, restapp.Actors); err != nil {
		panic(err)
	}
}

// 注册消息监听器
func subscribes() *[]restapp.RegisterSubscribe {
	return handler.GetRegisterSubscribe()
}

// 注册Http控制器
func controllers() *[]restapp.Controller {
	return controller.GetRegisterController()
}

// 注册领域事件
func events() *[]restapp.RegisterEventType {
	return event.GetRegisterEventType()
}
