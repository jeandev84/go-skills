package router



import (
	"net/http"

	"github.com/gin-gonic/gin"
)


// gin.Context (хранится много информации) it as Container

func indexHandler(c *gin.Context) {

 	  c.String(http.StatusOK, "Hello world!")
}


func collectHandler(c *gin.Context) {

      // получить категории
      category := c.Param("category")


      news.Collect(caterory)


      c.String(http.StatusOK, "search is initialized")

}


func resultHandler(c *gin.Context) {
      
      // из какой категории мы хотим получить результат
      category := c.Param("category")
      

      // получаем topics
      topics := news.Result(category)


      // Возврашаем в формате JSON
      c.JSON(http.StatusOK, topics)
}


