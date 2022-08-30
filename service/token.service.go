package service

import (
	"context"

	"github.com/OneKiwiTech/gofiber/database"
	"github.com/OneKiwiTech/gofiber/model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func CreateRefreshToken(token model.RefreshToken) (model.RefreshToken, error) {
	res, err := database.Collection("refresh_tokens").InsertOne(context.Background(), token)
	if err != nil {
		//panic(err)
		return model.RefreshToken{}, err
	}
	token, _ = GetRefreshTokenByID(res.InsertedID.(primitive.ObjectID))
	return token, nil
}

func GetRefreshTokenByID(id primitive.ObjectID) (model.RefreshToken, error) {
	var token model.RefreshToken
	err := database.Collection("refresh_tokens").FindOne(context.Background(), bson.M{"_id": id}).Decode(&token)
	if err != nil {
		return token, err
	}
	return token, nil
}
