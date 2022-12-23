// Code generated by go-swagger; DO NOT EDIT.

package equinixmetal

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// New creates a new equinixmetal API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

/*
Client for equinixmetal API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientOption is the option for Client methods
type ClientOption func(*runtime.ClientOperation)

// ClientService is the interface for Client methods
type ClientService interface {
	ListEquinixMetalSizesNoCredentialsV2(params *ListEquinixMetalSizesNoCredentialsV2Params, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*ListEquinixMetalSizesNoCredentialsV2OK, error)

	ListProjectEquinixMetalSizes(params *ListProjectEquinixMetalSizesParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*ListProjectEquinixMetalSizesOK, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
ListEquinixMetalSizesNoCredentialsV2 lists sizes from equinix metal
*/
func (a *Client) ListEquinixMetalSizesNoCredentialsV2(params *ListEquinixMetalSizesNoCredentialsV2Params, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*ListEquinixMetalSizesNoCredentialsV2OK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewListEquinixMetalSizesNoCredentialsV2Params()
	}
	op := &runtime.ClientOperation{
		ID:                 "listEquinixMetalSizesNoCredentialsV2",
		Method:             "GET",
		PathPattern:        "/api/v2/projects/{project_id}/clusters/{cluster_id}/providers/equinixmetal/sizes",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &ListEquinixMetalSizesNoCredentialsV2Reader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*ListEquinixMetalSizesNoCredentialsV2OK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*ListEquinixMetalSizesNoCredentialsV2Default)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
ListProjectEquinixMetalSizes lists sizes from equinix metal
*/
func (a *Client) ListProjectEquinixMetalSizes(params *ListProjectEquinixMetalSizesParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*ListProjectEquinixMetalSizesOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewListProjectEquinixMetalSizesParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "listProjectEquinixMetalSizes",
		Method:             "GET",
		PathPattern:        "/api/v2/projects/{project_id}/providers/equinixmetal/sizes",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &ListProjectEquinixMetalSizesReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*ListProjectEquinixMetalSizesOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*ListProjectEquinixMetalSizesDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}