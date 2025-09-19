package tests

import (
	"testing"

	"github.com/HivenOrg/Hiven/utils"
)

func TestPasswordFunctions(t *testing.T) {

	plainPassword := "1234abcd"
	wrongPassword := "abcd1234"
	hashedPassword, err := utils.HashPassword(plainPassword)

	if err != nil {
		t.Fatal("unable to hash password")
	}

	matched := utils.CheckPassword(hashedPassword, plainPassword)

	if !matched {
		t.Fatal("password matching failed: correct password + hashed password")
	}

	matched = utils.CheckPassword(hashedPassword, wrongPassword)
	if matched {
		t.Fatal("password matching succeeded: incorrect password + hashed password")
	}
}

func TestGenerateAndValidateJwt(t *testing.T) {

	// Generate token
	token, err := utils.GenerateJWT(101, "testkey", 1)
	if err != nil {
		t.Fatal("unable to generate JWT")
	}

	// Validate token
	userIdFromToken, err := utils.ValidateJWT(token, "testkey")
	if err != nil {
		t.Fatal("unable to validate JWT")
	}

	// Checking user id in token
	if userIdFromToken != 101 {
		t.Fatal("incorrect userID in token")
	}

	// Checking validation with wrong key
	_, err = utils.ValidateJWT(token, "wrongkey")
	if err == nil {
		t.Fatal("JWT validation succeeded for wrong key")
	}
}

func TestExpiredToken(t *testing.T) {

	expiredToken := "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE3NTgxMDc5ODgsImlkIjoxMDF9.Vep41lkCQLSZBs5zwAkvac32B6ymchI_CMveP3FJwgs"

	_, err := utils.ValidateJWT(expiredToken, "testkey")
	if err == nil {
		t.Fatal("Validated expired token")
	}
}
