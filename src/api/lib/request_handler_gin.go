package lib

import "github.com/gin-gonic/gin"

type RequestHandlerGin struct {
	Gin *gin.Engine
}

func NewRequestHandlerGin() RequestHandlerGin {
	engine := gin.Default()
	return RequestHandlerGin{Gin: engine}
}