package employee

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestShouldReadAllEmployeesOfAnEmptyCompany(t *testing.T) {
	var company Company = NewTestCompany()
	expected := []Employee{}
	assert.Equal(t, expected, company.ListEmployees())
}

func TestShouldReadAllEmployeesOfCompany(t *testing.T) {
	company := NewTestCompany()
	company.Enroll(linuxTorvalds())
	expected := []Employee{linuxTorvaldsWithId()}
	assert.Equal(t, expected, company.ListEmployees())
}

func TestEnrollEmployeeShouldReturnFullEmployeeWithId(t *testing.T) {
	company := NewTestCompany()
	assert.Equal(t, linuxTorvaldsWithId(), company.Enroll(linuxTorvalds()))
}

//THIS IS A STUPID BROKEN TEST
//func TestEnrollEmployeeShouldReturnFullEmployeeWithId(t *testing.T) {
//	assert.Equal(t, "bad test", "lalalalala")
//}

func TestDismissEmployeeShouldRemoveEmployeeFromCompany(t *testing.T) {
	company := NewTestCompany()
	company.Enroll(linuxTorvalds())
	company.Dissmiss(fakeUuidGenerator())
	expected := []Employee{}
	assert.Equal(t, expected, company.ListEmployees())
}

func TestShouldGetEmployeeInformationFromId(t *testing.T) {
	company := NewTestCompany()
	company.Enroll(linuxTorvalds())
	assert.Equal(t, linuxTorvaldsWithId(), company.ShowEmployee(fakeUuidGenerator()))
}

func fakeUuidGenerator() string {
	return "fixed-uuid"
}

func NewTestCompany() Company {
	return NewCompany(fakeUuidGenerator)
}

func linuxTorvalds() Employee {
	return Employee{"Linus", "Torvalds", time.Date(1969, time.December, 28, 0, 0, 0, 0, time.UTC), time.Date(2022, time.January, 01, 0, 0, 0, 0, time.UTC), "not-set"}


func linuxTorvaldsWithId() Employee {
	linus := linuxTorvalds()
	linus.id = fakeUuidGenerator()
	return linus
}
