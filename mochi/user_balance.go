package mochi

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"time"

	"github.com/consolelabs/mochi-go-sdk/mochi/model"
)

func (c *Client) GetUserBalances(profileID string) ([]model.UserTokenBalance, error) {
	client := &http.Client{}
	requestURL := fmt.Sprintf("%s/api/v1/mochi-wallet/%v/balances/icy", c.cfg.MochiPay.BaseURL, profileID)
	request, err := http.NewRequest(http.MethodGet, requestURL, nil)
	if err != nil {
		return nil, err
	}

	resp, err := client.Do(request)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()
	resBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode != http.StatusOK {
		var errMsg model.ErrorMessage
		if err := json.Unmarshal(resBody, &errMsg); err != nil {
			return nil, fmt.Errorf("[GetUserBalances] failed to decode error, error %v", err.Error())
		}
		return nil, fmt.Errorf("[GetUserBalances] failed to call , error %v", errMsg.Msg)
	}

	var respData model.UserTokenBalanceResponses

	if err := json.Unmarshal(resBody, &respData); err != nil {
		return nil, errors.New("[GetUserBalances] invalid decoded, error " + err.Error())
	}

	return respData.Data, nil
}

func (c *Client) GetBatchBalances(profileIds []string) (*model.BatchBalancesResponse, error) {
	var client = &http.Client{
		Timeout: 10 * time.Second,
	}

	url := fmt.Sprintf("%s/api/v1/mochi-wallet/balances/multiple", c.cfg.MochiPay.BaseURL)
	body, err := json.Marshal(struct {
		ProfileIDs []string `json:"profile_ids"`
	}{ProfileIDs: profileIds})
	if err != nil {
		return nil, errors.New("[GetBatchBalances] failed to marshal profile ids")
	}
	r, err := client.Post(url, "application/json", bytes.NewBuffer(body))
	if err != nil {
		return nil, errors.New("[GetBatchBalances] client.Post failed to make request to server")
	}
	defer r.Body.Close()

	res := &model.BatchBalancesResponse{}
	if err := json.NewDecoder(r.Body).Decode(res); err != nil {
		return nil, errors.New("[GetBatchBalances] Unable to decode response body")
	}

	return res, nil
}
