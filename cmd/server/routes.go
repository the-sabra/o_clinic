package server

import (
	"net/http"
	"o_clinic/internal/db"
	"o_clinic/pkg/initializer/doctor"

	"github.com/labstack/echo/v4"
)
 
// RegisterRoutes registers all the routes for the API.
func RegisterRoutes(e *echo.Echo, db  *db.Storage) {
	e.GET("/", func(c echo.Context) error { 
		return c.String(http.StatusOK, "Hello, World!") 
	}) 
 
	// Add more routes here
	api := e.Group("/api")
	
	
	doctor.Initialize(api, db)
} 

