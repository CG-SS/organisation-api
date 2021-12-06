package organisation_api

import (
	"log"
	"net/url"
	"os"
)

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

var DefaultConfig = &ClientConfig{
	RootUrl:        defaultRootUrl,
	DebugLog:       nil,
	IsDebugEnabled: false,
}

var DebugConfig = &ClientConfig{
	RootUrl:        defaultRootUrl,
	DebugLog:       log.New(os.Stdout, "DEBUG\t", log.Ldate|log.Ltime),
	IsDebugEnabled: true,
}
