package organisation_api

import (
	"bytes"
	"encoding/json"
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
