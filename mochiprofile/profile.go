package mochiprofile

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func (c *Client) GetByDiscordID(id string) (*ProfileResponse, error) {
	var client = &http.Client{}
	request, err := http.NewRequest(http.MethodGet, fmt.Sprintf("%v/api/v1/profiles/get-by-discord/%v", c.cfg.BaseURL, id), nil)
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

	res := &ProfileResponse{}
	err = json.Unmarshal(resBody, res)
	if err != nil {
		return nil, err
	}

	return res, nil
}
