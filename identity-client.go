package identity_client

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
)

// Identity defines the identity type
type Identity struct {
	URL        string
	CookieName string
}

// Get is a GET method client for remote api calls that takes a JWT token from
// a cookie & includes this token in the headers, whilst trying to be as agnostic with
// the returned data type as possible.
//
//		identity := Identity{
//			URL: "http://127.0.01:5000/users",
//			CookieName: "token",
//		}
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
	req, err := http.NewRequest("GET", i.URL, nil)
	if err != nil {
		return nil, err
	}
	tokenCookie, err := r.Cookie(i.CookieName)
	token := fmt.Sprintf("Bearer %s", tokenCookie.Value)
	req.Header.Set("Authorization", token)
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	if resp.Status >= "400" {
		log.Printf("identity api responded with status %s", resp.Status)
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
