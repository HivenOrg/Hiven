package tests

import (
	"fmt"
	"testing"

	"github.com/gofiber/fiber/v2"
)

func TestHiveRoutesMiddleware(t *testing.T) {

	if err := resetTestDB(testDB, cfg.STAGE); err != nil {
		t.Fatal("testDB reset failed")
	}

	res, err := sendRequest(
		testApp,
		"POST",
		"/hive/create",
		nil,
		payload{
			"name":    "New York Apartment",
			"address": "57th Street, near Central Park",
		},
	)

	if err != nil {
		t.Fatalf("request failed: %v", err)
	}

	if res.StatusCode != fiber.StatusUnauthorized {
		t.Fatalf("expected status code %d, got %d", fiber.StatusUnauthorized, res.StatusCode)
	}
}

func TestCreateHiveInvalidRequestBody(t *testing.T) {

	if err := resetTestDB(testDB, cfg.STAGE); err != nil {
		t.Fatal("testDB reset failed")
	}

	_, token, err := createDummyUser(testDB, cfg.JWT_SECRET_KEY)
	if err != nil {
		t.Fatalf("unable to create dummy user: %v", err)
	}

	invalidRequestBodies := []payload{
		// Empty body
		{},

		// name related
		{"address": "48 Camden High Street Camden Town London"},                         // no name
		{"nam": "", "address": "48 Camden High Street Camden Town London"},              // mispelled
		{"name": "", "address": "48 Camden High Street Camden Town London"},             // empty
		{"name": []int{0, 1, 2}, "address": "48 Camden High Street Camden Town London"}, // invalid

		// address related
		{"name": "London Home"}, // no address
		{"name": "London Home", "adddres": "48 Camden High Street Camden Town London"}, // misspelled
		{"name": "London Home", "address": ""},                                         // empty
		{"name": "London Home", "address": []int{0, 1, 2}},                             // invalid
	}

	for _, body := range invalidRequestBodies {

		res, err := sendRequest(
			testApp,
			"POST",
			"/hive/create",
			headers{"Authorization": "Bearer " + token},
			body,
		)

		if err != nil {
			t.Fatalf("request failed: %v", err)
		}

		if res.StatusCode != fiber.StatusBadRequest {
			t.Fatalf("expected status code %d, got %d", fiber.StatusBadRequest, res.StatusCode)
		}
	}
}

func TestCreateHive(t *testing.T) {

	if err := resetTestDB(testDB, cfg.STAGE); err != nil {
		t.Fatal("testDB reset failed")
	}

	_, token, err := createDummyUser(testDB, cfg.JWT_SECRET_KEY)
	if err != nil {
		t.Fatalf("unable to create dummy user: %v", err)
	}

	res, err := sendRequest(
		testApp,
		"POST",
		"/hive/create",
		headers{"Authorization": "Bearer " + token},
		payload{"name": "London Home", "address": "48 Camden High Street Camden Town London"},
	)

	if err != nil {
		t.Fatalf("request failed: %v", err)
	}

	if res.StatusCode != fiber.StatusCreated {
		t.Fatalf("expected status code %d, got %d", fiber.StatusCreated, res.StatusCode)
	}
}

