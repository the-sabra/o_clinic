package server

import (
	"bytes"

	"o_clinic/internal/db"
	"o_clinic/pkg/utils"
	"time"

	"github.com/go-playground/validator"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"

	echoSwagger "github.com/swaggo/echo-swagger"

	_ "o_clinic/internal/handler/docs"
)

// Server represents the server instance.
type Server struct{}
  
// NewServer creates a new instance of Server.
func NewServer() *Server {
	return &Server{}
}

// ListenAndServe starts the server and listens on the configured port.

func (s *Server) ListenAndServe() error {
	config := utils.NewConfig()
	e := echo.New()
	 
	e.GET("/swagger/*", echoSwagger.WrapHandler)

	e.Validator = &utils.Validator{Instance: validator.New()}

	setupMiddlewares(e)

	// e.GET("/", func(c echo.Context) error {
	// 	return c.String(200, "Hello, World!")
	// })
	
	db, err := db.NewStorage(config.DBConnString) 
	if err != nil {
		return err
	}
	 
	RegisterRoutes(e,db) // Register routes
    
	
	return e.Start(":" + config.Port)
}

// setupMiddlewares sets up middlewares for logging, timming and recovering.
func setupMiddlewares(e *echo.Echo) {
	e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format:           `${time_rfc3339} | ${method} ${uri} | ${status} | ${custom} | ${remote_ip} | ${user_agent}` + "\n",
		CustomTimeFormat: time.RFC3339,
		CustomTagFunc: func(c echo.Context, buf *bytes.Buffer) (int, error) {
			start := c.Get("start_time").(time.Time)
			end := time.Now()
			latencyStr := utils.CustomLatency(start, end)
			return buf.WriteString(latencyStr)
		},
		Output: nil,
	}))

	e.Use(func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			c.Set("start_time", time.Now())
			return next(c)
		}
	})

	e.Use(middleware.Recover())
}
