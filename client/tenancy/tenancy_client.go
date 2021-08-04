package tenancy

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// New creates a new tenancy API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

/*
Client for tenancy API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientOption is the option for Client methods
type ClientOption func(*runtime.ClientOperation)

// ClientService is the interface for Client methods
type ClientService interface {
	TenancyTenantGroupsBulkDelete(params *TenancyTenantGroupsBulkDeleteParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*TenancyTenantGroupsBulkDeleteNoContent, error)

	TenancyTenantGroupsBulkPartialUpdate(params *TenancyTenantGroupsBulkPartialUpdateParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*TenancyTenantGroupsBulkPartialUpdateOK, error)

	TenancyTenantGroupsBulkUpdate(params *TenancyTenantGroupsBulkUpdateParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*TenancyTenantGroupsBulkUpdateOK, error)

	TenancyTenantGroupsCreate(params *TenancyTenantGroupsCreateParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*TenancyTenantGroupsCreateCreated, error)

	TenancyTenantGroupsDelete(params *TenancyTenantGroupsDeleteParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*TenancyTenantGroupsDeleteNoContent, error)

	TenancyTenantGroupsList(params *TenancyTenantGroupsListParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*TenancyTenantGroupsListOK, error)

	TenancyTenantGroupsPartialUpdate(params *TenancyTenantGroupsPartialUpdateParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*TenancyTenantGroupsPartialUpdateOK, error)

	TenancyTenantGroupsRead(params *TenancyTenantGroupsReadParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*TenancyTenantGroupsReadOK, error)

	TenancyTenantGroupsUpdate(params *TenancyTenantGroupsUpdateParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*TenancyTenantGroupsUpdateOK, error)

	TenancyTenantsBulkDelete(params *TenancyTenantsBulkDeleteParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*TenancyTenantsBulkDeleteNoContent, error)

	TenancyTenantsBulkPartialUpdate(params *TenancyTenantsBulkPartialUpdateParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*TenancyTenantsBulkPartialUpdateOK, error)

	TenancyTenantsBulkUpdate(params *TenancyTenantsBulkUpdateParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*TenancyTenantsBulkUpdateOK, error)

	TenancyTenantsCreate(params *TenancyTenantsCreateParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*TenancyTenantsCreateCreated, error)

	TenancyTenantsDelete(params *TenancyTenantsDeleteParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*TenancyTenantsDeleteNoContent, error)

	TenancyTenantsList(params *TenancyTenantsListParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*TenancyTenantsListOK, error)

	TenancyTenantsPartialUpdate(params *TenancyTenantsPartialUpdateParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*TenancyTenantsPartialUpdateOK, error)

	TenancyTenantsRead(params *TenancyTenantsReadParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*TenancyTenantsReadOK, error)

	TenancyTenantsUpdate(params *TenancyTenantsUpdateParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*TenancyTenantsUpdateOK, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
  TenancyTenantGroupsBulkDelete tenancy tenant groups bulk delete API
*/
func (a *Client) TenancyTenantGroupsBulkDelete(params *TenancyTenantGroupsBulkDeleteParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*TenancyTenantGroupsBulkDeleteNoContent, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewTenancyTenantGroupsBulkDeleteParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "tenancy_tenant-groups_bulk_delete",
		Method:             "DELETE",
		PathPattern:        "/tenancy/tenant-groups/",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &TenancyTenantGroupsBulkDeleteReader{formats: a.formats},
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
	success, ok := result.(*TenancyTenantGroupsBulkDeleteNoContent)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for tenancy_tenant-groups_bulk_delete: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  TenancyTenantGroupsBulkPartialUpdate tenancy tenant groups bulk partial update API
*/
func (a *Client) TenancyTenantGroupsBulkPartialUpdate(params *TenancyTenantGroupsBulkPartialUpdateParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*TenancyTenantGroupsBulkPartialUpdateOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewTenancyTenantGroupsBulkPartialUpdateParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "tenancy_tenant-groups_bulk_partial_update",
		Method:             "PATCH",
		PathPattern:        "/tenancy/tenant-groups/",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &TenancyTenantGroupsBulkPartialUpdateReader{formats: a.formats},
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
	success, ok := result.(*TenancyTenantGroupsBulkPartialUpdateOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for tenancy_tenant-groups_bulk_partial_update: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  TenancyTenantGroupsBulkUpdate tenancy tenant groups bulk update API
*/
func (a *Client) TenancyTenantGroupsBulkUpdate(params *TenancyTenantGroupsBulkUpdateParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*TenancyTenantGroupsBulkUpdateOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewTenancyTenantGroupsBulkUpdateParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "tenancy_tenant-groups_bulk_update",
		Method:             "PUT",
		PathPattern:        "/tenancy/tenant-groups/",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &TenancyTenantGroupsBulkUpdateReader{formats: a.formats},
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
	success, ok := result.(*TenancyTenantGroupsBulkUpdateOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for tenancy_tenant-groups_bulk_update: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  TenancyTenantGroupsCreate tenancy tenant groups create API
*/
func (a *Client) TenancyTenantGroupsCreate(params *TenancyTenantGroupsCreateParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*TenancyTenantGroupsCreateCreated, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewTenancyTenantGroupsCreateParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "tenancy_tenant-groups_create",
		Method:             "POST",
		PathPattern:        "/tenancy/tenant-groups/",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &TenancyTenantGroupsCreateReader{formats: a.formats},
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
	success, ok := result.(*TenancyTenantGroupsCreateCreated)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for tenancy_tenant-groups_create: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  TenancyTenantGroupsDelete tenancy tenant groups delete API
*/
func (a *Client) TenancyTenantGroupsDelete(params *TenancyTenantGroupsDeleteParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*TenancyTenantGroupsDeleteNoContent, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewTenancyTenantGroupsDeleteParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "tenancy_tenant-groups_delete",
		Method:             "DELETE",
		PathPattern:        "/tenancy/tenant-groups/{id}/",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &TenancyTenantGroupsDeleteReader{formats: a.formats},
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
	success, ok := result.(*TenancyTenantGroupsDeleteNoContent)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for tenancy_tenant-groups_delete: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  TenancyTenantGroupsList tenancy tenant groups list API
*/
func (a *Client) TenancyTenantGroupsList(params *TenancyTenantGroupsListParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*TenancyTenantGroupsListOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewTenancyTenantGroupsListParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "tenancy_tenant-groups_list",
		Method:             "GET",
		PathPattern:        "/tenancy/tenant-groups/",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &TenancyTenantGroupsListReader{formats: a.formats},
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
	success, ok := result.(*TenancyTenantGroupsListOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for tenancy_tenant-groups_list: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  TenancyTenantGroupsPartialUpdate tenancy tenant groups partial update API
*/
func (a *Client) TenancyTenantGroupsPartialUpdate(params *TenancyTenantGroupsPartialUpdateParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*TenancyTenantGroupsPartialUpdateOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewTenancyTenantGroupsPartialUpdateParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "tenancy_tenant-groups_partial_update",
		Method:             "PATCH",
		PathPattern:        "/tenancy/tenant-groups/{id}/",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &TenancyTenantGroupsPartialUpdateReader{formats: a.formats},
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
	success, ok := result.(*TenancyTenantGroupsPartialUpdateOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for tenancy_tenant-groups_partial_update: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  TenancyTenantGroupsRead tenancy tenant groups read API
*/
func (a *Client) TenancyTenantGroupsRead(params *TenancyTenantGroupsReadParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*TenancyTenantGroupsReadOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewTenancyTenantGroupsReadParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "tenancy_tenant-groups_read",
		Method:             "GET",
		PathPattern:        "/tenancy/tenant-groups/{id}/",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &TenancyTenantGroupsReadReader{formats: a.formats},
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
	success, ok := result.(*TenancyTenantGroupsReadOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for tenancy_tenant-groups_read: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  TenancyTenantGroupsUpdate tenancy tenant groups update API
*/
func (a *Client) TenancyTenantGroupsUpdate(params *TenancyTenantGroupsUpdateParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*TenancyTenantGroupsUpdateOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewTenancyTenantGroupsUpdateParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "tenancy_tenant-groups_update",
		Method:             "PUT",
		PathPattern:        "/tenancy/tenant-groups/{id}/",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &TenancyTenantGroupsUpdateReader{formats: a.formats},
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
	success, ok := result.(*TenancyTenantGroupsUpdateOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for tenancy_tenant-groups_update: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  TenancyTenantsBulkDelete tenancy tenants bulk delete API
*/
func (a *Client) TenancyTenantsBulkDelete(params *TenancyTenantsBulkDeleteParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*TenancyTenantsBulkDeleteNoContent, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewTenancyTenantsBulkDeleteParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "tenancy_tenants_bulk_delete",
		Method:             "DELETE",
		PathPattern:        "/tenancy/tenants/",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &TenancyTenantsBulkDeleteReader{formats: a.formats},
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
	success, ok := result.(*TenancyTenantsBulkDeleteNoContent)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for tenancy_tenants_bulk_delete: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  TenancyTenantsBulkPartialUpdate tenancy tenants bulk partial update API
*/
func (a *Client) TenancyTenantsBulkPartialUpdate(params *TenancyTenantsBulkPartialUpdateParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*TenancyTenantsBulkPartialUpdateOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewTenancyTenantsBulkPartialUpdateParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "tenancy_tenants_bulk_partial_update",
		Method:             "PATCH",
		PathPattern:        "/tenancy/tenants/",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &TenancyTenantsBulkPartialUpdateReader{formats: a.formats},
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
	success, ok := result.(*TenancyTenantsBulkPartialUpdateOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for tenancy_tenants_bulk_partial_update: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  TenancyTenantsBulkUpdate tenancy tenants bulk update API
*/
func (a *Client) TenancyTenantsBulkUpdate(params *TenancyTenantsBulkUpdateParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*TenancyTenantsBulkUpdateOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewTenancyTenantsBulkUpdateParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "tenancy_tenants_bulk_update",
		Method:             "PUT",
		PathPattern:        "/tenancy/tenants/",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &TenancyTenantsBulkUpdateReader{formats: a.formats},
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
	success, ok := result.(*TenancyTenantsBulkUpdateOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for tenancy_tenants_bulk_update: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  TenancyTenantsCreate tenancy tenants create API
*/
func (a *Client) TenancyTenantsCreate(params *TenancyTenantsCreateParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*TenancyTenantsCreateCreated, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewTenancyTenantsCreateParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "tenancy_tenants_create",
		Method:             "POST",
		PathPattern:        "/tenancy/tenants/",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &TenancyTenantsCreateReader{formats: a.formats},
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
	success, ok := result.(*TenancyTenantsCreateCreated)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for tenancy_tenants_create: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  TenancyTenantsDelete tenancy tenants delete API
*/
func (a *Client) TenancyTenantsDelete(params *TenancyTenantsDeleteParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*TenancyTenantsDeleteNoContent, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewTenancyTenantsDeleteParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "tenancy_tenants_delete",
		Method:             "DELETE",
		PathPattern:        "/tenancy/tenants/{id}/",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &TenancyTenantsDeleteReader{formats: a.formats},
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
	success, ok := result.(*TenancyTenantsDeleteNoContent)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for tenancy_tenants_delete: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  TenancyTenantsList tenancy tenants list API
*/
func (a *Client) TenancyTenantsList(params *TenancyTenantsListParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*TenancyTenantsListOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewTenancyTenantsListParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "tenancy_tenants_list",
		Method:             "GET",
		PathPattern:        "/tenancy/tenants/",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &TenancyTenantsListReader{formats: a.formats},
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
	success, ok := result.(*TenancyTenantsListOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for tenancy_tenants_list: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  TenancyTenantsPartialUpdate tenancy tenants partial update API
*/
func (a *Client) TenancyTenantsPartialUpdate(params *TenancyTenantsPartialUpdateParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*TenancyTenantsPartialUpdateOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewTenancyTenantsPartialUpdateParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "tenancy_tenants_partial_update",
		Method:             "PATCH",
		PathPattern:        "/tenancy/tenants/{id}/",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &TenancyTenantsPartialUpdateReader{formats: a.formats},
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
	success, ok := result.(*TenancyTenantsPartialUpdateOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for tenancy_tenants_partial_update: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  TenancyTenantsRead tenancy tenants read API
*/
func (a *Client) TenancyTenantsRead(params *TenancyTenantsReadParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*TenancyTenantsReadOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewTenancyTenantsReadParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "tenancy_tenants_read",
		Method:             "GET",
		PathPattern:        "/tenancy/tenants/{id}/",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &TenancyTenantsReadReader{formats: a.formats},
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
	success, ok := result.(*TenancyTenantsReadOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for tenancy_tenants_read: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  TenancyTenantsUpdate tenancy tenants update API
*/
func (a *Client) TenancyTenantsUpdate(params *TenancyTenantsUpdateParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*TenancyTenantsUpdateOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewTenancyTenantsUpdateParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "tenancy_tenants_update",
		Method:             "PUT",
		PathPattern:        "/tenancy/tenants/{id}/",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &TenancyTenantsUpdateReader{formats: a.formats},
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
	success, ok := result.(*TenancyTenantsUpdateOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for tenancy_tenants_update: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
