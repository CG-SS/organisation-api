package organisation_api

import (
	"log"
	"net/url"
	"os"
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

var DefaultConfig = &ClientConfig{
	RootUrl:        defaultRootUrl,
	InfoLog:        log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime),
	ErrorLog:       log.New(os.Stdout, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile),
	DebugLog:       nil,
	IsDebugEnabled: false,
}

var DebugConfig = &ClientConfig{
	RootUrl:        defaultRootUrl,
	InfoLog:        log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime),
	ErrorLog:       log.New(os.Stdout, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile),
	DebugLog:       log.New(os.Stdout, "DEBUG\t", log.Ldate|log.Ltime),
	IsDebugEnabled: true,
}
