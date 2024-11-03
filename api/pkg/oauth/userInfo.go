package oauth

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func GetUserInfo[T any](accessToken, url, acceptHeader string) (*T, error) {
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}

	req.Header.Set("Authorization", "Bearer "+accessToken)
	req.Header.Set("Accept", acceptHeader)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("failed to get user info: %s", resp.Status)
	}

	var t T
	if err := json.NewDecoder(resp.Body).Decode(&t); err != nil {
		return nil, err
	}

	return &t, nil
}
