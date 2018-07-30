package middlewares

import (
	"github.com/vitelabs/go-vite/vite"
	"github.com/gin-gonic/gin"
	"github.com/vitelabs/go-vite/common/types"
)

func SetVite (vite *vite.Vite, genesisAddr types.Address) gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Set("vite", vite)
		c.Set("genesisAddr", genesisAddr)
		c.Next()
	}
}
