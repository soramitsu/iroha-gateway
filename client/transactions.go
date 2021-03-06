// API "Iroha-Gateway Server": transactions Resource Client
//
// Code generated by goagen v1.1.0-dirty, DO NOT EDIT.
//
// Command:
// $ goagen
// --design=github.com/soramitsu/iroha-gateway/design
// --out=$(GOPATH)/src/github.com/soramitsu/iroha-gateway
// --version=v1.1.0-dirty

package client

import (
	"fmt"
	"golang.org/x/net/context"
	"net/http"
	"net/url"
	"strconv"
)

// GetAllTransactionsPath computes a request path to the getAll action of transactions.
func GetAllTransactionsPath(currencyURI string) string {
	param0 := currencyURI

	return fmt.Sprintf("/transactions/%s", param0)
}

// GetAllTransactions makes a request to the getAll action endpoint of the transactions resource
func (c *Client) GetAllTransactions(ctx context.Context, path string, creatorPubkey string, target string, isCommitted *bool) (*http.Response, error) {
	req, err := c.NewGetAllTransactionsRequest(ctx, path, creatorPubkey, target, isCommitted)
	if err != nil {
		return nil, err
	}
	return c.Client.Do(ctx, req)
}

// NewGetAllTransactionsRequest create the request corresponding to the getAll action endpoint of the transactions resource.
func (c *Client) NewGetAllTransactionsRequest(ctx context.Context, path string, creatorPubkey string, target string, isCommitted *bool) (*http.Request, error) {
	scheme := c.Scheme
	if scheme == "" {
		scheme = "http"
	}
	u := url.URL{Host: c.Host, Scheme: scheme, Path: path}
	values := u.Query()
	values.Set("creator_pubkey", creatorPubkey)
	values.Set("target", target)
	if isCommitted != nil {
		tmp25 := strconv.FormatBool(*isCommitted)
		values.Set("is_committed", tmp25)
	}
	u.RawQuery = values.Encode()
	req, err := http.NewRequest("GET", u.String(), nil)
	if err != nil {
		return nil, err
	}
	return req, nil
}
