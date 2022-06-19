package commands

import (
	"github.com/urfave/cli/v2"
	// "fmt"
)

// func BuildAction(c *cli.Context) err{
// 	fmt.Println("welcome to Goup")

// 	return nil
// }

func GetApp(name string) *cli.App {
	app := &cli.App{
		Name: name,
		Flags: []cli.Flag{
			// 往哪个网站传
			&cli.StringFlag{
				Name:        "site",
				Required:    false,
				Value:       "cnblog",
				DefaultText: "default cnblog",
				Usage:       "which blog to upload",
				// Destination: &site,
			},
			&cli.StringFlag{
				Name: "file",
				Required: true,
				Aliases:  []string{"f"},
				Usage: "file path what to be upload",
			},
			// TODO:文章类别
			&cli.StringFlag{
				Name:        "category",
				Required:    false,
				Hidden: true,
				Usage:       "Which category do blogs fall into",
				// Destination: &category,
			},
			// 文章标签
			&cli.StringSliceFlag{
				Name:        "tag",
				Usage:       "Which tag do blog fall into",
				Required:    false,
				// Destination: &tags,
			},
			// TODO:是否是传目录
			&cli.BoolFlag{
				Name:        "dir",
				Required:    false,
				Usage:       "whether upload dir or not",
				// Destination: &isdir,
				/*
				Value: 代表默认值
				Destination: 绑定某个变量
				*/
			},
		},
		// Action: BuildAction,
	}
	return app
}
