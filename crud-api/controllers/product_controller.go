package controllers

import (
	"crud-api/models"
	"crud-api/repository"
	"crud-api/utils"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type ProductController struct {
	Repo *repository.ProductRepository
}

func NewProductController(repo *repository.ProductRepository) *ProductController {
	return &ProductController{Repo: repo}
}

// Create Product
func (ctrl *ProductController) CreateProduct(c *gin.Context) {
	var product models.Product
	if err := c.ShouldBindJSON(&product); err != nil {
		utils.RespondError(c, http.StatusBadRequest, "Invalid request payload")
		return
	}

	if err := ctrl.Repo.CreateProduct(&product); err != nil {
		utils.RespondError(c, http.StatusInternalServerError, "Failed to create product")
		return
	}

	utils.RespondSuccess(c, http.StatusCreated, product)
}

// Get Single Product
func (ctrl *ProductController) GetProduct(c *gin.Context) {
	id := c.Param("id")
	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		utils.RespondError(c, http.StatusBadRequest, "Invalid ID")
		return
	}

	product, err := ctrl.Repo.GetProductByID(objID)
	if err != nil {
		utils.RespondError(c, http.StatusNotFound, "Product not found")
		return
	}

	utils.RespondSuccess(c, http.StatusOK, product)
}

// Get All Products
func (ctrl *ProductController) GetProducts(c *gin.Context) {
	limit, _ := strconv.Atoi(c.Query("limit"))
	offset, _ := strconv.Atoi(c.Query("offset"))

	limit64 := int64(limit)
	offset64 := int64(offset)

	products, err := ctrl.Repo.GetProducts(limit64, offset64)
	if err != nil {
		utils.RespondError(c, http.StatusInternalServerError, "Failed to fetch products")
		return
	}

	utils.RespondSuccess(c, http.StatusOK, products)
}

// Get Products by category or type
func (ctrl *ProductController) GetProductsByCategoryOrType(c *gin.Context) {
	category := c.Query("category")
	productType := c.Query("type")

	products, err := ctrl.Repo.GetProductsByCategoryOrType(category, productType)
	if err != nil {
		utils.RespondError(c, http.StatusInternalServerError, "Failed to fetch products")
		return
	}

	utils.RespondSuccess(c, http.StatusOK, products)
}

// Update the existing product
func (ctrl *ProductController) UpdateProduct(c *gin.Context) {
	id := c.Param("id")
	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		utils.RespondError(c, http.StatusBadRequest, "Invalid ID")
		return
	}

	var update bson.M
	if err := c.ShouldBindJSON(&update); err != nil {
		utils.RespondError(c, http.StatusBadRequest, "Invalid request payload")
		return
	}

	if err := ctrl.Repo.UpdateProduct(objID, update); err != nil {
		utils.RespondError(c, http.StatusInternalServerError, "Failed to update product")
		return
	}

	utils.RespondSuccess(c, http.StatusOK, gin.H{"message": "Product updated successfully"})
}

// Delete the product
func (ctrl *ProductController) DeleteProduct(c *gin.Context) {
	id := c.Param("id")
	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		utils.RespondError(c, http.StatusBadRequest, "Invalid ID")
		return
	}

	if err := ctrl.Repo.DeleteProduct(objID); err != nil {
		utils.RespondError(c, http.StatusInternalServerError, "Failed to delete product")
		return
	}

	utils.RespondSuccess(c, http.StatusOK, gin.H{"message": "Product deleted successfully"})
}
