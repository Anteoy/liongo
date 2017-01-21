package modle

type User struct {
	Id       int    `json:"id"`
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

func NewUser(id int, name string, password string, email string) *User {
	return &User{id, name, password, email}
}
