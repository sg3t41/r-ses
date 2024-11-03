package oauth

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func GetAccessToken(code, clientID, clientSecret, url string) (string, error) {
	req, err := http.NewRequest("POST", url, nil)
	if err != nil {
		return "", err
	}

	req.Header.Set("Accept", "application/json")

	q := req.URL.Query()
	q.Add("client_id", clientID)
	q.Add("client_secret", clientSecret)
	q.Add("code", code)
	req.URL.RawQuery = q.Encode()

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return "", fmt.Errorf("failed to get access token: %s", resp.Status)
	}

	var result struct {
		AccessToken string `json:"access_token"`
	}
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return "", err
	}

	return result.AccessToken, nil
}
