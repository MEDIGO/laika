package api

import (
	"errors"
	"time"

	"github.com/DataDog/datadog-go/statsd"
	"github.com/labstack/echo"

	"github.com/MEDIGO/laika/store"
)

type User struct {
	ID        int64      `json:"id"`
	Username  string     `json:"username,omitempty"`
	Password  string     `json:"password,omitempty"`
	CreatedAt time.Time  `json:"created_at,omitempty"`
	UpdatedAt *time.Time `json:"updated_at,omitempty"`
}

func (u *User) Validate() error {
	if u.Username == "" {
		return errors.New("missing username")
	}

	if u.Password == "" {
		return errors.New("missing password")
	}

	return nil
}

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
			return NotFound(c, err)
		}
		return InternalServerError(c, err)
	}

	apiUser := &User{
		ID:        user.ID,
		Username:  user.Username,
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
	}

	return OK(c, apiUser)
}

func (r *UserResource) Create(c echo.Context) error {
	in := new(User)
	if err := c.Bind(&in); err != nil {
		return BadRequest(c, err)
	}

	if err := in.Validate(); err != nil {
		return BadRequest(c, err)
	}

	user := &store.User{
		Username:     in.Username,
		PasswordHash: in.Password,
	}

	if err := r.store.CreateUser(user); err != nil {
		return InternalServerError(c, err)
	}

	in.Password = ""

	return Created(c, in)
}
