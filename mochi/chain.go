package mochi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/consolelabs/mochi-go-sdk/mochi/model"
)

func (c *Client) GetChains() ([]model.Chain, error) {
	var client = &http.Client{}
	request, err := http.NewRequest(http.MethodGet, fmt.Sprintf("%v/api/v1/chains", c.cfg.MochiPay.BaseURL), nil)
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

	var res model.ChainsResponse
	err = json.Unmarshal(resBody, &res)
	if err != nil {
		return nil, err
	}

	return res.Data, nil
}
