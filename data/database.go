package data

var employees []Employee

func GetAllEmployees() []Employee {
	return employees
}

func GetEmployee(id int) *Employee {
	var employee Employee
	for _, employee = range employees {
		if employee.Id == id {
			return &employee
		}
	}
	return nil
}

func CreateNewEmployee(newEmployee Employee) {
	employees = append(employees, newEmployee)
}

func Init() {
	employees = append(employees, Employee{Id: 1, Age: 51, Name: "Oliver", City: "Test"})
	employees = append(employees, Employee{Id: 2, Age: 32, Name: "June", City: "Testy"})
	employees = append(employees, Employee{Id: 3, Age: 12, Name: "Dodger", City: "New York"})
}
