package main

import (
	"net/http"
	"os"

	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	"github.com/yudai/pp"
	"go.elastic.co/apm"
	"go.elastic.co/apm/module/apmechov4"
)

func init() {
	godotenv.Load()
	pp.Println(os.Getenv("ELASTIC_APM_SERVICE_NAME"))
}

func main() {
	e := echo.New()

	tracer := apm.DefaultTracer
	// tracer.Service.Name = "service-1"
	// tracer.Service.Version = "1.0"
	// tracer.Service.Environment = "dev"
	e.Use(apmechov4.Middleware(apmechov4.WithTracer(tracer)))

	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "[GET] Hello, World!")
	})

	e.POST("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "[POST] Hello, World!")
	})

	e.Logger.Fatal(e.Start(":8000"))
}
