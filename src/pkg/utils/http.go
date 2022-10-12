package utils

import (
	"github.com/gin-contrib/requestid"
	"github.com/gin-gonic/gin"
)

func GetRequestID(c *gin.Context) string {
	return requestid.Get(c)
}
