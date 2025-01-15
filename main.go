package main

import (
	"dev-crud/docs"
	"dev-crud/src/controllers"
	"dev-crud/src/database"
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"

	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func main(){
    enviroment := enviroment()
    
    database.Connect(enviroment["dbString"])
	database.Migrate()

    initDocs()
    router := router()
    router.Run(":" + enviroment["port"])
}

func router() *gin.Engine {
    router := gin.Default()

    api := router.Group("/api/v1")
    {
        dev := api.Group("/dev")
        {
            dev.POST("", controllers.CreateDev)
            dev.GET("", controllers.GetDevs)
            dev.GET("/:id", controllers.GetDevById)
            dev.DELETE("/:id", controllers.RemoveDev)
        }

        health := api.Group("/health")
        {
            health.GET("", controllers.Health)
        }
    }

    router.GET("/docs/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

    return router
}

func enviroment() map[string]string {
     err := godotenv.Load()

    if err != nil {
        log.Fatal("Error loading .env file")
    }

    return map[string]string{
        "port":  os.Getenv("PORT"),
        "dbString": os.Getenv("DB_STRING"),
    }
}

func initDocs() {
    docs.SwaggerInfo.Title = "Dev Crud API"
    docs.SwaggerInfo.Description = "This is a sample server."
    docs.SwaggerInfo.Version = "1.0"
    docs.SwaggerInfo.Host = "localhost"
    docs.SwaggerInfo.BasePath = "/api/v1"
    docs.SwaggerInfo.Schemes = []string{"http", "https"}
}