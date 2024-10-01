package main

import (
	"fmt"
	"time"
)

func main() {
	// User used for storing id and name
	type User struct {
		ID   int
		Name string
	}

	numUsers := 100000000
	fmt.Println("numUsers", numUsers)
	users := []User{}
	for i := 0; i < numUsers; i++ {
		u := User{
			ID:   i,
			Name: fmt.Sprintf("user_%d", i),
		}
		users = append(users, u)
	}
	userID := numUsers - 1 // worst T()
	begin := time.Now()
	// linear search
	found := false
	var foundUser User
	for _, user := range users {
		if user.ID == userID {
			found = true
			foundUser = user
			break
		}
	}
	if found {
		fmt.Printf("linear found user:%+v in:%s\n", foundUser, time.Since(begin))
	} else {
		fmt.Println("user not found")
	}
}
