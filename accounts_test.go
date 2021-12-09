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
	"time"
)

var c = "GB"
var accType = "Personal"
var version int64 = 0
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
	Version:        &version,
}

type roundTripAux func(r *http.Request) (*http.Response, error)

func (f roundTripAux) RoundTrip(r *http.Request) (*http.Response, error) {
	return f(r)
}

func TestOrganisationApiClient_CreateAccount(t *testing.T) {
	c := &OrganisationApiClient{
		Client: &http.Client{
			Timeout: 10 * time.Second,
			Transport: &http.Transport{
				TLSHandshakeTimeout: 5 * time.Second,
			},
		},
		ClientConfig: &ClientConfig{
			RootUrl:        defaultRootUrl,
			DebugLog:       nil,
			IsDebugEnabled: false,
		},
	}

	c.Client.Transport = roundTripAux(
		func(r *http.Request) (*http.Response, error) {
			if r.URL.Path != "/v1/organisation/accounts" {
				t.Fatal("Wrong request path! Got", r.URL.Path)
			}
			if r.Method != http.MethodPost {
				t.Fatal("Wrong method! Got", r.Method)
			}

			j, err := json.Marshal(mockAccountData)
			if err != nil {
				t.Fatal(err)
			}

			return &http.Response{
				StatusCode: http.StatusCreated,
				Body:       ioutil.NopCloser(strings.NewReader(string(j))),
			}, nil
		},
	)

	clientResponse, err := c.CreateAccount(mockAccountData)
	if err != nil {
		t.Fatal("Got client error", err)
	}
	if !clientResponse.Success {
		t.Fatal("Failed to create account! Got status code", clientResponse.StatusCode)
	}
}

func TestOrganisationApiClient_CreateAccountWithContext(t *testing.T) {
	u, err := url.Parse("https://localhost:12345/")
	if err != nil {
		t.Fatal(err)
	}

	c := &OrganisationApiClient{
		Client: &http.Client{},
		ClientConfig: &ClientConfig{
			RootUrl:        u,
			DebugLog:       nil,
			IsDebugEnabled: false,
		},
	}

	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Millisecond)
	defer cancel()

	_, err = c.CreateAccountWithContext(mockAccountData, ctx)
	if err == nil {
		t.Fatal("Should've failed!")
	}
	if !strings.Contains(err.Error(), "context deadline exceeded") {
		t.Fatal("Error should be context timeout related!")
	}
}

func TestOrganisationApiClient_DeleteAccount(t *testing.T) {
	c := &OrganisationApiClient{
		Client: &http.Client{
			Timeout: 10 * time.Second,
			Transport: &http.Transport{
				TLSHandshakeTimeout: 5 * time.Second,
			},
		},
		ClientConfig: &ClientConfig{
			RootUrl:        defaultRootUrl,
			DebugLog:       nil,
			IsDebugEnabled: false,
		},
	}

	c.Client.Transport = roundTripAux(
		func(r *http.Request) (*http.Response, error) {
			if r.URL.Path != "/v1/organisation/accounts/123e4567-e89b-12d3-a456-426614174129" {
				t.Fatal("Wrong request path! Got", r.URL.Path)
			}
			if r.Method != http.MethodDelete {
				t.Fatal("Wrong method! Got", r.Method)
			}

			return &http.Response{
				StatusCode: http.StatusNoContent,
				Body:       nil,
			}, nil
		},
	)

	clientResponse, err := c.DeleteAccount(mockAccountData.ID, *mockAccountData.Version)
	if err != nil {
		t.Fatal("Got client error", err)
	}
	if !clientResponse.Success {
		t.Fatal("Failed to delete account! Got status code", clientResponse.StatusCode)
	}
}

func TestOrganisationApiClient_DeleteAccountWithContext(t *testing.T) {
	u, err := url.Parse("https://localhost:12345/")
	if err != nil {
		t.Fatal(err)
	}

	c := &OrganisationApiClient{
		Client: &http.Client{},
		ClientConfig: &ClientConfig{
			RootUrl:        u,
			DebugLog:       nil,
			IsDebugEnabled: false,
		},
	}

	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Millisecond)
	defer cancel()

	_, err = c.DeleteAccountWithContext(mockAccountData.ID, *mockAccountData.Version, ctx)
	if err == nil {
		t.Fatal("Should've failed!")
	}
	if !strings.Contains(err.Error(), "context deadline exceeded") {
		t.Fatal("Error should be context timeout related!")
	}
}

func TestOrganisationApiClient_FetchAccount(t *testing.T) {
	c := &OrganisationApiClient{
		Client: &http.Client{
			Timeout: 10 * time.Second,
			Transport: &http.Transport{
				TLSHandshakeTimeout: 5 * time.Second,
			},
		},
		ClientConfig: &ClientConfig{
			RootUrl:        defaultRootUrl,
			DebugLog:       nil,
			IsDebugEnabled: false,
		},
	}

	c.Client.Transport = roundTripAux(
		func(r *http.Request) (*http.Response, error) {
			if r.URL.Path != "/v1/organisation/accounts/123e4567-e89b-12d3-a456-426614174129" {
				t.Fatal("Wrong request path! Got", r.URL.Path)
			}
			if r.Method != http.MethodGet {
				t.Fatal("Wrong method! Got", r.Method)
			}

			j, err := json.Marshal(mockAccountData)
			if err != nil {
				t.Fatal(err)
			}

			return &http.Response{
				StatusCode: http.StatusOK,
				Body:       ioutil.NopCloser(strings.NewReader(string(j))),
			}, nil
		},
	)

	clientResponse, err := c.FetchAccount(mockAccountData.ID)
	if err != nil {
		t.Fatal("Got client error", err)
	}
	if !clientResponse.Success {
		t.Fatal("Failed to fetch account! Got status code", clientResponse.StatusCode)
	}
}

func TestOrganisationApiClient_FetchAccountWithContext(t *testing.T) {
	u, err := url.Parse("https://localhost:12345/")
	if err != nil {
		t.Fatal(err)
	}

	c := &OrganisationApiClient{
		Client: &http.Client{},
		ClientConfig: &ClientConfig{
			RootUrl:        u,
			DebugLog:       nil,
			IsDebugEnabled: false,
		},
	}

	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Millisecond)
	defer cancel()

	_, err = c.FetchAccountWithContext(mockAccountData.ID, ctx)
	if err == nil {
		t.Fatal("Should've failed!")
	}
	if !strings.Contains(err.Error(), "context deadline exceeded") {
		t.Fatal("Error should be context timeout related!")
	}
}
