package circuits

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// New creates a new circuits API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

/*
Client for circuits API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientOption is the option for Client methods
type ClientOption func(*runtime.ClientOperation)

// ClientService is the interface for Client methods
type ClientService interface {
	CircuitsCircuitTerminationsBulkDelete(params *CircuitsCircuitTerminationsBulkDeleteParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*CircuitsCircuitTerminationsBulkDeleteNoContent, error)

	CircuitsCircuitTerminationsBulkPartialUpdate(params *CircuitsCircuitTerminationsBulkPartialUpdateParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*CircuitsCircuitTerminationsBulkPartialUpdateOK, error)

	CircuitsCircuitTerminationsBulkUpdate(params *CircuitsCircuitTerminationsBulkUpdateParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*CircuitsCircuitTerminationsBulkUpdateOK, error)

	CircuitsCircuitTerminationsCreate(params *CircuitsCircuitTerminationsCreateParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*CircuitsCircuitTerminationsCreateCreated, error)

	CircuitsCircuitTerminationsDelete(params *CircuitsCircuitTerminationsDeleteParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*CircuitsCircuitTerminationsDeleteNoContent, error)

	CircuitsCircuitTerminationsList(params *CircuitsCircuitTerminationsListParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*CircuitsCircuitTerminationsListOK, error)

	CircuitsCircuitTerminationsPartialUpdate(params *CircuitsCircuitTerminationsPartialUpdateParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*CircuitsCircuitTerminationsPartialUpdateOK, error)

	CircuitsCircuitTerminationsRead(params *CircuitsCircuitTerminationsReadParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*CircuitsCircuitTerminationsReadOK, error)

	CircuitsCircuitTerminationsTrace(params *CircuitsCircuitTerminationsTraceParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*CircuitsCircuitTerminationsTraceOK, error)

	CircuitsCircuitTerminationsUpdate(params *CircuitsCircuitTerminationsUpdateParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*CircuitsCircuitTerminationsUpdateOK, error)

	CircuitsCircuitTypesBulkDelete(params *CircuitsCircuitTypesBulkDeleteParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*CircuitsCircuitTypesBulkDeleteNoContent, error)

	CircuitsCircuitTypesBulkPartialUpdate(params *CircuitsCircuitTypesBulkPartialUpdateParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*CircuitsCircuitTypesBulkPartialUpdateOK, error)

	CircuitsCircuitTypesBulkUpdate(params *CircuitsCircuitTypesBulkUpdateParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*CircuitsCircuitTypesBulkUpdateOK, error)

	CircuitsCircuitTypesCreate(params *CircuitsCircuitTypesCreateParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*CircuitsCircuitTypesCreateCreated, error)

	CircuitsCircuitTypesDelete(params *CircuitsCircuitTypesDeleteParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*CircuitsCircuitTypesDeleteNoContent, error)

	CircuitsCircuitTypesList(params *CircuitsCircuitTypesListParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*CircuitsCircuitTypesListOK, error)

	CircuitsCircuitTypesPartialUpdate(params *CircuitsCircuitTypesPartialUpdateParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*CircuitsCircuitTypesPartialUpdateOK, error)

	CircuitsCircuitTypesRead(params *CircuitsCircuitTypesReadParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*CircuitsCircuitTypesReadOK, error)

	CircuitsCircuitTypesUpdate(params *CircuitsCircuitTypesUpdateParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*CircuitsCircuitTypesUpdateOK, error)

	CircuitsCircuitsBulkDelete(params *CircuitsCircuitsBulkDeleteParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*CircuitsCircuitsBulkDeleteNoContent, error)

	CircuitsCircuitsBulkPartialUpdate(params *CircuitsCircuitsBulkPartialUpdateParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*CircuitsCircuitsBulkPartialUpdateOK, error)

	CircuitsCircuitsBulkUpdate(params *CircuitsCircuitsBulkUpdateParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*CircuitsCircuitsBulkUpdateOK, error)

	CircuitsCircuitsCreate(params *CircuitsCircuitsCreateParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*CircuitsCircuitsCreateCreated, error)

	CircuitsCircuitsDelete(params *CircuitsCircuitsDeleteParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*CircuitsCircuitsDeleteNoContent, error)

	CircuitsCircuitsList(params *CircuitsCircuitsListParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*CircuitsCircuitsListOK, error)

	CircuitsCircuitsPartialUpdate(params *CircuitsCircuitsPartialUpdateParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*CircuitsCircuitsPartialUpdateOK, error)

	CircuitsCircuitsRead(params *CircuitsCircuitsReadParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*CircuitsCircuitsReadOK, error)

	CircuitsCircuitsUpdate(params *CircuitsCircuitsUpdateParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*CircuitsCircuitsUpdateOK, error)

	CircuitsProvidersBulkDelete(params *CircuitsProvidersBulkDeleteParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*CircuitsProvidersBulkDeleteNoContent, error)

	CircuitsProvidersBulkPartialUpdate(params *CircuitsProvidersBulkPartialUpdateParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*CircuitsProvidersBulkPartialUpdateOK, error)

	CircuitsProvidersBulkUpdate(params *CircuitsProvidersBulkUpdateParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*CircuitsProvidersBulkUpdateOK, error)

	CircuitsProvidersCreate(params *CircuitsProvidersCreateParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*CircuitsProvidersCreateCreated, error)

	CircuitsProvidersDelete(params *CircuitsProvidersDeleteParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*CircuitsProvidersDeleteNoContent, error)

	CircuitsProvidersList(params *CircuitsProvidersListParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*CircuitsProvidersListOK, error)

	CircuitsProvidersPartialUpdate(params *CircuitsProvidersPartialUpdateParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*CircuitsProvidersPartialUpdateOK, error)

	CircuitsProvidersRead(params *CircuitsProvidersReadParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*CircuitsProvidersReadOK, error)

	CircuitsProvidersUpdate(params *CircuitsProvidersUpdateParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*CircuitsProvidersUpdateOK, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
  CircuitsCircuitTerminationsBulkDelete circuits circuit terminations bulk delete API
*/
func (a *Client) CircuitsCircuitTerminationsBulkDelete(params *CircuitsCircuitTerminationsBulkDeleteParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*CircuitsCircuitTerminationsBulkDeleteNoContent, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewCircuitsCircuitTerminationsBulkDeleteParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "circuits_circuit-terminations_bulk_delete",
		Method:             "DELETE",
		PathPattern:        "/circuits/circuit-terminations/",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &CircuitsCircuitTerminationsBulkDeleteReader{formats: a.formats},
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
	success, ok := result.(*CircuitsCircuitTerminationsBulkDeleteNoContent)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for circuits_circuit-terminations_bulk_delete: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  CircuitsCircuitTerminationsBulkPartialUpdate circuits circuit terminations bulk partial update API
*/
func (a *Client) CircuitsCircuitTerminationsBulkPartialUpdate(params *CircuitsCircuitTerminationsBulkPartialUpdateParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*CircuitsCircuitTerminationsBulkPartialUpdateOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewCircuitsCircuitTerminationsBulkPartialUpdateParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "circuits_circuit-terminations_bulk_partial_update",
		Method:             "PATCH",
		PathPattern:        "/circuits/circuit-terminations/",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &CircuitsCircuitTerminationsBulkPartialUpdateReader{formats: a.formats},
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
	success, ok := result.(*CircuitsCircuitTerminationsBulkPartialUpdateOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for circuits_circuit-terminations_bulk_partial_update: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  CircuitsCircuitTerminationsBulkUpdate circuits circuit terminations bulk update API
*/
func (a *Client) CircuitsCircuitTerminationsBulkUpdate(params *CircuitsCircuitTerminationsBulkUpdateParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*CircuitsCircuitTerminationsBulkUpdateOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewCircuitsCircuitTerminationsBulkUpdateParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "circuits_circuit-terminations_bulk_update",
		Method:             "PUT",
		PathPattern:        "/circuits/circuit-terminations/",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &CircuitsCircuitTerminationsBulkUpdateReader{formats: a.formats},
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
	success, ok := result.(*CircuitsCircuitTerminationsBulkUpdateOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for circuits_circuit-terminations_bulk_update: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  CircuitsCircuitTerminationsCreate circuits circuit terminations create API
*/
func (a *Client) CircuitsCircuitTerminationsCreate(params *CircuitsCircuitTerminationsCreateParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*CircuitsCircuitTerminationsCreateCreated, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewCircuitsCircuitTerminationsCreateParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "circuits_circuit-terminations_create",
		Method:             "POST",
		PathPattern:        "/circuits/circuit-terminations/",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &CircuitsCircuitTerminationsCreateReader{formats: a.formats},
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
	success, ok := result.(*CircuitsCircuitTerminationsCreateCreated)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for circuits_circuit-terminations_create: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  CircuitsCircuitTerminationsDelete circuits circuit terminations delete API
*/
func (a *Client) CircuitsCircuitTerminationsDelete(params *CircuitsCircuitTerminationsDeleteParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*CircuitsCircuitTerminationsDeleteNoContent, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewCircuitsCircuitTerminationsDeleteParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "circuits_circuit-terminations_delete",
		Method:             "DELETE",
		PathPattern:        "/circuits/circuit-terminations/{id}/",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &CircuitsCircuitTerminationsDeleteReader{formats: a.formats},
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
	success, ok := result.(*CircuitsCircuitTerminationsDeleteNoContent)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for circuits_circuit-terminations_delete: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  CircuitsCircuitTerminationsList circuits circuit terminations list API
*/
func (a *Client) CircuitsCircuitTerminationsList(params *CircuitsCircuitTerminationsListParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*CircuitsCircuitTerminationsListOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewCircuitsCircuitTerminationsListParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "circuits_circuit-terminations_list",
		Method:             "GET",
		PathPattern:        "/circuits/circuit-terminations/",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &CircuitsCircuitTerminationsListReader{formats: a.formats},
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
	success, ok := result.(*CircuitsCircuitTerminationsListOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for circuits_circuit-terminations_list: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  CircuitsCircuitTerminationsPartialUpdate circuits circuit terminations partial update API
*/
func (a *Client) CircuitsCircuitTerminationsPartialUpdate(params *CircuitsCircuitTerminationsPartialUpdateParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*CircuitsCircuitTerminationsPartialUpdateOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewCircuitsCircuitTerminationsPartialUpdateParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "circuits_circuit-terminations_partial_update",
		Method:             "PATCH",
		PathPattern:        "/circuits/circuit-terminations/{id}/",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &CircuitsCircuitTerminationsPartialUpdateReader{formats: a.formats},
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
	success, ok := result.(*CircuitsCircuitTerminationsPartialUpdateOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for circuits_circuit-terminations_partial_update: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  CircuitsCircuitTerminationsRead circuits circuit terminations read API
*/
func (a *Client) CircuitsCircuitTerminationsRead(params *CircuitsCircuitTerminationsReadParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*CircuitsCircuitTerminationsReadOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewCircuitsCircuitTerminationsReadParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "circuits_circuit-terminations_read",
		Method:             "GET",
		PathPattern:        "/circuits/circuit-terminations/{id}/",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &CircuitsCircuitTerminationsReadReader{formats: a.formats},
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
	success, ok := result.(*CircuitsCircuitTerminationsReadOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for circuits_circuit-terminations_read: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  CircuitsCircuitTerminationsTrace Trace a complete cable path and return each segment as a three-tuple of (termination, cable, termination).
*/
func (a *Client) CircuitsCircuitTerminationsTrace(params *CircuitsCircuitTerminationsTraceParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*CircuitsCircuitTerminationsTraceOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewCircuitsCircuitTerminationsTraceParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "circuits_circuit-terminations_trace",
		Method:             "GET",
		PathPattern:        "/circuits/circuit-terminations/{id}/trace/",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &CircuitsCircuitTerminationsTraceReader{formats: a.formats},
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
	success, ok := result.(*CircuitsCircuitTerminationsTraceOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for circuits_circuit-terminations_trace: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  CircuitsCircuitTerminationsUpdate circuits circuit terminations update API
*/
func (a *Client) CircuitsCircuitTerminationsUpdate(params *CircuitsCircuitTerminationsUpdateParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*CircuitsCircuitTerminationsUpdateOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewCircuitsCircuitTerminationsUpdateParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "circuits_circuit-terminations_update",
		Method:             "PUT",
		PathPattern:        "/circuits/circuit-terminations/{id}/",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &CircuitsCircuitTerminationsUpdateReader{formats: a.formats},
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
	success, ok := result.(*CircuitsCircuitTerminationsUpdateOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for circuits_circuit-terminations_update: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  CircuitsCircuitTypesBulkDelete circuits circuit types bulk delete API
*/
func (a *Client) CircuitsCircuitTypesBulkDelete(params *CircuitsCircuitTypesBulkDeleteParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*CircuitsCircuitTypesBulkDeleteNoContent, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewCircuitsCircuitTypesBulkDeleteParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "circuits_circuit-types_bulk_delete",
		Method:             "DELETE",
		PathPattern:        "/circuits/circuit-types/",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &CircuitsCircuitTypesBulkDeleteReader{formats: a.formats},
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
	success, ok := result.(*CircuitsCircuitTypesBulkDeleteNoContent)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for circuits_circuit-types_bulk_delete: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  CircuitsCircuitTypesBulkPartialUpdate circuits circuit types bulk partial update API
*/
func (a *Client) CircuitsCircuitTypesBulkPartialUpdate(params *CircuitsCircuitTypesBulkPartialUpdateParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*CircuitsCircuitTypesBulkPartialUpdateOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewCircuitsCircuitTypesBulkPartialUpdateParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "circuits_circuit-types_bulk_partial_update",
		Method:             "PATCH",
		PathPattern:        "/circuits/circuit-types/",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &CircuitsCircuitTypesBulkPartialUpdateReader{formats: a.formats},
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
	success, ok := result.(*CircuitsCircuitTypesBulkPartialUpdateOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for circuits_circuit-types_bulk_partial_update: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  CircuitsCircuitTypesBulkUpdate circuits circuit types bulk update API
*/
func (a *Client) CircuitsCircuitTypesBulkUpdate(params *CircuitsCircuitTypesBulkUpdateParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*CircuitsCircuitTypesBulkUpdateOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewCircuitsCircuitTypesBulkUpdateParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "circuits_circuit-types_bulk_update",
		Method:             "PUT",
		PathPattern:        "/circuits/circuit-types/",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &CircuitsCircuitTypesBulkUpdateReader{formats: a.formats},
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
	success, ok := result.(*CircuitsCircuitTypesBulkUpdateOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for circuits_circuit-types_bulk_update: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  CircuitsCircuitTypesCreate circuits circuit types create API
*/
func (a *Client) CircuitsCircuitTypesCreate(params *CircuitsCircuitTypesCreateParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*CircuitsCircuitTypesCreateCreated, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewCircuitsCircuitTypesCreateParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "circuits_circuit-types_create",
		Method:             "POST",
		PathPattern:        "/circuits/circuit-types/",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &CircuitsCircuitTypesCreateReader{formats: a.formats},
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
	success, ok := result.(*CircuitsCircuitTypesCreateCreated)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for circuits_circuit-types_create: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  CircuitsCircuitTypesDelete circuits circuit types delete API
*/
func (a *Client) CircuitsCircuitTypesDelete(params *CircuitsCircuitTypesDeleteParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*CircuitsCircuitTypesDeleteNoContent, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewCircuitsCircuitTypesDeleteParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "circuits_circuit-types_delete",
		Method:             "DELETE",
		PathPattern:        "/circuits/circuit-types/{id}/",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &CircuitsCircuitTypesDeleteReader{formats: a.formats},
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
	success, ok := result.(*CircuitsCircuitTypesDeleteNoContent)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for circuits_circuit-types_delete: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  CircuitsCircuitTypesList circuits circuit types list API
*/
func (a *Client) CircuitsCircuitTypesList(params *CircuitsCircuitTypesListParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*CircuitsCircuitTypesListOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewCircuitsCircuitTypesListParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "circuits_circuit-types_list",
		Method:             "GET",
		PathPattern:        "/circuits/circuit-types/",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &CircuitsCircuitTypesListReader{formats: a.formats},
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
	success, ok := result.(*CircuitsCircuitTypesListOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for circuits_circuit-types_list: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  CircuitsCircuitTypesPartialUpdate circuits circuit types partial update API
*/
func (a *Client) CircuitsCircuitTypesPartialUpdate(params *CircuitsCircuitTypesPartialUpdateParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*CircuitsCircuitTypesPartialUpdateOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewCircuitsCircuitTypesPartialUpdateParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "circuits_circuit-types_partial_update",
		Method:             "PATCH",
		PathPattern:        "/circuits/circuit-types/{id}/",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &CircuitsCircuitTypesPartialUpdateReader{formats: a.formats},
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
	success, ok := result.(*CircuitsCircuitTypesPartialUpdateOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for circuits_circuit-types_partial_update: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  CircuitsCircuitTypesRead circuits circuit types read API
*/
func (a *Client) CircuitsCircuitTypesRead(params *CircuitsCircuitTypesReadParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*CircuitsCircuitTypesReadOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewCircuitsCircuitTypesReadParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "circuits_circuit-types_read",
		Method:             "GET",
		PathPattern:        "/circuits/circuit-types/{id}/",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &CircuitsCircuitTypesReadReader{formats: a.formats},
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
	success, ok := result.(*CircuitsCircuitTypesReadOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for circuits_circuit-types_read: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  CircuitsCircuitTypesUpdate circuits circuit types update API
*/
func (a *Client) CircuitsCircuitTypesUpdate(params *CircuitsCircuitTypesUpdateParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*CircuitsCircuitTypesUpdateOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewCircuitsCircuitTypesUpdateParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "circuits_circuit-types_update",
		Method:             "PUT",
		PathPattern:        "/circuits/circuit-types/{id}/",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &CircuitsCircuitTypesUpdateReader{formats: a.formats},
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
	success, ok := result.(*CircuitsCircuitTypesUpdateOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for circuits_circuit-types_update: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  CircuitsCircuitsBulkDelete circuits circuits bulk delete API
*/
func (a *Client) CircuitsCircuitsBulkDelete(params *CircuitsCircuitsBulkDeleteParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*CircuitsCircuitsBulkDeleteNoContent, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewCircuitsCircuitsBulkDeleteParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "circuits_circuits_bulk_delete",
		Method:             "DELETE",
		PathPattern:        "/circuits/circuits/",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &CircuitsCircuitsBulkDeleteReader{formats: a.formats},
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
	success, ok := result.(*CircuitsCircuitsBulkDeleteNoContent)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for circuits_circuits_bulk_delete: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  CircuitsCircuitsBulkPartialUpdate circuits circuits bulk partial update API
*/
func (a *Client) CircuitsCircuitsBulkPartialUpdate(params *CircuitsCircuitsBulkPartialUpdateParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*CircuitsCircuitsBulkPartialUpdateOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewCircuitsCircuitsBulkPartialUpdateParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "circuits_circuits_bulk_partial_update",
		Method:             "PATCH",
		PathPattern:        "/circuits/circuits/",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &CircuitsCircuitsBulkPartialUpdateReader{formats: a.formats},
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
	success, ok := result.(*CircuitsCircuitsBulkPartialUpdateOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for circuits_circuits_bulk_partial_update: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  CircuitsCircuitsBulkUpdate circuits circuits bulk update API
*/
func (a *Client) CircuitsCircuitsBulkUpdate(params *CircuitsCircuitsBulkUpdateParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*CircuitsCircuitsBulkUpdateOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewCircuitsCircuitsBulkUpdateParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "circuits_circuits_bulk_update",
		Method:             "PUT",
		PathPattern:        "/circuits/circuits/",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &CircuitsCircuitsBulkUpdateReader{formats: a.formats},
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
	success, ok := result.(*CircuitsCircuitsBulkUpdateOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for circuits_circuits_bulk_update: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  CircuitsCircuitsCreate circuits circuits create API
*/
func (a *Client) CircuitsCircuitsCreate(params *CircuitsCircuitsCreateParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*CircuitsCircuitsCreateCreated, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewCircuitsCircuitsCreateParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "circuits_circuits_create",
		Method:             "POST",
		PathPattern:        "/circuits/circuits/",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &CircuitsCircuitsCreateReader{formats: a.formats},
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
	success, ok := result.(*CircuitsCircuitsCreateCreated)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for circuits_circuits_create: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  CircuitsCircuitsDelete circuits circuits delete API
*/
func (a *Client) CircuitsCircuitsDelete(params *CircuitsCircuitsDeleteParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*CircuitsCircuitsDeleteNoContent, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewCircuitsCircuitsDeleteParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "circuits_circuits_delete",
		Method:             "DELETE",
		PathPattern:        "/circuits/circuits/{id}/",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &CircuitsCircuitsDeleteReader{formats: a.formats},
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
	success, ok := result.(*CircuitsCircuitsDeleteNoContent)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for circuits_circuits_delete: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  CircuitsCircuitsList circuits circuits list API
*/
func (a *Client) CircuitsCircuitsList(params *CircuitsCircuitsListParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*CircuitsCircuitsListOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewCircuitsCircuitsListParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "circuits_circuits_list",
		Method:             "GET",
		PathPattern:        "/circuits/circuits/",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &CircuitsCircuitsListReader{formats: a.formats},
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
	success, ok := result.(*CircuitsCircuitsListOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for circuits_circuits_list: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  CircuitsCircuitsPartialUpdate circuits circuits partial update API
*/
func (a *Client) CircuitsCircuitsPartialUpdate(params *CircuitsCircuitsPartialUpdateParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*CircuitsCircuitsPartialUpdateOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewCircuitsCircuitsPartialUpdateParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "circuits_circuits_partial_update",
		Method:             "PATCH",
		PathPattern:        "/circuits/circuits/{id}/",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &CircuitsCircuitsPartialUpdateReader{formats: a.formats},
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
	success, ok := result.(*CircuitsCircuitsPartialUpdateOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for circuits_circuits_partial_update: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  CircuitsCircuitsRead circuits circuits read API
*/
func (a *Client) CircuitsCircuitsRead(params *CircuitsCircuitsReadParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*CircuitsCircuitsReadOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewCircuitsCircuitsReadParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "circuits_circuits_read",
		Method:             "GET",
		PathPattern:        "/circuits/circuits/{id}/",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &CircuitsCircuitsReadReader{formats: a.formats},
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
	success, ok := result.(*CircuitsCircuitsReadOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for circuits_circuits_read: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  CircuitsCircuitsUpdate circuits circuits update API
*/
func (a *Client) CircuitsCircuitsUpdate(params *CircuitsCircuitsUpdateParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*CircuitsCircuitsUpdateOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewCircuitsCircuitsUpdateParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "circuits_circuits_update",
		Method:             "PUT",
		PathPattern:        "/circuits/circuits/{id}/",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &CircuitsCircuitsUpdateReader{formats: a.formats},
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
	success, ok := result.(*CircuitsCircuitsUpdateOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for circuits_circuits_update: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  CircuitsProvidersBulkDelete circuits providers bulk delete API
*/
func (a *Client) CircuitsProvidersBulkDelete(params *CircuitsProvidersBulkDeleteParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*CircuitsProvidersBulkDeleteNoContent, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewCircuitsProvidersBulkDeleteParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "circuits_providers_bulk_delete",
		Method:             "DELETE",
		PathPattern:        "/circuits/providers/",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &CircuitsProvidersBulkDeleteReader{formats: a.formats},
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
	success, ok := result.(*CircuitsProvidersBulkDeleteNoContent)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for circuits_providers_bulk_delete: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  CircuitsProvidersBulkPartialUpdate circuits providers bulk partial update API
*/
func (a *Client) CircuitsProvidersBulkPartialUpdate(params *CircuitsProvidersBulkPartialUpdateParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*CircuitsProvidersBulkPartialUpdateOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewCircuitsProvidersBulkPartialUpdateParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "circuits_providers_bulk_partial_update",
		Method:             "PATCH",
		PathPattern:        "/circuits/providers/",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &CircuitsProvidersBulkPartialUpdateReader{formats: a.formats},
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
	success, ok := result.(*CircuitsProvidersBulkPartialUpdateOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for circuits_providers_bulk_partial_update: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  CircuitsProvidersBulkUpdate circuits providers bulk update API
*/
func (a *Client) CircuitsProvidersBulkUpdate(params *CircuitsProvidersBulkUpdateParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*CircuitsProvidersBulkUpdateOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewCircuitsProvidersBulkUpdateParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "circuits_providers_bulk_update",
		Method:             "PUT",
		PathPattern:        "/circuits/providers/",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &CircuitsProvidersBulkUpdateReader{formats: a.formats},
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
	success, ok := result.(*CircuitsProvidersBulkUpdateOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for circuits_providers_bulk_update: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  CircuitsProvidersCreate circuits providers create API
*/
func (a *Client) CircuitsProvidersCreate(params *CircuitsProvidersCreateParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*CircuitsProvidersCreateCreated, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewCircuitsProvidersCreateParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "circuits_providers_create",
		Method:             "POST",
		PathPattern:        "/circuits/providers/",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &CircuitsProvidersCreateReader{formats: a.formats},
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
	success, ok := result.(*CircuitsProvidersCreateCreated)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for circuits_providers_create: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  CircuitsProvidersDelete circuits providers delete API
*/
func (a *Client) CircuitsProvidersDelete(params *CircuitsProvidersDeleteParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*CircuitsProvidersDeleteNoContent, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewCircuitsProvidersDeleteParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "circuits_providers_delete",
		Method:             "DELETE",
		PathPattern:        "/circuits/providers/{id}/",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &CircuitsProvidersDeleteReader{formats: a.formats},
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
	success, ok := result.(*CircuitsProvidersDeleteNoContent)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for circuits_providers_delete: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  CircuitsProvidersList circuits providers list API
*/
func (a *Client) CircuitsProvidersList(params *CircuitsProvidersListParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*CircuitsProvidersListOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewCircuitsProvidersListParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "circuits_providers_list",
		Method:             "GET",
		PathPattern:        "/circuits/providers/",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &CircuitsProvidersListReader{formats: a.formats},
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
	success, ok := result.(*CircuitsProvidersListOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for circuits_providers_list: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  CircuitsProvidersPartialUpdate circuits providers partial update API
*/
func (a *Client) CircuitsProvidersPartialUpdate(params *CircuitsProvidersPartialUpdateParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*CircuitsProvidersPartialUpdateOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewCircuitsProvidersPartialUpdateParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "circuits_providers_partial_update",
		Method:             "PATCH",
		PathPattern:        "/circuits/providers/{id}/",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &CircuitsProvidersPartialUpdateReader{formats: a.formats},
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
	success, ok := result.(*CircuitsProvidersPartialUpdateOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for circuits_providers_partial_update: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  CircuitsProvidersRead circuits providers read API
*/
func (a *Client) CircuitsProvidersRead(params *CircuitsProvidersReadParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*CircuitsProvidersReadOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewCircuitsProvidersReadParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "circuits_providers_read",
		Method:             "GET",
		PathPattern:        "/circuits/providers/{id}/",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &CircuitsProvidersReadReader{formats: a.formats},
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
	success, ok := result.(*CircuitsProvidersReadOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for circuits_providers_read: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  CircuitsProvidersUpdate circuits providers update API
*/
func (a *Client) CircuitsProvidersUpdate(params *CircuitsProvidersUpdateParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*CircuitsProvidersUpdateOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewCircuitsProvidersUpdateParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "circuits_providers_update",
		Method:             "PUT",
		PathPattern:        "/circuits/providers/{id}/",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &CircuitsProvidersUpdateReader{formats: a.formats},
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
	success, ok := result.(*CircuitsProvidersUpdateOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for circuits_providers_update: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
