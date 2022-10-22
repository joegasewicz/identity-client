package identity_client

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log"
	"net/http"
)

type IClient interface {
	Do(req *http.Request) (*http.Response, error)
}

// Identity defines the identity type
//
//	identity := Identity{
//		URL: "http://127.0.01:5000/users",
//		CookieName: "token", // set to "" if a cookie is not required
//		// If you do not have cookie then set the token value manually
//		Token: "Bearer ..." // or identity_client.AddTokenPrefix(r.Header.Get("Authorization"))
//		Client &http.Client{} // Optional
//	}
type Identity struct {
	// URL is the url that the http request will be sent to (Required)
	URL string
	// CookieName is the name of the cookie key (Optional)
	CookieName string
	// Client is the http client (Optional)
	Client IClient
	// Token requires the token value with or without a Bearer prefix.
	// (identity_client will always add a `Bearer` prefix to any requests).
	Token string
}

// Get is a GET method client for remote api calls that takes a JWT token from
// a cookie & includes this token in the headers, whilst trying to be as agnostic with
// the returned data type as possible.
//
//		data, err := identity.Get(r)
//	 	if data == nil { // bail out here }
//
// Then you can cast each value to the expected type, for example
//
//	d := data.(map[string]interface{})
//	var email := d["email"].(string)
//
// The data returned will be of the following type
//
//	map[string]interface{}
func (i *Identity) Get(r *http.Request) (data interface{}, err error) {
	return i.request(r, "GET", nil)
}

// Post is a GET method client for remote api calls that takes a JWT token from
// a cookie & includes this token in the headers, whilst trying to be as agnostic with
// the returned data type as possible.
//
//		jsonData := map[string]interface{}{"name": "John"}
//		data, err := identity.Post(r, jsonData)
//	 	if data == nil { // bail out here }
//
// Then you can cast each value to the expected type, for example
//
//	d := data.(map[string]interface{})
//	var email := d["email"].(string)
//
// The data returned will be of the following type
//
//	map[string]interface{}
func (i *Identity) Post(r *http.Request, jsonData map[string]interface{}) (data interface{}, err error) {
	r.Header.Set("Content-Type", "application/json")
	j, err := json.Marshal(jsonData)
	if err != nil {
		return nil, err
	}
	b := bytes.NewBuffer(j)
	return i.request(r, "POST", b)
}

func (i *Identity) request(r *http.Request, method string, jsonData *bytes.Buffer) (data interface{}, err error) {
	var req *http.Request
	var token string
	if i.Client == nil {
		i.Client = &http.Client{}
	}
	// make sure the body is not of type bytes if no jsonData is passed
	if jsonData == nil {
		req, err = http.NewRequest(method, i.URL, nil)
	} else {
		req, err = http.NewRequest(method, i.URL, jsonData)
	}
	if err != nil {
		return nil, err
	}
	if i.CookieName != "" {
		tokenCookie, err := r.Cookie(i.CookieName)
		if err != nil {
			return nil, err
		}
		token = fmt.Sprintf("Bearer %s", tokenCookie.Value)
	}
	if i.Token != "" {
		token = i.Token
	}
	req.Header.Set("Authorization", token)
	resp, err := i.Client.Do(req)
	if err != nil {
		return nil, err
	}
	if resp.Status >= "400" {
		log.Printf("identity api responded with status %s", resp.Status)
		err = errors.New("The request returned a status of " + resp.Status)
		return nil, err
	}
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(body, &data)
	if err != nil {
		return nil, err
	}
	return data, err
}
