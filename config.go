package organisation_api

import (
	"log"
	"net/url"
)

type ClientConfig struct {
	RootUrl        *url.URL
	InfoLog        *log.Logger
	ErrorLog       *log.Logger
	DebugLog       *log.Logger
	IsDebugEnabled bool
}

var defaultRootUrl = func() *url.URL {
	d, err := url.Parse("http://localhost:8080/v1/organisation/")
	if err != nil {
		panic(err)
	}

	return d
}()
