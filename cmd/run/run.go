package run

import (
	"github.com/vitelabs/vite-explorer-server"
	"github.com/gin-gonic/gin"
)

func Run (options gin.H)  {
	//if _, ok := options[]
	viteserver.StartServer(options)
}