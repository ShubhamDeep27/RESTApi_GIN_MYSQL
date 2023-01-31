package util

import (
	"errors"
	"regexp"
	"rest/gin/models"
	"unicode"
)

func ValidateMobileNumber(phone string) bool {

	const phoneregex = `^\d{10}$`
	return regexp.MustCompile(phoneregex).MatchString(phone)
}

func ValidateName(name string) bool {

	const nameregex = `^[a-zA-Z]+(([',. -][a-zA-Z ])?[a-zA-Z]*)*$`
	return regexp.MustCompile(nameregex).MatchString(name)
}

func ValidateAge(age int) bool {
	return age >= 0 && age <= 120
}

func ValidateSalary(salary int) bool {
	return salary >= 0
}
func ValidatePassword(pwd string) bool {

	var (
		hasMinLen  = false
		hasUpper   = false
		hasLower   = false
		hasNumber  = false
		hasSpecial = false
	)
	if len(pwd) >= 7 {
		hasMinLen = true
	}
	for _, char := range pwd {
		switch {
		case unicode.IsUpper(char):
			hasUpper = true
		case unicode.IsLower(char):
			hasLower = true
		case unicode.IsNumber(char):
			hasNumber = true
		case unicode.IsPunct(char) || unicode.IsSymbol(char):
			hasSpecial = true
		}
	}
	return hasMinLen && hasUpper && hasLower && hasNumber && hasSpecial
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
	if !ValidatePassword(employee.Password) {
		return errors.New("invalid Password")
	}
	return nil

}
