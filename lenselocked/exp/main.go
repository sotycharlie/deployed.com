package main

import (
	"fmt"

	"deployed.com/lenselocked/models"
	_ "github.com/lib/pq"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "2001"
	dbname   = "new_db"
)

func main() {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)

	// this function would ret,urn the *sql.
	us := models.NewUserService(psqlInfo)
	user := models.User{
		Name:     "Micheal Scott",
		Email:    "Micheal@dumehe.com",
		Password: "bestboss",
	}
	if err := us.Create(&user); err != nil {
		panic(err)
	}
	// Verify that the user has a Remember and RememberHash
	fmt.Printf("%+v\n", user)
	if user.Remember == "" {
		panic("Invalid remember token")
	}
	// Now verify that we can lookup a user with that remember
	// token
	user2, err := us.ByRemember(user.Remember)
	if err != nil {
		panic(err)
	}
	fmt.Printf("%+v\n", *user2)
}

/*
if err := us.Delete(user.ID); err != nil {
		panic(err)
	}
	fmt.Println(err)

	user.Email = "micheal@djcjcdcnijdxinj.com"
	if err := us.Update(&user); err != nil {
		panic(err)
	}
	NewEmail, err := us.ByEmail("micheal@djcjcdcnijdxinj.com")
	if err != nil {
		panic(err)
	}
	fmt.Println(NewEmail)
*/
