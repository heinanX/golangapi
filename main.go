package main

import (
	"heinanx/goapi20250224/data"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func handleGetAllEmployees(c *gin.Context) {
	emps := data.GetAllEmployees()
	c.IndentedJSON(http.StatusOK, emps)
}

func handleGetOneEmployee(c *gin.Context) {
	id := c.Param("id")
	numId, _ := strconv.Atoi(id)
	employee := data.GetEmployee(numId)

	if employee == nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"error": "no existing file"})
	} else {
		c.IndentedJSON(http.StatusOK, employee)
	}
}

func main() {
	data.Init()

	router := gin.Default()

	router.GET("/api/employee", handleGetAllEmployees)
	router.GET("/api/employee/:id", handleGetOneEmployee)

	router.Run()
}
