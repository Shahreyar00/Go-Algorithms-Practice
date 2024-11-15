package repository

import (
	"context"
	"crud-api/models"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type UserRepository struct {
	Collection *mongo.Collection
}

func NewUserRepository(db *mongo.Database) *UserRepository {
	return &UserRepository{
		Collection: db.Collection("users"),
	}
}

// CRUD methods
func (r *UserRepository) CreateUser(user models.User) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	_, err := r.Collection.InsertOne(ctx, user)
	return err
}

func (r *UserRepository) GetUserByID(id primitive.ObjectID) (models.User, error) {
	var user models.User
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	err := r.Collection.FindOne(ctx, bson.M{"_id": id}).Decode(&user)
	return user, err
}

func (r *UserRepository) UpdateUser(id primitive.ObjectID, user models.User) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	_, err := r.Collection.UpdateOne(ctx, bson.M{"_id": id}, bson.M{"$set": user})
	return err
}

func (r *UserRepository) DeleteUser(id primitive.ObjectID) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	_, err := r.Collection.DeleteOne(ctx, bson.M{"_id": id})
	return err
}
