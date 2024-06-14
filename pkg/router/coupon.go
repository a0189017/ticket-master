package router

import (
	"Tickermaster/pkg/utils"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func CouponList(c *gin.Context) {
	handler := func(c *gin.Context, db *gorm.DB) error {
		c.JSON(200, gin.H{
			"message": "CouponList",
		})
		return nil
	}
	utils.Wrapper(c, handler)
}

func CouponRegister(c *gin.Context) {

}

func CouponGrab(c *gin.Context) {

}
