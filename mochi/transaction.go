package mochi

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"strings"
	"time"

	"github.com/consolelabs/mochi-go-sdk/mochi/model"
)

func (c *Client) GetTransactions(req model.GetTransactionsRequest) (*model.GetTransactionsResponse, error) {
	var client = &http.Client{
		Timeout: 10 * time.Second,
	}

	queryParams := url.Values{}
	var pageSize int64 = 10
	if req.Size != 0 {
		pageSize = req.Size
	}
	queryParams.Add("size", fmt.Sprintf("%v", pageSize))
	queryParams.Add("page", fmt.Sprintf("%v", req.Page))

	if len(req.ActionList) > 0 {
		actions := ""
		for i, a := range req.ActionList {
			if i == 0 {
				actions += string(a)
				continue
			}
			actions += fmt.Sprintf("|%s", a)
		}
		queryParams.Add("action", actions)
	}
	if req.Type != "" {
		queryParams.Add("type", string(req.Type))
	}
	if req.Status != "" {
		queryParams.Add("status", string(req.Status))
	}

	if req.TokenAddress != "" {
		queryParams.Add("token_address", req.TokenAddress)
	}

	if len(req.Platforms) != 0 {
		platforms := make([]string, 0)
		for _, p := range req.Platforms {
			platforms = append(platforms, string(p))
		}

		queryParams.Add("platforms", strings.Join(platforms, "|"))
	}

	if len(req.ChainIDs) != 0 {
		queryParams.Add("chain_ids", strings.Join(req.ChainIDs, "|"))
	}

	if req.ProfileID != "" {
		queryParams.Add("profile_id", req.ProfileID)
	}

	if req.IsSender != nil {
		queryParams.Add("is_sender", fmt.Sprintf("%v", *req.IsSender))
	}

	if req.SortBy != "" {
		queryParams.Add("sort_by", req.SortBy)
	}

	r, err := client.Get(fmt.Sprintf("%s/api/v1/transactions?", c.cfg.MochiPay.BaseURL) + queryParams.Encode())
	if err != nil {
		return nil, fmt.Errorf("[GetTransactions] Unable to send request %v", err)
	}
	defer r.Body.Close()

	res := &model.GetTransactionsResponse{}
	if err := json.NewDecoder(r.Body).Decode(&res); err != nil {
		return nil, fmt.Errorf("[GetTransactions] Unable to decode response body %v", err)
	}

	return res, nil
}
