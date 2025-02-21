package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

type User struct {
	ID        primitive.ObjectID `bson:"_id,omitempty" json:"id,omitempty"`
	Name      string             `bson:"name" json:"name"`
	Surname   string             `bson:"surname" json:"surname"`
	BirthDate time.Time          `bson:"birthDate" json:"birthDate"`
	Email     string             `bson:"email" json:"email"`
	Password  string             `bson:"password" json:"password"`
}
