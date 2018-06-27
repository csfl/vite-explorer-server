package main

import (
	controllerAccount "vite-explorer-server/controller/account"
	controllerAccountchain "vite-explorer-server/controller/accountchain"
	controllerSnapshotchain "vite-explorer-server/controller/snapshotchain"
	controllerToken "vite-explorer-server/controller/token"

	"github.com/gin-gonic/gin"
)

var (
	port = "8081"
)

func registerAccountRouter(engine *gin.Engine) {
	router := engine.Group("/api/account")

	router.GET("/detail", controllerAccount.Detail)
}

func registerAccountChainRouter(engine *gin.Engine)  {
	router := engine.Group("/api/accountchain")

	router.POST("/blocklist", controllerAccountchain.BlockList)
	router.GET("/block", controllerAccountchain.Block)
}

func registerSnapshotChainRouter(engine *gin.Engine)  {
	router := engine.Group("/api/snapshotchain")

	router.POST("/blocklist", controllerSnapshotchain.BlockList)
	router.GET("/block", controllerSnapshotchain.Block)
}

func registerTokenRouter (engine *gin.Engine) {
	router := engine.Group("/api/token")

	router.POST("/list", controllerToken.List)
}


func main ()  {
	router := gin.New()

	// Auto log
	router.Use(gin.Logger())

	// Recover from error
	router.Use(gin.Recovery())

	registerAccountRouter(router)

	registerAccountChainRouter(router)

	registerSnapshotChainRouter(router)

	registerTokenRouter(router)

	router.Run(":" + port)
}
