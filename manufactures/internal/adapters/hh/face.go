package hh

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"time"
)

const (
	baseURL      = "https://api.hh.ru"
	authURL      = "https://hh.ru/oauth/token"
	clientID     = "YOUR_CLIENT_ID"     // Замените на ваш Client ID
	clientSecret = "YOUR_CLIENT_SECRET" // Замените на ваш Client Secret
	userAgent    = "PostmanRuntime/7.43.4"
)

// HHClient представляет клиент для работы с API hh.ru
type HHClient struct {
	httpClient *http.Client
	token      string
}

// NewHHClient создает новый экземпляр клиента для API hh.ru
func NewHHClient() HHClient {

	return HHClient{
		httpClient: &http.Client{
			Timeout: time.Second * 15,
		},
	}
}

// AuthResponse представляет ответ сервера на запрос авторизации
type AuthResponse struct {
	AccessToken string `json:"access_token"`
	TokenType   string `json:"token_type"`
	ExpiresIn   int    `json:"expires_in"`
}

// Authenticate выполняет OAuth2 аутентификацию в API hh.ru
func (c *HHClient) Authenticate() error {
	data := url.Values{}
	data.Set("grant_type", "client_credentials")
	data.Set("client_id", clientID)
	data.Set("client_secret", clientSecret)

	req, err := http.NewRequest("POST", authURL, bytes.NewBufferString(data.Encode()))
	if err != nil {
		return fmt.Errorf("error creating auth request: %v", err)
	}

	req.Header.Set("User-Agent", userAgent)
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return fmt.Errorf("error making auth request: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		body, _ := ioutil.ReadAll(resp.Body)
		return fmt.Errorf("auth failed with status %d: %s", resp.StatusCode, string(body))
	}

	var authResp AuthResponse
	if err := json.NewDecoder(resp.Body).Decode(&authResp); err != nil {
		return fmt.Errorf("error decoding auth response: %v", err)
	}

	c.token = authResp.AccessToken
	return nil
}
