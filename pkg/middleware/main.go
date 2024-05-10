package middleware

import (
	"Tickermaster/pkg/constants"
	"fmt"
	"net/http"
	"runtime"
	"strings"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func SetLogger(c *gin.Context) {

}

func Recover(c *gin.Context) {
	defer func() {
		if r := recover(); r != nil {
			if p, ok := r.(runtime.Error); ok {
				err := fmt.Errorf(fmt.Sprintf("panic error: %s", p))
				c.JSON(http.StatusInternalServerError, gin.H{"error": err})
				c.Abort()
			}
		}
	}()
	c.Next()
}

func SetDatabase(database *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Set(constants.FieldDatabase, database)
	}
}

func VerifyToken(c *gin.Context) {
	authorization := c.Request.Header.Get("Authorization")
	parts := strings.SplitN(authorization, " ", 2)
	if len(parts) != 2 || parts[0] != "Bearer" {
		c.JSON(http.StatusForbidden, "")
		c.Abort()
	}
	token := parts[1]
	c.Set(constants.AccessToken, token)

}
