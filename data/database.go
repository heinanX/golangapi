package data

import (
	"errors"

	"github.com/glebarez/sqlite"
	"gorm.io/gorm"
)

var db *gorm.DB

func GetAllEmployees() []Employee {
	var employees []Employee
	db.Find(&employees) // SIMPLIFIED BY GORM aka. SELECT * FROM Employees
	return employees
}

func GetEmployee(id int) *Employee {
	var employee Employee
	err := db.First(&employee, id).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil
	}
	return &employee

}

func CreateNewEmployee(employee Employee) *Employee {
	db.Create(&employee)
	return &employee
}
func UpdateEmployee(employee Employee) bool {
	var dbEmployee Employee
	err := db.First(&dbEmployee, employee.Id).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return false
	}
	dbEmployee.Age = employee.Age
	dbEmployee.Name = employee.Name
	dbEmployee.City = employee.City
	db.Save(&employee)
	return true
}

// func CreateNewEmployee(newEmployee Employee) {
// 	employees = append(employees, newEmployee)
// }

func Init() {
	db, _ = gorm.Open(sqlite.Open("employees.sqlite"), &gorm.Config{})
	db.AutoMigrate(&Employee{}) // functionality: Is there a table inside the database, called Employee? If not, create one.
	// it also checks if there are columns that don't match (I might have added a new property to a table), sync them

	var count int64
	db.Model(&Employee{}).Count(&count)
	if count == 0 {
		db.Create(&Employee{Age: 2, Name: "Oliver", City: "Test"})
		db.Create(&Employee{Age: 5, Name: "Dodger", City: "New York"})
		db.Create(&Employee{Age: 4, Name: "Tito", City: "Mexico"})
	}
}
