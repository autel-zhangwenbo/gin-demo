package routers

import (
	"fmt"
	"gin-demo/internal/configs"
	"gin-demo/internal/controllers"
	doc "gin-demo/swagger"
	"github.com/gin-gonic/gin"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func SetupRouter(r *gin.Engine) {
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	v1 := r.Group("/api/v1")

	articles := v1.Group("article")
	{
		articles.GET("", controllers.ListArticles)
		articles.GET ("/:id", controllers.GetArticle)
		articles.POST("/:id", controllers.CreateArticle)
		articles.PUT("/:id", controllers.UpdateArticle)
		articles.DELETE("/:id", controllers.DeleteArticle)
	}

	doc.SwaggerInfo.Host = fmt.Sprintf("%s:%s", configs.Cfg.AppCfg.HttpHost, configs.Cfg.AppCfg.HttpPort)
	doc.SwaggerInfo.BasePath = "/api/v1"
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))
	httpPort := ":" + configs.Cfg.AppCfg.HttpPort
	r.Run(httpPort)

}
