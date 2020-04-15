package api

import (
	"github.com/Pantani/logger"
	_ "github.com/Pantani/replacer/docs"
	"github.com/Pantani/replacer/internal/config"
	"github.com/Pantani/replacer/internal/storage"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func Provide(storage *storage.Storage) {
	gin.SetMode(config.Configuration.Api.Mode)
	engine := gin.Default()

	makeRoutes(engine, storage)

	port := config.Configuration.Port
	logger.Info("Running application", logger.Params{"port": port})
	if err := engine.Run(":" + port); err != nil {
		logger.Fatal("Application failed", err)
	}
}

func makeRoutes(engine *gin.Engine, storage *storage.Storage) {
	engine.GET("swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	engine.GET("/status", statusHandler)

	// Names
	engine.PUT("/names/:name", insertName(storage))
	engine.GET("/names/:name", getName(storage))
	engine.GET("/names", getAllNames(storage))
	engine.DELETE("/names", deleteAllNames(storage))
	engine.DELETE("/names/:name", deleteName(storage))

	// Annotates
	engine.POST("/annotate", generateAnnotate(storage))
}
