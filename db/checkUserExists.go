package db

import (
	"go.mongodb.org/mongo-driver/bson"
	"project/go-vets-backend/models"
)

func CheckUserExists(email string) (models.User, bool) {
	col := DB.Collection("users")
	filter := bson.M{"email": email}
	var result models.User
	err := col.FindOne(Ctx, filter).Decode(&result)
	if err != nil {
		return result, false
	}
	return result, true
}