func TestUpdateDisplayImageInvalidBody(t *testing.T) {

	if err := resetTestDB(testDB, cfg.STAGE); err != nil {
		t.Fatal("testDB reset failed")
	}

	if err := resetTestStorage(testStorage); err != nil {
		t.Fatal("testStorage reset failed")
	}

	user, token, err := createDummyUser(testDB, cfg.JWT_SECRET_KEY)
	if err != nil {
		t.Fatalf("unable to create dummy user: %v", err)
	}

	hive, err := createDummyHive(testDB, user.ID)
	if err != nil {
		t.Fatalf("unable to create dummy hive: %v", err)
	}

	correctFile := "test_data/logo.png"
	incorrectFile := "test_data/logo.heic"
	invalidHiveID := hive.ID + 999

	type TestCase struct {
		hiveID               *uint
		filePath             *string
		expectedResponseCode int
	}

	invalidTestCases := []TestCase{
		{nil, nil, fiber.StatusBadRequest},                    // empty
		{nil, &correctFile, fiber.StatusBadRequest},           // missing hive_id
		{&invalidHiveID, &correctFile, fiber.StatusForbidden}, // invalid hive_id
		{&hive.ID, nil, fiber.StatusBadRequest},               // missing file
		{&hive.ID, &incorrectFile, fiber.StatusBadRequest},    // invalid file format
	}

	for _, tc := range invalidTestCases {

		fields := map[string]string{}
		files := map[string]string{}

		if tc.hiveID != nil {
			fields["hive_id"] = fmt.Sprintf("%d", *tc.hiveID)
		}

		if tc.filePath != nil {
			files["file"] = *tc.filePath
		}

		res, err := sendFormRequest(
			testApp,
			"POST",
			"/hive/update-display-image",
			headers{"Authorization": "Bearer " + token},
			fields,
			files,
		)

		if err != nil {
			t.Fatalf("request failed: %v", err)
		}

		if res.StatusCode != tc.expectedResponseCode {
			t.Fatalf("expected status code %d, got %d", tc.expectedResponseCode, res.StatusCode)
		}
	}
}

func TestUpdateDisplayImage(t *testing.T) {

	if err := resetTestDB(testDB, cfg.STAGE); err != nil {
		t.Fatal("testDB reset failed")
	}

	if err := resetTestStorage(testStorage); err != nil {
		t.Fatal("testStorage reset failed")
	}

	user, token, err := createDummyUser(testDB, cfg.JWT_SECRET_KEY)
	if err != nil {
		t.Fatalf("unable to create dummy user: %v", err)
	}

	hive, err := createDummyHive(testDB, user.ID)
	if err != nil {
		t.Fatalf("unable to create dummy hive: %v", err)
	}

	res, err := sendFormRequest(
		testApp,
		"POST",
		"/hive/update-display-image",
		headers{"Authorization": "Bearer " + token},
		map[string]string{"hive_id": fmt.Sprintf("%d", hive.ID)},
		map[string]string{"file": "test_data/logo.png"},
	)

	if err != nil {
		t.Fatalf("request failed: %v", err)
	}

	if res.StatusCode != fiber.StatusOK {
		t.Fatalf("expected status code %d, got %d", fiber.StatusOK, res.StatusCode)
	}
}

func TestGetHivesList(t *testing.T) {

	if err := resetTestDB(testDB, cfg.STAGE); err != nil {
		t.Fatal("testDB reset failed")
	}

	if err := resetTestStorage(testStorage); err != nil {
		t.Fatal("testStorage reset failed")
	}

	user, token, err := createDummyUser(testDB, cfg.JWT_SECRET_KEY)
	if err != nil {
		t.Fatalf("unable to create dummy user: %v", err)
	}

	hive, err := createDummyHive(testDB, user.ID)
	if err != nil {
		t.Fatalf("unable to create dummy hive: %v", err)
	}

	res, err := sendRequest(
		testApp,
		"GET",
		"/hive/my-hives",
		headers{"Authorization": "Bearer " + token},
		nil,
	)

	if err != nil {
		t.Fatalf("request failed: %v", err)
	}

	if res.StatusCode != fiber.StatusOK {
		t.Fatalf("expected status code %d, got %d", fiber.StatusOK, res.StatusCode)
	}

	// Parsing response
	body, err := parseResponse(res)
	if err != nil {
		t.Fatalf("failed to parse response: %v", err)
	}
	resHives, ok := body["my_hives"]
	if !ok {
		t.Fatalf("expected 'my_hives' key in response body")
	}

	// Checking response contains a list of length 1
	hivesSlice, ok := resHives.([]any)
	if !ok {
		t.Fatalf("'my_hives' should be an array, got %T", resHives)
	}
	if len(hivesSlice) != 1 {
		t.Fatalf("expected 1 hive in response, got %d", len(hivesSlice))
	}

	// Matching hiveID
	firstHive, ok := hivesSlice[0].(map[string]any)
	if !ok {
		t.Fatalf("hive element should be an object, got %T", hivesSlice[0])
	}
	res_id, ok := firstHive["id"].(float64)
	if !ok {
		t.Fatalf("expected id to be a number, got %T", firstHive["id"])
	}
	if uint(res_id) != hive.ID {
		t.Fatal("incorrect hive_id in response")
	}
}

