package jwt

import (
	"github.com/gin-gonic/gin"
	cfg "github.com/dxcenter/chess/config"
	ginJwt "github.com/appleboy/gin-jwt"
	"time"
	"strings"
)

var jwt *ginJwt.GinJWTMiddleware

func GetJwtMiddleware() *ginJwt.GinJWTMiddleware {
	return jwt
}

func InitJwtMiddleware() {
	jwt = &ginJwt.GinJWTMiddleware{
		Realm:      "DXChess",
		Key:        []byte(cfg.Get().Secret),
		Timeout:    time.Hour,
		MaxRefresh: time.Hour,
		Authenticator: func(login string, password string, c *gin.Context) (string, bool) {
			login = strings.ToLower(login)
			for _, user := range cfg.Get().Users {
				if login == strings.ToLower(user.Login) && password == user.Password {
					return login, true
				}
			}

			return login, false
		},
		Authorizator: func(login string, c *gin.Context) bool {
			return true
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

