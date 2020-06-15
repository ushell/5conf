package router

import (
	"5conf/api"
	"5conf/library/e"
	"5conf/library/http"
	"github.com/gin-gonic/gin"
)

func InitRouter(router *gin.Engine) {
	router.GET("/", func (ctx *gin.Context) {
		ctx.JSON(200, http.Response{Code: e.SUCCESS, Message: "ok"})
	})

	// 中间件
	//router.Use(middleware.Logger())

	v1 := router.Group("api/v1")
	{
		config := new(api.Config)
		v1.POST("/config", config.CreateConfig)
		v1.GET("/config/:id", config.GetConfig)
		v1.PUT("/config/:id", config.UpdateConfig)
		v1.DELETE("/config/:id", config.DeleteConfig)

		project := new(api.Project)
		v1.POST("/project", project.CreateProject)
		v1.GET("/project/:id", project.GetProject)
		v1.PUT("/project/:id", project.UpdateProject)
		v1.DELETE("/project/:id", project.DeleteProject)

		env := new(api.Environment)
		v1.POST("/environment", env.CreateEnv)
		v1.GET("/environment/:id", env.GetEnv)
		v1.PUT("/environment/:id", env.UpdateEnv)
		v1.DELETE("/environment/:id", env.DeleteEnv)

	}
}
