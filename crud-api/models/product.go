package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Product struct {
	ID          primitive.ObjectID `json:"id,omitempty" bson:"_id,omitempty"`
	Name        string             `json:"name" bson:"name"`
	Description string             `json:"description" bson:"description"`
	Category    string             `json:"category" bson:"category"`
	Type        string             `json:"type" bson:"type"`
	Price       float64            `json:"price" bson:"price"`
	Stock       int                `json:"stock" bson:"stock"`
	ImageURL    string             `json:"image_url" bson:"image_url"`
	CreatedAt   int64              `json:"created_at" bson:"created_at"`
	UpdatedAt   int64              `json:"updated_at" bson:"updated_at"`
}
