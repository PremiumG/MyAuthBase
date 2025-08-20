package utils

import (
	"fmt"
	"io"
	"net/http"
	"strings"
)

func SendEmail(email string, magiclink string) {
	url := "https://sandbox.api.mailtrap.io/api/send/3918002"
	method := "POST"

	payload := strings.NewReader(fmt.Sprintf(`{
  "from": {
    "email": "hello@example.com",
    "name": "Mailtrap Test"
  },
  "to": [
    {
      "email": "%s"
    }
  ],
  "subject": "You are awesome!",
  "text": "%s",
  "category": "Integration Test"
}`, email, magiclink))
	client := &http.Client{}
	req, err := http.NewRequest(method, url, payload)

	if err != nil {
		fmt.Println(err)
		return
	}
	req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", AppConfig.EmailerKey))
	req.Header.Add("Content-Type", "application/json")

	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(string(body))
}
