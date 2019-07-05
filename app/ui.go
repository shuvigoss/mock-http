package app

import (
	"github.com/gin-gonic/gin"
	"mock-http/model"
	. "mock-http/util"
	"strconv"
)

type AddPathReq struct {
	Path string `json:"path"`
	Desc string `json:"desc"`
}

type UpdatePathReq struct {
	Path    string             `json:"path"`
	Methods []UpdatePathMethod `json:"methods"`
}

type UpdatePathMethod struct {
	Method string `json:"method"`
	Rsp    string `json:"response"`
}

func GetPath(c *gin.Context) {
	var paths []Path
	if err := DB.Preload("Methods").Find(&paths).Error; err != nil {
		c.JSON(200, model.FailWithData(model.StatusDatabaseError, err.Error()))
		return
	}
	c.JSON(200, model.Success(paths))
}

func AddPath(c *gin.Context) {
	var req AddPathReq
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(200, model.FailWithData(model.StatusParamError, err.Error()))
		return
	}

	p := Path{Path: req.Path, Desc: req.Desc}

	create := DB.Create(&p)
	if err := create.Error; err != nil {
		c.JSON(200, model.FailWithData(model.StatusDatabaseError, err.Error()))
		return
	}

	c.JSON(200, model.Success(p))
}

func UpdatePath(c *gin.Context) {
	id := c.Param("id")
	unitId, err := strconv.ParseUint(id, 10, 32)
	if err != nil {
		c.JSON(200, model.FailWithData(model.StatusParamError, err.Error()))
		return
	}

	var req UpdatePathReq
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(200, model.FailWithData(model.StatusParamError, err.Error()))
		return
	}

	tx := DB.Begin()
	if err := DB.Where("path_id = ?", id).Delete(Method{}).Error; err != nil {
		tx.Rollback()
		c.JSON(200, model.FailWithData(model.StatusDatabaseError, err.Error()))
		return
	}

	for _, me := range req.Methods {
		method := Method{Method: me.Method, Response: me.Rsp}
		method.PathId = uint(unitId)
		DB.Create(&method)
	}

	tx.Commit()
	c.JSON(200, model.Success(nil))
}

func DeletePath(c *gin.Context) {
	id := c.Param("id")
	unitId, err := strconv.ParseUint(id, 10, 32)
	if err != nil {
		c.JSON(200, model.FailWithData(model.StatusParamError, err.Error()))
		return
	}
	path := Path{ID: uint(unitId)}
	if err := DB.Delete(&path).Error; err != nil {
		c.JSON(200, model.FailWithData(model.StatusDatabaseError, err.Error()))
		return
	}

	c.JSON(200, model.Success(nil))

}
