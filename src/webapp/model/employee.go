package model

type Employee struct {
	Id    int64  `json:"id"`
	Name  string `json:"employee_name"`
	Email string `json:"email"`
}
