package middlewares

import (
	"strings"

	"github.com/RidwanRamdhani/chronos-laundry/backend/utils"
	"github.com/gin-gonic/gin"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")

		if authHeader == "" {
			utils.Unauthorized(c, "Authorization header missing")
			c.Abort()
			return
		}

		parts := strings.Split(authHeader, " ")
		if len(parts) != 2 || strings.ToLower(parts[0]) != "bearer" {
			utils.Unauthorized(c, "Invalid Authorization header format")
			c.Abort()
			return
		}

		tokenString := parts[1]

		claims, err := utils.VerifyToken(tokenString)
		if err != nil {
			utils.Unauthorized(c, "Invalid or expired token: "+err.Error())
			c.Abort()
			return
		}

		c.Set("admin_id", claims.AdminID)
		c.Set("admin_username", claims.Username)
		c.Set("admin_email", claims.Email)
		c.Set("admin_full_name", claims.FullName)

		c.Next()
	}
}
