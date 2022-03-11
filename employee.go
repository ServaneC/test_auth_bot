package employee

import "time"

type Company struct {
	employees     map[string]Employee
	uuidGenerator func() string
}

func (c Company) ListEmployees() []Employee {
	output := []Employee{}
	for _, employee := range c.employees {
		output = append(output, employee)
	}
	return output
}

func NewCompany(uuidGenerator func() string) Company {
	return Company{map[string]Employee{}, uuidGenerator}
}

func (c *Company) Enroll(employee Employee) Employee {
	employee.id = c.uuidGenerator()
	c.employees[employee.id] = employee
	return employee
}

func (c *Company) Dissmiss(uuid string) {
	delete(c.employees, uuid)
}

func (c Company) ShowEmployee(uuid string) Employee {
	return c.employees[uuid]
}

func (e Employee) Id() string {
	return e.id
}

type Employee struct {
	FirstName string
	LastName  string
	Birthday  time.Time
	EntryDate time.Time
	id        string
}
