package main

import (
	controllerAccount "github.com/vitelabs/vite-explorer-server/controller/account"
	controllerAccountchain "github.com/vitelabs/vite-explorer-server/controller/accountchain"
	controllerSnapshotchain "github.com/vitelabs/vite-explorer-server/controller/snapshotchain"
	controllerToken "github.com/vitelabs/vite-explorer-server/controller/token"

	"github.com/gin-gonic/gin"
	"flag"
	"github.com/vitelabs/vite-explorer-server/config"
	"github.com/vitelabs/vite-explorer-server/vitelog"
	"github.com/sirupsen/logrus"
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
	router.POST("chainheight", controllerSnapshotchain.ChainHeight)
}

func registerTokenRouter (engine *gin.Engine) {
	router := engine.Group("/api/token")

	router.POST("/list", controllerToken.List)
	router.GET("/detail", controllerToken.Detail)
}

func main ()  {
	var env string
	flag.StringVar(&env, "env", "dev", "env info")
	flag.Parse()

	viteconfig.LoadConfig(env)
	vitelog.InitLogger()

	router := gin.New()
	// Auto log
	if env == "production" {
		gin.SetMode(gin.ReleaseMode)
		router.Use(gin.LoggerWithWriter(vitelog.Logger.Writer()))
		router.Use(gin.RecoveryWithWriter(vitelog.Logger.WriterLevel(logrus.ErrorLevel)))
	} else {
		router.Use(gin.Logger())
		router.Use(gin.Recovery())
	}

	// Recover from error

	registerAccountRouter(router)

	registerAccountChainRouter(router)

	registerSnapshotChainRouter(router)

	registerTokenRouter(router)

	vitelog.Logger.Info("Server start listen in " + port)

	router.Run(":" + port)
}