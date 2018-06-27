package utils

import (
	"github.com/gin-gonic/gin"
)

func Respond (c *gin.Context, data interface{}, msg string, err error, code int) {
	c.JSON(200, gin.H{
		"code": code, // success
		"msg": msg,
		"data": data,
		"error": err,
	})
}

func RespondSuccess (c *gin.Context, data interface{}, msg string)  {
	Respond(c, data, msg, nil, 0)
}

func RespondFailed (c *gin.Context, code int,err error , msg string)  {
	Respond(c, nil, msg, err, code)
}


func RespondError (c *gin.Context, statusCode int, err error )  {
	c.String(statusCode,  err.Error())
}
