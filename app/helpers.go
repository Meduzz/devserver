package app

import (
	"encoding/json"
	"os"
	"strings"

	"github.com/Meduzz/devserver/model"
	"github.com/Meduzz/helper/http/client"
)

func Load(uri string) (*model.App, error) {
	if strings.HasPrefix(uri, "http") {
		req, err := client.GET(uri)

		if err != nil {
			return nil, err
		}

		res, err := req.DoDefault()

		if err != nil {
			return nil, err
		}

		bs, err := res.AsBytes()

		app := &model.App{
			Name: "TODO",
		}
		err = json.Unmarshal(bs, app)

		if err != nil {
			return nil, err
		}

		return app, nil
	} else {
		return LoadFile(uri)
	}
}

func LoadFile(uri string) (*model.App, error) {
	app := &model.App{
		Name: "TODO",
	}

	bs, err := os.ReadFile(uri)

	if err != nil {
		if os.IsNotExist(err) {
			return app, nil
		}

		return nil, err
	}

	err = json.Unmarshal(bs, app)

	if err != nil {
		return nil, err
	}

	return app, nil
}

func Store(app *model.App, uri string) error {
	bs, err := json.Marshal(app)

	if err != nil {
		return err
	}

	return os.WriteFile(uri, bs, 0644)
}
