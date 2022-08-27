package model

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type RefreshToken struct {
	ID        primitive.ObjectID `bson:"_id,omitempty"`
	AuthUUID  string             `bson:"auth_uuid"`
	UserID    primitive.ObjectID `bson:"user_id"`
	CreatedAt time.Time          `bson:"createdAt"`
	ExpireAt  time.Time          `bson:"expireAt"`
}
