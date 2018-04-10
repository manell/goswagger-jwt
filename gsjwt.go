package gsjwt

import (
	"fmt"
	"strings"

	"github.com/dgrijalva/jwt-go"
)

// ReturnFunction is a function definition that recieves a JWT parsed into a map
// and returns an empty interface or an error
type ReturnFunction func(jwt.Claims) (interface{}, error)

// Auth contains some configuration related to the JWT parsing, as well as a callback
// function for processing the resulting parsed JWT
type Auth struct {
	Key            []byte
	ReturnFunction ReturnFunction
}

// Authenticate parses a JWT from and Authorization header and executes a callback function
// using the parsed JWT values. Returns an error if the authentications fails
func (a *Auth) Authenticate(header string) (interface{}, error) {
	parts := strings.Split(header, " ")
	if len(parts) != 2 || strings.ToLower(parts[0]) != "bearer" {
		return nil, fmt.Errorf("Invalid Authorization header: %s", header)
	}

	token, err := jwt.Parse(parts[1], func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}
		return a.Key, nil
	})
	if err != nil {
		return nil, err
	}

	return a.ReturnFunction(token.Claims)
}
