package server

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func InternalServerErrorResponse(err error) gin.H {
	return gin.H{
		"error": fmt.Sprintf("internal server error: %v", err.Error()),
	}
}

func NotParamErrorResponse(param string) gin.H {
	return gin.H{
		"error": fmt.Sprintf("The param '%s' not found", param),
	}
}

func NotFoundObjectErrorResponse(obj string) gin.H {
	return gin.H{
		"error": fmt.Sprintf("%s not found", obj),
	}
}
