package middleware

import (
	"fmt"
	"log"

	"apigocovid/pkg/utils"

	"github.com/gin-gonic/gin"
)

func RequestLoggerMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		log.Println(fmt.Sprintf("RequestID: %s, Method: %s, URI: %s", utils.GetRequestID(ctx), ctx.Request.Method, ctx.Request.URL.String()))
	}
}
