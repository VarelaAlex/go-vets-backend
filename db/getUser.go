package db

import (
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"log"
	"project/go-vets-backend/models"
)

func GetUser(ID string) (models.User, error) {
	col := DB.Collection("users")
	objID, _ := primitive.ObjectIDFromHex(ID)
	filter := bson.M{"_id": objID}
	var user models.User
	err := col.FindOne(Ctx, filter).Decode(&user)
	user.Password = ""
	if err != nil {
		log.Println("User not found: " + err.Error())
		return user, err
	}
	return user, nil
}
