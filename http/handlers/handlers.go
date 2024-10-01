package handlers

import (
	"fmt"
	"go-demo/http/model"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

var users []model.User
var userEmails map[string]struct{}

func init() {
	users = []model.User{}
	userEmails = make(map[string]struct{})
}

// NewUser is a gin handler func that creates a new user from the data that it receives
// it can return a 400 if the JSON body validation has errors
// it returns the newly created user id
func NewUser(c *gin.Context) {
	log.Println("newUser")
	user := model.User{}
	resp := model.Response{}
	err := c.ShouldBindJSON(&user)
	if err != nil {
		resp.Error = fmt.Sprintf("bind JSON error, %s", err.Error())
		log.Println(resp.Error)
		c.JSON(http.StatusBadRequest, resp)
		return
	}

	// attach a newly generated uuid to the user
	user.ID = uuid.New()

	// check that the email is unique
	_, exists := userEmails[user.Email]
	if exists {
		resp.Error = "email not unique"
		log.Println(resp.Error)
		c.JSON(http.StatusBadRequest, resp)
		return
	}

	// save the user
	users = append(users, user)
	userEmails[user.Email] = struct{}{}

	// return the user id
	resp.Success = true
	resp.Data = user.ID
	log.Printf("newUser OK, userID: %s\n", user.ID)

	c.JSON(http.StatusOK, resp)
}

// GetAllUsers is a handler func that returns all the users
func GetAllUsers(c *gin.Context) {
	log.Println("getAllUsers")
	resp := model.Response{
		Success: true,
		Data:    users,
	}
	log.Println("getAllUsers OK")
	c.JSON(http.StatusOK, resp)
}

// GetUserById is a handler func that searches for a user by it's id
// it can return a 400 if userID is not a valid UUID
// it returns the user if it finds it
// or it returns a 404
func GetUserById(c *gin.Context) {
	userID := c.Param("userID")
	log.Printf("getUserById userID: %s\n", userID)
	resp := model.Response{}

	err := uuid.Validate(userID)
	if err != nil {
		resp.Error = fmt.Sprintf("uuid.Validate error, %s", err.Error())
		log.Println(resp.Error)
		c.JSON(http.StatusBadRequest, resp)
		return
	}

	for _, user := range users {
		if user.ID.String() == userID {
			resp.Success = true
			resp.Data = user
			break
		}
	}

	if resp.Data == nil {
		resp.Error = "userID not found in database"
		log.Println(resp.Error)
		c.JSON(http.StatusNotFound, resp)
		return
	}

	log.Println("getUserById OK")
	c.JSON(http.StatusOK, resp)
}
