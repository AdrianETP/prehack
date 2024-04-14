package main

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type User struct {
	Id       primitive.ObjectID `bson:"_id, omitempty" json:id`
	UserName string             `json: UserName`
	Password string             `json: Password`
}
