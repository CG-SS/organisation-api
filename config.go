package organisation_api

import (
	"log"
	"net/url"
	"os"
)

// ClientConfig Struct representing the client config.
type ClientConfig struct {
	RootUrl        *url.URL
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

// DefaultConfig Default config with no logging or debugging.
var DefaultConfig = &ClientConfig{
	RootUrl:        defaultRootUrl,
	DebugLog:       nil,
	IsDebugEnabled: false,
}

// DebugConfig Config meant for debugging capabilities, with logging.
var DebugConfig = &ClientConfig{
	RootUrl:        defaultRootUrl,
	DebugLog:       log.New(os.Stdout, "DEBUG\t", log.Ldate|log.Ltime),
	IsDebugEnabled: true,
}
