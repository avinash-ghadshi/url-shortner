package response

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Response struct {
	ErrorCode    int
	ErrorMessage string
	Data         interface{}
}

func Send200(ctx *gin.Context, data interface{}) {
	var resp = Response{
		ErrorCode:    http.StatusOK,
		ErrorMessage: "",
		Data:         data,
	}
	ctx.JSON(http.StatusOK, resp)
}

func Send400(ctx *gin.Context, msg string) {
	var resp = Response{
		ErrorCode:    http.StatusBadRequest,
		ErrorMessage: msg,
		Data:         nil,
	}
	ctx.JSON(http.StatusBadRequest, resp)
}

func Send404(ctx *gin.Context) {
	var resp = Response{
		ErrorCode:    http.StatusNotFound,
		ErrorMessage: "Requested URL not found",
		Data:         nil,
	}
	ctx.JSON(http.StatusNotFound, resp)
}

func Send422(ctx *gin.Context) {
	var resp = Response{
		ErrorCode:    http.StatusUnprocessableEntity,
		ErrorMessage: "Failed to delete URL",
		Data:         nil,
	}
	ctx.JSON(http.StatusNotFound, resp)
}
