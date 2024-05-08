package router

import (
	"Tickermaster/pkg/middleware"

	"github.com/gin-contrib/gzip"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func New(database *gorm.DB) *gin.Engine {
	gin.SetMode(gin.DebugMode)

	engine := gin.New()
	engine.Use(gzip.Gzip(gzip.DefaultCompression))
	engine.Use(middleware.SetLogger)
	if database != nil {
		engine.Use(middleware.SetDatabase(database))
	}
	engine.Use(middleware.Recover)

	// cors allow
	/*
		config := cors.Config{
			AllowMethods: []string{
				"PUT",
				"PATCH",
				"OPTIONS",
				"POST",
				"GET",
				"DELETE",
			},
			AllowHeaders: []string{
				"Content-Type",
				"X-Amz-Date",
				"Authorization",
				"X-Api-Key",
				"X-Amz-Security-Token",
			},
			AllowOrigins:     "",
			MaxAge:           15 * time.Minute,
			ExposeHeaders:    "",
			AllowCredentials: false,
		}
		engine.Use(cors.New(config))
	*/

	return engine
}
