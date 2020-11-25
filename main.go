package main

import (
	"hexagonal/controllers"
	dataSource "hexagonal/dataSource/mysql"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/prometheus/common/log"
)

var Router = gin.Default()

func init() {

	if err := godotenv.Load(); err != nil {
		log.Error(err)
	}

	if err := dataSource.InitMysql(); err != nil {
		log.Error(err)
		return
	}

	users := Router.Group("/users")
	{
		usersCtl := new(controllers.UsersCtrl)

		users.GET("/", usersCtl.Index)
		users.GET("/:user", usersCtl.Show)
		users.DELETE("/:user", usersCtl.Delete)
		users.POST("/", usersCtl.Store)
	}

	Router.NoRoute(func(c *gin.Context) {
		c.JSON(404, gin.H{"message": "Not found"})
	})

}

func main() {
	Router.Run(":5000")
}
