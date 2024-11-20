package repository

import (
	"context"
	"crud-api/models"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type LocationRepository struct {
	Collection *mongo.Collection
}

func NewLocationRepository(db *mongo.Database) *LocationRepository {
	repo := &LocationRepository{
		Collection: db.Collection("locations"),
	}

	if err := repo.CreateGeoIndex(); err != nil {
		panic(err)
	}

	return repo
}

func (r *LocationRepository) AddLocation(location models.Location) error {
	_, err := r.Collection.InsertOne(context.TODO(), location)
	return err
}

func (r *LocationRepository) GetNearbyLocations(latitude, longitude, radius float64) ([]models.Location, error) {
	filter := bson.M{
		"location": bson.M{
			"$nearSphere": bson.M{
				"$geometry": bson.M{
					"type":        "Point",
					"coordinates": []float64{longitude, latitude},
				},
				"$maxDistance": radius,
			},
		},
	}

	cursor, err := r.Collection.Find(context.TODO(), filter)
	if err != nil {
		return nil, err
	}

	var locations []models.Location
	if err := cursor.All(context.TODO(), &locations); err != nil {
		return nil, err
	}

	return locations, nil
}

func (r *LocationRepository) CreateGeoIndex() error {
	indexModel := mongo.IndexModel{
		Keys: bson.M{"location": "2dsphere"},
	}

	_, err := r.Collection.Indexes().CreateOne(context.TODO(), indexModel)
	return err
}
