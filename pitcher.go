package pitcher

import "net/http"

// Pitcher represents connection to Versatile API.
type Pitcher interface {
	// Connect allows to make connection to Centrifugo server.
	Signin(email string, password string) error
}

type pitcher struct {
	URL    string
	client *http.Client
}

func (c *pitcher) Signin(email string, password string) error {

	return nil
}

// NewClient .
func NewClient(u string) Pitcher {
	c := &pitcher{
		URL:    u,
		client: &http.Client{},
	}
	return c
}
