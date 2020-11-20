package main

import "github.com/gin-gonic/gin"

func main() {
	r := gin.Default()
	r.GET("/users/:name", getUsersHandler)
	r.POST("/users", addUserHandler)
	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}

func getUsersHandler(c *gin.Context) {
	name := c.Param("name")
	lastname := c.Query("lastname")
	c.JSON(200, gin.H{
		"name": name + " " + lastname,
	})
}

func addUserHandler(c *gin.Context) {
	c.JSON(201, gin.H{
		"name":     "lore",
		"lastname": "Gonzalez",
	})
}
