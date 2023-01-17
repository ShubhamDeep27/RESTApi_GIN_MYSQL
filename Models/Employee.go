package models

type Employee struct {
	ID       uint   `json:"id"`
	Name     string `json:"name"`
	Mobile   string `json:"mobile"`
	Address  string `json:"address"`
	Salary   int    `json:"salary"`
	Age      int    `json:"age"`
	Username string `json:"username"`
	Password string `json:"password"`
}

func (b *Employee) TableName() string {
	return "Employee"
}
