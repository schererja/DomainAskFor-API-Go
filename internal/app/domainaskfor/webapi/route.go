package webapi

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/schererja/domainaskfor-go/internal/app/domainaskfor/webhandler"
)

func Route(router *gin.Engine) *gin.Engine {
	router.Use(cors.Default())
	router.GET("/ping", func(context *gin.Context) {
		context.JSON(200, gin.H{
			"message": "pong",
		})
	})

	v1WhoIs := router.Group("/v1/whois")
	{
		v1WhoIs.GET("/:domain", webhandler.ShowHandler)
	}
	v1Synonyms := router.Group("/v1/synonyms")
	{
		v1Synonyms.GET("/:word", webhandler.ShowSynonymsHandler)
	}

	return router
}
