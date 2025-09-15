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
