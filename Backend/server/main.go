package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)


func testHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
			"message": "ok",
		})
		// if err := c.ShouldBindJSON($req) != nil
		// c.JSON(http.statusBadRequest, gin.H{"error" : "invalid request"})
}

func loginHandler(c *gin.Context) {
	var req Auth
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error" : "invalid request"})
		return
	}
	fmt.Println("Username", req.Username)
	fmt.Println("Password", req.Password)
	
	c.JSON(http.StatusOK, gin.H{
	"message":"login success!",
	"Username":req.Username,
})
}

type Auth struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func main() {
	router := gin.Default()
	router.GET("/test", testHandler)
	router.POST("/login", loginHandler)
	router.Run()
}