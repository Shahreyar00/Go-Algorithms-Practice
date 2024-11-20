package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type LatLong struct {
	Lat float64 `json:"lat" bson:"lat"`
	Lon float64 `json:"lon" bson:"lon"`
}

type Location struct {
	ID        primitive.ObjectID `json:"id,omitempty" bson:"_id,omitempty"`
	Latitude  float64            `json:"latitude" bson:"latitude"`
	Longitude float64            `json:"longitude" bson:"longitude"`
	PinCode   string             `json:"pin_code" bson:"pin_code"`
	Region    string             `json:"region,omitempty" bson:"region,omitempty"`
	Location  GeoJSONPoint       `json:"location" bson:"location"`
}

type GeoJSONPoint struct {
	Type        string    `json:"type" bson:"type"`
	Coordinates []float64 `json:"coordinates" bson:"coordinates"`
}
