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

func handleCreateEmployee(c *gin.Context) {
	var employee data.Employee
	if err := c.BindJSON(&employee); err != nil {
		return
	}
	employee.Id = 0
	data.CreateNewEmployee(employee)
	c.IndentedJSON(http.StatusCreated, employee)
}

func handleUpdateEmployeeById(c *gin.Context) {
	id := c.Param("id")
	var employee data.Employee
	if err := c.BindJSON(&employee); err != nil {
		return
	}
	employee.Id, _ = strconv.Atoi(id)

	if !data.UpdateEmployee(employee) {
		c.IndentedJSON(http.StatusNotFound, gin.H{"error": "not found"})
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
	router.POST("/api/employees/", handleCreateEmployee)
	router.PUT("/api/employees/:id", handleUpdateEmployeeById)

	router.Run()
}
