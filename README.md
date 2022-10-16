# Identity Client
Agnostic identity client

GET method client for remote api calls that takes a JWT token from
a cookie & includes this token in the headers, whilst trying to be as agnostic with
the returned data type as possible.
## Install
```
go get -u github.com/joegasewicz/identity-client
```

### Usage
```go
identity := Identity{
	URL: "http://127.0.01:5000/users",
	CookieName: "token",
}
data, err := identity.Get(r)
// cast the data to your required entity type e.g.
user := data.(*UserModel)
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