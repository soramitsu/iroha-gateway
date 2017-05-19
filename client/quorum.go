// Code generated by goagen v1.2.0, DO NOT EDIT.
//
// API "Iroha-Gateway Server": quorum Resource Client
//
// Command:
// $ goagen
// --design=github.com/soramitsu/iroha-gateway/design
// --out=$(GOPATH)/src/github.com/soramitsu/iroha-gateway
// --version=v1.2.0-dirty

package client

import (
	"bytes"
	"context"
	"fmt"
	"net/http"
	"net/url"
)

// UpdateQuorumPath computes a request path to the update action of quorum.
func UpdateQuorumPath(uuid string) string {
	param0 := uuid

	return fmt.Sprintf("/accounts/%s/quorum", param0)
}

// UpdateQuorum makes a request to the update action endpoint of the quorum resource
func (c *Client) UpdateQuorum(ctx context.Context, path string, payload *UpdateQuorumRequest, contentType string) (*http.Response, error) {
	req, err := c.NewUpdateQuorumRequest(ctx, path, payload, contentType)
	if err != nil {
		return nil, err
	}
	return c.Client.Do(ctx, req)
}

// NewUpdateQuorumRequest create the request corresponding to the update action endpoint of the quorum resource.
func (c *Client) NewUpdateQuorumRequest(ctx context.Context, path string, payload *UpdateQuorumRequest, contentType string) (*http.Request, error) {
	var body bytes.Buffer
	if contentType == "" {
		contentType = "*/*" // Use default encoder
	}
	err := c.Encoder.Encode(payload, &body, contentType)
	if err != nil {
		return nil, fmt.Errorf("failed to encode body: %s", err)
	}
	scheme := c.Scheme
	if scheme == "" {
		scheme = "http"
	}
	u := url.URL{Host: c.Host, Scheme: scheme, Path: path}
	req, err := http.NewRequest("PUT", u.String(), &body)
	if err != nil {
		return nil, err
	}
	header := req.Header
	if contentType == "*/*" {
		header.Set("Content-Type", "application/json")
	} else {
		header.Set("Content-Type", contentType)
	}
	return req, nil
}
