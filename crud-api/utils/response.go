package utils

import "github.com/gin-gonic/gin"

func RespondSuccess(c *gin.Context, code int, data interface{}) {
	c.JSON(code, gin.H{"success": true, "data": data})
}

func RespondError(c *gin.Context, code int, message string) {
	c.JSON(code, gin.H{"success": false, "error": message})
}
