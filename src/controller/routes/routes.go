package routes

import (
	"github.com/mendesx3/crudzinho-golang/src/controller"

	"github.com/gin-gonic/gin"
)

func InitRoutes(r *gin.RouterGroup) {
	r.POST("/user", controller.PostUser)
	r.PUT("/user", controller.PutUserBy)
	r.GET("/user", controller.GetAllUser)
	r.GET("/user/:id", controller.GetUserBy)
	r.DELETE("/user/:id", controller.DeleteUserBy)
}
