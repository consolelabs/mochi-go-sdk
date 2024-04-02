package mochi

import (
	"bytes"
	"crypto/ed25519"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strconv"
	"time"

	"github.com/consolelabs/mochi-go-sdk/mochi/model"
)

func (c *Client) Transfer(req *model.TransferRequest) ([]model.TransferTransaction, error) {
	requestBody, err := json.Marshal(req)
	if err != nil {
		return nil, err
	}

	var client = &http.Client{}
	request, err := http.NewRequest(http.MethodPost, fmt.Sprintf("%s/api/v1/applications/%v/transfer", c.cfg.MochiPay.BaseURL, c.cfg.MochiPay.ApplicationID), bytes.NewBuffer(requestBody))
	if err != nil {
		return nil, err
	}

	messageHeader := strconv.FormatInt(time.Now().Unix(), 10)
	privateKey, err := hex.DecodeString(c.cfg.MochiPay.APIKey)
	if err != nil {
		return nil, fmt.Errorf("[Transfer] invalid API Key, error %v", err)
	}

	signature := ed25519.Sign(privateKey, []byte(messageHeader))
	request.Header.Add("X-Message", messageHeader)
	request.Header.Add("X-Application", c.cfg.MochiPay.ApplicationName)
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
		var errMsg model.ErrorMessage
		if err := json.Unmarshal(resBody, &errMsg); err != nil {
			return nil, fmt.Errorf("[Transfer] failed to decode error, error %v", err.Error())
		}
		return nil, fmt.Errorf("[Transfer] failed to call , error %v", err.Error())
	}

	var respData model.TransferTransactionResponse
	if err := json.Unmarshal(resBody, &respData); err != nil {
		return nil, fmt.Errorf("[Transfer] failed to decode, error %v", err.Error())
	}

	return respData.Data, nil
}
