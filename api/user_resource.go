package api

import (
	"github.com/DataDog/datadog-go/statsd"
	"github.com/labstack/echo"

	"github.com/MEDIGO/laika/models"
	"github.com/MEDIGO/laika/store"
)

type UserResource struct {
	store store.Store
	stats *statsd.Client
}

func NewUserResource(store store.Store, stats *statsd.Client) *UserResource {
	return &UserResource{store, stats}
}

func (r *UserResource) Get(c echo.Context) error {
	username := c.Param("username")

	user, err := r.store.GetUserByUsername(username)
	if err != nil {
		if err == store.ErrNoRows {
			return NotFound(c)
		}
		return InternalServerError(c, err)
	}

	return OK(c, user)
}

func (r *UserResource) Create(c echo.Context) error {
	input := struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}{}

	if err := c.Bind(&input); err != nil {
		return BadRequest(c, "Payload must be a valid JSON object")
	}

	if input.Username == "" {
		return Invalid(c, "Username is required")
	}

	if input.Password == "" {
		return Invalid(c, "Password is required")
	}

	user := &models.User{
		Username: input.Username,
		Password: input.Password,
	}

	if err := r.store.CreateUser(user); err != nil {
		return InternalServerError(c, err)
	}

	return Created(c, user)
}
