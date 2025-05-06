package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"todolist/db"
	"todolist/models"
)

func main() {
	db.InitDb()
	router := gin.Default()
	router.POST("/addTodo", createTodo)
	router.GET("/getTodos", getTodos)
	router.GET("/getTodo/:id", getTodoById)
	router.PUT("/edit/:id", updateTodo)
	router.DELETE("/delete/:id", deleteTodoById)
	router.POST("/login", login)
	router.Run(":8080")

}

func createTodo(c *gin.Context) {
	var todo models.Todo
	err := c.BindJSON(&todo)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	err = todo.AddTodo()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not add todo"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Todo successfully created"})
}

func getTodos(c *gin.Context) {
	todos, err := models.GetAllTodos()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not get todos"})
	}
	c.JSON(http.StatusOK, gin.H{"todos": todos})
}

func updateTodo(c *gin.Context) {
	idParam := c.Param("id")
	id, err := strconv.ParseInt(idParam, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid id"})
		return
	}

	var todo models.Todo
	err = c.BindJSON(&todo)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	err = models.EditTodo(int(id), todo) // or change EditTodo to accept int64
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not update todo"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Todo updated successfully"})
}

func getTodoById(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	todo, err := models.GetTodoId(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Todo not found"})
		return
	}

	c.JSON(http.StatusOK, todo)
}

func deleteTodoById(c *gin.Context) {
	idstr := c.Param("id")
	id, err := strconv.ParseInt(idstr, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid id"})
	}
	err = models.DeleteTodo(int(id))
	if err != nil {
		return
	}
}

func login(c *gin.Context) {
	var user models.User
	err := c.BindJSON(&user)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
	}
	user.AddUser()
}
