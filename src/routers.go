package main

import (
	"urlshortener/server/libs/common"
	"urlshortener/server/middleware/logger"
	ushort "urlshortener/server/modules/urlshortner"

	"github.com/gin-gonic/gin"
)

func setupRouter() *gin.Engine {

	gin.SetMode(gin.DebugMode)

	r := gin.New()
	r.Use(gin.Recovery(), logger.Logger())
	// r.Use(gin.Recovery(), logger.Logger(), gindump.Dump())
	r.SetTrustedProxies(common.Config.App.Proxy)
	r.HandleMethodNotAllowed = true

	v1 := r.Group("/api")
	ushort.UshortApis(v1.Group("/urlshort"))

	return r
}
