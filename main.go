package main

import (
	"Goup/commands"
	"Goup/constants"
	"Goup/service"
	"Goup/utils"
	"fmt"
	"log"
	"os"
	"regexp"
	// "time"
	"github.com/urfave/cli/v2"
	// "time"
	// "github.com/fatih/color"
)

const VERSION = "v1.0"
const TITLE = "goup"
const DESC = "a tool for upload blog"



func main() {
		var cnb service.CnBlog
		cnb.Init(constants.USERNAME, constants.PASSWD, constants.BLOGADDR, constants.URL)
		app := commands.GetApp(TITLE)

		app.Action = func(ctx *cli.Context) error {

			if isDir := ctx.Bool("dir"); isDir {
				fmt.Println("sorry, we not'support apply dir currently, coming soon!")
				return nil
			}
			md := ctx.String("file")

			if md == "" {
				fmt.Println("file can't be null")
				return nil
			}
			if isFile, err := utils.CheckFile(md); err == nil {
				if !isFile {
					fmt.Println("the param file should be the path onf a file")
					return nil
				}
				tmpMd := []string{md}
				nPost := utils.ParseMd(&cnb, tmpMd...)
				cnb.UpBlog(nPost)
			} else {
				fmt.Println("the param file is invalid, please check it")
				return nil
			}

			return nil
		}

		err := app.Run(os.Args)
		if err != nil {
			log.Fatal(err)
		}

}
