package main

import (
	"embed"
	"github.com/liuxd6825/dapr-ddd-cli/pkg/action"
	"github.com/liuxd6825/dapr-ddd-cli/pkg/resource"
	_ "github.com/liuxd6825/dapr-ddd-cli/pkg/resource"
	"github.com/urfave/cli/v2"
	"log"
	"os"
)

//go:embed "static/*"
var local embed.FS

func main() {
	resource.Init(local)
	app := &cli.App{
		Version: "v1.7.1-1.0-alpha2",
		Commands: []*cli.Command{
			{
				Name:    "init",
				Aliases: []string{"i"},
				Usage:   "初始化项目结构与代码",
				Action:  action.Action,
				Flags: []cli.Flag{
					&cli.StringFlag{
						Name:    "model",
						Aliases: []string{"m"},
						Value:   "",
						Usage:   "DDDML文件目录",
					},
					&cli.StringFlag{
						Name:    "lang",
						Aliases: []string{"l"},
						Value:   "go",
						Usage:   "开发语言,默认是go，可选：go,java",
					},
					&cli.StringFlag{
						Name:    "out",
						Aliases: []string{"o"},
						Value:   "",
						Usage:   "生成源代码目录",
					},
					&cli.StringFlag{
						Name:    "service",
						Aliases: []string{"s"},
						Value:   "",
						Usage:   "服务类型 cmd,query",
					},
					&cli.StringFlag{
						Name:    "layer",
						Aliases: []string{""},
						Value:   "",
						Usage:   "按层更新 ui,app,domain,infra,other 多个以逗号分隔",
					},
					&cli.StringFlag{
						Name:    "aggregate",
						Aliases: []string{"a"},
						Value:   "",
						Usage:   "只更新指定聚合根，多个以逗号分隔",
					},
				},
			},
		},
	}
	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
