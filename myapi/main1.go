package main
import (
 "net/http"

 "github.com/gin-gonic/gin"
)
func helloHandler(c *gin.Context) {
 c.JSON(http.StatusOK, gin.H{
 "message": "hello",
 })
}

func getTodoHandler (c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, err)
		return
	}
	t, ok := todos[id]
	if !ok {
		c.JSON(http.StatusOK, gin,H{})
		return
		c.JSON(http.SttusOK,t)
	}
}
func main() {
 r := gin.Default()
 r.GET("/todos", getTodosHandler)
 r.GET("/todos,:id", getTodosHandler)
 r.POST("/todos",createTodosHandler)
 r.Run(":1234") 
}
