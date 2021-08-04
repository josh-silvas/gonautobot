package ipam

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// New creates a new ipam API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

/*
Client for ipam API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientOption is the option for Client methods
type ClientOption func(*runtime.ClientOperation)

// ClientService is the interface for Client methods
type ClientService interface {
	IpamAggregatesBulkDelete(params *IpamAggregatesBulkDeleteParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*IpamAggregatesBulkDeleteNoContent, error)

	IpamAggregatesBulkPartialUpdate(params *IpamAggregatesBulkPartialUpdateParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*IpamAggregatesBulkPartialUpdateOK, error)

	IpamAggregatesBulkUpdate(params *IpamAggregatesBulkUpdateParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*IpamAggregatesBulkUpdateOK, error)

	IpamAggregatesCreate(params *IpamAggregatesCreateParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*IpamAggregatesCreateCreated, error)

	IpamAggregatesDelete(params *IpamAggregatesDeleteParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*IpamAggregatesDeleteNoContent, error)

	IpamAggregatesList(params *IpamAggregatesListParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*IpamAggregatesListOK, error)

	IpamAggregatesPartialUpdate(params *IpamAggregatesPartialUpdateParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*IpamAggregatesPartialUpdateOK, error)

	IpamAggregatesRead(params *IpamAggregatesReadParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*IpamAggregatesReadOK, error)

	IpamAggregatesUpdate(params *IpamAggregatesUpdateParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*IpamAggregatesUpdateOK, error)

	IpamIPAddressesBulkDelete(params *IpamIPAddressesBulkDeleteParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*IpamIPAddressesBulkDeleteNoContent, error)

	IpamIPAddressesBulkPartialUpdate(params *IpamIPAddressesBulkPartialUpdateParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*IpamIPAddressesBulkPartialUpdateOK, error)

	IpamIPAddressesBulkUpdate(params *IpamIPAddressesBulkUpdateParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*IpamIPAddressesBulkUpdateOK, error)

	IpamIPAddressesCreate(params *IpamIPAddressesCreateParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*IpamIPAddressesCreateCreated, error)

	IpamIPAddressesDelete(params *IpamIPAddressesDeleteParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*IpamIPAddressesDeleteNoContent, error)

	IpamIPAddressesList(params *IpamIPAddressesListParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*IpamIPAddressesListOK, error)

	IpamIPAddressesPartialUpdate(params *IpamIPAddressesPartialUpdateParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*IpamIPAddressesPartialUpdateOK, error)

	IpamIPAddressesRead(params *IpamIPAddressesReadParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*IpamIPAddressesReadOK, error)

	IpamIPAddressesUpdate(params *IpamIPAddressesUpdateParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*IpamIPAddressesUpdateOK, error)

	IpamPrefixesAvailableIpsCreate(params *IpamPrefixesAvailableIpsCreateParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*IpamPrefixesAvailableIpsCreateCreated, error)

	IpamPrefixesAvailableIpsRead(params *IpamPrefixesAvailableIpsReadParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*IpamPrefixesAvailableIpsReadOK, error)

	IpamPrefixesAvailablePrefixesCreate(params *IpamPrefixesAvailablePrefixesCreateParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*IpamPrefixesAvailablePrefixesCreateCreated, error)

	IpamPrefixesAvailablePrefixesRead(params *IpamPrefixesAvailablePrefixesReadParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*IpamPrefixesAvailablePrefixesReadOK, error)

	IpamPrefixesBulkDelete(params *IpamPrefixesBulkDeleteParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*IpamPrefixesBulkDeleteNoContent, error)

	IpamPrefixesBulkPartialUpdate(params *IpamPrefixesBulkPartialUpdateParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*IpamPrefixesBulkPartialUpdateOK, error)

	IpamPrefixesBulkUpdate(params *IpamPrefixesBulkUpdateParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*IpamPrefixesBulkUpdateOK, error)

	IpamPrefixesCreate(params *IpamPrefixesCreateParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*IpamPrefixesCreateCreated, error)

	IpamPrefixesDelete(params *IpamPrefixesDeleteParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*IpamPrefixesDeleteNoContent, error)

	IpamPrefixesList(params *IpamPrefixesListParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*IpamPrefixesListOK, error)

	IpamPrefixesPartialUpdate(params *IpamPrefixesPartialUpdateParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*IpamPrefixesPartialUpdateOK, error)

	IpamPrefixesRead(params *IpamPrefixesReadParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*IpamPrefixesReadOK, error)

	IpamPrefixesUpdate(params *IpamPrefixesUpdateParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*IpamPrefixesUpdateOK, error)

	IpamRirsBulkDelete(params *IpamRirsBulkDeleteParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*IpamRirsBulkDeleteNoContent, error)

	IpamRirsBulkPartialUpdate(params *IpamRirsBulkPartialUpdateParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*IpamRirsBulkPartialUpdateOK, error)

	IpamRirsBulkUpdate(params *IpamRirsBulkUpdateParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*IpamRirsBulkUpdateOK, error)

	IpamRirsCreate(params *IpamRirsCreateParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*IpamRirsCreateCreated, error)

	IpamRirsDelete(params *IpamRirsDeleteParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*IpamRirsDeleteNoContent, error)

	IpamRirsList(params *IpamRirsListParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*IpamRirsListOK, error)

	IpamRirsPartialUpdate(params *IpamRirsPartialUpdateParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*IpamRirsPartialUpdateOK, error)

	IpamRirsRead(params *IpamRirsReadParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*IpamRirsReadOK, error)

	IpamRirsUpdate(params *IpamRirsUpdateParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*IpamRirsUpdateOK, error)

	IpamRolesBulkDelete(params *IpamRolesBulkDeleteParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*IpamRolesBulkDeleteNoContent, error)

	IpamRolesBulkPartialUpdate(params *IpamRolesBulkPartialUpdateParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*IpamRolesBulkPartialUpdateOK, error)

	IpamRolesBulkUpdate(params *IpamRolesBulkUpdateParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*IpamRolesBulkUpdateOK, error)

	IpamRolesCreate(params *IpamRolesCreateParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*IpamRolesCreateCreated, error)

	IpamRolesDelete(params *IpamRolesDeleteParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*IpamRolesDeleteNoContent, error)

	IpamRolesList(params *IpamRolesListParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*IpamRolesListOK, error)

	IpamRolesPartialUpdate(params *IpamRolesPartialUpdateParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*IpamRolesPartialUpdateOK, error)

	IpamRolesRead(params *IpamRolesReadParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*IpamRolesReadOK, error)

	IpamRolesUpdate(params *IpamRolesUpdateParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*IpamRolesUpdateOK, error)

	IpamRouteTargetsBulkDelete(params *IpamRouteTargetsBulkDeleteParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*IpamRouteTargetsBulkDeleteNoContent, error)

	IpamRouteTargetsBulkPartialUpdate(params *IpamRouteTargetsBulkPartialUpdateParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*IpamRouteTargetsBulkPartialUpdateOK, error)

	IpamRouteTargetsBulkUpdate(params *IpamRouteTargetsBulkUpdateParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*IpamRouteTargetsBulkUpdateOK, error)

	IpamRouteTargetsCreate(params *IpamRouteTargetsCreateParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*IpamRouteTargetsCreateCreated, error)

	IpamRouteTargetsDelete(params *IpamRouteTargetsDeleteParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*IpamRouteTargetsDeleteNoContent, error)

	IpamRouteTargetsList(params *IpamRouteTargetsListParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*IpamRouteTargetsListOK, error)

	IpamRouteTargetsPartialUpdate(params *IpamRouteTargetsPartialUpdateParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*IpamRouteTargetsPartialUpdateOK, error)

	IpamRouteTargetsRead(params *IpamRouteTargetsReadParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*IpamRouteTargetsReadOK, error)

	IpamRouteTargetsUpdate(params *IpamRouteTargetsUpdateParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*IpamRouteTargetsUpdateOK, error)

	IpamServicesBulkDelete(params *IpamServicesBulkDeleteParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*IpamServicesBulkDeleteNoContent, error)

	IpamServicesBulkPartialUpdate(params *IpamServicesBulkPartialUpdateParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*IpamServicesBulkPartialUpdateOK, error)

	IpamServicesBulkUpdate(params *IpamServicesBulkUpdateParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*IpamServicesBulkUpdateOK, error)

	IpamServicesCreate(params *IpamServicesCreateParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*IpamServicesCreateCreated, error)

	IpamServicesDelete(params *IpamServicesDeleteParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*IpamServicesDeleteNoContent, error)

	IpamServicesList(params *IpamServicesListParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*IpamServicesListOK, error)

	IpamServicesPartialUpdate(params *IpamServicesPartialUpdateParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*IpamServicesPartialUpdateOK, error)

	IpamServicesRead(params *IpamServicesReadParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*IpamServicesReadOK, error)

	IpamServicesUpdate(params *IpamServicesUpdateParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*IpamServicesUpdateOK, error)

	IpamVlanGroupsBulkDelete(params *IpamVlanGroupsBulkDeleteParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*IpamVlanGroupsBulkDeleteNoContent, error)

	IpamVlanGroupsBulkPartialUpdate(params *IpamVlanGroupsBulkPartialUpdateParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*IpamVlanGroupsBulkPartialUpdateOK, error)

	IpamVlanGroupsBulkUpdate(params *IpamVlanGroupsBulkUpdateParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*IpamVlanGroupsBulkUpdateOK, error)

	IpamVlanGroupsCreate(params *IpamVlanGroupsCreateParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*IpamVlanGroupsCreateCreated, error)

	IpamVlanGroupsDelete(params *IpamVlanGroupsDeleteParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*IpamVlanGroupsDeleteNoContent, error)

	IpamVlanGroupsList(params *IpamVlanGroupsListParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*IpamVlanGroupsListOK, error)

	IpamVlanGroupsPartialUpdate(params *IpamVlanGroupsPartialUpdateParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*IpamVlanGroupsPartialUpdateOK, error)

	IpamVlanGroupsRead(params *IpamVlanGroupsReadParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*IpamVlanGroupsReadOK, error)

	IpamVlanGroupsUpdate(params *IpamVlanGroupsUpdateParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*IpamVlanGroupsUpdateOK, error)

	IpamVlansBulkDelete(params *IpamVlansBulkDeleteParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*IpamVlansBulkDeleteNoContent, error)

	IpamVlansBulkPartialUpdate(params *IpamVlansBulkPartialUpdateParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*IpamVlansBulkPartialUpdateOK, error)

	IpamVlansBulkUpdate(params *IpamVlansBulkUpdateParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*IpamVlansBulkUpdateOK, error)

	IpamVlansCreate(params *IpamVlansCreateParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*IpamVlansCreateCreated, error)

	IpamVlansDelete(params *IpamVlansDeleteParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*IpamVlansDeleteNoContent, error)

	IpamVlansList(params *IpamVlansListParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*IpamVlansListOK, error)

	IpamVlansPartialUpdate(params *IpamVlansPartialUpdateParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*IpamVlansPartialUpdateOK, error)

	IpamVlansRead(params *IpamVlansReadParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*IpamVlansReadOK, error)

	IpamVlansUpdate(params *IpamVlansUpdateParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*IpamVlansUpdateOK, error)

	IpamVrfsBulkDelete(params *IpamVrfsBulkDeleteParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*IpamVrfsBulkDeleteNoContent, error)

	IpamVrfsBulkPartialUpdate(params *IpamVrfsBulkPartialUpdateParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*IpamVrfsBulkPartialUpdateOK, error)

	IpamVrfsBulkUpdate(params *IpamVrfsBulkUpdateParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*IpamVrfsBulkUpdateOK, error)

	IpamVrfsCreate(params *IpamVrfsCreateParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*IpamVrfsCreateCreated, error)

	IpamVrfsDelete(params *IpamVrfsDeleteParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*IpamVrfsDeleteNoContent, error)

	IpamVrfsList(params *IpamVrfsListParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*IpamVrfsListOK, error)

	IpamVrfsPartialUpdate(params *IpamVrfsPartialUpdateParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*IpamVrfsPartialUpdateOK, error)

	IpamVrfsRead(params *IpamVrfsReadParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*IpamVrfsReadOK, error)

	IpamVrfsUpdate(params *IpamVrfsUpdateParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*IpamVrfsUpdateOK, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
  IpamAggregatesBulkDelete ipam aggregates bulk delete API
*/
func (a *Client) IpamAggregatesBulkDelete(params *IpamAggregatesBulkDeleteParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*IpamAggregatesBulkDeleteNoContent, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewIpamAggregatesBulkDeleteParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "ipam_aggregates_bulk_delete",
		Method:             "DELETE",
		PathPattern:        "/ipam/aggregates/",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &IpamAggregatesBulkDeleteReader{formats: a.formats},
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
	success, ok := result.(*IpamAggregatesBulkDeleteNoContent)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for ipam_aggregates_bulk_delete: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  IpamAggregatesBulkPartialUpdate ipam aggregates bulk partial update API
*/
func (a *Client) IpamAggregatesBulkPartialUpdate(params *IpamAggregatesBulkPartialUpdateParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*IpamAggregatesBulkPartialUpdateOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewIpamAggregatesBulkPartialUpdateParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "ipam_aggregates_bulk_partial_update",
		Method:             "PATCH",
		PathPattern:        "/ipam/aggregates/",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &IpamAggregatesBulkPartialUpdateReader{formats: a.formats},
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
	success, ok := result.(*IpamAggregatesBulkPartialUpdateOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for ipam_aggregates_bulk_partial_update: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  IpamAggregatesBulkUpdate ipam aggregates bulk update API
*/
func (a *Client) IpamAggregatesBulkUpdate(params *IpamAggregatesBulkUpdateParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*IpamAggregatesBulkUpdateOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewIpamAggregatesBulkUpdateParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "ipam_aggregates_bulk_update",
		Method:             "PUT",
		PathPattern:        "/ipam/aggregates/",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &IpamAggregatesBulkUpdateReader{formats: a.formats},
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
	success, ok := result.(*IpamAggregatesBulkUpdateOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for ipam_aggregates_bulk_update: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  IpamAggregatesCreate ipam aggregates create API
*/
func (a *Client) IpamAggregatesCreate(params *IpamAggregatesCreateParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*IpamAggregatesCreateCreated, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewIpamAggregatesCreateParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "ipam_aggregates_create",
		Method:             "POST",
		PathPattern:        "/ipam/aggregates/",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &IpamAggregatesCreateReader{formats: a.formats},
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
	success, ok := result.(*IpamAggregatesCreateCreated)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for ipam_aggregates_create: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  IpamAggregatesDelete ipam aggregates delete API
*/
func (a *Client) IpamAggregatesDelete(params *IpamAggregatesDeleteParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*IpamAggregatesDeleteNoContent, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewIpamAggregatesDeleteParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "ipam_aggregates_delete",
		Method:             "DELETE",
		PathPattern:        "/ipam/aggregates/{id}/",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &IpamAggregatesDeleteReader{formats: a.formats},
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
	success, ok := result.(*IpamAggregatesDeleteNoContent)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for ipam_aggregates_delete: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  IpamAggregatesList ipam aggregates list API
*/
func (a *Client) IpamAggregatesList(params *IpamAggregatesListParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*IpamAggregatesListOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewIpamAggregatesListParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "ipam_aggregates_list",
		Method:             "GET",
		PathPattern:        "/ipam/aggregates/",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &IpamAggregatesListReader{formats: a.formats},
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
	success, ok := result.(*IpamAggregatesListOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for ipam_aggregates_list: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  IpamAggregatesPartialUpdate ipam aggregates partial update API
*/
func (a *Client) IpamAggregatesPartialUpdate(params *IpamAggregatesPartialUpdateParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*IpamAggregatesPartialUpdateOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewIpamAggregatesPartialUpdateParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "ipam_aggregates_partial_update",
		Method:             "PATCH",
		PathPattern:        "/ipam/aggregates/{id}/",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &IpamAggregatesPartialUpdateReader{formats: a.formats},
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
	success, ok := result.(*IpamAggregatesPartialUpdateOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for ipam_aggregates_partial_update: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  IpamAggregatesRead ipam aggregates read API
*/
func (a *Client) IpamAggregatesRead(params *IpamAggregatesReadParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*IpamAggregatesReadOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewIpamAggregatesReadParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "ipam_aggregates_read",
		Method:             "GET",
		PathPattern:        "/ipam/aggregates/{id}/",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &IpamAggregatesReadReader{formats: a.formats},
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
	success, ok := result.(*IpamAggregatesReadOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for ipam_aggregates_read: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  IpamAggregatesUpdate ipam aggregates update API
*/
func (a *Client) IpamAggregatesUpdate(params *IpamAggregatesUpdateParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*IpamAggregatesUpdateOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewIpamAggregatesUpdateParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "ipam_aggregates_update",
		Method:             "PUT",
		PathPattern:        "/ipam/aggregates/{id}/",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &IpamAggregatesUpdateReader{formats: a.formats},
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
	success, ok := result.(*IpamAggregatesUpdateOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for ipam_aggregates_update: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  IpamIPAddressesBulkDelete ipam ip addresses bulk delete API
*/
func (a *Client) IpamIPAddressesBulkDelete(params *IpamIPAddressesBulkDeleteParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*IpamIPAddressesBulkDeleteNoContent, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewIpamIPAddressesBulkDeleteParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "ipam_ip-addresses_bulk_delete",
		Method:             "DELETE",
		PathPattern:        "/ipam/ip-addresses/",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &IpamIPAddressesBulkDeleteReader{formats: a.formats},
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
	success, ok := result.(*IpamIPAddressesBulkDeleteNoContent)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for ipam_ip-addresses_bulk_delete: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  IpamIPAddressesBulkPartialUpdate ipam ip addresses bulk partial update API
*/
func (a *Client) IpamIPAddressesBulkPartialUpdate(params *IpamIPAddressesBulkPartialUpdateParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*IpamIPAddressesBulkPartialUpdateOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewIpamIPAddressesBulkPartialUpdateParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "ipam_ip-addresses_bulk_partial_update",
		Method:             "PATCH",
		PathPattern:        "/ipam/ip-addresses/",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &IpamIPAddressesBulkPartialUpdateReader{formats: a.formats},
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
	success, ok := result.(*IpamIPAddressesBulkPartialUpdateOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for ipam_ip-addresses_bulk_partial_update: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  IpamIPAddressesBulkUpdate ipam ip addresses bulk update API
*/
func (a *Client) IpamIPAddressesBulkUpdate(params *IpamIPAddressesBulkUpdateParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*IpamIPAddressesBulkUpdateOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewIpamIPAddressesBulkUpdateParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "ipam_ip-addresses_bulk_update",
		Method:             "PUT",
		PathPattern:        "/ipam/ip-addresses/",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &IpamIPAddressesBulkUpdateReader{formats: a.formats},
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
	success, ok := result.(*IpamIPAddressesBulkUpdateOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for ipam_ip-addresses_bulk_update: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  IpamIPAddressesCreate ipam ip addresses create API
*/
func (a *Client) IpamIPAddressesCreate(params *IpamIPAddressesCreateParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*IpamIPAddressesCreateCreated, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewIpamIPAddressesCreateParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "ipam_ip-addresses_create",
		Method:             "POST",
		PathPattern:        "/ipam/ip-addresses/",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &IpamIPAddressesCreateReader{formats: a.formats},
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
	success, ok := result.(*IpamIPAddressesCreateCreated)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for ipam_ip-addresses_create: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  IpamIPAddressesDelete ipam ip addresses delete API
*/
func (a *Client) IpamIPAddressesDelete(params *IpamIPAddressesDeleteParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*IpamIPAddressesDeleteNoContent, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewIpamIPAddressesDeleteParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "ipam_ip-addresses_delete",
		Method:             "DELETE",
		PathPattern:        "/ipam/ip-addresses/{id}/",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &IpamIPAddressesDeleteReader{formats: a.formats},
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
	success, ok := result.(*IpamIPAddressesDeleteNoContent)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for ipam_ip-addresses_delete: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  IpamIPAddressesList ipam ip addresses list API
*/
func (a *Client) IpamIPAddressesList(params *IpamIPAddressesListParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*IpamIPAddressesListOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewIpamIPAddressesListParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "ipam_ip-addresses_list",
		Method:             "GET",
		PathPattern:        "/ipam/ip-addresses/",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &IpamIPAddressesListReader{formats: a.formats},
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
	success, ok := result.(*IpamIPAddressesListOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for ipam_ip-addresses_list: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  IpamIPAddressesPartialUpdate ipam ip addresses partial update API
*/
func (a *Client) IpamIPAddressesPartialUpdate(params *IpamIPAddressesPartialUpdateParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*IpamIPAddressesPartialUpdateOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewIpamIPAddressesPartialUpdateParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "ipam_ip-addresses_partial_update",
		Method:             "PATCH",
		PathPattern:        "/ipam/ip-addresses/{id}/",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &IpamIPAddressesPartialUpdateReader{formats: a.formats},
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
	success, ok := result.(*IpamIPAddressesPartialUpdateOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for ipam_ip-addresses_partial_update: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  IpamIPAddressesRead ipam ip addresses read API
*/
func (a *Client) IpamIPAddressesRead(params *IpamIPAddressesReadParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*IpamIPAddressesReadOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewIpamIPAddressesReadParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "ipam_ip-addresses_read",
		Method:             "GET",
		PathPattern:        "/ipam/ip-addresses/{id}/",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &IpamIPAddressesReadReader{formats: a.formats},
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
	success, ok := result.(*IpamIPAddressesReadOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for ipam_ip-addresses_read: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  IpamIPAddressesUpdate ipam ip addresses update API
*/
func (a *Client) IpamIPAddressesUpdate(params *IpamIPAddressesUpdateParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*IpamIPAddressesUpdateOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewIpamIPAddressesUpdateParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "ipam_ip-addresses_update",
		Method:             "PUT",
		PathPattern:        "/ipam/ip-addresses/{id}/",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &IpamIPAddressesUpdateReader{formats: a.formats},
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
	success, ok := result.(*IpamIPAddressesUpdateOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for ipam_ip-addresses_update: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  IpamPrefixesAvailableIpsCreate A convenience method for returning available IP addresses within a prefix. By default, the number of IPs
returned will be equivalent to PAGINATE_COUNT. An arbitrary limit (up to MAX_PAGE_SIZE, if set) may be passed,
however results will not be paginated.

The advisory lock decorator uses a PostgreSQL advisory lock to prevent this API from being
invoked in parallel, which results in a race condition where multiple insertions can occur.
*/
func (a *Client) IpamPrefixesAvailableIpsCreate(params *IpamPrefixesAvailableIpsCreateParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*IpamPrefixesAvailableIpsCreateCreated, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewIpamPrefixesAvailableIpsCreateParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "ipam_prefixes_available-ips_create",
		Method:             "POST",
		PathPattern:        "/ipam/prefixes/{id}/available-ips/",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &IpamPrefixesAvailableIpsCreateReader{formats: a.formats},
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
	success, ok := result.(*IpamPrefixesAvailableIpsCreateCreated)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for ipam_prefixes_available-ips_create: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  IpamPrefixesAvailableIpsRead A convenience method for returning available IP addresses within a prefix. By default, the number of IPs
returned will be equivalent to PAGINATE_COUNT. An arbitrary limit (up to MAX_PAGE_SIZE, if set) may be passed,
however results will not be paginated.

The advisory lock decorator uses a PostgreSQL advisory lock to prevent this API from being
invoked in parallel, which results in a race condition where multiple insertions can occur.
*/
func (a *Client) IpamPrefixesAvailableIpsRead(params *IpamPrefixesAvailableIpsReadParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*IpamPrefixesAvailableIpsReadOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewIpamPrefixesAvailableIpsReadParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "ipam_prefixes_available-ips_read",
		Method:             "GET",
		PathPattern:        "/ipam/prefixes/{id}/available-ips/",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &IpamPrefixesAvailableIpsReadReader{formats: a.formats},
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
	success, ok := result.(*IpamPrefixesAvailableIpsReadOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for ipam_prefixes_available-ips_read: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  IpamPrefixesAvailablePrefixesCreate as convenience method for returning available child prefixes within a parent

  The advisory lock decorator uses a PostgreSQL advisory lock to prevent this API from being
invoked in parallel, which results in a race condition where multiple insertions can occur.
*/
func (a *Client) IpamPrefixesAvailablePrefixesCreate(params *IpamPrefixesAvailablePrefixesCreateParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*IpamPrefixesAvailablePrefixesCreateCreated, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewIpamPrefixesAvailablePrefixesCreateParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "ipam_prefixes_available-prefixes_create",
		Method:             "POST",
		PathPattern:        "/ipam/prefixes/{id}/available-prefixes/",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &IpamPrefixesAvailablePrefixesCreateReader{formats: a.formats},
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
	success, ok := result.(*IpamPrefixesAvailablePrefixesCreateCreated)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for ipam_prefixes_available-prefixes_create: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  IpamPrefixesAvailablePrefixesRead as convenience method for returning available child prefixes within a parent

  The advisory lock decorator uses a PostgreSQL advisory lock to prevent this API from being
invoked in parallel, which results in a race condition where multiple insertions can occur.
*/
func (a *Client) IpamPrefixesAvailablePrefixesRead(params *IpamPrefixesAvailablePrefixesReadParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*IpamPrefixesAvailablePrefixesReadOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewIpamPrefixesAvailablePrefixesReadParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "ipam_prefixes_available-prefixes_read",
		Method:             "GET",
		PathPattern:        "/ipam/prefixes/{id}/available-prefixes/",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &IpamPrefixesAvailablePrefixesReadReader{formats: a.formats},
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
	success, ok := result.(*IpamPrefixesAvailablePrefixesReadOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for ipam_prefixes_available-prefixes_read: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  IpamPrefixesBulkDelete ipam prefixes bulk delete API
*/
func (a *Client) IpamPrefixesBulkDelete(params *IpamPrefixesBulkDeleteParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*IpamPrefixesBulkDeleteNoContent, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewIpamPrefixesBulkDeleteParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "ipam_prefixes_bulk_delete",
		Method:             "DELETE",
		PathPattern:        "/ipam/prefixes/",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &IpamPrefixesBulkDeleteReader{formats: a.formats},
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
	success, ok := result.(*IpamPrefixesBulkDeleteNoContent)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for ipam_prefixes_bulk_delete: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  IpamPrefixesBulkPartialUpdate ipam prefixes bulk partial update API
*/
func (a *Client) IpamPrefixesBulkPartialUpdate(params *IpamPrefixesBulkPartialUpdateParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*IpamPrefixesBulkPartialUpdateOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewIpamPrefixesBulkPartialUpdateParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "ipam_prefixes_bulk_partial_update",
		Method:             "PATCH",
		PathPattern:        "/ipam/prefixes/",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &IpamPrefixesBulkPartialUpdateReader{formats: a.formats},
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
	success, ok := result.(*IpamPrefixesBulkPartialUpdateOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for ipam_prefixes_bulk_partial_update: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  IpamPrefixesBulkUpdate ipam prefixes bulk update API
*/
func (a *Client) IpamPrefixesBulkUpdate(params *IpamPrefixesBulkUpdateParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*IpamPrefixesBulkUpdateOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewIpamPrefixesBulkUpdateParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "ipam_prefixes_bulk_update",
		Method:             "PUT",
		PathPattern:        "/ipam/prefixes/",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &IpamPrefixesBulkUpdateReader{formats: a.formats},
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
	success, ok := result.(*IpamPrefixesBulkUpdateOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for ipam_prefixes_bulk_update: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  IpamPrefixesCreate ipam prefixes create API
*/
func (a *Client) IpamPrefixesCreate(params *IpamPrefixesCreateParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*IpamPrefixesCreateCreated, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewIpamPrefixesCreateParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "ipam_prefixes_create",
		Method:             "POST",
		PathPattern:        "/ipam/prefixes/",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &IpamPrefixesCreateReader{formats: a.formats},
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
	success, ok := result.(*IpamPrefixesCreateCreated)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for ipam_prefixes_create: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  IpamPrefixesDelete ipam prefixes delete API
*/
func (a *Client) IpamPrefixesDelete(params *IpamPrefixesDeleteParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*IpamPrefixesDeleteNoContent, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewIpamPrefixesDeleteParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "ipam_prefixes_delete",
		Method:             "DELETE",
		PathPattern:        "/ipam/prefixes/{id}/",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &IpamPrefixesDeleteReader{formats: a.formats},
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
	success, ok := result.(*IpamPrefixesDeleteNoContent)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for ipam_prefixes_delete: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  IpamPrefixesList ipam prefixes list API
*/
func (a *Client) IpamPrefixesList(params *IpamPrefixesListParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*IpamPrefixesListOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewIpamPrefixesListParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "ipam_prefixes_list",
		Method:             "GET",
		PathPattern:        "/ipam/prefixes/",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &IpamPrefixesListReader{formats: a.formats},
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
	success, ok := result.(*IpamPrefixesListOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for ipam_prefixes_list: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  IpamPrefixesPartialUpdate ipam prefixes partial update API
*/
func (a *Client) IpamPrefixesPartialUpdate(params *IpamPrefixesPartialUpdateParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*IpamPrefixesPartialUpdateOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewIpamPrefixesPartialUpdateParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "ipam_prefixes_partial_update",
		Method:             "PATCH",
		PathPattern:        "/ipam/prefixes/{id}/",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &IpamPrefixesPartialUpdateReader{formats: a.formats},
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
	success, ok := result.(*IpamPrefixesPartialUpdateOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for ipam_prefixes_partial_update: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  IpamPrefixesRead ipam prefixes read API
*/
func (a *Client) IpamPrefixesRead(params *IpamPrefixesReadParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*IpamPrefixesReadOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewIpamPrefixesReadParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "ipam_prefixes_read",
		Method:             "GET",
		PathPattern:        "/ipam/prefixes/{id}/",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &IpamPrefixesReadReader{formats: a.formats},
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
	success, ok := result.(*IpamPrefixesReadOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for ipam_prefixes_read: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  IpamPrefixesUpdate ipam prefixes update API
*/
func (a *Client) IpamPrefixesUpdate(params *IpamPrefixesUpdateParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*IpamPrefixesUpdateOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewIpamPrefixesUpdateParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "ipam_prefixes_update",
		Method:             "PUT",
		PathPattern:        "/ipam/prefixes/{id}/",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &IpamPrefixesUpdateReader{formats: a.formats},
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
	success, ok := result.(*IpamPrefixesUpdateOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for ipam_prefixes_update: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  IpamRirsBulkDelete ipam rirs bulk delete API
*/
func (a *Client) IpamRirsBulkDelete(params *IpamRirsBulkDeleteParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*IpamRirsBulkDeleteNoContent, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewIpamRirsBulkDeleteParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "ipam_rirs_bulk_delete",
		Method:             "DELETE",
		PathPattern:        "/ipam/rirs/",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &IpamRirsBulkDeleteReader{formats: a.formats},
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
	success, ok := result.(*IpamRirsBulkDeleteNoContent)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for ipam_rirs_bulk_delete: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  IpamRirsBulkPartialUpdate ipam rirs bulk partial update API
*/
func (a *Client) IpamRirsBulkPartialUpdate(params *IpamRirsBulkPartialUpdateParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*IpamRirsBulkPartialUpdateOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewIpamRirsBulkPartialUpdateParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "ipam_rirs_bulk_partial_update",
		Method:             "PATCH",
		PathPattern:        "/ipam/rirs/",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &IpamRirsBulkPartialUpdateReader{formats: a.formats},
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
	success, ok := result.(*IpamRirsBulkPartialUpdateOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for ipam_rirs_bulk_partial_update: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  IpamRirsBulkUpdate ipam rirs bulk update API
*/
func (a *Client) IpamRirsBulkUpdate(params *IpamRirsBulkUpdateParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*IpamRirsBulkUpdateOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewIpamRirsBulkUpdateParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "ipam_rirs_bulk_update",
		Method:             "PUT",
		PathPattern:        "/ipam/rirs/",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &IpamRirsBulkUpdateReader{formats: a.formats},
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
	success, ok := result.(*IpamRirsBulkUpdateOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for ipam_rirs_bulk_update: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  IpamRirsCreate ipam rirs create API
*/
func (a *Client) IpamRirsCreate(params *IpamRirsCreateParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*IpamRirsCreateCreated, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewIpamRirsCreateParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "ipam_rirs_create",
		Method:             "POST",
		PathPattern:        "/ipam/rirs/",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &IpamRirsCreateReader{formats: a.formats},
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
	success, ok := result.(*IpamRirsCreateCreated)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for ipam_rirs_create: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  IpamRirsDelete ipam rirs delete API
*/
func (a *Client) IpamRirsDelete(params *IpamRirsDeleteParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*IpamRirsDeleteNoContent, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewIpamRirsDeleteParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "ipam_rirs_delete",
		Method:             "DELETE",
		PathPattern:        "/ipam/rirs/{id}/",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &IpamRirsDeleteReader{formats: a.formats},
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
	success, ok := result.(*IpamRirsDeleteNoContent)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for ipam_rirs_delete: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  IpamRirsList ipam rirs list API
*/
func (a *Client) IpamRirsList(params *IpamRirsListParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*IpamRirsListOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewIpamRirsListParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "ipam_rirs_list",
		Method:             "GET",
		PathPattern:        "/ipam/rirs/",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &IpamRirsListReader{formats: a.formats},
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
	success, ok := result.(*IpamRirsListOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for ipam_rirs_list: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  IpamRirsPartialUpdate ipam rirs partial update API
*/
func (a *Client) IpamRirsPartialUpdate(params *IpamRirsPartialUpdateParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*IpamRirsPartialUpdateOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewIpamRirsPartialUpdateParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "ipam_rirs_partial_update",
		Method:             "PATCH",
		PathPattern:        "/ipam/rirs/{id}/",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &IpamRirsPartialUpdateReader{formats: a.formats},
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
	success, ok := result.(*IpamRirsPartialUpdateOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for ipam_rirs_partial_update: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  IpamRirsRead ipam rirs read API
*/
func (a *Client) IpamRirsRead(params *IpamRirsReadParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*IpamRirsReadOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewIpamRirsReadParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "ipam_rirs_read",
		Method:             "GET",
		PathPattern:        "/ipam/rirs/{id}/",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &IpamRirsReadReader{formats: a.formats},
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
	success, ok := result.(*IpamRirsReadOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for ipam_rirs_read: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  IpamRirsUpdate ipam rirs update API
*/
func (a *Client) IpamRirsUpdate(params *IpamRirsUpdateParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*IpamRirsUpdateOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewIpamRirsUpdateParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "ipam_rirs_update",
		Method:             "PUT",
		PathPattern:        "/ipam/rirs/{id}/",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &IpamRirsUpdateReader{formats: a.formats},
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
	success, ok := result.(*IpamRirsUpdateOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for ipam_rirs_update: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  IpamRolesBulkDelete ipam roles bulk delete API
*/
func (a *Client) IpamRolesBulkDelete(params *IpamRolesBulkDeleteParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*IpamRolesBulkDeleteNoContent, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewIpamRolesBulkDeleteParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "ipam_roles_bulk_delete",
		Method:             "DELETE",
		PathPattern:        "/ipam/roles/",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &IpamRolesBulkDeleteReader{formats: a.formats},
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
	success, ok := result.(*IpamRolesBulkDeleteNoContent)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for ipam_roles_bulk_delete: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  IpamRolesBulkPartialUpdate ipam roles bulk partial update API
*/
func (a *Client) IpamRolesBulkPartialUpdate(params *IpamRolesBulkPartialUpdateParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*IpamRolesBulkPartialUpdateOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewIpamRolesBulkPartialUpdateParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "ipam_roles_bulk_partial_update",
		Method:             "PATCH",
		PathPattern:        "/ipam/roles/",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &IpamRolesBulkPartialUpdateReader{formats: a.formats},
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
	success, ok := result.(*IpamRolesBulkPartialUpdateOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for ipam_roles_bulk_partial_update: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  IpamRolesBulkUpdate ipam roles bulk update API
*/
func (a *Client) IpamRolesBulkUpdate(params *IpamRolesBulkUpdateParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*IpamRolesBulkUpdateOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewIpamRolesBulkUpdateParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "ipam_roles_bulk_update",
		Method:             "PUT",
		PathPattern:        "/ipam/roles/",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &IpamRolesBulkUpdateReader{formats: a.formats},
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
	success, ok := result.(*IpamRolesBulkUpdateOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for ipam_roles_bulk_update: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  IpamRolesCreate ipam roles create API
*/
func (a *Client) IpamRolesCreate(params *IpamRolesCreateParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*IpamRolesCreateCreated, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewIpamRolesCreateParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "ipam_roles_create",
		Method:             "POST",
		PathPattern:        "/ipam/roles/",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &IpamRolesCreateReader{formats: a.formats},
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
	success, ok := result.(*IpamRolesCreateCreated)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for ipam_roles_create: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  IpamRolesDelete ipam roles delete API
*/
func (a *Client) IpamRolesDelete(params *IpamRolesDeleteParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*IpamRolesDeleteNoContent, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewIpamRolesDeleteParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "ipam_roles_delete",
		Method:             "DELETE",
		PathPattern:        "/ipam/roles/{id}/",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &IpamRolesDeleteReader{formats: a.formats},
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
	success, ok := result.(*IpamRolesDeleteNoContent)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for ipam_roles_delete: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  IpamRolesList ipam roles list API
*/
func (a *Client) IpamRolesList(params *IpamRolesListParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*IpamRolesListOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewIpamRolesListParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "ipam_roles_list",
		Method:             "GET",
		PathPattern:        "/ipam/roles/",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &IpamRolesListReader{formats: a.formats},
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
	success, ok := result.(*IpamRolesListOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for ipam_roles_list: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  IpamRolesPartialUpdate ipam roles partial update API
*/
func (a *Client) IpamRolesPartialUpdate(params *IpamRolesPartialUpdateParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*IpamRolesPartialUpdateOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewIpamRolesPartialUpdateParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "ipam_roles_partial_update",
		Method:             "PATCH",
		PathPattern:        "/ipam/roles/{id}/",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &IpamRolesPartialUpdateReader{formats: a.formats},
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
	success, ok := result.(*IpamRolesPartialUpdateOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for ipam_roles_partial_update: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  IpamRolesRead ipam roles read API
*/
func (a *Client) IpamRolesRead(params *IpamRolesReadParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*IpamRolesReadOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewIpamRolesReadParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "ipam_roles_read",
		Method:             "GET",
		PathPattern:        "/ipam/roles/{id}/",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &IpamRolesReadReader{formats: a.formats},
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
	success, ok := result.(*IpamRolesReadOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for ipam_roles_read: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  IpamRolesUpdate ipam roles update API
*/
func (a *Client) IpamRolesUpdate(params *IpamRolesUpdateParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*IpamRolesUpdateOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewIpamRolesUpdateParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "ipam_roles_update",
		Method:             "PUT",
		PathPattern:        "/ipam/roles/{id}/",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &IpamRolesUpdateReader{formats: a.formats},
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
	success, ok := result.(*IpamRolesUpdateOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for ipam_roles_update: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  IpamRouteTargetsBulkDelete ipam route targets bulk delete API
*/
func (a *Client) IpamRouteTargetsBulkDelete(params *IpamRouteTargetsBulkDeleteParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*IpamRouteTargetsBulkDeleteNoContent, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewIpamRouteTargetsBulkDeleteParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "ipam_route-targets_bulk_delete",
		Method:             "DELETE",
		PathPattern:        "/ipam/route-targets/",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &IpamRouteTargetsBulkDeleteReader{formats: a.formats},
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
	success, ok := result.(*IpamRouteTargetsBulkDeleteNoContent)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for ipam_route-targets_bulk_delete: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  IpamRouteTargetsBulkPartialUpdate ipam route targets bulk partial update API
*/
func (a *Client) IpamRouteTargetsBulkPartialUpdate(params *IpamRouteTargetsBulkPartialUpdateParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*IpamRouteTargetsBulkPartialUpdateOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewIpamRouteTargetsBulkPartialUpdateParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "ipam_route-targets_bulk_partial_update",
		Method:             "PATCH",
		PathPattern:        "/ipam/route-targets/",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &IpamRouteTargetsBulkPartialUpdateReader{formats: a.formats},
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
	success, ok := result.(*IpamRouteTargetsBulkPartialUpdateOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for ipam_route-targets_bulk_partial_update: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  IpamRouteTargetsBulkUpdate ipam route targets bulk update API
*/
func (a *Client) IpamRouteTargetsBulkUpdate(params *IpamRouteTargetsBulkUpdateParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*IpamRouteTargetsBulkUpdateOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewIpamRouteTargetsBulkUpdateParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "ipam_route-targets_bulk_update",
		Method:             "PUT",
		PathPattern:        "/ipam/route-targets/",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &IpamRouteTargetsBulkUpdateReader{formats: a.formats},
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
	success, ok := result.(*IpamRouteTargetsBulkUpdateOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for ipam_route-targets_bulk_update: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  IpamRouteTargetsCreate ipam route targets create API
*/
func (a *Client) IpamRouteTargetsCreate(params *IpamRouteTargetsCreateParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*IpamRouteTargetsCreateCreated, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewIpamRouteTargetsCreateParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "ipam_route-targets_create",
		Method:             "POST",
		PathPattern:        "/ipam/route-targets/",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &IpamRouteTargetsCreateReader{formats: a.formats},
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
	success, ok := result.(*IpamRouteTargetsCreateCreated)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for ipam_route-targets_create: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  IpamRouteTargetsDelete ipam route targets delete API
*/
func (a *Client) IpamRouteTargetsDelete(params *IpamRouteTargetsDeleteParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*IpamRouteTargetsDeleteNoContent, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewIpamRouteTargetsDeleteParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "ipam_route-targets_delete",
		Method:             "DELETE",
		PathPattern:        "/ipam/route-targets/{id}/",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &IpamRouteTargetsDeleteReader{formats: a.formats},
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
	success, ok := result.(*IpamRouteTargetsDeleteNoContent)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for ipam_route-targets_delete: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  IpamRouteTargetsList ipam route targets list API
*/
func (a *Client) IpamRouteTargetsList(params *IpamRouteTargetsListParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*IpamRouteTargetsListOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewIpamRouteTargetsListParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "ipam_route-targets_list",
		Method:             "GET",
		PathPattern:        "/ipam/route-targets/",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &IpamRouteTargetsListReader{formats: a.formats},
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
	success, ok := result.(*IpamRouteTargetsListOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for ipam_route-targets_list: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  IpamRouteTargetsPartialUpdate ipam route targets partial update API
*/
func (a *Client) IpamRouteTargetsPartialUpdate(params *IpamRouteTargetsPartialUpdateParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*IpamRouteTargetsPartialUpdateOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewIpamRouteTargetsPartialUpdateParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "ipam_route-targets_partial_update",
		Method:             "PATCH",
		PathPattern:        "/ipam/route-targets/{id}/",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &IpamRouteTargetsPartialUpdateReader{formats: a.formats},
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
	success, ok := result.(*IpamRouteTargetsPartialUpdateOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for ipam_route-targets_partial_update: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  IpamRouteTargetsRead ipam route targets read API
*/
func (a *Client) IpamRouteTargetsRead(params *IpamRouteTargetsReadParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*IpamRouteTargetsReadOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewIpamRouteTargetsReadParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "ipam_route-targets_read",
		Method:             "GET",
		PathPattern:        "/ipam/route-targets/{id}/",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &IpamRouteTargetsReadReader{formats: a.formats},
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
	success, ok := result.(*IpamRouteTargetsReadOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for ipam_route-targets_read: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  IpamRouteTargetsUpdate ipam route targets update API
*/
func (a *Client) IpamRouteTargetsUpdate(params *IpamRouteTargetsUpdateParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*IpamRouteTargetsUpdateOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewIpamRouteTargetsUpdateParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "ipam_route-targets_update",
		Method:             "PUT",
		PathPattern:        "/ipam/route-targets/{id}/",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &IpamRouteTargetsUpdateReader{formats: a.formats},
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
	success, ok := result.(*IpamRouteTargetsUpdateOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for ipam_route-targets_update: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  IpamServicesBulkDelete ipam services bulk delete API
*/
func (a *Client) IpamServicesBulkDelete(params *IpamServicesBulkDeleteParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*IpamServicesBulkDeleteNoContent, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewIpamServicesBulkDeleteParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "ipam_services_bulk_delete",
		Method:             "DELETE",
		PathPattern:        "/ipam/services/",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &IpamServicesBulkDeleteReader{formats: a.formats},
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
	success, ok := result.(*IpamServicesBulkDeleteNoContent)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for ipam_services_bulk_delete: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  IpamServicesBulkPartialUpdate ipam services bulk partial update API
*/
func (a *Client) IpamServicesBulkPartialUpdate(params *IpamServicesBulkPartialUpdateParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*IpamServicesBulkPartialUpdateOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewIpamServicesBulkPartialUpdateParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "ipam_services_bulk_partial_update",
		Method:             "PATCH",
		PathPattern:        "/ipam/services/",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &IpamServicesBulkPartialUpdateReader{formats: a.formats},
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
	success, ok := result.(*IpamServicesBulkPartialUpdateOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for ipam_services_bulk_partial_update: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  IpamServicesBulkUpdate ipam services bulk update API
*/
func (a *Client) IpamServicesBulkUpdate(params *IpamServicesBulkUpdateParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*IpamServicesBulkUpdateOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewIpamServicesBulkUpdateParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "ipam_services_bulk_update",
		Method:             "PUT",
		PathPattern:        "/ipam/services/",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &IpamServicesBulkUpdateReader{formats: a.formats},
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
	success, ok := result.(*IpamServicesBulkUpdateOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for ipam_services_bulk_update: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  IpamServicesCreate ipam services create API
*/
func (a *Client) IpamServicesCreate(params *IpamServicesCreateParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*IpamServicesCreateCreated, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewIpamServicesCreateParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "ipam_services_create",
		Method:             "POST",
		PathPattern:        "/ipam/services/",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &IpamServicesCreateReader{formats: a.formats},
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
	success, ok := result.(*IpamServicesCreateCreated)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for ipam_services_create: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  IpamServicesDelete ipam services delete API
*/
func (a *Client) IpamServicesDelete(params *IpamServicesDeleteParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*IpamServicesDeleteNoContent, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewIpamServicesDeleteParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "ipam_services_delete",
		Method:             "DELETE",
		PathPattern:        "/ipam/services/{id}/",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &IpamServicesDeleteReader{formats: a.formats},
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
	success, ok := result.(*IpamServicesDeleteNoContent)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for ipam_services_delete: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  IpamServicesList ipam services list API
*/
func (a *Client) IpamServicesList(params *IpamServicesListParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*IpamServicesListOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewIpamServicesListParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "ipam_services_list",
		Method:             "GET",
		PathPattern:        "/ipam/services/",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &IpamServicesListReader{formats: a.formats},
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
	success, ok := result.(*IpamServicesListOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for ipam_services_list: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  IpamServicesPartialUpdate ipam services partial update API
*/
func (a *Client) IpamServicesPartialUpdate(params *IpamServicesPartialUpdateParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*IpamServicesPartialUpdateOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewIpamServicesPartialUpdateParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "ipam_services_partial_update",
		Method:             "PATCH",
		PathPattern:        "/ipam/services/{id}/",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &IpamServicesPartialUpdateReader{formats: a.formats},
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
	success, ok := result.(*IpamServicesPartialUpdateOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for ipam_services_partial_update: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  IpamServicesRead ipam services read API
*/
func (a *Client) IpamServicesRead(params *IpamServicesReadParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*IpamServicesReadOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewIpamServicesReadParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "ipam_services_read",
		Method:             "GET",
		PathPattern:        "/ipam/services/{id}/",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &IpamServicesReadReader{formats: a.formats},
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
	success, ok := result.(*IpamServicesReadOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for ipam_services_read: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  IpamServicesUpdate ipam services update API
*/
func (a *Client) IpamServicesUpdate(params *IpamServicesUpdateParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*IpamServicesUpdateOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewIpamServicesUpdateParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "ipam_services_update",
		Method:             "PUT",
		PathPattern:        "/ipam/services/{id}/",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &IpamServicesUpdateReader{formats: a.formats},
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
	success, ok := result.(*IpamServicesUpdateOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for ipam_services_update: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  IpamVlanGroupsBulkDelete ipam vlan groups bulk delete API
*/
func (a *Client) IpamVlanGroupsBulkDelete(params *IpamVlanGroupsBulkDeleteParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*IpamVlanGroupsBulkDeleteNoContent, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewIpamVlanGroupsBulkDeleteParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "ipam_vlan-groups_bulk_delete",
		Method:             "DELETE",
		PathPattern:        "/ipam/vlan-groups/",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &IpamVlanGroupsBulkDeleteReader{formats: a.formats},
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
	success, ok := result.(*IpamVlanGroupsBulkDeleteNoContent)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for ipam_vlan-groups_bulk_delete: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  IpamVlanGroupsBulkPartialUpdate ipam vlan groups bulk partial update API
*/
func (a *Client) IpamVlanGroupsBulkPartialUpdate(params *IpamVlanGroupsBulkPartialUpdateParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*IpamVlanGroupsBulkPartialUpdateOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewIpamVlanGroupsBulkPartialUpdateParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "ipam_vlan-groups_bulk_partial_update",
		Method:             "PATCH",
		PathPattern:        "/ipam/vlan-groups/",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &IpamVlanGroupsBulkPartialUpdateReader{formats: a.formats},
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
	success, ok := result.(*IpamVlanGroupsBulkPartialUpdateOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for ipam_vlan-groups_bulk_partial_update: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  IpamVlanGroupsBulkUpdate ipam vlan groups bulk update API
*/
func (a *Client) IpamVlanGroupsBulkUpdate(params *IpamVlanGroupsBulkUpdateParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*IpamVlanGroupsBulkUpdateOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewIpamVlanGroupsBulkUpdateParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "ipam_vlan-groups_bulk_update",
		Method:             "PUT",
		PathPattern:        "/ipam/vlan-groups/",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &IpamVlanGroupsBulkUpdateReader{formats: a.formats},
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
	success, ok := result.(*IpamVlanGroupsBulkUpdateOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for ipam_vlan-groups_bulk_update: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  IpamVlanGroupsCreate ipam vlan groups create API
*/
func (a *Client) IpamVlanGroupsCreate(params *IpamVlanGroupsCreateParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*IpamVlanGroupsCreateCreated, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewIpamVlanGroupsCreateParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "ipam_vlan-groups_create",
		Method:             "POST",
		PathPattern:        "/ipam/vlan-groups/",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &IpamVlanGroupsCreateReader{formats: a.formats},
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
	success, ok := result.(*IpamVlanGroupsCreateCreated)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for ipam_vlan-groups_create: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  IpamVlanGroupsDelete ipam vlan groups delete API
*/
func (a *Client) IpamVlanGroupsDelete(params *IpamVlanGroupsDeleteParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*IpamVlanGroupsDeleteNoContent, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewIpamVlanGroupsDeleteParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "ipam_vlan-groups_delete",
		Method:             "DELETE",
		PathPattern:        "/ipam/vlan-groups/{id}/",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &IpamVlanGroupsDeleteReader{formats: a.formats},
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
	success, ok := result.(*IpamVlanGroupsDeleteNoContent)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for ipam_vlan-groups_delete: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  IpamVlanGroupsList ipam vlan groups list API
*/
func (a *Client) IpamVlanGroupsList(params *IpamVlanGroupsListParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*IpamVlanGroupsListOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewIpamVlanGroupsListParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "ipam_vlan-groups_list",
		Method:             "GET",
		PathPattern:        "/ipam/vlan-groups/",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &IpamVlanGroupsListReader{formats: a.formats},
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
	success, ok := result.(*IpamVlanGroupsListOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for ipam_vlan-groups_list: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  IpamVlanGroupsPartialUpdate ipam vlan groups partial update API
*/
func (a *Client) IpamVlanGroupsPartialUpdate(params *IpamVlanGroupsPartialUpdateParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*IpamVlanGroupsPartialUpdateOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewIpamVlanGroupsPartialUpdateParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "ipam_vlan-groups_partial_update",
		Method:             "PATCH",
		PathPattern:        "/ipam/vlan-groups/{id}/",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &IpamVlanGroupsPartialUpdateReader{formats: a.formats},
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
	success, ok := result.(*IpamVlanGroupsPartialUpdateOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for ipam_vlan-groups_partial_update: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  IpamVlanGroupsRead ipam vlan groups read API
*/
func (a *Client) IpamVlanGroupsRead(params *IpamVlanGroupsReadParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*IpamVlanGroupsReadOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewIpamVlanGroupsReadParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "ipam_vlan-groups_read",
		Method:             "GET",
		PathPattern:        "/ipam/vlan-groups/{id}/",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &IpamVlanGroupsReadReader{formats: a.formats},
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
	success, ok := result.(*IpamVlanGroupsReadOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for ipam_vlan-groups_read: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  IpamVlanGroupsUpdate ipam vlan groups update API
*/
func (a *Client) IpamVlanGroupsUpdate(params *IpamVlanGroupsUpdateParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*IpamVlanGroupsUpdateOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewIpamVlanGroupsUpdateParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "ipam_vlan-groups_update",
		Method:             "PUT",
		PathPattern:        "/ipam/vlan-groups/{id}/",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &IpamVlanGroupsUpdateReader{formats: a.formats},
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
	success, ok := result.(*IpamVlanGroupsUpdateOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for ipam_vlan-groups_update: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  IpamVlansBulkDelete ipam vlans bulk delete API
*/
func (a *Client) IpamVlansBulkDelete(params *IpamVlansBulkDeleteParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*IpamVlansBulkDeleteNoContent, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewIpamVlansBulkDeleteParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "ipam_vlans_bulk_delete",
		Method:             "DELETE",
		PathPattern:        "/ipam/vlans/",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &IpamVlansBulkDeleteReader{formats: a.formats},
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
	success, ok := result.(*IpamVlansBulkDeleteNoContent)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for ipam_vlans_bulk_delete: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  IpamVlansBulkPartialUpdate ipam vlans bulk partial update API
*/
func (a *Client) IpamVlansBulkPartialUpdate(params *IpamVlansBulkPartialUpdateParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*IpamVlansBulkPartialUpdateOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewIpamVlansBulkPartialUpdateParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "ipam_vlans_bulk_partial_update",
		Method:             "PATCH",
		PathPattern:        "/ipam/vlans/",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &IpamVlansBulkPartialUpdateReader{formats: a.formats},
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
	success, ok := result.(*IpamVlansBulkPartialUpdateOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for ipam_vlans_bulk_partial_update: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  IpamVlansBulkUpdate ipam vlans bulk update API
*/
func (a *Client) IpamVlansBulkUpdate(params *IpamVlansBulkUpdateParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*IpamVlansBulkUpdateOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewIpamVlansBulkUpdateParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "ipam_vlans_bulk_update",
		Method:             "PUT",
		PathPattern:        "/ipam/vlans/",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &IpamVlansBulkUpdateReader{formats: a.formats},
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
	success, ok := result.(*IpamVlansBulkUpdateOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for ipam_vlans_bulk_update: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  IpamVlansCreate ipam vlans create API
*/
func (a *Client) IpamVlansCreate(params *IpamVlansCreateParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*IpamVlansCreateCreated, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewIpamVlansCreateParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "ipam_vlans_create",
		Method:             "POST",
		PathPattern:        "/ipam/vlans/",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &IpamVlansCreateReader{formats: a.formats},
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
	success, ok := result.(*IpamVlansCreateCreated)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for ipam_vlans_create: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  IpamVlansDelete ipam vlans delete API
*/
func (a *Client) IpamVlansDelete(params *IpamVlansDeleteParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*IpamVlansDeleteNoContent, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewIpamVlansDeleteParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "ipam_vlans_delete",
		Method:             "DELETE",
		PathPattern:        "/ipam/vlans/{id}/",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &IpamVlansDeleteReader{formats: a.formats},
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
	success, ok := result.(*IpamVlansDeleteNoContent)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for ipam_vlans_delete: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  IpamVlansList ipam vlans list API
*/
func (a *Client) IpamVlansList(params *IpamVlansListParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*IpamVlansListOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewIpamVlansListParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "ipam_vlans_list",
		Method:             "GET",
		PathPattern:        "/ipam/vlans/",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &IpamVlansListReader{formats: a.formats},
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
	success, ok := result.(*IpamVlansListOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for ipam_vlans_list: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  IpamVlansPartialUpdate ipam vlans partial update API
*/
func (a *Client) IpamVlansPartialUpdate(params *IpamVlansPartialUpdateParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*IpamVlansPartialUpdateOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewIpamVlansPartialUpdateParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "ipam_vlans_partial_update",
		Method:             "PATCH",
		PathPattern:        "/ipam/vlans/{id}/",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &IpamVlansPartialUpdateReader{formats: a.formats},
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
	success, ok := result.(*IpamVlansPartialUpdateOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for ipam_vlans_partial_update: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  IpamVlansRead ipam vlans read API
*/
func (a *Client) IpamVlansRead(params *IpamVlansReadParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*IpamVlansReadOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewIpamVlansReadParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "ipam_vlans_read",
		Method:             "GET",
		PathPattern:        "/ipam/vlans/{id}/",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &IpamVlansReadReader{formats: a.formats},
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
	success, ok := result.(*IpamVlansReadOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for ipam_vlans_read: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  IpamVlansUpdate ipam vlans update API
*/
func (a *Client) IpamVlansUpdate(params *IpamVlansUpdateParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*IpamVlansUpdateOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewIpamVlansUpdateParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "ipam_vlans_update",
		Method:             "PUT",
		PathPattern:        "/ipam/vlans/{id}/",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &IpamVlansUpdateReader{formats: a.formats},
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
	success, ok := result.(*IpamVlansUpdateOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for ipam_vlans_update: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  IpamVrfsBulkDelete ipam vrfs bulk delete API
*/
func (a *Client) IpamVrfsBulkDelete(params *IpamVrfsBulkDeleteParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*IpamVrfsBulkDeleteNoContent, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewIpamVrfsBulkDeleteParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "ipam_vrfs_bulk_delete",
		Method:             "DELETE",
		PathPattern:        "/ipam/vrfs/",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &IpamVrfsBulkDeleteReader{formats: a.formats},
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
	success, ok := result.(*IpamVrfsBulkDeleteNoContent)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for ipam_vrfs_bulk_delete: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  IpamVrfsBulkPartialUpdate ipam vrfs bulk partial update API
*/
func (a *Client) IpamVrfsBulkPartialUpdate(params *IpamVrfsBulkPartialUpdateParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*IpamVrfsBulkPartialUpdateOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewIpamVrfsBulkPartialUpdateParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "ipam_vrfs_bulk_partial_update",
		Method:             "PATCH",
		PathPattern:        "/ipam/vrfs/",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &IpamVrfsBulkPartialUpdateReader{formats: a.formats},
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
	success, ok := result.(*IpamVrfsBulkPartialUpdateOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for ipam_vrfs_bulk_partial_update: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  IpamVrfsBulkUpdate ipam vrfs bulk update API
*/
func (a *Client) IpamVrfsBulkUpdate(params *IpamVrfsBulkUpdateParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*IpamVrfsBulkUpdateOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewIpamVrfsBulkUpdateParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "ipam_vrfs_bulk_update",
		Method:             "PUT",
		PathPattern:        "/ipam/vrfs/",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &IpamVrfsBulkUpdateReader{formats: a.formats},
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
	success, ok := result.(*IpamVrfsBulkUpdateOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for ipam_vrfs_bulk_update: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  IpamVrfsCreate ipam vrfs create API
*/
func (a *Client) IpamVrfsCreate(params *IpamVrfsCreateParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*IpamVrfsCreateCreated, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewIpamVrfsCreateParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "ipam_vrfs_create",
		Method:             "POST",
		PathPattern:        "/ipam/vrfs/",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &IpamVrfsCreateReader{formats: a.formats},
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
	success, ok := result.(*IpamVrfsCreateCreated)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for ipam_vrfs_create: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  IpamVrfsDelete ipam vrfs delete API
*/
func (a *Client) IpamVrfsDelete(params *IpamVrfsDeleteParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*IpamVrfsDeleteNoContent, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewIpamVrfsDeleteParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "ipam_vrfs_delete",
		Method:             "DELETE",
		PathPattern:        "/ipam/vrfs/{id}/",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &IpamVrfsDeleteReader{formats: a.formats},
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
	success, ok := result.(*IpamVrfsDeleteNoContent)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for ipam_vrfs_delete: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  IpamVrfsList ipam vrfs list API
*/
func (a *Client) IpamVrfsList(params *IpamVrfsListParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*IpamVrfsListOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewIpamVrfsListParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "ipam_vrfs_list",
		Method:             "GET",
		PathPattern:        "/ipam/vrfs/",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &IpamVrfsListReader{formats: a.formats},
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
	success, ok := result.(*IpamVrfsListOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for ipam_vrfs_list: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  IpamVrfsPartialUpdate ipam vrfs partial update API
*/
func (a *Client) IpamVrfsPartialUpdate(params *IpamVrfsPartialUpdateParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*IpamVrfsPartialUpdateOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewIpamVrfsPartialUpdateParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "ipam_vrfs_partial_update",
		Method:             "PATCH",
		PathPattern:        "/ipam/vrfs/{id}/",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &IpamVrfsPartialUpdateReader{formats: a.formats},
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
	success, ok := result.(*IpamVrfsPartialUpdateOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for ipam_vrfs_partial_update: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  IpamVrfsRead ipam vrfs read API
*/
func (a *Client) IpamVrfsRead(params *IpamVrfsReadParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*IpamVrfsReadOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewIpamVrfsReadParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "ipam_vrfs_read",
		Method:             "GET",
		PathPattern:        "/ipam/vrfs/{id}/",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &IpamVrfsReadReader{formats: a.formats},
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
	success, ok := result.(*IpamVrfsReadOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for ipam_vrfs_read: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  IpamVrfsUpdate ipam vrfs update API
*/
func (a *Client) IpamVrfsUpdate(params *IpamVrfsUpdateParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*IpamVrfsUpdateOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewIpamVrfsUpdateParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "ipam_vrfs_update",
		Method:             "PUT",
		PathPattern:        "/ipam/vrfs/{id}/",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &IpamVrfsUpdateReader{formats: a.formats},
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
	success, ok := result.(*IpamVrfsUpdateOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for ipam_vrfs_update: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
