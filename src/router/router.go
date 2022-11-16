package router

import (
	"github.com/gin-gonic/gin"
	"github.com/mendesx3/crudzinho-golang/src/controller/user"
)

func InitRoutes(r *gin.RouterGroup) {
	u := r.Group("/user")
	u.POST("/", user.PostUser)
	u.PUT("/:id", user.PutUserBy)
	u.GET("/", user.GetAllUser)
	u.GET("/:id", user.GetUserBy)
	u.DELETE("/:id", user.DeleteUserBy)
}
