package commands

import (
	"fmt"

	"github.com/Meduzz/devserver/config"
	"github.com/gin-gonic/gin"
	"github.com/spf13/cobra"
)

func init() {
	cms := &cobra.Command{}
	cms.Use = "cms"
	cms.Example = "cms --file=app.json [--port=8765]"
	cms.Short = "start handy in cms mode"
	cms.Flags().String("file", "app.json", "define which app (config) to load and modify in the cms mode")
	cms.Flags().Int("port", 8765, "set port that the cms mode will start up in")
	cms.MarkFlagFilename("file", "json")
	cms.RunE = cmsHandler

	Root.AddCommand(cms)
}

func cmsHandler(cmd *cobra.Command, args []string) error {
	file, err := cmd.Flags().GetString("file")

	if err != nil {
		return err
	}

	if file == "" {
		return fmt.Errorf("file flag must not be empty")
	}

	/*
		TODO
		Build an app context for the app we're editing.
		Initiate the CMS with the app context
		In the CMS define:
		- a database for the app context config file?
		- a bootstrap endpoint for the SPA
		- a read app context config file endpoint
		- a write app context config file endpoint
	*/

	port, err := cmd.Flags().GetInt("port")

	if err != nil {
		return err
	}

	srv := gin.Default()

	return srv.Run(fmt.Sprintf(":%d", port))
}

var (
	cfg = config.Config{
		App: &config.App{
			Name:   "CMS",
			Prefix: "/cms",
		},
		Endpoints: append([]*config.Endpoint{},
			&config.Endpoint{
				Kind:   "static",
				Method: "GET",
				Path:   "/",
				Static: &config.Static{
					Datasource: "dist/index.html",
				},
			},
			&config.Endpoint{
				Kind:   "static",
				Method: "GET",
				Path:   "/dbs",
				Static: &config.Static{
					Datasource: "dist/index.html",
				},
			},
			&config.Endpoint{
				Kind:   "static",
				Method: "GET",
				Path:   "/pages",
				Static: &config.Static{
					Datasource: "dist/index.html",
				},
			},
			&config.Endpoint{
				Kind:   "static",
				Method: "GET",
				Path:   "/assets/index.js",
				Static: &config.Static{
					Datasource: "dist/assets/index.js",
				},
			},
			&config.Endpoint{
				Kind:   "static",
				Method: "GET",
				Path:   "/assets/index.css",
				Static: &config.Static{
					Datasource: "dist/assets/index.css",
				},
			},
			&config.Endpoint{
				Kind:   "api",
				Method: "GET",
				Path:   "/bootstrap",
				Api: &config.Api{
					Datasource: "bootstrap",
				},
			},
			&config.Endpoint{
				Kind:   "api",
				Method: "GET",
				Path:   "/config",
				Api: &config.Api{
					Datasource: "config",
				},
			},
			&config.Endpoint{
				Kind:   "api",
				Method: "PUT",
				Path:   "/config",
				Api: &config.Api{
					Datasource: "config",
				},
			},
		),
		Datasources: append([]*config.Datasource{},
			&config.Datasource{
				Kind:     "property",
				Encoding: "json",
				Name:     "bootstrap",
				Property: &config.Property{
					Field:  "bootstrap",
					Base64: false,
				},
			},
			&config.Datasource{
				Kind:     "file",
				Encoding: "json",
				Name:     "config",
				File: &config.File{
					URL: ".", // TODO is this smart? Something magic is needed anyway.
				},
			},
			&config.Datasource{
				Kind:     "file",
				Encoding: "html",
				Name:     "dist/index.html",
				File: &config.File{
					URL: "dist/index.html",
				},
			},
			&config.Datasource{
				Kind:     "file",
				Encoding: "js",
				Name:     "dist/assets/index.js",
				File: &config.File{
					URL: "dist/assets/index.js",
				},
			},
			&config.Datasource{
				Kind:     "file",
				Encoding: "css",
				Name:     "dist/assets/index.css",
				File: &config.File{
					URL: "dist/assets/index.css",
				},
			},
		), // TODO now the fun part... the extra properties not part of the config...
	}
)
