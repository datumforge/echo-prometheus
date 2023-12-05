package main

import (
	"net/http"

	echoPrometheus "github.com/datumforge/echo-prometheus/v5"
	echo "github.com/datumforge/echox"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

func main() {
	e := echo.New()

	e.Use(echoPrometheus.MetricsMiddleware())
	e.GET("/metrics", echo.WrapHandler(promhttp.Handler()))

	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})
	e.Logger.Error(e.Start(":1323"))
}
