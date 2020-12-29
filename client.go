package zoom

import (
	"bytes"
	"encoding/json"
	"errors"
	"github.com/joaosoft/tag"
	"io/ioutil"
	"net/http"
	"net/url"
)

func Init(config *Config, options ...Option) {
	var uri = url.URL{
		Scheme: apiScheme,
		Host:   apiURI,
		Path:   apiVersion,
	}

	defaultClient = &Client{
		endpoint: uri.String(),
		config:   config,
	}

	defaultClient.Reconfigure(options...)
}

func (c *Client) executeRequest(opts *requestOpts) (*http.Response, error) {
	client := c.httpClient()
	req, err := c.addRequestAuth(c.httpRequest(opts))
	if err != nil {
		return nil, err
	}

	req.Header.Add(headerContentType, headerContentTypeApplicationJson)

	return client.Do(req)
}

func (c *Client) httpRequest(opts *requestOpts) (*http.Request, error) {
	var buf bytes.Buffer

	// encode body parameters if any
	if err := json.NewEncoder(&buf).Encode(&opts.DataParameters); err != nil {
		return nil, err
	}

	// set URL parameters
	values := tag.Load(opts.URLParameters, "url")

	// set request URL
	requestURL := c.endpoint + opts.Path
	if len(values) > 0 {
		requestURL += "?" + values.Encode()
	}

	// create HTTP request
	return http.NewRequest(string(opts.Method), requestURL, &buf)
}

func (c *Client) httpClient() *http.Client {
	client := &http.Client{Transport: c.Transport}
	if c.Timeout > 0 {
		client.Timeout = c.Timeout
	}

	return client
}

func (c *Client) request(opts *requestOpts) error {
	if !c.config.Enabled {
		return nil
	}

	// execute HTTP request
	resp, err := c.executeRequest(opts)
	if err != nil {
		return err
	}

	// if there is no body in response
	if opts.HeadResponse {
		if err = c.requestHeadOnly(resp); err != nil {
			return err
		}
	} else {
		if err = c.requestWithBody(opts, resp); err != nil {
			return err
		}
	}

	return nil
}

func (c *Client) requestWithBody(opts *requestOpts, resp *http.Response) error {
	// read HTTP response
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	// check for Zoom errors in the response
	if err := checkError(body); err != nil {
		return err
	}

	// unmarshal the response body into the return object
	return json.Unmarshal(body, &opts.Ret)
}

func (c *Client) requestHeadOnly(resp *http.Response) error {
	if resp.StatusCode >= http.StatusBadRequest {
		return errors.New(resp.Status)
	}

	return nil
}
