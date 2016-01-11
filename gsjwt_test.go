package gsjwt

import (
	"errors"
	"fmt"
	"testing"

	"time"

	"github.com/dgrijalva/jwt-go"
)

var Key = []byte("Sensual key")

func createToken(t *testing.T) string {
	token := jwt.New(jwt.SigningMethodHS256)

	token.Claims["sub"] = "bar"
	token.Claims["exp"] = time.Now().Add(time.Hour * 72).Unix()

	tokenString, err := token.SignedString(Key)
	if err != nil {
		t.Fatal(err)
	}

	return tokenString
}

// converts string to &string
func stringAddr(s string) *string {
	return &s
}

func TestAuthenticate(t *testing.T) {
	tests := []struct {
		result *string
		err    error
		header string
	}{
		{stringAddr("bar"), nil, fmt.Sprintf("Bearer %s", createToken(t))},
		{nil, errors.New("Invalid field :"), fmt.Sprintf("Bearer: %s", createToken(t))},
		{nil, errors.New("Invalid token"), fmt.Sprintf("Bearer %s", "Stupid token")},
		{nil, errors.New("Invalid header"), ""},
	}

	auth := &Auth{
		Key: Key,
		ReturnFunction: func(values map[string]interface{}) (interface{}, error) {
			return values["sub"], nil
		},
	}

	for _, test := range tests {
		value, err := auth.Authenticate(test.header)
		if (test.err == nil) && (err != nil) {
			t.Fatal("Invalid errors " + err.Error())
		}

		if (test.result == nil) && (value != nil) {
			t.Log(test.result)
			t.Fatal("Invalid result ")
		}

		if (test.result != nil) && (*test.result != value) {
			t.Log(test.result)
			t.Fatal("Invalid result ")
		}
	}
}
