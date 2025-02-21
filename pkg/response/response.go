package response

import "github.com/gin-gonic/gin"

type Response struct {
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"`
	Error   string      `json:"error,omitempty"`
}

func JSON(ctx *gin.Context, status int, data interface{}) {
	ctx.JSON(status, Response{
		Message: "success",
		Data:    data,
	})
}

func Error(ctx *gin.Context, status int, err string) {
	ctx.JSON(status, Response{
		Message: "fail",
		Error:   err,
	})
}
