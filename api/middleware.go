package api

import (
	"github.com/DataDog/datadog-go/statsd"
	"github.com/MEDIGO/feature-flag/model"
	log "github.com/Sirupsen/logrus"
	"github.com/labstack/echo"
	"net/http"
	"time"
)

func LogRequestMiddleware() echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c *echo.Context) error {
			start := time.Now()
			err := next(c)
			duration := time.Since(start)

			log.WithFields(log.Fields{
				"request":      c.Request().RequestURI,
				"method":       c.Request().Method,
				"remote":       c.Request().RemoteAddr,
				"status":       c.Response().Status(),
				"request_time": duration,
			}).Info("request handled")

			return err
		}
	}
}

func InstrumentMiddleware(stats *statsd.Client) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c *echo.Context) error {
			start := time.Now()
			err := next(c)
			duration := time.Since(start)

			stats.Histogram("core.request_time", duration.Seconds(), nil, 1)

			return err
		}
	}
}

func ResponseEncoderMiddleware() echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c *echo.Context) error {
			err := next(c)

			switch v := err.(type) {
			case Response:
				if v.Status >= 500 {
					log.Error(v.Error())
				}
				return c.JSON(v.Status, v.Payload)
			default:
				log.Error(err)
				return c.JSON(http.StatusInternalServerError, model.APIError{err.Error()})
			}
		}
	}
}
