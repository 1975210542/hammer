package middleware

import(
	"github.com/gin-gonic/gin"
	"fmt"
	"github.com/appleboy/gin-jwt"
	"time"
)


func InitJwtMiddleware(c *gin.Context) *jwt.GinJWTMiddleware{
	return &jwt.GinJWTMiddleware{
		Realm:      "test zone",
		Key:        []byte("secret key"),
		Timeout:    time.Hour,
		MaxRefresh: time.Hour,
		Authenticator: func(userID string, password string, c *gin.Context) (string, bool) {
			if (userID == "admin" && password == "admin") || (userID == "test" && password == "test") {
				return userID, true
			}

			return userID, false
		},
		Authorizator: func(userID string, c *gin.Context) bool {
			if userID == "admin" {
				return true
			}

			return false
		},
		Unauthorized: func(c *gin.Context, code int, message string) {
			c.JSON(code, gin.H{
				"code":    code,
				"message": message,
			})
		},
		// TokenLookup is a string in the form of "<source>:<name>" that is used
		// to extract token from the request.
		// Optional. Default value "header:Authorization".
		// Possible values:
		// - "header:<name>"
		// - "query:<name>"
		// - "cookie:<name>"
		TokenLookup: "header:Authorization",
		// TokenLookup: "query:token",
		// TokenLookup: "cookie:token",

		// TokenHeadName is a string in the header. Default value is "Bearer"
		TokenHeadName: "Bearer",

		// TimeFunc provides the current time. You can override it to use another time value. This is useful for testing or if your server uses a different time zone than your tokens.
		TimeFunc: time.Now,
	}
}



func Auth(c *gin.Context) {
	InitJwtMiddleware(c).LoginHandler(c)
	fmt.Println("In auth middleware")
	c.Next()
}
