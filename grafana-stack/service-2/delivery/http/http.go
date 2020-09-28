package http

import (
	"net/http"
	"time"

	"github.com/labstack/echo-contrib/jaegertracing"
	"github.com/labstack/echo/v4"

	clientOpentracing "service-2/client/opentracing"
	"service-2/client/service1"
)

func New() *echo.Echo {
	e := echo.New()

	e.Use(jaegertracing.TraceWithConfig(jaegertracing.TraceConfig{
		Tracer:     clientOpentracing.GetTracer(),
		Skipper:    nil,
		IsBodyDump: true,
	}))

	e.GET("/", func(c echo.Context) error {
		jaegertracing.TraceFunction(c, slowFunc, "Test String")

		clientSvc1 := service1.GetClient()
		resp, err := clientSvc1.Hello(c.Request().Context(), &service1.HelloRequest{
			Message: "I'm service 2",
		})

		if err != nil {
			return err
		}

		return c.JSON(200, resp.GetMessage())
	})

	e.POST("/", func(c echo.Context) error {
		jaegertracing.TraceFunction(c, slowFunc, "Test String")

		return c.String(http.StatusOK, "[POST] Hello, World!")
	})

	return e
}

func slowFunc(s string) {
	time.Sleep(200 * time.Millisecond)
	return
}
