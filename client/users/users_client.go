package users

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// New creates a new users API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

/*
Client for users API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientOption is the option for Client methods
type ClientOption func(*runtime.ClientOperation)

// ClientService is the interface for Client methods
type ClientService interface {
	UsersConfigList(params *UsersConfigListParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*UsersConfigListOK, error)

	UsersGroupsBulkDelete(params *UsersGroupsBulkDeleteParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*UsersGroupsBulkDeleteNoContent, error)

	UsersGroupsBulkPartialUpdate(params *UsersGroupsBulkPartialUpdateParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*UsersGroupsBulkPartialUpdateOK, error)

	UsersGroupsBulkUpdate(params *UsersGroupsBulkUpdateParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*UsersGroupsBulkUpdateOK, error)

	UsersGroupsCreate(params *UsersGroupsCreateParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*UsersGroupsCreateCreated, error)

	UsersGroupsDelete(params *UsersGroupsDeleteParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*UsersGroupsDeleteNoContent, error)

	UsersGroupsList(params *UsersGroupsListParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*UsersGroupsListOK, error)

	UsersGroupsPartialUpdate(params *UsersGroupsPartialUpdateParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*UsersGroupsPartialUpdateOK, error)

	UsersGroupsRead(params *UsersGroupsReadParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*UsersGroupsReadOK, error)

	UsersGroupsUpdate(params *UsersGroupsUpdateParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*UsersGroupsUpdateOK, error)

	UsersPermissionsBulkDelete(params *UsersPermissionsBulkDeleteParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*UsersPermissionsBulkDeleteNoContent, error)

	UsersPermissionsBulkPartialUpdate(params *UsersPermissionsBulkPartialUpdateParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*UsersPermissionsBulkPartialUpdateOK, error)

	UsersPermissionsBulkUpdate(params *UsersPermissionsBulkUpdateParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*UsersPermissionsBulkUpdateOK, error)

	UsersPermissionsCreate(params *UsersPermissionsCreateParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*UsersPermissionsCreateCreated, error)

	UsersPermissionsDelete(params *UsersPermissionsDeleteParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*UsersPermissionsDeleteNoContent, error)

	UsersPermissionsList(params *UsersPermissionsListParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*UsersPermissionsListOK, error)

	UsersPermissionsPartialUpdate(params *UsersPermissionsPartialUpdateParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*UsersPermissionsPartialUpdateOK, error)

	UsersPermissionsRead(params *UsersPermissionsReadParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*UsersPermissionsReadOK, error)

	UsersPermissionsUpdate(params *UsersPermissionsUpdateParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*UsersPermissionsUpdateOK, error)

	UsersUsersBulkDelete(params *UsersUsersBulkDeleteParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*UsersUsersBulkDeleteNoContent, error)

	UsersUsersBulkPartialUpdate(params *UsersUsersBulkPartialUpdateParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*UsersUsersBulkPartialUpdateOK, error)

	UsersUsersBulkUpdate(params *UsersUsersBulkUpdateParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*UsersUsersBulkUpdateOK, error)

	UsersUsersCreate(params *UsersUsersCreateParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*UsersUsersCreateCreated, error)

	UsersUsersDelete(params *UsersUsersDeleteParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*UsersUsersDeleteNoContent, error)

	UsersUsersList(params *UsersUsersListParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*UsersUsersListOK, error)

	UsersUsersPartialUpdate(params *UsersUsersPartialUpdateParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*UsersUsersPartialUpdateOK, error)

	UsersUsersRead(params *UsersUsersReadParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*UsersUsersReadOK, error)

	UsersUsersUpdate(params *UsersUsersUpdateParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*UsersUsersUpdateOK, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
  UsersConfigList Return the config_data for the currently authenticated User.
*/
func (a *Client) UsersConfigList(params *UsersConfigListParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*UsersConfigListOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewUsersConfigListParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "users_config_list",
		Method:             "GET",
		PathPattern:        "/users/config/",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &UsersConfigListReader{formats: a.formats},
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
	success, ok := result.(*UsersConfigListOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for users_config_list: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  UsersGroupsBulkDelete users groups bulk delete API
*/
func (a *Client) UsersGroupsBulkDelete(params *UsersGroupsBulkDeleteParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*UsersGroupsBulkDeleteNoContent, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewUsersGroupsBulkDeleteParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "users_groups_bulk_delete",
		Method:             "DELETE",
		PathPattern:        "/users/groups/",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &UsersGroupsBulkDeleteReader{formats: a.formats},
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
	success, ok := result.(*UsersGroupsBulkDeleteNoContent)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for users_groups_bulk_delete: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  UsersGroupsBulkPartialUpdate users groups bulk partial update API
*/
func (a *Client) UsersGroupsBulkPartialUpdate(params *UsersGroupsBulkPartialUpdateParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*UsersGroupsBulkPartialUpdateOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewUsersGroupsBulkPartialUpdateParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "users_groups_bulk_partial_update",
		Method:             "PATCH",
		PathPattern:        "/users/groups/",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &UsersGroupsBulkPartialUpdateReader{formats: a.formats},
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
	success, ok := result.(*UsersGroupsBulkPartialUpdateOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for users_groups_bulk_partial_update: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  UsersGroupsBulkUpdate users groups bulk update API
*/
func (a *Client) UsersGroupsBulkUpdate(params *UsersGroupsBulkUpdateParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*UsersGroupsBulkUpdateOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewUsersGroupsBulkUpdateParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "users_groups_bulk_update",
		Method:             "PUT",
		PathPattern:        "/users/groups/",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &UsersGroupsBulkUpdateReader{formats: a.formats},
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
	success, ok := result.(*UsersGroupsBulkUpdateOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for users_groups_bulk_update: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  UsersGroupsCreate users groups create API
*/
func (a *Client) UsersGroupsCreate(params *UsersGroupsCreateParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*UsersGroupsCreateCreated, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewUsersGroupsCreateParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "users_groups_create",
		Method:             "POST",
		PathPattern:        "/users/groups/",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &UsersGroupsCreateReader{formats: a.formats},
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
	success, ok := result.(*UsersGroupsCreateCreated)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for users_groups_create: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  UsersGroupsDelete users groups delete API
*/
func (a *Client) UsersGroupsDelete(params *UsersGroupsDeleteParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*UsersGroupsDeleteNoContent, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewUsersGroupsDeleteParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "users_groups_delete",
		Method:             "DELETE",
		PathPattern:        "/users/groups/{id}/",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &UsersGroupsDeleteReader{formats: a.formats},
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
	success, ok := result.(*UsersGroupsDeleteNoContent)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for users_groups_delete: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  UsersGroupsList users groups list API
*/
func (a *Client) UsersGroupsList(params *UsersGroupsListParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*UsersGroupsListOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewUsersGroupsListParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "users_groups_list",
		Method:             "GET",
		PathPattern:        "/users/groups/",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &UsersGroupsListReader{formats: a.formats},
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
	success, ok := result.(*UsersGroupsListOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for users_groups_list: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  UsersGroupsPartialUpdate users groups partial update API
*/
func (a *Client) UsersGroupsPartialUpdate(params *UsersGroupsPartialUpdateParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*UsersGroupsPartialUpdateOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewUsersGroupsPartialUpdateParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "users_groups_partial_update",
		Method:             "PATCH",
		PathPattern:        "/users/groups/{id}/",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &UsersGroupsPartialUpdateReader{formats: a.formats},
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
	success, ok := result.(*UsersGroupsPartialUpdateOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for users_groups_partial_update: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  UsersGroupsRead users groups read API
*/
func (a *Client) UsersGroupsRead(params *UsersGroupsReadParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*UsersGroupsReadOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewUsersGroupsReadParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "users_groups_read",
		Method:             "GET",
		PathPattern:        "/users/groups/{id}/",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &UsersGroupsReadReader{formats: a.formats},
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
	success, ok := result.(*UsersGroupsReadOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for users_groups_read: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  UsersGroupsUpdate users groups update API
*/
func (a *Client) UsersGroupsUpdate(params *UsersGroupsUpdateParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*UsersGroupsUpdateOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewUsersGroupsUpdateParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "users_groups_update",
		Method:             "PUT",
		PathPattern:        "/users/groups/{id}/",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &UsersGroupsUpdateReader{formats: a.formats},
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
	success, ok := result.(*UsersGroupsUpdateOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for users_groups_update: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  UsersPermissionsBulkDelete users permissions bulk delete API
*/
func (a *Client) UsersPermissionsBulkDelete(params *UsersPermissionsBulkDeleteParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*UsersPermissionsBulkDeleteNoContent, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewUsersPermissionsBulkDeleteParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "users_permissions_bulk_delete",
		Method:             "DELETE",
		PathPattern:        "/users/permissions/",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &UsersPermissionsBulkDeleteReader{formats: a.formats},
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
	success, ok := result.(*UsersPermissionsBulkDeleteNoContent)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for users_permissions_bulk_delete: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  UsersPermissionsBulkPartialUpdate users permissions bulk partial update API
*/
func (a *Client) UsersPermissionsBulkPartialUpdate(params *UsersPermissionsBulkPartialUpdateParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*UsersPermissionsBulkPartialUpdateOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewUsersPermissionsBulkPartialUpdateParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "users_permissions_bulk_partial_update",
		Method:             "PATCH",
		PathPattern:        "/users/permissions/",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &UsersPermissionsBulkPartialUpdateReader{formats: a.formats},
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
	success, ok := result.(*UsersPermissionsBulkPartialUpdateOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for users_permissions_bulk_partial_update: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  UsersPermissionsBulkUpdate users permissions bulk update API
*/
func (a *Client) UsersPermissionsBulkUpdate(params *UsersPermissionsBulkUpdateParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*UsersPermissionsBulkUpdateOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewUsersPermissionsBulkUpdateParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "users_permissions_bulk_update",
		Method:             "PUT",
		PathPattern:        "/users/permissions/",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &UsersPermissionsBulkUpdateReader{formats: a.formats},
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
	success, ok := result.(*UsersPermissionsBulkUpdateOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for users_permissions_bulk_update: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  UsersPermissionsCreate users permissions create API
*/
func (a *Client) UsersPermissionsCreate(params *UsersPermissionsCreateParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*UsersPermissionsCreateCreated, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewUsersPermissionsCreateParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "users_permissions_create",
		Method:             "POST",
		PathPattern:        "/users/permissions/",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &UsersPermissionsCreateReader{formats: a.formats},
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
	success, ok := result.(*UsersPermissionsCreateCreated)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for users_permissions_create: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  UsersPermissionsDelete users permissions delete API
*/
func (a *Client) UsersPermissionsDelete(params *UsersPermissionsDeleteParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*UsersPermissionsDeleteNoContent, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewUsersPermissionsDeleteParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "users_permissions_delete",
		Method:             "DELETE",
		PathPattern:        "/users/permissions/{id}/",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &UsersPermissionsDeleteReader{formats: a.formats},
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
	success, ok := result.(*UsersPermissionsDeleteNoContent)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for users_permissions_delete: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  UsersPermissionsList users permissions list API
*/
func (a *Client) UsersPermissionsList(params *UsersPermissionsListParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*UsersPermissionsListOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewUsersPermissionsListParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "users_permissions_list",
		Method:             "GET",
		PathPattern:        "/users/permissions/",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &UsersPermissionsListReader{formats: a.formats},
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
	success, ok := result.(*UsersPermissionsListOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for users_permissions_list: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  UsersPermissionsPartialUpdate users permissions partial update API
*/
func (a *Client) UsersPermissionsPartialUpdate(params *UsersPermissionsPartialUpdateParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*UsersPermissionsPartialUpdateOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewUsersPermissionsPartialUpdateParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "users_permissions_partial_update",
		Method:             "PATCH",
		PathPattern:        "/users/permissions/{id}/",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &UsersPermissionsPartialUpdateReader{formats: a.formats},
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
	success, ok := result.(*UsersPermissionsPartialUpdateOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for users_permissions_partial_update: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  UsersPermissionsRead users permissions read API
*/
func (a *Client) UsersPermissionsRead(params *UsersPermissionsReadParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*UsersPermissionsReadOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewUsersPermissionsReadParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "users_permissions_read",
		Method:             "GET",
		PathPattern:        "/users/permissions/{id}/",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &UsersPermissionsReadReader{formats: a.formats},
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
	success, ok := result.(*UsersPermissionsReadOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for users_permissions_read: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  UsersPermissionsUpdate users permissions update API
*/
func (a *Client) UsersPermissionsUpdate(params *UsersPermissionsUpdateParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*UsersPermissionsUpdateOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewUsersPermissionsUpdateParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "users_permissions_update",
		Method:             "PUT",
		PathPattern:        "/users/permissions/{id}/",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &UsersPermissionsUpdateReader{formats: a.formats},
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
	success, ok := result.(*UsersPermissionsUpdateOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for users_permissions_update: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  UsersUsersBulkDelete users users bulk delete API
*/
func (a *Client) UsersUsersBulkDelete(params *UsersUsersBulkDeleteParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*UsersUsersBulkDeleteNoContent, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewUsersUsersBulkDeleteParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "users_users_bulk_delete",
		Method:             "DELETE",
		PathPattern:        "/users/users/",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &UsersUsersBulkDeleteReader{formats: a.formats},
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
	success, ok := result.(*UsersUsersBulkDeleteNoContent)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for users_users_bulk_delete: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  UsersUsersBulkPartialUpdate users users bulk partial update API
*/
func (a *Client) UsersUsersBulkPartialUpdate(params *UsersUsersBulkPartialUpdateParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*UsersUsersBulkPartialUpdateOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewUsersUsersBulkPartialUpdateParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "users_users_bulk_partial_update",
		Method:             "PATCH",
		PathPattern:        "/users/users/",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &UsersUsersBulkPartialUpdateReader{formats: a.formats},
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
	success, ok := result.(*UsersUsersBulkPartialUpdateOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for users_users_bulk_partial_update: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  UsersUsersBulkUpdate users users bulk update API
*/
func (a *Client) UsersUsersBulkUpdate(params *UsersUsersBulkUpdateParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*UsersUsersBulkUpdateOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewUsersUsersBulkUpdateParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "users_users_bulk_update",
		Method:             "PUT",
		PathPattern:        "/users/users/",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &UsersUsersBulkUpdateReader{formats: a.formats},
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
	success, ok := result.(*UsersUsersBulkUpdateOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for users_users_bulk_update: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  UsersUsersCreate users users create API
*/
func (a *Client) UsersUsersCreate(params *UsersUsersCreateParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*UsersUsersCreateCreated, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewUsersUsersCreateParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "users_users_create",
		Method:             "POST",
		PathPattern:        "/users/users/",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &UsersUsersCreateReader{formats: a.formats},
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
	success, ok := result.(*UsersUsersCreateCreated)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for users_users_create: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  UsersUsersDelete users users delete API
*/
func (a *Client) UsersUsersDelete(params *UsersUsersDeleteParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*UsersUsersDeleteNoContent, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewUsersUsersDeleteParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "users_users_delete",
		Method:             "DELETE",
		PathPattern:        "/users/users/{id}/",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &UsersUsersDeleteReader{formats: a.formats},
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
	success, ok := result.(*UsersUsersDeleteNoContent)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for users_users_delete: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  UsersUsersList users users list API
*/
func (a *Client) UsersUsersList(params *UsersUsersListParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*UsersUsersListOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewUsersUsersListParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "users_users_list",
		Method:             "GET",
		PathPattern:        "/users/users/",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &UsersUsersListReader{formats: a.formats},
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
	success, ok := result.(*UsersUsersListOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for users_users_list: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  UsersUsersPartialUpdate users users partial update API
*/
func (a *Client) UsersUsersPartialUpdate(params *UsersUsersPartialUpdateParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*UsersUsersPartialUpdateOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewUsersUsersPartialUpdateParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "users_users_partial_update",
		Method:             "PATCH",
		PathPattern:        "/users/users/{id}/",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &UsersUsersPartialUpdateReader{formats: a.formats},
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
	success, ok := result.(*UsersUsersPartialUpdateOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for users_users_partial_update: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  UsersUsersRead users users read API
*/
func (a *Client) UsersUsersRead(params *UsersUsersReadParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*UsersUsersReadOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewUsersUsersReadParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "users_users_read",
		Method:             "GET",
		PathPattern:        "/users/users/{id}/",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &UsersUsersReadReader{formats: a.formats},
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
	success, ok := result.(*UsersUsersReadOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for users_users_read: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  UsersUsersUpdate users users update API
*/
func (a *Client) UsersUsersUpdate(params *UsersUsersUpdateParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*UsersUsersUpdateOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewUsersUsersUpdateParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "users_users_update",
		Method:             "PUT",
		PathPattern:        "/users/users/{id}/",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &UsersUsersUpdateReader{formats: a.formats},
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
	success, ok := result.(*UsersUsersUpdateOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for users_users_update: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
