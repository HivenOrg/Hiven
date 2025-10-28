package tests

import (
	"testing"

	"github.com/gofiber/fiber/v2"
)

func TestLoginInvalidRequestBody(t *testing.T) {

	if err := resetTestDB(testDB, cfg.STAGE); err != nil {
		t.Fatal("testDB reset failed")
	}

	invalidRequestBodies := []payload{
		// Empty body
		{},

		// Email related
		{"password": "shadowhunter123"},                                  // no email
		{"mail": "harry.hunter@test.com", "password": "shadowhunter123"}, // misspelled email
		{"email": []int{1, 2}, "password": "shadowhunter123"},            // invalid email type
		{"email": "harry.hunter", "password": "shadowhunter123"},         // invalid email format
		{"email": "", "password": "shadowhunter123"},                     // empty email

		// Password related
		{"email": "harry.hunter@test.com"},                               // no password
		{"email": "harry.hunter@test.com", "passwrd": "shadowhunter123"}, // misspelled password
		{"email": "harry.hunter@test.com", "password": []int{1, 2}},      // invalid password type
		{"email": "harry.hunter@test.com", "password": "12345"},          // less than required password length
		{"email": "harry.hunter@test.com", "password": ""},               // empty password
	}

	for _, reqBody := range invalidRequestBodies {

		res, err := sendRequest(
			testApp,
			"POST",
			"/auth/login",
			nil,
			reqBody,
		)

		if err != nil {
			t.Fatalf("request failed: %v", err)
		}

		if res.StatusCode != fiber.StatusBadRequest {
			t.Fatalf("expected status code %d, got %d\nPayload: %+v", fiber.StatusBadRequest, res.StatusCode, reqBody)
		}
	}
}

func TestLoginEmailDoesNotExist(t *testing.T) {
	if err := resetTestDB(testDB, cfg.STAGE); err != nil {
		t.Fatal("testDB reset failed")
	}

	res, err := sendRequest(
		testApp,
		"POST",
		"/auth/login",
		nil,
		payload{
			"email":    "harry.hunter@test.com",
			"password": "shadowhunter123",
		},
	)

	if err != nil {
		t.Fatalf("request failed: %v", err)
	}

	if res.StatusCode != fiber.StatusUnauthorized {
		t.Fatalf("expected status code %d, got %d", fiber.StatusUnauthorized, res.StatusCode)
	}
}

func TestLoginIncorrectPassword(t *testing.T) {
	if err := resetTestDB(testDB, cfg.STAGE); err != nil {
		t.Fatal("testDB reset failed")
	}

	dummyUser, _, err := createDummyUser(testDB, cfg.JWT_SECRET_KEY)
	if err != nil {
		t.Fatalf("unable to create dummy user: %v", err)
	}

	res, err := sendRequest(
		testApp,
		"POST",
		"/auth/login",
		nil,
		payload{
			"email":    dummyUser.Email,
			"password": "wrong password",
		},
	)

	if err != nil {
		t.Fatalf("request failed: %v", err)
	}

	if res.StatusCode != fiber.StatusUnauthorized {
		t.Fatalf("expected status code %d, got %d", fiber.StatusUnauthorized, res.StatusCode)
	}
}

func TestLogin(t *testing.T) {
	if err := resetTestDB(testDB, cfg.STAGE); err != nil {
		t.Fatal("testDB reset failed")
	}

	dummyUser, _, err := createDummyUser(testDB, cfg.JWT_SECRET_KEY)
	if err != nil {
		t.Fatalf("unable to create dummy user: %v", err)
	}

	res, err := sendRequest(
		testApp,
		"POST",
		"/auth/login",
		nil,
		payload{
			"email":    dummyUser.Email,
			"password": dummyUser.Password,
		},
	)

	if err != nil {
		t.Fatalf("request failed: %v", err)
	}

	if res.StatusCode != fiber.StatusOK {
		t.Fatalf("expected status code %d, got %d", fiber.StatusOK, res.StatusCode)
	}

	_, err = checkTokenFromResponse(res, cfg.JWT_SECRET_KEY)
	if err != nil {
		t.Fatal("invalid token")
	}
}
