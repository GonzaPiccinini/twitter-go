package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type User struct {
	ID primitive.ObjectID `bson:"_id,omitempty" json:"id"`

	Avatar    string    `bson:"avatar" json:"avatar,omitempty"`
	Banner    string    `bson:"banner" json:"banner,omitempty"`
	Biography string    `bson:"biography" json:"biography,omitempty"`
	Birthdate time.Time `bson:"birthdate" json:"birthdate,omitempty"`
	Email     string    `bson:"email" json:"email"`
	Firstname string    `bson:"firstname" json:"firstname,omitempty"`
	Lastname  string    `bson:"lastname" json:"lastname,omitempty"`
	Password  string    `bson:"password" json:"password,omitempty"`
	Ubication string    `bson:"ubication" json:"ubication,omitempty"`
	Web       string    `bson:"web" json:"web,omitempty"`
}
