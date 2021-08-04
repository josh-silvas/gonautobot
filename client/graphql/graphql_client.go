package graphql

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// New creates a new graphql API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

/*
Client for graphql API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientOption is the option for Client methods
type ClientOption func(*runtime.ClientOperation)

// ClientService is the interface for Client methods
type ClientService interface {
	GraphqlCreate(params *GraphqlCreateParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GraphqlCreateOK, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
  GraphqlCreate Query the database using a GraphQL query
*/
func (a *Client) GraphqlCreate(params *GraphqlCreateParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GraphqlCreateOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGraphqlCreateParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "graphql_create",
		Method:             "POST",
		PathPattern:        "/graphql/",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GraphqlCreateReader{formats: a.formats},
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
	success, ok := result.(*GraphqlCreateOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for graphql_create: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
