package api

import (
	"strconv"
	"time"

	"github.com/DataDog/datadog-go/statsd"
	"github.com/MEDIGO/laika/store"
	log "github.com/Sirupsen/logrus"
	"github.com/google/uuid"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"golang.org/x/crypto/bcrypt"
)

// StateMiddleware attaches the state to the request.
func StateMiddleware(store store.Store) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			state, err := store.State()
			if err != nil {
				return InternalServerError(c, err)
			}
			c.Set("state", state)
			return next(c)
		}
	}
}

// TraceMiddleware attaches an ID to the current request.
func TraceMiddleware() echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			if id, err := uuid.NewRandom(); err == nil {
				c.Set("request_id", id.String())
			}
			return next(c)
		}
	}
}

// LogMiddleware logs information about the current request.
func LogMiddleware() echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			defer func(start time.Time) {
				log.WithFields(log.Fields{
					"request_id":                    RequestID(c),
					"request_uri":                   c.Request().URL.String(),
					"request_method":                c.Request().Method,
					"request_remote_addr":           c.Request().RemoteAddr,
					"request_status_code":           c.Response().Status,
					"request_duration_microseconds": int(time.Since(start).Seconds() * 1000000),
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
					"method:" + c.Request().Method,
					"status:" + strconv.Itoa(c.Response().Status),
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
	return middleware.BasicAuth(func(username, password string, c echo.Context) (bool, error) {
		if username == rootUsername {
			return password == rootPassword, nil
		}

		for _, user := range getState(c).Users {
			if user.Username == username {
				return bcrypt.CompareHashAndPassword([]byte(user.PasswordHash), []byte(password)) == nil, nil
			}
		}

		return false, nil
	})
}
