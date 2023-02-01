package models

type Employee struct {
	ID       int64  `json:"id"`
	Name     string `json:"name"`
	Mobile   string `json:"mobile"`
	Address  string `json:"address"`
	Salary   int    `json:"salary"`
	Age      int    `json:"age"`
	Username string `json:"username"`
	Password string `json:"password"`
}

type CreateEmployee struct {
	Name     string `json:"name"`
	Mobile   string `json:"mobile"`
	Address  string `json:"address"`
	Salary   int    `json:"salary"`
	Age      int    `json:"age"`
	Username string `json:"username"`
	Password string `json:"password"`
}
type GetAllEmployeeResponse struct {
	Data    []*Employee `json:"data"`
	Status  string      `json:"status"`
	Message string      `json:"message"`
}
type GetEmployeeByIdResponse struct {
	Data    *Employee `json:"data"`
	Status  string    `json:"status"`
	Message string    `json:"message"`
}
type CreateEmployeeResponse struct {
	Data    *CreateEmployee `json:"data"`
	Status  string          `json:"status"`
	Message string          `json:"message"`
}
type UpdateEmployeeResponse struct {
	Data    *CreateEmployee `json:"data"`
	Status  string          `json:"status"`
	Message string          `json:"message"`
}
type ErrorResponse struct {
	ErrorCode string `json:"error_code"`
	ErrorMsg  string `json:"error_msg"`
}

func (b *CreateEmployee) TableName() string {

	return "Employee"
}
func (b *Employee) TableName() string {

	return "Employee"
}
