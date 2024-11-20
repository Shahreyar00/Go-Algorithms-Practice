package utils

import (
	"crud-api/models"
)

func IsPointInPolygon(latitude, longitude float64, polygon []models.LatLong) bool {
	inside := false
	n := len(polygon)
	for i := 0; i < n; i++ {
		j := (i + 1) % n
		if (polygon[i].Lon > longitude) != (polygon[j].Lon > longitude) &&
			(latitude < (polygon[j].Lat-polygon[i].Lat)*(longitude-polygon[i].Lon)/(polygon[j].Lon-polygon[i].Lon)+polygon[i].Lat) {
			inside = !inside
		}
	}
	return inside
}
