package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Book struct {
	Id       primitive.ObjectID `json:"id,omitempty" bson:"_id"`
	Author   string             `json:"author,omitempty" validate:"required"`
	Category BookCategory       `json:"category,omitempty" validate:"required"`
	Title    string             `json:"title,omitempty" validate:"required"`
}

type BookCategory int

const (
	Action BookCategory = iota
	Romance
	Mistery
	SciFi
	Contemporary
)
