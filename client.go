package organisation_api

import (
	"net/http"
)

type OrganisationApiClient struct {
	*http.Client
	ClientConfig *ClientConfig
}

var DefaultClient = &OrganisationApiClient{
	Client:       http.DefaultClient,
	ClientConfig: DefaultConfig,
}

var DebugClient = &OrganisationApiClient{
	Client:       http.DefaultClient,
	ClientConfig: DebugConfig,
}
