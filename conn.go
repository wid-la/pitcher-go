package pitcher

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
)

type connection interface {
	Close()
	hasToken() bool
	setToken(string)
	Get(string) ([]byte, *ErrorResponse)
	Post(string, interface{}) ([]byte, *ErrorResponse)
	// WriteMessage([]byte) error
	// ReadMessage() ([]byte, error)
}

type hConn struct {
	client       *http.Client
	baseURL      string
	sessionToken string
}

func createNewConnection(url string) connection {
	return &hConn{client: &http.Client{}, baseURL: url}
}

func (c *hConn) Close() {

}

func (c *hConn) hasToken() bool {
	fmt.Println("- ? Has Token : ", c.sessionToken)
	return c.sessionToken != ""
}

func (c *hConn) setToken(token string) {
	c.sessionToken = token
}

func (c *hConn) request(method string, path string, content io.Reader) ([]byte, *ErrorResponse) {
	url, _ := c.buildURL(path)
	request, err := http.NewRequest(method, url.String(), content)
	if err != nil {
		fmt.Println("===> ERROR CODE : ", err.Error())
		return nil, nil
	}

	c.setDefaultHeaders(request)
	response, err := c.client.Do(request)
	if err != nil {
		fmt.Println("===> ERROR CODE : ", err.Error())
		return nil, nil
	}

	if response.StatusCode >= 400 && response.StatusCode < 600 {
		return nil, c.processErrorMessage(response.Body)
	}

	return c.streamToByte(response.Body), nil
}

func (c *hConn) Get(path string) ([]byte, *ErrorResponse) {
	return c.request("GET", path, nil)
}

func (c *hConn) Post(path string, obj interface{}) ([]byte, *ErrorResponse) {
	payloadJSON, _ := json.Marshal(obj)
	return c.request("POST", path, bytes.NewBuffer(payloadJSON))
}

func (c *hConn) buildURL(pathOrURL string) (*url.URL, error) {
	u, err := url.ParseRequestURI(pathOrURL)
	if err != nil {
		u, err = url.Parse(c.baseURL)
		if err != nil {
			return nil, err
		}
		return u.Parse(pathOrURL)
	}
	return u, nil
}

func (c *hConn) setDefaultHeaders(request *http.Request) {
	request.Header.Set("Accept", "application/json")
	request.Header.Set("Content-Type", "application/json")
	request.Header.Set("Auth-Server-Token", c.sessionToken)
	// if c.hasToken() {
	// 	cookie := &http.Cookie{Name: "auth_token", Value: c.sessionToken}
	// 	request.AddCookie(cookie)
	// }
}

func (c *hConn) processErrorMessage(body io.ReadCloser) *ErrorResponse {
	res := ErrorResponse{}
	// readio := bytes.NewReader(body)
	err := json.NewDecoder(body).Decode(&res)
	if err != nil {
		fmt.Printf("===!!!!==== ERROR Read Error : %s \n", err.Error())
	}
	fmt.Println("===> MESSAGE : ", res.Description)
	return &res
}

func (c *hConn) streamToByte(stream io.Reader) []byte {
	buf := new(bytes.Buffer)
	buf.ReadFrom(stream)
	return buf.Bytes()
}
