package user

import "fmt"

// public
type User struct {
	// public
	Name string

	// private (because of lower case of first letter)
	age int
}

// public
func (u *User) ChangeAge(newAge int) {
	u.age = newAge
}

func (u *User) CallPrivateMethodChangeName(newName string) {
	u.changeName(newName)
}

// private (because of lower case of first letter)
func (u *User) changeName(newName string) {
	h := human{name: newName, age: 11, test: "test"}
	fmt.Printf("%+v", h)
	u.Name = newName
}
