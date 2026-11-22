package config

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

func load(){
	err := godotenv.Load()
	if err !=  nil{
		//  The Error is logged if the file is not found
		fmt.Println("Error Loading the dotenv file")
	}
}

//  Helper funtions 
// GetString func --> get the Right env form the given  env file
func GetString(key string, fallback string) string{
	//  it loads the file from the load function
	load()
	value , ok :=os.LookupEnv(key)
	if !ok{
		return fallback
	}

	return  value
}