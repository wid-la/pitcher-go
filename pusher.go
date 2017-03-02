package pitcher

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

// Pusher represents connection to Versatile PUSH API.
type Pusher interface {
	Publish(channel string, payload Payload) error
}

type pusher struct {
	URL    string
	token  string
	appID  string
	client *http.Client
}

// Payload .
type Payload struct {
	Title    string `json:"title,omitempty"`
	Subtitle string `json:"subtitle,omitempty"`
	Body     string `json:"body,omitempty"`
}

func (p *pusher) Publish(channel string, payload Payload) error {
	payloadJSON, _ := json.Marshal(payload)
	url := fmt.Sprintf("%s/%s/%s", p.URL, p.appID, channel)

	request, err := http.NewRequest("POST", url, bytes.NewBuffer(payloadJSON))
	if err != nil {
		fmt.Println("===> ERROR CODE : ", err.Error())
		return err
	}

	request.Header.Set("Accept", "application/json")
	request.Header.Set("Content-Type", "application/json")
	request.Header.Set("Auth-Server-Token", p.token)

	response, err := p.client.Do(request)
	if err != nil {
		fmt.Println("===> ERROR CODE : ", err.Error())
		return err
	}

	fmt.Println(response)

	return nil
}

// NewPusher .
func NewPusher(url string, token string, appID string) Pusher {
	p := &pusher{
		URL:    url,
		token:  token,
		appID:  appID,
		client: &http.Client{},
	}
	return p
}
