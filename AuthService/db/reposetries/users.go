package db

import (
	"database/sql"
	"fmt"
)

type UserRepository interface { // this decouples the whole arch
	Create() error
}

type UserRepositoryImp struct {
	db *sql.DB // used to make Db Queries

}

func (u *UserRepositoryImp) Create() error {
	fmt.Println("The User is being created in the DB   ")
	return nil

}