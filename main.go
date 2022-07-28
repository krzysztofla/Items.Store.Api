package main

import (
	docs "github.com/Items.Store.Api/docs"
	"github.com/Items.Store.Api/handlers"
	"github.com/gin-gonic/gin"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// @title           Items Store Go Example API
// @version         1.0
// @description     This is a sample server celler server.
// @termsOfService  TBD

// @contact.name   API Support
// @contact.url    TBD
// @contact.email  krzysztof.lach@icloud.com

// @license.name  Apache 2.0
// @license.url   http://www.apache.org/licenses/LICENSE-2.0.html

// @host      localhost:8080
// @BasePath  /api/v1

// @securityDefinitions.basic  BasicAuth
func main() {
	r := SetUpRouter()

	docs.SwaggerInfo.BasePath = "/api/v1"
	v1 := r.Group("/api/v1")
	{
		eg := v1.Group("items")
		{
			eg.Use(gin.Recovery())
			eg.Use(gin.Logger())

			eg.GET("metrics", handlers.PrometheusHandler())
			eg.GET("all", handlers.GetAllItems)
			eg.GET(":id", handlers.GetItemById)

			eg.POST("", handlers.Create)
			eg.PUT("", handlers.Update)
			eg.DELETE(":id", handlers.Delete)
		}

	}
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))

	r.Run()
}

func SetUpRouter() *gin.Engine {
	router := gin.Default()
	return router
}
