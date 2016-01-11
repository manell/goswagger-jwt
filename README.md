# goswagger-jwt
JWT parsing tool for [go-swagger](https://github.com/go-swagger/go-swagger)

## Installation
Install the package in your ```$GOPATH``` running:
```
go get github.com/manell/goswagger-jwt
```

## Usage
Provided a swagger definition with the following authentication:
```yml
securityDefinitions:
  bearer:
    type: apiKey
    name: Authorization
    in: header
```

Create a new Auth instance with a callback function that validates the field sub in a JWT:
```go
auth := gsjwt.Auth{
	Key: []byte("My secret hmac key"),
	ReturnFunction: func(values map[string]interface{}) (interface{}, error) {
		sub, ok := values["sub"]
		if !ok {
			return nil, errors.New("Sub not provided")
		}
		return sub, nil
	}
}
```

Then just configure the api:
```go
api.BearerAuth = auth.Authenticate
```
