package main

import (
	"AuthinGo/app"
	"fmt"
)

func main() {
	fmt.Println("This is the Main Server Running")
	cfg:= app.Config{
		Addr: ":3001",
	}
	app := app.Application{
		Config: cfg,
		
	}
	fmt.Println("The Server is Running on Port :" , cfg.Addr )

	app.Run()
}