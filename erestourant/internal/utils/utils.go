package utils

import (
	"github.com/gin-gonic/gin"
)

func RespondError(ctx *gin.Context, statusCode int, errCode string, message string, details ...string) {
	resp := gin.H{
		"error":   errCode,
		"message": message,
	}

	if len(details) > 0 {
		resp["details"] = details[0]
	}

	ctx.JSON(statusCode, resp)
}
