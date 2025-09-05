// package middleware

// import (
// 	"strings"

// 	"github.com/JacobGeorgeMathew/GO_Notes_CRED_Application_Backend/utils"
// 	"github.com/gofiber/fiber/v2"
// )

// func AuthMiddleware() fiber.Handler {
// 	return func(c *fiber.Ctx) error {
// 		// Get token from Authorization header
// 		authHeader := c.Get("Authorization")
// 		if authHeader == "" {
// 			return c.Status(401).JSON(fiber.Map{
// 				"error": "Authorization header required",
// 			})
// 		}

// 		// Check if header starts with "Bearer "
// 		tokenString := strings.TrimPrefix(authHeader, "Bearer ")
// 		if tokenString == authHeader {
// 			return c.Status(401).JSON(fiber.Map{
// 				"error": "Bearer token required",
// 			})
// 		}

// 		// Validate token
// 		claims, err := utils.ValidateJWT(tokenString)
// 		if err != nil {
// 			return c.Status(401).JSON(fiber.Map{
// 				"error": "Invalid token",
// 			})
// 		}

// 		// Store user ID in context
// 		c.Locals("userID", claims.UserID)
// 		return c.Next()
// 	}
// }

package middleware

import (
	"time"
	"github.com/JacobGeorgeMathew/GO_Notes_CRED_Application_Backend/utils"
	"github.com/gofiber/fiber/v2"
)

// 🍪 Cookie-based Auth Middleware
func AuthMiddleware() fiber.Handler {
	return func(c *fiber.Ctx) error {
		// Try to get token from cookie first
		tokenString := c.Cookies("jwt_token")
		
		// If no cookie, check Authorization header as fallback
		if tokenString == "" {
			authHeader := c.Get("Authorization")
			if authHeader != "" {
				// Remove "Bearer " prefix
				if len(authHeader) > 7 && authHeader[:7] == "Bearer " {
					tokenString = authHeader[7:]
				}
			}
		}
		
		// If still no token, return error
		if tokenString == "" {
			return c.Status(401).JSON(fiber.Map{
				"error": "Authentication required",
			})
		}

		// Validate token
		claims, err := utils.ValidateJWT(tokenString)
		if err != nil {
			return c.Status(401).JSON(fiber.Map{
				"error": "Invalid token",
			})
		}

		// Store user ID in context
		c.Locals("userID", claims.UserID)
		return c.Next()
	}
}

// 🔥 Alternative: Cookie-only auth middleware (more secure)
func CookieAuthMiddleware() fiber.Handler {
	return func(c *fiber.Ctx) error {
		// Only accept tokens from cookies (no Authorization header)
		tokenString := c.Cookies("jwt_token")
		
		if tokenString == "" {
			return c.Status(401).JSON(fiber.Map{
				"error": "Authentication cookie required",
			})
		}

		// Validate token
		claims, err := utils.ValidateJWT(tokenString)
		if err != nil {
			// Clear invalid cookie
			c.Cookie(&fiber.Cookie{
				Name:     "jwt_token",
				Value:    "",
				Expires:  time.Now().Add(-time.Hour),
				HTTPOnly: true,
			})
			
			return c.Status(401).JSON(fiber.Map{
				"error": "Invalid authentication cookie",
			})
		}

		// Store user ID in context
		c.Locals("userID", claims.UserID)
		return c.Next()
	}
}