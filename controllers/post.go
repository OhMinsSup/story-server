package controllers

import (
	"github.com/OhMinsSup/story-server/dto"
	"github.com/OhMinsSup/story-server/helpers"
	"github.com/OhMinsSup/story-server/services"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"net/http"
)

func UpdatePostController(ctx *gin.Context) {
	var body dto.WritePostBody
	if err := ctx.BindJSON(&body); err != nil {
		ctx.AbortWithStatus(http.StatusBadRequest)
		return
	}

	db := ctx.MustGet("db").(*gorm.DB)
	result, code, err := services.UpdatePostService(body, db, ctx)
	if err != nil {
		ctx.AbortWithError(code, err)
		return
	}

	ctx.JSON(code, result)
}

func DeletePostController(ctx *gin.Context) {
	db := ctx.MustGet("db").(*gorm.DB)
	result, code, err := services.DeletePostService(db, ctx)
	if err != nil {
		ctx.AbortWithError(code, err)
		return
	}

	ctx.JSON(code, result)
}

func WritePostController(ctx *gin.Context) {
	var body dto.WritePostBody
	if err := ctx.BindJSON(&body); err != nil {
		ctx.AbortWithStatus(http.StatusBadRequest)
		return
	}

	db := ctx.MustGet("db").(*gorm.DB)
	result, code, err := services.WritePostService(body, db, ctx)
	if err != nil {
		ctx.AbortWithError(code, err)
		return
	}

	ctx.JSON(code, result)
}

func GetPostController(ctx *gin.Context) {
	db := ctx.MustGet("db").(*gorm.DB)
	result, code, err := services.GetPostService(db, ctx)
	if err != nil {
		ctx.AbortWithError(code, err)
		return
	}

	ctx.JSON(code, result)
}

func PostViewController(ctx *gin.Context) {
	ip := helpers.CreateHash(ctx.ClientIP())
	postId := ctx.Param("post_id")

	params := dto.PostViewParams{
		Ip:     ip,
		PostId: postId,
	}

	db := ctx.MustGet("db").(*gorm.DB)
	result, code, err := services.PostViewService(params, db, ctx)
	if err != nil {
		ctx.AbortWithError(code, err)
		return
	}

	ctx.JSON(code, result)
}