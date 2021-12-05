package organisation_api

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
	"net/url"
	"path"
	"strings"
)

func logMsg(logger *log.Logger, msg ...string) {
	if logger != nil {
		logger.Println(strings.Join(msg, " "))
	}
}

func buildAccountsUrl(c *OrganisationApiClient) (*url.URL, error) {
	clientRootUrlPath := c.ClientConfig.RootUrl.Path

	if c.ClientConfig.IsDebugEnabled {
		logMsg(c.ClientConfig.DebugLog, "Joining paths", clientRootUrlPath, "and", accountsPath)
	}

	requestUrl, err := url.Parse(path.Join(clientRootUrlPath, accountsPath))

	if err != nil {
		return nil, err
	}

	requestUrl.Host = c.ClientConfig.RootUrl.Host
	requestUrl.Scheme = c.ClientConfig.RootUrl.Scheme

	return requestUrl, nil
}

func fetchAccountDataFromBody(resp *http.Response) (*AccountData, error) {
	data := AccountData{}
	err := json.NewDecoder(resp.Body).Decode(&data)
	if err != nil {
		return nil, err
	}

	return &data, nil
}

func closeBody(body io.ReadCloser) {
	err := body.Close()
	if err != nil {
		panic(err)
	}
}
