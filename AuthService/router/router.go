package router

import (
	"AuthinGo/controller"

	"github.com/go-chi/chi"
)

  func SetupRouter (userRouter *UserRouter) *chi.Mux{
	chirouter := chi.NewRouter()
	chirouter.Get("/ping" , controller.PingHandler)
	userRouter.Register(chirouter)
	return chirouter
  }