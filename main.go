package main

import (
	"embed"
	"github.com/liuxd6825/dapr-ddd-cli/pkg/cmd-init"
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
		Version: "v1.7.1-1.0-alpha",
		Commands: []*cli.Command{
			{
				Name:    "init",
				Aliases: []string{""},
				Usage:   "初始化项目结构与代码",
				Action:  cmd_init.Acton,
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
						Value:   "./",
						Usage:   "生成源代码目录",
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
