package handler

import (
	"context"
	"fmt"
	"net/http"
	"o_clinic/internal/domain"

	"strconv"

	"github.com/labstack/echo/v4"
)

// DoctorService represents the service for managing doctors.
type DoctorService interface {
	CreateDoctor(ctx context.Context, doctor *domain.NewDoctor) (int64, error)
	GetDoctor(ctx context.Context, id int64) (*domain.Doctor, error)
	GetAllPagination(ctx context.Context, limit, offset int) ([]*domain.Doctor, error)
	UpdateDoctor(ctx context.Context, doctor *domain.Doctor) error
	DeleteDoctor(ctx context.Context, id int64) error
}

// DoctorHandler represents the HTTP handler for doctor.
type DoctorHandler struct {
	service DoctorService
}

// NewDoctorHandler creates a new doctor handler.
func NewDoctorHandler(service DoctorService) *DoctorHandler {
	return &DoctorHandler{service: service}
}

// Create godoc
// @Summary      Create doctor
// @Description  create doctor
// @Tags         doctor
// @Accept       json
// @Produce      json
// @Param        doctor body domain.NewDoctor true "Doctor"
// @Success      201  {object}   map[string]int64
// @Failure      400  {object}  map[string]string
// @Failure      500  {object}  map[string]string
// @Router       /doctor [post]
func (h *DoctorHandler) Create(c echo.Context) error {
	var doctor domain.NewDoctor
	if err := c.Bind(&doctor); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
	}

	id, err := h.service.CreateDoctor(c.Request().Context(), &doctor)
	fmt.Println("id", id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}

	return c.JSON(http.StatusCreated, map[string]int64{"id": id})
}

// GetByID godoc
// @Summary      Get doctor by id
// @Description  get doctor by id
// @Tags         doctor
// @Accept       json
// @Produce      json
// @Param        id path int true "Doctor ID"
// @Success      200  {object}   domain.Doctor
// @Failure      400  {object}  map[string]string
// @Failure      404  {object}  map[string]string
// @Failure      500  {object}  map[string]string
// @Router       /doctor/{id} [get]
func (h *DoctorHandler) GetByID(c echo.Context) error {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "invalid id"})
	}

	doctor, err := h.service.GetDoctor(c.Request().Context(), id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}

	if doctor == nil {
		return c.JSON(http.StatusNotFound, map[string]string{"error": "doctor not found"})
	}

	return c.JSON(http.StatusOK, doctor)
}

// GetAllPagination godoc
// @Summary      List doctors
// @Description  get doctors
// @Tags         doctor
// @Accept       json
// @Produce      json
// @Success      200  {array}   domain.Doctor
// @Failure      400  {object}  map[string]string
// @Failure      500  {object}  map[string]string
// @Param        page query int true "Page"
// @Param        take query int true "Take"
// @Router       /doctor [get]
func (h *DoctorHandler) GetAllPagination(c echo.Context) error {
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
	doctors, err := h.service.GetAllPagination(c.Request().Context(), int(limit), int(offset))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}

	return c.JSON(http.StatusOK, doctors)
}

// UpdateDoctor godoc
// @Summary      Update doctor
// @Description  update doctor
// @Tags         doctor
// @Accept       json
// @Produce      json
// @Param        id path int true "Doctor ID"
// @Param        doctor body domain.Doctor true "Doctor"
// @Success      200  {object}   map[string]string
// @Failure      400  {object}  map[string]string
// @Failure      500  {object}  map[string]string
// @Router       /doctor/{id} [put]
func (h *DoctorHandler) UpdateDoctor(c echo.Context) error {
	fmt.Println("ID", c.Param("id"))
	doctorID, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "invalid id"})
	}
	var doctor domain.Doctor
	if err := c.Bind(&doctor); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
	}
	doctor.ID = doctorID
	err = h.service.UpdateDoctor(c.Request().Context(), &doctor)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}
	return c.JSON(http.StatusOK, map[string]string{"message": "doctor updated successfully"})
}

// DeleteDoctor  godoc
// @Summary      Delete doctor
// @Description  delete doctor
// @Tags         doctor
// @Accept       json
// @Produce      json
// @Param        id path int true "Doctor ID"
// @Success      200  {object}   map[string]string
// @Failure      400  {object}  map[string]string
// @Failure      500  {object}  map[string]string
// @Router       /doctor/{id} [delete]
func (h *DoctorHandler) DeleteDoctor(c echo.Context) error {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "invalid id"})
	}

	err = h.service.DeleteDoctor(c.Request().Context(), id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}

	return c.JSON(http.StatusOK, map[string]string{"message": "doctor deleted successfully"})
}
