package utils

import "github.com/gin-gonic/gin"

func RespondJSON(c *gin.Context, status int, data interface{}, message string) {
	c.JSON(status, gin.H{
		"data":    data,
		"message": message,
		"status":  status,
	})
}
