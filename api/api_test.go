package api

import (
	"bytes"
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"net/url"
	"strconv"
	"testing"
)

type TestClient struct {
	server   *httptest.Server
	username string
	password string
}

func NewTestClient(t *testing.T, username string, password string) *TestClient {
	return &TestClient{
		server:   NewTestServer(t),
		username: username,
		password: password,
	}
}

func (tc *TestClient) Close() {
	tc.server.Close()
}

func (tc *TestClient) do(method, endpoint string, in interface{}, out interface{}) error {
	rel, err := url.Parse(endpoint)
	if err != nil {
		return err
	}

	base, err := url.Parse(tc.server.URL)
	if err != nil {
		return err
	}

	url := base.ResolveReference(rel)
	req, err := http.NewRequest(method, url.String(), nil)
	if err != nil {
		return err
	}

	req.SetBasicAuth(tc.username, tc.password)
	req.Header.Set("Content-Type", "application/json")

	if in != nil {
		payload, err := json.Marshal(in)
		if err != nil {
			return err
		}
		buf := bytes.NewBuffer(payload)
		req.Body = ioutil.NopCloser(buf)

		req.ContentLength = int64(len(payload))
		req.Header.Set("Content-Length", strconv.Itoa(len(payload)))
	}

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return err
	}

	defer res.Body.Close()

	if code := res.StatusCode; 200 <= code && code <= 299 {
		if out != nil {
			return json.NewDecoder(res.Body).Decode(out)
		}

		return nil
	}

	e := new(Error)
	if err := json.NewDecoder(res.Body).Decode(e); err != nil {
		return err
	}

	return errors.New(e.Message)
}

func (tc *TestClient) get(endpoint string, out interface{}) error {
	return tc.do("GET", endpoint, nil, out)
}

func (tc *TestClient) post(endpoint string, in, out interface{}) error {
	return tc.do("POST", endpoint, in, out)
}

func (tc *TestClient) put(endpoint string, in, out interface{}) error {
	return tc.do("PUT", endpoint, in, out)
}

func (tc *TestClient) patch(endpoint string, in, out interface{}) error {
	return tc.do("PATCH", endpoint, in, out)
}

func (tc *TestClient) delete(endpoint string, out interface{}) error {
	return tc.do("DELETE", endpoint, nil, out)
}
