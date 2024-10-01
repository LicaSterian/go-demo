package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"
)

const createNewUserAddr = "http://localhost:8080/users/new"

var client http.Client

type User struct {
	Name  string `json:"name"`
	Email string `json:"email"`
}

type NewUserResponse struct {
	Success bool        `json:"success"`
	Error   string      `json:"error"`
	Data    interface{} `json:"data"`
}

func main() {
	client = http.Client{}

	numConcurrentRequests := 10

	startTime := time.Now().UTC()

	for i := 0; i < numConcurrentRequests; i++ {
		userEmail := fmt.Sprintf("user_%d@hcltech.com", i)
		createNewUser(userEmail)
	}

	took := time.Since(startTime)
	log.Printf("%d sequential requests took: %v", numConcurrentRequests, took)
}

func createNewUser(userEmail string) {
	newUser := User{
		Name:  "John Doe",
		Email: userEmail,
	}
	reqBodyBytes, err := json.Marshal(newUser)
	if err != nil {
		log.Println("json.Marshal error", err.Error())
		return
	}
	req, err := http.NewRequest(http.MethodPost, createNewUserAddr, bytes.NewBuffer(reqBodyBytes))
	if err != nil {
		log.Println("http.NewRequest error", err.Error())
		return
	}

	resp, err := client.Do(req)
	if err != nil {
		log.Println("client.Do error", err.Error())
		return
	}
	defer resp.Body.Close()

	newUserResp := NewUserResponse{}
	decoder := json.NewDecoder(resp.Body)
	err = decoder.Decode(&newUserResp)
	if err != nil {
		log.Println("decoder.Decode error", err.Error())
		return
	}

	if !newUserResp.Success {
		log.Printf("error creating new user, error: %s\n", newUserResp.Error)
		return
	}

	log.Printf("user successfully created, userID: %v\n", newUserResp.Data)
}
