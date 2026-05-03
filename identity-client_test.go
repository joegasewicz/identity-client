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
		if req.Method != "POST" {
			t.Errorf("expected method to be POST but got %s", req.Method)
		}
		if req.Header.Get("Content-Type") != "application/json" {
			t.Errorf("expected Content-Type to be application/json but got %s", req.Header.Get("Content-Type"))
		}
		body, err := io.ReadAll(req.Body)
		if err != nil {
			t.Errorf("expected request body to be readable but got %s", err)
		}
		if string(body) != `{"name":"John"}` {
			t.Errorf("expected request body to be %s but got %s", `{"name":"John"}`, string(body))
		}
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

func TestPutReturnsErrorForNotCookie(t *testing.T) {

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
	_, err := i.Put(req, jsonData)
	expectedErr := "http: named cookie not present"
	if err.Error() != expectedErr {
		t.Logf("expected error to be %s but got %s", expectedErr, err.Error())
		t.Fail()
	}

}

func TestPutReturnsData(t *testing.T) {

	json := `{"data": 123}`

	GetDoFunc = func(req *http.Request) (*http.Response, error) {
		if req.Method != "PUT" {
			t.Errorf("expected method to be PUT but got %s", req.Method)
		}
		if req.Header.Get("Content-Type") != "application/json" {
			t.Errorf("expected Content-Type to be application/json but got %s", req.Header.Get("Content-Type"))
		}
		body, err := io.ReadAll(req.Body)
		if err != nil {
			t.Errorf("expected request body to be readable but got %s", err)
		}
		if string(body) != `{"name":"John"}` {
			t.Errorf("expected request body to be %s but got %s", `{"name":"John"}`, string(body))
		}
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
	data, err := i.Put(req, jsonData)
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

func TestPutWithToken(t *testing.T) {
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
	data, err := i.Put(req, jsonData)
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
		t.Logf("expected token to equal %s but got %s", testToken, result)
		t.Fail()
	}
}

func TestDeleteReturnsErrorForNotCookie(t *testing.T) {

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
	_, err := i.Delete(req)
	expectedErr := "http: named cookie not present"
	if err.Error() != expectedErr {
		t.Logf("expected error to be %s but got %s", expectedErr, err.Error())
		t.Fail()
	}

}

func TestDeleteReturnsData(t *testing.T) {

	json := `{"data": 123}`

	GetDoFunc = func(req *http.Request) (*http.Response, error) {
		if req.Method != "DELETE" {
			t.Errorf("expected method to be DELETE but got %s", req.Method)
		}
		if req.Body != nil {
			t.Errorf("expected request body to be nil")
		}
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

	data, err := i.Delete(req)
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

func TestDeleteWithToken(t *testing.T) {
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

	data, err := i.Delete(req)
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
		t.Logf("expected token to equal %s but got %s", testToken, result)
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
