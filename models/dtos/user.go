package dtos

type AddUser struct {
	FirstName string
	LastName  string
	Email     string
	Avatar    string
	Password  string
}

type LoginUser struct {
	Email    string
	Password string
}
type UpdateUser struct {
	FirstName string
	LastName  string
	Email     string
}
