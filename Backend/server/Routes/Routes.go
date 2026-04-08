package Routes

import (
	"server/Controllers"

	"github.com/gin-gonic/gin"
)

//SetupRouter ... Configure routes

func SetupRouter() *gin.Engine {
 r := gin.Default()
 grp1 := r.Group("/user-api")
 {
  grp1.GET("todos", Controllers.GetTodos)
  grp1.POST("todos", Controllers.CreateTodo)
  grp1.GET("todos/:id", Controllers.GetTodobyID)
  grp1.PATCH("todos/:id", Controllers.UpdatTodo)
  grp1.DELETE("todos/:id", Controllers.DeleteTodo)
 }
 return r
}