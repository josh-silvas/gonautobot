package core

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/sirupsen/logrus"
	"io"
	"net/http"
	"net/url"
	"strings"
)

// Client : Requests data type client.
type Client struct {
	Client *http.Client
	Log    *logrus.Logger
	Token  string
	URL    string
}

// Request : Crafts an HTTP request to Nautobot.
func (c *Client) Request(method, uri string, body interface{}, params *url.Values) (*http.Request, error) {
	var buf io.ReadWriter
	if body != nil {
		buf = new(bytes.Buffer)
		enc := json.NewEncoder(buf)
		enc.SetEscapeHTML(false)
		if err := enc.Encode(body); err != nil {
			return nil, err
		}
	}
	uri = strings.TrimPrefix(uri, "/")

	req, err := http.NewRequestWithContext(context.Background(), method, fmt.Sprintf("%s%s", c.URL, uri), buf)
	if err != nil {
		return nil, err
	}

	if body != nil {
		req.Header.Set("Content-Type", "application/json")
	}
	req.Header.Set("Accept", "application/json")
	req.Header.Set("Authorization", "Token "+c.Token)
	if params != nil && len(*params) > 0 {
		req.URL.RawQuery = params.Encode()
	}

	return req, nil
}

// Do : Performs an HTTP requests and returns the raw *http.Response.
func (c *Client) Do(req *http.Request) (*http.Response, error) {
	resp, err := c.Client.Do(req)
	if err != nil {
		return nil, err
	}
	return resp, HasError(resp)
}

// UnmarshalDo : Performs an HTTP request to the resource defined in the request.
func (c *Client) UnmarshalDo(req *http.Request, v interface{}) error {
	resp, err := c.Do(req)
	if resp != nil && resp.Body != nil {
		defer func(r *http.Response) {
			if err := r.Body.Close(); err != nil {
				c.Log.Warn(err)
			}
		}(resp)
	}
	if err != nil {
		return err
	}

	switch v := v.(type) {
	case nil:
	case io.Writer:
		_, err = io.Copy(v, resp.Body)
	default:
		decErr := json.NewDecoder(resp.Body).Decode(v)
		if errors.Is(decErr, io.EOF) {
			decErr = nil // ignore EOF errors caused by empty response body
		}
		if decErr != nil {
			err = fmt.Errorf("error.UnmarshalDo.json.NewDecoder(%w)", decErr)
		}
	}
	return err
}
