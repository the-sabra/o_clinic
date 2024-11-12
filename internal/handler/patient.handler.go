package handler

import (
	"context"
	"database/sql"
	"net/http"
	"o_clinic/internal/domain"

	"strconv"

	"github.com/labstack/echo/v4"
)

// PatientService represents the service for managing patient.
type PatientService interface {
	CreatePatient(ctx context.Context, doctor *domain.NewPatient) (int64, error)
	GetPatient(ctx context.Context, id int64) (*domain.Patient, error)
	GetAllPagination(ctx context.Context, limit, offset int) ([]*domain.Patient, error)
	UpdatePatient(ctx context.Context, doctor *domain.Patient) error
	DeletePatient(ctx context.Context, id int64) error
}

// PatientHandler represents the HTTP handler for patient.
type PatientHandler struct {
	service PatientService
}

// NewPatientHandler creates a new patient handler.
func NewPatientHandler(service PatientService) *PatientHandler {
	return &PatientHandler{service: service}
}


// Create godoc
// @Summary      Create Patient
// @Description  create Patient
// @Tags         Patient
// @Accept       json
// @Produce      json
// @Param        Patient body domain.NewPatient true "Patient"
// @Success      201  {object}   map[string]int64
// @Failure      400  {object}  map[string]string
// @Failure      500  {object}  map[string]string
// @Router       /patient [post]
func (h *PatientHandler) Create(c echo.Context) error {
	var doctor domain.NewPatient
	if err := c.Bind(&doctor); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
	}

	id, err := h.service.CreatePatient(c.Request().Context(), &doctor)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}

	return c.JSON(http.StatusCreated, map[string]int64{"id": id})
}

// GetByID godoc
// @Summary      Get Patient by id
// @Description  get Patient by id
// @Tags         Patient
// @Accept       json
// @Produce      json
// @Param        id path int true "Patient ID"
// @Success      200  {object}   domain.Patient
// @Failure      400  {object}  map[string]string
// @Failure      500  {object}  map[string]string
// @Router       /patient/{id} [get]
func (h *PatientHandler) GetByID(c echo.Context) error {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
	}

	doctor, err := h.service.GetPatient(c.Request().Context(), id)
	if err != nil {
		if err == sql.ErrNoRows {
			return c.JSON(http.StatusNotFound, map[string]string{"error": "patient not found"})
		}
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}

	return c.JSON(http.StatusOK, doctor)
}

// GetAllPagination godoc
// @Summary      Get all Patient with pagination
// @Description  get all Patient with pagination
// @Tags         Patient
// @Accept       json
// @Produce      json
// @Param        page query int true "Page"
// @Param        take query int true "Take"
// @Success      200  {object}   []domain.Patient
// @Failure      400  {object}  map[string]string
// @Failure      500  {object}  map[string]string
// @Router       /patient [get]
func (h *PatientHandler) GetAllPagination(c echo.Context) error {
	page, err := strconv.ParseInt(c.QueryParam("page"), 10, 64)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "invalid page"})
	}
	take, err := strconv.ParseInt(c.QueryParam("take"), 10, 64)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "invalid take"})
	}
	offset := (page - 1) * take
	limit := take
	patients, err := h.service.GetAllPagination(c.Request().Context(), int(limit), int(offset))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}

	return c.JSON(http.StatusOK, patients)
}

// Update godoc
// @Summary      Update Patient
// @Description  update Patient
// @Tags         Patient
// @Accept       json
// @Produce      json
// @Param        id path int true "Patient ID"
// @Param        Patient body domain.Patient true "Patient"
// @Success      200  {object}   map[string]string
// @Failure      400  {object}  map[string]string
// @Failure      500  {object}  map[string]string
// @Router       /patient/{id} [put]
func (h *PatientHandler) Update(c echo.Context) error {
	var doctor domain.Patient
	if err := c.Bind(&doctor); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
	}

	err := h.service.UpdatePatient(c.Request().Context(), &doctor)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}

	return c.JSON(http.StatusOK, map[string]string{"message": "success"})
}

// Delete godoc
// @Summary      Delete Patient
// @Description  delete Patient
// @Tags         Patient
// @Accept       json
// @Produce      json
// @Param        id path int true "Patient ID"
// @Success      200  {object}   map[string]string
// @Failure      400  {object}  map[string]string
// @Failure      500  {object}  map[string]string
// @Router       /patient/{id} [delete]
func (h *PatientHandler) Delete(c echo.Context) error {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
	}

	err = h.service.DeletePatient(c.Request().Context(), id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}

	return c.JSON(http.StatusOK, map[string]string{"message": "success"})
}