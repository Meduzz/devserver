package server

import (
	"github.com/gin-gonic/gin"

	"../plugin"
	"github.com/urfave/cli"
)

func NewServer(plugins ...plugin.Plugin) cli.Command {
	cmd := cli.Command{}
	cmd.Name = "server"
	cmd.Usage = "server [-bind :1090]"
	cmd.Flags = []cli.Flag{
		cli.StringFlag{
			Name:  "bind",
			Value: ":1090",
			Usage: "-bind :1090",
		},
	}
	cmd.Action = func(ctx *cli.Context) {
		bind := ctx.String("bind")

		srv := gin.Default()
		srv.LoadHTMLGlob("templates/*")

		srv.GET("/", func(ctx *gin.Context) {
			ctx.HTML(200, "index.html", gin.H{})
		})
		srv.GET("/menu", func(ctx *gin.Context) {
			menu := make(map[string]string)

			for _, p := range plugins {
				menu[p.Prefix()] = p.Name()
			}

			ctx.JSON(200, menu)
		})

		srv.Static("/static", "./static")

		for _, p := range plugins {
			grp := srv.Group(p.Prefix())
			p.Bind(grp)
		}

		srv.Run(bind)
	}

	return cmd
}
