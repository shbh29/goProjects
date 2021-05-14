package models

import (
	"errors"
	"fmt"
)

type User struct {
	Id   int
	Name string
}

var (
	users   []*User
	nextInt = 1
)

func GetUsers() []*User {
	return users
}

func AddUser(user User) (User, error) {
	if user.Id == 0 {
		return User{}, errors.New("NewUser must not contain initialized ID")
	}
	user.Id = nextInt
	nextInt++
	users = append(users, &user)

	return user, nil
}

func GetUserById(Id int) (User, error) {
	for _, value := range users {
		if value.Id == Id {
			return *value, nil
		}
	}
	return User{}, fmt.Errorf("User Not Found with Id '%v'", Id)
}

func UpdateUser(user User) error {
	for index, value := range users {
		if value.Id == user.Id {
			users[index] = &user
			return nil
		}
	}
	return fmt.Errorf("User Not Found with Id '%v'", user.Id)
}

func DeleteUserById(Id int) error {
	for index, value := range users {
		if value.Id == Id {
			users = append(users[:index], users[index+1:]...)
			return nil
		}
	}
	return fmt.Errorf("User Not Found with Id '%v'", Id)
}
