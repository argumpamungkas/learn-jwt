package router

import (
	"DTS/Chapter-3/sesi/sesi2-go-jwt/controllers"

	"github.com/gin-gonic/gin"
)

// Method grup merupakan method dari gin framework, dan mempermudah dalam melakukan route grouping. jadi nantinya ketika client
// mengirimkan request dengan path /users, maka request tsb akan masuk kedalam scope userRouter/scope setelah deklarasi userRouter
// scope tersebut dikhususkan untuk routing endpoint users.
func StartApp() *gin.Engine {
	r := gin.Default()

	userRouter := r.Group("/users")
	{
		userRouter.POST("/register", controllers.UserRegister)

		userRouter.POST("/login", controllers.UserLogin)
	}

	return r
}
