package middleware

import (
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/intaen/apigocovid/pkg/utils"
)

func RequestLoggerMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		log.Println(fmt.Sprintf("RequestID: %s, Method: %s, URI: %s", utils.GetRequestID(ctx), ctx.Request.Method, ctx.Request.URL.String()))
	}
}
