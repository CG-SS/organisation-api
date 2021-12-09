//go:build !integration
// +build !integration

package organisation_api

import (
	"context"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
	"testing"
)

func TestOrganisationApiClient_createRequest(t *testing.T) {
	testCases := []struct {
		name   string
		url    string
		method string
	}{
		{"Creates GET request", "https://localhost/", http.MethodGet},
		{"Creates POST request", "https://localhost/", http.MethodPost},
		{"Creates DELETE request", "https://localhost/", http.MethodDelete},
		{"Creates non safe request", "http://localhost/", http.MethodGet},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			u, err := url.Parse(tc.url)
			if err != nil {
				t.Fatal(err)
			}

			req, err := createRequest(context.Background(), tc.method, *u, nil)
			if err != nil {
				t.Fatal(err)
			}
			if req.Method != tc.method || req.URL.Path != u.Path {
				t.Fatal("Incorrect method or URL!")
			}
		})
	}
}

func TestOrganisationApiClient_fetchAccountDataFromBody(t *testing.T) {
	c := DefaultClient
	country := "GB"
	accType := "Personal"

	testCases := []struct {
		name     string
		expected dataHolder
	}{
		{"Fetch account from body", dataHolder{AccountData{
			Attributes: &AccountAttributes{
				AccountClassification: &accType,
				AccountNumber:         "10000004",
				BankID:                "400302",
				BankIDCode:            "GBDSC",
				BaseCurrency:          "GBP",
				Bic:                   "NWBKGB42",
				Country:               &country,
				Iban:                  "GB28NWBK40030212764204",
				Name:                  []string{"Kelvin", "Klein"},
			},
			ID:             "123e4567-e89b-12d3-a456-426614174129",
			OrganisationID: "123e4567-e89b-12d3-a456-426614174111",
			Type:           "accounts",
		}},
		},
		{"Fetch another account from body", dataHolder{AccountData{
			Attributes: &AccountAttributes{
				AccountClassification: &accType,
				AccountNumber:         "10000004",
				BankID:                "400302",
				BankIDCode:            "GBDSC",
				BaseCurrency:          "GBP",
				Bic:                   "NWBKGB42",
				Country:               &country,
				Iban:                  "GB28NWBK40030212764204",
				Name:                  []string{"Kelvin", "Klein"},
			},
			ID:             "456e4567-e89b-12d3-a456-426614174100",
			OrganisationID: "456e4567-e89b-12d3-a456-426614174100",
			Type:           "accounts",
		}},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			j, err := json.Marshal(tc.expected)
			if err != nil {
				t.Fatal(err)
			}

			resp := http.Response{
				Body: ioutil.NopCloser(strings.NewReader(string(j))),
			}
			dataFromBody, err := fetchAccountDataFromBody(c, &resp)
			if err != nil {
				t.Fatal(err)
			}
			d := dataHolder{*dataFromBody}
			// maybe compare all fields
			if d.Data.ID != tc.expected.Data.ID {
				t.Fatal("Expected", tc.expected, "got", dataFromBody)
			}
		})
	}
}
