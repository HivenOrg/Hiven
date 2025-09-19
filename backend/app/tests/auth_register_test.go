package tests

import (
	"testing"

	"github.com/gofiber/fiber/v2"
)

func TestRegisterInvalidRequestBody(t *testing.T) {

	if err := resetTestDB(testDB, cfg.STAGE); err != nil {
		t.Fatal("testDB reset failed")
	}

	invalidRequestBodies := []payload{
		// Empty body
		{},

		// Email related
		{"password": "shadowhunter123", "firstname": "Harry", "lastname": "Hunter", "phone_number": "+20000000000"},                                  // no email
		{"mail": "harry.hunter@test.com", "password": "shadowhunter123", "firstname": "Harry", "lastname": "Hunter", "phone_number": "+20000000000"}, // misspelled email
		{"email": []int{1, 2}, "password": "shadowhunter123", "firstname": "Harry", "lastname": "Hunter", "phone_number": "+20000000000"},            // invalid email type
		{"email": "harry.hunter", "password": "shadowhunter123", "firstname": "Harry", "lastname": "Hunter", "phone_number": "+20000000000"},         // invalid email format
		{"email": "", "password": "shadowhunter123", "firstname": "Harry", "lastname": "Hunter", "phone_number": "+20000000000"},                     // empty email

		// Password related
		{"email": "harry.hunter@test.com", "firstname": "Harry", "lastname": "Hunter", "phone_number": "+20000000000"},                               // no password
		{"email": "harry.hunter@test.com", "passwrd": "shadowhunter123", "firstname": "Harry", "lastname": "Hunter", "phone_number": "+20000000000"}, // misspelled password
		{"email": "harry.hunter@test.com", "password": []int{1, 2}, "firstname": "Harry", "lastname": "Hunter", "phone_number": "+20000000000"},      // invalid password type
		{"email": "harry.hunter@test.com", "password": "12345", "firstname": "Harry", "lastname": "Hunter", "phone_number": "+20000000000"},          // less than required password length
		{"email": "harry.hunter@test.com", "password": "", "firstname": "Harry", "lastname": "Hunter", "phone_number": "+20000000000"},               // empty password

		// First name related
		{"email": "harry.hunter@test.com", "password": "shadowhunter123", "lastname": "Hunter", "phone_number": "+20000000000"},                           // no first name
		{"email": "harry.hunter@test.com", "password": "shadowhunter123", "firstname": "", "lastname": "Hunter", "phone_number": "+20000000000"},          // empty first name
		{"email": "harry.hunter@test.com", "password": "shadowhunter123", "firstname": []int{1, 2}, "lastname": "Hunter", "phone_number": "+20000000000"}, // invalid firstname type

		// Last name related
		{"email": "harry.hunter@test.com", "password": "shadowhunter123", "firstname": "Harry", "phone_number": "+20000000000"},                          // no last name
		{"email": "harry.hunter@test.com", "password": "shadowhunter123", "firstname": "Harry", "lastname": "", "phone_number": "+20000000000"},          // empty last name
		{"email": "harry.hunter@test.com", "password": "shadowhunter123", "firstname": "Harry", "lastname": []int{1, 2}, "phone_number": "+20000000000"}, // invalid lastname type

		// Phone number related
		{"email": "harry.hunter@test.com", "password": "shadowhunter123", "firstname": "Harry", "lastname": "Hunter"},                               // no phone number
		{"email": "harry.hunter@test.com", "password": "shadowhunter123", "firstname": "Harry", "lastname": "Hunter", "phone_number": ""},           // empty phone number
		{"email": "harry.hunter@test.com", "password": "shadowhunter123", "firstname": "Harry", "lastname": "Hunter", "phone_number": "notaphone"},  // invalid phone number format
		{"email": "harry.hunter@test.com", "password": "shadowhunter123", "firstname": "Harry", "lastname": "Hunter", "phone_number": []int{1, 2}},  // invalid phone number type
		{"email": "harry.hunter@test.com", "password": "shadowhunter123", "firstname": "Harry", "lastname": "Hunter", "phone_number": "2000000000"}, // not in e164 format
	}

	for _, reqBody := range invalidRequestBodies {

		res, err := sendRequest(
			testApp,
			"POST",
			"/auth/register",
			headers{"Content-Type": "application/json"},
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

func TestRegisterEmailAlreadyInUse(t *testing.T) {
	if err := resetTestDB(testDB, cfg.STAGE); err != nil {
		t.Fatal("testDB reset failed")
	}

	dummyUser, err := createDummyUser(testDB)
	if err != nil {
		t.Fatalf("unable to create dummy user: %v", err)
	}

	res, err := sendRequest(
		testApp,
		"POST",
		"/auth/register",
		headers{"Content-Type": "application/json"},
		payload{
			"email":        dummyUser.Email,
			"password":     "shadowhunter123",
			"firstname":    "Harry",
			"lastname":     "Hunter",
			"phone_number": "+12000000000",
		},
	)

	if err != nil {
		t.Fatalf("request failed: %v", err)
	}

	if res.StatusCode != fiber.StatusConflict {
		t.Fatalf("expected status code %d, got %d", fiber.StatusConflict, res.StatusCode)
	}
}

func TestRegisterPhoneNumberAlreadyInUse(t *testing.T) {
	if err := resetTestDB(testDB, cfg.STAGE); err != nil {
		t.Fatal("testDB reset failed")
	}

	dummyUser, err := createDummyUser(testDB)
	if err != nil {
		t.Fatalf("unable to create dummy user: %v", err)
	}

	res, err := sendRequest(
		testApp,
		"POST",
		"/auth/register",
		headers{"Content-Type": "application/json"},
		payload{
			"email":        "harry.hunter@test.com",
			"password":     "shadowhunter123",
			"firstname":    "Harry",
			"lastname":     "Hunter",
			"phone_number": dummyUser.PhoneNumber,
		},
	)

	if err != nil {
		t.Fatalf("request failed: %v", err)
	}

	if res.StatusCode != fiber.StatusConflict {
		t.Fatalf("expected status code %d, got %d", fiber.StatusConflict, res.StatusCode)
	}
}

func TestRegistration(t *testing.T) {
	if err := resetTestDB(testDB, cfg.STAGE); err != nil {
		t.Fatal("testDB reset failed")
	}

	res, err := sendRequest(
		testApp,
		"POST",
		"/auth/register",
		headers{"Content-Type": "application/json"},
		payload{
			"email":        "harry.hunter@test.com",
			"password":     "shadowhunter123",
			"firstname":    "Harry",
			"lastname":     "Hunter",
			"phone_number": "+12000000000",
		},
	)

	if err != nil {
		t.Fatalf("request failed: %v", err)
	}

	if res.StatusCode != fiber.StatusCreated {
		t.Fatalf("expected status code %d, got %d", fiber.StatusCreated, res.StatusCode)
	}

	_, err = checkTokenFromResponse(res, cfg.JWT_SECRET_KEY)
	if err != nil {
		t.Fatal("invalid token")
	}
}
