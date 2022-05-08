package main

import (
	"embed"
	"fmt"
	"github.com/dapr/dapr-ddd-cli/pkg/cmd-init"
	"github.com/dapr/dapr-ddd-cli/pkg/resource"
	_ "github.com/dapr/dapr-ddd-cli/pkg/resource"
	"github.com/urfave/cli/v2"
	"log"
	"os"
)

//go:embed "static/*"
var local embed.FS

func main() {
	resource.Init(local)

	app := &cli.App{
		Commands: []*cli.Command{
			{
				Name:    "init",
				Aliases: []string{""},
				Usage:   "初始代项目",
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
			{
				Name:    "complete",
				Aliases: []string{"c"},
				Usage:   "complete a task on the list",
				Action: func(c *cli.Context) error {
					fmt.Println("completed task: ", c.Args().First())
					return nil
				},
			},
			{
				Name:    "template",
				Aliases: []string{"t"},
				Usage:   "options for task templates",
				Subcommands: []*cli.Command{
					{
						Name:  "add",
						Usage: "add a new template",
						Action: func(c *cli.Context) error {
							fmt.Println("new task template: ", c.Args().First())
							return nil
						},
					},
					{
						Name:  "remove",
						Usage: "remove an existing template",
						Action: func(c *cli.Context) error {
							fmt.Println("removed task template: ", c.Args().First())
							return nil
						},
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
