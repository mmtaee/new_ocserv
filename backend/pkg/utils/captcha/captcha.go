package captcha

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
	"net/url"
	"time"
)

func Verify(secret, token string) bool {
	type captchaResponse struct {
		Success     bool      `json:"success"`
		ChallengeTS time.Time `json:"challenge_ts"`
		Hostname    string    `json:"hostname"`
		ErrorCodes  []string  `json:"error-codes"`
	}

	formData := url.Values{}
	formData.Set("secret", secret)
	formData.Set("response", token)

	client := &http.Client{Timeout: 10 * time.Second}
	resp, err := client.PostForm("https://www.google.com/recaptcha/api/siteverify", formData)
	if err != nil {
		return false
	}
	defer func() {
		if err = resp.Body.Close(); err != nil {
			log.Printf("Error closing response body: %v", err)
		}
	}()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return false
	}

	var response captchaResponse
	if err = json.Unmarshal(body, &response); err != nil {
		return false
	}
	return response.Success
}
