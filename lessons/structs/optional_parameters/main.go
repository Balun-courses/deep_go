package main

type Option func(*User)

type User struct {
	Name    string
	Surname string
	Email   *string
	Phone   *string
	Address *string
}

func NewUser(name string, surname string, email, phone, address *string) User {
	return User{
		Name:    name,
		Surname: surname,
		Email:   email,
		Phone:   phone,
		Address: address,
	}
}

func main() {
	email := "test@test.ru"
	phone := ""

	user1 := NewUser("Ivan", "Ivanov", &email, &phone, nil)
	_ = user1
}
