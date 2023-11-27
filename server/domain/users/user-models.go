package users

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// For now this is the input/output DTO, domain object, and mongo db object
type User struct {
	ID        primitive.ObjectID `bson:"_id"`
	FirstName string             `json:"firstName" validate:"required,min=1"`
	LastName  string             `json:"lastName" validate:"required,min=1"`
	Email     string             `json:"email" validate:"required,email"`
}
