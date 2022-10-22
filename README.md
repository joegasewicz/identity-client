# Identity Client
Agnostic identity client

HTTP client for remote api calls that takes a JWT token from
either a cookie or a token value & includes this token in the headers, whilst trying to be as agnostic with
the returned data type as possible.
## Install
```
go get -u github.com/joegasewicz/identity-client
```
### Setup
```go
identity := Identity{
	URL: "http://127.0.01:5000/users",
	CookieName: "token", // Optional
	// If you do not have cookie & require authorisation then set the token value manually
	Token: "Bearer ..."
	Client &http.Client{} // Optional (for testing your client)
}
```
### Methods
#### Get
```go
data, err := identity.Get(r)
```
#### Post
```go
jsonData := map[string]interface{}{"name": "John"}
data, err := identity.Post(r, jsonData)
```

The data returned will be of the following type
```go
map[string]interface{}
```

You can cast each value to the expected type, for example
```go
d := data.(map[string]interface{})
var email := d["email"].(string)
```


### Utils
#### AddTokenPrefix
When you add a token value to the `Identity` type, you might want to add
a `Bearer` prefix to the token value (or if you are unsure whether the incoming 
request token has a Bearer prefix)
```go
i := identity_client.Identity{
    URL:   "http://127.0.01:5000/users",
    Token: identity_client.AddTokenPrefix(r.Header.Get("Authorization")),
}
```