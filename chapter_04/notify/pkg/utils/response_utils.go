package utils

import (
	"github.com/gin-gonic/gin"
	"net/http"
)
import "notify/api/server/service/error"

func WriteErr(c *gin.Context, err cerror.BError) {
	c.JSON(err.Code, gin.H{
		"code":   err.Err.BCode,
		"reason": err.Err.Reason,
		"data":   err.Err.Data,
	})
}

func WriteData(c *gin.Context, data interface{}) {
	c.JSON(http.StatusOK, gin.H{
		"code":   "0",
		"reason": "success",
		"data":   data,
	})
}
