package server

import (
	"cadence-poc/api/controllers"
	"github.com/gin-gonic/gin"
)

func InitRouter(rg *gin.RouterGroup) {
	initUsersRoutes(rg)
	initValidatorsRoutes(rg)

}

func initUsersRoutes(rg *gin.RouterGroup) {
	g := rg.Group("/user")
	{
		g.POST("/register", controllers.UserRegister)
		g.POST("/login", controllers.UserLogin)
	}
}

func initValidatorsRoutes(rg *gin.RouterGroup) {
	rg.POST("/deploy/:proto", controllers.ValidatorsDeploy)
	rg.POST("/unstake/:proto", controllers.ValidatorsUnstake)
}
