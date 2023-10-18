package mochipay

import (
	"bytes"
	"crypto/ed25519"
	"encoding/hex"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"strconv"
	"time"
)

func (c *Client) Transfer(req *TransferRequest) ([]TransactionResult, error) {
	requestBody, err := json.Marshal(req)
	if err != nil {
		return nil, err
	}

	var client = &http.Client{}
	request, err := http.NewRequest(http.MethodPost, fmt.Sprintf("%s/api/v1/applications/%v/transfer", c.cfg.BaseURL, c.cfg.ApplicationID), bytes.NewBuffer(requestBody))
	if err != nil {
		return nil, err
	}

	messageHeader := strconv.FormatInt(time.Now().Unix(), 10)
	privateKey, err := hex.DecodeString(c.cfg.APIKey)
	signature := ed25519.Sign(privateKey, []byte(messageHeader))
	request.Header.Add("X-Message", messageHeader)
	request.Header.Add("X-Application", "Guess Game")
	request.Header.Add("X-Signature", hex.EncodeToString(signature))

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
		var errMsg ErrorMessage
		if err := json.Unmarshal(resBody, &errMsg); err != nil {
			return nil, errors.New("invalid decoded, error " + err.Error())
		}
		return nil, errors.New("invalid call, code " + strconv.Itoa(resp.StatusCode) + " " + errMsg.Msg)
	}

	var respData TransactionResponse
	if err := json.Unmarshal(resBody, &respData); err != nil {
		return nil, errors.New("invalid decoded, error " + err.Error())
	}

	return respData.Data, nil
}