func TestEditHiveInvalidBodies(t *testing.T) {

	if err := resetTestDB(testDB, cfg.STAGE); err != nil {
		t.Fatal("testDB reset failed")
	}

	if err := resetTestStorage(testStorage); err != nil {
		t.Fatal("testStorage reset failed")
	}

	owner, token, err := createDummyUser(testDB, cfg.JWT_SECRET_KEY)
	if err != nil {
		t.Fatalf("unable to create dummy user: %v", err)
	}

	hive, err := createDummyHive(testDB, owner.ID)
	if err != nil {
		t.Fatalf("unable to create dummy hive: %v", err)
	}

	invalidRequestBodies := []payload{
		// Empty body
		{},

		// hive_id
		{"name": "updated", "address": "updated"},                             // missing
		{"hive_id": []uint{1, 2, 3}, "name": "updated", "address": "updated"}, // invalid format
		{"hive_id": 0, "name": "updated", "address": "updated"},               // 0 always fails

		// name
		{"hive_id": hive.ID, "address": "updated"},                          // missing
		{"hive_id": hive.ID, "name": []uint{1, 2, 3}, "address": "updated"}, // invalid format
		{"hive_id": hive.ID, "name": "", "address": "updated"},              // empty

		// address
		{"hive_id": hive.ID, "name": "updated"},                             // missing
		{"hive_id": hive.ID, "name": "updated", "address": []uint{1, 2, 3}}, // invalid format
		{"hive_id": hive.ID, "name": "updated", "address": ""},              // empty
	}

	for _, body := range invalidRequestBodies {
		res, err := sendRequest(
			testApp,
			"PUT",
			"/hive/edit",
			headers{"Authorization": "Bearer " + token},
			body,
		)

		if err != nil {
			t.Fatalf("request failed: %v", err)
		}

		if res.StatusCode != fiber.StatusBadRequest {
			t.Fatalf("expected status code %d, got %d", fiber.StatusBadRequest, res.StatusCode)
		}
	}
}

func TestEditHiveUnauthorized(t *testing.T) {

	if err := resetTestDB(testDB, cfg.STAGE); err != nil {
		t.Fatal("testDB reset failed")
	}

	if err := resetTestStorage(testStorage); err != nil {
		t.Fatal("testStorage reset failed")
	}

	// owner
	owner, _, err := createDummyUser(testDB, cfg.JWT_SECRET_KEY)
	if err != nil {
		t.Fatalf("unable to create dummy user: %v", err)
	}

	// member
	_, token, err := createDummyUser(testDB, cfg.JWT_SECRET_KEY)
	if err != nil {
		t.Fatalf("unable to create dummy user: %v", err)
	}

	hive, err := createDummyHive(testDB, owner.ID)
	if err != nil {
		t.Fatalf("unable to create dummy hive: %v", err)
	}

	res, err := sendRequest(
		testApp,
		"PUT",
		"/hive/edit",
		headers{"Authorization": "Bearer " + token},
		payload{"hive_id": hive.ID, "name": "updated", "address": "updated"},
	)

	if err != nil {
		t.Fatalf("request failed: %v", err)
	}

	if res.StatusCode != fiber.StatusForbidden {
		t.Fatalf("expected status code %d, got %d", fiber.StatusForbidden, res.StatusCode)
	}
}

