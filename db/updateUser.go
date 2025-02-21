package db

import (
	"go.mongodb.org/mongo-driver/bson"
	"project/go-vets-backend/models"
)

func UpdateUser(user models.User) error {
	col := DB.Collection("users")
	update := bson.D{{Key: "$set",
		Value: bson.D{
			{Key: "name", Value: user.Name},
			{Key: "surname", Value: user.Surname},
			{Key: "birthDate", Value: user.BirthDate},
		},
	}}
	filter := bson.M{"email": user.Email}
	_, err := col.UpdateOne(Ctx, filter, update)
	return err
}
