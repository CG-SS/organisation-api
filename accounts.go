package organisation_api

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"path"
	"strconv"
)

const accountsPath = "accounts"

func (c *OrganisationApiClient) CreateAccount(data AccountData) (*AccountData, error) {
	requestUrl, err := buildAccountsUrl(c)

	if err != nil {
		logMsg(c.ClientConfig.ErrorLog, err.Error())
		return nil, err
	}

	values := map[string]AccountData{"data": data}
	jsonValue, err := json.Marshal(values)

	if err != nil {
		logMsg(c.ClientConfig.ErrorLog, err.Error())
		return nil, err
	}

	resp, err := c.Post(requestUrl.String(), "application/json", bytes.NewBuffer(jsonValue))

	if err != nil {
		logMsg(c.ClientConfig.ErrorLog, err.Error())
		return nil, err
	}

	defer closeBody(resp.Body)

	return fetchAccountDataFromBody(resp)
}

func (c *OrganisationApiClient) FetchAccount(id string) (*AccountData, error) {
	requestUrl, err := buildAccountsUrl(c)

	if err != nil {
		logMsg(c.ClientConfig.ErrorLog, err.Error())
		return nil, err
	}

	requestUrl.Path = path.Join(requestUrl.Path, id)

	resp, err := c.Get(requestUrl.String())

	if err != nil {
		logMsg(c.ClientConfig.ErrorLog, err.Error())
		return nil, err
	}

	defer closeBody(resp.Body)

	return fetchAccountDataFromBody(resp)
}

func (c *OrganisationApiClient) DeleteAccount(id string, version int) (bool, error) {
	requestUrl, err := buildAccountsUrl(c)

	if err != nil {
		logMsg(c.ClientConfig.ErrorLog, err.Error())
		return false, err
	}

	requestUrl.Path = path.Join(requestUrl.Path, id)
	requestUrl.RawQuery = fmt.Sprintf("version=%d", version)

	request := http.Request{
		Method: http.MethodDelete,
		URL:    requestUrl,
	}

	resp, err := c.Do(&request)
	if err != nil {
		return false, err
	}

	statusCode, err := strconv.Atoi(resp.Status)
	if err != nil {
		return false, err
	}

	return statusCode == http.StatusNoContent, nil
}
