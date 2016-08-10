package pitcher

import (
	"bytes"
	"encoding/json"
	"fmt"
)

// Pitcher represents connection to Versatile API.
type Pitcher interface {
	// Connect allows to make connection to Centrifugo server.
	GetURL() string
	Signin(email string, password string) error
	CurrentSession() *Session
	GetData() []Data
	PublishData(Data, interface{})
}

type pitcher struct {
	URL  string
	conn connection
}

func (c *pitcher) Signin(email string, password string) error {
	req := &SigninRequest{Email: email, Password: password}
	o, err := c.conn.Post("users/signin", req)
	if err != nil {
		return nil
	}
	res := &Session{}
	r := bytes.NewReader(o)
	json.NewDecoder(r).Decode(&res)
	c.conn.setToken(res.Token)
	return nil
}

func (c *pitcher) CurrentSession() *Session {
	o, err := c.conn.Get("users/me/session")
	if err != nil {
		return nil
	}
	res := &Session{}
	r := bytes.NewReader(o)
	json.NewDecoder(r).Decode(&res)
	return res
}

func (c *pitcher) GetURL() string {
	return c.URL
}

func (c *pitcher) GetData() []Data {
	o, err := c.conn.Get("data")
	if err != nil {
		return nil
	}
	res := []Data{}
	r := bytes.NewReader(o)
	json.NewDecoder(r).Decode(&res)
	return res
}

func (c *pitcher) PublishData(d Data, o interface{}) {
	// data/{key}/publish
	req := &DataPublishRequest{}
	req.Value = o
	path := fmt.Sprintf("data/%s/publish", d.ID)
	_, err := c.conn.Post(path, req)
	if err != nil {
		fmt.Println("Error Publish")
	}
}

// NewClient .
func NewClient(u string) Pitcher {
	co := createNewConnection(u)
	c := &pitcher{
		URL:  u,
		conn: co,
	}
	return c
}
