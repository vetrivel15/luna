package user

import (
	"errors"
)

type user struct {
	id        string
	firstName string
	lastName  string
	password  string
}

var (
	users = make(map[string]user)
)

func New(id, firstName, lastName, password string) (bool, error) {

	if IsExist(id) {
		return false, errors.New("User already exists")
	}

	newuser := user{
		id:        id,
		firstName: firstName,
		lastName:  lastName,
		password:  password,
	}

	users[id] = newuser
	return true, nil
}

func IsExist(id string) bool {
	_, ok := users[id]
	return ok
}

func Login(id, password string) (bool, error) {

	allUsers, exists := users[id]

	if !exists {
		return false, errors.New("User not found")
	}

	if allUsers.password != password {
		return false, errors.New("wrong password")
	}

	return true, nil
}

func Delete(id, password string) (bool, error) {
	allUsers, exists := users[id]

	if !exists {
		return false, errors.New("User not found")
	}

	if allUsers.password != password {
		return false, errors.New("wrong password")
	}

	delete(users, id)

	return true, nil
}
