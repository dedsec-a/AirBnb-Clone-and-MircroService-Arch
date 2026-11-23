package router

import (
	"AuthinGo/controller"

	"github.com/go-chi/chi"
)

type UserRouter struct {
	//  might accept a User Controller as a fucntion..?
	UserController *controller.UserController 
}

func NewUserRouter(_userController *controller.UserController) *UserRouter {  
	return &UserRouter{
		UserController: _userController,
	}

}

func (ur *UserRouter) Register (r chi.Router){
	r.Post("/signup" , ur.UserController.RegisterUser) // this is where the route is being registerd

}