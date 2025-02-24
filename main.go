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

type PageView struct {
	Title   string
	Heading string
}

func handleStartPage(c *gin.Context) {
	c.HTML(http.StatusOK, "index.html", &PageView{Title: "GolangApi", Heading: "Welcome to my first Golang api"})
	//c.String(http.StatusOK, "hello world")
}

func main() {
	data.Init()

	router := gin.Default()
	router.LoadHTMLGlob("templates/**")

	router.GET("/", handleStartPage)
	router.GET("/api/employees", handleGetAllEmployees)
	router.GET("/api/employees/:id", handleGetOneEmployee)

	router.Run()
}
