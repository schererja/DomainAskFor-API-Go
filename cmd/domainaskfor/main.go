package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/schererja/domainaskfor-go/internal/app/domainaskfor/webapi"
)

func main() {
	log.Print("Starting DomainAskFor")
	engine := gin.Default()
	webapi.Route(engine)
	engine.Run()
}
