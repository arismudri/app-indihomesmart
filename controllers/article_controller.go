package controllers

import (
	"app-indihomesmart/helpers"
	"app-indihomesmart/infra/logger"
	"app-indihomesmart/models"
	"app-indihomesmart/repository"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func (base *Controller) ArticleGetList(ctx *gin.Context) {
	var article []*models.Article
	offset, _ := strconv.Atoi(ctx.DefaultQuery("offset", "0"))
	limit, _ := strconv.Atoi(ctx.DefaultQuery("limit", "10"))

	repository.Get(&article, limit, offset)

	response := helpers.Response{
		Data: &article,
	}
	ctx.JSON(http.StatusOK, response.Success())
}

func (base *Controller) ArticleGetById(ctx *gin.Context) {
	var article models.Article
	var response helpers.Response

	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		logger.Errorf("error : %v", err)
		response.Message = err.Error()
		ctx.JSON(http.StatusOK, response.Fail())
		return
	}

	article.Id = id
	err = repository.GetOne(&article)

	if err != nil {
		logger.Errorf("error : %v", err)
		response.Message = err.Error()
		ctx.JSON(http.StatusOK, response.Fail())
		return
	}

	response.Data = &article

	ctx.JSON(http.StatusOK, response.Success())
}

func (base *Controller) ArticleAdd(ctx *gin.Context) {
	response := helpers.Response{}

	article := new(models.Article)

	err := ctx.ShouldBindJSON(&article)
	if err != nil {
		logger.Errorf("error : %v", err)
		response.Message = err.Error()
		ctx.JSON(http.StatusOK, response.Fail())
		return
	}
	err = base.DB.Debug().Create(&article).Error
	if err != nil {
		logger.Errorf("error : %v", err)
		response.Message = err.Error()
		ctx.JSON(http.StatusOK, response.Fail())
		return
	}

	response.Data = &article
	ctx.JSON(http.StatusOK, response.Success())
}

func (base *Controller) ArticleEdit(ctx *gin.Context) {
	response := helpers.Response{}

	article := new(models.Article)
	dataUpdate := new(models.Article)

	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		logger.Errorf("error : %v", err)
		response.Message = err.Error()
		ctx.JSON(http.StatusOK, response.Fail())
		return
	}

	article.Id = id

	err = ctx.ShouldBindJSON(&article)
	if err != nil {
		logger.Errorf("error : %v", err)
		response.Message = err.Error()
		ctx.JSON(http.StatusOK, response.Fail())
		return
	}

	dataUpdate.Title = article.Title
	dataUpdate.Content = article.Content
	dataUpdate.Category = article.Category
	dataUpdate.Status = article.Status

	repository.Update(&article, dataUpdate)

	response.Data = &article
	ctx.JSON(http.StatusOK, response.Success())
}

func (base *Controller) ArticleDelete(ctx *gin.Context) {
	var article models.Article
	var response helpers.Response

	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		logger.Errorf("error : %v", err)
		response.Message = err.Error()
		ctx.JSON(http.StatusOK, response.Fail())
		return
	}

	article.Id = id
	repository.Delete(&article)

	ctx.JSON(http.StatusOK, response.Success())
}
