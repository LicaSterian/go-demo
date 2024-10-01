package model

import "github.com/google/uuid"

type User struct {
	ID    uuid.UUID `json:"id" bson:"_id"`
	Name  string    `json:"name" bson:"name" binding:"required"`
	Email string    `json:"email" bson:"email" binding:"required"`
}
