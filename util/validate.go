package util

import (
	"errors"
	"fmt"
	"regexp"
	"rest/gin/models"
)

func ValidateMobileNumber(phone string) bool {

	const phoneregex = `^\d{10}$`
	return regexp.MustCompile(phoneregex).MatchString(phone)
}

func ValidateName(name string) bool {

	const nameregex = `^[a-zA-Z]+(([',. -][a-zA-Z ])?[a-zA-Z]*)*$`
	fmt.Print(name, regexp.MustCompile(nameregex).MatchString(name))
	return regexp.MustCompile(nameregex).MatchString(name)
}

func ValidateAge(age int) bool {
	return age >= 0 && age <= 120
}

func ValidateSalary(salary int) bool {
	return salary >= 0
}
func EmployeeValidation(employee *models.CreateEmployee) error {

	if !ValidateName(employee.Name) {
		return errors.New("invalid name")
	}
	if !ValidateMobileNumber(employee.Mobile) {
		return errors.New("invalid mobile number")
	}
	if !ValidateAge(employee.Age) {
		return errors.New("invalid age")
	}
	if !ValidateSalary(employee.Salary) {
		return errors.New("invalid salary")
	}
	return nil

}
