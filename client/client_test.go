package client

import (
	"net/http"
	"net/http/httptest"

	"github.com/stretchr/testify/suite"
)

type ClientSuite struct {
	suite.Suite
	client Client
	server *httptest.Server
	mux    *http.ServeMux
}

func (s *ClientSuite) SetupTest() {
	s.mux = http.NewServeMux()
	s.server = httptest.NewServer(s.mux)

	client, _ := NewClient(s.server.URL)
	s.client = client
}

func (s *ClientSuite) TearDownTest() {
	s.server.Close()
}
