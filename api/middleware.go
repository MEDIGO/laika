package api

import (
	"net/http"
	"time"

	"github.com/DataDog/datadog-go/statsd"
	log "github.com/Sirupsen/logrus"
	"github.com/labstack/echo"
)

func LogRequestMiddleware() echo.MiddlewareFunc {
	return func(next echo.Handler) echo.Handler {
		return echo.HandlerFunc(func(c echo.Context) error {
			start := time.Now()
			err := next.Handle(c)
			duration := time.Since(start)

			log.WithFields(log.Fields{
				"request":      c.Request().URI(),
				"method":       c.Request().Method(),
				"remote":       c.Request().RemoteAddress(),
				"status":       c.Response().Status(),
				"request_time": duration,
			}).Info("request handled")

			return err
		})
	}
}

func InstrumentMiddleware(stats *statsd.Client) echo.MiddlewareFunc {
	return func(next echo.Handler) echo.Handler {
		return echo.HandlerFunc(func(c echo.Context) error {
			start := time.Now()
			err := next.Handle(c)
			duration := time.Since(start)

			stats.Count("laika.request_total", 1, nil, 1)
			stats.Histogram("laika.request_duration_milliseconds", duration.Seconds() * 1000, nil, 1)

			return err
		})
	}
}

func ResponseEncoderMiddleware() echo.MiddlewareFunc {
	return func(next echo.Handler) echo.Handler {
		return echo.HandlerFunc(func(c echo.Context) error {
			err := next.Handle(c)

			switch v := err.(type) {
			case Response:
				if v.Status >= 500 {
					log.Error(v.Error())
				}
				return c.JSON(v.Status, v.Payload)
			case *echo.HTTPError:
				return c.JSON(v.Code, APIError{v.Error()})
			default:
				if err != nil {
					log.Error(err)
					return c.JSON(http.StatusInternalServerError, APIError{err.Error()})
				}
				return nil
			}
		})
	}
}
