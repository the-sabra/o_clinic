package doctor

import (
	"o_clinic/internal/db"
	"o_clinic/internal/handler"
	"o_clinic/internal/repository"
	"o_clinic/internal/service"

	"github.com/labstack/echo/v4"
)

// Initialize sets up the doctor routes, handlers, services, and repositories.
func Initialize(e *echo.Group, db *db.Storage) {
    // Initialize repositories
    doctorRepo := repository.NewDoctorRepository(db)

    // Initialize services
    doctorService := service.NewDoctorService(doctorRepo)

    // Initialize handlers
    doctorHandler := handler.NewDoctorHandler(doctorService)

	// Routes
	doctorGroup := e.Group("/doctor")
	doctorGroup.POST("", doctorHandler.Create)
	doctorGroup.GET("/:id", doctorHandler.GetByID)
	doctorGroup.GET("", doctorHandler.GetAllPagination)
	doctorGroup.PUT("/:id", doctorHandler.UpdateDoctor)
	doctorGroup.DELETE("/:id", doctorHandler.DeleteDoctor)
}   
 