package utils

import (
	"Tickermaster/pkg/constants"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func Wrapper(c *gin.Context, handlerFunc func(c *gin.Context, d *gorm.DB) error, event ...string) {
	d := c.MustGet(constants.FieldDatabase)
	err := handlerFunc(c, d.(*gorm.DB))
	if err != nil {

		c.JSON(http.StatusBadRequest, gin.H{"error": err})
		c.Abort()
	} else {
		c.Next()
	}
}
