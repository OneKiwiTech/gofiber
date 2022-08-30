package service

import (
	"context"

	"github.com/OneKiwiTech/gofiber/database"
	"github.com/OneKiwiTech/gofiber/model"
	validation "github.com/go-ozzo/ozzo-validation"
	"github.com/go-ozzo/ozzo-validation/is"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

const USERS = "users"

func Validate(user model.User) error {
	return validation.ValidateStruct(&user,
		validation.Field(&user.Name, validation.Required, validation.Length(3, 64)),
		validation.Field(&user.Email, validation.Required, is.Email),
		validation.Field(&user.Password, validation.Required, validation.Length(8, 18)),
	)
}

func CreateUser(user model.User) (model.User, error) {
	res, err := database.Collection(USERS).InsertOne(context.Background(), user)
	if err != nil {
		//panic(err)
		return model.User{}, err
	}
	user, _ = GetUserByID(res.InsertedID.(primitive.ObjectID))
	return user, nil
}

func GetUserByEmail(email string) (model.User, error) {
	var user model.User
	err := database.Collection(USERS).FindOne(context.Background(), bson.M{"email": email}).Decode(&user)
	if err != nil {
		return user, err
	}
	return user, nil
}

func GetUserByID(id primitive.ObjectID) (model.User, error) {
	var user model.User
	err := database.Collection(USERS).FindOne(context.Background(), bson.M{"_id": id}).Decode(&user)
	if err != nil {
		return user, err
	}
	return user, nil
}
