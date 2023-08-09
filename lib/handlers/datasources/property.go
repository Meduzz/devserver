package datasources

import (
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/Meduzz/devserver/config"
	"github.com/Meduzz/devserver/lib/common"
	"github.com/Meduzz/devserver/lib/handlers/model"
	"github.com/Meduzz/helper/fp/slice"
	"github.com/gin-gonic/gin"
	"github.com/tidwall/gjson"
	"github.com/tidwall/sjson"
)

type (
	objectPropertyDs struct {
		ctx *common.AppContext
		cfg *config.Property
	}

	listPropertyDs struct {
		ctx *common.AppContext
		cfg *config.Property
	}

	keydPropertyDs struct {
		ctx *common.AppContext
		cfg *config.Property
	}
)

func NewProperty(ctx *common.AppContext, cfg *config.Property) model.Datasource {
	switch cfg.Kind {
	case "list":
		return &listPropertyDs{ctx, cfg}
	case "keyd":
		return &keydPropertyDs{ctx, cfg}
	default:
		return &objectPropertyDs{ctx, cfg}
	}
}

func (p *objectPropertyDs) Handler() (gin.HandlerFunc, error) {
	item, err := loadProperty(p.ctx.File, p.cfg.Field)

	if err != nil {
		return nil, err
	}

	return func(ctx *gin.Context) {
		switch ctx.Request.Method {
		case http.MethodGet:
			ctx.JSON(200, item.Value())
		case http.MethodPost:
		case http.MethodPut:
			data, err := ctx.GetRawData()

			if err != nil {
				ctx.AbortWithError(500, err)
				return
			}

			err = storeProperty(p.ctx.File, p.cfg.Field, data)

			if err != nil {
				ctx.AbortWithError(500, err)
				return
			}

			ctx.Status(200)
		case http.MethodDelete:
			err := deleteProperty(p.ctx.File, p.cfg.Field)

			if err != nil {
				ctx.AbortWithError(500, err)
				return
			}

			ctx.Status(200)
		}
	}, nil
}

func (p *objectPropertyDs) Context() *common.AppContext {
	return p.ctx
}

func (p *listPropertyDs) Handler() (gin.HandlerFunc, error) {
	item, err := loadProperty(p.ctx.File, p.cfg.Field)

	if err != nil {
		return nil, err
	}

	return func(ctx *gin.Context) {
		switch ctx.Request.Method {
		case http.MethodGet:
			skip := ctx.DefaultQuery("skip", "0")
			take := ctx.DefaultQuery("take", "15")

			iSkip, err := strconv.Atoi(skip)

			if err != nil {
				ctx.AbortWithError(400, err)
				return
			}

			iTake, err := strconv.Atoi(take)

			if err != nil {
				ctx.AbortWithError(400, err)
				return
			}

			items := slice.Skip(item.Array(), iSkip)
			items = slice.Take(items, iTake)
			jsonItems := slice.Map(items, func(item gjson.Result) interface{} {
				return item.Value()
			})

			ctx.JSON(200, jsonItems)
		case http.MethodPost:
			bs, err := ctx.GetRawData()

			if err != nil {
				ctx.AbortWithError(400, err)
				return
			}

			err = appendProperty(p.ctx.File, p.cfg.Field, bs)

			if err != nil {
				ctx.AbortWithError(500, err)
				return
			}

			ctx.Data(200, "application/json", bs)
		case http.MethodPut:
			bs, err := ctx.GetRawData()

			if err != nil {
				ctx.AbortWithError(400, err)
				return
			}

			idString, ok := ctx.GetQuery(p.cfg.IdField)

			if !ok {
				log.Println("IdField was not set")
				ctx.Status(400)
				return
			}

		case http.MethodDelete:
		}
	}, nil
}

func (p *listPropertyDs) Context() *common.AppContext {
	return p.ctx
}

func (p *keydPropertyDs) Handler() (gin.HandlerFunc, error) {
	item, err := loadProperty(p.ctx.File, p.cfg.Field)

	if err != nil {
		return nil, err
	}

	return func(ctx *gin.Context) {
		switch ctx.Request.Method {
		case http.MethodGet:
			ctx.JSON(200, item.Value())
		case http.MethodPost:
		case http.MethodPut:
		case http.MethodDelete:
		}
	}, nil
}

func (p *keydPropertyDs) Context() *common.AppContext {
	return p.ctx
}

func loadProperty(file, field string) (gjson.Result, error) {
	bs, err := common.LoadProperty(file, field)

	if err != nil {
		return gjson.Result{}, err
	}

	return gjson.ParseBytes(bs), nil
}

func storeProperty(file, field string, data []byte) error {
	bs, err := common.LoadFileContent(file)

	if err != nil {
		return err
	}

	raw, err := sjson.SetBytes(bs, field, data)

	if err != nil {
		return err
	}

	return common.StoreFileContent(file, raw)
}

func deleteProperty(file, field string) error {
	bs, err := common.LoadFileContent(file)

	if err != nil {
		return err
	}

	raw, err := sjson.DeleteBytes(bs, field)

	if err != nil {
		return err
	}

	return common.StoreFileContent(file, raw)
}

func appendProperty(file, field string, data []byte) error {
	bs, err := common.LoadFileContent(file)

	if err != nil {
		return err
	}

	fieldArray := fmt.Sprintf("%s.-1", field)

	raw, err := sjson.SetBytes(bs, fieldArray, data)

	if err != nil {
		return err
	}

	return common.StoreFileContent(file, raw)
}
