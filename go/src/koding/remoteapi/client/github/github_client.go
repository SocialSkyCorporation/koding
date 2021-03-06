package github

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// New creates a new github API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) *Client {
	return &Client{transport: transport, formats: formats}
}

/*
Client for github API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

/*
GithubAPI github api API
*/
func (a *Client) GithubAPI(params *GithubAPIParams, authInfo runtime.ClientAuthInfoWriter) (*GithubAPIOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGithubAPIParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "Github.api",
		Method:             "POST",
		PathPattern:        "/remote.api/Github.api",
		ProducesMediaTypes: []string{""},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &GithubAPIReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*GithubAPIOK), nil

}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
