package router

import (
	"github.com/gin-gonic/gin"
	"github.com/Hariharan148/Go-JWT-Auth-Gin/controllers"
)

func AuthRouter(incomming *gin.Engine){
	incomming.POST("user/signup", controllers.SignUp())
	incomming.POST("user/login", controllers.Login())
}