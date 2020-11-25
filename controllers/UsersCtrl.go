package controllers

import (
	"net/http"
	"strconv"

	"hexagonal/entities"
	"hexagonal/interactors"

	"github.com/gin-gonic/gin"
)

var (
	userItr = interactors.UserItr{}
)

type UsersCtrl struct{}

func (ctrl *UsersCtrl) Index(c *gin.Context) {
	users, _ := userItr.FindAll()
	c.JSON(http.StatusOK, users)
}

func (ctrl *UsersCtrl) Show(c *gin.Context) {
	idStr := c.Params.ByName("user")
	id, _ := strconv.ParseUint(idStr, 10, 64)

	user, _ := userItr.GetById(id)
	c.JSON(http.StatusOK, user)
}

func (ctrl *UsersCtrl) Delete(c *gin.Context) {
	c.JSON(http.StatusOK, "successfully deleted")
}

func (ctrl *UsersCtrl) Store(c *gin.Context) {
	var user entities.User
	c.BindJSON(&user)

	userItr.Create(&user)
	c.JSON(http.StatusOK, nil)
}
