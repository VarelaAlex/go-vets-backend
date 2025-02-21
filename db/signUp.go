package db

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"project/go-vets-backend/models"
	"project/go-vets-backend/utils"
)

func SignUp(user models.User) (string, error) {
	col := DB.Collection("users")
	user.Password, _ = utils.EncryptPassword(user.Password)
	result, err := col.InsertOne(Ctx, user)
	if err != nil {
		return "", err
	}
	objID, _ := result.InsertedID.(primitive.ObjectID)
	return objID.String(), nil
}
