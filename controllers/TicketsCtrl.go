package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type TicketsCtrl struct{}

func (ctrl *TicketsCtrl) Index(c *gin.Context) {
	// users := interactors.UserFindAll()
	c.JSON(http.StatusOK, nil)
}

func (ctrl *TicketsCtrl) Show(c *gin.Context) {
	c.JSON(http.StatusOK, nil)
}

func (ctrl *TicketsCtrl) Delete(c *gin.Context) {
	c.JSON(http.StatusOK, "successfully deleted")
}

func (ctrl *TicketsCtrl) Store(c *gin.Context) {
	c.JSON(http.StatusOK, nil)
}
