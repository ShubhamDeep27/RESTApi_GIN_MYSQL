package util

import (
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

func EmployeeValidation(employee *models.CreateEmployee) models.ErrorResponse {

	if !ValidateName(employee.Name) {
		return models.ErrorResponse{ErrorCode: "5031", ErrorMsg: GET_ERROR_MSGS["5031"]}
	}
	if !ValidateMobileNumber(employee.Mobile) {
		return models.ErrorResponse{ErrorCode: "5032", ErrorMsg: GET_ERROR_MSGS["5032"]}
	}
	if !ValidateAge(employee.Age) {
		return models.ErrorResponse{ErrorCode: "5033", ErrorMsg: GET_ERROR_MSGS["5033"]}
	}
	if !ValidateSalary(employee.Salary) {
		return models.ErrorResponse{ErrorCode: "5034", ErrorMsg: GET_ERROR_MSGS["5034"]}
	}
	if !ValidatePassword(employee.Password) {
		return models.ErrorResponse{ErrorCode: "5035", ErrorMsg: GET_ERROR_MSGS["5035"]}
	}
	return models.ErrorResponse{}

}
