package api

import (
	"net/http"
	"strconv"
	"time"

	"github.com/DataDog/datadog-go/statsd"
	log "github.com/Sirupsen/logrus"
	"github.com/labstack/echo"
)

// LogMiddleware logs information about the current request.
func LogMiddleware() echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			defer func(start time.Time) {
				log.WithFields(log.Fields{
					"uri":                   c.Request().URI(),
					"method":                c.Request().Method(),
					"remote_addr":           c.Request().RemoteAddress(),
					"status":                c.Response().Status(),
					"duration_microseconds": int(time.Since(start).Seconds() * 1000000),
				}).Info("request handled")
			}(time.Now())

			return next(c)
		}
	}
}

// InstrumentMiddleware collects metrics about the current request.
func InstrumentMiddleware(stats *statsd.Client) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			defer func(start time.Time) {
				tags := []string{
					"method:" + c.Request().Method(),
					"status:" + strconv.Itoa(c.Response().Status()),
				}

				stats.Count("laika.request_total", 1, tags, 1)
				stats.Histogram("laika.request_duration_microseconds", float64(int(time.Since(start).Seconds()*1000000)), tags, 1)
			}(time.Now())

			return next(c)
		}
	}
}

func ResponseEncoderMiddleware() echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			err := next(c)

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
		}
	}
}
