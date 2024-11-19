package repository

import (
	"context"
	"crud-api/models"
	"crud-api/utils"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type ProductRepository struct {
	Collection *mongo.Collection
}

func NewProductRepository(db *mongo.Database) *ProductRepository {
	return &ProductRepository{
		Collection: db.Collection("products"),
	}
}

// Create a new product
func (r *ProductRepository) CreateProduct(product *models.Product) error {
	product.ID = primitive.NewObjectID() // Generate a new ObjectID
	product.CreatedAt = time.Now().Unix()
	product.UpdatedAt = product.CreatedAt

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	result, err := r.Collection.InsertOne(ctx, product)
	if err != nil {
		return err
	}

	// Optionally: Use the inserted ID from the result
	if oid, ok := result.InsertedID.(primitive.ObjectID); ok {
		product.ID = oid
	}
	return nil
}

// Get a single product by ID
func (r *ProductRepository) GetProductByID(id primitive.ObjectID) (models.Product, error) {
	var product models.Product

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	err := utils.Retry(func() error {
		return r.Collection.FindOne(ctx, bson.M{"_id": id}).Decode(&product)
	})
	return product, err
}

// Get a list of products with pagination
func (r *ProductRepository) GetProducts(limit, offset int64) ([]models.Product, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	findOptions := options.Find()
	findOptions.SetLimit(limit)
	findOptions.SetSkip(offset)

	cursor, err := r.Collection.Find(ctx, bson.M{}, findOptions)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	var products []models.Product
	if err := cursor.All(ctx, &products); err != nil {
		return nil, err
	}

	return products, nil
}

// Get products by category or type
func (r *ProductRepository) GetProductsByCategoryOrType(category, productType string) ([]models.Product, error) {
	var products []models.Product
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	filter := bson.M{}
	if category != "" {
		filter["category"] = category
	}
	if productType != "" {
		filter["type"] = productType
	}

	cursor, err := r.Collection.Find(ctx, filter)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	for cursor.Next(ctx) {
		var product models.Product
		if err := cursor.Decode(&product); err != nil {
			return nil, err
		}
		products = append(products, product)
	}
	return products, nil
}

// Update a product by ID
func (r *ProductRepository) UpdateProduct(id primitive.ObjectID, update bson.M) error {
	update["updated_at"] = time.Now().Unix()

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	return utils.Retry(func() error {
		_, err := r.Collection.UpdateOne(ctx, bson.M{"_id": id}, bson.M{"$set": update})
		return err
	})
}

// Delete a product by ID
func (r *ProductRepository) DeleteProduct(id primitive.ObjectID) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	return utils.Retry(func() error {
		_, err := r.Collection.DeleteOne(ctx, bson.M{"_id": id})
		return err
	})
}
