package virtualization

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// New creates a new virtualization API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

/*
Client for virtualization API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientOption is the option for Client methods
type ClientOption func(*runtime.ClientOperation)

// ClientService is the interface for Client methods
type ClientService interface {
	VirtualizationClusterGroupsBulkDelete(params *VirtualizationClusterGroupsBulkDeleteParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*VirtualizationClusterGroupsBulkDeleteNoContent, error)

	VirtualizationClusterGroupsBulkPartialUpdate(params *VirtualizationClusterGroupsBulkPartialUpdateParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*VirtualizationClusterGroupsBulkPartialUpdateOK, error)

	VirtualizationClusterGroupsBulkUpdate(params *VirtualizationClusterGroupsBulkUpdateParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*VirtualizationClusterGroupsBulkUpdateOK, error)

	VirtualizationClusterGroupsCreate(params *VirtualizationClusterGroupsCreateParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*VirtualizationClusterGroupsCreateCreated, error)

	VirtualizationClusterGroupsDelete(params *VirtualizationClusterGroupsDeleteParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*VirtualizationClusterGroupsDeleteNoContent, error)

	VirtualizationClusterGroupsList(params *VirtualizationClusterGroupsListParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*VirtualizationClusterGroupsListOK, error)

	VirtualizationClusterGroupsPartialUpdate(params *VirtualizationClusterGroupsPartialUpdateParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*VirtualizationClusterGroupsPartialUpdateOK, error)

	VirtualizationClusterGroupsRead(params *VirtualizationClusterGroupsReadParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*VirtualizationClusterGroupsReadOK, error)

	VirtualizationClusterGroupsUpdate(params *VirtualizationClusterGroupsUpdateParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*VirtualizationClusterGroupsUpdateOK, error)

	VirtualizationClusterTypesBulkDelete(params *VirtualizationClusterTypesBulkDeleteParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*VirtualizationClusterTypesBulkDeleteNoContent, error)

	VirtualizationClusterTypesBulkPartialUpdate(params *VirtualizationClusterTypesBulkPartialUpdateParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*VirtualizationClusterTypesBulkPartialUpdateOK, error)

	VirtualizationClusterTypesBulkUpdate(params *VirtualizationClusterTypesBulkUpdateParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*VirtualizationClusterTypesBulkUpdateOK, error)

	VirtualizationClusterTypesCreate(params *VirtualizationClusterTypesCreateParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*VirtualizationClusterTypesCreateCreated, error)

	VirtualizationClusterTypesDelete(params *VirtualizationClusterTypesDeleteParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*VirtualizationClusterTypesDeleteNoContent, error)

	VirtualizationClusterTypesList(params *VirtualizationClusterTypesListParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*VirtualizationClusterTypesListOK, error)

	VirtualizationClusterTypesPartialUpdate(params *VirtualizationClusterTypesPartialUpdateParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*VirtualizationClusterTypesPartialUpdateOK, error)

	VirtualizationClusterTypesRead(params *VirtualizationClusterTypesReadParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*VirtualizationClusterTypesReadOK, error)

	VirtualizationClusterTypesUpdate(params *VirtualizationClusterTypesUpdateParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*VirtualizationClusterTypesUpdateOK, error)

	VirtualizationClustersBulkDelete(params *VirtualizationClustersBulkDeleteParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*VirtualizationClustersBulkDeleteNoContent, error)

	VirtualizationClustersBulkPartialUpdate(params *VirtualizationClustersBulkPartialUpdateParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*VirtualizationClustersBulkPartialUpdateOK, error)

	VirtualizationClustersBulkUpdate(params *VirtualizationClustersBulkUpdateParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*VirtualizationClustersBulkUpdateOK, error)

	VirtualizationClustersCreate(params *VirtualizationClustersCreateParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*VirtualizationClustersCreateCreated, error)

	VirtualizationClustersDelete(params *VirtualizationClustersDeleteParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*VirtualizationClustersDeleteNoContent, error)

	VirtualizationClustersList(params *VirtualizationClustersListParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*VirtualizationClustersListOK, error)

	VirtualizationClustersPartialUpdate(params *VirtualizationClustersPartialUpdateParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*VirtualizationClustersPartialUpdateOK, error)

	VirtualizationClustersRead(params *VirtualizationClustersReadParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*VirtualizationClustersReadOK, error)

	VirtualizationClustersUpdate(params *VirtualizationClustersUpdateParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*VirtualizationClustersUpdateOK, error)

	VirtualizationInterfacesBulkDelete(params *VirtualizationInterfacesBulkDeleteParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*VirtualizationInterfacesBulkDeleteNoContent, error)

	VirtualizationInterfacesBulkPartialUpdate(params *VirtualizationInterfacesBulkPartialUpdateParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*VirtualizationInterfacesBulkPartialUpdateOK, error)

	VirtualizationInterfacesBulkUpdate(params *VirtualizationInterfacesBulkUpdateParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*VirtualizationInterfacesBulkUpdateOK, error)

	VirtualizationInterfacesCreate(params *VirtualizationInterfacesCreateParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*VirtualizationInterfacesCreateCreated, error)

	VirtualizationInterfacesDelete(params *VirtualizationInterfacesDeleteParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*VirtualizationInterfacesDeleteNoContent, error)

	VirtualizationInterfacesList(params *VirtualizationInterfacesListParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*VirtualizationInterfacesListOK, error)

	VirtualizationInterfacesPartialUpdate(params *VirtualizationInterfacesPartialUpdateParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*VirtualizationInterfacesPartialUpdateOK, error)

	VirtualizationInterfacesRead(params *VirtualizationInterfacesReadParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*VirtualizationInterfacesReadOK, error)

	VirtualizationInterfacesUpdate(params *VirtualizationInterfacesUpdateParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*VirtualizationInterfacesUpdateOK, error)

	VirtualizationVirtualMachinesBulkDelete(params *VirtualizationVirtualMachinesBulkDeleteParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*VirtualizationVirtualMachinesBulkDeleteNoContent, error)

	VirtualizationVirtualMachinesBulkPartialUpdate(params *VirtualizationVirtualMachinesBulkPartialUpdateParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*VirtualizationVirtualMachinesBulkPartialUpdateOK, error)

	VirtualizationVirtualMachinesBulkUpdate(params *VirtualizationVirtualMachinesBulkUpdateParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*VirtualizationVirtualMachinesBulkUpdateOK, error)

	VirtualizationVirtualMachinesCreate(params *VirtualizationVirtualMachinesCreateParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*VirtualizationVirtualMachinesCreateCreated, error)

	VirtualizationVirtualMachinesDelete(params *VirtualizationVirtualMachinesDeleteParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*VirtualizationVirtualMachinesDeleteNoContent, error)

	VirtualizationVirtualMachinesList(params *VirtualizationVirtualMachinesListParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*VirtualizationVirtualMachinesListOK, error)

	VirtualizationVirtualMachinesPartialUpdate(params *VirtualizationVirtualMachinesPartialUpdateParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*VirtualizationVirtualMachinesPartialUpdateOK, error)

	VirtualizationVirtualMachinesRead(params *VirtualizationVirtualMachinesReadParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*VirtualizationVirtualMachinesReadOK, error)

	VirtualizationVirtualMachinesUpdate(params *VirtualizationVirtualMachinesUpdateParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*VirtualizationVirtualMachinesUpdateOK, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
  VirtualizationClusterGroupsBulkDelete virtualization cluster groups bulk delete API
*/
func (a *Client) VirtualizationClusterGroupsBulkDelete(params *VirtualizationClusterGroupsBulkDeleteParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*VirtualizationClusterGroupsBulkDeleteNoContent, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewVirtualizationClusterGroupsBulkDeleteParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "virtualization_cluster-groups_bulk_delete",
		Method:             "DELETE",
		PathPattern:        "/virtualization/cluster-groups/",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &VirtualizationClusterGroupsBulkDeleteReader{formats: a.formats},
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
	success, ok := result.(*VirtualizationClusterGroupsBulkDeleteNoContent)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for virtualization_cluster-groups_bulk_delete: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  VirtualizationClusterGroupsBulkPartialUpdate virtualization cluster groups bulk partial update API
*/
func (a *Client) VirtualizationClusterGroupsBulkPartialUpdate(params *VirtualizationClusterGroupsBulkPartialUpdateParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*VirtualizationClusterGroupsBulkPartialUpdateOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewVirtualizationClusterGroupsBulkPartialUpdateParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "virtualization_cluster-groups_bulk_partial_update",
		Method:             "PATCH",
		PathPattern:        "/virtualization/cluster-groups/",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &VirtualizationClusterGroupsBulkPartialUpdateReader{formats: a.formats},
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
	success, ok := result.(*VirtualizationClusterGroupsBulkPartialUpdateOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for virtualization_cluster-groups_bulk_partial_update: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  VirtualizationClusterGroupsBulkUpdate virtualization cluster groups bulk update API
*/
func (a *Client) VirtualizationClusterGroupsBulkUpdate(params *VirtualizationClusterGroupsBulkUpdateParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*VirtualizationClusterGroupsBulkUpdateOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewVirtualizationClusterGroupsBulkUpdateParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "virtualization_cluster-groups_bulk_update",
		Method:             "PUT",
		PathPattern:        "/virtualization/cluster-groups/",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &VirtualizationClusterGroupsBulkUpdateReader{formats: a.formats},
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
	success, ok := result.(*VirtualizationClusterGroupsBulkUpdateOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for virtualization_cluster-groups_bulk_update: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  VirtualizationClusterGroupsCreate virtualization cluster groups create API
*/
func (a *Client) VirtualizationClusterGroupsCreate(params *VirtualizationClusterGroupsCreateParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*VirtualizationClusterGroupsCreateCreated, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewVirtualizationClusterGroupsCreateParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "virtualization_cluster-groups_create",
		Method:             "POST",
		PathPattern:        "/virtualization/cluster-groups/",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &VirtualizationClusterGroupsCreateReader{formats: a.formats},
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
	success, ok := result.(*VirtualizationClusterGroupsCreateCreated)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for virtualization_cluster-groups_create: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  VirtualizationClusterGroupsDelete virtualization cluster groups delete API
*/
func (a *Client) VirtualizationClusterGroupsDelete(params *VirtualizationClusterGroupsDeleteParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*VirtualizationClusterGroupsDeleteNoContent, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewVirtualizationClusterGroupsDeleteParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "virtualization_cluster-groups_delete",
		Method:             "DELETE",
		PathPattern:        "/virtualization/cluster-groups/{id}/",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &VirtualizationClusterGroupsDeleteReader{formats: a.formats},
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
	success, ok := result.(*VirtualizationClusterGroupsDeleteNoContent)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for virtualization_cluster-groups_delete: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  VirtualizationClusterGroupsList virtualization cluster groups list API
*/
func (a *Client) VirtualizationClusterGroupsList(params *VirtualizationClusterGroupsListParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*VirtualizationClusterGroupsListOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewVirtualizationClusterGroupsListParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "virtualization_cluster-groups_list",
		Method:             "GET",
		PathPattern:        "/virtualization/cluster-groups/",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &VirtualizationClusterGroupsListReader{formats: a.formats},
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
	success, ok := result.(*VirtualizationClusterGroupsListOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for virtualization_cluster-groups_list: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  VirtualizationClusterGroupsPartialUpdate virtualization cluster groups partial update API
*/
func (a *Client) VirtualizationClusterGroupsPartialUpdate(params *VirtualizationClusterGroupsPartialUpdateParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*VirtualizationClusterGroupsPartialUpdateOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewVirtualizationClusterGroupsPartialUpdateParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "virtualization_cluster-groups_partial_update",
		Method:             "PATCH",
		PathPattern:        "/virtualization/cluster-groups/{id}/",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &VirtualizationClusterGroupsPartialUpdateReader{formats: a.formats},
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
	success, ok := result.(*VirtualizationClusterGroupsPartialUpdateOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for virtualization_cluster-groups_partial_update: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  VirtualizationClusterGroupsRead virtualization cluster groups read API
*/
func (a *Client) VirtualizationClusterGroupsRead(params *VirtualizationClusterGroupsReadParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*VirtualizationClusterGroupsReadOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewVirtualizationClusterGroupsReadParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "virtualization_cluster-groups_read",
		Method:             "GET",
		PathPattern:        "/virtualization/cluster-groups/{id}/",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &VirtualizationClusterGroupsReadReader{formats: a.formats},
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
	success, ok := result.(*VirtualizationClusterGroupsReadOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for virtualization_cluster-groups_read: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  VirtualizationClusterGroupsUpdate virtualization cluster groups update API
*/
func (a *Client) VirtualizationClusterGroupsUpdate(params *VirtualizationClusterGroupsUpdateParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*VirtualizationClusterGroupsUpdateOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewVirtualizationClusterGroupsUpdateParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "virtualization_cluster-groups_update",
		Method:             "PUT",
		PathPattern:        "/virtualization/cluster-groups/{id}/",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &VirtualizationClusterGroupsUpdateReader{formats: a.formats},
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
	success, ok := result.(*VirtualizationClusterGroupsUpdateOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for virtualization_cluster-groups_update: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  VirtualizationClusterTypesBulkDelete virtualization cluster types bulk delete API
*/
func (a *Client) VirtualizationClusterTypesBulkDelete(params *VirtualizationClusterTypesBulkDeleteParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*VirtualizationClusterTypesBulkDeleteNoContent, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewVirtualizationClusterTypesBulkDeleteParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "virtualization_cluster-types_bulk_delete",
		Method:             "DELETE",
		PathPattern:        "/virtualization/cluster-types/",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &VirtualizationClusterTypesBulkDeleteReader{formats: a.formats},
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
	success, ok := result.(*VirtualizationClusterTypesBulkDeleteNoContent)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for virtualization_cluster-types_bulk_delete: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  VirtualizationClusterTypesBulkPartialUpdate virtualization cluster types bulk partial update API
*/
func (a *Client) VirtualizationClusterTypesBulkPartialUpdate(params *VirtualizationClusterTypesBulkPartialUpdateParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*VirtualizationClusterTypesBulkPartialUpdateOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewVirtualizationClusterTypesBulkPartialUpdateParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "virtualization_cluster-types_bulk_partial_update",
		Method:             "PATCH",
		PathPattern:        "/virtualization/cluster-types/",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &VirtualizationClusterTypesBulkPartialUpdateReader{formats: a.formats},
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
	success, ok := result.(*VirtualizationClusterTypesBulkPartialUpdateOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for virtualization_cluster-types_bulk_partial_update: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  VirtualizationClusterTypesBulkUpdate virtualization cluster types bulk update API
*/
func (a *Client) VirtualizationClusterTypesBulkUpdate(params *VirtualizationClusterTypesBulkUpdateParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*VirtualizationClusterTypesBulkUpdateOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewVirtualizationClusterTypesBulkUpdateParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "virtualization_cluster-types_bulk_update",
		Method:             "PUT",
		PathPattern:        "/virtualization/cluster-types/",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &VirtualizationClusterTypesBulkUpdateReader{formats: a.formats},
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
	success, ok := result.(*VirtualizationClusterTypesBulkUpdateOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for virtualization_cluster-types_bulk_update: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  VirtualizationClusterTypesCreate virtualization cluster types create API
*/
func (a *Client) VirtualizationClusterTypesCreate(params *VirtualizationClusterTypesCreateParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*VirtualizationClusterTypesCreateCreated, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewVirtualizationClusterTypesCreateParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "virtualization_cluster-types_create",
		Method:             "POST",
		PathPattern:        "/virtualization/cluster-types/",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &VirtualizationClusterTypesCreateReader{formats: a.formats},
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
	success, ok := result.(*VirtualizationClusterTypesCreateCreated)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for virtualization_cluster-types_create: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  VirtualizationClusterTypesDelete virtualization cluster types delete API
*/
func (a *Client) VirtualizationClusterTypesDelete(params *VirtualizationClusterTypesDeleteParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*VirtualizationClusterTypesDeleteNoContent, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewVirtualizationClusterTypesDeleteParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "virtualization_cluster-types_delete",
		Method:             "DELETE",
		PathPattern:        "/virtualization/cluster-types/{id}/",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &VirtualizationClusterTypesDeleteReader{formats: a.formats},
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
	success, ok := result.(*VirtualizationClusterTypesDeleteNoContent)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for virtualization_cluster-types_delete: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  VirtualizationClusterTypesList virtualization cluster types list API
*/
func (a *Client) VirtualizationClusterTypesList(params *VirtualizationClusterTypesListParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*VirtualizationClusterTypesListOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewVirtualizationClusterTypesListParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "virtualization_cluster-types_list",
		Method:             "GET",
		PathPattern:        "/virtualization/cluster-types/",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &VirtualizationClusterTypesListReader{formats: a.formats},
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
	success, ok := result.(*VirtualizationClusterTypesListOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for virtualization_cluster-types_list: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  VirtualizationClusterTypesPartialUpdate virtualization cluster types partial update API
*/
func (a *Client) VirtualizationClusterTypesPartialUpdate(params *VirtualizationClusterTypesPartialUpdateParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*VirtualizationClusterTypesPartialUpdateOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewVirtualizationClusterTypesPartialUpdateParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "virtualization_cluster-types_partial_update",
		Method:             "PATCH",
		PathPattern:        "/virtualization/cluster-types/{id}/",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &VirtualizationClusterTypesPartialUpdateReader{formats: a.formats},
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
	success, ok := result.(*VirtualizationClusterTypesPartialUpdateOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for virtualization_cluster-types_partial_update: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  VirtualizationClusterTypesRead virtualization cluster types read API
*/
func (a *Client) VirtualizationClusterTypesRead(params *VirtualizationClusterTypesReadParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*VirtualizationClusterTypesReadOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewVirtualizationClusterTypesReadParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "virtualization_cluster-types_read",
		Method:             "GET",
		PathPattern:        "/virtualization/cluster-types/{id}/",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &VirtualizationClusterTypesReadReader{formats: a.formats},
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
	success, ok := result.(*VirtualizationClusterTypesReadOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for virtualization_cluster-types_read: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  VirtualizationClusterTypesUpdate virtualization cluster types update API
*/
func (a *Client) VirtualizationClusterTypesUpdate(params *VirtualizationClusterTypesUpdateParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*VirtualizationClusterTypesUpdateOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewVirtualizationClusterTypesUpdateParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "virtualization_cluster-types_update",
		Method:             "PUT",
		PathPattern:        "/virtualization/cluster-types/{id}/",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &VirtualizationClusterTypesUpdateReader{formats: a.formats},
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
	success, ok := result.(*VirtualizationClusterTypesUpdateOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for virtualization_cluster-types_update: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  VirtualizationClustersBulkDelete virtualization clusters bulk delete API
*/
func (a *Client) VirtualizationClustersBulkDelete(params *VirtualizationClustersBulkDeleteParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*VirtualizationClustersBulkDeleteNoContent, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewVirtualizationClustersBulkDeleteParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "virtualization_clusters_bulk_delete",
		Method:             "DELETE",
		PathPattern:        "/virtualization/clusters/",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &VirtualizationClustersBulkDeleteReader{formats: a.formats},
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
	success, ok := result.(*VirtualizationClustersBulkDeleteNoContent)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for virtualization_clusters_bulk_delete: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  VirtualizationClustersBulkPartialUpdate virtualization clusters bulk partial update API
*/
func (a *Client) VirtualizationClustersBulkPartialUpdate(params *VirtualizationClustersBulkPartialUpdateParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*VirtualizationClustersBulkPartialUpdateOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewVirtualizationClustersBulkPartialUpdateParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "virtualization_clusters_bulk_partial_update",
		Method:             "PATCH",
		PathPattern:        "/virtualization/clusters/",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &VirtualizationClustersBulkPartialUpdateReader{formats: a.formats},
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
	success, ok := result.(*VirtualizationClustersBulkPartialUpdateOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for virtualization_clusters_bulk_partial_update: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  VirtualizationClustersBulkUpdate virtualization clusters bulk update API
*/
func (a *Client) VirtualizationClustersBulkUpdate(params *VirtualizationClustersBulkUpdateParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*VirtualizationClustersBulkUpdateOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewVirtualizationClustersBulkUpdateParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "virtualization_clusters_bulk_update",
		Method:             "PUT",
		PathPattern:        "/virtualization/clusters/",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &VirtualizationClustersBulkUpdateReader{formats: a.formats},
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
	success, ok := result.(*VirtualizationClustersBulkUpdateOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for virtualization_clusters_bulk_update: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  VirtualizationClustersCreate virtualization clusters create API
*/
func (a *Client) VirtualizationClustersCreate(params *VirtualizationClustersCreateParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*VirtualizationClustersCreateCreated, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewVirtualizationClustersCreateParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "virtualization_clusters_create",
		Method:             "POST",
		PathPattern:        "/virtualization/clusters/",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &VirtualizationClustersCreateReader{formats: a.formats},
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
	success, ok := result.(*VirtualizationClustersCreateCreated)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for virtualization_clusters_create: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  VirtualizationClustersDelete virtualization clusters delete API
*/
func (a *Client) VirtualizationClustersDelete(params *VirtualizationClustersDeleteParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*VirtualizationClustersDeleteNoContent, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewVirtualizationClustersDeleteParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "virtualization_clusters_delete",
		Method:             "DELETE",
		PathPattern:        "/virtualization/clusters/{id}/",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &VirtualizationClustersDeleteReader{formats: a.formats},
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
	success, ok := result.(*VirtualizationClustersDeleteNoContent)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for virtualization_clusters_delete: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  VirtualizationClustersList virtualization clusters list API
*/
func (a *Client) VirtualizationClustersList(params *VirtualizationClustersListParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*VirtualizationClustersListOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewVirtualizationClustersListParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "virtualization_clusters_list",
		Method:             "GET",
		PathPattern:        "/virtualization/clusters/",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &VirtualizationClustersListReader{formats: a.formats},
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
	success, ok := result.(*VirtualizationClustersListOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for virtualization_clusters_list: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  VirtualizationClustersPartialUpdate virtualization clusters partial update API
*/
func (a *Client) VirtualizationClustersPartialUpdate(params *VirtualizationClustersPartialUpdateParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*VirtualizationClustersPartialUpdateOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewVirtualizationClustersPartialUpdateParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "virtualization_clusters_partial_update",
		Method:             "PATCH",
		PathPattern:        "/virtualization/clusters/{id}/",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &VirtualizationClustersPartialUpdateReader{formats: a.formats},
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
	success, ok := result.(*VirtualizationClustersPartialUpdateOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for virtualization_clusters_partial_update: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  VirtualizationClustersRead virtualization clusters read API
*/
func (a *Client) VirtualizationClustersRead(params *VirtualizationClustersReadParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*VirtualizationClustersReadOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewVirtualizationClustersReadParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "virtualization_clusters_read",
		Method:             "GET",
		PathPattern:        "/virtualization/clusters/{id}/",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &VirtualizationClustersReadReader{formats: a.formats},
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
	success, ok := result.(*VirtualizationClustersReadOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for virtualization_clusters_read: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  VirtualizationClustersUpdate virtualization clusters update API
*/
func (a *Client) VirtualizationClustersUpdate(params *VirtualizationClustersUpdateParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*VirtualizationClustersUpdateOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewVirtualizationClustersUpdateParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "virtualization_clusters_update",
		Method:             "PUT",
		PathPattern:        "/virtualization/clusters/{id}/",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &VirtualizationClustersUpdateReader{formats: a.formats},
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
	success, ok := result.(*VirtualizationClustersUpdateOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for virtualization_clusters_update: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  VirtualizationInterfacesBulkDelete virtualization interfaces bulk delete API
*/
func (a *Client) VirtualizationInterfacesBulkDelete(params *VirtualizationInterfacesBulkDeleteParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*VirtualizationInterfacesBulkDeleteNoContent, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewVirtualizationInterfacesBulkDeleteParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "virtualization_interfaces_bulk_delete",
		Method:             "DELETE",
		PathPattern:        "/virtualization/interfaces/",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &VirtualizationInterfacesBulkDeleteReader{formats: a.formats},
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
	success, ok := result.(*VirtualizationInterfacesBulkDeleteNoContent)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for virtualization_interfaces_bulk_delete: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  VirtualizationInterfacesBulkPartialUpdate virtualization interfaces bulk partial update API
*/
func (a *Client) VirtualizationInterfacesBulkPartialUpdate(params *VirtualizationInterfacesBulkPartialUpdateParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*VirtualizationInterfacesBulkPartialUpdateOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewVirtualizationInterfacesBulkPartialUpdateParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "virtualization_interfaces_bulk_partial_update",
		Method:             "PATCH",
		PathPattern:        "/virtualization/interfaces/",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &VirtualizationInterfacesBulkPartialUpdateReader{formats: a.formats},
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
	success, ok := result.(*VirtualizationInterfacesBulkPartialUpdateOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for virtualization_interfaces_bulk_partial_update: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  VirtualizationInterfacesBulkUpdate virtualization interfaces bulk update API
*/
func (a *Client) VirtualizationInterfacesBulkUpdate(params *VirtualizationInterfacesBulkUpdateParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*VirtualizationInterfacesBulkUpdateOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewVirtualizationInterfacesBulkUpdateParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "virtualization_interfaces_bulk_update",
		Method:             "PUT",
		PathPattern:        "/virtualization/interfaces/",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &VirtualizationInterfacesBulkUpdateReader{formats: a.formats},
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
	success, ok := result.(*VirtualizationInterfacesBulkUpdateOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for virtualization_interfaces_bulk_update: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  VirtualizationInterfacesCreate virtualization interfaces create API
*/
func (a *Client) VirtualizationInterfacesCreate(params *VirtualizationInterfacesCreateParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*VirtualizationInterfacesCreateCreated, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewVirtualizationInterfacesCreateParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "virtualization_interfaces_create",
		Method:             "POST",
		PathPattern:        "/virtualization/interfaces/",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &VirtualizationInterfacesCreateReader{formats: a.formats},
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
	success, ok := result.(*VirtualizationInterfacesCreateCreated)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for virtualization_interfaces_create: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  VirtualizationInterfacesDelete virtualization interfaces delete API
*/
func (a *Client) VirtualizationInterfacesDelete(params *VirtualizationInterfacesDeleteParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*VirtualizationInterfacesDeleteNoContent, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewVirtualizationInterfacesDeleteParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "virtualization_interfaces_delete",
		Method:             "DELETE",
		PathPattern:        "/virtualization/interfaces/{id}/",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &VirtualizationInterfacesDeleteReader{formats: a.formats},
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
	success, ok := result.(*VirtualizationInterfacesDeleteNoContent)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for virtualization_interfaces_delete: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  VirtualizationInterfacesList virtualization interfaces list API
*/
func (a *Client) VirtualizationInterfacesList(params *VirtualizationInterfacesListParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*VirtualizationInterfacesListOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewVirtualizationInterfacesListParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "virtualization_interfaces_list",
		Method:             "GET",
		PathPattern:        "/virtualization/interfaces/",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &VirtualizationInterfacesListReader{formats: a.formats},
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
	success, ok := result.(*VirtualizationInterfacesListOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for virtualization_interfaces_list: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  VirtualizationInterfacesPartialUpdate virtualization interfaces partial update API
*/
func (a *Client) VirtualizationInterfacesPartialUpdate(params *VirtualizationInterfacesPartialUpdateParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*VirtualizationInterfacesPartialUpdateOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewVirtualizationInterfacesPartialUpdateParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "virtualization_interfaces_partial_update",
		Method:             "PATCH",
		PathPattern:        "/virtualization/interfaces/{id}/",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &VirtualizationInterfacesPartialUpdateReader{formats: a.formats},
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
	success, ok := result.(*VirtualizationInterfacesPartialUpdateOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for virtualization_interfaces_partial_update: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  VirtualizationInterfacesRead virtualization interfaces read API
*/
func (a *Client) VirtualizationInterfacesRead(params *VirtualizationInterfacesReadParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*VirtualizationInterfacesReadOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewVirtualizationInterfacesReadParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "virtualization_interfaces_read",
		Method:             "GET",
		PathPattern:        "/virtualization/interfaces/{id}/",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &VirtualizationInterfacesReadReader{formats: a.formats},
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
	success, ok := result.(*VirtualizationInterfacesReadOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for virtualization_interfaces_read: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  VirtualizationInterfacesUpdate virtualization interfaces update API
*/
func (a *Client) VirtualizationInterfacesUpdate(params *VirtualizationInterfacesUpdateParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*VirtualizationInterfacesUpdateOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewVirtualizationInterfacesUpdateParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "virtualization_interfaces_update",
		Method:             "PUT",
		PathPattern:        "/virtualization/interfaces/{id}/",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &VirtualizationInterfacesUpdateReader{formats: a.formats},
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
	success, ok := result.(*VirtualizationInterfacesUpdateOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for virtualization_interfaces_update: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  VirtualizationVirtualMachinesBulkDelete virtualization virtual machines bulk delete API
*/
func (a *Client) VirtualizationVirtualMachinesBulkDelete(params *VirtualizationVirtualMachinesBulkDeleteParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*VirtualizationVirtualMachinesBulkDeleteNoContent, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewVirtualizationVirtualMachinesBulkDeleteParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "virtualization_virtual-machines_bulk_delete",
		Method:             "DELETE",
		PathPattern:        "/virtualization/virtual-machines/",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &VirtualizationVirtualMachinesBulkDeleteReader{formats: a.formats},
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
	success, ok := result.(*VirtualizationVirtualMachinesBulkDeleteNoContent)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for virtualization_virtual-machines_bulk_delete: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  VirtualizationVirtualMachinesBulkPartialUpdate virtualization virtual machines bulk partial update API
*/
func (a *Client) VirtualizationVirtualMachinesBulkPartialUpdate(params *VirtualizationVirtualMachinesBulkPartialUpdateParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*VirtualizationVirtualMachinesBulkPartialUpdateOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewVirtualizationVirtualMachinesBulkPartialUpdateParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "virtualization_virtual-machines_bulk_partial_update",
		Method:             "PATCH",
		PathPattern:        "/virtualization/virtual-machines/",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &VirtualizationVirtualMachinesBulkPartialUpdateReader{formats: a.formats},
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
	success, ok := result.(*VirtualizationVirtualMachinesBulkPartialUpdateOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for virtualization_virtual-machines_bulk_partial_update: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  VirtualizationVirtualMachinesBulkUpdate virtualization virtual machines bulk update API
*/
func (a *Client) VirtualizationVirtualMachinesBulkUpdate(params *VirtualizationVirtualMachinesBulkUpdateParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*VirtualizationVirtualMachinesBulkUpdateOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewVirtualizationVirtualMachinesBulkUpdateParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "virtualization_virtual-machines_bulk_update",
		Method:             "PUT",
		PathPattern:        "/virtualization/virtual-machines/",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &VirtualizationVirtualMachinesBulkUpdateReader{formats: a.formats},
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
	success, ok := result.(*VirtualizationVirtualMachinesBulkUpdateOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for virtualization_virtual-machines_bulk_update: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  VirtualizationVirtualMachinesCreate virtualization virtual machines create API
*/
func (a *Client) VirtualizationVirtualMachinesCreate(params *VirtualizationVirtualMachinesCreateParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*VirtualizationVirtualMachinesCreateCreated, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewVirtualizationVirtualMachinesCreateParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "virtualization_virtual-machines_create",
		Method:             "POST",
		PathPattern:        "/virtualization/virtual-machines/",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &VirtualizationVirtualMachinesCreateReader{formats: a.formats},
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
	success, ok := result.(*VirtualizationVirtualMachinesCreateCreated)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for virtualization_virtual-machines_create: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  VirtualizationVirtualMachinesDelete virtualization virtual machines delete API
*/
func (a *Client) VirtualizationVirtualMachinesDelete(params *VirtualizationVirtualMachinesDeleteParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*VirtualizationVirtualMachinesDeleteNoContent, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewVirtualizationVirtualMachinesDeleteParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "virtualization_virtual-machines_delete",
		Method:             "DELETE",
		PathPattern:        "/virtualization/virtual-machines/{id}/",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &VirtualizationVirtualMachinesDeleteReader{formats: a.formats},
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
	success, ok := result.(*VirtualizationVirtualMachinesDeleteNoContent)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for virtualization_virtual-machines_delete: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  VirtualizationVirtualMachinesList virtualization virtual machines list API
*/
func (a *Client) VirtualizationVirtualMachinesList(params *VirtualizationVirtualMachinesListParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*VirtualizationVirtualMachinesListOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewVirtualizationVirtualMachinesListParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "virtualization_virtual-machines_list",
		Method:             "GET",
		PathPattern:        "/virtualization/virtual-machines/",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &VirtualizationVirtualMachinesListReader{formats: a.formats},
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
	success, ok := result.(*VirtualizationVirtualMachinesListOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for virtualization_virtual-machines_list: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  VirtualizationVirtualMachinesPartialUpdate virtualization virtual machines partial update API
*/
func (a *Client) VirtualizationVirtualMachinesPartialUpdate(params *VirtualizationVirtualMachinesPartialUpdateParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*VirtualizationVirtualMachinesPartialUpdateOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewVirtualizationVirtualMachinesPartialUpdateParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "virtualization_virtual-machines_partial_update",
		Method:             "PATCH",
		PathPattern:        "/virtualization/virtual-machines/{id}/",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &VirtualizationVirtualMachinesPartialUpdateReader{formats: a.formats},
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
	success, ok := result.(*VirtualizationVirtualMachinesPartialUpdateOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for virtualization_virtual-machines_partial_update: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  VirtualizationVirtualMachinesRead virtualization virtual machines read API
*/
func (a *Client) VirtualizationVirtualMachinesRead(params *VirtualizationVirtualMachinesReadParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*VirtualizationVirtualMachinesReadOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewVirtualizationVirtualMachinesReadParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "virtualization_virtual-machines_read",
		Method:             "GET",
		PathPattern:        "/virtualization/virtual-machines/{id}/",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &VirtualizationVirtualMachinesReadReader{formats: a.formats},
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
	success, ok := result.(*VirtualizationVirtualMachinesReadOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for virtualization_virtual-machines_read: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  VirtualizationVirtualMachinesUpdate virtualization virtual machines update API
*/
func (a *Client) VirtualizationVirtualMachinesUpdate(params *VirtualizationVirtualMachinesUpdateParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*VirtualizationVirtualMachinesUpdateOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewVirtualizationVirtualMachinesUpdateParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "virtualization_virtual-machines_update",
		Method:             "PUT",
		PathPattern:        "/virtualization/virtual-machines/{id}/",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &VirtualizationVirtualMachinesUpdateReader{formats: a.formats},
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
	success, ok := result.(*VirtualizationVirtualMachinesUpdateOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for virtualization_virtual-machines_update: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
