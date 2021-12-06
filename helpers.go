package organisation_api

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
	"net/url"
	"path"
)

func logMsg(logger *log.Logger, msg ...interface{}) {
	if logger != nil {
		logger.Println(msg...)
	}
}

func buildAccountsUrl(c *OrganisationApiClient) (*url.URL, error) {
	clientRootUrlPath := c.ClientConfig.RootUrl.Path

	logMsg(c.ClientConfig.DebugLog, "Joining paths", clientRootUrlPath, "and", accountsPath)

	requestUrl, err := url.Parse(path.Join(clientRootUrlPath, accountsPath))

	if err != nil {
		return nil, err
	}

	requestUrl.Host = c.ClientConfig.RootUrl.Host
	requestUrl.Scheme = c.ClientConfig.RootUrl.Scheme

	return requestUrl, nil
}

func fetchAccountDataFromBody(c *OrganisationApiClient, resp *http.Response) (*AccountData, error) {
	b, err := io.ReadAll(resp.Body)

	if err != nil {
		return nil, err
	}

	if c.ClientConfig.IsDebugEnabled {
		rawBody := string(b)
		logMsg(c.ClientConfig.DebugLog, "Received raw msg", rawBody)
	}

	data := dataHolder{}
	err = json.Unmarshal(b, &data)
	if err != nil {
		return nil, err
	}

	logMsg(c.ClientConfig.DebugLog, "Unmarshalled data from body", data)

	return &data.Data, nil
}

func closeBody(body io.ReadCloser) {
	err := body.Close()
	if err != nil {
		panic(err)
	}
}
