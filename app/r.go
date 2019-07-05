package app

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func (a App) Route() {
	r := a.R

	//sfs, _ := fs.New()

	r.GET("/static/*filename", gin.WrapH(a.UI.StaticHandler()))
	r.GET("/", func(context *gin.Context) {
		context.Redirect(http.StatusMovedPermanently, "/static")
	})

	api := r.Group("/api")
	api.POST("/path", AddPath)
	api.GET("/path", GetPath)
	api.PUT("/path/:id", UpdatePath)
	api.DELETE("/path/:id", DeletePath)

	mr := a.MR
	mr.Any("/*router", Mock)
}
