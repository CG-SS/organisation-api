package organisation_api

import (
	"net/http"
	"time"
)

// OrganisationApiClient Struct for the API client. It uses http.Client as composition.
type OrganisationApiClient struct {
	*http.Client
	ClientConfig *ClientConfig
}

// DefaultClient Default client with timeout defined.
var DefaultClient = &OrganisationApiClient{
	Client: &http.Client{
		Timeout: 10 * time.Second,
		Transport: &http.Transport{
			TLSHandshakeTimeout: 5 * time.Second,
		},
	},
	ClientConfig: DefaultConfig,
}

// DebugClient Default client for debugging.
var DebugClient = &OrganisationApiClient{
	Client:       http.DefaultClient,
	ClientConfig: DebugConfig,
}
