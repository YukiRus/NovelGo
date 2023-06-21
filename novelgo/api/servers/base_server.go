package servers

import "github.com/gin-gonic/gin"

type BaseServer interface {
	Ok(c gin.Context) Response
}

type BaseServices struct {
	context gin.Context
}

type Response struct {
	Code    int
	Message string
	Data    interface{}
}

func (bs *BaseServices) Ok(data interface{}) Response {
	return Response{
		Code:    200,
		Message: "success",
		Data:    data,
	}
}
