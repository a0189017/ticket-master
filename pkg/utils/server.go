package utils

import (
	"Tickermaster/pkg/config"
	"Tickermaster/pkg/constants"
	"Tickermaster/pkg/errors"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func ParseSystemError(c *gin.Context, err error) (systemError errors.SystemError) {
	if h, ok := err.(errors.SystemError); ok {
		systemError = h
	} else {
		systemError = errors.New(errors.ErrorInfo{Err: err})
	}
	c.Set(constants.StackTrace, systemError.StackTrace)
	// stacktrace should not show in test&prod environment
	if !config.GetConfig().Settings.DebugMode {
		systemError.StackTrace = nil
	}
	return
}

func Wrapper(c *gin.Context, handlerFunc func(c *gin.Context, d *gorm.DB) error, event ...string) {
	d := c.MustGet(constants.FieldDatabase)
	err := handlerFunc(c, d.(*gorm.DB))
	if err != nil {
		systemError := ParseSystemError(c, err)

		c.JSON(systemError.StatusCode, gin.H{"error": systemError})
		c.Abort()
	} else {
		c.Next()
	}
}
