package middleware

import (
	"golang-firestore/config"

	"github.com/gin-gonic/gin"
)

/**
 * @brief      this function for middleware api authentication
 * @return     object
 */
func ApiAuthentication() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With, User-Agent, x-token, client-secret, Subdomain, subdomain, access-control-allow-origin, access-control-allow-headers")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST,DELETE,GET,PUT")
		c.Writer.Header().Set("Content-Type", "application/json")

		key := c.Request.Header.Get("app-key")

		// check if key not nil
		if key != "" {
			if key != config.Env("APP_KEY") {
				c.AbortWithStatus(403)
				return
			}
			c.Next()
		} else {
			c.AbortWithStatus(403)
			return
		}
	}
}
