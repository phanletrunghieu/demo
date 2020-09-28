package http

import (
	"net/http"
	"time"

	"github.com/labstack/echo-contrib/jaegertracing"
	"github.com/labstack/echo/v4"

	clientOpentracing "service-1/client/opentracing"
	"service-1/client/service2"
)

func New() *echo.Echo {
	e := echo.New()

	e.Use(jaegertracing.TraceWithConfig(jaegertracing.TraceConfig{
		Tracer:     clientOpentracing.GetTracer(),
		Skipper:    nil,
		IsBodyDump: true,
	}))

	e.GET("/", func(c echo.Context) error {
		jaegertracing.TraceFunction(c, slowFunc, "Func 1")

		clientSvc2 := service2.GetClient()
		resp, err := clientSvc2.Hello(c.Request().Context(), &service2.HelloRequest{
			Message: "I'm service 1",
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
