package service

import (
	db "AuthinGo/db/reposetries"
	"fmt"
)

type UserService interface {
	CreateUser() error
}

type UserServiceImp struct {
	userRepo  db.UserRepository // this is depending on the interfacr rather than the concrete class 
}

func NewUSerService(_userReop db.UserRepository) UserService {
	return &UserServiceImp{
		userRepo: _userReop,
	}
}

func (u *UserServiceImp) CreateUser() error {
	fmt.Println("Creating the User in the User Service")
	 u.userRepo.Create() // this is calling the Db layer now
	return nil
}