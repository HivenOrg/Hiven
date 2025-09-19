package tests

import (
	"testing"

	"github.com/gofiber/fiber/v2"
)

func TestMiddlewareMissingAuthorizationHeader(t *testing.T) {

	if err := resetTestDB(testDB, cfg.STAGE); err != nil {
		t.Fatal("testDB reset failed")
	}

	res, err := sendRequest(
		testApp,
		"GET",
		"/test/protected",
		nil,
		nil,
	)

	if err != nil {
		t.Fatalf("request failed: %v", err)
	}

	if res.StatusCode != fiber.StatusUnauthorized {
		t.Fatalf("expected status code %d, got %d", fiber.StatusUnauthorized, res.StatusCode)
	}
}

func TestMiddlewareInvalidTokenType(t *testing.T) {

	if err := resetTestDB(testDB, cfg.STAGE); err != nil {
		t.Fatal("testDB reset failed")
	}

	_, token, err := createDummyUser(testDB, cfg.JWT_SECRET_KEY)
	if err != nil {
		t.Fatalf("unable to create dummy user: %v", err)
	}

	res, err := sendRequest(
		testApp,
		"GET",
		"/test/protected",
		headers{"Authorization": "not_bearer " + token},
		nil,
	)

	if err != nil {
		t.Fatalf("request failed: %v", err)
	}

	if res.StatusCode != fiber.StatusUnauthorized {
		t.Fatalf("expected status code %d, got %d", fiber.StatusUnauthorized, res.StatusCode)
	}
}

func TestMiddlewareExpiredToken(t *testing.T) {

	if err := resetTestDB(testDB, cfg.STAGE); err != nil {
		t.Fatal("testDB reset failed")
	}

	expiredToken := "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE3NTgxMDc5ODgsImlkIjoxMDF9.Vep41lkCQLSZBs5zwAkvac32B6ymchI_CMveP3FJwgs"

	res, err := sendRequest(
		testApp,
		"GET",
		"/test/protected",
		headers{"Authorization": "Bearer " + expiredToken},
		nil,
	)

	if err != nil {
		t.Fatalf("request failed: %v", err)
	}

	if res.StatusCode != fiber.StatusUnauthorized {
		t.Fatalf("expected status code %d, got %d", fiber.StatusUnauthorized, res.StatusCode)
	}
}

func TestMiddlewareUserDoesNotExists(t *testing.T) {

	if err := resetTestDB(testDB, cfg.STAGE); err != nil {
		t.Fatal("testDB reset failed")
	}

	_, token, err := createDummyUser(testDB, cfg.JWT_SECRET_KEY)
	if err != nil {
		t.Fatalf("unable to create dummy user: %v", err)
	}

	if err := resetTestDB(testDB, cfg.STAGE); err != nil {
		t.Fatal("testDB reset failed")
	}

	res, err := sendRequest(
		testApp,
		"GET",
		"/test/protected",
		headers{"Authorization": "Bearer " + token},
		nil,
	)

	if err != nil {
		t.Fatalf("request failed: %v", err)
	}

	if res.StatusCode != fiber.StatusUnauthorized {
		t.Fatalf("expected status code %d, got %d", fiber.StatusUnauthorized, res.StatusCode)
	}
}

func TestMiddleware(t *testing.T) {

	if err := resetTestDB(testDB, cfg.STAGE); err != nil {
		t.Fatal("testDB reset failed")
	}

	_, token, err := createDummyUser(testDB, cfg.JWT_SECRET_KEY)
	if err != nil {
		t.Fatalf("unable to create dummy user: %v", err)
	}

	res, err := sendRequest(
		testApp,
		"GET",
		"/test/protected",
		headers{"Authorization": "Bearer " + token},
		nil,
	)

	if err != nil {
		t.Fatalf("request failed: %v", err)
	}

	if res.StatusCode != fiber.StatusOK {
		t.Fatalf("expected status code %d, got %d", fiber.StatusOK, res.StatusCode)
	}
}
