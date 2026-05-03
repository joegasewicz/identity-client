package identity_client

import (
	"bytes"
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestGetReturnsErrorForNotCookie(t *testing.T) {

	json := `{}`

	GetDoFunc = func(req *http.Request) (*http.Response, error) {
		r := io.NopCloser(bytes.NewReader([]byte(json)))
		return &http.Response{
			StatusCode: 200,
			Body:       r,
		}, nil
	}

	i := Identity{
		URL:        "",
		CookieName: "token",
		Client:     &MockClient{},
	}

	req := httptest.NewRequest("GET", "http://domain.com", nil)
	_, err := i.Get(req)
	expectedErr := "http: named cookie not present"
	if err.Error() != expectedErr {
		t.Logf("expected error to be %s but got %s", expectedErr, err.Error())
		t.Fail()
	}

}

func TestGetReturnsData(t *testing.T) {

	json := `{"data": 123}`

	GetDoFunc = func(req *http.Request) (*http.Response, error) {
		r := io.NopCloser(bytes.NewReader([]byte(json)))
		return &http.Response{
			StatusCode: 200,
			Body:       r,
		}, nil
	}

	i := Identity{
		URL:        "",
		CookieName: "",
		Client:     &MockClient{},
	}
	req := httptest.NewRequest("GET", "http://domain.com", nil)
	req.Header.Set("token", "<VALUE>")

	data, err := i.Get(req)
	if err != nil {
		t.Logf("expected error to be nil but got %e", err)
		t.Fail()
	}
	if data == nil {
		t.Log("Expected data but got nil")
		t.Fail()
	}

	result := data.(map[string]interface{})
	if result["data"].(float64) != 123 {
		t.Logf("expected data to equal 123 but got %s", result)
		t.Fail()
	}
}

func TestPostReturnsErrorForNotCookie(t *testing.T) {

	json := `{}`

	GetDoFunc = func(req *http.Request) (*http.Response, error) {
		r := io.NopCloser(bytes.NewReader([]byte(json)))
		return &http.Response{
			StatusCode: 200,
			Body:       r,
		}, nil
	}

	i := Identity{
		URL:        "",
		CookieName: "token",
		Client:     &MockClient{},
	}

	req := httptest.NewRequest("GET", "http://domain.com", nil)
	jsonData := map[string]interface{}{
		"name": "John",
	}
	_, err := i.Post(req, jsonData)
	expectedErr := "http: named cookie not present"
	if err.Error() != expectedErr {
		t.Logf("expected error to be %s but got %s", expectedErr, err.Error())
		t.Fail()
	}

}

func TestPostReturnsData(t *testing.T) {

	json := `{"data": 123}`

	GetDoFunc = func(req *http.Request) (*http.Response, error) {
		r := io.NopCloser(bytes.NewReader([]byte(json)))
		return &http.Response{
			StatusCode: 200,
			Body:       r,
		}, nil
	}

	i := Identity{
		URL:        "",
		CookieName: "",
		Client:     &MockClient{},
	}
	req := httptest.NewRequest("GET", "http://domain.com", nil)
	req.Header.Set("token", "<VALUE>")

	jsonData := map[string]interface{}{
		"name": "John",
	}
	data, err := i.Post(req, jsonData)
	if err != nil {
		t.Logf("expected error to be nil but got %e", err)
		t.Fail()
	}
	if data == nil {
		t.Log("Expected data but got nil")
		t.Fail()
	}

	result := data.(map[string]interface{})
	if result["data"].(float64) != 123 {
		t.Logf("expected data to equal 123 but got %s", result)
		t.Fail()
	}
}

func TestPostWithToken(t *testing.T) {
	testToken := "Bearer <TOKEN>"
	GetDoFunc = func(req *http.Request) (*http.Response, error) {
		json := fmt.Sprintf(`{"token": "%s"}`, req.Header.Get("Authorization"))
		r := io.NopCloser(bytes.NewReader([]byte(json)))

		return &http.Response{
			StatusCode: 200,
			Body:       r,
		}, nil
	}

	i := Identity{
		URL:    "",
		Token:  testToken,
		Client: &MockClient{},
	}
	req := httptest.NewRequest("GET", "http://domain.com", nil)
	req.Header.Set("token", "<VALUE>")

	jsonData := map[string]interface{}{
		"name": "John",
	}
	data, err := i.Post(req, jsonData)
	if err != nil {
		t.Logf("expected error to be nil but got %e", err)
		t.Fail()
	}
	if data == nil {
		t.Log("Expected data but got nil")
		t.Fail()
	}

	result := data.(map[string]interface{})
	if result["token"].(string) != testToken {
		t.Logf("expected data to equal 123 but got %s", result)
		t.Fail()
	}
}
