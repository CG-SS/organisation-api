package organisation_api

import "net/http"

type OrganisationApiClient struct {
	*http.Client
	ClientConfig *ClientConfig
}
