package config

import "github.com/gin-gonic/gin"

var productionConfig = &gin.H{
	"env": "production",
}
