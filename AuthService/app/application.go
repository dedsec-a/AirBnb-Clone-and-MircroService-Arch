package app

import (
	"AuthinGo/config"
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
}

func (app *Application) Run() error{
	server := &http.Server{
		Addr: app.Config.Addr,
		Handler:  nil  , // setup the chi router here
		ReadTimeout: 10*time.Second,
		WriteTimeout: 10*time.Second,

	}

	return  server.ListenAndServe() // makes our app instance Listen to the incoming request
}



//  