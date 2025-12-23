package utils

import (
	"testing"
)

const sampleSecret = "this-is-secret"

func TestJwtToken(t *testing.T) {

	email := "asd@gmail.com"
	userId := "1"

	jwtToken, err := newToken(userId, email, issuer, accessSubject, 5, sampleSecret)

	if err != nil {
		t.Error(err)
		return
	}

	decodeJson, err := parseToken(jwtToken, issuer, accessSubject, sampleSecret)

	if err != nil {
		t.Error(err)
		return
	}

	if decodeJson.UserId != userId || decodeJson.Email != email {
		t.Errorf("Expected: Email %s, UserID %s. Got Email: %s, UserID: %s.", email, userId, decodeJson.Email, decodeJson.UserId)
	}

}

func TestTokenExpiration(t *testing.T) {

	email := "asd@gmail.com"
	userId := "1"

	jwtToken, err := newToken(userId, email, issuer, accessSubject, 0, sampleSecret) // JWT gets invalid

	if err != nil {
		t.Error(err)
		return
	}

	_, err = parseToken(jwtToken, issuer, accessSubject, sampleSecret)

	if err == nil {
		t.Error("expected token to be expired")
		return
	}

}
