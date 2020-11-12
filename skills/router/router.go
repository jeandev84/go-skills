package router


import (
	"github.com/gin-gonic/gin"
)



func New() *gin.Engine {
    
    // default route
	r := gin.Default()


	return r
}