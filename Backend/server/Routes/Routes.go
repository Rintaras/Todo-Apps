package Routes

import (
	"todo-apps/backend/server/Controllers"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

// SetupRouter ... Configure routes
func SetupRouter() *gin.Engine {
	r := gin.Default()

	// Swagger UI（別ポート）や Frontend（:3000）からブラウザで叩くため CORS を許可
	r.Use(cors.New(cors.Config{
		AllowAllOrigins:  true,
		AllowMethods:     []string{"GET", "POST", "PATCH", "PUT", "DELETE", "OPTIONS", "HEAD"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Accept", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: false,
	}))

	grp1 := r.Group("/api")
	{
		grp1.GET("todos", Controllers.GetTodos)
		grp1.POST("todos", Controllers.CreateTodo)
		grp1.GET("todos/:id", Controllers.GetTodoByID)
		grp1.PATCH("todos/:id", Controllers.UpdateTodo)
		grp1.DELETE("todos/:id", Controllers.DeleteTodo)
	}
	return r
}
