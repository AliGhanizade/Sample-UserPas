package controller

import (
	"net/http"
	model "userpas/model"

	"github.com/gin-gonic/gin"
	"github.com/spf13/cast"
)

func (c *UserController) Show(ctx *gin.Context) {
	var User model.User

	err := User.Get()
	if err != nil {
		ctx.JSON(http.StatusBadRequest, err)

	}
	ctx.JSON(http.StatusOK, User)
}
func (c *UserController) Admin(ctx *gin.Context) {
	var Users model.Users
	var User model.User
	err := ctx.ShouldBindJSON(&User)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, err)
	}
	err = Users.GetAdmin(cast.ToInt16(User.Rulse))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, err)
	}
	ctx.JSON(http.StatusOK, Users)
}
