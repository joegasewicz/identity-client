package identity_client

import "net/http"

var (
	GetDoFunc func(req *http.Request) (*http.Response, error)
)

type MockClient struct{}

func (m *MockClient) Do(req *http.Request) (*http.Response, error) {
	return GetDoFunc(req)
}
