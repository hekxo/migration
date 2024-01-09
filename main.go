package main

import (
	"fmt"
	_ "github.com/lib/pq"
	"log"
	"migration/database"
	"migration/models"
)

func main() {
	db, err := database.ConnectDB()
	if err != nil {
		log.Fatal(err)
	}
	defer database.CloseDB()

	if db == nil {
		log.Fatal("Database connection is nil")
	}

	user := &models.User{Name: "John Doe", Email: "john@example.com"}

	// Create
	err = database.CreateUser(user)
	if err != nil {
		log.Fatal(err)
	}

	// Read
	retrievedUser, err := database.GetUserByID(user.ID)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Retrieved User: %+v\n", *retrievedUser)

	// Update
	err = database.UpdateUserNameByID(user.ID, "Updated Name")
	if err != nil {
		log.Fatal(err)
	}

	// Read updated user
	retrievedUser, err = database.GetUserByID(user.ID)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Updated User: %+v\n", *retrievedUser)

	// Delete
	err = database.DeleteUserByID(user.ID)
	if err != nil {
		log.Fatal(err)
	}

	// Read all users
	allUsers, err := database.GetAllUsers()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("All Users:", allUsers)
}
