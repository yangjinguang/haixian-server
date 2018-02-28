package xRes

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func OK(ct *gin.Context, data interface{}) {
	ct.JSON(http.StatusOK, gin.H{
		"success": true,
		"data":    data,
	})
}

func BadRequest(ct *gin.Context, message interface{}) {
	if message == nil {
		message = "Bad Request"
	}
	ct.JSON(http.StatusBadRequest, gin.H{
		"success": false,
		"message": message,
	})
	ct.Abort()
}

func Unauthorized(ct *gin.Context, message interface{}) {
	if message == nil {
		message = "Unauthorized"
	}
	ct.JSON(http.StatusUnauthorized, gin.H{
		"success": false,
		"message": message,
	})
	ct.Abort()
}

func Created(ct *gin.Context, message interface{}) {
	if message == nil {
		message = "Created Success"
	}
	ct.JSON(http.StatusCreated, gin.H{
		"success": false,
		"message": message,
	})
	ct.Abort()
}

func MethodNotAllowed(ct *gin.Context, message interface{}) {
	if message == nil {
		message = "Method Not Allowed"
	}
	ct.JSON(http.StatusMethodNotAllowed, gin.H{
		"success": false,
		"message": message,
	})
	ct.Abort()
}

func NotFound(ct *gin.Context, message interface{}) {
	if message == nil {
		message = "Resource Not Found"
	}
	ct.JSON(http.StatusNotFound, gin.H{
		"success": false,
		"message": message,
	})
	ct.Abort()
}

func Forbidden(ct *gin.Context, message interface{}) {
	if message == nil {
		message = "Resource Forbidden"
	}
	ct.JSON(http.StatusForbidden, gin.H{
		"success": false,
		"message": message,
	})
	ct.Abort()
}
