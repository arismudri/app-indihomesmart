package controllers

import (
	"app-indihomesmart/helpers"
	"app-indihomesmart/infra/logger"
	"app-indihomesmart/models"
	"app-indihomesmart/repository"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

func (base *Controller) UserGetList(ctx *gin.Context) {
	var user []*models.User
	offset, _ := strconv.Atoi(ctx.DefaultQuery("offset", "0"))
	limit, _ := strconv.Atoi(ctx.DefaultQuery("limit", "10"))

	repository.Get(&user, limit, offset)

	response := helpers.Response{
		Data: &user,
	}
	ctx.JSON(http.StatusOK, response.Success())
}

func (base *Controller) UserGetById(ctx *gin.Context) {
	var user models.User
	var response helpers.Response

	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		logger.Errorf("error : %v", err)
		response.Message = err.Error()
		ctx.JSON(http.StatusOK, response.Fail())
		return
	}

	user.Id = id
	err = repository.GetOne(&user)

	if err != nil {
		logger.Errorf("error : %v", err)
		response.Message = err.Error()
		ctx.JSON(http.StatusOK, response.Fail())
		return
	}

	user.Password = ""

	response.Data = &user
	ctx.JSON(http.StatusOK, response.Success())
}

func (base *Controller) UserAdd(ctx *gin.Context) {
	response := helpers.Response{}

	user := new(models.User)

	err := ctx.ShouldBindJSON(&user)
	if err != nil {
		logger.Errorf("error : %v", err)
		response.Message = err.Error()
		ctx.JSON(http.StatusOK, response.Fail())
		return
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), 8)
	if err != nil {
		logger.Errorf("error : %v", err)
		response.Message = err.Error()
		ctx.JSON(http.StatusOK, response.Fail())
		return
	}

	user.Password = string(hashedPassword)

	err = base.DB.Debug().Create(&user).Error
	if err != nil {
		logger.Errorf("error : %v", err)
		response.Message = err.Error()
		ctx.JSON(http.StatusOK, response.Fail())
		return
	}

	user.Password = ""

	response.Data = &user
	ctx.JSON(http.StatusOK, response.Success())
}

func (base *Controller) UserEdit(ctx *gin.Context) {
	response := helpers.Response{}

	user := new(models.User)
	dataUpdate := new(models.User)

	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		logger.Errorf("error : %v", err)
		response.Message = err.Error()
		ctx.JSON(http.StatusOK, response.Fail())
		return
	}

	user.Id = id

	err = ctx.ShouldBindJSON(&user)
	if err != nil {
		logger.Errorf("error : %v", err)
		response.Message = err.Error()
		ctx.JSON(http.StatusOK, response.Fail())
		return
	}

	dataUpdate.Name = user.Name
	dataUpdate.Bio = user.Bio

	repository.Update(&user, dataUpdate)

	user.Password = ""

	response.Data = &user
	ctx.JSON(http.StatusOK, response.Success())
}

func (base *Controller) UserDelete(ctx *gin.Context) {
	var user models.User
	var response helpers.Response

	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		logger.Errorf("error : %v", err)
		response.Message = err.Error()
		ctx.JSON(http.StatusOK, response.Fail())
		return
	}

	user.Id = id
	user.Password = ""
	repository.Delete(&user)

	ctx.JSON(http.StatusOK, response.Success())
}