func TestEditHive(t *testing.T) {

	if err := resetTestDB(testDB, cfg.STAGE); err != nil {
		t.Fatal("testDB reset failed")
	}

	if err := resetTestStorage(testStorage); err != nil {
		t.Fatal("testStorage reset failed")
	}

	user, token, err := createDummyUser(testDB, cfg.JWT_SECRET_KEY)
	if err != nil {
		t.Fatalf("unable to create dummy user: %v", err)
	}

	hive, err := createDummyHive(testDB, user.ID)
	if err != nil {
		t.Fatalf("unable to create dummy hive: %v", err)
	}

	res, err := sendRequest(
		testApp,
		"PUT",
		"/hive/edit",
		headers{"Authorization": "Bearer " + token},
		payload{"hive_id": hive.ID, "name": "updated", "address": "updated"},
	)

	if err != nil {
		t.Fatalf("request failed: %v", err)
	}

	if res.StatusCode != fiber.StatusOK {
		t.Fatalf("expected status code %d, got %d", fiber.StatusOK, res.StatusCode)
	}
}

func TestDeleteHiveInvalidQueryParameter(t *testing.T) {

	if err := resetTestDB(testDB, cfg.STAGE); err != nil {
		t.Fatal("testDB reset failed")
	}

	if err := resetTestStorage(testStorage); err != nil {
		t.Fatal("testStorage reset failed")
	}

	_, token, err := createDummyUser(testDB, cfg.JWT_SECRET_KEY)
	if err != nil {
		t.Fatalf("unable to create dummy user: %v", err)
	}

	urls := []string{
		"/hive/delete",        //missing
		"/hive/delete?id=abc", //invalid
	}

	for _, url := range urls {
		res, err := sendRequest(
			testApp,
			"DELETE",
			url,
			headers{"Authorization": "Bearer " + token},
			nil,
		)

		if err != nil {
			t.Fatalf("request failed: %v", err)
		}

		if res.StatusCode != fiber.StatusBadRequest {
			t.Fatalf("expected status code %d, got %d", fiber.StatusBadRequest, res.StatusCode)
		}
	}
}

func TestDeleteHiveUnauthorized(t *testing.T) {

	if err := resetTestDB(testDB, cfg.STAGE); err != nil {
		t.Fatal("testDB reset failed")
	}

	if err := resetTestStorage(testStorage); err != nil {
		t.Fatal("testStorage reset failed")
	}

	// owner
	owner, _, err := createDummyUser(testDB, cfg.JWT_SECRET_KEY)
	if err != nil {
		t.Fatalf("unable to create dummy user: %v", err)
	}

	// member
	_, token, err := createDummyUser(testDB, cfg.JWT_SECRET_KEY)
	if err != nil {
		t.Fatalf("unable to create dummy user: %v", err)
	}

	hive, err := createDummyHive(testDB, owner.ID)
	if err != nil {
		t.Fatalf("unable to create dummy hive: %v", err)
	}

	url := fmt.Sprintf("/hive/delete?id=%d", hive.ID)

	res, err := sendRequest(
		testApp,
		"DELETE",
		url,
		headers{"Authorization": "Bearer " + token},
		nil,
	)

	if err != nil {
		t.Fatalf("request failed: %v", err)
	}

	if res.StatusCode != fiber.StatusForbidden {
		t.Fatalf("expected status code %d, got %d", fiber.StatusForbidden, res.StatusCode)
	}
}

func TestDeleteHive(t *testing.T) {

	if err := resetTestDB(testDB, cfg.STAGE); err != nil {
		t.Fatal("testDB reset failed")
	}

	if err := resetTestStorage(testStorage); err != nil {
		t.Fatal("testStorage reset failed")
	}

	// owner
	owner, token, err := createDummyUser(testDB, cfg.JWT_SECRET_KEY)
	if err != nil {
		t.Fatalf("unable to create dummy user: %v", err)
	}

	hive, err := createDummyHive(testDB, owner.ID)
	if err != nil {
		t.Fatalf("unable to create dummy hive: %v", err)
	}

	url := fmt.Sprintf("/hive/delete?id=%d", hive.ID)

	res, err := sendRequest(
		testApp,
		"DELETE",
		url,
		headers{"Authorization": "Bearer " + token},
		nil,
	)

	if err != nil {
		t.Fatalf("request failed: %v", err)
	}

	if res.StatusCode != fiber.StatusNoContent {
		t.Fatalf("expected status code %d, got %d", fiber.StatusNoContent, res.StatusCode)
	}
}
