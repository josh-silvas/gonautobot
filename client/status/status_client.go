package status

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// New creates a new status API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

/*
Client for status API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientOption is the option for Client methods
type ClientOption func(*runtime.ClientOperation)

// ClientService is the interface for Client methods
type ClientService interface {
	StatusList(params *StatusListParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*StatusListOK, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
  StatusList A lightweight read-only endpoint for conveying Nautobot's current operational status.
*/
func (a *Client) StatusList(params *StatusListParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*StatusListOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewStatusListParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "status_list",
		Method:             "GET",
		PathPattern:        "/status/",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &StatusListReader{formats: a.formats},
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
	success, ok := result.(*StatusListOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for status_list: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
