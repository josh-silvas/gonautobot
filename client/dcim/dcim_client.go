package dcim

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// New creates a new dcim API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

/*
Client for dcim API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientOption is the option for Client methods
type ClientOption func(*runtime.ClientOperation)

// ClientService is the interface for Client methods
type ClientService interface {
	DcimCablesBulkDelete(params *DcimCablesBulkDeleteParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*DcimCablesBulkDeleteNoContent, error)

	DcimCablesBulkPartialUpdate(params *DcimCablesBulkPartialUpdateParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*DcimCablesBulkPartialUpdateOK, error)

	DcimCablesBulkUpdate(params *DcimCablesBulkUpdateParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*DcimCablesBulkUpdateOK, error)

	DcimCablesCreate(params *DcimCablesCreateParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*DcimCablesCreateCreated, error)

	DcimCablesDelete(params *DcimCablesDeleteParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*DcimCablesDeleteNoContent, error)

	DcimCablesList(params *DcimCablesListParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*DcimCablesListOK, error)

	DcimCablesPartialUpdate(params *DcimCablesPartialUpdateParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*DcimCablesPartialUpdateOK, error)

	DcimCablesRead(params *DcimCablesReadParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*DcimCablesReadOK, error)

	DcimCablesUpdate(params *DcimCablesUpdateParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*DcimCablesUpdateOK, error)

	DcimConnectedDeviceList(params *DcimConnectedDeviceListParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*DcimConnectedDeviceListOK, error)

	DcimConsoleConnectionsList(params *DcimConsoleConnectionsListParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*DcimConsoleConnectionsListOK, error)

	DcimConsolePortTemplatesBulkDelete(params *DcimConsolePortTemplatesBulkDeleteParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*DcimConsolePortTemplatesBulkDeleteNoContent, error)

	DcimConsolePortTemplatesBulkPartialUpdate(params *DcimConsolePortTemplatesBulkPartialUpdateParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*DcimConsolePortTemplatesBulkPartialUpdateOK, error)

	DcimConsolePortTemplatesBulkUpdate(params *DcimConsolePortTemplatesBulkUpdateParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*DcimConsolePortTemplatesBulkUpdateOK, error)

	DcimConsolePortTemplatesCreate(params *DcimConsolePortTemplatesCreateParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*DcimConsolePortTemplatesCreateCreated, error)

	DcimConsolePortTemplatesDelete(params *DcimConsolePortTemplatesDeleteParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*DcimConsolePortTemplatesDeleteNoContent, error)

	DcimConsolePortTemplatesList(params *DcimConsolePortTemplatesListParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*DcimConsolePortTemplatesListOK, error)

	DcimConsolePortTemplatesPartialUpdate(params *DcimConsolePortTemplatesPartialUpdateParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*DcimConsolePortTemplatesPartialUpdateOK, error)

	DcimConsolePortTemplatesRead(params *DcimConsolePortTemplatesReadParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*DcimConsolePortTemplatesReadOK, error)

	DcimConsolePortTemplatesUpdate(params *DcimConsolePortTemplatesUpdateParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*DcimConsolePortTemplatesUpdateOK, error)

	DcimConsolePortsBulkDelete(params *DcimConsolePortsBulkDeleteParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*DcimConsolePortsBulkDeleteNoContent, error)

	DcimConsolePortsBulkPartialUpdate(params *DcimConsolePortsBulkPartialUpdateParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*DcimConsolePortsBulkPartialUpdateOK, error)

	DcimConsolePortsBulkUpdate(params *DcimConsolePortsBulkUpdateParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*DcimConsolePortsBulkUpdateOK, error)

	DcimConsolePortsCreate(params *DcimConsolePortsCreateParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*DcimConsolePortsCreateCreated, error)

	DcimConsolePortsDelete(params *DcimConsolePortsDeleteParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*DcimConsolePortsDeleteNoContent, error)

	DcimConsolePortsList(params *DcimConsolePortsListParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*DcimConsolePortsListOK, error)

	DcimConsolePortsPartialUpdate(params *DcimConsolePortsPartialUpdateParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*DcimConsolePortsPartialUpdateOK, error)

	DcimConsolePortsRead(params *DcimConsolePortsReadParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*DcimConsolePortsReadOK, error)

	DcimConsolePortsTrace(params *DcimConsolePortsTraceParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*DcimConsolePortsTraceOK, error)

	DcimConsolePortsUpdate(params *DcimConsolePortsUpdateParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*DcimConsolePortsUpdateOK, error)

	DcimConsoleServerPortTemplatesBulkDelete(params *DcimConsoleServerPortTemplatesBulkDeleteParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*DcimConsoleServerPortTemplatesBulkDeleteNoContent, error)

	DcimConsoleServerPortTemplatesBulkPartialUpdate(params *DcimConsoleServerPortTemplatesBulkPartialUpdateParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*DcimConsoleServerPortTemplatesBulkPartialUpdateOK, error)

	DcimConsoleServerPortTemplatesBulkUpdate(params *DcimConsoleServerPortTemplatesBulkUpdateParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*DcimConsoleServerPortTemplatesBulkUpdateOK, error)

	DcimConsoleServerPortTemplatesCreate(params *DcimConsoleServerPortTemplatesCreateParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*DcimConsoleServerPortTemplatesCreateCreated, error)

	DcimConsoleServerPortTemplatesDelete(params *DcimConsoleServerPortTemplatesDeleteParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*DcimConsoleServerPortTemplatesDeleteNoContent, error)

	DcimConsoleServerPortTemplatesList(params *DcimConsoleServerPortTemplatesListParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*DcimConsoleServerPortTemplatesListOK, error)

	DcimConsoleServerPortTemplatesPartialUpdate(params *DcimConsoleServerPortTemplatesPartialUpdateParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*DcimConsoleServerPortTemplatesPartialUpdateOK, error)

	DcimConsoleServerPortTemplatesRead(params *DcimConsoleServerPortTemplatesReadParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*DcimConsoleServerPortTemplatesReadOK, error)

	DcimConsoleServerPortTemplatesUpdate(params *DcimConsoleServerPortTemplatesUpdateParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*DcimConsoleServerPortTemplatesUpdateOK, error)

	DcimConsoleServerPortsBulkDelete(params *DcimConsoleServerPortsBulkDeleteParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*DcimConsoleServerPortsBulkDeleteNoContent, error)

	DcimConsoleServerPortsBulkPartialUpdate(params *DcimConsoleServerPortsBulkPartialUpdateParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*DcimConsoleServerPortsBulkPartialUpdateOK, error)

	DcimConsoleServerPortsBulkUpdate(params *DcimConsoleServerPortsBulkUpdateParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*DcimConsoleServerPortsBulkUpdateOK, error)

	DcimConsoleServerPortsCreate(params *DcimConsoleServerPortsCreateParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*DcimConsoleServerPortsCreateCreated, error)

	DcimConsoleServerPortsDelete(params *DcimConsoleServerPortsDeleteParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*DcimConsoleServerPortsDeleteNoContent, error)

	DcimConsoleServerPortsList(params *DcimConsoleServerPortsListParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*DcimConsoleServerPortsListOK, error)

	DcimConsoleServerPortsPartialUpdate(params *DcimConsoleServerPortsPartialUpdateParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*DcimConsoleServerPortsPartialUpdateOK, error)

	DcimConsoleServerPortsRead(params *DcimConsoleServerPortsReadParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*DcimConsoleServerPortsReadOK, error)

	DcimConsoleServerPortsTrace(params *DcimConsoleServerPortsTraceParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*DcimConsoleServerPortsTraceOK, error)

	DcimConsoleServerPortsUpdate(params *DcimConsoleServerPortsUpdateParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*DcimConsoleServerPortsUpdateOK, error)

	DcimDeviceBayTemplatesBulkDelete(params *DcimDeviceBayTemplatesBulkDeleteParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*DcimDeviceBayTemplatesBulkDeleteNoContent, error)

	DcimDeviceBayTemplatesBulkPartialUpdate(params *DcimDeviceBayTemplatesBulkPartialUpdateParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*DcimDeviceBayTemplatesBulkPartialUpdateOK, error)

	DcimDeviceBayTemplatesBulkUpdate(params *DcimDeviceBayTemplatesBulkUpdateParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*DcimDeviceBayTemplatesBulkUpdateOK, error)

	DcimDeviceBayTemplatesCreate(params *DcimDeviceBayTemplatesCreateParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*DcimDeviceBayTemplatesCreateCreated, error)

	DcimDeviceBayTemplatesDelete(params *DcimDeviceBayTemplatesDeleteParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*DcimDeviceBayTemplatesDeleteNoContent, error)

	DcimDeviceBayTemplatesList(params *DcimDeviceBayTemplatesListParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*DcimDeviceBayTemplatesListOK, error)

	DcimDeviceBayTemplatesPartialUpdate(params *DcimDeviceBayTemplatesPartialUpdateParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*DcimDeviceBayTemplatesPartialUpdateOK, error)

	DcimDeviceBayTemplatesRead(params *DcimDeviceBayTemplatesReadParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*DcimDeviceBayTemplatesReadOK, error)

	DcimDeviceBayTemplatesUpdate(params *DcimDeviceBayTemplatesUpdateParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*DcimDeviceBayTemplatesUpdateOK, error)

	DcimDeviceBaysBulkDelete(params *DcimDeviceBaysBulkDeleteParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*DcimDeviceBaysBulkDeleteNoContent, error)

	DcimDeviceBaysBulkPartialUpdate(params *DcimDeviceBaysBulkPartialUpdateParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*DcimDeviceBaysBulkPartialUpdateOK, error)

	DcimDeviceBaysBulkUpdate(params *DcimDeviceBaysBulkUpdateParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*DcimDeviceBaysBulkUpdateOK, error)

	DcimDeviceBaysCreate(params *DcimDeviceBaysCreateParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*DcimDeviceBaysCreateCreated, error)

	DcimDeviceBaysDelete(params *DcimDeviceBaysDeleteParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*DcimDeviceBaysDeleteNoContent, error)

	DcimDeviceBaysList(params *DcimDeviceBaysListParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*DcimDeviceBaysListOK, error)

	DcimDeviceBaysPartialUpdate(params *DcimDeviceBaysPartialUpdateParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*DcimDeviceBaysPartialUpdateOK, error)

	DcimDeviceBaysRead(params *DcimDeviceBaysReadParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*DcimDeviceBaysReadOK, error)

	DcimDeviceBaysUpdate(params *DcimDeviceBaysUpdateParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*DcimDeviceBaysUpdateOK, error)

	DcimDeviceRolesBulkDelete(params *DcimDeviceRolesBulkDeleteParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*DcimDeviceRolesBulkDeleteNoContent, error)

	DcimDeviceRolesBulkPartialUpdate(params *DcimDeviceRolesBulkPartialUpdateParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*DcimDeviceRolesBulkPartialUpdateOK, error)

	DcimDeviceRolesBulkUpdate(params *DcimDeviceRolesBulkUpdateParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*DcimDeviceRolesBulkUpdateOK, error)

	DcimDeviceRolesCreate(params *DcimDeviceRolesCreateParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*DcimDeviceRolesCreateCreated, error)

	DcimDeviceRolesDelete(params *DcimDeviceRolesDeleteParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*DcimDeviceRolesDeleteNoContent, error)

	DcimDeviceRolesList(params *DcimDeviceRolesListParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*DcimDeviceRolesListOK, error)

	DcimDeviceRolesPartialUpdate(params *DcimDeviceRolesPartialUpdateParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*DcimDeviceRolesPartialUpdateOK, error)

	DcimDeviceRolesRead(params *DcimDeviceRolesReadParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*DcimDeviceRolesReadOK, error)

	DcimDeviceRolesUpdate(params *DcimDeviceRolesUpdateParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*DcimDeviceRolesUpdateOK, error)

	DcimDeviceTypesBulkDelete(params *DcimDeviceTypesBulkDeleteParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*DcimDeviceTypesBulkDeleteNoContent, error)

	DcimDeviceTypesBulkPartialUpdate(params *DcimDeviceTypesBulkPartialUpdateParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*DcimDeviceTypesBulkPartialUpdateOK, error)

	DcimDeviceTypesBulkUpdate(params *DcimDeviceTypesBulkUpdateParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*DcimDeviceTypesBulkUpdateOK, error)

	DcimDeviceTypesCreate(params *DcimDeviceTypesCreateParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*DcimDeviceTypesCreateCreated, error)

	DcimDeviceTypesDelete(params *DcimDeviceTypesDeleteParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*DcimDeviceTypesDeleteNoContent, error)

	DcimDeviceTypesList(params *DcimDeviceTypesListParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*DcimDeviceTypesListOK, error)

	DcimDeviceTypesPartialUpdate(params *DcimDeviceTypesPartialUpdateParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*DcimDeviceTypesPartialUpdateOK, error)

	DcimDeviceTypesRead(params *DcimDeviceTypesReadParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*DcimDeviceTypesReadOK, error)

	DcimDeviceTypesUpdate(params *DcimDeviceTypesUpdateParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*DcimDeviceTypesUpdateOK, error)

	DcimDevicesBulkDelete(params *DcimDevicesBulkDeleteParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*DcimDevicesBulkDeleteNoContent, error)

	DcimDevicesBulkPartialUpdate(params *DcimDevicesBulkPartialUpdateParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*DcimDevicesBulkPartialUpdateOK, error)

	DcimDevicesBulkUpdate(params *DcimDevicesBulkUpdateParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*DcimDevicesBulkUpdateOK, error)

	DcimDevicesCreate(params *DcimDevicesCreateParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*DcimDevicesCreateCreated, error)

	DcimDevicesDelete(params *DcimDevicesDeleteParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*DcimDevicesDeleteNoContent, error)

	DcimDevicesList(params *DcimDevicesListParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*DcimDevicesListOK, error)

	DcimDevicesNapalm(params *DcimDevicesNapalmParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*DcimDevicesNapalmOK, error)

	DcimDevicesPartialUpdate(params *DcimDevicesPartialUpdateParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*DcimDevicesPartialUpdateOK, error)

	DcimDevicesRead(params *DcimDevicesReadParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*DcimDevicesReadOK, error)

	DcimDevicesUpdate(params *DcimDevicesUpdateParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*DcimDevicesUpdateOK, error)

	DcimFrontPortTemplatesBulkDelete(params *DcimFrontPortTemplatesBulkDeleteParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*DcimFrontPortTemplatesBulkDeleteNoContent, error)

	DcimFrontPortTemplatesBulkPartialUpdate(params *DcimFrontPortTemplatesBulkPartialUpdateParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*DcimFrontPortTemplatesBulkPartialUpdateOK, error)

	DcimFrontPortTemplatesBulkUpdate(params *DcimFrontPortTemplatesBulkUpdateParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*DcimFrontPortTemplatesBulkUpdateOK, error)

	DcimFrontPortTemplatesCreate(params *DcimFrontPortTemplatesCreateParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*DcimFrontPortTemplatesCreateCreated, error)

	DcimFrontPortTemplatesDelete(params *DcimFrontPortTemplatesDeleteParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*DcimFrontPortTemplatesDeleteNoContent, error)

	DcimFrontPortTemplatesList(params *DcimFrontPortTemplatesListParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*DcimFrontPortTemplatesListOK, error)

	DcimFrontPortTemplatesPartialUpdate(params *DcimFrontPortTemplatesPartialUpdateParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*DcimFrontPortTemplatesPartialUpdateOK, error)

	DcimFrontPortTemplatesRead(params *DcimFrontPortTemplatesReadParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*DcimFrontPortTemplatesReadOK, error)

	DcimFrontPortTemplatesUpdate(params *DcimFrontPortTemplatesUpdateParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*DcimFrontPortTemplatesUpdateOK, error)

	DcimFrontPortsBulkDelete(params *DcimFrontPortsBulkDeleteParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*DcimFrontPortsBulkDeleteNoContent, error)

	DcimFrontPortsBulkPartialUpdate(params *DcimFrontPortsBulkPartialUpdateParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*DcimFrontPortsBulkPartialUpdateOK, error)

	DcimFrontPortsBulkUpdate(params *DcimFrontPortsBulkUpdateParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*DcimFrontPortsBulkUpdateOK, error)

	DcimFrontPortsCreate(params *DcimFrontPortsCreateParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*DcimFrontPortsCreateCreated, error)

	DcimFrontPortsDelete(params *DcimFrontPortsDeleteParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*DcimFrontPortsDeleteNoContent, error)

	DcimFrontPortsList(params *DcimFrontPortsListParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*DcimFrontPortsListOK, error)

	DcimFrontPortsPartialUpdate(params *DcimFrontPortsPartialUpdateParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*DcimFrontPortsPartialUpdateOK, error)

	DcimFrontPortsPaths(params *DcimFrontPortsPathsParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*DcimFrontPortsPathsOK, error)

	DcimFrontPortsRead(params *DcimFrontPortsReadParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*DcimFrontPortsReadOK, error)

	DcimFrontPortsUpdate(params *DcimFrontPortsUpdateParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*DcimFrontPortsUpdateOK, error)

	DcimInterfaceConnectionsList(params *DcimInterfaceConnectionsListParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*DcimInterfaceConnectionsListOK, error)

	DcimInterfaceTemplatesBulkDelete(params *DcimInterfaceTemplatesBulkDeleteParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*DcimInterfaceTemplatesBulkDeleteNoContent, error)

	DcimInterfaceTemplatesBulkPartialUpdate(params *DcimInterfaceTemplatesBulkPartialUpdateParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*DcimInterfaceTemplatesBulkPartialUpdateOK, error)

	DcimInterfaceTemplatesBulkUpdate(params *DcimInterfaceTemplatesBulkUpdateParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*DcimInterfaceTemplatesBulkUpdateOK, error)

	DcimInterfaceTemplatesCreate(params *DcimInterfaceTemplatesCreateParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*DcimInterfaceTemplatesCreateCreated, error)

	DcimInterfaceTemplatesDelete(params *DcimInterfaceTemplatesDeleteParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*DcimInterfaceTemplatesDeleteNoContent, error)

	DcimInterfaceTemplatesList(params *DcimInterfaceTemplatesListParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*DcimInterfaceTemplatesListOK, error)

	DcimInterfaceTemplatesPartialUpdate(params *DcimInterfaceTemplatesPartialUpdateParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*DcimInterfaceTemplatesPartialUpdateOK, error)

	DcimInterfaceTemplatesRead(params *DcimInterfaceTemplatesReadParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*DcimInterfaceTemplatesReadOK, error)

	DcimInterfaceTemplatesUpdate(params *DcimInterfaceTemplatesUpdateParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*DcimInterfaceTemplatesUpdateOK, error)

	DcimInterfacesBulkDelete(params *DcimInterfacesBulkDeleteParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*DcimInterfacesBulkDeleteNoContent, error)

	DcimInterfacesBulkPartialUpdate(params *DcimInterfacesBulkPartialUpdateParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*DcimInterfacesBulkPartialUpdateOK, error)

	DcimInterfacesBulkUpdate(params *DcimInterfacesBulkUpdateParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*DcimInterfacesBulkUpdateOK, error)

	DcimInterfacesCreate(params *DcimInterfacesCreateParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*DcimInterfacesCreateCreated, error)

	DcimInterfacesDelete(params *DcimInterfacesDeleteParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*DcimInterfacesDeleteNoContent, error)

	DcimInterfacesList(params *DcimInterfacesListParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*DcimInterfacesListOK, error)

	DcimInterfacesPartialUpdate(params *DcimInterfacesPartialUpdateParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*DcimInterfacesPartialUpdateOK, error)

	DcimInterfacesRead(params *DcimInterfacesReadParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*DcimInterfacesReadOK, error)

	DcimInterfacesTrace(params *DcimInterfacesTraceParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*DcimInterfacesTraceOK, error)

	DcimInterfacesUpdate(params *DcimInterfacesUpdateParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*DcimInterfacesUpdateOK, error)

	DcimInventoryItemsBulkDelete(params *DcimInventoryItemsBulkDeleteParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*DcimInventoryItemsBulkDeleteNoContent, error)

	DcimInventoryItemsBulkPartialUpdate(params *DcimInventoryItemsBulkPartialUpdateParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*DcimInventoryItemsBulkPartialUpdateOK, error)

	DcimInventoryItemsBulkUpdate(params *DcimInventoryItemsBulkUpdateParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*DcimInventoryItemsBulkUpdateOK, error)

	DcimInventoryItemsCreate(params *DcimInventoryItemsCreateParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*DcimInventoryItemsCreateCreated, error)

	DcimInventoryItemsDelete(params *DcimInventoryItemsDeleteParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*DcimInventoryItemsDeleteNoContent, error)

	DcimInventoryItemsList(params *DcimInventoryItemsListParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*DcimInventoryItemsListOK, error)

	DcimInventoryItemsPartialUpdate(params *DcimInventoryItemsPartialUpdateParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*DcimInventoryItemsPartialUpdateOK, error)

	DcimInventoryItemsRead(params *DcimInventoryItemsReadParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*DcimInventoryItemsReadOK, error)

	DcimInventoryItemsUpdate(params *DcimInventoryItemsUpdateParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*DcimInventoryItemsUpdateOK, error)

	DcimManufacturersBulkDelete(params *DcimManufacturersBulkDeleteParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*DcimManufacturersBulkDeleteNoContent, error)

	DcimManufacturersBulkPartialUpdate(params *DcimManufacturersBulkPartialUpdateParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*DcimManufacturersBulkPartialUpdateOK, error)

	DcimManufacturersBulkUpdate(params *DcimManufacturersBulkUpdateParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*DcimManufacturersBulkUpdateOK, error)

	DcimManufacturersCreate(params *DcimManufacturersCreateParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*DcimManufacturersCreateCreated, error)

	DcimManufacturersDelete(params *DcimManufacturersDeleteParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*DcimManufacturersDeleteNoContent, error)

	DcimManufacturersList(params *DcimManufacturersListParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*DcimManufacturersListOK, error)

	DcimManufacturersPartialUpdate(params *DcimManufacturersPartialUpdateParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*DcimManufacturersPartialUpdateOK, error)

	DcimManufacturersRead(params *DcimManufacturersReadParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*DcimManufacturersReadOK, error)

	DcimManufacturersUpdate(params *DcimManufacturersUpdateParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*DcimManufacturersUpdateOK, error)

	DcimPlatformsBulkDelete(params *DcimPlatformsBulkDeleteParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*DcimPlatformsBulkDeleteNoContent, error)

	DcimPlatformsBulkPartialUpdate(params *DcimPlatformsBulkPartialUpdateParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*DcimPlatformsBulkPartialUpdateOK, error)

	DcimPlatformsBulkUpdate(params *DcimPlatformsBulkUpdateParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*DcimPlatformsBulkUpdateOK, error)

	DcimPlatformsCreate(params *DcimPlatformsCreateParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*DcimPlatformsCreateCreated, error)

	DcimPlatformsDelete(params *DcimPlatformsDeleteParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*DcimPlatformsDeleteNoContent, error)

	DcimPlatformsList(params *DcimPlatformsListParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*DcimPlatformsListOK, error)

	DcimPlatformsPartialUpdate(params *DcimPlatformsPartialUpdateParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*DcimPlatformsPartialUpdateOK, error)

	DcimPlatformsRead(params *DcimPlatformsReadParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*DcimPlatformsReadOK, error)

	DcimPlatformsUpdate(params *DcimPlatformsUpdateParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*DcimPlatformsUpdateOK, error)

	DcimPowerConnectionsList(params *DcimPowerConnectionsListParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*DcimPowerConnectionsListOK, error)

	DcimPowerFeedsBulkDelete(params *DcimPowerFeedsBulkDeleteParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*DcimPowerFeedsBulkDeleteNoContent, error)

	DcimPowerFeedsBulkPartialUpdate(params *DcimPowerFeedsBulkPartialUpdateParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*DcimPowerFeedsBulkPartialUpdateOK, error)

	DcimPowerFeedsBulkUpdate(params *DcimPowerFeedsBulkUpdateParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*DcimPowerFeedsBulkUpdateOK, error)

	DcimPowerFeedsCreate(params *DcimPowerFeedsCreateParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*DcimPowerFeedsCreateCreated, error)

	DcimPowerFeedsDelete(params *DcimPowerFeedsDeleteParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*DcimPowerFeedsDeleteNoContent, error)

	DcimPowerFeedsList(params *DcimPowerFeedsListParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*DcimPowerFeedsListOK, error)

	DcimPowerFeedsPartialUpdate(params *DcimPowerFeedsPartialUpdateParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*DcimPowerFeedsPartialUpdateOK, error)

	DcimPowerFeedsRead(params *DcimPowerFeedsReadParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*DcimPowerFeedsReadOK, error)

	DcimPowerFeedsTrace(params *DcimPowerFeedsTraceParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*DcimPowerFeedsTraceOK, error)

	DcimPowerFeedsUpdate(params *DcimPowerFeedsUpdateParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*DcimPowerFeedsUpdateOK, error)

	DcimPowerOutletTemplatesBulkDelete(params *DcimPowerOutletTemplatesBulkDeleteParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*DcimPowerOutletTemplatesBulkDeleteNoContent, error)

	DcimPowerOutletTemplatesBulkPartialUpdate(params *DcimPowerOutletTemplatesBulkPartialUpdateParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*DcimPowerOutletTemplatesBulkPartialUpdateOK, error)

	DcimPowerOutletTemplatesBulkUpdate(params *DcimPowerOutletTemplatesBulkUpdateParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*DcimPowerOutletTemplatesBulkUpdateOK, error)

	DcimPowerOutletTemplatesCreate(params *DcimPowerOutletTemplatesCreateParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*DcimPowerOutletTemplatesCreateCreated, error)

	DcimPowerOutletTemplatesDelete(params *DcimPowerOutletTemplatesDeleteParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*DcimPowerOutletTemplatesDeleteNoContent, error)

	DcimPowerOutletTemplatesList(params *DcimPowerOutletTemplatesListParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*DcimPowerOutletTemplatesListOK, error)

	DcimPowerOutletTemplatesPartialUpdate(params *DcimPowerOutletTemplatesPartialUpdateParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*DcimPowerOutletTemplatesPartialUpdateOK, error)

	DcimPowerOutletTemplatesRead(params *DcimPowerOutletTemplatesReadParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*DcimPowerOutletTemplatesReadOK, error)

	DcimPowerOutletTemplatesUpdate(params *DcimPowerOutletTemplatesUpdateParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*DcimPowerOutletTemplatesUpdateOK, error)

	DcimPowerOutletsBulkDelete(params *DcimPowerOutletsBulkDeleteParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*DcimPowerOutletsBulkDeleteNoContent, error)

	DcimPowerOutletsBulkPartialUpdate(params *DcimPowerOutletsBulkPartialUpdateParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*DcimPowerOutletsBulkPartialUpdateOK, error)

	DcimPowerOutletsBulkUpdate(params *DcimPowerOutletsBulkUpdateParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*DcimPowerOutletsBulkUpdateOK, error)

	DcimPowerOutletsCreate(params *DcimPowerOutletsCreateParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*DcimPowerOutletsCreateCreated, error)

	DcimPowerOutletsDelete(params *DcimPowerOutletsDeleteParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*DcimPowerOutletsDeleteNoContent, error)

	DcimPowerOutletsList(params *DcimPowerOutletsListParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*DcimPowerOutletsListOK, error)

	DcimPowerOutletsPartialUpdate(params *DcimPowerOutletsPartialUpdateParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*DcimPowerOutletsPartialUpdateOK, error)

	DcimPowerOutletsRead(params *DcimPowerOutletsReadParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*DcimPowerOutletsReadOK, error)

	DcimPowerOutletsTrace(params *DcimPowerOutletsTraceParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*DcimPowerOutletsTraceOK, error)

	DcimPowerOutletsUpdate(params *DcimPowerOutletsUpdateParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*DcimPowerOutletsUpdateOK, error)

	DcimPowerPanelsBulkDelete(params *DcimPowerPanelsBulkDeleteParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*DcimPowerPanelsBulkDeleteNoContent, error)

	DcimPowerPanelsBulkPartialUpdate(params *DcimPowerPanelsBulkPartialUpdateParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*DcimPowerPanelsBulkPartialUpdateOK, error)

	DcimPowerPanelsBulkUpdate(params *DcimPowerPanelsBulkUpdateParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*DcimPowerPanelsBulkUpdateOK, error)

	DcimPowerPanelsCreate(params *DcimPowerPanelsCreateParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*DcimPowerPanelsCreateCreated, error)

	DcimPowerPanelsDelete(params *DcimPowerPanelsDeleteParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*DcimPowerPanelsDeleteNoContent, error)

	DcimPowerPanelsList(params *DcimPowerPanelsListParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*DcimPowerPanelsListOK, error)

	DcimPowerPanelsPartialUpdate(params *DcimPowerPanelsPartialUpdateParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*DcimPowerPanelsPartialUpdateOK, error)

	DcimPowerPanelsRead(params *DcimPowerPanelsReadParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*DcimPowerPanelsReadOK, error)

	DcimPowerPanelsUpdate(params *DcimPowerPanelsUpdateParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*DcimPowerPanelsUpdateOK, error)

	DcimPowerPortTemplatesBulkDelete(params *DcimPowerPortTemplatesBulkDeleteParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*DcimPowerPortTemplatesBulkDeleteNoContent, error)

	DcimPowerPortTemplatesBulkPartialUpdate(params *DcimPowerPortTemplatesBulkPartialUpdateParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*DcimPowerPortTemplatesBulkPartialUpdateOK, error)

	DcimPowerPortTemplatesBulkUpdate(params *DcimPowerPortTemplatesBulkUpdateParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*DcimPowerPortTemplatesBulkUpdateOK, error)

	DcimPowerPortTemplatesCreate(params *DcimPowerPortTemplatesCreateParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*DcimPowerPortTemplatesCreateCreated, error)

	DcimPowerPortTemplatesDelete(params *DcimPowerPortTemplatesDeleteParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*DcimPowerPortTemplatesDeleteNoContent, error)

	DcimPowerPortTemplatesList(params *DcimPowerPortTemplatesListParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*DcimPowerPortTemplatesListOK, error)

	DcimPowerPortTemplatesPartialUpdate(params *DcimPowerPortTemplatesPartialUpdateParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*DcimPowerPortTemplatesPartialUpdateOK, error)

	DcimPowerPortTemplatesRead(params *DcimPowerPortTemplatesReadParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*DcimPowerPortTemplatesReadOK, error)

	DcimPowerPortTemplatesUpdate(params *DcimPowerPortTemplatesUpdateParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*DcimPowerPortTemplatesUpdateOK, error)

	DcimPowerPortsBulkDelete(params *DcimPowerPortsBulkDeleteParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*DcimPowerPortsBulkDeleteNoContent, error)

	DcimPowerPortsBulkPartialUpdate(params *DcimPowerPortsBulkPartialUpdateParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*DcimPowerPortsBulkPartialUpdateOK, error)

	DcimPowerPortsBulkUpdate(params *DcimPowerPortsBulkUpdateParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*DcimPowerPortsBulkUpdateOK, error)

	DcimPowerPortsCreate(params *DcimPowerPortsCreateParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*DcimPowerPortsCreateCreated, error)

	DcimPowerPortsDelete(params *DcimPowerPortsDeleteParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*DcimPowerPortsDeleteNoContent, error)

	DcimPowerPortsList(params *DcimPowerPortsListParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*DcimPowerPortsListOK, error)

	DcimPowerPortsPartialUpdate(params *DcimPowerPortsPartialUpdateParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*DcimPowerPortsPartialUpdateOK, error)

	DcimPowerPortsRead(params *DcimPowerPortsReadParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*DcimPowerPortsReadOK, error)

	DcimPowerPortsTrace(params *DcimPowerPortsTraceParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*DcimPowerPortsTraceOK, error)

	DcimPowerPortsUpdate(params *DcimPowerPortsUpdateParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*DcimPowerPortsUpdateOK, error)

	DcimRackGroupsBulkDelete(params *DcimRackGroupsBulkDeleteParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*DcimRackGroupsBulkDeleteNoContent, error)

	DcimRackGroupsBulkPartialUpdate(params *DcimRackGroupsBulkPartialUpdateParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*DcimRackGroupsBulkPartialUpdateOK, error)

	DcimRackGroupsBulkUpdate(params *DcimRackGroupsBulkUpdateParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*DcimRackGroupsBulkUpdateOK, error)

	DcimRackGroupsCreate(params *DcimRackGroupsCreateParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*DcimRackGroupsCreateCreated, error)

	DcimRackGroupsDelete(params *DcimRackGroupsDeleteParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*DcimRackGroupsDeleteNoContent, error)

	DcimRackGroupsList(params *DcimRackGroupsListParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*DcimRackGroupsListOK, error)

	DcimRackGroupsPartialUpdate(params *DcimRackGroupsPartialUpdateParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*DcimRackGroupsPartialUpdateOK, error)

	DcimRackGroupsRead(params *DcimRackGroupsReadParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*DcimRackGroupsReadOK, error)

	DcimRackGroupsUpdate(params *DcimRackGroupsUpdateParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*DcimRackGroupsUpdateOK, error)

	DcimRackReservationsBulkDelete(params *DcimRackReservationsBulkDeleteParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*DcimRackReservationsBulkDeleteNoContent, error)

	DcimRackReservationsBulkPartialUpdate(params *DcimRackReservationsBulkPartialUpdateParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*DcimRackReservationsBulkPartialUpdateOK, error)

	DcimRackReservationsBulkUpdate(params *DcimRackReservationsBulkUpdateParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*DcimRackReservationsBulkUpdateOK, error)

	DcimRackReservationsCreate(params *DcimRackReservationsCreateParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*DcimRackReservationsCreateCreated, error)

	DcimRackReservationsDelete(params *DcimRackReservationsDeleteParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*DcimRackReservationsDeleteNoContent, error)

	DcimRackReservationsList(params *DcimRackReservationsListParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*DcimRackReservationsListOK, error)

	DcimRackReservationsPartialUpdate(params *DcimRackReservationsPartialUpdateParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*DcimRackReservationsPartialUpdateOK, error)

	DcimRackReservationsRead(params *DcimRackReservationsReadParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*DcimRackReservationsReadOK, error)

	DcimRackReservationsUpdate(params *DcimRackReservationsUpdateParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*DcimRackReservationsUpdateOK, error)

	DcimRackRolesBulkDelete(params *DcimRackRolesBulkDeleteParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*DcimRackRolesBulkDeleteNoContent, error)

	DcimRackRolesBulkPartialUpdate(params *DcimRackRolesBulkPartialUpdateParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*DcimRackRolesBulkPartialUpdateOK, error)

	DcimRackRolesBulkUpdate(params *DcimRackRolesBulkUpdateParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*DcimRackRolesBulkUpdateOK, error)

	DcimRackRolesCreate(params *DcimRackRolesCreateParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*DcimRackRolesCreateCreated, error)

	DcimRackRolesDelete(params *DcimRackRolesDeleteParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*DcimRackRolesDeleteNoContent, error)

	DcimRackRolesList(params *DcimRackRolesListParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*DcimRackRolesListOK, error)

	DcimRackRolesPartialUpdate(params *DcimRackRolesPartialUpdateParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*DcimRackRolesPartialUpdateOK, error)

	DcimRackRolesRead(params *DcimRackRolesReadParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*DcimRackRolesReadOK, error)

	DcimRackRolesUpdate(params *DcimRackRolesUpdateParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*DcimRackRolesUpdateOK, error)

	DcimRacksBulkDelete(params *DcimRacksBulkDeleteParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*DcimRacksBulkDeleteNoContent, error)

	DcimRacksBulkPartialUpdate(params *DcimRacksBulkPartialUpdateParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*DcimRacksBulkPartialUpdateOK, error)

	DcimRacksBulkUpdate(params *DcimRacksBulkUpdateParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*DcimRacksBulkUpdateOK, error)

	DcimRacksCreate(params *DcimRacksCreateParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*DcimRacksCreateCreated, error)

	DcimRacksDelete(params *DcimRacksDeleteParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*DcimRacksDeleteNoContent, error)

	DcimRacksElevationRead(params *DcimRacksElevationReadParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*DcimRacksElevationReadOK, error)

	DcimRacksList(params *DcimRacksListParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*DcimRacksListOK, error)

	DcimRacksPartialUpdate(params *DcimRacksPartialUpdateParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*DcimRacksPartialUpdateOK, error)

	DcimRacksRead(params *DcimRacksReadParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*DcimRacksReadOK, error)

	DcimRacksUpdate(params *DcimRacksUpdateParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*DcimRacksUpdateOK, error)

	DcimRearPortTemplatesBulkDelete(params *DcimRearPortTemplatesBulkDeleteParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*DcimRearPortTemplatesBulkDeleteNoContent, error)

	DcimRearPortTemplatesBulkPartialUpdate(params *DcimRearPortTemplatesBulkPartialUpdateParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*DcimRearPortTemplatesBulkPartialUpdateOK, error)

	DcimRearPortTemplatesBulkUpdate(params *DcimRearPortTemplatesBulkUpdateParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*DcimRearPortTemplatesBulkUpdateOK, error)

	DcimRearPortTemplatesCreate(params *DcimRearPortTemplatesCreateParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*DcimRearPortTemplatesCreateCreated, error)

	DcimRearPortTemplatesDelete(params *DcimRearPortTemplatesDeleteParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*DcimRearPortTemplatesDeleteNoContent, error)

	DcimRearPortTemplatesList(params *DcimRearPortTemplatesListParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*DcimRearPortTemplatesListOK, error)

	DcimRearPortTemplatesPartialUpdate(params *DcimRearPortTemplatesPartialUpdateParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*DcimRearPortTemplatesPartialUpdateOK, error)

	DcimRearPortTemplatesRead(params *DcimRearPortTemplatesReadParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*DcimRearPortTemplatesReadOK, error)

	DcimRearPortTemplatesUpdate(params *DcimRearPortTemplatesUpdateParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*DcimRearPortTemplatesUpdateOK, error)

	DcimRearPortsBulkDelete(params *DcimRearPortsBulkDeleteParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*DcimRearPortsBulkDeleteNoContent, error)

	DcimRearPortsBulkPartialUpdate(params *DcimRearPortsBulkPartialUpdateParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*DcimRearPortsBulkPartialUpdateOK, error)

	DcimRearPortsBulkUpdate(params *DcimRearPortsBulkUpdateParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*DcimRearPortsBulkUpdateOK, error)

	DcimRearPortsCreate(params *DcimRearPortsCreateParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*DcimRearPortsCreateCreated, error)

	DcimRearPortsDelete(params *DcimRearPortsDeleteParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*DcimRearPortsDeleteNoContent, error)

	DcimRearPortsList(params *DcimRearPortsListParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*DcimRearPortsListOK, error)

	DcimRearPortsPartialUpdate(params *DcimRearPortsPartialUpdateParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*DcimRearPortsPartialUpdateOK, error)

	DcimRearPortsPaths(params *DcimRearPortsPathsParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*DcimRearPortsPathsOK, error)

	DcimRearPortsRead(params *DcimRearPortsReadParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*DcimRearPortsReadOK, error)

	DcimRearPortsUpdate(params *DcimRearPortsUpdateParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*DcimRearPortsUpdateOK, error)

	DcimRegionsBulkDelete(params *DcimRegionsBulkDeleteParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*DcimRegionsBulkDeleteNoContent, error)

	DcimRegionsBulkPartialUpdate(params *DcimRegionsBulkPartialUpdateParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*DcimRegionsBulkPartialUpdateOK, error)

	DcimRegionsBulkUpdate(params *DcimRegionsBulkUpdateParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*DcimRegionsBulkUpdateOK, error)

	DcimRegionsCreate(params *DcimRegionsCreateParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*DcimRegionsCreateCreated, error)

	DcimRegionsDelete(params *DcimRegionsDeleteParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*DcimRegionsDeleteNoContent, error)

	DcimRegionsList(params *DcimRegionsListParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*DcimRegionsListOK, error)

	DcimRegionsPartialUpdate(params *DcimRegionsPartialUpdateParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*DcimRegionsPartialUpdateOK, error)

	DcimRegionsRead(params *DcimRegionsReadParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*DcimRegionsReadOK, error)

	DcimRegionsUpdate(params *DcimRegionsUpdateParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*DcimRegionsUpdateOK, error)

	DcimSitesBulkDelete(params *DcimSitesBulkDeleteParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*DcimSitesBulkDeleteNoContent, error)

	DcimSitesBulkPartialUpdate(params *DcimSitesBulkPartialUpdateParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*DcimSitesBulkPartialUpdateOK, error)

	DcimSitesBulkUpdate(params *DcimSitesBulkUpdateParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*DcimSitesBulkUpdateOK, error)

	DcimSitesCreate(params *DcimSitesCreateParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*DcimSitesCreateCreated, error)

	DcimSitesDelete(params *DcimSitesDeleteParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*DcimSitesDeleteNoContent, error)

	DcimSitesList(params *DcimSitesListParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*DcimSitesListOK, error)

	DcimSitesPartialUpdate(params *DcimSitesPartialUpdateParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*DcimSitesPartialUpdateOK, error)

	DcimSitesRead(params *DcimSitesReadParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*DcimSitesReadOK, error)

	DcimSitesUpdate(params *DcimSitesUpdateParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*DcimSitesUpdateOK, error)

	DcimVirtualChassisBulkDelete(params *DcimVirtualChassisBulkDeleteParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*DcimVirtualChassisBulkDeleteNoContent, error)

	DcimVirtualChassisBulkPartialUpdate(params *DcimVirtualChassisBulkPartialUpdateParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*DcimVirtualChassisBulkPartialUpdateOK, error)

	DcimVirtualChassisBulkUpdate(params *DcimVirtualChassisBulkUpdateParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*DcimVirtualChassisBulkUpdateOK, error)

	DcimVirtualChassisCreate(params *DcimVirtualChassisCreateParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*DcimVirtualChassisCreateCreated, error)

	DcimVirtualChassisDelete(params *DcimVirtualChassisDeleteParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*DcimVirtualChassisDeleteNoContent, error)

	DcimVirtualChassisList(params *DcimVirtualChassisListParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*DcimVirtualChassisListOK, error)

	DcimVirtualChassisPartialUpdate(params *DcimVirtualChassisPartialUpdateParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*DcimVirtualChassisPartialUpdateOK, error)

	DcimVirtualChassisRead(params *DcimVirtualChassisReadParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*DcimVirtualChassisReadOK, error)

	DcimVirtualChassisUpdate(params *DcimVirtualChassisUpdateParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*DcimVirtualChassisUpdateOK, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
  DcimCablesBulkDelete dcim cables bulk delete API
*/
func (a *Client) DcimCablesBulkDelete(params *DcimCablesBulkDeleteParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*DcimCablesBulkDeleteNoContent, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDcimCablesBulkDeleteParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "dcim_cables_bulk_delete",
		Method:             "DELETE",
		PathPattern:        "/dcim/cables/",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &DcimCablesBulkDeleteReader{formats: a.formats},
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
	success, ok := result.(*DcimCablesBulkDeleteNoContent)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for dcim_cables_bulk_delete: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  DcimCablesBulkPartialUpdate dcim cables bulk partial update API
*/
func (a *Client) DcimCablesBulkPartialUpdate(params *DcimCablesBulkPartialUpdateParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*DcimCablesBulkPartialUpdateOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDcimCablesBulkPartialUpdateParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "dcim_cables_bulk_partial_update",
		Method:             "PATCH",
		PathPattern:        "/dcim/cables/",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &DcimCablesBulkPartialUpdateReader{formats: a.formats},
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
	success, ok := result.(*DcimCablesBulkPartialUpdateOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for dcim_cables_bulk_partial_update: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  DcimCablesBulkUpdate dcim cables bulk update API
*/
func (a *Client) DcimCablesBulkUpdate(params *DcimCablesBulkUpdateParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*DcimCablesBulkUpdateOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDcimCablesBulkUpdateParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "dcim_cables_bulk_update",
		Method:             "PUT",
		PathPattern:        "/dcim/cables/",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &DcimCablesBulkUpdateReader{formats: a.formats},
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
	success, ok := result.(*DcimCablesBulkUpdateOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for dcim_cables_bulk_update: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  DcimCablesCreate dcim cables create API
*/
func (a *Client) DcimCablesCreate(params *DcimCablesCreateParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*DcimCablesCreateCreated, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDcimCablesCreateParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "dcim_cables_create",
		Method:             "POST",
		PathPattern:        "/dcim/cables/",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &DcimCablesCreateReader{formats: a.formats},
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
	success, ok := result.(*DcimCablesCreateCreated)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for dcim_cables_create: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  DcimCablesDelete dcim cables delete API
*/
func (a *Client) DcimCablesDelete(params *DcimCablesDeleteParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*DcimCablesDeleteNoContent, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDcimCablesDeleteParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "dcim_cables_delete",
		Method:             "DELETE",
		PathPattern:        "/dcim/cables/{id}/",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &DcimCablesDeleteReader{formats: a.formats},
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
	success, ok := result.(*DcimCablesDeleteNoContent)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for dcim_cables_delete: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  DcimCablesList dcim cables list API
*/
func (a *Client) DcimCablesList(params *DcimCablesListParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*DcimCablesListOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDcimCablesListParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "dcim_cables_list",
		Method:             "GET",
		PathPattern:        "/dcim/cables/",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &DcimCablesListReader{formats: a.formats},
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
	success, ok := result.(*DcimCablesListOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for dcim_cables_list: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  DcimCablesPartialUpdate dcim cables partial update API
*/
func (a *Client) DcimCablesPartialUpdate(params *DcimCablesPartialUpdateParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*DcimCablesPartialUpdateOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDcimCablesPartialUpdateParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "dcim_cables_partial_update",
		Method:             "PATCH",
		PathPattern:        "/dcim/cables/{id}/",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &DcimCablesPartialUpdateReader{formats: a.formats},
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
	success, ok := result.(*DcimCablesPartialUpdateOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for dcim_cables_partial_update: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  DcimCablesRead dcim cables read API
*/
func (a *Client) DcimCablesRead(params *DcimCablesReadParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*DcimCablesReadOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDcimCablesReadParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "dcim_cables_read",
		Method:             "GET",
		PathPattern:        "/dcim/cables/{id}/",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &DcimCablesReadReader{formats: a.formats},
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
	success, ok := result.(*DcimCablesReadOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for dcim_cables_read: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  DcimCablesUpdate dcim cables update API
*/
func (a *Client) DcimCablesUpdate(params *DcimCablesUpdateParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*DcimCablesUpdateOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDcimCablesUpdateParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "dcim_cables_update",
		Method:             "PUT",
		PathPattern:        "/dcim/cables/{id}/",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &DcimCablesUpdateReader{formats: a.formats},
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
	success, ok := result.(*DcimCablesUpdateOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for dcim_cables_update: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  DcimConnectedDeviceList This endpoint allows a user to determine what device (if any) is connected to a given peer device and peer
interface. This is useful in a situation where a device boots with no configuration, but can detect its neighbors
via a protocol such as LLDP. Two query parameters must be included in the request:

* `peer_device`: The name of the peer device
* `peer_interface`: The name of the peer interface
*/
func (a *Client) DcimConnectedDeviceList(params *DcimConnectedDeviceListParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*DcimConnectedDeviceListOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDcimConnectedDeviceListParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "dcim_connected-device_list",
		Method:             "GET",
		PathPattern:        "/dcim/connected-device/",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &DcimConnectedDeviceListReader{formats: a.formats},
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
	success, ok := result.(*DcimConnectedDeviceListOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for dcim_connected-device_list: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  DcimConsoleConnectionsList dcim console connections list API
*/
func (a *Client) DcimConsoleConnectionsList(params *DcimConsoleConnectionsListParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*DcimConsoleConnectionsListOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDcimConsoleConnectionsListParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "dcim_console-connections_list",
		Method:             "GET",
		PathPattern:        "/dcim/console-connections/",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &DcimConsoleConnectionsListReader{formats: a.formats},
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
	success, ok := result.(*DcimConsoleConnectionsListOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for dcim_console-connections_list: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  DcimConsolePortTemplatesBulkDelete dcim console port templates bulk delete API
*/
func (a *Client) DcimConsolePortTemplatesBulkDelete(params *DcimConsolePortTemplatesBulkDeleteParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*DcimConsolePortTemplatesBulkDeleteNoContent, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDcimConsolePortTemplatesBulkDeleteParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "dcim_console-port-templates_bulk_delete",
		Method:             "DELETE",
		PathPattern:        "/dcim/console-port-templates/",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &DcimConsolePortTemplatesBulkDeleteReader{formats: a.formats},
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
	success, ok := result.(*DcimConsolePortTemplatesBulkDeleteNoContent)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for dcim_console-port-templates_bulk_delete: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  DcimConsolePortTemplatesBulkPartialUpdate dcim console port templates bulk partial update API
*/
func (a *Client) DcimConsolePortTemplatesBulkPartialUpdate(params *DcimConsolePortTemplatesBulkPartialUpdateParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*DcimConsolePortTemplatesBulkPartialUpdateOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDcimConsolePortTemplatesBulkPartialUpdateParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "dcim_console-port-templates_bulk_partial_update",
		Method:             "PATCH",
		PathPattern:        "/dcim/console-port-templates/",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &DcimConsolePortTemplatesBulkPartialUpdateReader{formats: a.formats},
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
	success, ok := result.(*DcimConsolePortTemplatesBulkPartialUpdateOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for dcim_console-port-templates_bulk_partial_update: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  DcimConsolePortTemplatesBulkUpdate dcim console port templates bulk update API
*/
func (a *Client) DcimConsolePortTemplatesBulkUpdate(params *DcimConsolePortTemplatesBulkUpdateParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*DcimConsolePortTemplatesBulkUpdateOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDcimConsolePortTemplatesBulkUpdateParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "dcim_console-port-templates_bulk_update",
		Method:             "PUT",
		PathPattern:        "/dcim/console-port-templates/",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &DcimConsolePortTemplatesBulkUpdateReader{formats: a.formats},
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
	success, ok := result.(*DcimConsolePortTemplatesBulkUpdateOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for dcim_console-port-templates_bulk_update: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  DcimConsolePortTemplatesCreate dcim console port templates create API
*/
func (a *Client) DcimConsolePortTemplatesCreate(params *DcimConsolePortTemplatesCreateParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*DcimConsolePortTemplatesCreateCreated, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDcimConsolePortTemplatesCreateParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "dcim_console-port-templates_create",
		Method:             "POST",
		PathPattern:        "/dcim/console-port-templates/",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &DcimConsolePortTemplatesCreateReader{formats: a.formats},
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
	success, ok := result.(*DcimConsolePortTemplatesCreateCreated)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for dcim_console-port-templates_create: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  DcimConsolePortTemplatesDelete dcim console port templates delete API
*/
func (a *Client) DcimConsolePortTemplatesDelete(params *DcimConsolePortTemplatesDeleteParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*DcimConsolePortTemplatesDeleteNoContent, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDcimConsolePortTemplatesDeleteParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "dcim_console-port-templates_delete",
		Method:             "DELETE",
		PathPattern:        "/dcim/console-port-templates/{id}/",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &DcimConsolePortTemplatesDeleteReader{formats: a.formats},
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
	success, ok := result.(*DcimConsolePortTemplatesDeleteNoContent)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for dcim_console-port-templates_delete: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  DcimConsolePortTemplatesList dcim console port templates list API
*/
func (a *Client) DcimConsolePortTemplatesList(params *DcimConsolePortTemplatesListParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*DcimConsolePortTemplatesListOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDcimConsolePortTemplatesListParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "dcim_console-port-templates_list",
		Method:             "GET",
		PathPattern:        "/dcim/console-port-templates/",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &DcimConsolePortTemplatesListReader{formats: a.formats},
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
	success, ok := result.(*DcimConsolePortTemplatesListOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for dcim_console-port-templates_list: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  DcimConsolePortTemplatesPartialUpdate dcim console port templates partial update API
*/
func (a *Client) DcimConsolePortTemplatesPartialUpdate(params *DcimConsolePortTemplatesPartialUpdateParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*DcimConsolePortTemplatesPartialUpdateOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDcimConsolePortTemplatesPartialUpdateParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "dcim_console-port-templates_partial_update",
		Method:             "PATCH",
		PathPattern:        "/dcim/console-port-templates/{id}/",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &DcimConsolePortTemplatesPartialUpdateReader{formats: a.formats},
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
	success, ok := result.(*DcimConsolePortTemplatesPartialUpdateOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for dcim_console-port-templates_partial_update: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  DcimConsolePortTemplatesRead dcim console port templates read API
*/
func (a *Client) DcimConsolePortTemplatesRead(params *DcimConsolePortTemplatesReadParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*DcimConsolePortTemplatesReadOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDcimConsolePortTemplatesReadParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "dcim_console-port-templates_read",
		Method:             "GET",
		PathPattern:        "/dcim/console-port-templates/{id}/",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &DcimConsolePortTemplatesReadReader{formats: a.formats},
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
	success, ok := result.(*DcimConsolePortTemplatesReadOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for dcim_console-port-templates_read: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  DcimConsolePortTemplatesUpdate dcim console port templates update API
*/
func (a *Client) DcimConsolePortTemplatesUpdate(params *DcimConsolePortTemplatesUpdateParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*DcimConsolePortTemplatesUpdateOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDcimConsolePortTemplatesUpdateParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "dcim_console-port-templates_update",
		Method:             "PUT",
		PathPattern:        "/dcim/console-port-templates/{id}/",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &DcimConsolePortTemplatesUpdateReader{formats: a.formats},
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
	success, ok := result.(*DcimConsolePortTemplatesUpdateOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for dcim_console-port-templates_update: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  DcimConsolePortsBulkDelete dcim console ports bulk delete API
*/
func (a *Client) DcimConsolePortsBulkDelete(params *DcimConsolePortsBulkDeleteParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*DcimConsolePortsBulkDeleteNoContent, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDcimConsolePortsBulkDeleteParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "dcim_console-ports_bulk_delete",
		Method:             "DELETE",
		PathPattern:        "/dcim/console-ports/",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &DcimConsolePortsBulkDeleteReader{formats: a.formats},
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
	success, ok := result.(*DcimConsolePortsBulkDeleteNoContent)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for dcim_console-ports_bulk_delete: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  DcimConsolePortsBulkPartialUpdate dcim console ports bulk partial update API
*/
func (a *Client) DcimConsolePortsBulkPartialUpdate(params *DcimConsolePortsBulkPartialUpdateParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*DcimConsolePortsBulkPartialUpdateOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDcimConsolePortsBulkPartialUpdateParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "dcim_console-ports_bulk_partial_update",
		Method:             "PATCH",
		PathPattern:        "/dcim/console-ports/",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &DcimConsolePortsBulkPartialUpdateReader{formats: a.formats},
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
	success, ok := result.(*DcimConsolePortsBulkPartialUpdateOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for dcim_console-ports_bulk_partial_update: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  DcimConsolePortsBulkUpdate dcim console ports bulk update API
*/
func (a *Client) DcimConsolePortsBulkUpdate(params *DcimConsolePortsBulkUpdateParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*DcimConsolePortsBulkUpdateOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDcimConsolePortsBulkUpdateParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "dcim_console-ports_bulk_update",
		Method:             "PUT",
		PathPattern:        "/dcim/console-ports/",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &DcimConsolePortsBulkUpdateReader{formats: a.formats},
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
	success, ok := result.(*DcimConsolePortsBulkUpdateOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for dcim_console-ports_bulk_update: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  DcimConsolePortsCreate dcim console ports create API
*/
func (a *Client) DcimConsolePortsCreate(params *DcimConsolePortsCreateParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*DcimConsolePortsCreateCreated, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDcimConsolePortsCreateParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "dcim_console-ports_create",
		Method:             "POST",
		PathPattern:        "/dcim/console-ports/",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &DcimConsolePortsCreateReader{formats: a.formats},
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
	success, ok := result.(*DcimConsolePortsCreateCreated)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for dcim_console-ports_create: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  DcimConsolePortsDelete dcim console ports delete API
*/
func (a *Client) DcimConsolePortsDelete(params *DcimConsolePortsDeleteParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*DcimConsolePortsDeleteNoContent, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDcimConsolePortsDeleteParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "dcim_console-ports_delete",
		Method:             "DELETE",
		PathPattern:        "/dcim/console-ports/{id}/",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &DcimConsolePortsDeleteReader{formats: a.formats},
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
	success, ok := result.(*DcimConsolePortsDeleteNoContent)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for dcim_console-ports_delete: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  DcimConsolePortsList dcim console ports list API
*/
func (a *Client) DcimConsolePortsList(params *DcimConsolePortsListParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*DcimConsolePortsListOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDcimConsolePortsListParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "dcim_console-ports_list",
		Method:             "GET",
		PathPattern:        "/dcim/console-ports/",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &DcimConsolePortsListReader{formats: a.formats},
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
	success, ok := result.(*DcimConsolePortsListOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for dcim_console-ports_list: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  DcimConsolePortsPartialUpdate dcim console ports partial update API
*/
func (a *Client) DcimConsolePortsPartialUpdate(params *DcimConsolePortsPartialUpdateParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*DcimConsolePortsPartialUpdateOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDcimConsolePortsPartialUpdateParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "dcim_console-ports_partial_update",
		Method:             "PATCH",
		PathPattern:        "/dcim/console-ports/{id}/",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &DcimConsolePortsPartialUpdateReader{formats: a.formats},
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
	success, ok := result.(*DcimConsolePortsPartialUpdateOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for dcim_console-ports_partial_update: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  DcimConsolePortsRead dcim console ports read API
*/
func (a *Client) DcimConsolePortsRead(params *DcimConsolePortsReadParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*DcimConsolePortsReadOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDcimConsolePortsReadParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "dcim_console-ports_read",
		Method:             "GET",
		PathPattern:        "/dcim/console-ports/{id}/",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &DcimConsolePortsReadReader{formats: a.formats},
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
	success, ok := result.(*DcimConsolePortsReadOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for dcim_console-ports_read: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  DcimConsolePortsTrace Trace a complete cable path and return each segment as a three-tuple of (termination, cable, termination).
*/
func (a *Client) DcimConsolePortsTrace(params *DcimConsolePortsTraceParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*DcimConsolePortsTraceOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDcimConsolePortsTraceParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "dcim_console-ports_trace",
		Method:             "GET",
		PathPattern:        "/dcim/console-ports/{id}/trace/",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &DcimConsolePortsTraceReader{formats: a.formats},
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
	success, ok := result.(*DcimConsolePortsTraceOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for dcim_console-ports_trace: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  DcimConsolePortsUpdate dcim console ports update API
*/
func (a *Client) DcimConsolePortsUpdate(params *DcimConsolePortsUpdateParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*DcimConsolePortsUpdateOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDcimConsolePortsUpdateParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "dcim_console-ports_update",
		Method:             "PUT",
		PathPattern:        "/dcim/console-ports/{id}/",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &DcimConsolePortsUpdateReader{formats: a.formats},
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
	success, ok := result.(*DcimConsolePortsUpdateOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for dcim_console-ports_update: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  DcimConsoleServerPortTemplatesBulkDelete dcim console server port templates bulk delete API
*/
func (a *Client) DcimConsoleServerPortTemplatesBulkDelete(params *DcimConsoleServerPortTemplatesBulkDeleteParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*DcimConsoleServerPortTemplatesBulkDeleteNoContent, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDcimConsoleServerPortTemplatesBulkDeleteParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "dcim_console-server-port-templates_bulk_delete",
		Method:             "DELETE",
		PathPattern:        "/dcim/console-server-port-templates/",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &DcimConsoleServerPortTemplatesBulkDeleteReader{formats: a.formats},
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
	success, ok := result.(*DcimConsoleServerPortTemplatesBulkDeleteNoContent)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for dcim_console-server-port-templates_bulk_delete: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  DcimConsoleServerPortTemplatesBulkPartialUpdate dcim console server port templates bulk partial update API
*/
func (a *Client) DcimConsoleServerPortTemplatesBulkPartialUpdate(params *DcimConsoleServerPortTemplatesBulkPartialUpdateParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*DcimConsoleServerPortTemplatesBulkPartialUpdateOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDcimConsoleServerPortTemplatesBulkPartialUpdateParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "dcim_console-server-port-templates_bulk_partial_update",
		Method:             "PATCH",
		PathPattern:        "/dcim/console-server-port-templates/",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &DcimConsoleServerPortTemplatesBulkPartialUpdateReader{formats: a.formats},
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
	success, ok := result.(*DcimConsoleServerPortTemplatesBulkPartialUpdateOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for dcim_console-server-port-templates_bulk_partial_update: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  DcimConsoleServerPortTemplatesBulkUpdate dcim console server port templates bulk update API
*/
func (a *Client) DcimConsoleServerPortTemplatesBulkUpdate(params *DcimConsoleServerPortTemplatesBulkUpdateParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*DcimConsoleServerPortTemplatesBulkUpdateOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDcimConsoleServerPortTemplatesBulkUpdateParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "dcim_console-server-port-templates_bulk_update",
		Method:             "PUT",
		PathPattern:        "/dcim/console-server-port-templates/",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &DcimConsoleServerPortTemplatesBulkUpdateReader{formats: a.formats},
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
	success, ok := result.(*DcimConsoleServerPortTemplatesBulkUpdateOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for dcim_console-server-port-templates_bulk_update: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  DcimConsoleServerPortTemplatesCreate dcim console server port templates create API
*/
func (a *Client) DcimConsoleServerPortTemplatesCreate(params *DcimConsoleServerPortTemplatesCreateParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*DcimConsoleServerPortTemplatesCreateCreated, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDcimConsoleServerPortTemplatesCreateParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "dcim_console-server-port-templates_create",
		Method:             "POST",
		PathPattern:        "/dcim/console-server-port-templates/",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &DcimConsoleServerPortTemplatesCreateReader{formats: a.formats},
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
	success, ok := result.(*DcimConsoleServerPortTemplatesCreateCreated)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for dcim_console-server-port-templates_create: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  DcimConsoleServerPortTemplatesDelete dcim console server port templates delete API
*/
func (a *Client) DcimConsoleServerPortTemplatesDelete(params *DcimConsoleServerPortTemplatesDeleteParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*DcimConsoleServerPortTemplatesDeleteNoContent, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDcimConsoleServerPortTemplatesDeleteParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "dcim_console-server-port-templates_delete",
		Method:             "DELETE",
		PathPattern:        "/dcim/console-server-port-templates/{id}/",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &DcimConsoleServerPortTemplatesDeleteReader{formats: a.formats},
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
	success, ok := result.(*DcimConsoleServerPortTemplatesDeleteNoContent)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for dcim_console-server-port-templates_delete: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  DcimConsoleServerPortTemplatesList dcim console server port templates list API
*/
func (a *Client) DcimConsoleServerPortTemplatesList(params *DcimConsoleServerPortTemplatesListParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*DcimConsoleServerPortTemplatesListOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDcimConsoleServerPortTemplatesListParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "dcim_console-server-port-templates_list",
		Method:             "GET",
		PathPattern:        "/dcim/console-server-port-templates/",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &DcimConsoleServerPortTemplatesListReader{formats: a.formats},
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
	success, ok := result.(*DcimConsoleServerPortTemplatesListOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for dcim_console-server-port-templates_list: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  DcimConsoleServerPortTemplatesPartialUpdate dcim console server port templates partial update API
*/
func (a *Client) DcimConsoleServerPortTemplatesPartialUpdate(params *DcimConsoleServerPortTemplatesPartialUpdateParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*DcimConsoleServerPortTemplatesPartialUpdateOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDcimConsoleServerPortTemplatesPartialUpdateParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "dcim_console-server-port-templates_partial_update",
		Method:             "PATCH",
		PathPattern:        "/dcim/console-server-port-templates/{id}/",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &DcimConsoleServerPortTemplatesPartialUpdateReader{formats: a.formats},
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
	success, ok := result.(*DcimConsoleServerPortTemplatesPartialUpdateOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for dcim_console-server-port-templates_partial_update: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  DcimConsoleServerPortTemplatesRead dcim console server port templates read API
*/
func (a *Client) DcimConsoleServerPortTemplatesRead(params *DcimConsoleServerPortTemplatesReadParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*DcimConsoleServerPortTemplatesReadOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDcimConsoleServerPortTemplatesReadParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "dcim_console-server-port-templates_read",
		Method:             "GET",
		PathPattern:        "/dcim/console-server-port-templates/{id}/",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &DcimConsoleServerPortTemplatesReadReader{formats: a.formats},
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
	success, ok := result.(*DcimConsoleServerPortTemplatesReadOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for dcim_console-server-port-templates_read: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  DcimConsoleServerPortTemplatesUpdate dcim console server port templates update API
*/
func (a *Client) DcimConsoleServerPortTemplatesUpdate(params *DcimConsoleServerPortTemplatesUpdateParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*DcimConsoleServerPortTemplatesUpdateOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDcimConsoleServerPortTemplatesUpdateParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "dcim_console-server-port-templates_update",
		Method:             "PUT",
		PathPattern:        "/dcim/console-server-port-templates/{id}/",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &DcimConsoleServerPortTemplatesUpdateReader{formats: a.formats},
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
	success, ok := result.(*DcimConsoleServerPortTemplatesUpdateOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for dcim_console-server-port-templates_update: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  DcimConsoleServerPortsBulkDelete dcim console server ports bulk delete API
*/
func (a *Client) DcimConsoleServerPortsBulkDelete(params *DcimConsoleServerPortsBulkDeleteParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*DcimConsoleServerPortsBulkDeleteNoContent, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDcimConsoleServerPortsBulkDeleteParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "dcim_console-server-ports_bulk_delete",
		Method:             "DELETE",
		PathPattern:        "/dcim/console-server-ports/",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &DcimConsoleServerPortsBulkDeleteReader{formats: a.formats},
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
	success, ok := result.(*DcimConsoleServerPortsBulkDeleteNoContent)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for dcim_console-server-ports_bulk_delete: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  DcimConsoleServerPortsBulkPartialUpdate dcim console server ports bulk partial update API
*/
func (a *Client) DcimConsoleServerPortsBulkPartialUpdate(params *DcimConsoleServerPortsBulkPartialUpdateParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*DcimConsoleServerPortsBulkPartialUpdateOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDcimConsoleServerPortsBulkPartialUpdateParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "dcim_console-server-ports_bulk_partial_update",
		Method:             "PATCH",
		PathPattern:        "/dcim/console-server-ports/",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &DcimConsoleServerPortsBulkPartialUpdateReader{formats: a.formats},
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
	success, ok := result.(*DcimConsoleServerPortsBulkPartialUpdateOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for dcim_console-server-ports_bulk_partial_update: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  DcimConsoleServerPortsBulkUpdate dcim console server ports bulk update API
*/
func (a *Client) DcimConsoleServerPortsBulkUpdate(params *DcimConsoleServerPortsBulkUpdateParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*DcimConsoleServerPortsBulkUpdateOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDcimConsoleServerPortsBulkUpdateParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "dcim_console-server-ports_bulk_update",
		Method:             "PUT",
		PathPattern:        "/dcim/console-server-ports/",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &DcimConsoleServerPortsBulkUpdateReader{formats: a.formats},
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
	success, ok := result.(*DcimConsoleServerPortsBulkUpdateOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for dcim_console-server-ports_bulk_update: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  DcimConsoleServerPortsCreate dcim console server ports create API
*/
func (a *Client) DcimConsoleServerPortsCreate(params *DcimConsoleServerPortsCreateParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*DcimConsoleServerPortsCreateCreated, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDcimConsoleServerPortsCreateParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "dcim_console-server-ports_create",
		Method:             "POST",
		PathPattern:        "/dcim/console-server-ports/",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &DcimConsoleServerPortsCreateReader{formats: a.formats},
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
	success, ok := result.(*DcimConsoleServerPortsCreateCreated)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for dcim_console-server-ports_create: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  DcimConsoleServerPortsDelete dcim console server ports delete API
*/
func (a *Client) DcimConsoleServerPortsDelete(params *DcimConsoleServerPortsDeleteParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*DcimConsoleServerPortsDeleteNoContent, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDcimConsoleServerPortsDeleteParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "dcim_console-server-ports_delete",
		Method:             "DELETE",
		PathPattern:        "/dcim/console-server-ports/{id}/",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &DcimConsoleServerPortsDeleteReader{formats: a.formats},
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
	success, ok := result.(*DcimConsoleServerPortsDeleteNoContent)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for dcim_console-server-ports_delete: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  DcimConsoleServerPortsList dcim console server ports list API
*/
func (a *Client) DcimConsoleServerPortsList(params *DcimConsoleServerPortsListParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*DcimConsoleServerPortsListOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDcimConsoleServerPortsListParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "dcim_console-server-ports_list",
		Method:             "GET",
		PathPattern:        "/dcim/console-server-ports/",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &DcimConsoleServerPortsListReader{formats: a.formats},
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
	success, ok := result.(*DcimConsoleServerPortsListOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for dcim_console-server-ports_list: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  DcimConsoleServerPortsPartialUpdate dcim console server ports partial update API
*/
func (a *Client) DcimConsoleServerPortsPartialUpdate(params *DcimConsoleServerPortsPartialUpdateParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*DcimConsoleServerPortsPartialUpdateOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDcimConsoleServerPortsPartialUpdateParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "dcim_console-server-ports_partial_update",
		Method:             "PATCH",
		PathPattern:        "/dcim/console-server-ports/{id}/",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &DcimConsoleServerPortsPartialUpdateReader{formats: a.formats},
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
	success, ok := result.(*DcimConsoleServerPortsPartialUpdateOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for dcim_console-server-ports_partial_update: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  DcimConsoleServerPortsRead dcim console server ports read API
*/
func (a *Client) DcimConsoleServerPortsRead(params *DcimConsoleServerPortsReadParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*DcimConsoleServerPortsReadOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDcimConsoleServerPortsReadParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "dcim_console-server-ports_read",
		Method:             "GET",
		PathPattern:        "/dcim/console-server-ports/{id}/",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &DcimConsoleServerPortsReadReader{formats: a.formats},
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
	success, ok := result.(*DcimConsoleServerPortsReadOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for dcim_console-server-ports_read: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  DcimConsoleServerPortsTrace Trace a complete cable path and return each segment as a three-tuple of (termination, cable, termination).
*/
func (a *Client) DcimConsoleServerPortsTrace(params *DcimConsoleServerPortsTraceParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*DcimConsoleServerPortsTraceOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDcimConsoleServerPortsTraceParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "dcim_console-server-ports_trace",
		Method:             "GET",
		PathPattern:        "/dcim/console-server-ports/{id}/trace/",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &DcimConsoleServerPortsTraceReader{formats: a.formats},
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
	success, ok := result.(*DcimConsoleServerPortsTraceOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for dcim_console-server-ports_trace: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  DcimConsoleServerPortsUpdate dcim console server ports update API
*/
func (a *Client) DcimConsoleServerPortsUpdate(params *DcimConsoleServerPortsUpdateParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*DcimConsoleServerPortsUpdateOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDcimConsoleServerPortsUpdateParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "dcim_console-server-ports_update",
		Method:             "PUT",
		PathPattern:        "/dcim/console-server-ports/{id}/",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &DcimConsoleServerPortsUpdateReader{formats: a.formats},
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
	success, ok := result.(*DcimConsoleServerPortsUpdateOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for dcim_console-server-ports_update: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  DcimDeviceBayTemplatesBulkDelete dcim device bay templates bulk delete API
*/
func (a *Client) DcimDeviceBayTemplatesBulkDelete(params *DcimDeviceBayTemplatesBulkDeleteParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*DcimDeviceBayTemplatesBulkDeleteNoContent, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDcimDeviceBayTemplatesBulkDeleteParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "dcim_device-bay-templates_bulk_delete",
		Method:             "DELETE",
		PathPattern:        "/dcim/device-bay-templates/",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &DcimDeviceBayTemplatesBulkDeleteReader{formats: a.formats},
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
	success, ok := result.(*DcimDeviceBayTemplatesBulkDeleteNoContent)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for dcim_device-bay-templates_bulk_delete: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  DcimDeviceBayTemplatesBulkPartialUpdate dcim device bay templates bulk partial update API
*/
func (a *Client) DcimDeviceBayTemplatesBulkPartialUpdate(params *DcimDeviceBayTemplatesBulkPartialUpdateParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*DcimDeviceBayTemplatesBulkPartialUpdateOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDcimDeviceBayTemplatesBulkPartialUpdateParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "dcim_device-bay-templates_bulk_partial_update",
		Method:             "PATCH",
		PathPattern:        "/dcim/device-bay-templates/",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &DcimDeviceBayTemplatesBulkPartialUpdateReader{formats: a.formats},
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
	success, ok := result.(*DcimDeviceBayTemplatesBulkPartialUpdateOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for dcim_device-bay-templates_bulk_partial_update: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  DcimDeviceBayTemplatesBulkUpdate dcim device bay templates bulk update API
*/
func (a *Client) DcimDeviceBayTemplatesBulkUpdate(params *DcimDeviceBayTemplatesBulkUpdateParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*DcimDeviceBayTemplatesBulkUpdateOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDcimDeviceBayTemplatesBulkUpdateParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "dcim_device-bay-templates_bulk_update",
		Method:             "PUT",
		PathPattern:        "/dcim/device-bay-templates/",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &DcimDeviceBayTemplatesBulkUpdateReader{formats: a.formats},
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
	success, ok := result.(*DcimDeviceBayTemplatesBulkUpdateOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for dcim_device-bay-templates_bulk_update: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  DcimDeviceBayTemplatesCreate dcim device bay templates create API
*/
func (a *Client) DcimDeviceBayTemplatesCreate(params *DcimDeviceBayTemplatesCreateParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*DcimDeviceBayTemplatesCreateCreated, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDcimDeviceBayTemplatesCreateParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "dcim_device-bay-templates_create",
		Method:             "POST",
		PathPattern:        "/dcim/device-bay-templates/",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &DcimDeviceBayTemplatesCreateReader{formats: a.formats},
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
	success, ok := result.(*DcimDeviceBayTemplatesCreateCreated)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for dcim_device-bay-templates_create: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  DcimDeviceBayTemplatesDelete dcim device bay templates delete API
*/
func (a *Client) DcimDeviceBayTemplatesDelete(params *DcimDeviceBayTemplatesDeleteParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*DcimDeviceBayTemplatesDeleteNoContent, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDcimDeviceBayTemplatesDeleteParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "dcim_device-bay-templates_delete",
		Method:             "DELETE",
		PathPattern:        "/dcim/device-bay-templates/{id}/",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &DcimDeviceBayTemplatesDeleteReader{formats: a.formats},
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
	success, ok := result.(*DcimDeviceBayTemplatesDeleteNoContent)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for dcim_device-bay-templates_delete: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  DcimDeviceBayTemplatesList dcim device bay templates list API
*/
func (a *Client) DcimDeviceBayTemplatesList(params *DcimDeviceBayTemplatesListParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*DcimDeviceBayTemplatesListOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDcimDeviceBayTemplatesListParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "dcim_device-bay-templates_list",
		Method:             "GET",
		PathPattern:        "/dcim/device-bay-templates/",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &DcimDeviceBayTemplatesListReader{formats: a.formats},
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
	success, ok := result.(*DcimDeviceBayTemplatesListOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for dcim_device-bay-templates_list: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  DcimDeviceBayTemplatesPartialUpdate dcim device bay templates partial update API
*/
func (a *Client) DcimDeviceBayTemplatesPartialUpdate(params *DcimDeviceBayTemplatesPartialUpdateParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*DcimDeviceBayTemplatesPartialUpdateOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDcimDeviceBayTemplatesPartialUpdateParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "dcim_device-bay-templates_partial_update",
		Method:             "PATCH",
		PathPattern:        "/dcim/device-bay-templates/{id}/",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &DcimDeviceBayTemplatesPartialUpdateReader{formats: a.formats},
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
	success, ok := result.(*DcimDeviceBayTemplatesPartialUpdateOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for dcim_device-bay-templates_partial_update: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  DcimDeviceBayTemplatesRead dcim device bay templates read API
*/
func (a *Client) DcimDeviceBayTemplatesRead(params *DcimDeviceBayTemplatesReadParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*DcimDeviceBayTemplatesReadOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDcimDeviceBayTemplatesReadParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "dcim_device-bay-templates_read",
		Method:             "GET",
		PathPattern:        "/dcim/device-bay-templates/{id}/",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &DcimDeviceBayTemplatesReadReader{formats: a.formats},
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
	success, ok := result.(*DcimDeviceBayTemplatesReadOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for dcim_device-bay-templates_read: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  DcimDeviceBayTemplatesUpdate dcim device bay templates update API
*/
func (a *Client) DcimDeviceBayTemplatesUpdate(params *DcimDeviceBayTemplatesUpdateParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*DcimDeviceBayTemplatesUpdateOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDcimDeviceBayTemplatesUpdateParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "dcim_device-bay-templates_update",
		Method:             "PUT",
		PathPattern:        "/dcim/device-bay-templates/{id}/",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &DcimDeviceBayTemplatesUpdateReader{formats: a.formats},
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
	success, ok := result.(*DcimDeviceBayTemplatesUpdateOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for dcim_device-bay-templates_update: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  DcimDeviceBaysBulkDelete dcim device bays bulk delete API
*/
func (a *Client) DcimDeviceBaysBulkDelete(params *DcimDeviceBaysBulkDeleteParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*DcimDeviceBaysBulkDeleteNoContent, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDcimDeviceBaysBulkDeleteParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "dcim_device-bays_bulk_delete",
		Method:             "DELETE",
		PathPattern:        "/dcim/device-bays/",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &DcimDeviceBaysBulkDeleteReader{formats: a.formats},
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
	success, ok := result.(*DcimDeviceBaysBulkDeleteNoContent)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for dcim_device-bays_bulk_delete: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  DcimDeviceBaysBulkPartialUpdate dcim device bays bulk partial update API
*/
func (a *Client) DcimDeviceBaysBulkPartialUpdate(params *DcimDeviceBaysBulkPartialUpdateParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*DcimDeviceBaysBulkPartialUpdateOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDcimDeviceBaysBulkPartialUpdateParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "dcim_device-bays_bulk_partial_update",
		Method:             "PATCH",
		PathPattern:        "/dcim/device-bays/",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &DcimDeviceBaysBulkPartialUpdateReader{formats: a.formats},
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
	success, ok := result.(*DcimDeviceBaysBulkPartialUpdateOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for dcim_device-bays_bulk_partial_update: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  DcimDeviceBaysBulkUpdate dcim device bays bulk update API
*/
func (a *Client) DcimDeviceBaysBulkUpdate(params *DcimDeviceBaysBulkUpdateParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*DcimDeviceBaysBulkUpdateOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDcimDeviceBaysBulkUpdateParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "dcim_device-bays_bulk_update",
		Method:             "PUT",
		PathPattern:        "/dcim/device-bays/",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &DcimDeviceBaysBulkUpdateReader{formats: a.formats},
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
	success, ok := result.(*DcimDeviceBaysBulkUpdateOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for dcim_device-bays_bulk_update: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  DcimDeviceBaysCreate dcim device bays create API
*/
func (a *Client) DcimDeviceBaysCreate(params *DcimDeviceBaysCreateParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*DcimDeviceBaysCreateCreated, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDcimDeviceBaysCreateParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "dcim_device-bays_create",
		Method:             "POST",
		PathPattern:        "/dcim/device-bays/",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &DcimDeviceBaysCreateReader{formats: a.formats},
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
	success, ok := result.(*DcimDeviceBaysCreateCreated)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for dcim_device-bays_create: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  DcimDeviceBaysDelete dcim device bays delete API
*/
func (a *Client) DcimDeviceBaysDelete(params *DcimDeviceBaysDeleteParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*DcimDeviceBaysDeleteNoContent, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDcimDeviceBaysDeleteParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "dcim_device-bays_delete",
		Method:             "DELETE",
		PathPattern:        "/dcim/device-bays/{id}/",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &DcimDeviceBaysDeleteReader{formats: a.formats},
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
	success, ok := result.(*DcimDeviceBaysDeleteNoContent)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for dcim_device-bays_delete: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  DcimDeviceBaysList dcim device bays list API
*/
func (a *Client) DcimDeviceBaysList(params *DcimDeviceBaysListParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*DcimDeviceBaysListOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDcimDeviceBaysListParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "dcim_device-bays_list",
		Method:             "GET",
		PathPattern:        "/dcim/device-bays/",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &DcimDeviceBaysListReader{formats: a.formats},
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
	success, ok := result.(*DcimDeviceBaysListOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for dcim_device-bays_list: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  DcimDeviceBaysPartialUpdate dcim device bays partial update API
*/
func (a *Client) DcimDeviceBaysPartialUpdate(params *DcimDeviceBaysPartialUpdateParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*DcimDeviceBaysPartialUpdateOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDcimDeviceBaysPartialUpdateParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "dcim_device-bays_partial_update",
		Method:             "PATCH",
		PathPattern:        "/dcim/device-bays/{id}/",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &DcimDeviceBaysPartialUpdateReader{formats: a.formats},
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
	success, ok := result.(*DcimDeviceBaysPartialUpdateOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for dcim_device-bays_partial_update: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  DcimDeviceBaysRead dcim device bays read API
*/
func (a *Client) DcimDeviceBaysRead(params *DcimDeviceBaysReadParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*DcimDeviceBaysReadOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDcimDeviceBaysReadParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "dcim_device-bays_read",
		Method:             "GET",
		PathPattern:        "/dcim/device-bays/{id}/",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &DcimDeviceBaysReadReader{formats: a.formats},
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
	success, ok := result.(*DcimDeviceBaysReadOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for dcim_device-bays_read: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  DcimDeviceBaysUpdate dcim device bays update API
*/
func (a *Client) DcimDeviceBaysUpdate(params *DcimDeviceBaysUpdateParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*DcimDeviceBaysUpdateOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDcimDeviceBaysUpdateParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "dcim_device-bays_update",
		Method:             "PUT",
		PathPattern:        "/dcim/device-bays/{id}/",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &DcimDeviceBaysUpdateReader{formats: a.formats},
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
	success, ok := result.(*DcimDeviceBaysUpdateOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for dcim_device-bays_update: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  DcimDeviceRolesBulkDelete dcim device roles bulk delete API
*/
func (a *Client) DcimDeviceRolesBulkDelete(params *DcimDeviceRolesBulkDeleteParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*DcimDeviceRolesBulkDeleteNoContent, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDcimDeviceRolesBulkDeleteParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "dcim_device-roles_bulk_delete",
		Method:             "DELETE",
		PathPattern:        "/dcim/device-roles/",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &DcimDeviceRolesBulkDeleteReader{formats: a.formats},
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
	success, ok := result.(*DcimDeviceRolesBulkDeleteNoContent)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for dcim_device-roles_bulk_delete: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  DcimDeviceRolesBulkPartialUpdate dcim device roles bulk partial update API
*/
func (a *Client) DcimDeviceRolesBulkPartialUpdate(params *DcimDeviceRolesBulkPartialUpdateParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*DcimDeviceRolesBulkPartialUpdateOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDcimDeviceRolesBulkPartialUpdateParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "dcim_device-roles_bulk_partial_update",
		Method:             "PATCH",
		PathPattern:        "/dcim/device-roles/",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &DcimDeviceRolesBulkPartialUpdateReader{formats: a.formats},
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
	success, ok := result.(*DcimDeviceRolesBulkPartialUpdateOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for dcim_device-roles_bulk_partial_update: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  DcimDeviceRolesBulkUpdate dcim device roles bulk update API
*/
func (a *Client) DcimDeviceRolesBulkUpdate(params *DcimDeviceRolesBulkUpdateParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*DcimDeviceRolesBulkUpdateOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDcimDeviceRolesBulkUpdateParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "dcim_device-roles_bulk_update",
		Method:             "PUT",
		PathPattern:        "/dcim/device-roles/",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &DcimDeviceRolesBulkUpdateReader{formats: a.formats},
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
	success, ok := result.(*DcimDeviceRolesBulkUpdateOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for dcim_device-roles_bulk_update: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  DcimDeviceRolesCreate dcim device roles create API
*/
func (a *Client) DcimDeviceRolesCreate(params *DcimDeviceRolesCreateParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*DcimDeviceRolesCreateCreated, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDcimDeviceRolesCreateParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "dcim_device-roles_create",
		Method:             "POST",
		PathPattern:        "/dcim/device-roles/",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &DcimDeviceRolesCreateReader{formats: a.formats},
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
	success, ok := result.(*DcimDeviceRolesCreateCreated)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for dcim_device-roles_create: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  DcimDeviceRolesDelete dcim device roles delete API
*/
func (a *Client) DcimDeviceRolesDelete(params *DcimDeviceRolesDeleteParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*DcimDeviceRolesDeleteNoContent, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDcimDeviceRolesDeleteParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "dcim_device-roles_delete",
		Method:             "DELETE",
		PathPattern:        "/dcim/device-roles/{id}/",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &DcimDeviceRolesDeleteReader{formats: a.formats},
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
	success, ok := result.(*DcimDeviceRolesDeleteNoContent)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for dcim_device-roles_delete: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  DcimDeviceRolesList dcim device roles list API
*/
func (a *Client) DcimDeviceRolesList(params *DcimDeviceRolesListParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*DcimDeviceRolesListOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDcimDeviceRolesListParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "dcim_device-roles_list",
		Method:             "GET",
		PathPattern:        "/dcim/device-roles/",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &DcimDeviceRolesListReader{formats: a.formats},
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
	success, ok := result.(*DcimDeviceRolesListOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for dcim_device-roles_list: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  DcimDeviceRolesPartialUpdate dcim device roles partial update API
*/
func (a *Client) DcimDeviceRolesPartialUpdate(params *DcimDeviceRolesPartialUpdateParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*DcimDeviceRolesPartialUpdateOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDcimDeviceRolesPartialUpdateParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "dcim_device-roles_partial_update",
		Method:             "PATCH",
		PathPattern:        "/dcim/device-roles/{id}/",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &DcimDeviceRolesPartialUpdateReader{formats: a.formats},
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
	success, ok := result.(*DcimDeviceRolesPartialUpdateOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for dcim_device-roles_partial_update: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  DcimDeviceRolesRead dcim device roles read API
*/
func (a *Client) DcimDeviceRolesRead(params *DcimDeviceRolesReadParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*DcimDeviceRolesReadOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDcimDeviceRolesReadParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "dcim_device-roles_read",
		Method:             "GET",
		PathPattern:        "/dcim/device-roles/{id}/",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &DcimDeviceRolesReadReader{formats: a.formats},
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
	success, ok := result.(*DcimDeviceRolesReadOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for dcim_device-roles_read: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  DcimDeviceRolesUpdate dcim device roles update API
*/
func (a *Client) DcimDeviceRolesUpdate(params *DcimDeviceRolesUpdateParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*DcimDeviceRolesUpdateOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDcimDeviceRolesUpdateParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "dcim_device-roles_update",
		Method:             "PUT",
		PathPattern:        "/dcim/device-roles/{id}/",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &DcimDeviceRolesUpdateReader{formats: a.formats},
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
	success, ok := result.(*DcimDeviceRolesUpdateOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for dcim_device-roles_update: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  DcimDeviceTypesBulkDelete dcim device types bulk delete API
*/
func (a *Client) DcimDeviceTypesBulkDelete(params *DcimDeviceTypesBulkDeleteParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*DcimDeviceTypesBulkDeleteNoContent, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDcimDeviceTypesBulkDeleteParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "dcim_device-types_bulk_delete",
		Method:             "DELETE",
		PathPattern:        "/dcim/device-types/",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &DcimDeviceTypesBulkDeleteReader{formats: a.formats},
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
	success, ok := result.(*DcimDeviceTypesBulkDeleteNoContent)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for dcim_device-types_bulk_delete: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  DcimDeviceTypesBulkPartialUpdate dcim device types bulk partial update API
*/
func (a *Client) DcimDeviceTypesBulkPartialUpdate(params *DcimDeviceTypesBulkPartialUpdateParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*DcimDeviceTypesBulkPartialUpdateOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDcimDeviceTypesBulkPartialUpdateParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "dcim_device-types_bulk_partial_update",
		Method:             "PATCH",
		PathPattern:        "/dcim/device-types/",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &DcimDeviceTypesBulkPartialUpdateReader{formats: a.formats},
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
	success, ok := result.(*DcimDeviceTypesBulkPartialUpdateOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for dcim_device-types_bulk_partial_update: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  DcimDeviceTypesBulkUpdate dcim device types bulk update API
*/
func (a *Client) DcimDeviceTypesBulkUpdate(params *DcimDeviceTypesBulkUpdateParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*DcimDeviceTypesBulkUpdateOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDcimDeviceTypesBulkUpdateParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "dcim_device-types_bulk_update",
		Method:             "PUT",
		PathPattern:        "/dcim/device-types/",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &DcimDeviceTypesBulkUpdateReader{formats: a.formats},
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
	success, ok := result.(*DcimDeviceTypesBulkUpdateOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for dcim_device-types_bulk_update: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  DcimDeviceTypesCreate dcim device types create API
*/
func (a *Client) DcimDeviceTypesCreate(params *DcimDeviceTypesCreateParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*DcimDeviceTypesCreateCreated, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDcimDeviceTypesCreateParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "dcim_device-types_create",
		Method:             "POST",
		PathPattern:        "/dcim/device-types/",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &DcimDeviceTypesCreateReader{formats: a.formats},
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
	success, ok := result.(*DcimDeviceTypesCreateCreated)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for dcim_device-types_create: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  DcimDeviceTypesDelete dcim device types delete API
*/
func (a *Client) DcimDeviceTypesDelete(params *DcimDeviceTypesDeleteParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*DcimDeviceTypesDeleteNoContent, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDcimDeviceTypesDeleteParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "dcim_device-types_delete",
		Method:             "DELETE",
		PathPattern:        "/dcim/device-types/{id}/",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &DcimDeviceTypesDeleteReader{formats: a.formats},
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
	success, ok := result.(*DcimDeviceTypesDeleteNoContent)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for dcim_device-types_delete: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  DcimDeviceTypesList dcim device types list API
*/
func (a *Client) DcimDeviceTypesList(params *DcimDeviceTypesListParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*DcimDeviceTypesListOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDcimDeviceTypesListParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "dcim_device-types_list",
		Method:             "GET",
		PathPattern:        "/dcim/device-types/",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &DcimDeviceTypesListReader{formats: a.formats},
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
	success, ok := result.(*DcimDeviceTypesListOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for dcim_device-types_list: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  DcimDeviceTypesPartialUpdate dcim device types partial update API
*/
func (a *Client) DcimDeviceTypesPartialUpdate(params *DcimDeviceTypesPartialUpdateParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*DcimDeviceTypesPartialUpdateOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDcimDeviceTypesPartialUpdateParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "dcim_device-types_partial_update",
		Method:             "PATCH",
		PathPattern:        "/dcim/device-types/{id}/",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &DcimDeviceTypesPartialUpdateReader{formats: a.formats},
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
	success, ok := result.(*DcimDeviceTypesPartialUpdateOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for dcim_device-types_partial_update: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  DcimDeviceTypesRead dcim device types read API
*/
func (a *Client) DcimDeviceTypesRead(params *DcimDeviceTypesReadParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*DcimDeviceTypesReadOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDcimDeviceTypesReadParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "dcim_device-types_read",
		Method:             "GET",
		PathPattern:        "/dcim/device-types/{id}/",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &DcimDeviceTypesReadReader{formats: a.formats},
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
	success, ok := result.(*DcimDeviceTypesReadOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for dcim_device-types_read: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  DcimDeviceTypesUpdate dcim device types update API
*/
func (a *Client) DcimDeviceTypesUpdate(params *DcimDeviceTypesUpdateParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*DcimDeviceTypesUpdateOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDcimDeviceTypesUpdateParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "dcim_device-types_update",
		Method:             "PUT",
		PathPattern:        "/dcim/device-types/{id}/",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &DcimDeviceTypesUpdateReader{formats: a.formats},
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
	success, ok := result.(*DcimDeviceTypesUpdateOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for dcim_device-types_update: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  DcimDevicesBulkDelete dcim devices bulk delete API
*/
func (a *Client) DcimDevicesBulkDelete(params *DcimDevicesBulkDeleteParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*DcimDevicesBulkDeleteNoContent, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDcimDevicesBulkDeleteParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "dcim_devices_bulk_delete",
		Method:             "DELETE",
		PathPattern:        "/dcim/devices/",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &DcimDevicesBulkDeleteReader{formats: a.formats},
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
	success, ok := result.(*DcimDevicesBulkDeleteNoContent)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for dcim_devices_bulk_delete: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  DcimDevicesBulkPartialUpdate dcim devices bulk partial update API
*/
func (a *Client) DcimDevicesBulkPartialUpdate(params *DcimDevicesBulkPartialUpdateParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*DcimDevicesBulkPartialUpdateOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDcimDevicesBulkPartialUpdateParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "dcim_devices_bulk_partial_update",
		Method:             "PATCH",
		PathPattern:        "/dcim/devices/",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &DcimDevicesBulkPartialUpdateReader{formats: a.formats},
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
	success, ok := result.(*DcimDevicesBulkPartialUpdateOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for dcim_devices_bulk_partial_update: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  DcimDevicesBulkUpdate dcim devices bulk update API
*/
func (a *Client) DcimDevicesBulkUpdate(params *DcimDevicesBulkUpdateParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*DcimDevicesBulkUpdateOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDcimDevicesBulkUpdateParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "dcim_devices_bulk_update",
		Method:             "PUT",
		PathPattern:        "/dcim/devices/",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &DcimDevicesBulkUpdateReader{formats: a.formats},
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
	success, ok := result.(*DcimDevicesBulkUpdateOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for dcim_devices_bulk_update: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  DcimDevicesCreate dcim devices create API
*/
func (a *Client) DcimDevicesCreate(params *DcimDevicesCreateParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*DcimDevicesCreateCreated, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDcimDevicesCreateParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "dcim_devices_create",
		Method:             "POST",
		PathPattern:        "/dcim/devices/",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &DcimDevicesCreateReader{formats: a.formats},
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
	success, ok := result.(*DcimDevicesCreateCreated)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for dcim_devices_create: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  DcimDevicesDelete dcim devices delete API
*/
func (a *Client) DcimDevicesDelete(params *DcimDevicesDeleteParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*DcimDevicesDeleteNoContent, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDcimDevicesDeleteParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "dcim_devices_delete",
		Method:             "DELETE",
		PathPattern:        "/dcim/devices/{id}/",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &DcimDevicesDeleteReader{formats: a.formats},
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
	success, ok := result.(*DcimDevicesDeleteNoContent)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for dcim_devices_delete: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  DcimDevicesList dcim devices list API
*/
func (a *Client) DcimDevicesList(params *DcimDevicesListParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*DcimDevicesListOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDcimDevicesListParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "dcim_devices_list",
		Method:             "GET",
		PathPattern:        "/dcim/devices/",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &DcimDevicesListReader{formats: a.formats},
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
	success, ok := result.(*DcimDevicesListOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for dcim_devices_list: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  DcimDevicesNapalm Execute a NAPALM method on a Device
*/
func (a *Client) DcimDevicesNapalm(params *DcimDevicesNapalmParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*DcimDevicesNapalmOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDcimDevicesNapalmParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "dcim_devices_napalm",
		Method:             "GET",
		PathPattern:        "/dcim/devices/{id}/napalm/",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &DcimDevicesNapalmReader{formats: a.formats},
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
	success, ok := result.(*DcimDevicesNapalmOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for dcim_devices_napalm: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  DcimDevicesPartialUpdate dcim devices partial update API
*/
func (a *Client) DcimDevicesPartialUpdate(params *DcimDevicesPartialUpdateParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*DcimDevicesPartialUpdateOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDcimDevicesPartialUpdateParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "dcim_devices_partial_update",
		Method:             "PATCH",
		PathPattern:        "/dcim/devices/{id}/",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &DcimDevicesPartialUpdateReader{formats: a.formats},
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
	success, ok := result.(*DcimDevicesPartialUpdateOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for dcim_devices_partial_update: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  DcimDevicesRead dcim devices read API
*/
func (a *Client) DcimDevicesRead(params *DcimDevicesReadParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*DcimDevicesReadOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDcimDevicesReadParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "dcim_devices_read",
		Method:             "GET",
		PathPattern:        "/dcim/devices/{id}/",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &DcimDevicesReadReader{formats: a.formats},
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
	success, ok := result.(*DcimDevicesReadOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for dcim_devices_read: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  DcimDevicesUpdate dcim devices update API
*/
func (a *Client) DcimDevicesUpdate(params *DcimDevicesUpdateParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*DcimDevicesUpdateOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDcimDevicesUpdateParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "dcim_devices_update",
		Method:             "PUT",
		PathPattern:        "/dcim/devices/{id}/",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &DcimDevicesUpdateReader{formats: a.formats},
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
	success, ok := result.(*DcimDevicesUpdateOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for dcim_devices_update: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  DcimFrontPortTemplatesBulkDelete dcim front port templates bulk delete API
*/
func (a *Client) DcimFrontPortTemplatesBulkDelete(params *DcimFrontPortTemplatesBulkDeleteParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*DcimFrontPortTemplatesBulkDeleteNoContent, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDcimFrontPortTemplatesBulkDeleteParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "dcim_front-port-templates_bulk_delete",
		Method:             "DELETE",
		PathPattern:        "/dcim/front-port-templates/",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &DcimFrontPortTemplatesBulkDeleteReader{formats: a.formats},
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
	success, ok := result.(*DcimFrontPortTemplatesBulkDeleteNoContent)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for dcim_front-port-templates_bulk_delete: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  DcimFrontPortTemplatesBulkPartialUpdate dcim front port templates bulk partial update API
*/
func (a *Client) DcimFrontPortTemplatesBulkPartialUpdate(params *DcimFrontPortTemplatesBulkPartialUpdateParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*DcimFrontPortTemplatesBulkPartialUpdateOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDcimFrontPortTemplatesBulkPartialUpdateParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "dcim_front-port-templates_bulk_partial_update",
		Method:             "PATCH",
		PathPattern:        "/dcim/front-port-templates/",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &DcimFrontPortTemplatesBulkPartialUpdateReader{formats: a.formats},
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
	success, ok := result.(*DcimFrontPortTemplatesBulkPartialUpdateOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for dcim_front-port-templates_bulk_partial_update: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  DcimFrontPortTemplatesBulkUpdate dcim front port templates bulk update API
*/
func (a *Client) DcimFrontPortTemplatesBulkUpdate(params *DcimFrontPortTemplatesBulkUpdateParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*DcimFrontPortTemplatesBulkUpdateOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDcimFrontPortTemplatesBulkUpdateParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "dcim_front-port-templates_bulk_update",
		Method:             "PUT",
		PathPattern:        "/dcim/front-port-templates/",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &DcimFrontPortTemplatesBulkUpdateReader{formats: a.formats},
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
	success, ok := result.(*DcimFrontPortTemplatesBulkUpdateOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for dcim_front-port-templates_bulk_update: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  DcimFrontPortTemplatesCreate dcim front port templates create API
*/
func (a *Client) DcimFrontPortTemplatesCreate(params *DcimFrontPortTemplatesCreateParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*DcimFrontPortTemplatesCreateCreated, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDcimFrontPortTemplatesCreateParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "dcim_front-port-templates_create",
		Method:             "POST",
		PathPattern:        "/dcim/front-port-templates/",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &DcimFrontPortTemplatesCreateReader{formats: a.formats},
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
	success, ok := result.(*DcimFrontPortTemplatesCreateCreated)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for dcim_front-port-templates_create: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  DcimFrontPortTemplatesDelete dcim front port templates delete API
*/
func (a *Client) DcimFrontPortTemplatesDelete(params *DcimFrontPortTemplatesDeleteParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*DcimFrontPortTemplatesDeleteNoContent, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDcimFrontPortTemplatesDeleteParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "dcim_front-port-templates_delete",
		Method:             "DELETE",
		PathPattern:        "/dcim/front-port-templates/{id}/",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &DcimFrontPortTemplatesDeleteReader{formats: a.formats},
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
	success, ok := result.(*DcimFrontPortTemplatesDeleteNoContent)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for dcim_front-port-templates_delete: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  DcimFrontPortTemplatesList dcim front port templates list API
*/
func (a *Client) DcimFrontPortTemplatesList(params *DcimFrontPortTemplatesListParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*DcimFrontPortTemplatesListOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDcimFrontPortTemplatesListParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "dcim_front-port-templates_list",
		Method:             "GET",
		PathPattern:        "/dcim/front-port-templates/",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &DcimFrontPortTemplatesListReader{formats: a.formats},
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
	success, ok := result.(*DcimFrontPortTemplatesListOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for dcim_front-port-templates_list: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  DcimFrontPortTemplatesPartialUpdate dcim front port templates partial update API
*/
func (a *Client) DcimFrontPortTemplatesPartialUpdate(params *DcimFrontPortTemplatesPartialUpdateParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*DcimFrontPortTemplatesPartialUpdateOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDcimFrontPortTemplatesPartialUpdateParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "dcim_front-port-templates_partial_update",
		Method:             "PATCH",
		PathPattern:        "/dcim/front-port-templates/{id}/",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &DcimFrontPortTemplatesPartialUpdateReader{formats: a.formats},
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
	success, ok := result.(*DcimFrontPortTemplatesPartialUpdateOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for dcim_front-port-templates_partial_update: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  DcimFrontPortTemplatesRead dcim front port templates read API
*/
func (a *Client) DcimFrontPortTemplatesRead(params *DcimFrontPortTemplatesReadParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*DcimFrontPortTemplatesReadOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDcimFrontPortTemplatesReadParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "dcim_front-port-templates_read",
		Method:             "GET",
		PathPattern:        "/dcim/front-port-templates/{id}/",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &DcimFrontPortTemplatesReadReader{formats: a.formats},
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
	success, ok := result.(*DcimFrontPortTemplatesReadOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for dcim_front-port-templates_read: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  DcimFrontPortTemplatesUpdate dcim front port templates update API
*/
func (a *Client) DcimFrontPortTemplatesUpdate(params *DcimFrontPortTemplatesUpdateParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*DcimFrontPortTemplatesUpdateOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDcimFrontPortTemplatesUpdateParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "dcim_front-port-templates_update",
		Method:             "PUT",
		PathPattern:        "/dcim/front-port-templates/{id}/",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &DcimFrontPortTemplatesUpdateReader{formats: a.formats},
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
	success, ok := result.(*DcimFrontPortTemplatesUpdateOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for dcim_front-port-templates_update: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  DcimFrontPortsBulkDelete dcim front ports bulk delete API
*/
func (a *Client) DcimFrontPortsBulkDelete(params *DcimFrontPortsBulkDeleteParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*DcimFrontPortsBulkDeleteNoContent, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDcimFrontPortsBulkDeleteParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "dcim_front-ports_bulk_delete",
		Method:             "DELETE",
		PathPattern:        "/dcim/front-ports/",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &DcimFrontPortsBulkDeleteReader{formats: a.formats},
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
	success, ok := result.(*DcimFrontPortsBulkDeleteNoContent)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for dcim_front-ports_bulk_delete: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  DcimFrontPortsBulkPartialUpdate dcim front ports bulk partial update API
*/
func (a *Client) DcimFrontPortsBulkPartialUpdate(params *DcimFrontPortsBulkPartialUpdateParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*DcimFrontPortsBulkPartialUpdateOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDcimFrontPortsBulkPartialUpdateParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "dcim_front-ports_bulk_partial_update",
		Method:             "PATCH",
		PathPattern:        "/dcim/front-ports/",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &DcimFrontPortsBulkPartialUpdateReader{formats: a.formats},
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
	success, ok := result.(*DcimFrontPortsBulkPartialUpdateOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for dcim_front-ports_bulk_partial_update: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  DcimFrontPortsBulkUpdate dcim front ports bulk update API
*/
func (a *Client) DcimFrontPortsBulkUpdate(params *DcimFrontPortsBulkUpdateParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*DcimFrontPortsBulkUpdateOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDcimFrontPortsBulkUpdateParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "dcim_front-ports_bulk_update",
		Method:             "PUT",
		PathPattern:        "/dcim/front-ports/",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &DcimFrontPortsBulkUpdateReader{formats: a.formats},
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
	success, ok := result.(*DcimFrontPortsBulkUpdateOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for dcim_front-ports_bulk_update: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  DcimFrontPortsCreate dcim front ports create API
*/
func (a *Client) DcimFrontPortsCreate(params *DcimFrontPortsCreateParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*DcimFrontPortsCreateCreated, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDcimFrontPortsCreateParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "dcim_front-ports_create",
		Method:             "POST",
		PathPattern:        "/dcim/front-ports/",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &DcimFrontPortsCreateReader{formats: a.formats},
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
	success, ok := result.(*DcimFrontPortsCreateCreated)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for dcim_front-ports_create: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  DcimFrontPortsDelete dcim front ports delete API
*/
func (a *Client) DcimFrontPortsDelete(params *DcimFrontPortsDeleteParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*DcimFrontPortsDeleteNoContent, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDcimFrontPortsDeleteParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "dcim_front-ports_delete",
		Method:             "DELETE",
		PathPattern:        "/dcim/front-ports/{id}/",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &DcimFrontPortsDeleteReader{formats: a.formats},
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
	success, ok := result.(*DcimFrontPortsDeleteNoContent)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for dcim_front-ports_delete: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  DcimFrontPortsList dcim front ports list API
*/
func (a *Client) DcimFrontPortsList(params *DcimFrontPortsListParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*DcimFrontPortsListOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDcimFrontPortsListParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "dcim_front-ports_list",
		Method:             "GET",
		PathPattern:        "/dcim/front-ports/",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &DcimFrontPortsListReader{formats: a.formats},
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
	success, ok := result.(*DcimFrontPortsListOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for dcim_front-ports_list: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  DcimFrontPortsPartialUpdate dcim front ports partial update API
*/
func (a *Client) DcimFrontPortsPartialUpdate(params *DcimFrontPortsPartialUpdateParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*DcimFrontPortsPartialUpdateOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDcimFrontPortsPartialUpdateParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "dcim_front-ports_partial_update",
		Method:             "PATCH",
		PathPattern:        "/dcim/front-ports/{id}/",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &DcimFrontPortsPartialUpdateReader{formats: a.formats},
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
	success, ok := result.(*DcimFrontPortsPartialUpdateOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for dcim_front-ports_partial_update: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  DcimFrontPortsPaths Return all CablePaths which traverse a given pass-through port.
*/
func (a *Client) DcimFrontPortsPaths(params *DcimFrontPortsPathsParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*DcimFrontPortsPathsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDcimFrontPortsPathsParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "dcim_front-ports_paths",
		Method:             "GET",
		PathPattern:        "/dcim/front-ports/{id}/paths/",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &DcimFrontPortsPathsReader{formats: a.formats},
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
	success, ok := result.(*DcimFrontPortsPathsOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for dcim_front-ports_paths: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  DcimFrontPortsRead dcim front ports read API
*/
func (a *Client) DcimFrontPortsRead(params *DcimFrontPortsReadParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*DcimFrontPortsReadOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDcimFrontPortsReadParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "dcim_front-ports_read",
		Method:             "GET",
		PathPattern:        "/dcim/front-ports/{id}/",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &DcimFrontPortsReadReader{formats: a.formats},
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
	success, ok := result.(*DcimFrontPortsReadOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for dcim_front-ports_read: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  DcimFrontPortsUpdate dcim front ports update API
*/
func (a *Client) DcimFrontPortsUpdate(params *DcimFrontPortsUpdateParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*DcimFrontPortsUpdateOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDcimFrontPortsUpdateParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "dcim_front-ports_update",
		Method:             "PUT",
		PathPattern:        "/dcim/front-ports/{id}/",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &DcimFrontPortsUpdateReader{formats: a.formats},
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
	success, ok := result.(*DcimFrontPortsUpdateOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for dcim_front-ports_update: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  DcimInterfaceConnectionsList dcim interface connections list API
*/
func (a *Client) DcimInterfaceConnectionsList(params *DcimInterfaceConnectionsListParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*DcimInterfaceConnectionsListOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDcimInterfaceConnectionsListParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "dcim_interface-connections_list",
		Method:             "GET",
		PathPattern:        "/dcim/interface-connections/",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &DcimInterfaceConnectionsListReader{formats: a.formats},
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
	success, ok := result.(*DcimInterfaceConnectionsListOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for dcim_interface-connections_list: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  DcimInterfaceTemplatesBulkDelete dcim interface templates bulk delete API
*/
func (a *Client) DcimInterfaceTemplatesBulkDelete(params *DcimInterfaceTemplatesBulkDeleteParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*DcimInterfaceTemplatesBulkDeleteNoContent, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDcimInterfaceTemplatesBulkDeleteParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "dcim_interface-templates_bulk_delete",
		Method:             "DELETE",
		PathPattern:        "/dcim/interface-templates/",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &DcimInterfaceTemplatesBulkDeleteReader{formats: a.formats},
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
	success, ok := result.(*DcimInterfaceTemplatesBulkDeleteNoContent)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for dcim_interface-templates_bulk_delete: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  DcimInterfaceTemplatesBulkPartialUpdate dcim interface templates bulk partial update API
*/
func (a *Client) DcimInterfaceTemplatesBulkPartialUpdate(params *DcimInterfaceTemplatesBulkPartialUpdateParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*DcimInterfaceTemplatesBulkPartialUpdateOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDcimInterfaceTemplatesBulkPartialUpdateParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "dcim_interface-templates_bulk_partial_update",
		Method:             "PATCH",
		PathPattern:        "/dcim/interface-templates/",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &DcimInterfaceTemplatesBulkPartialUpdateReader{formats: a.formats},
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
	success, ok := result.(*DcimInterfaceTemplatesBulkPartialUpdateOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for dcim_interface-templates_bulk_partial_update: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  DcimInterfaceTemplatesBulkUpdate dcim interface templates bulk update API
*/
func (a *Client) DcimInterfaceTemplatesBulkUpdate(params *DcimInterfaceTemplatesBulkUpdateParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*DcimInterfaceTemplatesBulkUpdateOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDcimInterfaceTemplatesBulkUpdateParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "dcim_interface-templates_bulk_update",
		Method:             "PUT",
		PathPattern:        "/dcim/interface-templates/",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &DcimInterfaceTemplatesBulkUpdateReader{formats: a.formats},
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
	success, ok := result.(*DcimInterfaceTemplatesBulkUpdateOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for dcim_interface-templates_bulk_update: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  DcimInterfaceTemplatesCreate dcim interface templates create API
*/
func (a *Client) DcimInterfaceTemplatesCreate(params *DcimInterfaceTemplatesCreateParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*DcimInterfaceTemplatesCreateCreated, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDcimInterfaceTemplatesCreateParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "dcim_interface-templates_create",
		Method:             "POST",
		PathPattern:        "/dcim/interface-templates/",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &DcimInterfaceTemplatesCreateReader{formats: a.formats},
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
	success, ok := result.(*DcimInterfaceTemplatesCreateCreated)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for dcim_interface-templates_create: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  DcimInterfaceTemplatesDelete dcim interface templates delete API
*/
func (a *Client) DcimInterfaceTemplatesDelete(params *DcimInterfaceTemplatesDeleteParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*DcimInterfaceTemplatesDeleteNoContent, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDcimInterfaceTemplatesDeleteParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "dcim_interface-templates_delete",
		Method:             "DELETE",
		PathPattern:        "/dcim/interface-templates/{id}/",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &DcimInterfaceTemplatesDeleteReader{formats: a.formats},
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
	success, ok := result.(*DcimInterfaceTemplatesDeleteNoContent)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for dcim_interface-templates_delete: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  DcimInterfaceTemplatesList dcim interface templates list API
*/
func (a *Client) DcimInterfaceTemplatesList(params *DcimInterfaceTemplatesListParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*DcimInterfaceTemplatesListOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDcimInterfaceTemplatesListParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "dcim_interface-templates_list",
		Method:             "GET",
		PathPattern:        "/dcim/interface-templates/",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &DcimInterfaceTemplatesListReader{formats: a.formats},
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
	success, ok := result.(*DcimInterfaceTemplatesListOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for dcim_interface-templates_list: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  DcimInterfaceTemplatesPartialUpdate dcim interface templates partial update API
*/
func (a *Client) DcimInterfaceTemplatesPartialUpdate(params *DcimInterfaceTemplatesPartialUpdateParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*DcimInterfaceTemplatesPartialUpdateOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDcimInterfaceTemplatesPartialUpdateParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "dcim_interface-templates_partial_update",
		Method:             "PATCH",
		PathPattern:        "/dcim/interface-templates/{id}/",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &DcimInterfaceTemplatesPartialUpdateReader{formats: a.formats},
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
	success, ok := result.(*DcimInterfaceTemplatesPartialUpdateOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for dcim_interface-templates_partial_update: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  DcimInterfaceTemplatesRead dcim interface templates read API
*/
func (a *Client) DcimInterfaceTemplatesRead(params *DcimInterfaceTemplatesReadParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*DcimInterfaceTemplatesReadOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDcimInterfaceTemplatesReadParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "dcim_interface-templates_read",
		Method:             "GET",
		PathPattern:        "/dcim/interface-templates/{id}/",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &DcimInterfaceTemplatesReadReader{formats: a.formats},
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
	success, ok := result.(*DcimInterfaceTemplatesReadOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for dcim_interface-templates_read: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  DcimInterfaceTemplatesUpdate dcim interface templates update API
*/
func (a *Client) DcimInterfaceTemplatesUpdate(params *DcimInterfaceTemplatesUpdateParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*DcimInterfaceTemplatesUpdateOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDcimInterfaceTemplatesUpdateParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "dcim_interface-templates_update",
		Method:             "PUT",
		PathPattern:        "/dcim/interface-templates/{id}/",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &DcimInterfaceTemplatesUpdateReader{formats: a.formats},
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
	success, ok := result.(*DcimInterfaceTemplatesUpdateOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for dcim_interface-templates_update: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  DcimInterfacesBulkDelete dcim interfaces bulk delete API
*/
func (a *Client) DcimInterfacesBulkDelete(params *DcimInterfacesBulkDeleteParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*DcimInterfacesBulkDeleteNoContent, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDcimInterfacesBulkDeleteParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "dcim_interfaces_bulk_delete",
		Method:             "DELETE",
		PathPattern:        "/dcim/interfaces/",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &DcimInterfacesBulkDeleteReader{formats: a.formats},
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
	success, ok := result.(*DcimInterfacesBulkDeleteNoContent)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for dcim_interfaces_bulk_delete: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  DcimInterfacesBulkPartialUpdate dcim interfaces bulk partial update API
*/
func (a *Client) DcimInterfacesBulkPartialUpdate(params *DcimInterfacesBulkPartialUpdateParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*DcimInterfacesBulkPartialUpdateOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDcimInterfacesBulkPartialUpdateParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "dcim_interfaces_bulk_partial_update",
		Method:             "PATCH",
		PathPattern:        "/dcim/interfaces/",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &DcimInterfacesBulkPartialUpdateReader{formats: a.formats},
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
	success, ok := result.(*DcimInterfacesBulkPartialUpdateOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for dcim_interfaces_bulk_partial_update: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  DcimInterfacesBulkUpdate dcim interfaces bulk update API
*/
func (a *Client) DcimInterfacesBulkUpdate(params *DcimInterfacesBulkUpdateParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*DcimInterfacesBulkUpdateOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDcimInterfacesBulkUpdateParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "dcim_interfaces_bulk_update",
		Method:             "PUT",
		PathPattern:        "/dcim/interfaces/",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &DcimInterfacesBulkUpdateReader{formats: a.formats},
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
	success, ok := result.(*DcimInterfacesBulkUpdateOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for dcim_interfaces_bulk_update: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  DcimInterfacesCreate dcim interfaces create API
*/
func (a *Client) DcimInterfacesCreate(params *DcimInterfacesCreateParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*DcimInterfacesCreateCreated, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDcimInterfacesCreateParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "dcim_interfaces_create",
		Method:             "POST",
		PathPattern:        "/dcim/interfaces/",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &DcimInterfacesCreateReader{formats: a.formats},
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
	success, ok := result.(*DcimInterfacesCreateCreated)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for dcim_interfaces_create: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  DcimInterfacesDelete dcim interfaces delete API
*/
func (a *Client) DcimInterfacesDelete(params *DcimInterfacesDeleteParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*DcimInterfacesDeleteNoContent, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDcimInterfacesDeleteParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "dcim_interfaces_delete",
		Method:             "DELETE",
		PathPattern:        "/dcim/interfaces/{id}/",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &DcimInterfacesDeleteReader{formats: a.formats},
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
	success, ok := result.(*DcimInterfacesDeleteNoContent)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for dcim_interfaces_delete: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  DcimInterfacesList dcim interfaces list API
*/
func (a *Client) DcimInterfacesList(params *DcimInterfacesListParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*DcimInterfacesListOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDcimInterfacesListParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "dcim_interfaces_list",
		Method:             "GET",
		PathPattern:        "/dcim/interfaces/",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &DcimInterfacesListReader{formats: a.formats},
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
	success, ok := result.(*DcimInterfacesListOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for dcim_interfaces_list: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  DcimInterfacesPartialUpdate dcim interfaces partial update API
*/
func (a *Client) DcimInterfacesPartialUpdate(params *DcimInterfacesPartialUpdateParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*DcimInterfacesPartialUpdateOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDcimInterfacesPartialUpdateParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "dcim_interfaces_partial_update",
		Method:             "PATCH",
		PathPattern:        "/dcim/interfaces/{id}/",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &DcimInterfacesPartialUpdateReader{formats: a.formats},
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
	success, ok := result.(*DcimInterfacesPartialUpdateOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for dcim_interfaces_partial_update: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  DcimInterfacesRead dcim interfaces read API
*/
func (a *Client) DcimInterfacesRead(params *DcimInterfacesReadParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*DcimInterfacesReadOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDcimInterfacesReadParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "dcim_interfaces_read",
		Method:             "GET",
		PathPattern:        "/dcim/interfaces/{id}/",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &DcimInterfacesReadReader{formats: a.formats},
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
	success, ok := result.(*DcimInterfacesReadOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for dcim_interfaces_read: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  DcimInterfacesTrace Trace a complete cable path and return each segment as a three-tuple of (termination, cable, termination).
*/
func (a *Client) DcimInterfacesTrace(params *DcimInterfacesTraceParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*DcimInterfacesTraceOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDcimInterfacesTraceParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "dcim_interfaces_trace",
		Method:             "GET",
		PathPattern:        "/dcim/interfaces/{id}/trace/",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &DcimInterfacesTraceReader{formats: a.formats},
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
	success, ok := result.(*DcimInterfacesTraceOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for dcim_interfaces_trace: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  DcimInterfacesUpdate dcim interfaces update API
*/
func (a *Client) DcimInterfacesUpdate(params *DcimInterfacesUpdateParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*DcimInterfacesUpdateOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDcimInterfacesUpdateParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "dcim_interfaces_update",
		Method:             "PUT",
		PathPattern:        "/dcim/interfaces/{id}/",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &DcimInterfacesUpdateReader{formats: a.formats},
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
	success, ok := result.(*DcimInterfacesUpdateOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for dcim_interfaces_update: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  DcimInventoryItemsBulkDelete dcim inventory items bulk delete API
*/
func (a *Client) DcimInventoryItemsBulkDelete(params *DcimInventoryItemsBulkDeleteParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*DcimInventoryItemsBulkDeleteNoContent, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDcimInventoryItemsBulkDeleteParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "dcim_inventory-items_bulk_delete",
		Method:             "DELETE",
		PathPattern:        "/dcim/inventory-items/",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &DcimInventoryItemsBulkDeleteReader{formats: a.formats},
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
	success, ok := result.(*DcimInventoryItemsBulkDeleteNoContent)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for dcim_inventory-items_bulk_delete: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  DcimInventoryItemsBulkPartialUpdate dcim inventory items bulk partial update API
*/
func (a *Client) DcimInventoryItemsBulkPartialUpdate(params *DcimInventoryItemsBulkPartialUpdateParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*DcimInventoryItemsBulkPartialUpdateOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDcimInventoryItemsBulkPartialUpdateParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "dcim_inventory-items_bulk_partial_update",
		Method:             "PATCH",
		PathPattern:        "/dcim/inventory-items/",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &DcimInventoryItemsBulkPartialUpdateReader{formats: a.formats},
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
	success, ok := result.(*DcimInventoryItemsBulkPartialUpdateOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for dcim_inventory-items_bulk_partial_update: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  DcimInventoryItemsBulkUpdate dcim inventory items bulk update API
*/
func (a *Client) DcimInventoryItemsBulkUpdate(params *DcimInventoryItemsBulkUpdateParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*DcimInventoryItemsBulkUpdateOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDcimInventoryItemsBulkUpdateParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "dcim_inventory-items_bulk_update",
		Method:             "PUT",
		PathPattern:        "/dcim/inventory-items/",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &DcimInventoryItemsBulkUpdateReader{formats: a.formats},
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
	success, ok := result.(*DcimInventoryItemsBulkUpdateOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for dcim_inventory-items_bulk_update: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  DcimInventoryItemsCreate dcim inventory items create API
*/
func (a *Client) DcimInventoryItemsCreate(params *DcimInventoryItemsCreateParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*DcimInventoryItemsCreateCreated, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDcimInventoryItemsCreateParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "dcim_inventory-items_create",
		Method:             "POST",
		PathPattern:        "/dcim/inventory-items/",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &DcimInventoryItemsCreateReader{formats: a.formats},
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
	success, ok := result.(*DcimInventoryItemsCreateCreated)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for dcim_inventory-items_create: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  DcimInventoryItemsDelete dcim inventory items delete API
*/
func (a *Client) DcimInventoryItemsDelete(params *DcimInventoryItemsDeleteParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*DcimInventoryItemsDeleteNoContent, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDcimInventoryItemsDeleteParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "dcim_inventory-items_delete",
		Method:             "DELETE",
		PathPattern:        "/dcim/inventory-items/{id}/",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &DcimInventoryItemsDeleteReader{formats: a.formats},
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
	success, ok := result.(*DcimInventoryItemsDeleteNoContent)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for dcim_inventory-items_delete: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  DcimInventoryItemsList dcim inventory items list API
*/
func (a *Client) DcimInventoryItemsList(params *DcimInventoryItemsListParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*DcimInventoryItemsListOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDcimInventoryItemsListParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "dcim_inventory-items_list",
		Method:             "GET",
		PathPattern:        "/dcim/inventory-items/",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &DcimInventoryItemsListReader{formats: a.formats},
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
	success, ok := result.(*DcimInventoryItemsListOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for dcim_inventory-items_list: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  DcimInventoryItemsPartialUpdate dcim inventory items partial update API
*/
func (a *Client) DcimInventoryItemsPartialUpdate(params *DcimInventoryItemsPartialUpdateParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*DcimInventoryItemsPartialUpdateOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDcimInventoryItemsPartialUpdateParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "dcim_inventory-items_partial_update",
		Method:             "PATCH",
		PathPattern:        "/dcim/inventory-items/{id}/",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &DcimInventoryItemsPartialUpdateReader{formats: a.formats},
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
	success, ok := result.(*DcimInventoryItemsPartialUpdateOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for dcim_inventory-items_partial_update: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  DcimInventoryItemsRead dcim inventory items read API
*/
func (a *Client) DcimInventoryItemsRead(params *DcimInventoryItemsReadParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*DcimInventoryItemsReadOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDcimInventoryItemsReadParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "dcim_inventory-items_read",
		Method:             "GET",
		PathPattern:        "/dcim/inventory-items/{id}/",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &DcimInventoryItemsReadReader{formats: a.formats},
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
	success, ok := result.(*DcimInventoryItemsReadOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for dcim_inventory-items_read: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  DcimInventoryItemsUpdate dcim inventory items update API
*/
func (a *Client) DcimInventoryItemsUpdate(params *DcimInventoryItemsUpdateParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*DcimInventoryItemsUpdateOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDcimInventoryItemsUpdateParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "dcim_inventory-items_update",
		Method:             "PUT",
		PathPattern:        "/dcim/inventory-items/{id}/",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &DcimInventoryItemsUpdateReader{formats: a.formats},
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
	success, ok := result.(*DcimInventoryItemsUpdateOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for dcim_inventory-items_update: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  DcimManufacturersBulkDelete dcim manufacturers bulk delete API
*/
func (a *Client) DcimManufacturersBulkDelete(params *DcimManufacturersBulkDeleteParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*DcimManufacturersBulkDeleteNoContent, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDcimManufacturersBulkDeleteParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "dcim_manufacturers_bulk_delete",
		Method:             "DELETE",
		PathPattern:        "/dcim/manufacturers/",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &DcimManufacturersBulkDeleteReader{formats: a.formats},
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
	success, ok := result.(*DcimManufacturersBulkDeleteNoContent)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for dcim_manufacturers_bulk_delete: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  DcimManufacturersBulkPartialUpdate dcim manufacturers bulk partial update API
*/
func (a *Client) DcimManufacturersBulkPartialUpdate(params *DcimManufacturersBulkPartialUpdateParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*DcimManufacturersBulkPartialUpdateOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDcimManufacturersBulkPartialUpdateParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "dcim_manufacturers_bulk_partial_update",
		Method:             "PATCH",
		PathPattern:        "/dcim/manufacturers/",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &DcimManufacturersBulkPartialUpdateReader{formats: a.formats},
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
	success, ok := result.(*DcimManufacturersBulkPartialUpdateOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for dcim_manufacturers_bulk_partial_update: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  DcimManufacturersBulkUpdate dcim manufacturers bulk update API
*/
func (a *Client) DcimManufacturersBulkUpdate(params *DcimManufacturersBulkUpdateParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*DcimManufacturersBulkUpdateOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDcimManufacturersBulkUpdateParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "dcim_manufacturers_bulk_update",
		Method:             "PUT",
		PathPattern:        "/dcim/manufacturers/",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &DcimManufacturersBulkUpdateReader{formats: a.formats},
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
	success, ok := result.(*DcimManufacturersBulkUpdateOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for dcim_manufacturers_bulk_update: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  DcimManufacturersCreate dcim manufacturers create API
*/
func (a *Client) DcimManufacturersCreate(params *DcimManufacturersCreateParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*DcimManufacturersCreateCreated, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDcimManufacturersCreateParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "dcim_manufacturers_create",
		Method:             "POST",
		PathPattern:        "/dcim/manufacturers/",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &DcimManufacturersCreateReader{formats: a.formats},
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
	success, ok := result.(*DcimManufacturersCreateCreated)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for dcim_manufacturers_create: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  DcimManufacturersDelete dcim manufacturers delete API
*/
func (a *Client) DcimManufacturersDelete(params *DcimManufacturersDeleteParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*DcimManufacturersDeleteNoContent, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDcimManufacturersDeleteParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "dcim_manufacturers_delete",
		Method:             "DELETE",
		PathPattern:        "/dcim/manufacturers/{id}/",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &DcimManufacturersDeleteReader{formats: a.formats},
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
	success, ok := result.(*DcimManufacturersDeleteNoContent)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for dcim_manufacturers_delete: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  DcimManufacturersList dcim manufacturers list API
*/
func (a *Client) DcimManufacturersList(params *DcimManufacturersListParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*DcimManufacturersListOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDcimManufacturersListParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "dcim_manufacturers_list",
		Method:             "GET",
		PathPattern:        "/dcim/manufacturers/",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &DcimManufacturersListReader{formats: a.formats},
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
	success, ok := result.(*DcimManufacturersListOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for dcim_manufacturers_list: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  DcimManufacturersPartialUpdate dcim manufacturers partial update API
*/
func (a *Client) DcimManufacturersPartialUpdate(params *DcimManufacturersPartialUpdateParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*DcimManufacturersPartialUpdateOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDcimManufacturersPartialUpdateParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "dcim_manufacturers_partial_update",
		Method:             "PATCH",
		PathPattern:        "/dcim/manufacturers/{id}/",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &DcimManufacturersPartialUpdateReader{formats: a.formats},
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
	success, ok := result.(*DcimManufacturersPartialUpdateOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for dcim_manufacturers_partial_update: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  DcimManufacturersRead dcim manufacturers read API
*/
func (a *Client) DcimManufacturersRead(params *DcimManufacturersReadParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*DcimManufacturersReadOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDcimManufacturersReadParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "dcim_manufacturers_read",
		Method:             "GET",
		PathPattern:        "/dcim/manufacturers/{id}/",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &DcimManufacturersReadReader{formats: a.formats},
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
	success, ok := result.(*DcimManufacturersReadOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for dcim_manufacturers_read: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  DcimManufacturersUpdate dcim manufacturers update API
*/
func (a *Client) DcimManufacturersUpdate(params *DcimManufacturersUpdateParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*DcimManufacturersUpdateOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDcimManufacturersUpdateParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "dcim_manufacturers_update",
		Method:             "PUT",
		PathPattern:        "/dcim/manufacturers/{id}/",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &DcimManufacturersUpdateReader{formats: a.formats},
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
	success, ok := result.(*DcimManufacturersUpdateOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for dcim_manufacturers_update: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  DcimPlatformsBulkDelete dcim platforms bulk delete API
*/
func (a *Client) DcimPlatformsBulkDelete(params *DcimPlatformsBulkDeleteParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*DcimPlatformsBulkDeleteNoContent, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDcimPlatformsBulkDeleteParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "dcim_platforms_bulk_delete",
		Method:             "DELETE",
		PathPattern:        "/dcim/platforms/",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &DcimPlatformsBulkDeleteReader{formats: a.formats},
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
	success, ok := result.(*DcimPlatformsBulkDeleteNoContent)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for dcim_platforms_bulk_delete: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  DcimPlatformsBulkPartialUpdate dcim platforms bulk partial update API
*/
func (a *Client) DcimPlatformsBulkPartialUpdate(params *DcimPlatformsBulkPartialUpdateParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*DcimPlatformsBulkPartialUpdateOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDcimPlatformsBulkPartialUpdateParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "dcim_platforms_bulk_partial_update",
		Method:             "PATCH",
		PathPattern:        "/dcim/platforms/",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &DcimPlatformsBulkPartialUpdateReader{formats: a.formats},
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
	success, ok := result.(*DcimPlatformsBulkPartialUpdateOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for dcim_platforms_bulk_partial_update: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  DcimPlatformsBulkUpdate dcim platforms bulk update API
*/
func (a *Client) DcimPlatformsBulkUpdate(params *DcimPlatformsBulkUpdateParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*DcimPlatformsBulkUpdateOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDcimPlatformsBulkUpdateParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "dcim_platforms_bulk_update",
		Method:             "PUT",
		PathPattern:        "/dcim/platforms/",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &DcimPlatformsBulkUpdateReader{formats: a.formats},
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
	success, ok := result.(*DcimPlatformsBulkUpdateOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for dcim_platforms_bulk_update: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  DcimPlatformsCreate dcim platforms create API
*/
func (a *Client) DcimPlatformsCreate(params *DcimPlatformsCreateParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*DcimPlatformsCreateCreated, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDcimPlatformsCreateParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "dcim_platforms_create",
		Method:             "POST",
		PathPattern:        "/dcim/platforms/",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &DcimPlatformsCreateReader{formats: a.formats},
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
	success, ok := result.(*DcimPlatformsCreateCreated)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for dcim_platforms_create: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  DcimPlatformsDelete dcim platforms delete API
*/
func (a *Client) DcimPlatformsDelete(params *DcimPlatformsDeleteParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*DcimPlatformsDeleteNoContent, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDcimPlatformsDeleteParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "dcim_platforms_delete",
		Method:             "DELETE",
		PathPattern:        "/dcim/platforms/{id}/",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &DcimPlatformsDeleteReader{formats: a.formats},
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
	success, ok := result.(*DcimPlatformsDeleteNoContent)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for dcim_platforms_delete: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  DcimPlatformsList dcim platforms list API
*/
func (a *Client) DcimPlatformsList(params *DcimPlatformsListParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*DcimPlatformsListOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDcimPlatformsListParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "dcim_platforms_list",
		Method:             "GET",
		PathPattern:        "/dcim/platforms/",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &DcimPlatformsListReader{formats: a.formats},
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
	success, ok := result.(*DcimPlatformsListOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for dcim_platforms_list: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  DcimPlatformsPartialUpdate dcim platforms partial update API
*/
func (a *Client) DcimPlatformsPartialUpdate(params *DcimPlatformsPartialUpdateParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*DcimPlatformsPartialUpdateOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDcimPlatformsPartialUpdateParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "dcim_platforms_partial_update",
		Method:             "PATCH",
		PathPattern:        "/dcim/platforms/{id}/",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &DcimPlatformsPartialUpdateReader{formats: a.formats},
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
	success, ok := result.(*DcimPlatformsPartialUpdateOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for dcim_platforms_partial_update: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  DcimPlatformsRead dcim platforms read API
*/
func (a *Client) DcimPlatformsRead(params *DcimPlatformsReadParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*DcimPlatformsReadOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDcimPlatformsReadParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "dcim_platforms_read",
		Method:             "GET",
		PathPattern:        "/dcim/platforms/{id}/",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &DcimPlatformsReadReader{formats: a.formats},
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
	success, ok := result.(*DcimPlatformsReadOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for dcim_platforms_read: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  DcimPlatformsUpdate dcim platforms update API
*/
func (a *Client) DcimPlatformsUpdate(params *DcimPlatformsUpdateParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*DcimPlatformsUpdateOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDcimPlatformsUpdateParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "dcim_platforms_update",
		Method:             "PUT",
		PathPattern:        "/dcim/platforms/{id}/",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &DcimPlatformsUpdateReader{formats: a.formats},
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
	success, ok := result.(*DcimPlatformsUpdateOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for dcim_platforms_update: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  DcimPowerConnectionsList dcim power connections list API
*/
func (a *Client) DcimPowerConnectionsList(params *DcimPowerConnectionsListParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*DcimPowerConnectionsListOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDcimPowerConnectionsListParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "dcim_power-connections_list",
		Method:             "GET",
		PathPattern:        "/dcim/power-connections/",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &DcimPowerConnectionsListReader{formats: a.formats},
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
	success, ok := result.(*DcimPowerConnectionsListOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for dcim_power-connections_list: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  DcimPowerFeedsBulkDelete dcim power feeds bulk delete API
*/
func (a *Client) DcimPowerFeedsBulkDelete(params *DcimPowerFeedsBulkDeleteParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*DcimPowerFeedsBulkDeleteNoContent, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDcimPowerFeedsBulkDeleteParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "dcim_power-feeds_bulk_delete",
		Method:             "DELETE",
		PathPattern:        "/dcim/power-feeds/",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &DcimPowerFeedsBulkDeleteReader{formats: a.formats},
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
	success, ok := result.(*DcimPowerFeedsBulkDeleteNoContent)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for dcim_power-feeds_bulk_delete: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  DcimPowerFeedsBulkPartialUpdate dcim power feeds bulk partial update API
*/
func (a *Client) DcimPowerFeedsBulkPartialUpdate(params *DcimPowerFeedsBulkPartialUpdateParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*DcimPowerFeedsBulkPartialUpdateOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDcimPowerFeedsBulkPartialUpdateParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "dcim_power-feeds_bulk_partial_update",
		Method:             "PATCH",
		PathPattern:        "/dcim/power-feeds/",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &DcimPowerFeedsBulkPartialUpdateReader{formats: a.formats},
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
	success, ok := result.(*DcimPowerFeedsBulkPartialUpdateOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for dcim_power-feeds_bulk_partial_update: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  DcimPowerFeedsBulkUpdate dcim power feeds bulk update API
*/
func (a *Client) DcimPowerFeedsBulkUpdate(params *DcimPowerFeedsBulkUpdateParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*DcimPowerFeedsBulkUpdateOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDcimPowerFeedsBulkUpdateParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "dcim_power-feeds_bulk_update",
		Method:             "PUT",
		PathPattern:        "/dcim/power-feeds/",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &DcimPowerFeedsBulkUpdateReader{formats: a.formats},
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
	success, ok := result.(*DcimPowerFeedsBulkUpdateOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for dcim_power-feeds_bulk_update: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  DcimPowerFeedsCreate dcim power feeds create API
*/
func (a *Client) DcimPowerFeedsCreate(params *DcimPowerFeedsCreateParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*DcimPowerFeedsCreateCreated, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDcimPowerFeedsCreateParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "dcim_power-feeds_create",
		Method:             "POST",
		PathPattern:        "/dcim/power-feeds/",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &DcimPowerFeedsCreateReader{formats: a.formats},
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
	success, ok := result.(*DcimPowerFeedsCreateCreated)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for dcim_power-feeds_create: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  DcimPowerFeedsDelete dcim power feeds delete API
*/
func (a *Client) DcimPowerFeedsDelete(params *DcimPowerFeedsDeleteParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*DcimPowerFeedsDeleteNoContent, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDcimPowerFeedsDeleteParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "dcim_power-feeds_delete",
		Method:             "DELETE",
		PathPattern:        "/dcim/power-feeds/{id}/",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &DcimPowerFeedsDeleteReader{formats: a.formats},
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
	success, ok := result.(*DcimPowerFeedsDeleteNoContent)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for dcim_power-feeds_delete: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  DcimPowerFeedsList dcim power feeds list API
*/
func (a *Client) DcimPowerFeedsList(params *DcimPowerFeedsListParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*DcimPowerFeedsListOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDcimPowerFeedsListParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "dcim_power-feeds_list",
		Method:             "GET",
		PathPattern:        "/dcim/power-feeds/",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &DcimPowerFeedsListReader{formats: a.formats},
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
	success, ok := result.(*DcimPowerFeedsListOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for dcim_power-feeds_list: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  DcimPowerFeedsPartialUpdate dcim power feeds partial update API
*/
func (a *Client) DcimPowerFeedsPartialUpdate(params *DcimPowerFeedsPartialUpdateParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*DcimPowerFeedsPartialUpdateOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDcimPowerFeedsPartialUpdateParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "dcim_power-feeds_partial_update",
		Method:             "PATCH",
		PathPattern:        "/dcim/power-feeds/{id}/",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &DcimPowerFeedsPartialUpdateReader{formats: a.formats},
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
	success, ok := result.(*DcimPowerFeedsPartialUpdateOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for dcim_power-feeds_partial_update: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  DcimPowerFeedsRead dcim power feeds read API
*/
func (a *Client) DcimPowerFeedsRead(params *DcimPowerFeedsReadParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*DcimPowerFeedsReadOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDcimPowerFeedsReadParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "dcim_power-feeds_read",
		Method:             "GET",
		PathPattern:        "/dcim/power-feeds/{id}/",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &DcimPowerFeedsReadReader{formats: a.formats},
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
	success, ok := result.(*DcimPowerFeedsReadOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for dcim_power-feeds_read: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  DcimPowerFeedsTrace Trace a complete cable path and return each segment as a three-tuple of (termination, cable, termination).
*/
func (a *Client) DcimPowerFeedsTrace(params *DcimPowerFeedsTraceParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*DcimPowerFeedsTraceOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDcimPowerFeedsTraceParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "dcim_power-feeds_trace",
		Method:             "GET",
		PathPattern:        "/dcim/power-feeds/{id}/trace/",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &DcimPowerFeedsTraceReader{formats: a.formats},
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
	success, ok := result.(*DcimPowerFeedsTraceOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for dcim_power-feeds_trace: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  DcimPowerFeedsUpdate dcim power feeds update API
*/
func (a *Client) DcimPowerFeedsUpdate(params *DcimPowerFeedsUpdateParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*DcimPowerFeedsUpdateOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDcimPowerFeedsUpdateParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "dcim_power-feeds_update",
		Method:             "PUT",
		PathPattern:        "/dcim/power-feeds/{id}/",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &DcimPowerFeedsUpdateReader{formats: a.formats},
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
	success, ok := result.(*DcimPowerFeedsUpdateOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for dcim_power-feeds_update: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  DcimPowerOutletTemplatesBulkDelete dcim power outlet templates bulk delete API
*/
func (a *Client) DcimPowerOutletTemplatesBulkDelete(params *DcimPowerOutletTemplatesBulkDeleteParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*DcimPowerOutletTemplatesBulkDeleteNoContent, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDcimPowerOutletTemplatesBulkDeleteParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "dcim_power-outlet-templates_bulk_delete",
		Method:             "DELETE",
		PathPattern:        "/dcim/power-outlet-templates/",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &DcimPowerOutletTemplatesBulkDeleteReader{formats: a.formats},
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
	success, ok := result.(*DcimPowerOutletTemplatesBulkDeleteNoContent)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for dcim_power-outlet-templates_bulk_delete: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  DcimPowerOutletTemplatesBulkPartialUpdate dcim power outlet templates bulk partial update API
*/
func (a *Client) DcimPowerOutletTemplatesBulkPartialUpdate(params *DcimPowerOutletTemplatesBulkPartialUpdateParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*DcimPowerOutletTemplatesBulkPartialUpdateOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDcimPowerOutletTemplatesBulkPartialUpdateParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "dcim_power-outlet-templates_bulk_partial_update",
		Method:             "PATCH",
		PathPattern:        "/dcim/power-outlet-templates/",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &DcimPowerOutletTemplatesBulkPartialUpdateReader{formats: a.formats},
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
	success, ok := result.(*DcimPowerOutletTemplatesBulkPartialUpdateOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for dcim_power-outlet-templates_bulk_partial_update: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  DcimPowerOutletTemplatesBulkUpdate dcim power outlet templates bulk update API
*/
func (a *Client) DcimPowerOutletTemplatesBulkUpdate(params *DcimPowerOutletTemplatesBulkUpdateParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*DcimPowerOutletTemplatesBulkUpdateOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDcimPowerOutletTemplatesBulkUpdateParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "dcim_power-outlet-templates_bulk_update",
		Method:             "PUT",
		PathPattern:        "/dcim/power-outlet-templates/",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &DcimPowerOutletTemplatesBulkUpdateReader{formats: a.formats},
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
	success, ok := result.(*DcimPowerOutletTemplatesBulkUpdateOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for dcim_power-outlet-templates_bulk_update: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  DcimPowerOutletTemplatesCreate dcim power outlet templates create API
*/
func (a *Client) DcimPowerOutletTemplatesCreate(params *DcimPowerOutletTemplatesCreateParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*DcimPowerOutletTemplatesCreateCreated, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDcimPowerOutletTemplatesCreateParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "dcim_power-outlet-templates_create",
		Method:             "POST",
		PathPattern:        "/dcim/power-outlet-templates/",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &DcimPowerOutletTemplatesCreateReader{formats: a.formats},
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
	success, ok := result.(*DcimPowerOutletTemplatesCreateCreated)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for dcim_power-outlet-templates_create: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  DcimPowerOutletTemplatesDelete dcim power outlet templates delete API
*/
func (a *Client) DcimPowerOutletTemplatesDelete(params *DcimPowerOutletTemplatesDeleteParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*DcimPowerOutletTemplatesDeleteNoContent, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDcimPowerOutletTemplatesDeleteParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "dcim_power-outlet-templates_delete",
		Method:             "DELETE",
		PathPattern:        "/dcim/power-outlet-templates/{id}/",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &DcimPowerOutletTemplatesDeleteReader{formats: a.formats},
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
	success, ok := result.(*DcimPowerOutletTemplatesDeleteNoContent)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for dcim_power-outlet-templates_delete: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  DcimPowerOutletTemplatesList dcim power outlet templates list API
*/
func (a *Client) DcimPowerOutletTemplatesList(params *DcimPowerOutletTemplatesListParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*DcimPowerOutletTemplatesListOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDcimPowerOutletTemplatesListParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "dcim_power-outlet-templates_list",
		Method:             "GET",
		PathPattern:        "/dcim/power-outlet-templates/",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &DcimPowerOutletTemplatesListReader{formats: a.formats},
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
	success, ok := result.(*DcimPowerOutletTemplatesListOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for dcim_power-outlet-templates_list: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  DcimPowerOutletTemplatesPartialUpdate dcim power outlet templates partial update API
*/
func (a *Client) DcimPowerOutletTemplatesPartialUpdate(params *DcimPowerOutletTemplatesPartialUpdateParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*DcimPowerOutletTemplatesPartialUpdateOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDcimPowerOutletTemplatesPartialUpdateParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "dcim_power-outlet-templates_partial_update",
		Method:             "PATCH",
		PathPattern:        "/dcim/power-outlet-templates/{id}/",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &DcimPowerOutletTemplatesPartialUpdateReader{formats: a.formats},
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
	success, ok := result.(*DcimPowerOutletTemplatesPartialUpdateOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for dcim_power-outlet-templates_partial_update: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  DcimPowerOutletTemplatesRead dcim power outlet templates read API
*/
func (a *Client) DcimPowerOutletTemplatesRead(params *DcimPowerOutletTemplatesReadParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*DcimPowerOutletTemplatesReadOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDcimPowerOutletTemplatesReadParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "dcim_power-outlet-templates_read",
		Method:             "GET",
		PathPattern:        "/dcim/power-outlet-templates/{id}/",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &DcimPowerOutletTemplatesReadReader{formats: a.formats},
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
	success, ok := result.(*DcimPowerOutletTemplatesReadOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for dcim_power-outlet-templates_read: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  DcimPowerOutletTemplatesUpdate dcim power outlet templates update API
*/
func (a *Client) DcimPowerOutletTemplatesUpdate(params *DcimPowerOutletTemplatesUpdateParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*DcimPowerOutletTemplatesUpdateOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDcimPowerOutletTemplatesUpdateParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "dcim_power-outlet-templates_update",
		Method:             "PUT",
		PathPattern:        "/dcim/power-outlet-templates/{id}/",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &DcimPowerOutletTemplatesUpdateReader{formats: a.formats},
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
	success, ok := result.(*DcimPowerOutletTemplatesUpdateOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for dcim_power-outlet-templates_update: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  DcimPowerOutletsBulkDelete dcim power outlets bulk delete API
*/
func (a *Client) DcimPowerOutletsBulkDelete(params *DcimPowerOutletsBulkDeleteParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*DcimPowerOutletsBulkDeleteNoContent, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDcimPowerOutletsBulkDeleteParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "dcim_power-outlets_bulk_delete",
		Method:             "DELETE",
		PathPattern:        "/dcim/power-outlets/",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &DcimPowerOutletsBulkDeleteReader{formats: a.formats},
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
	success, ok := result.(*DcimPowerOutletsBulkDeleteNoContent)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for dcim_power-outlets_bulk_delete: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  DcimPowerOutletsBulkPartialUpdate dcim power outlets bulk partial update API
*/
func (a *Client) DcimPowerOutletsBulkPartialUpdate(params *DcimPowerOutletsBulkPartialUpdateParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*DcimPowerOutletsBulkPartialUpdateOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDcimPowerOutletsBulkPartialUpdateParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "dcim_power-outlets_bulk_partial_update",
		Method:             "PATCH",
		PathPattern:        "/dcim/power-outlets/",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &DcimPowerOutletsBulkPartialUpdateReader{formats: a.formats},
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
	success, ok := result.(*DcimPowerOutletsBulkPartialUpdateOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for dcim_power-outlets_bulk_partial_update: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  DcimPowerOutletsBulkUpdate dcim power outlets bulk update API
*/
func (a *Client) DcimPowerOutletsBulkUpdate(params *DcimPowerOutletsBulkUpdateParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*DcimPowerOutletsBulkUpdateOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDcimPowerOutletsBulkUpdateParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "dcim_power-outlets_bulk_update",
		Method:             "PUT",
		PathPattern:        "/dcim/power-outlets/",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &DcimPowerOutletsBulkUpdateReader{formats: a.formats},
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
	success, ok := result.(*DcimPowerOutletsBulkUpdateOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for dcim_power-outlets_bulk_update: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  DcimPowerOutletsCreate dcim power outlets create API
*/
func (a *Client) DcimPowerOutletsCreate(params *DcimPowerOutletsCreateParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*DcimPowerOutletsCreateCreated, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDcimPowerOutletsCreateParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "dcim_power-outlets_create",
		Method:             "POST",
		PathPattern:        "/dcim/power-outlets/",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &DcimPowerOutletsCreateReader{formats: a.formats},
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
	success, ok := result.(*DcimPowerOutletsCreateCreated)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for dcim_power-outlets_create: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  DcimPowerOutletsDelete dcim power outlets delete API
*/
func (a *Client) DcimPowerOutletsDelete(params *DcimPowerOutletsDeleteParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*DcimPowerOutletsDeleteNoContent, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDcimPowerOutletsDeleteParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "dcim_power-outlets_delete",
		Method:             "DELETE",
		PathPattern:        "/dcim/power-outlets/{id}/",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &DcimPowerOutletsDeleteReader{formats: a.formats},
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
	success, ok := result.(*DcimPowerOutletsDeleteNoContent)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for dcim_power-outlets_delete: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  DcimPowerOutletsList dcim power outlets list API
*/
func (a *Client) DcimPowerOutletsList(params *DcimPowerOutletsListParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*DcimPowerOutletsListOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDcimPowerOutletsListParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "dcim_power-outlets_list",
		Method:             "GET",
		PathPattern:        "/dcim/power-outlets/",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &DcimPowerOutletsListReader{formats: a.formats},
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
	success, ok := result.(*DcimPowerOutletsListOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for dcim_power-outlets_list: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  DcimPowerOutletsPartialUpdate dcim power outlets partial update API
*/
func (a *Client) DcimPowerOutletsPartialUpdate(params *DcimPowerOutletsPartialUpdateParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*DcimPowerOutletsPartialUpdateOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDcimPowerOutletsPartialUpdateParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "dcim_power-outlets_partial_update",
		Method:             "PATCH",
		PathPattern:        "/dcim/power-outlets/{id}/",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &DcimPowerOutletsPartialUpdateReader{formats: a.formats},
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
	success, ok := result.(*DcimPowerOutletsPartialUpdateOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for dcim_power-outlets_partial_update: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  DcimPowerOutletsRead dcim power outlets read API
*/
func (a *Client) DcimPowerOutletsRead(params *DcimPowerOutletsReadParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*DcimPowerOutletsReadOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDcimPowerOutletsReadParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "dcim_power-outlets_read",
		Method:             "GET",
		PathPattern:        "/dcim/power-outlets/{id}/",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &DcimPowerOutletsReadReader{formats: a.formats},
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
	success, ok := result.(*DcimPowerOutletsReadOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for dcim_power-outlets_read: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  DcimPowerOutletsTrace Trace a complete cable path and return each segment as a three-tuple of (termination, cable, termination).
*/
func (a *Client) DcimPowerOutletsTrace(params *DcimPowerOutletsTraceParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*DcimPowerOutletsTraceOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDcimPowerOutletsTraceParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "dcim_power-outlets_trace",
		Method:             "GET",
		PathPattern:        "/dcim/power-outlets/{id}/trace/",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &DcimPowerOutletsTraceReader{formats: a.formats},
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
	success, ok := result.(*DcimPowerOutletsTraceOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for dcim_power-outlets_trace: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  DcimPowerOutletsUpdate dcim power outlets update API
*/
func (a *Client) DcimPowerOutletsUpdate(params *DcimPowerOutletsUpdateParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*DcimPowerOutletsUpdateOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDcimPowerOutletsUpdateParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "dcim_power-outlets_update",
		Method:             "PUT",
		PathPattern:        "/dcim/power-outlets/{id}/",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &DcimPowerOutletsUpdateReader{formats: a.formats},
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
	success, ok := result.(*DcimPowerOutletsUpdateOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for dcim_power-outlets_update: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  DcimPowerPanelsBulkDelete dcim power panels bulk delete API
*/
func (a *Client) DcimPowerPanelsBulkDelete(params *DcimPowerPanelsBulkDeleteParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*DcimPowerPanelsBulkDeleteNoContent, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDcimPowerPanelsBulkDeleteParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "dcim_power-panels_bulk_delete",
		Method:             "DELETE",
		PathPattern:        "/dcim/power-panels/",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &DcimPowerPanelsBulkDeleteReader{formats: a.formats},
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
	success, ok := result.(*DcimPowerPanelsBulkDeleteNoContent)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for dcim_power-panels_bulk_delete: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  DcimPowerPanelsBulkPartialUpdate dcim power panels bulk partial update API
*/
func (a *Client) DcimPowerPanelsBulkPartialUpdate(params *DcimPowerPanelsBulkPartialUpdateParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*DcimPowerPanelsBulkPartialUpdateOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDcimPowerPanelsBulkPartialUpdateParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "dcim_power-panels_bulk_partial_update",
		Method:             "PATCH",
		PathPattern:        "/dcim/power-panels/",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &DcimPowerPanelsBulkPartialUpdateReader{formats: a.formats},
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
	success, ok := result.(*DcimPowerPanelsBulkPartialUpdateOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for dcim_power-panels_bulk_partial_update: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  DcimPowerPanelsBulkUpdate dcim power panels bulk update API
*/
func (a *Client) DcimPowerPanelsBulkUpdate(params *DcimPowerPanelsBulkUpdateParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*DcimPowerPanelsBulkUpdateOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDcimPowerPanelsBulkUpdateParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "dcim_power-panels_bulk_update",
		Method:             "PUT",
		PathPattern:        "/dcim/power-panels/",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &DcimPowerPanelsBulkUpdateReader{formats: a.formats},
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
	success, ok := result.(*DcimPowerPanelsBulkUpdateOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for dcim_power-panels_bulk_update: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  DcimPowerPanelsCreate dcim power panels create API
*/
func (a *Client) DcimPowerPanelsCreate(params *DcimPowerPanelsCreateParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*DcimPowerPanelsCreateCreated, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDcimPowerPanelsCreateParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "dcim_power-panels_create",
		Method:             "POST",
		PathPattern:        "/dcim/power-panels/",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &DcimPowerPanelsCreateReader{formats: a.formats},
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
	success, ok := result.(*DcimPowerPanelsCreateCreated)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for dcim_power-panels_create: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  DcimPowerPanelsDelete dcim power panels delete API
*/
func (a *Client) DcimPowerPanelsDelete(params *DcimPowerPanelsDeleteParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*DcimPowerPanelsDeleteNoContent, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDcimPowerPanelsDeleteParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "dcim_power-panels_delete",
		Method:             "DELETE",
		PathPattern:        "/dcim/power-panels/{id}/",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &DcimPowerPanelsDeleteReader{formats: a.formats},
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
	success, ok := result.(*DcimPowerPanelsDeleteNoContent)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for dcim_power-panels_delete: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  DcimPowerPanelsList dcim power panels list API
*/
func (a *Client) DcimPowerPanelsList(params *DcimPowerPanelsListParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*DcimPowerPanelsListOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDcimPowerPanelsListParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "dcim_power-panels_list",
		Method:             "GET",
		PathPattern:        "/dcim/power-panels/",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &DcimPowerPanelsListReader{formats: a.formats},
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
	success, ok := result.(*DcimPowerPanelsListOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for dcim_power-panels_list: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  DcimPowerPanelsPartialUpdate dcim power panels partial update API
*/
func (a *Client) DcimPowerPanelsPartialUpdate(params *DcimPowerPanelsPartialUpdateParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*DcimPowerPanelsPartialUpdateOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDcimPowerPanelsPartialUpdateParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "dcim_power-panels_partial_update",
		Method:             "PATCH",
		PathPattern:        "/dcim/power-panels/{id}/",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &DcimPowerPanelsPartialUpdateReader{formats: a.formats},
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
	success, ok := result.(*DcimPowerPanelsPartialUpdateOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for dcim_power-panels_partial_update: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  DcimPowerPanelsRead dcim power panels read API
*/
func (a *Client) DcimPowerPanelsRead(params *DcimPowerPanelsReadParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*DcimPowerPanelsReadOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDcimPowerPanelsReadParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "dcim_power-panels_read",
		Method:             "GET",
		PathPattern:        "/dcim/power-panels/{id}/",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &DcimPowerPanelsReadReader{formats: a.formats},
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
	success, ok := result.(*DcimPowerPanelsReadOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for dcim_power-panels_read: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  DcimPowerPanelsUpdate dcim power panels update API
*/
func (a *Client) DcimPowerPanelsUpdate(params *DcimPowerPanelsUpdateParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*DcimPowerPanelsUpdateOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDcimPowerPanelsUpdateParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "dcim_power-panels_update",
		Method:             "PUT",
		PathPattern:        "/dcim/power-panels/{id}/",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &DcimPowerPanelsUpdateReader{formats: a.formats},
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
	success, ok := result.(*DcimPowerPanelsUpdateOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for dcim_power-panels_update: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  DcimPowerPortTemplatesBulkDelete dcim power port templates bulk delete API
*/
func (a *Client) DcimPowerPortTemplatesBulkDelete(params *DcimPowerPortTemplatesBulkDeleteParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*DcimPowerPortTemplatesBulkDeleteNoContent, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDcimPowerPortTemplatesBulkDeleteParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "dcim_power-port-templates_bulk_delete",
		Method:             "DELETE",
		PathPattern:        "/dcim/power-port-templates/",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &DcimPowerPortTemplatesBulkDeleteReader{formats: a.formats},
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
	success, ok := result.(*DcimPowerPortTemplatesBulkDeleteNoContent)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for dcim_power-port-templates_bulk_delete: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  DcimPowerPortTemplatesBulkPartialUpdate dcim power port templates bulk partial update API
*/
func (a *Client) DcimPowerPortTemplatesBulkPartialUpdate(params *DcimPowerPortTemplatesBulkPartialUpdateParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*DcimPowerPortTemplatesBulkPartialUpdateOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDcimPowerPortTemplatesBulkPartialUpdateParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "dcim_power-port-templates_bulk_partial_update",
		Method:             "PATCH",
		PathPattern:        "/dcim/power-port-templates/",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &DcimPowerPortTemplatesBulkPartialUpdateReader{formats: a.formats},
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
	success, ok := result.(*DcimPowerPortTemplatesBulkPartialUpdateOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for dcim_power-port-templates_bulk_partial_update: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  DcimPowerPortTemplatesBulkUpdate dcim power port templates bulk update API
*/
func (a *Client) DcimPowerPortTemplatesBulkUpdate(params *DcimPowerPortTemplatesBulkUpdateParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*DcimPowerPortTemplatesBulkUpdateOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDcimPowerPortTemplatesBulkUpdateParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "dcim_power-port-templates_bulk_update",
		Method:             "PUT",
		PathPattern:        "/dcim/power-port-templates/",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &DcimPowerPortTemplatesBulkUpdateReader{formats: a.formats},
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
	success, ok := result.(*DcimPowerPortTemplatesBulkUpdateOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for dcim_power-port-templates_bulk_update: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  DcimPowerPortTemplatesCreate dcim power port templates create API
*/
func (a *Client) DcimPowerPortTemplatesCreate(params *DcimPowerPortTemplatesCreateParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*DcimPowerPortTemplatesCreateCreated, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDcimPowerPortTemplatesCreateParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "dcim_power-port-templates_create",
		Method:             "POST",
		PathPattern:        "/dcim/power-port-templates/",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &DcimPowerPortTemplatesCreateReader{formats: a.formats},
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
	success, ok := result.(*DcimPowerPortTemplatesCreateCreated)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for dcim_power-port-templates_create: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  DcimPowerPortTemplatesDelete dcim power port templates delete API
*/
func (a *Client) DcimPowerPortTemplatesDelete(params *DcimPowerPortTemplatesDeleteParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*DcimPowerPortTemplatesDeleteNoContent, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDcimPowerPortTemplatesDeleteParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "dcim_power-port-templates_delete",
		Method:             "DELETE",
		PathPattern:        "/dcim/power-port-templates/{id}/",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &DcimPowerPortTemplatesDeleteReader{formats: a.formats},
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
	success, ok := result.(*DcimPowerPortTemplatesDeleteNoContent)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for dcim_power-port-templates_delete: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  DcimPowerPortTemplatesList dcim power port templates list API
*/
func (a *Client) DcimPowerPortTemplatesList(params *DcimPowerPortTemplatesListParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*DcimPowerPortTemplatesListOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDcimPowerPortTemplatesListParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "dcim_power-port-templates_list",
		Method:             "GET",
		PathPattern:        "/dcim/power-port-templates/",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &DcimPowerPortTemplatesListReader{formats: a.formats},
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
	success, ok := result.(*DcimPowerPortTemplatesListOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for dcim_power-port-templates_list: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  DcimPowerPortTemplatesPartialUpdate dcim power port templates partial update API
*/
func (a *Client) DcimPowerPortTemplatesPartialUpdate(params *DcimPowerPortTemplatesPartialUpdateParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*DcimPowerPortTemplatesPartialUpdateOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDcimPowerPortTemplatesPartialUpdateParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "dcim_power-port-templates_partial_update",
		Method:             "PATCH",
		PathPattern:        "/dcim/power-port-templates/{id}/",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &DcimPowerPortTemplatesPartialUpdateReader{formats: a.formats},
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
	success, ok := result.(*DcimPowerPortTemplatesPartialUpdateOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for dcim_power-port-templates_partial_update: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  DcimPowerPortTemplatesRead dcim power port templates read API
*/
func (a *Client) DcimPowerPortTemplatesRead(params *DcimPowerPortTemplatesReadParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*DcimPowerPortTemplatesReadOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDcimPowerPortTemplatesReadParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "dcim_power-port-templates_read",
		Method:             "GET",
		PathPattern:        "/dcim/power-port-templates/{id}/",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &DcimPowerPortTemplatesReadReader{formats: a.formats},
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
	success, ok := result.(*DcimPowerPortTemplatesReadOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for dcim_power-port-templates_read: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  DcimPowerPortTemplatesUpdate dcim power port templates update API
*/
func (a *Client) DcimPowerPortTemplatesUpdate(params *DcimPowerPortTemplatesUpdateParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*DcimPowerPortTemplatesUpdateOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDcimPowerPortTemplatesUpdateParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "dcim_power-port-templates_update",
		Method:             "PUT",
		PathPattern:        "/dcim/power-port-templates/{id}/",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &DcimPowerPortTemplatesUpdateReader{formats: a.formats},
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
	success, ok := result.(*DcimPowerPortTemplatesUpdateOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for dcim_power-port-templates_update: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  DcimPowerPortsBulkDelete dcim power ports bulk delete API
*/
func (a *Client) DcimPowerPortsBulkDelete(params *DcimPowerPortsBulkDeleteParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*DcimPowerPortsBulkDeleteNoContent, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDcimPowerPortsBulkDeleteParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "dcim_power-ports_bulk_delete",
		Method:             "DELETE",
		PathPattern:        "/dcim/power-ports/",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &DcimPowerPortsBulkDeleteReader{formats: a.formats},
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
	success, ok := result.(*DcimPowerPortsBulkDeleteNoContent)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for dcim_power-ports_bulk_delete: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  DcimPowerPortsBulkPartialUpdate dcim power ports bulk partial update API
*/
func (a *Client) DcimPowerPortsBulkPartialUpdate(params *DcimPowerPortsBulkPartialUpdateParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*DcimPowerPortsBulkPartialUpdateOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDcimPowerPortsBulkPartialUpdateParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "dcim_power-ports_bulk_partial_update",
		Method:             "PATCH",
		PathPattern:        "/dcim/power-ports/",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &DcimPowerPortsBulkPartialUpdateReader{formats: a.formats},
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
	success, ok := result.(*DcimPowerPortsBulkPartialUpdateOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for dcim_power-ports_bulk_partial_update: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  DcimPowerPortsBulkUpdate dcim power ports bulk update API
*/
func (a *Client) DcimPowerPortsBulkUpdate(params *DcimPowerPortsBulkUpdateParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*DcimPowerPortsBulkUpdateOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDcimPowerPortsBulkUpdateParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "dcim_power-ports_bulk_update",
		Method:             "PUT",
		PathPattern:        "/dcim/power-ports/",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &DcimPowerPortsBulkUpdateReader{formats: a.formats},
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
	success, ok := result.(*DcimPowerPortsBulkUpdateOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for dcim_power-ports_bulk_update: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  DcimPowerPortsCreate dcim power ports create API
*/
func (a *Client) DcimPowerPortsCreate(params *DcimPowerPortsCreateParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*DcimPowerPortsCreateCreated, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDcimPowerPortsCreateParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "dcim_power-ports_create",
		Method:             "POST",
		PathPattern:        "/dcim/power-ports/",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &DcimPowerPortsCreateReader{formats: a.formats},
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
	success, ok := result.(*DcimPowerPortsCreateCreated)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for dcim_power-ports_create: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  DcimPowerPortsDelete dcim power ports delete API
*/
func (a *Client) DcimPowerPortsDelete(params *DcimPowerPortsDeleteParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*DcimPowerPortsDeleteNoContent, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDcimPowerPortsDeleteParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "dcim_power-ports_delete",
		Method:             "DELETE",
		PathPattern:        "/dcim/power-ports/{id}/",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &DcimPowerPortsDeleteReader{formats: a.formats},
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
	success, ok := result.(*DcimPowerPortsDeleteNoContent)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for dcim_power-ports_delete: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  DcimPowerPortsList dcim power ports list API
*/
func (a *Client) DcimPowerPortsList(params *DcimPowerPortsListParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*DcimPowerPortsListOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDcimPowerPortsListParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "dcim_power-ports_list",
		Method:             "GET",
		PathPattern:        "/dcim/power-ports/",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &DcimPowerPortsListReader{formats: a.formats},
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
	success, ok := result.(*DcimPowerPortsListOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for dcim_power-ports_list: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  DcimPowerPortsPartialUpdate dcim power ports partial update API
*/
func (a *Client) DcimPowerPortsPartialUpdate(params *DcimPowerPortsPartialUpdateParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*DcimPowerPortsPartialUpdateOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDcimPowerPortsPartialUpdateParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "dcim_power-ports_partial_update",
		Method:             "PATCH",
		PathPattern:        "/dcim/power-ports/{id}/",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &DcimPowerPortsPartialUpdateReader{formats: a.formats},
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
	success, ok := result.(*DcimPowerPortsPartialUpdateOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for dcim_power-ports_partial_update: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  DcimPowerPortsRead dcim power ports read API
*/
func (a *Client) DcimPowerPortsRead(params *DcimPowerPortsReadParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*DcimPowerPortsReadOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDcimPowerPortsReadParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "dcim_power-ports_read",
		Method:             "GET",
		PathPattern:        "/dcim/power-ports/{id}/",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &DcimPowerPortsReadReader{formats: a.formats},
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
	success, ok := result.(*DcimPowerPortsReadOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for dcim_power-ports_read: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  DcimPowerPortsTrace Trace a complete cable path and return each segment as a three-tuple of (termination, cable, termination).
*/
func (a *Client) DcimPowerPortsTrace(params *DcimPowerPortsTraceParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*DcimPowerPortsTraceOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDcimPowerPortsTraceParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "dcim_power-ports_trace",
		Method:             "GET",
		PathPattern:        "/dcim/power-ports/{id}/trace/",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &DcimPowerPortsTraceReader{formats: a.formats},
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
	success, ok := result.(*DcimPowerPortsTraceOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for dcim_power-ports_trace: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  DcimPowerPortsUpdate dcim power ports update API
*/
func (a *Client) DcimPowerPortsUpdate(params *DcimPowerPortsUpdateParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*DcimPowerPortsUpdateOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDcimPowerPortsUpdateParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "dcim_power-ports_update",
		Method:             "PUT",
		PathPattern:        "/dcim/power-ports/{id}/",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &DcimPowerPortsUpdateReader{formats: a.formats},
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
	success, ok := result.(*DcimPowerPortsUpdateOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for dcim_power-ports_update: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  DcimRackGroupsBulkDelete dcim rack groups bulk delete API
*/
func (a *Client) DcimRackGroupsBulkDelete(params *DcimRackGroupsBulkDeleteParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*DcimRackGroupsBulkDeleteNoContent, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDcimRackGroupsBulkDeleteParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "dcim_rack-groups_bulk_delete",
		Method:             "DELETE",
		PathPattern:        "/dcim/rack-groups/",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &DcimRackGroupsBulkDeleteReader{formats: a.formats},
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
	success, ok := result.(*DcimRackGroupsBulkDeleteNoContent)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for dcim_rack-groups_bulk_delete: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  DcimRackGroupsBulkPartialUpdate dcim rack groups bulk partial update API
*/
func (a *Client) DcimRackGroupsBulkPartialUpdate(params *DcimRackGroupsBulkPartialUpdateParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*DcimRackGroupsBulkPartialUpdateOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDcimRackGroupsBulkPartialUpdateParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "dcim_rack-groups_bulk_partial_update",
		Method:             "PATCH",
		PathPattern:        "/dcim/rack-groups/",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &DcimRackGroupsBulkPartialUpdateReader{formats: a.formats},
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
	success, ok := result.(*DcimRackGroupsBulkPartialUpdateOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for dcim_rack-groups_bulk_partial_update: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  DcimRackGroupsBulkUpdate dcim rack groups bulk update API
*/
func (a *Client) DcimRackGroupsBulkUpdate(params *DcimRackGroupsBulkUpdateParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*DcimRackGroupsBulkUpdateOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDcimRackGroupsBulkUpdateParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "dcim_rack-groups_bulk_update",
		Method:             "PUT",
		PathPattern:        "/dcim/rack-groups/",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &DcimRackGroupsBulkUpdateReader{formats: a.formats},
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
	success, ok := result.(*DcimRackGroupsBulkUpdateOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for dcim_rack-groups_bulk_update: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  DcimRackGroupsCreate dcim rack groups create API
*/
func (a *Client) DcimRackGroupsCreate(params *DcimRackGroupsCreateParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*DcimRackGroupsCreateCreated, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDcimRackGroupsCreateParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "dcim_rack-groups_create",
		Method:             "POST",
		PathPattern:        "/dcim/rack-groups/",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &DcimRackGroupsCreateReader{formats: a.formats},
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
	success, ok := result.(*DcimRackGroupsCreateCreated)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for dcim_rack-groups_create: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  DcimRackGroupsDelete dcim rack groups delete API
*/
func (a *Client) DcimRackGroupsDelete(params *DcimRackGroupsDeleteParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*DcimRackGroupsDeleteNoContent, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDcimRackGroupsDeleteParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "dcim_rack-groups_delete",
		Method:             "DELETE",
		PathPattern:        "/dcim/rack-groups/{id}/",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &DcimRackGroupsDeleteReader{formats: a.formats},
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
	success, ok := result.(*DcimRackGroupsDeleteNoContent)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for dcim_rack-groups_delete: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  DcimRackGroupsList dcim rack groups list API
*/
func (a *Client) DcimRackGroupsList(params *DcimRackGroupsListParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*DcimRackGroupsListOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDcimRackGroupsListParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "dcim_rack-groups_list",
		Method:             "GET",
		PathPattern:        "/dcim/rack-groups/",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &DcimRackGroupsListReader{formats: a.formats},
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
	success, ok := result.(*DcimRackGroupsListOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for dcim_rack-groups_list: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  DcimRackGroupsPartialUpdate dcim rack groups partial update API
*/
func (a *Client) DcimRackGroupsPartialUpdate(params *DcimRackGroupsPartialUpdateParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*DcimRackGroupsPartialUpdateOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDcimRackGroupsPartialUpdateParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "dcim_rack-groups_partial_update",
		Method:             "PATCH",
		PathPattern:        "/dcim/rack-groups/{id}/",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &DcimRackGroupsPartialUpdateReader{formats: a.formats},
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
	success, ok := result.(*DcimRackGroupsPartialUpdateOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for dcim_rack-groups_partial_update: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  DcimRackGroupsRead dcim rack groups read API
*/
func (a *Client) DcimRackGroupsRead(params *DcimRackGroupsReadParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*DcimRackGroupsReadOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDcimRackGroupsReadParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "dcim_rack-groups_read",
		Method:             "GET",
		PathPattern:        "/dcim/rack-groups/{id}/",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &DcimRackGroupsReadReader{formats: a.formats},
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
	success, ok := result.(*DcimRackGroupsReadOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for dcim_rack-groups_read: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  DcimRackGroupsUpdate dcim rack groups update API
*/
func (a *Client) DcimRackGroupsUpdate(params *DcimRackGroupsUpdateParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*DcimRackGroupsUpdateOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDcimRackGroupsUpdateParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "dcim_rack-groups_update",
		Method:             "PUT",
		PathPattern:        "/dcim/rack-groups/{id}/",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &DcimRackGroupsUpdateReader{formats: a.formats},
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
	success, ok := result.(*DcimRackGroupsUpdateOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for dcim_rack-groups_update: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  DcimRackReservationsBulkDelete dcim rack reservations bulk delete API
*/
func (a *Client) DcimRackReservationsBulkDelete(params *DcimRackReservationsBulkDeleteParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*DcimRackReservationsBulkDeleteNoContent, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDcimRackReservationsBulkDeleteParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "dcim_rack-reservations_bulk_delete",
		Method:             "DELETE",
		PathPattern:        "/dcim/rack-reservations/",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &DcimRackReservationsBulkDeleteReader{formats: a.formats},
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
	success, ok := result.(*DcimRackReservationsBulkDeleteNoContent)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for dcim_rack-reservations_bulk_delete: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  DcimRackReservationsBulkPartialUpdate dcim rack reservations bulk partial update API
*/
func (a *Client) DcimRackReservationsBulkPartialUpdate(params *DcimRackReservationsBulkPartialUpdateParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*DcimRackReservationsBulkPartialUpdateOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDcimRackReservationsBulkPartialUpdateParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "dcim_rack-reservations_bulk_partial_update",
		Method:             "PATCH",
		PathPattern:        "/dcim/rack-reservations/",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &DcimRackReservationsBulkPartialUpdateReader{formats: a.formats},
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
	success, ok := result.(*DcimRackReservationsBulkPartialUpdateOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for dcim_rack-reservations_bulk_partial_update: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  DcimRackReservationsBulkUpdate dcim rack reservations bulk update API
*/
func (a *Client) DcimRackReservationsBulkUpdate(params *DcimRackReservationsBulkUpdateParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*DcimRackReservationsBulkUpdateOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDcimRackReservationsBulkUpdateParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "dcim_rack-reservations_bulk_update",
		Method:             "PUT",
		PathPattern:        "/dcim/rack-reservations/",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &DcimRackReservationsBulkUpdateReader{formats: a.formats},
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
	success, ok := result.(*DcimRackReservationsBulkUpdateOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for dcim_rack-reservations_bulk_update: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  DcimRackReservationsCreate dcim rack reservations create API
*/
func (a *Client) DcimRackReservationsCreate(params *DcimRackReservationsCreateParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*DcimRackReservationsCreateCreated, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDcimRackReservationsCreateParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "dcim_rack-reservations_create",
		Method:             "POST",
		PathPattern:        "/dcim/rack-reservations/",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &DcimRackReservationsCreateReader{formats: a.formats},
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
	success, ok := result.(*DcimRackReservationsCreateCreated)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for dcim_rack-reservations_create: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  DcimRackReservationsDelete dcim rack reservations delete API
*/
func (a *Client) DcimRackReservationsDelete(params *DcimRackReservationsDeleteParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*DcimRackReservationsDeleteNoContent, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDcimRackReservationsDeleteParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "dcim_rack-reservations_delete",
		Method:             "DELETE",
		PathPattern:        "/dcim/rack-reservations/{id}/",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &DcimRackReservationsDeleteReader{formats: a.formats},
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
	success, ok := result.(*DcimRackReservationsDeleteNoContent)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for dcim_rack-reservations_delete: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  DcimRackReservationsList dcim rack reservations list API
*/
func (a *Client) DcimRackReservationsList(params *DcimRackReservationsListParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*DcimRackReservationsListOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDcimRackReservationsListParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "dcim_rack-reservations_list",
		Method:             "GET",
		PathPattern:        "/dcim/rack-reservations/",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &DcimRackReservationsListReader{formats: a.formats},
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
	success, ok := result.(*DcimRackReservationsListOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for dcim_rack-reservations_list: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  DcimRackReservationsPartialUpdate dcim rack reservations partial update API
*/
func (a *Client) DcimRackReservationsPartialUpdate(params *DcimRackReservationsPartialUpdateParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*DcimRackReservationsPartialUpdateOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDcimRackReservationsPartialUpdateParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "dcim_rack-reservations_partial_update",
		Method:             "PATCH",
		PathPattern:        "/dcim/rack-reservations/{id}/",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &DcimRackReservationsPartialUpdateReader{formats: a.formats},
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
	success, ok := result.(*DcimRackReservationsPartialUpdateOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for dcim_rack-reservations_partial_update: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  DcimRackReservationsRead dcim rack reservations read API
*/
func (a *Client) DcimRackReservationsRead(params *DcimRackReservationsReadParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*DcimRackReservationsReadOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDcimRackReservationsReadParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "dcim_rack-reservations_read",
		Method:             "GET",
		PathPattern:        "/dcim/rack-reservations/{id}/",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &DcimRackReservationsReadReader{formats: a.formats},
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
	success, ok := result.(*DcimRackReservationsReadOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for dcim_rack-reservations_read: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  DcimRackReservationsUpdate dcim rack reservations update API
*/
func (a *Client) DcimRackReservationsUpdate(params *DcimRackReservationsUpdateParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*DcimRackReservationsUpdateOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDcimRackReservationsUpdateParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "dcim_rack-reservations_update",
		Method:             "PUT",
		PathPattern:        "/dcim/rack-reservations/{id}/",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &DcimRackReservationsUpdateReader{formats: a.formats},
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
	success, ok := result.(*DcimRackReservationsUpdateOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for dcim_rack-reservations_update: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  DcimRackRolesBulkDelete dcim rack roles bulk delete API
*/
func (a *Client) DcimRackRolesBulkDelete(params *DcimRackRolesBulkDeleteParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*DcimRackRolesBulkDeleteNoContent, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDcimRackRolesBulkDeleteParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "dcim_rack-roles_bulk_delete",
		Method:             "DELETE",
		PathPattern:        "/dcim/rack-roles/",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &DcimRackRolesBulkDeleteReader{formats: a.formats},
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
	success, ok := result.(*DcimRackRolesBulkDeleteNoContent)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for dcim_rack-roles_bulk_delete: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  DcimRackRolesBulkPartialUpdate dcim rack roles bulk partial update API
*/
func (a *Client) DcimRackRolesBulkPartialUpdate(params *DcimRackRolesBulkPartialUpdateParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*DcimRackRolesBulkPartialUpdateOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDcimRackRolesBulkPartialUpdateParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "dcim_rack-roles_bulk_partial_update",
		Method:             "PATCH",
		PathPattern:        "/dcim/rack-roles/",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &DcimRackRolesBulkPartialUpdateReader{formats: a.formats},
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
	success, ok := result.(*DcimRackRolesBulkPartialUpdateOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for dcim_rack-roles_bulk_partial_update: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  DcimRackRolesBulkUpdate dcim rack roles bulk update API
*/
func (a *Client) DcimRackRolesBulkUpdate(params *DcimRackRolesBulkUpdateParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*DcimRackRolesBulkUpdateOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDcimRackRolesBulkUpdateParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "dcim_rack-roles_bulk_update",
		Method:             "PUT",
		PathPattern:        "/dcim/rack-roles/",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &DcimRackRolesBulkUpdateReader{formats: a.formats},
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
	success, ok := result.(*DcimRackRolesBulkUpdateOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for dcim_rack-roles_bulk_update: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  DcimRackRolesCreate dcim rack roles create API
*/
func (a *Client) DcimRackRolesCreate(params *DcimRackRolesCreateParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*DcimRackRolesCreateCreated, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDcimRackRolesCreateParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "dcim_rack-roles_create",
		Method:             "POST",
		PathPattern:        "/dcim/rack-roles/",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &DcimRackRolesCreateReader{formats: a.formats},
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
	success, ok := result.(*DcimRackRolesCreateCreated)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for dcim_rack-roles_create: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  DcimRackRolesDelete dcim rack roles delete API
*/
func (a *Client) DcimRackRolesDelete(params *DcimRackRolesDeleteParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*DcimRackRolesDeleteNoContent, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDcimRackRolesDeleteParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "dcim_rack-roles_delete",
		Method:             "DELETE",
		PathPattern:        "/dcim/rack-roles/{id}/",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &DcimRackRolesDeleteReader{formats: a.formats},
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
	success, ok := result.(*DcimRackRolesDeleteNoContent)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for dcim_rack-roles_delete: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  DcimRackRolesList dcim rack roles list API
*/
func (a *Client) DcimRackRolesList(params *DcimRackRolesListParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*DcimRackRolesListOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDcimRackRolesListParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "dcim_rack-roles_list",
		Method:             "GET",
		PathPattern:        "/dcim/rack-roles/",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &DcimRackRolesListReader{formats: a.formats},
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
	success, ok := result.(*DcimRackRolesListOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for dcim_rack-roles_list: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  DcimRackRolesPartialUpdate dcim rack roles partial update API
*/
func (a *Client) DcimRackRolesPartialUpdate(params *DcimRackRolesPartialUpdateParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*DcimRackRolesPartialUpdateOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDcimRackRolesPartialUpdateParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "dcim_rack-roles_partial_update",
		Method:             "PATCH",
		PathPattern:        "/dcim/rack-roles/{id}/",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &DcimRackRolesPartialUpdateReader{formats: a.formats},
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
	success, ok := result.(*DcimRackRolesPartialUpdateOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for dcim_rack-roles_partial_update: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  DcimRackRolesRead dcim rack roles read API
*/
func (a *Client) DcimRackRolesRead(params *DcimRackRolesReadParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*DcimRackRolesReadOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDcimRackRolesReadParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "dcim_rack-roles_read",
		Method:             "GET",
		PathPattern:        "/dcim/rack-roles/{id}/",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &DcimRackRolesReadReader{formats: a.formats},
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
	success, ok := result.(*DcimRackRolesReadOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for dcim_rack-roles_read: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  DcimRackRolesUpdate dcim rack roles update API
*/
func (a *Client) DcimRackRolesUpdate(params *DcimRackRolesUpdateParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*DcimRackRolesUpdateOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDcimRackRolesUpdateParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "dcim_rack-roles_update",
		Method:             "PUT",
		PathPattern:        "/dcim/rack-roles/{id}/",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &DcimRackRolesUpdateReader{formats: a.formats},
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
	success, ok := result.(*DcimRackRolesUpdateOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for dcim_rack-roles_update: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  DcimRacksBulkDelete dcim racks bulk delete API
*/
func (a *Client) DcimRacksBulkDelete(params *DcimRacksBulkDeleteParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*DcimRacksBulkDeleteNoContent, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDcimRacksBulkDeleteParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "dcim_racks_bulk_delete",
		Method:             "DELETE",
		PathPattern:        "/dcim/racks/",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &DcimRacksBulkDeleteReader{formats: a.formats},
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
	success, ok := result.(*DcimRacksBulkDeleteNoContent)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for dcim_racks_bulk_delete: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  DcimRacksBulkPartialUpdate dcim racks bulk partial update API
*/
func (a *Client) DcimRacksBulkPartialUpdate(params *DcimRacksBulkPartialUpdateParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*DcimRacksBulkPartialUpdateOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDcimRacksBulkPartialUpdateParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "dcim_racks_bulk_partial_update",
		Method:             "PATCH",
		PathPattern:        "/dcim/racks/",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &DcimRacksBulkPartialUpdateReader{formats: a.formats},
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
	success, ok := result.(*DcimRacksBulkPartialUpdateOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for dcim_racks_bulk_partial_update: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  DcimRacksBulkUpdate dcim racks bulk update API
*/
func (a *Client) DcimRacksBulkUpdate(params *DcimRacksBulkUpdateParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*DcimRacksBulkUpdateOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDcimRacksBulkUpdateParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "dcim_racks_bulk_update",
		Method:             "PUT",
		PathPattern:        "/dcim/racks/",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &DcimRacksBulkUpdateReader{formats: a.formats},
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
	success, ok := result.(*DcimRacksBulkUpdateOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for dcim_racks_bulk_update: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  DcimRacksCreate dcim racks create API
*/
func (a *Client) DcimRacksCreate(params *DcimRacksCreateParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*DcimRacksCreateCreated, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDcimRacksCreateParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "dcim_racks_create",
		Method:             "POST",
		PathPattern:        "/dcim/racks/",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &DcimRacksCreateReader{formats: a.formats},
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
	success, ok := result.(*DcimRacksCreateCreated)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for dcim_racks_create: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  DcimRacksDelete dcim racks delete API
*/
func (a *Client) DcimRacksDelete(params *DcimRacksDeleteParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*DcimRacksDeleteNoContent, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDcimRacksDeleteParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "dcim_racks_delete",
		Method:             "DELETE",
		PathPattern:        "/dcim/racks/{id}/",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &DcimRacksDeleteReader{formats: a.formats},
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
	success, ok := result.(*DcimRacksDeleteNoContent)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for dcim_racks_delete: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  DcimRacksElevationRead Rack elevation representing the list of rack units. Also supports rendering the elevation as an SVG.
*/
func (a *Client) DcimRacksElevationRead(params *DcimRacksElevationReadParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*DcimRacksElevationReadOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDcimRacksElevationReadParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "dcim_racks_elevation_read",
		Method:             "GET",
		PathPattern:        "/dcim/racks/{id}/elevation/",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &DcimRacksElevationReadReader{formats: a.formats},
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
	success, ok := result.(*DcimRacksElevationReadOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for dcim_racks_elevation_read: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  DcimRacksList dcim racks list API
*/
func (a *Client) DcimRacksList(params *DcimRacksListParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*DcimRacksListOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDcimRacksListParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "dcim_racks_list",
		Method:             "GET",
		PathPattern:        "/dcim/racks/",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &DcimRacksListReader{formats: a.formats},
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
	success, ok := result.(*DcimRacksListOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for dcim_racks_list: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  DcimRacksPartialUpdate dcim racks partial update API
*/
func (a *Client) DcimRacksPartialUpdate(params *DcimRacksPartialUpdateParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*DcimRacksPartialUpdateOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDcimRacksPartialUpdateParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "dcim_racks_partial_update",
		Method:             "PATCH",
		PathPattern:        "/dcim/racks/{id}/",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &DcimRacksPartialUpdateReader{formats: a.formats},
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
	success, ok := result.(*DcimRacksPartialUpdateOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for dcim_racks_partial_update: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  DcimRacksRead dcim racks read API
*/
func (a *Client) DcimRacksRead(params *DcimRacksReadParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*DcimRacksReadOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDcimRacksReadParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "dcim_racks_read",
		Method:             "GET",
		PathPattern:        "/dcim/racks/{id}/",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &DcimRacksReadReader{formats: a.formats},
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
	success, ok := result.(*DcimRacksReadOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for dcim_racks_read: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  DcimRacksUpdate dcim racks update API
*/
func (a *Client) DcimRacksUpdate(params *DcimRacksUpdateParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*DcimRacksUpdateOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDcimRacksUpdateParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "dcim_racks_update",
		Method:             "PUT",
		PathPattern:        "/dcim/racks/{id}/",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &DcimRacksUpdateReader{formats: a.formats},
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
	success, ok := result.(*DcimRacksUpdateOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for dcim_racks_update: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  DcimRearPortTemplatesBulkDelete dcim rear port templates bulk delete API
*/
func (a *Client) DcimRearPortTemplatesBulkDelete(params *DcimRearPortTemplatesBulkDeleteParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*DcimRearPortTemplatesBulkDeleteNoContent, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDcimRearPortTemplatesBulkDeleteParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "dcim_rear-port-templates_bulk_delete",
		Method:             "DELETE",
		PathPattern:        "/dcim/rear-port-templates/",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &DcimRearPortTemplatesBulkDeleteReader{formats: a.formats},
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
	success, ok := result.(*DcimRearPortTemplatesBulkDeleteNoContent)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for dcim_rear-port-templates_bulk_delete: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  DcimRearPortTemplatesBulkPartialUpdate dcim rear port templates bulk partial update API
*/
func (a *Client) DcimRearPortTemplatesBulkPartialUpdate(params *DcimRearPortTemplatesBulkPartialUpdateParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*DcimRearPortTemplatesBulkPartialUpdateOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDcimRearPortTemplatesBulkPartialUpdateParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "dcim_rear-port-templates_bulk_partial_update",
		Method:             "PATCH",
		PathPattern:        "/dcim/rear-port-templates/",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &DcimRearPortTemplatesBulkPartialUpdateReader{formats: a.formats},
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
	success, ok := result.(*DcimRearPortTemplatesBulkPartialUpdateOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for dcim_rear-port-templates_bulk_partial_update: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  DcimRearPortTemplatesBulkUpdate dcim rear port templates bulk update API
*/
func (a *Client) DcimRearPortTemplatesBulkUpdate(params *DcimRearPortTemplatesBulkUpdateParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*DcimRearPortTemplatesBulkUpdateOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDcimRearPortTemplatesBulkUpdateParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "dcim_rear-port-templates_bulk_update",
		Method:             "PUT",
		PathPattern:        "/dcim/rear-port-templates/",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &DcimRearPortTemplatesBulkUpdateReader{formats: a.formats},
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
	success, ok := result.(*DcimRearPortTemplatesBulkUpdateOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for dcim_rear-port-templates_bulk_update: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  DcimRearPortTemplatesCreate dcim rear port templates create API
*/
func (a *Client) DcimRearPortTemplatesCreate(params *DcimRearPortTemplatesCreateParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*DcimRearPortTemplatesCreateCreated, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDcimRearPortTemplatesCreateParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "dcim_rear-port-templates_create",
		Method:             "POST",
		PathPattern:        "/dcim/rear-port-templates/",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &DcimRearPortTemplatesCreateReader{formats: a.formats},
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
	success, ok := result.(*DcimRearPortTemplatesCreateCreated)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for dcim_rear-port-templates_create: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  DcimRearPortTemplatesDelete dcim rear port templates delete API
*/
func (a *Client) DcimRearPortTemplatesDelete(params *DcimRearPortTemplatesDeleteParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*DcimRearPortTemplatesDeleteNoContent, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDcimRearPortTemplatesDeleteParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "dcim_rear-port-templates_delete",
		Method:             "DELETE",
		PathPattern:        "/dcim/rear-port-templates/{id}/",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &DcimRearPortTemplatesDeleteReader{formats: a.formats},
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
	success, ok := result.(*DcimRearPortTemplatesDeleteNoContent)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for dcim_rear-port-templates_delete: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  DcimRearPortTemplatesList dcim rear port templates list API
*/
func (a *Client) DcimRearPortTemplatesList(params *DcimRearPortTemplatesListParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*DcimRearPortTemplatesListOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDcimRearPortTemplatesListParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "dcim_rear-port-templates_list",
		Method:             "GET",
		PathPattern:        "/dcim/rear-port-templates/",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &DcimRearPortTemplatesListReader{formats: a.formats},
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
	success, ok := result.(*DcimRearPortTemplatesListOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for dcim_rear-port-templates_list: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  DcimRearPortTemplatesPartialUpdate dcim rear port templates partial update API
*/
func (a *Client) DcimRearPortTemplatesPartialUpdate(params *DcimRearPortTemplatesPartialUpdateParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*DcimRearPortTemplatesPartialUpdateOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDcimRearPortTemplatesPartialUpdateParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "dcim_rear-port-templates_partial_update",
		Method:             "PATCH",
		PathPattern:        "/dcim/rear-port-templates/{id}/",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &DcimRearPortTemplatesPartialUpdateReader{formats: a.formats},
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
	success, ok := result.(*DcimRearPortTemplatesPartialUpdateOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for dcim_rear-port-templates_partial_update: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  DcimRearPortTemplatesRead dcim rear port templates read API
*/
func (a *Client) DcimRearPortTemplatesRead(params *DcimRearPortTemplatesReadParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*DcimRearPortTemplatesReadOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDcimRearPortTemplatesReadParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "dcim_rear-port-templates_read",
		Method:             "GET",
		PathPattern:        "/dcim/rear-port-templates/{id}/",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &DcimRearPortTemplatesReadReader{formats: a.formats},
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
	success, ok := result.(*DcimRearPortTemplatesReadOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for dcim_rear-port-templates_read: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  DcimRearPortTemplatesUpdate dcim rear port templates update API
*/
func (a *Client) DcimRearPortTemplatesUpdate(params *DcimRearPortTemplatesUpdateParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*DcimRearPortTemplatesUpdateOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDcimRearPortTemplatesUpdateParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "dcim_rear-port-templates_update",
		Method:             "PUT",
		PathPattern:        "/dcim/rear-port-templates/{id}/",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &DcimRearPortTemplatesUpdateReader{formats: a.formats},
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
	success, ok := result.(*DcimRearPortTemplatesUpdateOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for dcim_rear-port-templates_update: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  DcimRearPortsBulkDelete dcim rear ports bulk delete API
*/
func (a *Client) DcimRearPortsBulkDelete(params *DcimRearPortsBulkDeleteParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*DcimRearPortsBulkDeleteNoContent, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDcimRearPortsBulkDeleteParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "dcim_rear-ports_bulk_delete",
		Method:             "DELETE",
		PathPattern:        "/dcim/rear-ports/",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &DcimRearPortsBulkDeleteReader{formats: a.formats},
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
	success, ok := result.(*DcimRearPortsBulkDeleteNoContent)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for dcim_rear-ports_bulk_delete: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  DcimRearPortsBulkPartialUpdate dcim rear ports bulk partial update API
*/
func (a *Client) DcimRearPortsBulkPartialUpdate(params *DcimRearPortsBulkPartialUpdateParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*DcimRearPortsBulkPartialUpdateOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDcimRearPortsBulkPartialUpdateParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "dcim_rear-ports_bulk_partial_update",
		Method:             "PATCH",
		PathPattern:        "/dcim/rear-ports/",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &DcimRearPortsBulkPartialUpdateReader{formats: a.formats},
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
	success, ok := result.(*DcimRearPortsBulkPartialUpdateOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for dcim_rear-ports_bulk_partial_update: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  DcimRearPortsBulkUpdate dcim rear ports bulk update API
*/
func (a *Client) DcimRearPortsBulkUpdate(params *DcimRearPortsBulkUpdateParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*DcimRearPortsBulkUpdateOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDcimRearPortsBulkUpdateParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "dcim_rear-ports_bulk_update",
		Method:             "PUT",
		PathPattern:        "/dcim/rear-ports/",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &DcimRearPortsBulkUpdateReader{formats: a.formats},
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
	success, ok := result.(*DcimRearPortsBulkUpdateOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for dcim_rear-ports_bulk_update: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  DcimRearPortsCreate dcim rear ports create API
*/
func (a *Client) DcimRearPortsCreate(params *DcimRearPortsCreateParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*DcimRearPortsCreateCreated, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDcimRearPortsCreateParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "dcim_rear-ports_create",
		Method:             "POST",
		PathPattern:        "/dcim/rear-ports/",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &DcimRearPortsCreateReader{formats: a.formats},
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
	success, ok := result.(*DcimRearPortsCreateCreated)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for dcim_rear-ports_create: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  DcimRearPortsDelete dcim rear ports delete API
*/
func (a *Client) DcimRearPortsDelete(params *DcimRearPortsDeleteParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*DcimRearPortsDeleteNoContent, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDcimRearPortsDeleteParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "dcim_rear-ports_delete",
		Method:             "DELETE",
		PathPattern:        "/dcim/rear-ports/{id}/",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &DcimRearPortsDeleteReader{formats: a.formats},
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
	success, ok := result.(*DcimRearPortsDeleteNoContent)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for dcim_rear-ports_delete: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  DcimRearPortsList dcim rear ports list API
*/
func (a *Client) DcimRearPortsList(params *DcimRearPortsListParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*DcimRearPortsListOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDcimRearPortsListParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "dcim_rear-ports_list",
		Method:             "GET",
		PathPattern:        "/dcim/rear-ports/",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &DcimRearPortsListReader{formats: a.formats},
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
	success, ok := result.(*DcimRearPortsListOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for dcim_rear-ports_list: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  DcimRearPortsPartialUpdate dcim rear ports partial update API
*/
func (a *Client) DcimRearPortsPartialUpdate(params *DcimRearPortsPartialUpdateParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*DcimRearPortsPartialUpdateOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDcimRearPortsPartialUpdateParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "dcim_rear-ports_partial_update",
		Method:             "PATCH",
		PathPattern:        "/dcim/rear-ports/{id}/",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &DcimRearPortsPartialUpdateReader{formats: a.formats},
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
	success, ok := result.(*DcimRearPortsPartialUpdateOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for dcim_rear-ports_partial_update: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  DcimRearPortsPaths Return all CablePaths which traverse a given pass-through port.
*/
func (a *Client) DcimRearPortsPaths(params *DcimRearPortsPathsParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*DcimRearPortsPathsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDcimRearPortsPathsParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "dcim_rear-ports_paths",
		Method:             "GET",
		PathPattern:        "/dcim/rear-ports/{id}/paths/",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &DcimRearPortsPathsReader{formats: a.formats},
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
	success, ok := result.(*DcimRearPortsPathsOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for dcim_rear-ports_paths: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  DcimRearPortsRead dcim rear ports read API
*/
func (a *Client) DcimRearPortsRead(params *DcimRearPortsReadParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*DcimRearPortsReadOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDcimRearPortsReadParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "dcim_rear-ports_read",
		Method:             "GET",
		PathPattern:        "/dcim/rear-ports/{id}/",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &DcimRearPortsReadReader{formats: a.formats},
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
	success, ok := result.(*DcimRearPortsReadOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for dcim_rear-ports_read: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  DcimRearPortsUpdate dcim rear ports update API
*/
func (a *Client) DcimRearPortsUpdate(params *DcimRearPortsUpdateParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*DcimRearPortsUpdateOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDcimRearPortsUpdateParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "dcim_rear-ports_update",
		Method:             "PUT",
		PathPattern:        "/dcim/rear-ports/{id}/",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &DcimRearPortsUpdateReader{formats: a.formats},
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
	success, ok := result.(*DcimRearPortsUpdateOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for dcim_rear-ports_update: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  DcimRegionsBulkDelete dcim regions bulk delete API
*/
func (a *Client) DcimRegionsBulkDelete(params *DcimRegionsBulkDeleteParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*DcimRegionsBulkDeleteNoContent, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDcimRegionsBulkDeleteParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "dcim_regions_bulk_delete",
		Method:             "DELETE",
		PathPattern:        "/dcim/regions/",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &DcimRegionsBulkDeleteReader{formats: a.formats},
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
	success, ok := result.(*DcimRegionsBulkDeleteNoContent)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for dcim_regions_bulk_delete: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  DcimRegionsBulkPartialUpdate dcim regions bulk partial update API
*/
func (a *Client) DcimRegionsBulkPartialUpdate(params *DcimRegionsBulkPartialUpdateParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*DcimRegionsBulkPartialUpdateOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDcimRegionsBulkPartialUpdateParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "dcim_regions_bulk_partial_update",
		Method:             "PATCH",
		PathPattern:        "/dcim/regions/",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &DcimRegionsBulkPartialUpdateReader{formats: a.formats},
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
	success, ok := result.(*DcimRegionsBulkPartialUpdateOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for dcim_regions_bulk_partial_update: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  DcimRegionsBulkUpdate dcim regions bulk update API
*/
func (a *Client) DcimRegionsBulkUpdate(params *DcimRegionsBulkUpdateParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*DcimRegionsBulkUpdateOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDcimRegionsBulkUpdateParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "dcim_regions_bulk_update",
		Method:             "PUT",
		PathPattern:        "/dcim/regions/",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &DcimRegionsBulkUpdateReader{formats: a.formats},
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
	success, ok := result.(*DcimRegionsBulkUpdateOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for dcim_regions_bulk_update: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  DcimRegionsCreate dcim regions create API
*/
func (a *Client) DcimRegionsCreate(params *DcimRegionsCreateParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*DcimRegionsCreateCreated, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDcimRegionsCreateParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "dcim_regions_create",
		Method:             "POST",
		PathPattern:        "/dcim/regions/",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &DcimRegionsCreateReader{formats: a.formats},
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
	success, ok := result.(*DcimRegionsCreateCreated)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for dcim_regions_create: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  DcimRegionsDelete dcim regions delete API
*/
func (a *Client) DcimRegionsDelete(params *DcimRegionsDeleteParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*DcimRegionsDeleteNoContent, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDcimRegionsDeleteParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "dcim_regions_delete",
		Method:             "DELETE",
		PathPattern:        "/dcim/regions/{id}/",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &DcimRegionsDeleteReader{formats: a.formats},
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
	success, ok := result.(*DcimRegionsDeleteNoContent)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for dcim_regions_delete: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  DcimRegionsList dcim regions list API
*/
func (a *Client) DcimRegionsList(params *DcimRegionsListParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*DcimRegionsListOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDcimRegionsListParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "dcim_regions_list",
		Method:             "GET",
		PathPattern:        "/dcim/regions/",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &DcimRegionsListReader{formats: a.formats},
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
	success, ok := result.(*DcimRegionsListOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for dcim_regions_list: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  DcimRegionsPartialUpdate dcim regions partial update API
*/
func (a *Client) DcimRegionsPartialUpdate(params *DcimRegionsPartialUpdateParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*DcimRegionsPartialUpdateOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDcimRegionsPartialUpdateParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "dcim_regions_partial_update",
		Method:             "PATCH",
		PathPattern:        "/dcim/regions/{id}/",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &DcimRegionsPartialUpdateReader{formats: a.formats},
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
	success, ok := result.(*DcimRegionsPartialUpdateOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for dcim_regions_partial_update: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  DcimRegionsRead dcim regions read API
*/
func (a *Client) DcimRegionsRead(params *DcimRegionsReadParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*DcimRegionsReadOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDcimRegionsReadParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "dcim_regions_read",
		Method:             "GET",
		PathPattern:        "/dcim/regions/{id}/",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &DcimRegionsReadReader{formats: a.formats},
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
	success, ok := result.(*DcimRegionsReadOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for dcim_regions_read: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  DcimRegionsUpdate dcim regions update API
*/
func (a *Client) DcimRegionsUpdate(params *DcimRegionsUpdateParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*DcimRegionsUpdateOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDcimRegionsUpdateParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "dcim_regions_update",
		Method:             "PUT",
		PathPattern:        "/dcim/regions/{id}/",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &DcimRegionsUpdateReader{formats: a.formats},
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
	success, ok := result.(*DcimRegionsUpdateOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for dcim_regions_update: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  DcimSitesBulkDelete dcim sites bulk delete API
*/
func (a *Client) DcimSitesBulkDelete(params *DcimSitesBulkDeleteParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*DcimSitesBulkDeleteNoContent, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDcimSitesBulkDeleteParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "dcim_sites_bulk_delete",
		Method:             "DELETE",
		PathPattern:        "/dcim/sites/",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &DcimSitesBulkDeleteReader{formats: a.formats},
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
	success, ok := result.(*DcimSitesBulkDeleteNoContent)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for dcim_sites_bulk_delete: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  DcimSitesBulkPartialUpdate dcim sites bulk partial update API
*/
func (a *Client) DcimSitesBulkPartialUpdate(params *DcimSitesBulkPartialUpdateParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*DcimSitesBulkPartialUpdateOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDcimSitesBulkPartialUpdateParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "dcim_sites_bulk_partial_update",
		Method:             "PATCH",
		PathPattern:        "/dcim/sites/",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &DcimSitesBulkPartialUpdateReader{formats: a.formats},
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
	success, ok := result.(*DcimSitesBulkPartialUpdateOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for dcim_sites_bulk_partial_update: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  DcimSitesBulkUpdate dcim sites bulk update API
*/
func (a *Client) DcimSitesBulkUpdate(params *DcimSitesBulkUpdateParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*DcimSitesBulkUpdateOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDcimSitesBulkUpdateParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "dcim_sites_bulk_update",
		Method:             "PUT",
		PathPattern:        "/dcim/sites/",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &DcimSitesBulkUpdateReader{formats: a.formats},
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
	success, ok := result.(*DcimSitesBulkUpdateOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for dcim_sites_bulk_update: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  DcimSitesCreate dcim sites create API
*/
func (a *Client) DcimSitesCreate(params *DcimSitesCreateParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*DcimSitesCreateCreated, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDcimSitesCreateParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "dcim_sites_create",
		Method:             "POST",
		PathPattern:        "/dcim/sites/",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &DcimSitesCreateReader{formats: a.formats},
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
	success, ok := result.(*DcimSitesCreateCreated)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for dcim_sites_create: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  DcimSitesDelete dcim sites delete API
*/
func (a *Client) DcimSitesDelete(params *DcimSitesDeleteParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*DcimSitesDeleteNoContent, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDcimSitesDeleteParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "dcim_sites_delete",
		Method:             "DELETE",
		PathPattern:        "/dcim/sites/{id}/",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &DcimSitesDeleteReader{formats: a.formats},
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
	success, ok := result.(*DcimSitesDeleteNoContent)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for dcim_sites_delete: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  DcimSitesList dcim sites list API
*/
func (a *Client) DcimSitesList(params *DcimSitesListParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*DcimSitesListOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDcimSitesListParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "dcim_sites_list",
		Method:             "GET",
		PathPattern:        "/dcim/sites/",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &DcimSitesListReader{formats: a.formats},
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
	success, ok := result.(*DcimSitesListOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for dcim_sites_list: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  DcimSitesPartialUpdate dcim sites partial update API
*/
func (a *Client) DcimSitesPartialUpdate(params *DcimSitesPartialUpdateParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*DcimSitesPartialUpdateOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDcimSitesPartialUpdateParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "dcim_sites_partial_update",
		Method:             "PATCH",
		PathPattern:        "/dcim/sites/{id}/",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &DcimSitesPartialUpdateReader{formats: a.formats},
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
	success, ok := result.(*DcimSitesPartialUpdateOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for dcim_sites_partial_update: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  DcimSitesRead dcim sites read API
*/
func (a *Client) DcimSitesRead(params *DcimSitesReadParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*DcimSitesReadOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDcimSitesReadParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "dcim_sites_read",
		Method:             "GET",
		PathPattern:        "/dcim/sites/{id}/",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &DcimSitesReadReader{formats: a.formats},
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
	success, ok := result.(*DcimSitesReadOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for dcim_sites_read: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  DcimSitesUpdate dcim sites update API
*/
func (a *Client) DcimSitesUpdate(params *DcimSitesUpdateParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*DcimSitesUpdateOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDcimSitesUpdateParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "dcim_sites_update",
		Method:             "PUT",
		PathPattern:        "/dcim/sites/{id}/",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &DcimSitesUpdateReader{formats: a.formats},
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
	success, ok := result.(*DcimSitesUpdateOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for dcim_sites_update: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  DcimVirtualChassisBulkDelete dcim virtual chassis bulk delete API
*/
func (a *Client) DcimVirtualChassisBulkDelete(params *DcimVirtualChassisBulkDeleteParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*DcimVirtualChassisBulkDeleteNoContent, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDcimVirtualChassisBulkDeleteParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "dcim_virtual-chassis_bulk_delete",
		Method:             "DELETE",
		PathPattern:        "/dcim/virtual-chassis/",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &DcimVirtualChassisBulkDeleteReader{formats: a.formats},
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
	success, ok := result.(*DcimVirtualChassisBulkDeleteNoContent)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for dcim_virtual-chassis_bulk_delete: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  DcimVirtualChassisBulkPartialUpdate dcim virtual chassis bulk partial update API
*/
func (a *Client) DcimVirtualChassisBulkPartialUpdate(params *DcimVirtualChassisBulkPartialUpdateParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*DcimVirtualChassisBulkPartialUpdateOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDcimVirtualChassisBulkPartialUpdateParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "dcim_virtual-chassis_bulk_partial_update",
		Method:             "PATCH",
		PathPattern:        "/dcim/virtual-chassis/",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &DcimVirtualChassisBulkPartialUpdateReader{formats: a.formats},
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
	success, ok := result.(*DcimVirtualChassisBulkPartialUpdateOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for dcim_virtual-chassis_bulk_partial_update: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  DcimVirtualChassisBulkUpdate dcim virtual chassis bulk update API
*/
func (a *Client) DcimVirtualChassisBulkUpdate(params *DcimVirtualChassisBulkUpdateParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*DcimVirtualChassisBulkUpdateOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDcimVirtualChassisBulkUpdateParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "dcim_virtual-chassis_bulk_update",
		Method:             "PUT",
		PathPattern:        "/dcim/virtual-chassis/",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &DcimVirtualChassisBulkUpdateReader{formats: a.formats},
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
	success, ok := result.(*DcimVirtualChassisBulkUpdateOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for dcim_virtual-chassis_bulk_update: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  DcimVirtualChassisCreate dcim virtual chassis create API
*/
func (a *Client) DcimVirtualChassisCreate(params *DcimVirtualChassisCreateParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*DcimVirtualChassisCreateCreated, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDcimVirtualChassisCreateParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "dcim_virtual-chassis_create",
		Method:             "POST",
		PathPattern:        "/dcim/virtual-chassis/",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &DcimVirtualChassisCreateReader{formats: a.formats},
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
	success, ok := result.(*DcimVirtualChassisCreateCreated)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for dcim_virtual-chassis_create: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  DcimVirtualChassisDelete dcim virtual chassis delete API
*/
func (a *Client) DcimVirtualChassisDelete(params *DcimVirtualChassisDeleteParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*DcimVirtualChassisDeleteNoContent, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDcimVirtualChassisDeleteParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "dcim_virtual-chassis_delete",
		Method:             "DELETE",
		PathPattern:        "/dcim/virtual-chassis/{id}/",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &DcimVirtualChassisDeleteReader{formats: a.formats},
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
	success, ok := result.(*DcimVirtualChassisDeleteNoContent)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for dcim_virtual-chassis_delete: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  DcimVirtualChassisList dcim virtual chassis list API
*/
func (a *Client) DcimVirtualChassisList(params *DcimVirtualChassisListParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*DcimVirtualChassisListOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDcimVirtualChassisListParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "dcim_virtual-chassis_list",
		Method:             "GET",
		PathPattern:        "/dcim/virtual-chassis/",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &DcimVirtualChassisListReader{formats: a.formats},
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
	success, ok := result.(*DcimVirtualChassisListOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for dcim_virtual-chassis_list: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  DcimVirtualChassisPartialUpdate dcim virtual chassis partial update API
*/
func (a *Client) DcimVirtualChassisPartialUpdate(params *DcimVirtualChassisPartialUpdateParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*DcimVirtualChassisPartialUpdateOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDcimVirtualChassisPartialUpdateParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "dcim_virtual-chassis_partial_update",
		Method:             "PATCH",
		PathPattern:        "/dcim/virtual-chassis/{id}/",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &DcimVirtualChassisPartialUpdateReader{formats: a.formats},
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
	success, ok := result.(*DcimVirtualChassisPartialUpdateOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for dcim_virtual-chassis_partial_update: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  DcimVirtualChassisRead dcim virtual chassis read API
*/
func (a *Client) DcimVirtualChassisRead(params *DcimVirtualChassisReadParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*DcimVirtualChassisReadOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDcimVirtualChassisReadParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "dcim_virtual-chassis_read",
		Method:             "GET",
		PathPattern:        "/dcim/virtual-chassis/{id}/",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &DcimVirtualChassisReadReader{formats: a.formats},
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
	success, ok := result.(*DcimVirtualChassisReadOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for dcim_virtual-chassis_read: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  DcimVirtualChassisUpdate dcim virtual chassis update API
*/
func (a *Client) DcimVirtualChassisUpdate(params *DcimVirtualChassisUpdateParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*DcimVirtualChassisUpdateOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDcimVirtualChassisUpdateParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "dcim_virtual-chassis_update",
		Method:             "PUT",
		PathPattern:        "/dcim/virtual-chassis/{id}/",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &DcimVirtualChassisUpdateReader{formats: a.formats},
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
	success, ok := result.(*DcimVirtualChassisUpdateOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for dcim_virtual-chassis_update: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
