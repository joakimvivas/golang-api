package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type User struct {
	Id          primitive.ObjectID `json:"id,omitempty"`
	Project     string             `json:"project,omitempty" validate:"required"`
	URL         string             `json:"url,omitempty" validate:"required"`
	Description string             `json:"description,omitempty" validate:"required"`
}
