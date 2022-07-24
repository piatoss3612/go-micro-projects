package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Entry struct {
	ID          primitive.ObjectID `bson:"id" json:"-"`
	Dish        *string            `json:"dish"`
	Ingredients *string            `json:"ingredients"`
	Calories    *float64           `json:"calories"`
	Fat         *float64           `json:"fat"`
}
