package middlewares

import (
	"github.com/gin-gonic/gin"
)

func AddCorsMiddleware(ApiGatewayrUrl string) gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Header("Access-Control-Allow-Origin", ApiGatewayrUrl)
		c.Header("Access-Control-Allow-Credentials", "true")
		c.Header("Access-Control-Allow-Headers",
			`Content-Type, 
			Content-Length, 
			Accept-Encoding, 
			Authorization, 
			accept, 
			origin, 
			Cache-Control`)
		c.Header("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")

		c.Next()
	}
}
