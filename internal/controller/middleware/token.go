package middleware

import (
	"net/http"
	"strings"
	"time"

	"github.com/evrone/go-clean-template/internal/entity"
	"github.com/evrone/go-clean-template/internal/usecase"
	"github.com/evrone/go-clean-template/pkg/hashing"
	"github.com/evrone/go-clean-template/pkg/logger"
	"github.com/gin-gonic/gin"
)

var unauthorizedResponse = gin.H{"errors": nil, "message": "Unauthorized", "success": false}

func CheckAuthToken(u usecase.User) gin.HandlerFunc {
	return func(c *gin.Context) { // Get token from header
		tokenString := c.Request.Header.Get("Authorization")
		if tokenString == "" {
			c.AbortWithStatusJSON(http.StatusUnauthorized, unauthorizedResponse)
			return
		}
		if !strings.Contains(tokenString, "Bearer") {
			c.AbortWithStatusJSON(http.StatusUnauthorized, unauthorizedResponse)
			return
		}
		tokenString = strings.Replace(tokenString, "Bearer ", "", 1)

		prefix := tokenString[0:32]
		expTime := tokenString[32:46]
		suffix := tokenString[46:]

		t, _ := time.ParseInLocation("20060102150405", expTime, time.Local)

		expired := time.Now().Before(t)
		if !expired {
			logger.New("debug").Info("Expired")
			c.AbortWithStatusJSON(http.StatusUnauthorized, unauthorizedResponse)
			return
		}

		user, err := u.FindUserByPassword(c, suffix)

		if user == (entity.User{}) {
			logger.New("debug").Info("Not found user")
			c.AbortWithStatusJSON(http.StatusUnauthorized, unauthorizedResponse)
			return
		} else if err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, unauthorizedResponse)
			return
		}

		if hashing.Md5(user.Email) != prefix {
			logger.New("debug").Info("Email not match")
			c.AbortWithStatusJSON(http.StatusUnauthorized, unauthorizedResponse)
			return
		}

		c.Next()
	}
}
