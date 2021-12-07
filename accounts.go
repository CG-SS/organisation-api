package organisation_api

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"path"
)

const accountsPath = "accounts"

var defaultContext = context.Background()

// CreateAccount Creates a new resource given the AccountData. Uses defaultContext as the context.
func (c *OrganisationApiClient) CreateAccount(data AccountData) (*AccountData, error) {
	return c.CreateAccountWithContext(data, defaultContext)
}

// CreateAccountWithContext Creates a new resource given the AccountData with the given context.
func (c *OrganisationApiClient) CreateAccountWithContext(data AccountData, ctx context.Context) (*AccountData, error) {
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

	req, err := createRequest(ctx, http.MethodPost, *requestUrl, bytes.NewBuffer(jsonValue))
	if err != nil {
		logMsg(c.ClientConfig.DebugLog, err.Error())
		return nil, err
	}

	resp, err := c.Do(req)

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

// FetchAccount Fetches the account given an id. Uses defaultContext as the context.
func (c *OrganisationApiClient) FetchAccount(id string) (*AccountData, error) {
	return c.FetchAccountWithContext(id, defaultContext)
}

// FetchAccountWithContext Fetches the account given an id and context.
func (c *OrganisationApiClient) FetchAccountWithContext(id string, ctx context.Context) (*AccountData, error) {
	requestUrl, err := buildAccountsUrl(c)

	logMsg(c.ClientConfig.DebugLog, "Fetching msg", id)

	if err != nil {
		logMsg(c.ClientConfig.DebugLog, err.Error())
		return nil, err
	}

	requestUrl.Path = path.Join(requestUrl.Path, id)

	req, err := createRequest(ctx, http.MethodGet, *requestUrl, nil)
	if err != nil {
		logMsg(c.ClientConfig.DebugLog, err.Error())
		return nil, err
	}

	resp, err := c.Do(req)
	if err != nil {
		logMsg(c.ClientConfig.DebugLog, err.Error())
		return nil, err
	}

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

// DeleteAccount Deletes account with given id and version. Uses defaultContext as the context.
func (c *OrganisationApiClient) DeleteAccount(id string, version int) (bool, error) {
	return c.DeleteAccountWithContext(id, version, defaultContext)
}

// DeleteAccountWithContext Deletes account with given id, version and context.
func (c *OrganisationApiClient) DeleteAccountWithContext(id string, version int, ctx context.Context) (bool, error) {
	requestUrl, err := buildAccountsUrl(c)

	if err != nil {
		logMsg(c.ClientConfig.DebugLog, err.Error())
		return false, err
	}
	requestUrl.Path = path.Join(requestUrl.Path, id)
	requestUrl.RawQuery = fmt.Sprintf("version=%d", version)

	req, err := createRequest(ctx, http.MethodDelete, *requestUrl, nil)
	if err != nil {
		logMsg(c.ClientConfig.DebugLog, err.Error())
		return false, err
	}

	resp, err := c.Do(req)
	if err != nil {
		return false, err
	}

	logMsg(c.ClientConfig.DebugLog, "Received status: ", resp.Status)

	return resp.StatusCode == http.StatusNoContent, nil
}
