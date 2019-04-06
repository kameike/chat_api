// Code generated by go-swagger; DO NOT EDIT.

package account

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// New creates a new account API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) *Client {
	return &Client{transport: transport, formats: formats}
}

/*
Client for account API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

/*
PostAuth ログインs

サインアップもしくはアクセストークンの更新を行います
*/
func (a *Client) PostAuth(params *PostAuthParams) (*PostAuthOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPostAuthParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "PostAuth",
		Method:             "POST",
		PathPattern:        "/auth",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &PostAuthReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*PostAuthOK), nil

}

/*
PostProfile ユーザープロファイルのアップデートs

nameをアップデートできます。
*/
func (a *Client) PostProfile(params *PostProfileParams, authInfo runtime.ClientAuthInfoWriter) (*PostProfileOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPostProfileParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "PostProfile",
		Method:             "POST",
		PathPattern:        "/profile",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &PostProfileReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*PostProfileOK), nil

}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
