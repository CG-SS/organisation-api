package organisation_api

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"path"
)

const accountsPath = "accounts"

func (c *OrganisationApiClient) CreateAccount(data AccountData) (*AccountData, error) {
	requestUrl, err := buildAccountsUrl(c)

	if err != nil {
		logMsg(c.ClientConfig.DebugLog, err.Error())
		return nil, err
	}

	jsonValue, err := json.Marshal(dataHolder{Data: data})

	if err != nil {
		logMsg(c.ClientConfig.DebugLog, err.Error())
		return nil, err
	}

	resp, err := c.Post(requestUrl.String(), "application/json", bytes.NewBuffer(jsonValue))

	if err != nil {
		logMsg(c.ClientConfig.DebugLog, err.Error())
		return nil, err
	}

	statusCode := resp.StatusCode
	logMsg(c.ClientConfig.DebugLog, "Received status: ", resp.Status)
	if statusCode != http.StatusCreated {
		return nil, errors.New(fmt.Sprintf("Received status code %d!", statusCode))
	}

	defer closeBody(resp.Body)

	return fetchAccountDataFromBody(c, resp)
}

func (c *OrganisationApiClient) FetchAccount(id string) (*AccountData, error) {
	requestUrl, err := buildAccountsUrl(c)

	logMsg(c.ClientConfig.DebugLog, "Fetching msg", id)

	if err != nil {
		logMsg(c.ClientConfig.DebugLog, err.Error())
		return nil, err
	}

	requestUrl.Path = path.Join(requestUrl.Path, id)

	resp, err := c.Get(requestUrl.String())

	if err != nil {
		logMsg(c.ClientConfig.DebugLog, err.Error())
		return nil, err
	}

	statusCode := resp.StatusCode
	logMsg(c.ClientConfig.DebugLog, "Received status: ", resp.Status)
	if statusCode != http.StatusOK {
		return nil, errors.New(fmt.Sprintf("Received status code %d!", statusCode))
	}

	defer closeBody(resp.Body)

	return fetchAccountDataFromBody(c, resp)
}

func (c *OrganisationApiClient) DeleteAccount(id string, version int) (bool, error) {
	requestUrl, err := buildAccountsUrl(c)

	if err != nil {
		logMsg(c.ClientConfig.DebugLog, err.Error())
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

	logMsg(c.ClientConfig.DebugLog, "Received status: ", resp.Status)

	return resp.StatusCode == http.StatusNoContent, nil
}
