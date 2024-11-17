package repository

import (
	"context"
	"crud-api/models"
	"crud-api/utils"
	"fmt"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"golang.org/x/crypto/bcrypt"
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
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	user.Password = string(hashedPassword)

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	return utils.Retry(func() error {
		_, err := r.Collection.InsertOne(ctx, user)
		if err != nil {
			log.Printf("Error inserting user: %v, user: %v", err, user)
		}
		return err
	})
}

func (r *UserRepository) AuthenticateUser(email, password string) (models.User, error) {
	var user models.User

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	err := r.Collection.FindOne(ctx, bson.M{"email": email}).Decode(&user)
	if err != nil {
		return user, fmt.Errorf("user not found: %w", err)
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil {
		return user, fmt.Errorf("invalid password: %w", err)
	}

	return user, nil
}

func (r *UserRepository) GetUserByID(id primitive.ObjectID) (models.User, error) {
	var user models.User
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	err := utils.Retry(func() error {
		return r.Collection.FindOne(ctx, bson.M{"_id": id}).Decode(&user)
	})
	return user, err
}

func (r *UserRepository) UpdateUser(id primitive.ObjectID, user models.User) error {
	if user.Password != "" {
		hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
		if err != nil {
			return fmt.Errorf("failed to hash password: %w", err)
		}
		user.Password = string(hashedPassword)
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	updateFields := bson.M{"$set": user}
	return utils.Retry(func() error {
		_, err := r.Collection.UpdateOne(ctx, bson.M{"_id": id}, updateFields)
		return err
	})
}

func (r *UserRepository) DeleteUser(id primitive.ObjectID) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	return utils.Retry(func() error {
		_, err := r.Collection.DeleteOne(ctx, bson.M{"_id": id})
		return err
	})
}
