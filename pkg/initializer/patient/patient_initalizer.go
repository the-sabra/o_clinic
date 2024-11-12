package patient


import (
	"o_clinic/internal/db"
	"o_clinic/internal/handler"
	"o_clinic/internal/repository"
	"o_clinic/internal/service"

	"github.com/labstack/echo/v4"
)

// Initialize sets up the Patient routes, handlers, services, and repositories.
func Initialize(e *echo.Group, db *db.Storage) {
    // Initialize repositories
    patientRepo := repository.NewPatientRepository(db)

    // Initialize services
    patientService := service.NewPatientService(patientRepo)

    // Initialize handlers
    patientHandler := handler.NewPatientHandler(patientService)

    // Routes
    patientGroup := e.Group("/patient")
    patientGroup.POST("", patientHandler.Create)
    patientGroup.GET("/:id", patientHandler.GetByID)
    patientGroup.GET("", patientHandler.GetAllPagination)
    patientGroup.PUT("/:id", patientHandler.Update)
    patientGroup.DELETE("/:id", patientHandler.Delete)
}