package server

import (
	"database/sql"
	"log"

	"code.id.northwind/controllers"
	"code.id.northwind/repositories"
	"code.id.northwind/services"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

type HttpServer struct {
	config             *viper.Viper
	router             *gin.Engine
	categoryController *controllers.CategoryController
}

func InitHttpServer(config *viper.Viper, dbHandler *sql.DB) HttpServer {
	categoryRepository := repositories.NewCategoryRepository(dbHandler)

	categoryService := services.NewCategoryService(categoryRepository)

	categoryController := controllers.NewCategoryController(categoryService)

	router := gin.Default()

	// Router endpoint
	router.GET("/category", categoryController.GetListCategory)
	router.GET("/category/:id", categoryController.GetCategory)
	router.POST("/category", categoryController.CreateCategory)

	router.PUT("/category/:id", categoryController.UpdateCategory)
	router.DELETE("/category/:id", categoryController.DeleteCategory)

	return HttpServer{
		config:             config,
		router:             router,
		categoryController: categoryController,
	}
}

// Running gin HttpServer
func (hs HttpServer) Start() {
	err := hs.router.Run(hs.config.GetString("http.server_address"))

	if err != nil {
		log.Fatalf("Error while starting HTTP Server: %v", err)
	}
}
