package middleware

import (
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/upeshchalise/go_blogs/pkg/utils/jwt"
)

func VerifyJWT(Next gin.HandlerFunc) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		token := ctx.GetHeader("Authorization")
		if token == "" {
			ctx.JSON(401, gin.H{"error": "Authorization header is missing"})
			ctx.Abort()
			return
		}

		if !strings.HasPrefix(token, "Bearer ") {
			ctx.JSON(401, gin.H{"error": "Authorization header format must be Bearer {token}"})
			ctx.Abort()
			return
		}
		token = strings.TrimPrefix(token, "Bearer ")

		token, err := jwt.ValidateJwtToken(token, "thisissecrettoken")
		if err != nil {
			ctx.JSON(401, gin.H{"error": "Invalid token"})
			ctx.Abort()
			return
		}
		ctx.Set("userId", token)
		Next(ctx)
	}
}

// CORS middleware function definition
func CorsMiddleware() gin.HandlerFunc {
	// Define allowed origins as a comma-separated string
	originsString := "http://localhost:3000"
	var allowedOrigins []string
	if originsString != "" {
		// Split the originsString into individual origins and store them in allowedOrigins slice
		allowedOrigins = strings.Split(originsString, ",")
	}

	// Return the actual middleware handler function
	return func(c *gin.Context) {
		// Function to check if a given origin is allowed
		isOriginAllowed := func(origin string, allowedOrigins []string) bool {
			for _, allowedOrigin := range allowedOrigins {
				if origin == allowedOrigin {
					return true
				}
			}
			return false
		}

		// Get the Origin header from the request
		origin := c.Request.Header.Get("Origin")

		// Check if the origin is allowed
		if isOriginAllowed(origin, allowedOrigins) {
			// If the origin is allowed, set CORS headers in the response
			c.Writer.Header().Set("Access-Control-Allow-Origin", origin)
			c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
			c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
			c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT")
		}

		// Handle preflight OPTIONS requests by aborting with status 204
		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		// Call the next handler
		c.Next()
	}
}
