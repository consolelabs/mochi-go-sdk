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

func (c *Client) RequestPayment(req *PaymentRequest) error {
	requestBody, err := json.Marshal(req)
	if err != nil {
		return err
	}

	var client = &http.Client{}
	requestURL := fmt.Sprintf("%s/api/v1/applications/%v/requests", c.cfg.BaseURL, c.cfg.ApplicationID)
	request, err := http.NewRequest(http.MethodPost, requestURL, bytes.NewBuffer(requestBody))
	if err != nil {
		return err
	}

	messageHeader := strconv.FormatInt(time.Now().Unix(), 10)
	privateKey, err := hex.DecodeString(c.cfg.APIKey)
	signature := ed25519.Sign(privateKey, []byte(messageHeader))

	request.Header.Add("X-Message", messageHeader)
	request.Header.Add("X-Application", c.cfg.ApplicationName)
	request.Header.Add("X-Signature", hex.EncodeToString(signature))

	resp, err := client.Do(request)
	if err != nil {
		return err
	}

	defer resp.Body.Close()
	resBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	if resp.StatusCode != http.StatusOK {
		var errMsg ErrorMessage
		if err := json.Unmarshal(resBody, &errMsg); err != nil {
			return errors.New("invalid decoded, error " + err.Error())
		}
		return errors.New("invalid call, code " + strconv.Itoa(resp.StatusCode) + " " + errMsg.Msg)
	}

	return nil
}
