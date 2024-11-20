package controllers

import (
	"crud-api/models"
	"crud-api/repository"
	"crud-api/utils"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type LocationController struct {
	Repo *repository.LocationRepository
}

func NewLocationController(repo *repository.LocationRepository) *LocationController {
	return &LocationController{Repo: repo}
}

var Regions = []struct {
	Name   string
	Bounds []models.LatLong
}{
	{
		Name: "North",
		Bounds: []models.LatLong{
			{Lat: 28.6, Lon: 77.0},
			{Lat: 29.0, Lon: 77.2},
			{Lat: 28.8, Lon: 77.3},
			{Lat: 28.7, Lon: 77.1},
		},
	},
}

func (ctrl *LocationController) AddLocation(c *gin.Context) {
	var location models.Location
	if err := c.ShouldBindJSON(&location); err != nil {
		utils.RespondError(c, http.StatusBadRequest, "Invalid input")
		return
	}

	location.Location = models.GeoJSONPoint{
		Type:        "Point",
		Coordinates: []float64{location.Longitude, location.Latitude},
	}

	err := ctrl.Repo.AddLocation(location)
	if err != nil {
		utils.RespondError(c, http.StatusInternalServerError, "Failed to add location")
		return
	}

	utils.RespondSuccess(c, http.StatusCreated, "Location added successfully")
}

func (ctrl *LocationController) GetNearbyLocations(c *gin.Context) {
	latitude, err1 := strconv.ParseFloat(c.Query("latitude"), 64)
	longitude, err2 := strconv.ParseFloat(c.Query("longitude"), 64)
	radius, err3 := strconv.ParseFloat(c.Query("radius"), 64)

	if err1 != nil || err2 != nil || err3 != nil || radius <= 0 {
		utils.RespondError(c, http.StatusBadRequest, "Invalid query parameters")
		return
	}

	locations, err := ctrl.Repo.GetNearbyLocations(latitude, longitude, radius)
	if err != nil {
		log.Printf("Error: %v", err)
		utils.RespondError(c, http.StatusInternalServerError, "Failed to fetch nearby locations")
		return
	}

	utils.RespondSuccess(c, http.StatusOK, locations)
}

func (ctrl *LocationController) DetermineRegion(c *gin.Context) {
	latitude, err1 := strconv.ParseFloat(c.Query("latitude"), 64)
	longitude, err2 := strconv.ParseFloat(c.Query("longitude"), 64)

	if err1 != nil || err2 != nil {
		utils.RespondError(c, http.StatusBadRequest, "Invalid query parameters")
		return
	}

	for _, region := range Regions {
		fmt.Printf("Checking region: %s with bounds: %v\n", region.Name, region.Bounds)
		if utils.IsPointInPolygon(latitude, longitude, region.Bounds) {
			utils.RespondSuccess(c, http.StatusOK, gin.H{"region": region.Name})
			return
		}
	}

	utils.RespondError(c, http.StatusNotFound, "Region not found")
}
