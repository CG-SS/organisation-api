//go:build integration
// +build integration

package organisation_api

import (
	"testing"
)

var c = "GB"
var accType = "Personal"
var mockAccountData = AccountData{
	Attributes: &AccountAttributes{
		AccountClassification: &accType,
		AccountNumber:         "10000004",
		BankID:                "400302",
		BankIDCode:            "GBDSC",
		BaseCurrency:          "GBP",
		Bic:                   "NWBKGB42",
		Country:               &c,
		Iban:                  "GB28NWBK40030212764204",
		Name:                  []string{"Kelvin", "Klein"},
	},
	ID:             "123e4567-e89b-12d3-a456-426614174129",
	OrganisationID: "123e4567-e89b-12d3-a456-426614174111",
	Type:           "accounts",
}

func Test_Integration_OrganisationApiClient(t *testing.T) {
	client := DefaultClient

	testAccountCreation(t, client)
	testAccountFetching(t, client)
	testAccountDeletion(t, client)
}

func testAccountCreation(t *testing.T, client *OrganisationApiClient) {
	testCases := []struct {
		name       string
		payload    AccountData
		shouldFail bool
	}{
		{"Should create account", mockAccountData, false},
		{"Shouldn't create account", AccountData{}, true},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			r, err := client.CreateAccount(tc.payload)
			if err != nil {
				t.Fatal("Got client API error! Error:", err)
			}

			if r.Success == tc.shouldFail || !r.Success == !tc.shouldFail {
				t.Fatal("Fail condition didn't match! Got", err, "should've been", tc.shouldFail)
			}
		})
	}
}

func testAccountFetching(t *testing.T, client *OrganisationApiClient) {
	testCases := []struct {
		name       string
		uuid       string
		shouldFail bool
	}{
		{"Should fetch account", mockAccountData.ID, false},
		{"Shouldn't fetch account", "123", true},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			r, err := client.FetchAccount(tc.uuid)

			if err != nil {
				t.Fatal("Got client API error! Error:", err)
			}
			if r.Success == tc.shouldFail || !r.Success == !tc.shouldFail {
				t.Fatal("Fail condition didn't match! Got", err, "should've been", tc.shouldFail)
			}
		})
	}
}

func testAccountDeletion(t *testing.T, client *OrganisationApiClient) {
	testCases := []struct {
		name       string
		uuid       string
		shouldFail bool
	}{
		{"Should delete account", mockAccountData.ID, false},
		{"Should fail to delete account", "123", true},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			r, err := client.DeleteAccount(tc.uuid, 0)

			if err != nil {
				t.Fatal("Got client API error! Error:", err)
			}
			if r.Success == tc.shouldFail || !r.Success == !tc.shouldFail {
				t.Fatal("Fail condition didn't match! Got", err, "should've been", tc.shouldFail)
			}
		})
	}
}
