package app

import (
	"AuthinGo/config"
	db "AuthinGo/db/reposetries"
	"AuthinGo/router"
	"net/http"
	"time"
)

//  Defining the Type of application

type Config struct{
	Addr string
	 
}

func NewConfig() Config{ 
	port := config.GetString("PORT" , ":8080")

	return  Config{
		Addr: port,
	}
}

type Application struct{
	Config Config
	Storage db.Storage
}

func NewApp() *Application{
	return &Application{
		Config: NewConfig(),
		Storage: *db.NewStorage(),
	}
}

func (app *Application) Run() error{

	server := &http.Server{
		Addr: app.Config.Addr,
		Handler:  router.SetupRouter()  , // setup the chi router here
		ReadTimeout: 10*time.Second,
		WriteTimeout: 10*time.Second,

	}

	return  server.ListenAndServe() // makes our app instance Listen to the incoming request
}



//  