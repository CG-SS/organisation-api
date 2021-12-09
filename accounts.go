package organisation_api

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"path"
)

const accountsPath = "accounts"

var defaultContext = context.Background()

// CreateAccount Creates a new resource given the AccountData. Uses defaultContext as the context.
func (c *OrganisationApiClient) CreateAccount(data AccountData) (*ClientResponse, error) {
	return c.CreateAccountWithContext(data, defaultContext)
}

// CreateAccountWithContext Creates a new resource given the AccountData with the given context.
func (c *OrganisationApiClient) CreateAccountWithContext(data AccountData, ctx context.Context) (*ClientResponse, error) {
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
		return &ClientResponse{
			Data:       nil,
			StatusCode: resp.StatusCode,
			Success:    false,
		}, nil
	}

	defer closeBody(resp.Body)

	respData, err := fetchAccountDataFromBody(c, resp)
	if err != nil {
		logMsg(c.ClientConfig.DebugLog, err.Error())
		return nil, err
	}

	return &ClientResponse{
		Data:       respData,
		StatusCode: resp.StatusCode,
		Success:    true,
	}, nil
}

// FetchAccount Fetches the account given an id. Uses defaultContext as the context.
func (c *OrganisationApiClient) FetchAccount(id string) (*ClientResponse, error) {
	return c.FetchAccountWithContext(id, defaultContext)
}

// FetchAccountWithContext Fetches the account given an id and context.
func (c *OrganisationApiClient) FetchAccountWithContext(id string, ctx context.Context) (*ClientResponse, error) {
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
		return &ClientResponse{
			Data:       nil,
			StatusCode: resp.StatusCode,
			Success:    false,
		}, nil
	}

	defer closeBody(resp.Body)

	respData, err := fetchAccountDataFromBody(c, resp)
	if err != nil {
		logMsg(c.ClientConfig.DebugLog, err.Error())
		return nil, err
	}

	return &ClientResponse{
		Data:       respData,
		StatusCode: resp.StatusCode,
		Success:    true,
	}, nil
}

// DeleteAccount Deletes account with given id and version. Uses defaultContext as the context.
func (c *OrganisationApiClient) DeleteAccount(id string, version int64) (*ClientResponse, error) {
	return c.DeleteAccountWithContext(id, version, defaultContext)
}

// DeleteAccountWithContext Deletes account with given id, version and context.
func (c *OrganisationApiClient) DeleteAccountWithContext(id string, version int64, ctx context.Context) (*ClientResponse, error) {
	requestUrl, err := buildAccountsUrl(c)

	if err != nil {
		logMsg(c.ClientConfig.DebugLog, err.Error())
		return nil, err
	}
	requestUrl.Path = path.Join(requestUrl.Path, id)
	requestUrl.RawQuery = fmt.Sprintf("version=%d", version)

	req, err := createRequest(ctx, http.MethodDelete, *requestUrl, nil)
	if err != nil {
		logMsg(c.ClientConfig.DebugLog, err.Error())
		return nil, err
	}

	resp, err := c.Do(req)
	if err != nil {
		return nil, err
	}

	logMsg(c.ClientConfig.DebugLog, "Received status: ", resp.Status)

	return &ClientResponse{
		Data:       nil,
		StatusCode: resp.StatusCode,
		Success:    resp.StatusCode == http.StatusNoContent,
	}, nil
}
