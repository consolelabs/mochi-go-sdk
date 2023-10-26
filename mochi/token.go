package mochi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/consolelabs/mochi-go-sdk/mochi/model"
)

func (c *Client) GetTokens(in *model.GetTokenRequest) ([]model.Token, error) {
	var client = &http.Client{}
	request, err := http.NewRequest(http.MethodGet, fmt.Sprintf("%v/api/v1/tokens", c.cfg.MochiPay.BaseURL), nil)
	if err != nil {
		return nil, err
	}

	q := request.URL.Query()
	q.Add("chain_id", in.ChainID)
	q.Add("symbol", in.Symbol)
	request.URL.RawQuery = q.Encode()

	resp, err := client.Do(request)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()
	resBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var res model.TokensResponse
	err = json.Unmarshal(resBody, &res)
	if err != nil {
		return nil, err
	}

	return res.Data, nil
}
