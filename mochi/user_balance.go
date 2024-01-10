package mochi

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"strconv"

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
			return nil, errors.New("invalid decoded, error " + err.Error())
		}
		return nil, errors.New("invalid call, code " + strconv.Itoa(resp.StatusCode) + " " + errMsg.Msg)
	}

	var respData model.UserTokenBalanceResponses

	if err := json.Unmarshal(resBody, &respData); err != nil {
		return nil, errors.New("invalid decoded, error " + err.Error())
	}

	return respData.Data, nil
}
