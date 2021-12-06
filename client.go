package organisation_api

import (
	"net/http"
	"time"
)

type OrganisationApiClient struct {
	*http.Client
	ClientConfig *ClientConfig
}

var DefaultClient = &OrganisationApiClient{
	Client: &http.Client{
		Timeout: 10 * time.Second,
		Transport: &http.Transport{
			TLSHandshakeTimeout: 5 * time.Second,
		},
	},
	ClientConfig: DefaultConfig,
}

var DebugClient = &OrganisationApiClient{
	Client:       http.DefaultClient,
	ClientConfig: DebugConfig,
}
