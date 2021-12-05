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
