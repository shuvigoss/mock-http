package app

import (
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/render"
	"mock-http/model"
	. "mock-http/util"
	"strings"
)

func Mock(c *gin.Context) {
	uri := c.Request.RequestURI
	method := c.Request.Method
	var p Path

	if err := DB.Where("path = ?", uri).Preload("Methods").Find(&p).Error; err != nil {
		c.JSON(404, model.FailWithData(model.ServerError, err.Error()))
		return
	}

	for _, m := range p.Methods {
		if m.Method == method {
			if strings.HasPrefix(m.Response, "{") || strings.HasPrefix(m.Response, "[") {
				c.Writer.Header().Set("Content-Type", "application/json; charset=utf-8")
			} else if strings.HasPrefix(m.Response, "<") {
				c.Writer.Header().Set("Content-Type", "application/xml; charset=utf-8")
			} else {
				c.Writer.Header().Set("Content-Type", "text/plain; charset=utf-8")
			}
			c.Render(200, render.Data{Data: []byte(m.Response)})
			return
		}
	}

	c.JSON(404, model.FailWithData(model.ServerError, "未找到路由"))
}
