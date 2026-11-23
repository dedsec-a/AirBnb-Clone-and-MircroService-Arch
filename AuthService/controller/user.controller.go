package controller

import (
	"AuthinGo/service"
	"net/http"
)

type UserController struct {
	// it should have a userService 
	UsereService service.UserService
}


// constructor to create a new userService
func NewUserContoller(_userService service.UserService) *UserController{
	return &UserController{
		UsereService: _userService,
	}

}

func (u *UserController) RegisterUser(w http.ResponseWriter, r *http.Request) {
    if err := u.UsereService.CreateUser(); err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }

    w.WriteHeader(http.StatusCreated)
    w.Write([]byte("User Registration Endpoint"))
}