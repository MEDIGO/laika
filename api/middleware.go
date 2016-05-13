package api

import (
	"strconv"
	"time"

	"github.com/DataDog/datadog-go/statsd"
	"github.com/MEDIGO/laika/store"
	log "github.com/Sirupsen/logrus"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"golang.org/x/crypto/bcrypt"
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

// AuthMiddleware checks login credentials.
func AuthMiddleware(rootUsername, rootPassword string, s store.Store) echo.MiddlewareFunc {
	return middleware.BasicAuth(func(username, password string) bool {
		if username == rootUsername {
			return password == rootPassword
		}

		user, err := s.GetUserByUsername(username)
		if err != nil {
			log.Error("Failed to retrieve user: ", err)
			return false
		}

		return bcrypt.CompareHashAndPassword([]byte(user.PasswordHash), []byte(password)) == nil
	})
}
