package plugins

import (
	"encoding/json"

	"../db"
	"../plugin"
	"github.com/etcd-io/bbolt"

	"github.com/gin-gonic/gin"
)

type (
	pages   struct{}
	pageDao struct{}

	listView struct {
		Skip int
		Take int
	}

	page struct {
		ID    string   `json:"id"`
		Title string   `json:"title"`
		Tags  []string `json:"tags"`
	}

	pageDetail struct {
		ID      string   `json:"id"`
		Title   string   `json:"title"`
		Type    string   `json:"type"`
		Content string   `json:"content"`
		Tags    []string `json:"tags"`
	}
)

var dao = &pageDao{}

const bucket = "pages"

func Pages() plugin.Plugin {
	return &pages{}
}

func (p *pages) Prefix() string {
	return "/p/"
}

func (p *pages) Name() string {
	return "Pages"
}

func (p *pages) Bind(router *gin.RouterGroup) {
	router.GET("/", func(ctx *gin.Context) {
		ctx.HTML(200, "pages.html", gin.H{})
	})
	router.GET("/api/pages", func(ctx *gin.Context) {
		list := &listView{}

		err := ctx.BindQuery(list)

		if err != nil {
			ctx.AbortWithError(500, err)
			return
		}

		if list.Take == 0 {
			list.Take = 15
		}

		data := dao.List(list.Skip, list.Take)

		ctx.JSON(200, data)
	})
	router.POST("/api/page", func(ctx *gin.Context) {
		details := &pageDetail{}
		err := ctx.BindJSON(details)

		if err != nil {
			ctx.AbortWithError(400, err)
			return
		}

		err = dao.Store(details)

		if err != nil {
			ctx.AbortWithError(500, err)
			return
		}

		ctx.JSON(200, nil)
	})
	router.GET("/api/page/:id", func(ctx *gin.Context) {
		id := ctx.Param("id")

		detail, err := dao.Details(id)

		if err != nil {
			ctx.AbortWithError(500, err)
			return
		}

		ctx.JSON(200, detail)
	})
}

func (p *pageDao) List(skip, take int) []*page {
	pages := make([]*page, 0)

	db.ReadBucket(bucket, func(b *bbolt.Bucket) error {
		c := b.Cursor()

		_, v := c.First()
		first := &page{}
		json.Unmarshal(v, first)

		if skip == 0 {
			pages = append(pages, first)
			take = take - 1
		} else {
			skip = skip - 1

			for skip > 0 {
				c.Next()
			}
		}

		for take > 0 {
			_, v = c.Next()
			json.Unmarshal(v, first)
			pages = append(pages, first)
		}

		return nil
	})

	return pages
}

func (p *pageDao) Store(detail *pageDetail) error {
	return db.WriteBucket(bucket, func(b *bbolt.Bucket) error {
		data, err := json.Marshal(detail)

		if err != nil {
			return err
		}

		return b.Put([]byte(detail.ID), data)
	})
}

func (p *pageDao) Remove(id string) error {
	return db.WriteBucket(bucket, func(b *bbolt.Bucket) error {
		return b.Delete([]byte(id))
	})
}

func (p *pageDao) Details(id string) (*pageDetail, error) {
	detail := &pageDetail{}

	err := db.ReadBucket(bucket, func(b *bbolt.Bucket) error {
		body := b.Get([]byte(id))

		return json.Unmarshal(body, detail)
	})

	return detail, err
}
