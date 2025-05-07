package response

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func SuccessResponse(c *gin.Context, data any) {
	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"message": "success",
		"data":    data,
	})
}

func CreatedResponse(c *gin.Context, data any) {
	c.JSON(http.StatusCreated, gin.H{
		"success": true,
		"message": "created",
		"data":    data,
	})
}

func UpdateResponse(c *gin.Context, data any) {
	c.JSON(http.StatusAccepted, gin.H{
		"success": true,
		"message": "updated",
		"data":    data,
	})
}

func DeletedResponse(c *gin.Context) {
	c.JSON(http.StatusNoContent, gin.H{})
}

func ErrorResponse(c *gin.Context, error string) {
	c.JSON(http.StatusBadRequest, gin.H{
		"success": false,
		"error":   error,
		"data":    nil,
	})
}

func ValidationErrorResponse(c *gin.Context, error string) {
	c.AbortWithStatusJSON(http.StatusUnprocessableEntity, gin.H{
		"success": false,
		"error":   error,
		"data":    nil,
	})
}

func AuthenticationErrorResponse(c *gin.Context, error string) {
	c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
		"success": false,
		"error":   error,
		"data":    nil,
	})
}

func AuthorizationErrorResponse(c *gin.Context, error string) {
	c.AbortWithStatusJSON(http.StatusForbidden, gin.H{
		"success": false,
		"error":   error,
		"data":    nil,
	})
}
