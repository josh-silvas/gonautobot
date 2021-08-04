package virtualization

import (
	"context"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// NewVirtualizationVirtualMachinesListParams creates a new VirtualizationVirtualMachinesListParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewVirtualizationVirtualMachinesListParams() *VirtualizationVirtualMachinesListParams {
	return &VirtualizationVirtualMachinesListParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewVirtualizationVirtualMachinesListParamsWithTimeout creates a new VirtualizationVirtualMachinesListParams object
// with the ability to set a timeout on a request.
func NewVirtualizationVirtualMachinesListParamsWithTimeout(timeout time.Duration) *VirtualizationVirtualMachinesListParams {
	return &VirtualizationVirtualMachinesListParams{
		timeout: timeout,
	}
}

// NewVirtualizationVirtualMachinesListParamsWithContext creates a new VirtualizationVirtualMachinesListParams object
// with the ability to set a context for a request.
func NewVirtualizationVirtualMachinesListParamsWithContext(ctx context.Context) *VirtualizationVirtualMachinesListParams {
	return &VirtualizationVirtualMachinesListParams{
		Context: ctx,
	}
}

// NewVirtualizationVirtualMachinesListParamsWithHTTPClient creates a new VirtualizationVirtualMachinesListParams object
// with the ability to set a custom HTTPClient for a request.
func NewVirtualizationVirtualMachinesListParamsWithHTTPClient(client *http.Client) *VirtualizationVirtualMachinesListParams {
	return &VirtualizationVirtualMachinesListParams{
		HTTPClient: client,
	}
}

/* VirtualizationVirtualMachinesListParams contains all the parameters to send to the API endpoint
   for the virtualization virtual machines list operation.

   Typically these are written to a http.Request.
*/
type VirtualizationVirtualMachinesListParams struct {

	// Cluster.
	Cluster *string

	// Clustern.
	Clustern *string

	// ClusterGroup.
	ClusterGroup *string

	// ClusterGroupn.
	ClusterGroupn *string

	// ClusterGroupID.
	ClusterGroupID *string

	// ClusterGroupIDn.
	ClusterGroupIDn *string

	// ClusterID.
	ClusterID *string

	// ClusterIDn.
	ClusterIDn *string

	// ClusterType.
	ClusterType *string

	// ClusterTypen.
	ClusterTypen *string

	// ClusterTypeID.
	ClusterTypeID *string

	// ClusterTypeIDn.
	ClusterTypeIDn *string

	// Created.
	Created *string

	// CreatedGte.
	CreatedGte *string

	// CreatedLte.
	CreatedLte *string

	// Disk.
	Disk *string

	// DiskGt.
	DiskGt *string

	// DiskGte.
	DiskGte *string

	// DiskLt.
	DiskLt *string

	// DiskLte.
	DiskLte *string

	// Diskn.
	Diskn *string

	// HasPrimaryIP.
	HasPrimaryIP *string

	// ID.
	ID *string

	// IDIc.
	IDIc *string

	// IDIe.
	IDIe *string

	// IDIew.
	IDIew *string

	// IDIsw.
	IDIsw *string

	// IDn.
	IDn *string

	// IDNic.
	IDNic *string

	// IDNie.
	IDNie *string

	// IDNiew.
	IDNiew *string

	// IDNisw.
	IDNisw *string

	// LastUpdated.
	LastUpdated *string

	// LastUpdatedGte.
	LastUpdatedGte *string

	// LastUpdatedLte.
	LastUpdatedLte *string

	/* Limit.

	   Number of results to return per page.
	*/
	Limit *int64

	// LocalContextData.
	LocalContextData *string

	// LocalContextSchema.
	LocalContextSchema *string

	// LocalContextScheman.
	LocalContextScheman *string

	// LocalContextSchemaID.
	LocalContextSchemaID *string

	// LocalContextSchemaIDn.
	LocalContextSchemaIDn *string

	// MacAddress.
	MacAddress *string

	// MacAddressIc.
	MacAddressIc *string

	// MacAddressIe.
	MacAddressIe *string

	// MacAddressIew.
	MacAddressIew *string

	// MacAddressIsw.
	MacAddressIsw *string

	// MacAddressn.
	MacAddressn *string

	// MacAddressNic.
	MacAddressNic *string

	// MacAddressNie.
	MacAddressNie *string

	// MacAddressNiew.
	MacAddressNiew *string

	// MacAddressNisw.
	MacAddressNisw *string

	// Memory.
	Memory *string

	// MemoryGt.
	MemoryGt *string

	// MemoryGte.
	MemoryGte *string

	// MemoryLt.
	MemoryLt *string

	// MemoryLte.
	MemoryLte *string

	// Memoryn.
	Memoryn *string

	// Name.
	Name *string

	// NameIc.
	NameIc *string

	// NameIe.
	NameIe *string

	// NameIew.
	NameIew *string

	// NameIsw.
	NameIsw *string

	// Namen.
	Namen *string

	// NameNic.
	NameNic *string

	// NameNie.
	NameNie *string

	// NameNiew.
	NameNiew *string

	// NameNisw.
	NameNisw *string

	/* Offset.

	   The initial index from which to return the results.
	*/
	Offset *int64

	// Platform.
	Platform *string

	// Platformn.
	Platformn *string

	// PlatformID.
	PlatformID *string

	// PlatformIDn.
	PlatformIDn *string

	// Q.
	Q *string

	// Region.
	Region *string

	// Regionn.
	Regionn *string

	// RegionID.
	RegionID *string

	// RegionIDn.
	RegionIDn *string

	// Role.
	Role *string

	// Rolen.
	Rolen *string

	// RoleID.
	RoleID *string

	// RoleIDn.
	RoleIDn *string

	// Site.
	Site *string

	// Siten.
	Siten *string

	// SiteID.
	SiteID *string

	// SiteIDn.
	SiteIDn *string

	// Status.
	Status *string

	// Statusn.
	Statusn *string

	// Tag.
	Tag *string

	// Tagn.
	Tagn *string

	// Tenant.
	Tenant *string

	// Tenantn.
	Tenantn *string

	// TenantGroup.
	TenantGroup *string

	// TenantGroupn.
	TenantGroupn *string

	// TenantGroupID.
	TenantGroupID *string

	// TenantGroupIDn.
	TenantGroupIDn *string

	// TenantID.
	TenantID *string

	// TenantIDn.
	TenantIDn *string

	// Vcpus.
	Vcpus *string

	// VcpusGt.
	VcpusGt *string

	// VcpusGte.
	VcpusGte *string

	// VcpusLt.
	VcpusLt *string

	// VcpusLte.
	VcpusLte *string

	// Vcpusn.
	Vcpusn *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the virtualization virtual machines list params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *VirtualizationVirtualMachinesListParams) WithDefaults() *VirtualizationVirtualMachinesListParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the virtualization virtual machines list params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *VirtualizationVirtualMachinesListParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the virtualization virtual machines list params
func (o *VirtualizationVirtualMachinesListParams) WithTimeout(timeout time.Duration) *VirtualizationVirtualMachinesListParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the virtualization virtual machines list params
func (o *VirtualizationVirtualMachinesListParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the virtualization virtual machines list params
func (o *VirtualizationVirtualMachinesListParams) WithContext(ctx context.Context) *VirtualizationVirtualMachinesListParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the virtualization virtual machines list params
func (o *VirtualizationVirtualMachinesListParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the virtualization virtual machines list params
func (o *VirtualizationVirtualMachinesListParams) WithHTTPClient(client *http.Client) *VirtualizationVirtualMachinesListParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the virtualization virtual machines list params
func (o *VirtualizationVirtualMachinesListParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithCluster adds the cluster to the virtualization virtual machines list params
func (o *VirtualizationVirtualMachinesListParams) WithCluster(cluster *string) *VirtualizationVirtualMachinesListParams {
	o.SetCluster(cluster)
	return o
}

// SetCluster adds the cluster to the virtualization virtual machines list params
func (o *VirtualizationVirtualMachinesListParams) SetCluster(cluster *string) {
	o.Cluster = cluster
}

// WithClustern adds the clustern to the virtualization virtual machines list params
func (o *VirtualizationVirtualMachinesListParams) WithClustern(clustern *string) *VirtualizationVirtualMachinesListParams {
	o.SetClustern(clustern)
	return o
}

// SetClustern adds the clusterN to the virtualization virtual machines list params
func (o *VirtualizationVirtualMachinesListParams) SetClustern(clustern *string) {
	o.Clustern = clustern
}

// WithClusterGroup adds the clusterGroup to the virtualization virtual machines list params
func (o *VirtualizationVirtualMachinesListParams) WithClusterGroup(clusterGroup *string) *VirtualizationVirtualMachinesListParams {
	o.SetClusterGroup(clusterGroup)
	return o
}

// SetClusterGroup adds the clusterGroup to the virtualization virtual machines list params
func (o *VirtualizationVirtualMachinesListParams) SetClusterGroup(clusterGroup *string) {
	o.ClusterGroup = clusterGroup
}

// WithClusterGroupn adds the clusterGroupn to the virtualization virtual machines list params
func (o *VirtualizationVirtualMachinesListParams) WithClusterGroupn(clusterGroupn *string) *VirtualizationVirtualMachinesListParams {
	o.SetClusterGroupn(clusterGroupn)
	return o
}

// SetClusterGroupn adds the clusterGroupN to the virtualization virtual machines list params
func (o *VirtualizationVirtualMachinesListParams) SetClusterGroupn(clusterGroupn *string) {
	o.ClusterGroupn = clusterGroupn
}

// WithClusterGroupID adds the clusterGroupID to the virtualization virtual machines list params
func (o *VirtualizationVirtualMachinesListParams) WithClusterGroupID(clusterGroupID *string) *VirtualizationVirtualMachinesListParams {
	o.SetClusterGroupID(clusterGroupID)
	return o
}

// SetClusterGroupID adds the clusterGroupId to the virtualization virtual machines list params
func (o *VirtualizationVirtualMachinesListParams) SetClusterGroupID(clusterGroupID *string) {
	o.ClusterGroupID = clusterGroupID
}

// WithClusterGroupIDn adds the clusterGroupIDn to the virtualization virtual machines list params
func (o *VirtualizationVirtualMachinesListParams) WithClusterGroupIDn(clusterGroupIDn *string) *VirtualizationVirtualMachinesListParams {
	o.SetClusterGroupIDn(clusterGroupIDn)
	return o
}

// SetClusterGroupIDn adds the clusterGroupIdN to the virtualization virtual machines list params
func (o *VirtualizationVirtualMachinesListParams) SetClusterGroupIDn(clusterGroupIDn *string) {
	o.ClusterGroupIDn = clusterGroupIDn
}

// WithClusterID adds the clusterID to the virtualization virtual machines list params
func (o *VirtualizationVirtualMachinesListParams) WithClusterID(clusterID *string) *VirtualizationVirtualMachinesListParams {
	o.SetClusterID(clusterID)
	return o
}

// SetClusterID adds the clusterId to the virtualization virtual machines list params
func (o *VirtualizationVirtualMachinesListParams) SetClusterID(clusterID *string) {
	o.ClusterID = clusterID
}

// WithClusterIDn adds the clusterIDn to the virtualization virtual machines list params
func (o *VirtualizationVirtualMachinesListParams) WithClusterIDn(clusterIDn *string) *VirtualizationVirtualMachinesListParams {
	o.SetClusterIDn(clusterIDn)
	return o
}

// SetClusterIDn adds the clusterIdN to the virtualization virtual machines list params
func (o *VirtualizationVirtualMachinesListParams) SetClusterIDn(clusterIDn *string) {
	o.ClusterIDn = clusterIDn
}

// WithClusterType adds the clusterType to the virtualization virtual machines list params
func (o *VirtualizationVirtualMachinesListParams) WithClusterType(clusterType *string) *VirtualizationVirtualMachinesListParams {
	o.SetClusterType(clusterType)
	return o
}

// SetClusterType adds the clusterType to the virtualization virtual machines list params
func (o *VirtualizationVirtualMachinesListParams) SetClusterType(clusterType *string) {
	o.ClusterType = clusterType
}

// WithClusterTypen adds the clusterTypen to the virtualization virtual machines list params
func (o *VirtualizationVirtualMachinesListParams) WithClusterTypen(clusterTypen *string) *VirtualizationVirtualMachinesListParams {
	o.SetClusterTypen(clusterTypen)
	return o
}

// SetClusterTypen adds the clusterTypeN to the virtualization virtual machines list params
func (o *VirtualizationVirtualMachinesListParams) SetClusterTypen(clusterTypen *string) {
	o.ClusterTypen = clusterTypen
}

// WithClusterTypeID adds the clusterTypeID to the virtualization virtual machines list params
func (o *VirtualizationVirtualMachinesListParams) WithClusterTypeID(clusterTypeID *string) *VirtualizationVirtualMachinesListParams {
	o.SetClusterTypeID(clusterTypeID)
	return o
}

// SetClusterTypeID adds the clusterTypeId to the virtualization virtual machines list params
func (o *VirtualizationVirtualMachinesListParams) SetClusterTypeID(clusterTypeID *string) {
	o.ClusterTypeID = clusterTypeID
}

// WithClusterTypeIDn adds the clusterTypeIDn to the virtualization virtual machines list params
func (o *VirtualizationVirtualMachinesListParams) WithClusterTypeIDn(clusterTypeIDn *string) *VirtualizationVirtualMachinesListParams {
	o.SetClusterTypeIDn(clusterTypeIDn)
	return o
}

// SetClusterTypeIDn adds the clusterTypeIdN to the virtualization virtual machines list params
func (o *VirtualizationVirtualMachinesListParams) SetClusterTypeIDn(clusterTypeIDn *string) {
	o.ClusterTypeIDn = clusterTypeIDn
}

// WithCreated adds the created to the virtualization virtual machines list params
func (o *VirtualizationVirtualMachinesListParams) WithCreated(created *string) *VirtualizationVirtualMachinesListParams {
	o.SetCreated(created)
	return o
}

// SetCreated adds the created to the virtualization virtual machines list params
func (o *VirtualizationVirtualMachinesListParams) SetCreated(created *string) {
	o.Created = created
}

// WithCreatedGte adds the createdGte to the virtualization virtual machines list params
func (o *VirtualizationVirtualMachinesListParams) WithCreatedGte(createdGte *string) *VirtualizationVirtualMachinesListParams {
	o.SetCreatedGte(createdGte)
	return o
}

// SetCreatedGte adds the createdGte to the virtualization virtual machines list params
func (o *VirtualizationVirtualMachinesListParams) SetCreatedGte(createdGte *string) {
	o.CreatedGte = createdGte
}

// WithCreatedLte adds the createdLte to the virtualization virtual machines list params
func (o *VirtualizationVirtualMachinesListParams) WithCreatedLte(createdLte *string) *VirtualizationVirtualMachinesListParams {
	o.SetCreatedLte(createdLte)
	return o
}

// SetCreatedLte adds the createdLte to the virtualization virtual machines list params
func (o *VirtualizationVirtualMachinesListParams) SetCreatedLte(createdLte *string) {
	o.CreatedLte = createdLte
}

// WithDisk adds the disk to the virtualization virtual machines list params
func (o *VirtualizationVirtualMachinesListParams) WithDisk(disk *string) *VirtualizationVirtualMachinesListParams {
	o.SetDisk(disk)
	return o
}

// SetDisk adds the disk to the virtualization virtual machines list params
func (o *VirtualizationVirtualMachinesListParams) SetDisk(disk *string) {
	o.Disk = disk
}

// WithDiskGt adds the diskGt to the virtualization virtual machines list params
func (o *VirtualizationVirtualMachinesListParams) WithDiskGt(diskGt *string) *VirtualizationVirtualMachinesListParams {
	o.SetDiskGt(diskGt)
	return o
}

// SetDiskGt adds the diskGt to the virtualization virtual machines list params
func (o *VirtualizationVirtualMachinesListParams) SetDiskGt(diskGt *string) {
	o.DiskGt = diskGt
}

// WithDiskGte adds the diskGte to the virtualization virtual machines list params
func (o *VirtualizationVirtualMachinesListParams) WithDiskGte(diskGte *string) *VirtualizationVirtualMachinesListParams {
	o.SetDiskGte(diskGte)
	return o
}

// SetDiskGte adds the diskGte to the virtualization virtual machines list params
func (o *VirtualizationVirtualMachinesListParams) SetDiskGte(diskGte *string) {
	o.DiskGte = diskGte
}

// WithDiskLt adds the diskLt to the virtualization virtual machines list params
func (o *VirtualizationVirtualMachinesListParams) WithDiskLt(diskLt *string) *VirtualizationVirtualMachinesListParams {
	o.SetDiskLt(diskLt)
	return o
}

// SetDiskLt adds the diskLt to the virtualization virtual machines list params
func (o *VirtualizationVirtualMachinesListParams) SetDiskLt(diskLt *string) {
	o.DiskLt = diskLt
}

// WithDiskLte adds the diskLte to the virtualization virtual machines list params
func (o *VirtualizationVirtualMachinesListParams) WithDiskLte(diskLte *string) *VirtualizationVirtualMachinesListParams {
	o.SetDiskLte(diskLte)
	return o
}

// SetDiskLte adds the diskLte to the virtualization virtual machines list params
func (o *VirtualizationVirtualMachinesListParams) SetDiskLte(diskLte *string) {
	o.DiskLte = diskLte
}

// WithDiskn adds the diskn to the virtualization virtual machines list params
func (o *VirtualizationVirtualMachinesListParams) WithDiskn(diskn *string) *VirtualizationVirtualMachinesListParams {
	o.SetDiskn(diskn)
	return o
}

// SetDiskn adds the diskN to the virtualization virtual machines list params
func (o *VirtualizationVirtualMachinesListParams) SetDiskn(diskn *string) {
	o.Diskn = diskn
}

// WithHasPrimaryIP adds the hasPrimaryIP to the virtualization virtual machines list params
func (o *VirtualizationVirtualMachinesListParams) WithHasPrimaryIP(hasPrimaryIP *string) *VirtualizationVirtualMachinesListParams {
	o.SetHasPrimaryIP(hasPrimaryIP)
	return o
}

// SetHasPrimaryIP adds the hasPrimaryIp to the virtualization virtual machines list params
func (o *VirtualizationVirtualMachinesListParams) SetHasPrimaryIP(hasPrimaryIP *string) {
	o.HasPrimaryIP = hasPrimaryIP
}

// WithID adds the id to the virtualization virtual machines list params
func (o *VirtualizationVirtualMachinesListParams) WithID(id *string) *VirtualizationVirtualMachinesListParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the virtualization virtual machines list params
func (o *VirtualizationVirtualMachinesListParams) SetID(id *string) {
	o.ID = id
}

// WithIDIc adds the iDIc to the virtualization virtual machines list params
func (o *VirtualizationVirtualMachinesListParams) WithIDIc(iDIc *string) *VirtualizationVirtualMachinesListParams {
	o.SetIDIc(iDIc)
	return o
}

// SetIDIc adds the idIc to the virtualization virtual machines list params
func (o *VirtualizationVirtualMachinesListParams) SetIDIc(iDIc *string) {
	o.IDIc = iDIc
}

// WithIDIe adds the iDIe to the virtualization virtual machines list params
func (o *VirtualizationVirtualMachinesListParams) WithIDIe(iDIe *string) *VirtualizationVirtualMachinesListParams {
	o.SetIDIe(iDIe)
	return o
}

// SetIDIe adds the idIe to the virtualization virtual machines list params
func (o *VirtualizationVirtualMachinesListParams) SetIDIe(iDIe *string) {
	o.IDIe = iDIe
}

// WithIDIew adds the iDIew to the virtualization virtual machines list params
func (o *VirtualizationVirtualMachinesListParams) WithIDIew(iDIew *string) *VirtualizationVirtualMachinesListParams {
	o.SetIDIew(iDIew)
	return o
}

// SetIDIew adds the idIew to the virtualization virtual machines list params
func (o *VirtualizationVirtualMachinesListParams) SetIDIew(iDIew *string) {
	o.IDIew = iDIew
}

// WithIDIsw adds the iDIsw to the virtualization virtual machines list params
func (o *VirtualizationVirtualMachinesListParams) WithIDIsw(iDIsw *string) *VirtualizationVirtualMachinesListParams {
	o.SetIDIsw(iDIsw)
	return o
}

// SetIDIsw adds the idIsw to the virtualization virtual machines list params
func (o *VirtualizationVirtualMachinesListParams) SetIDIsw(iDIsw *string) {
	o.IDIsw = iDIsw
}

// WithIDn adds the iDn to the virtualization virtual machines list params
func (o *VirtualizationVirtualMachinesListParams) WithIDn(iDn *string) *VirtualizationVirtualMachinesListParams {
	o.SetIDn(iDn)
	return o
}

// SetIDn adds the idN to the virtualization virtual machines list params
func (o *VirtualizationVirtualMachinesListParams) SetIDn(iDn *string) {
	o.IDn = iDn
}

// WithIDNic adds the iDNic to the virtualization virtual machines list params
func (o *VirtualizationVirtualMachinesListParams) WithIDNic(iDNic *string) *VirtualizationVirtualMachinesListParams {
	o.SetIDNic(iDNic)
	return o
}

// SetIDNic adds the idNic to the virtualization virtual machines list params
func (o *VirtualizationVirtualMachinesListParams) SetIDNic(iDNic *string) {
	o.IDNic = iDNic
}

// WithIDNie adds the iDNie to the virtualization virtual machines list params
func (o *VirtualizationVirtualMachinesListParams) WithIDNie(iDNie *string) *VirtualizationVirtualMachinesListParams {
	o.SetIDNie(iDNie)
	return o
}

// SetIDNie adds the idNie to the virtualization virtual machines list params
func (o *VirtualizationVirtualMachinesListParams) SetIDNie(iDNie *string) {
	o.IDNie = iDNie
}

// WithIDNiew adds the iDNiew to the virtualization virtual machines list params
func (o *VirtualizationVirtualMachinesListParams) WithIDNiew(iDNiew *string) *VirtualizationVirtualMachinesListParams {
	o.SetIDNiew(iDNiew)
	return o
}

// SetIDNiew adds the idNiew to the virtualization virtual machines list params
func (o *VirtualizationVirtualMachinesListParams) SetIDNiew(iDNiew *string) {
	o.IDNiew = iDNiew
}

// WithIDNisw adds the iDNisw to the virtualization virtual machines list params
func (o *VirtualizationVirtualMachinesListParams) WithIDNisw(iDNisw *string) *VirtualizationVirtualMachinesListParams {
	o.SetIDNisw(iDNisw)
	return o
}

// SetIDNisw adds the idNisw to the virtualization virtual machines list params
func (o *VirtualizationVirtualMachinesListParams) SetIDNisw(iDNisw *string) {
	o.IDNisw = iDNisw
}

// WithLastUpdated adds the lastUpdated to the virtualization virtual machines list params
func (o *VirtualizationVirtualMachinesListParams) WithLastUpdated(lastUpdated *string) *VirtualizationVirtualMachinesListParams {
	o.SetLastUpdated(lastUpdated)
	return o
}

// SetLastUpdated adds the lastUpdated to the virtualization virtual machines list params
func (o *VirtualizationVirtualMachinesListParams) SetLastUpdated(lastUpdated *string) {
	o.LastUpdated = lastUpdated
}

// WithLastUpdatedGte adds the lastUpdatedGte to the virtualization virtual machines list params
func (o *VirtualizationVirtualMachinesListParams) WithLastUpdatedGte(lastUpdatedGte *string) *VirtualizationVirtualMachinesListParams {
	o.SetLastUpdatedGte(lastUpdatedGte)
	return o
}

// SetLastUpdatedGte adds the lastUpdatedGte to the virtualization virtual machines list params
func (o *VirtualizationVirtualMachinesListParams) SetLastUpdatedGte(lastUpdatedGte *string) {
	o.LastUpdatedGte = lastUpdatedGte
}

// WithLastUpdatedLte adds the lastUpdatedLte to the virtualization virtual machines list params
func (o *VirtualizationVirtualMachinesListParams) WithLastUpdatedLte(lastUpdatedLte *string) *VirtualizationVirtualMachinesListParams {
	o.SetLastUpdatedLte(lastUpdatedLte)
	return o
}

// SetLastUpdatedLte adds the lastUpdatedLte to the virtualization virtual machines list params
func (o *VirtualizationVirtualMachinesListParams) SetLastUpdatedLte(lastUpdatedLte *string) {
	o.LastUpdatedLte = lastUpdatedLte
}

// WithLimit adds the limit to the virtualization virtual machines list params
func (o *VirtualizationVirtualMachinesListParams) WithLimit(limit *int64) *VirtualizationVirtualMachinesListParams {
	o.SetLimit(limit)
	return o
}

// SetLimit adds the limit to the virtualization virtual machines list params
func (o *VirtualizationVirtualMachinesListParams) SetLimit(limit *int64) {
	o.Limit = limit
}

// WithLocalContextData adds the localContextData to the virtualization virtual machines list params
func (o *VirtualizationVirtualMachinesListParams) WithLocalContextData(localContextData *string) *VirtualizationVirtualMachinesListParams {
	o.SetLocalContextData(localContextData)
	return o
}

// SetLocalContextData adds the localContextData to the virtualization virtual machines list params
func (o *VirtualizationVirtualMachinesListParams) SetLocalContextData(localContextData *string) {
	o.LocalContextData = localContextData
}

// WithLocalContextSchema adds the localContextSchema to the virtualization virtual machines list params
func (o *VirtualizationVirtualMachinesListParams) WithLocalContextSchema(localContextSchema *string) *VirtualizationVirtualMachinesListParams {
	o.SetLocalContextSchema(localContextSchema)
	return o
}

// SetLocalContextSchema adds the localContextSchema to the virtualization virtual machines list params
func (o *VirtualizationVirtualMachinesListParams) SetLocalContextSchema(localContextSchema *string) {
	o.LocalContextSchema = localContextSchema
}

// WithLocalContextScheman adds the localContextScheman to the virtualization virtual machines list params
func (o *VirtualizationVirtualMachinesListParams) WithLocalContextScheman(localContextScheman *string) *VirtualizationVirtualMachinesListParams {
	o.SetLocalContextScheman(localContextScheman)
	return o
}

// SetLocalContextScheman adds the localContextSchemaN to the virtualization virtual machines list params
func (o *VirtualizationVirtualMachinesListParams) SetLocalContextScheman(localContextScheman *string) {
	o.LocalContextScheman = localContextScheman
}

// WithLocalContextSchemaID adds the localContextSchemaID to the virtualization virtual machines list params
func (o *VirtualizationVirtualMachinesListParams) WithLocalContextSchemaID(localContextSchemaID *string) *VirtualizationVirtualMachinesListParams {
	o.SetLocalContextSchemaID(localContextSchemaID)
	return o
}

// SetLocalContextSchemaID adds the localContextSchemaId to the virtualization virtual machines list params
func (o *VirtualizationVirtualMachinesListParams) SetLocalContextSchemaID(localContextSchemaID *string) {
	o.LocalContextSchemaID = localContextSchemaID
}

// WithLocalContextSchemaIDn adds the localContextSchemaIDn to the virtualization virtual machines list params
func (o *VirtualizationVirtualMachinesListParams) WithLocalContextSchemaIDn(localContextSchemaIDn *string) *VirtualizationVirtualMachinesListParams {
	o.SetLocalContextSchemaIDn(localContextSchemaIDn)
	return o
}

// SetLocalContextSchemaIDn adds the localContextSchemaIdN to the virtualization virtual machines list params
func (o *VirtualizationVirtualMachinesListParams) SetLocalContextSchemaIDn(localContextSchemaIDn *string) {
	o.LocalContextSchemaIDn = localContextSchemaIDn
}

// WithMacAddress adds the macAddress to the virtualization virtual machines list params
func (o *VirtualizationVirtualMachinesListParams) WithMacAddress(macAddress *string) *VirtualizationVirtualMachinesListParams {
	o.SetMacAddress(macAddress)
	return o
}

// SetMacAddress adds the macAddress to the virtualization virtual machines list params
func (o *VirtualizationVirtualMachinesListParams) SetMacAddress(macAddress *string) {
	o.MacAddress = macAddress
}

// WithMacAddressIc adds the macAddressIc to the virtualization virtual machines list params
func (o *VirtualizationVirtualMachinesListParams) WithMacAddressIc(macAddressIc *string) *VirtualizationVirtualMachinesListParams {
	o.SetMacAddressIc(macAddressIc)
	return o
}

// SetMacAddressIc adds the macAddressIc to the virtualization virtual machines list params
func (o *VirtualizationVirtualMachinesListParams) SetMacAddressIc(macAddressIc *string) {
	o.MacAddressIc = macAddressIc
}

// WithMacAddressIe adds the macAddressIe to the virtualization virtual machines list params
func (o *VirtualizationVirtualMachinesListParams) WithMacAddressIe(macAddressIe *string) *VirtualizationVirtualMachinesListParams {
	o.SetMacAddressIe(macAddressIe)
	return o
}

// SetMacAddressIe adds the macAddressIe to the virtualization virtual machines list params
func (o *VirtualizationVirtualMachinesListParams) SetMacAddressIe(macAddressIe *string) {
	o.MacAddressIe = macAddressIe
}

// WithMacAddressIew adds the macAddressIew to the virtualization virtual machines list params
func (o *VirtualizationVirtualMachinesListParams) WithMacAddressIew(macAddressIew *string) *VirtualizationVirtualMachinesListParams {
	o.SetMacAddressIew(macAddressIew)
	return o
}

// SetMacAddressIew adds the macAddressIew to the virtualization virtual machines list params
func (o *VirtualizationVirtualMachinesListParams) SetMacAddressIew(macAddressIew *string) {
	o.MacAddressIew = macAddressIew
}

// WithMacAddressIsw adds the macAddressIsw to the virtualization virtual machines list params
func (o *VirtualizationVirtualMachinesListParams) WithMacAddressIsw(macAddressIsw *string) *VirtualizationVirtualMachinesListParams {
	o.SetMacAddressIsw(macAddressIsw)
	return o
}

// SetMacAddressIsw adds the macAddressIsw to the virtualization virtual machines list params
func (o *VirtualizationVirtualMachinesListParams) SetMacAddressIsw(macAddressIsw *string) {
	o.MacAddressIsw = macAddressIsw
}

// WithMacAddressn adds the macAddressn to the virtualization virtual machines list params
func (o *VirtualizationVirtualMachinesListParams) WithMacAddressn(macAddressn *string) *VirtualizationVirtualMachinesListParams {
	o.SetMacAddressn(macAddressn)
	return o
}

// SetMacAddressn adds the macAddressN to the virtualization virtual machines list params
func (o *VirtualizationVirtualMachinesListParams) SetMacAddressn(macAddressn *string) {
	o.MacAddressn = macAddressn
}

// WithMacAddressNic adds the macAddressNic to the virtualization virtual machines list params
func (o *VirtualizationVirtualMachinesListParams) WithMacAddressNic(macAddressNic *string) *VirtualizationVirtualMachinesListParams {
	o.SetMacAddressNic(macAddressNic)
	return o
}

// SetMacAddressNic adds the macAddressNic to the virtualization virtual machines list params
func (o *VirtualizationVirtualMachinesListParams) SetMacAddressNic(macAddressNic *string) {
	o.MacAddressNic = macAddressNic
}

// WithMacAddressNie adds the macAddressNie to the virtualization virtual machines list params
func (o *VirtualizationVirtualMachinesListParams) WithMacAddressNie(macAddressNie *string) *VirtualizationVirtualMachinesListParams {
	o.SetMacAddressNie(macAddressNie)
	return o
}

// SetMacAddressNie adds the macAddressNie to the virtualization virtual machines list params
func (o *VirtualizationVirtualMachinesListParams) SetMacAddressNie(macAddressNie *string) {
	o.MacAddressNie = macAddressNie
}

// WithMacAddressNiew adds the macAddressNiew to the virtualization virtual machines list params
func (o *VirtualizationVirtualMachinesListParams) WithMacAddressNiew(macAddressNiew *string) *VirtualizationVirtualMachinesListParams {
	o.SetMacAddressNiew(macAddressNiew)
	return o
}

// SetMacAddressNiew adds the macAddressNiew to the virtualization virtual machines list params
func (o *VirtualizationVirtualMachinesListParams) SetMacAddressNiew(macAddressNiew *string) {
	o.MacAddressNiew = macAddressNiew
}

// WithMacAddressNisw adds the macAddressNisw to the virtualization virtual machines list params
func (o *VirtualizationVirtualMachinesListParams) WithMacAddressNisw(macAddressNisw *string) *VirtualizationVirtualMachinesListParams {
	o.SetMacAddressNisw(macAddressNisw)
	return o
}

// SetMacAddressNisw adds the macAddressNisw to the virtualization virtual machines list params
func (o *VirtualizationVirtualMachinesListParams) SetMacAddressNisw(macAddressNisw *string) {
	o.MacAddressNisw = macAddressNisw
}

// WithMemory adds the memory to the virtualization virtual machines list params
func (o *VirtualizationVirtualMachinesListParams) WithMemory(memory *string) *VirtualizationVirtualMachinesListParams {
	o.SetMemory(memory)
	return o
}

// SetMemory adds the memory to the virtualization virtual machines list params
func (o *VirtualizationVirtualMachinesListParams) SetMemory(memory *string) {
	o.Memory = memory
}

// WithMemoryGt adds the memoryGt to the virtualization virtual machines list params
func (o *VirtualizationVirtualMachinesListParams) WithMemoryGt(memoryGt *string) *VirtualizationVirtualMachinesListParams {
	o.SetMemoryGt(memoryGt)
	return o
}

// SetMemoryGt adds the memoryGt to the virtualization virtual machines list params
func (o *VirtualizationVirtualMachinesListParams) SetMemoryGt(memoryGt *string) {
	o.MemoryGt = memoryGt
}

// WithMemoryGte adds the memoryGte to the virtualization virtual machines list params
func (o *VirtualizationVirtualMachinesListParams) WithMemoryGte(memoryGte *string) *VirtualizationVirtualMachinesListParams {
	o.SetMemoryGte(memoryGte)
	return o
}

// SetMemoryGte adds the memoryGte to the virtualization virtual machines list params
func (o *VirtualizationVirtualMachinesListParams) SetMemoryGte(memoryGte *string) {
	o.MemoryGte = memoryGte
}

// WithMemoryLt adds the memoryLt to the virtualization virtual machines list params
func (o *VirtualizationVirtualMachinesListParams) WithMemoryLt(memoryLt *string) *VirtualizationVirtualMachinesListParams {
	o.SetMemoryLt(memoryLt)
	return o
}

// SetMemoryLt adds the memoryLt to the virtualization virtual machines list params
func (o *VirtualizationVirtualMachinesListParams) SetMemoryLt(memoryLt *string) {
	o.MemoryLt = memoryLt
}

// WithMemoryLte adds the memoryLte to the virtualization virtual machines list params
func (o *VirtualizationVirtualMachinesListParams) WithMemoryLte(memoryLte *string) *VirtualizationVirtualMachinesListParams {
	o.SetMemoryLte(memoryLte)
	return o
}

// SetMemoryLte adds the memoryLte to the virtualization virtual machines list params
func (o *VirtualizationVirtualMachinesListParams) SetMemoryLte(memoryLte *string) {
	o.MemoryLte = memoryLte
}

// WithMemoryn adds the memoryn to the virtualization virtual machines list params
func (o *VirtualizationVirtualMachinesListParams) WithMemoryn(memoryn *string) *VirtualizationVirtualMachinesListParams {
	o.SetMemoryn(memoryn)
	return o
}

// SetMemoryn adds the memoryN to the virtualization virtual machines list params
func (o *VirtualizationVirtualMachinesListParams) SetMemoryn(memoryn *string) {
	o.Memoryn = memoryn
}

// WithName adds the name to the virtualization virtual machines list params
func (o *VirtualizationVirtualMachinesListParams) WithName(name *string) *VirtualizationVirtualMachinesListParams {
	o.SetName(name)
	return o
}

// SetName adds the name to the virtualization virtual machines list params
func (o *VirtualizationVirtualMachinesListParams) SetName(name *string) {
	o.Name = name
}

// WithNameIc adds the nameIc to the virtualization virtual machines list params
func (o *VirtualizationVirtualMachinesListParams) WithNameIc(nameIc *string) *VirtualizationVirtualMachinesListParams {
	o.SetNameIc(nameIc)
	return o
}

// SetNameIc adds the nameIc to the virtualization virtual machines list params
func (o *VirtualizationVirtualMachinesListParams) SetNameIc(nameIc *string) {
	o.NameIc = nameIc
}

// WithNameIe adds the nameIe to the virtualization virtual machines list params
func (o *VirtualizationVirtualMachinesListParams) WithNameIe(nameIe *string) *VirtualizationVirtualMachinesListParams {
	o.SetNameIe(nameIe)
	return o
}

// SetNameIe adds the nameIe to the virtualization virtual machines list params
func (o *VirtualizationVirtualMachinesListParams) SetNameIe(nameIe *string) {
	o.NameIe = nameIe
}

// WithNameIew adds the nameIew to the virtualization virtual machines list params
func (o *VirtualizationVirtualMachinesListParams) WithNameIew(nameIew *string) *VirtualizationVirtualMachinesListParams {
	o.SetNameIew(nameIew)
	return o
}

// SetNameIew adds the nameIew to the virtualization virtual machines list params
func (o *VirtualizationVirtualMachinesListParams) SetNameIew(nameIew *string) {
	o.NameIew = nameIew
}

// WithNameIsw adds the nameIsw to the virtualization virtual machines list params
func (o *VirtualizationVirtualMachinesListParams) WithNameIsw(nameIsw *string) *VirtualizationVirtualMachinesListParams {
	o.SetNameIsw(nameIsw)
	return o
}

// SetNameIsw adds the nameIsw to the virtualization virtual machines list params
func (o *VirtualizationVirtualMachinesListParams) SetNameIsw(nameIsw *string) {
	o.NameIsw = nameIsw
}

// WithNamen adds the namen to the virtualization virtual machines list params
func (o *VirtualizationVirtualMachinesListParams) WithNamen(namen *string) *VirtualizationVirtualMachinesListParams {
	o.SetNamen(namen)
	return o
}

// SetNamen adds the nameN to the virtualization virtual machines list params
func (o *VirtualizationVirtualMachinesListParams) SetNamen(namen *string) {
	o.Namen = namen
}

// WithNameNic adds the nameNic to the virtualization virtual machines list params
func (o *VirtualizationVirtualMachinesListParams) WithNameNic(nameNic *string) *VirtualizationVirtualMachinesListParams {
	o.SetNameNic(nameNic)
	return o
}

// SetNameNic adds the nameNic to the virtualization virtual machines list params
func (o *VirtualizationVirtualMachinesListParams) SetNameNic(nameNic *string) {
	o.NameNic = nameNic
}

// WithNameNie adds the nameNie to the virtualization virtual machines list params
func (o *VirtualizationVirtualMachinesListParams) WithNameNie(nameNie *string) *VirtualizationVirtualMachinesListParams {
	o.SetNameNie(nameNie)
	return o
}

// SetNameNie adds the nameNie to the virtualization virtual machines list params
func (o *VirtualizationVirtualMachinesListParams) SetNameNie(nameNie *string) {
	o.NameNie = nameNie
}

// WithNameNiew adds the nameNiew to the virtualization virtual machines list params
func (o *VirtualizationVirtualMachinesListParams) WithNameNiew(nameNiew *string) *VirtualizationVirtualMachinesListParams {
	o.SetNameNiew(nameNiew)
	return o
}

// SetNameNiew adds the nameNiew to the virtualization virtual machines list params
func (o *VirtualizationVirtualMachinesListParams) SetNameNiew(nameNiew *string) {
	o.NameNiew = nameNiew
}

// WithNameNisw adds the nameNisw to the virtualization virtual machines list params
func (o *VirtualizationVirtualMachinesListParams) WithNameNisw(nameNisw *string) *VirtualizationVirtualMachinesListParams {
	o.SetNameNisw(nameNisw)
	return o
}

// SetNameNisw adds the nameNisw to the virtualization virtual machines list params
func (o *VirtualizationVirtualMachinesListParams) SetNameNisw(nameNisw *string) {
	o.NameNisw = nameNisw
}

// WithOffset adds the offset to the virtualization virtual machines list params
func (o *VirtualizationVirtualMachinesListParams) WithOffset(offset *int64) *VirtualizationVirtualMachinesListParams {
	o.SetOffset(offset)
	return o
}

// SetOffset adds the offset to the virtualization virtual machines list params
func (o *VirtualizationVirtualMachinesListParams) SetOffset(offset *int64) {
	o.Offset = offset
}

// WithPlatform adds the platform to the virtualization virtual machines list params
func (o *VirtualizationVirtualMachinesListParams) WithPlatform(platform *string) *VirtualizationVirtualMachinesListParams {
	o.SetPlatform(platform)
	return o
}

// SetPlatform adds the platform to the virtualization virtual machines list params
func (o *VirtualizationVirtualMachinesListParams) SetPlatform(platform *string) {
	o.Platform = platform
}

// WithPlatformn adds the platformn to the virtualization virtual machines list params
func (o *VirtualizationVirtualMachinesListParams) WithPlatformn(platformn *string) *VirtualizationVirtualMachinesListParams {
	o.SetPlatformn(platformn)
	return o
}

// SetPlatformn adds the platformN to the virtualization virtual machines list params
func (o *VirtualizationVirtualMachinesListParams) SetPlatformn(platformn *string) {
	o.Platformn = platformn
}

// WithPlatformID adds the platformID to the virtualization virtual machines list params
func (o *VirtualizationVirtualMachinesListParams) WithPlatformID(platformID *string) *VirtualizationVirtualMachinesListParams {
	o.SetPlatformID(platformID)
	return o
}

// SetPlatformID adds the platformId to the virtualization virtual machines list params
func (o *VirtualizationVirtualMachinesListParams) SetPlatformID(platformID *string) {
	o.PlatformID = platformID
}

// WithPlatformIDn adds the platformIDn to the virtualization virtual machines list params
func (o *VirtualizationVirtualMachinesListParams) WithPlatformIDn(platformIDn *string) *VirtualizationVirtualMachinesListParams {
	o.SetPlatformIDn(platformIDn)
	return o
}

// SetPlatformIDn adds the platformIdN to the virtualization virtual machines list params
func (o *VirtualizationVirtualMachinesListParams) SetPlatformIDn(platformIDn *string) {
	o.PlatformIDn = platformIDn
}

// WithQ adds the q to the virtualization virtual machines list params
func (o *VirtualizationVirtualMachinesListParams) WithQ(q *string) *VirtualizationVirtualMachinesListParams {
	o.SetQ(q)
	return o
}

// SetQ adds the q to the virtualization virtual machines list params
func (o *VirtualizationVirtualMachinesListParams) SetQ(q *string) {
	o.Q = q
}

// WithRegion adds the region to the virtualization virtual machines list params
func (o *VirtualizationVirtualMachinesListParams) WithRegion(region *string) *VirtualizationVirtualMachinesListParams {
	o.SetRegion(region)
	return o
}

// SetRegion adds the region to the virtualization virtual machines list params
func (o *VirtualizationVirtualMachinesListParams) SetRegion(region *string) {
	o.Region = region
}

// WithRegionn adds the regionn to the virtualization virtual machines list params
func (o *VirtualizationVirtualMachinesListParams) WithRegionn(regionn *string) *VirtualizationVirtualMachinesListParams {
	o.SetRegionn(regionn)
	return o
}

// SetRegionn adds the regionN to the virtualization virtual machines list params
func (o *VirtualizationVirtualMachinesListParams) SetRegionn(regionn *string) {
	o.Regionn = regionn
}

// WithRegionID adds the regionID to the virtualization virtual machines list params
func (o *VirtualizationVirtualMachinesListParams) WithRegionID(regionID *string) *VirtualizationVirtualMachinesListParams {
	o.SetRegionID(regionID)
	return o
}

// SetRegionID adds the regionId to the virtualization virtual machines list params
func (o *VirtualizationVirtualMachinesListParams) SetRegionID(regionID *string) {
	o.RegionID = regionID
}

// WithRegionIDn adds the regionIDn to the virtualization virtual machines list params
func (o *VirtualizationVirtualMachinesListParams) WithRegionIDn(regionIDn *string) *VirtualizationVirtualMachinesListParams {
	o.SetRegionIDn(regionIDn)
	return o
}

// SetRegionIDn adds the regionIdN to the virtualization virtual machines list params
func (o *VirtualizationVirtualMachinesListParams) SetRegionIDn(regionIDn *string) {
	o.RegionIDn = regionIDn
}

// WithRole adds the role to the virtualization virtual machines list params
func (o *VirtualizationVirtualMachinesListParams) WithRole(role *string) *VirtualizationVirtualMachinesListParams {
	o.SetRole(role)
	return o
}

// SetRole adds the role to the virtualization virtual machines list params
func (o *VirtualizationVirtualMachinesListParams) SetRole(role *string) {
	o.Role = role
}

// WithRolen adds the rolen to the virtualization virtual machines list params
func (o *VirtualizationVirtualMachinesListParams) WithRolen(rolen *string) *VirtualizationVirtualMachinesListParams {
	o.SetRolen(rolen)
	return o
}

// SetRolen adds the roleN to the virtualization virtual machines list params
func (o *VirtualizationVirtualMachinesListParams) SetRolen(rolen *string) {
	o.Rolen = rolen
}

// WithRoleID adds the roleID to the virtualization virtual machines list params
func (o *VirtualizationVirtualMachinesListParams) WithRoleID(roleID *string) *VirtualizationVirtualMachinesListParams {
	o.SetRoleID(roleID)
	return o
}

// SetRoleID adds the roleId to the virtualization virtual machines list params
func (o *VirtualizationVirtualMachinesListParams) SetRoleID(roleID *string) {
	o.RoleID = roleID
}

// WithRoleIDn adds the roleIDn to the virtualization virtual machines list params
func (o *VirtualizationVirtualMachinesListParams) WithRoleIDn(roleIDn *string) *VirtualizationVirtualMachinesListParams {
	o.SetRoleIDn(roleIDn)
	return o
}

// SetRoleIDn adds the roleIdN to the virtualization virtual machines list params
func (o *VirtualizationVirtualMachinesListParams) SetRoleIDn(roleIDn *string) {
	o.RoleIDn = roleIDn
}

// WithSite adds the site to the virtualization virtual machines list params
func (o *VirtualizationVirtualMachinesListParams) WithSite(site *string) *VirtualizationVirtualMachinesListParams {
	o.SetSite(site)
	return o
}

// SetSite adds the site to the virtualization virtual machines list params
func (o *VirtualizationVirtualMachinesListParams) SetSite(site *string) {
	o.Site = site
}

// WithSiten adds the siten to the virtualization virtual machines list params
func (o *VirtualizationVirtualMachinesListParams) WithSiten(siten *string) *VirtualizationVirtualMachinesListParams {
	o.SetSiten(siten)
	return o
}

// SetSiten adds the siteN to the virtualization virtual machines list params
func (o *VirtualizationVirtualMachinesListParams) SetSiten(siten *string) {
	o.Siten = siten
}

// WithSiteID adds the siteID to the virtualization virtual machines list params
func (o *VirtualizationVirtualMachinesListParams) WithSiteID(siteID *string) *VirtualizationVirtualMachinesListParams {
	o.SetSiteID(siteID)
	return o
}

// SetSiteID adds the siteId to the virtualization virtual machines list params
func (o *VirtualizationVirtualMachinesListParams) SetSiteID(siteID *string) {
	o.SiteID = siteID
}

// WithSiteIDn adds the siteIDn to the virtualization virtual machines list params
func (o *VirtualizationVirtualMachinesListParams) WithSiteIDn(siteIDn *string) *VirtualizationVirtualMachinesListParams {
	o.SetSiteIDn(siteIDn)
	return o
}

// SetSiteIDn adds the siteIdN to the virtualization virtual machines list params
func (o *VirtualizationVirtualMachinesListParams) SetSiteIDn(siteIDn *string) {
	o.SiteIDn = siteIDn
}

// WithStatus adds the status to the virtualization virtual machines list params
func (o *VirtualizationVirtualMachinesListParams) WithStatus(status *string) *VirtualizationVirtualMachinesListParams {
	o.SetStatus(status)
	return o
}

// SetStatus adds the status to the virtualization virtual machines list params
func (o *VirtualizationVirtualMachinesListParams) SetStatus(status *string) {
	o.Status = status
}

// WithStatusn adds the statusn to the virtualization virtual machines list params
func (o *VirtualizationVirtualMachinesListParams) WithStatusn(statusn *string) *VirtualizationVirtualMachinesListParams {
	o.SetStatusn(statusn)
	return o
}

// SetStatusn adds the statusN to the virtualization virtual machines list params
func (o *VirtualizationVirtualMachinesListParams) SetStatusn(statusn *string) {
	o.Statusn = statusn
}

// WithTag adds the tag to the virtualization virtual machines list params
func (o *VirtualizationVirtualMachinesListParams) WithTag(tag *string) *VirtualizationVirtualMachinesListParams {
	o.SetTag(tag)
	return o
}

// SetTag adds the tag to the virtualization virtual machines list params
func (o *VirtualizationVirtualMachinesListParams) SetTag(tag *string) {
	o.Tag = tag
}

// WithTagn adds the tagn to the virtualization virtual machines list params
func (o *VirtualizationVirtualMachinesListParams) WithTagn(tagn *string) *VirtualizationVirtualMachinesListParams {
	o.SetTagn(tagn)
	return o
}

// SetTagn adds the tagN to the virtualization virtual machines list params
func (o *VirtualizationVirtualMachinesListParams) SetTagn(tagn *string) {
	o.Tagn = tagn
}

// WithTenant adds the tenant to the virtualization virtual machines list params
func (o *VirtualizationVirtualMachinesListParams) WithTenant(tenant *string) *VirtualizationVirtualMachinesListParams {
	o.SetTenant(tenant)
	return o
}

// SetTenant adds the tenant to the virtualization virtual machines list params
func (o *VirtualizationVirtualMachinesListParams) SetTenant(tenant *string) {
	o.Tenant = tenant
}

// WithTenantn adds the tenantn to the virtualization virtual machines list params
func (o *VirtualizationVirtualMachinesListParams) WithTenantn(tenantn *string) *VirtualizationVirtualMachinesListParams {
	o.SetTenantn(tenantn)
	return o
}

// SetTenantn adds the tenantN to the virtualization virtual machines list params
func (o *VirtualizationVirtualMachinesListParams) SetTenantn(tenantn *string) {
	o.Tenantn = tenantn
}

// WithTenantGroup adds the tenantGroup to the virtualization virtual machines list params
func (o *VirtualizationVirtualMachinesListParams) WithTenantGroup(tenantGroup *string) *VirtualizationVirtualMachinesListParams {
	o.SetTenantGroup(tenantGroup)
	return o
}

// SetTenantGroup adds the tenantGroup to the virtualization virtual machines list params
func (o *VirtualizationVirtualMachinesListParams) SetTenantGroup(tenantGroup *string) {
	o.TenantGroup = tenantGroup
}

// WithTenantGroupn adds the tenantGroupn to the virtualization virtual machines list params
func (o *VirtualizationVirtualMachinesListParams) WithTenantGroupn(tenantGroupn *string) *VirtualizationVirtualMachinesListParams {
	o.SetTenantGroupn(tenantGroupn)
	return o
}

// SetTenantGroupn adds the tenantGroupN to the virtualization virtual machines list params
func (o *VirtualizationVirtualMachinesListParams) SetTenantGroupn(tenantGroupn *string) {
	o.TenantGroupn = tenantGroupn
}

// WithTenantGroupID adds the tenantGroupID to the virtualization virtual machines list params
func (o *VirtualizationVirtualMachinesListParams) WithTenantGroupID(tenantGroupID *string) *VirtualizationVirtualMachinesListParams {
	o.SetTenantGroupID(tenantGroupID)
	return o
}

// SetTenantGroupID adds the tenantGroupId to the virtualization virtual machines list params
func (o *VirtualizationVirtualMachinesListParams) SetTenantGroupID(tenantGroupID *string) {
	o.TenantGroupID = tenantGroupID
}

// WithTenantGroupIDn adds the tenantGroupIDn to the virtualization virtual machines list params
func (o *VirtualizationVirtualMachinesListParams) WithTenantGroupIDn(tenantGroupIDn *string) *VirtualizationVirtualMachinesListParams {
	o.SetTenantGroupIDn(tenantGroupIDn)
	return o
}

// SetTenantGroupIDn adds the tenantGroupIdN to the virtualization virtual machines list params
func (o *VirtualizationVirtualMachinesListParams) SetTenantGroupIDn(tenantGroupIDn *string) {
	o.TenantGroupIDn = tenantGroupIDn
}

// WithTenantID adds the tenantID to the virtualization virtual machines list params
func (o *VirtualizationVirtualMachinesListParams) WithTenantID(tenantID *string) *VirtualizationVirtualMachinesListParams {
	o.SetTenantID(tenantID)
	return o
}

// SetTenantID adds the tenantId to the virtualization virtual machines list params
func (o *VirtualizationVirtualMachinesListParams) SetTenantID(tenantID *string) {
	o.TenantID = tenantID
}

// WithTenantIDn adds the tenantIDn to the virtualization virtual machines list params
func (o *VirtualizationVirtualMachinesListParams) WithTenantIDn(tenantIDn *string) *VirtualizationVirtualMachinesListParams {
	o.SetTenantIDn(tenantIDn)
	return o
}

// SetTenantIDn adds the tenantIdN to the virtualization virtual machines list params
func (o *VirtualizationVirtualMachinesListParams) SetTenantIDn(tenantIDn *string) {
	o.TenantIDn = tenantIDn
}

// WithVcpus adds the vcpus to the virtualization virtual machines list params
func (o *VirtualizationVirtualMachinesListParams) WithVcpus(vcpus *string) *VirtualizationVirtualMachinesListParams {
	o.SetVcpus(vcpus)
	return o
}

// SetVcpus adds the vcpus to the virtualization virtual machines list params
func (o *VirtualizationVirtualMachinesListParams) SetVcpus(vcpus *string) {
	o.Vcpus = vcpus
}

// WithVcpusGt adds the vcpusGt to the virtualization virtual machines list params
func (o *VirtualizationVirtualMachinesListParams) WithVcpusGt(vcpusGt *string) *VirtualizationVirtualMachinesListParams {
	o.SetVcpusGt(vcpusGt)
	return o
}

// SetVcpusGt adds the vcpusGt to the virtualization virtual machines list params
func (o *VirtualizationVirtualMachinesListParams) SetVcpusGt(vcpusGt *string) {
	o.VcpusGt = vcpusGt
}

// WithVcpusGte adds the vcpusGte to the virtualization virtual machines list params
func (o *VirtualizationVirtualMachinesListParams) WithVcpusGte(vcpusGte *string) *VirtualizationVirtualMachinesListParams {
	o.SetVcpusGte(vcpusGte)
	return o
}

// SetVcpusGte adds the vcpusGte to the virtualization virtual machines list params
func (o *VirtualizationVirtualMachinesListParams) SetVcpusGte(vcpusGte *string) {
	o.VcpusGte = vcpusGte
}

// WithVcpusLt adds the vcpusLt to the virtualization virtual machines list params
func (o *VirtualizationVirtualMachinesListParams) WithVcpusLt(vcpusLt *string) *VirtualizationVirtualMachinesListParams {
	o.SetVcpusLt(vcpusLt)
	return o
}

// SetVcpusLt adds the vcpusLt to the virtualization virtual machines list params
func (o *VirtualizationVirtualMachinesListParams) SetVcpusLt(vcpusLt *string) {
	o.VcpusLt = vcpusLt
}

// WithVcpusLte adds the vcpusLte to the virtualization virtual machines list params
func (o *VirtualizationVirtualMachinesListParams) WithVcpusLte(vcpusLte *string) *VirtualizationVirtualMachinesListParams {
	o.SetVcpusLte(vcpusLte)
	return o
}

// SetVcpusLte adds the vcpusLte to the virtualization virtual machines list params
func (o *VirtualizationVirtualMachinesListParams) SetVcpusLte(vcpusLte *string) {
	o.VcpusLte = vcpusLte
}

// WithVcpusn adds the vcpusn to the virtualization virtual machines list params
func (o *VirtualizationVirtualMachinesListParams) WithVcpusn(vcpusn *string) *VirtualizationVirtualMachinesListParams {
	o.SetVcpusn(vcpusn)
	return o
}

// SetVcpusn adds the vcpusN to the virtualization virtual machines list params
func (o *VirtualizationVirtualMachinesListParams) SetVcpusn(vcpusn *string) {
	o.Vcpusn = vcpusn
}

// WriteToRequest writes these params to a swagger request
func (o *VirtualizationVirtualMachinesListParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Cluster != nil {

		// query param cluster
		var qrCluster string

		if o.Cluster != nil {
			qrCluster = *o.Cluster
		}
		qCluster := qrCluster
		if qCluster != "" {

			if err := r.SetQueryParam("cluster", qCluster); err != nil {
				return err
			}
		}
	}

	if o.Clustern != nil {

		// query param cluster__n
		var qrClustern string

		if o.Clustern != nil {
			qrClustern = *o.Clustern
		}
		qClustern := qrClustern
		if qClustern != "" {

			if err := r.SetQueryParam("cluster__n", qClustern); err != nil {
				return err
			}
		}
	}

	if o.ClusterGroup != nil {

		// query param cluster_group
		var qrClusterGroup string

		if o.ClusterGroup != nil {
			qrClusterGroup = *o.ClusterGroup
		}
		qClusterGroup := qrClusterGroup
		if qClusterGroup != "" {

			if err := r.SetQueryParam("cluster_group", qClusterGroup); err != nil {
				return err
			}
		}
	}

	if o.ClusterGroupn != nil {

		// query param cluster_group__n
		var qrClusterGroupn string

		if o.ClusterGroupn != nil {
			qrClusterGroupn = *o.ClusterGroupn
		}
		qClusterGroupn := qrClusterGroupn
		if qClusterGroupn != "" {

			if err := r.SetQueryParam("cluster_group__n", qClusterGroupn); err != nil {
				return err
			}
		}
	}

	if o.ClusterGroupID != nil {

		// query param cluster_group_id
		var qrClusterGroupID string

		if o.ClusterGroupID != nil {
			qrClusterGroupID = *o.ClusterGroupID
		}
		qClusterGroupID := qrClusterGroupID
		if qClusterGroupID != "" {

			if err := r.SetQueryParam("cluster_group_id", qClusterGroupID); err != nil {
				return err
			}
		}
	}

	if o.ClusterGroupIDn != nil {

		// query param cluster_group_id__n
		var qrClusterGroupIDn string

		if o.ClusterGroupIDn != nil {
			qrClusterGroupIDn = *o.ClusterGroupIDn
		}
		qClusterGroupIDn := qrClusterGroupIDn
		if qClusterGroupIDn != "" {

			if err := r.SetQueryParam("cluster_group_id__n", qClusterGroupIDn); err != nil {
				return err
			}
		}
	}

	if o.ClusterID != nil {

		// query param cluster_id
		var qrClusterID string

		if o.ClusterID != nil {
			qrClusterID = *o.ClusterID
		}
		qClusterID := qrClusterID
		if qClusterID != "" {

			if err := r.SetQueryParam("cluster_id", qClusterID); err != nil {
				return err
			}
		}
	}

	if o.ClusterIDn != nil {

		// query param cluster_id__n
		var qrClusterIDn string

		if o.ClusterIDn != nil {
			qrClusterIDn = *o.ClusterIDn
		}
		qClusterIDn := qrClusterIDn
		if qClusterIDn != "" {

			if err := r.SetQueryParam("cluster_id__n", qClusterIDn); err != nil {
				return err
			}
		}
	}

	if o.ClusterType != nil {

		// query param cluster_type
		var qrClusterType string

		if o.ClusterType != nil {
			qrClusterType = *o.ClusterType
		}
		qClusterType := qrClusterType
		if qClusterType != "" {

			if err := r.SetQueryParam("cluster_type", qClusterType); err != nil {
				return err
			}
		}
	}

	if o.ClusterTypen != nil {

		// query param cluster_type__n
		var qrClusterTypen string

		if o.ClusterTypen != nil {
			qrClusterTypen = *o.ClusterTypen
		}
		qClusterTypen := qrClusterTypen
		if qClusterTypen != "" {

			if err := r.SetQueryParam("cluster_type__n", qClusterTypen); err != nil {
				return err
			}
		}
	}

	if o.ClusterTypeID != nil {

		// query param cluster_type_id
		var qrClusterTypeID string

		if o.ClusterTypeID != nil {
			qrClusterTypeID = *o.ClusterTypeID
		}
		qClusterTypeID := qrClusterTypeID
		if qClusterTypeID != "" {

			if err := r.SetQueryParam("cluster_type_id", qClusterTypeID); err != nil {
				return err
			}
		}
	}

	if o.ClusterTypeIDn != nil {

		// query param cluster_type_id__n
		var qrClusterTypeIDn string

		if o.ClusterTypeIDn != nil {
			qrClusterTypeIDn = *o.ClusterTypeIDn
		}
		qClusterTypeIDn := qrClusterTypeIDn
		if qClusterTypeIDn != "" {

			if err := r.SetQueryParam("cluster_type_id__n", qClusterTypeIDn); err != nil {
				return err
			}
		}
	}

	if o.Created != nil {

		// query param created
		var qrCreated string

		if o.Created != nil {
			qrCreated = *o.Created
		}
		qCreated := qrCreated
		if qCreated != "" {

			if err := r.SetQueryParam("created", qCreated); err != nil {
				return err
			}
		}
	}

	if o.CreatedGte != nil {

		// query param created__gte
		var qrCreatedGte string

		if o.CreatedGte != nil {
			qrCreatedGte = *o.CreatedGte
		}
		qCreatedGte := qrCreatedGte
		if qCreatedGte != "" {

			if err := r.SetQueryParam("created__gte", qCreatedGte); err != nil {
				return err
			}
		}
	}

	if o.CreatedLte != nil {

		// query param created__lte
		var qrCreatedLte string

		if o.CreatedLte != nil {
			qrCreatedLte = *o.CreatedLte
		}
		qCreatedLte := qrCreatedLte
		if qCreatedLte != "" {

			if err := r.SetQueryParam("created__lte", qCreatedLte); err != nil {
				return err
			}
		}
	}

	if o.Disk != nil {

		// query param disk
		var qrDisk string

		if o.Disk != nil {
			qrDisk = *o.Disk
		}
		qDisk := qrDisk
		if qDisk != "" {

			if err := r.SetQueryParam("disk", qDisk); err != nil {
				return err
			}
		}
	}

	if o.DiskGt != nil {

		// query param disk__gt
		var qrDiskGt string

		if o.DiskGt != nil {
			qrDiskGt = *o.DiskGt
		}
		qDiskGt := qrDiskGt
		if qDiskGt != "" {

			if err := r.SetQueryParam("disk__gt", qDiskGt); err != nil {
				return err
			}
		}
	}

	if o.DiskGte != nil {

		// query param disk__gte
		var qrDiskGte string

		if o.DiskGte != nil {
			qrDiskGte = *o.DiskGte
		}
		qDiskGte := qrDiskGte
		if qDiskGte != "" {

			if err := r.SetQueryParam("disk__gte", qDiskGte); err != nil {
				return err
			}
		}
	}

	if o.DiskLt != nil {

		// query param disk__lt
		var qrDiskLt string

		if o.DiskLt != nil {
			qrDiskLt = *o.DiskLt
		}
		qDiskLt := qrDiskLt
		if qDiskLt != "" {

			if err := r.SetQueryParam("disk__lt", qDiskLt); err != nil {
				return err
			}
		}
	}

	if o.DiskLte != nil {

		// query param disk__lte
		var qrDiskLte string

		if o.DiskLte != nil {
			qrDiskLte = *o.DiskLte
		}
		qDiskLte := qrDiskLte
		if qDiskLte != "" {

			if err := r.SetQueryParam("disk__lte", qDiskLte); err != nil {
				return err
			}
		}
	}

	if o.Diskn != nil {

		// query param disk__n
		var qrDiskn string

		if o.Diskn != nil {
			qrDiskn = *o.Diskn
		}
		qDiskn := qrDiskn
		if qDiskn != "" {

			if err := r.SetQueryParam("disk__n", qDiskn); err != nil {
				return err
			}
		}
	}

	if o.HasPrimaryIP != nil {

		// query param has_primary_ip
		var qrHasPrimaryIP string

		if o.HasPrimaryIP != nil {
			qrHasPrimaryIP = *o.HasPrimaryIP
		}
		qHasPrimaryIP := qrHasPrimaryIP
		if qHasPrimaryIP != "" {

			if err := r.SetQueryParam("has_primary_ip", qHasPrimaryIP); err != nil {
				return err
			}
		}
	}

	if o.ID != nil {

		// query param id
		var qrID string

		if o.ID != nil {
			qrID = *o.ID
		}
		qID := qrID
		if qID != "" {

			if err := r.SetQueryParam("id", qID); err != nil {
				return err
			}
		}
	}

	if o.IDIc != nil {

		// query param id__ic
		var qrIDIc string

		if o.IDIc != nil {
			qrIDIc = *o.IDIc
		}
		qIDIc := qrIDIc
		if qIDIc != "" {

			if err := r.SetQueryParam("id__ic", qIDIc); err != nil {
				return err
			}
		}
	}

	if o.IDIe != nil {

		// query param id__ie
		var qrIDIe string

		if o.IDIe != nil {
			qrIDIe = *o.IDIe
		}
		qIDIe := qrIDIe
		if qIDIe != "" {

			if err := r.SetQueryParam("id__ie", qIDIe); err != nil {
				return err
			}
		}
	}

	if o.IDIew != nil {

		// query param id__iew
		var qrIDIew string

		if o.IDIew != nil {
			qrIDIew = *o.IDIew
		}
		qIDIew := qrIDIew
		if qIDIew != "" {

			if err := r.SetQueryParam("id__iew", qIDIew); err != nil {
				return err
			}
		}
	}

	if o.IDIsw != nil {

		// query param id__isw
		var qrIDIsw string

		if o.IDIsw != nil {
			qrIDIsw = *o.IDIsw
		}
		qIDIsw := qrIDIsw
		if qIDIsw != "" {

			if err := r.SetQueryParam("id__isw", qIDIsw); err != nil {
				return err
			}
		}
	}

	if o.IDn != nil {

		// query param id__n
		var qrIDn string

		if o.IDn != nil {
			qrIDn = *o.IDn
		}
		qIDn := qrIDn
		if qIDn != "" {

			if err := r.SetQueryParam("id__n", qIDn); err != nil {
				return err
			}
		}
	}

	if o.IDNic != nil {

		// query param id__nic
		var qrIDNic string

		if o.IDNic != nil {
			qrIDNic = *o.IDNic
		}
		qIDNic := qrIDNic
		if qIDNic != "" {

			if err := r.SetQueryParam("id__nic", qIDNic); err != nil {
				return err
			}
		}
	}

	if o.IDNie != nil {

		// query param id__nie
		var qrIDNie string

		if o.IDNie != nil {
			qrIDNie = *o.IDNie
		}
		qIDNie := qrIDNie
		if qIDNie != "" {

			if err := r.SetQueryParam("id__nie", qIDNie); err != nil {
				return err
			}
		}
	}

	if o.IDNiew != nil {

		// query param id__niew
		var qrIDNiew string

		if o.IDNiew != nil {
			qrIDNiew = *o.IDNiew
		}
		qIDNiew := qrIDNiew
		if qIDNiew != "" {

			if err := r.SetQueryParam("id__niew", qIDNiew); err != nil {
				return err
			}
		}
	}

	if o.IDNisw != nil {

		// query param id__nisw
		var qrIDNisw string

		if o.IDNisw != nil {
			qrIDNisw = *o.IDNisw
		}
		qIDNisw := qrIDNisw
		if qIDNisw != "" {

			if err := r.SetQueryParam("id__nisw", qIDNisw); err != nil {
				return err
			}
		}
	}

	if o.LastUpdated != nil {

		// query param last_updated
		var qrLastUpdated string

		if o.LastUpdated != nil {
			qrLastUpdated = *o.LastUpdated
		}
		qLastUpdated := qrLastUpdated
		if qLastUpdated != "" {

			if err := r.SetQueryParam("last_updated", qLastUpdated); err != nil {
				return err
			}
		}
	}

	if o.LastUpdatedGte != nil {

		// query param last_updated__gte
		var qrLastUpdatedGte string

		if o.LastUpdatedGte != nil {
			qrLastUpdatedGte = *o.LastUpdatedGte
		}
		qLastUpdatedGte := qrLastUpdatedGte
		if qLastUpdatedGte != "" {

			if err := r.SetQueryParam("last_updated__gte", qLastUpdatedGte); err != nil {
				return err
			}
		}
	}

	if o.LastUpdatedLte != nil {

		// query param last_updated__lte
		var qrLastUpdatedLte string

		if o.LastUpdatedLte != nil {
			qrLastUpdatedLte = *o.LastUpdatedLte
		}
		qLastUpdatedLte := qrLastUpdatedLte
		if qLastUpdatedLte != "" {

			if err := r.SetQueryParam("last_updated__lte", qLastUpdatedLte); err != nil {
				return err
			}
		}
	}

	if o.Limit != nil {

		// query param limit
		var qrLimit int64

		if o.Limit != nil {
			qrLimit = *o.Limit
		}
		qLimit := swag.FormatInt64(qrLimit)
		if qLimit != "" {

			if err := r.SetQueryParam("limit", qLimit); err != nil {
				return err
			}
		}
	}

	if o.LocalContextData != nil {

		// query param local_context_data
		var qrLocalContextData string

		if o.LocalContextData != nil {
			qrLocalContextData = *o.LocalContextData
		}
		qLocalContextData := qrLocalContextData
		if qLocalContextData != "" {

			if err := r.SetQueryParam("local_context_data", qLocalContextData); err != nil {
				return err
			}
		}
	}

	if o.LocalContextSchema != nil {

		// query param local_context_schema
		var qrLocalContextSchema string

		if o.LocalContextSchema != nil {
			qrLocalContextSchema = *o.LocalContextSchema
		}
		qLocalContextSchema := qrLocalContextSchema
		if qLocalContextSchema != "" {

			if err := r.SetQueryParam("local_context_schema", qLocalContextSchema); err != nil {
				return err
			}
		}
	}

	if o.LocalContextScheman != nil {

		// query param local_context_schema__n
		var qrLocalContextScheman string

		if o.LocalContextScheman != nil {
			qrLocalContextScheman = *o.LocalContextScheman
		}
		qLocalContextScheman := qrLocalContextScheman
		if qLocalContextScheman != "" {

			if err := r.SetQueryParam("local_context_schema__n", qLocalContextScheman); err != nil {
				return err
			}
		}
	}

	if o.LocalContextSchemaID != nil {

		// query param local_context_schema_id
		var qrLocalContextSchemaID string

		if o.LocalContextSchemaID != nil {
			qrLocalContextSchemaID = *o.LocalContextSchemaID
		}
		qLocalContextSchemaID := qrLocalContextSchemaID
		if qLocalContextSchemaID != "" {

			if err := r.SetQueryParam("local_context_schema_id", qLocalContextSchemaID); err != nil {
				return err
			}
		}
	}

	if o.LocalContextSchemaIDn != nil {

		// query param local_context_schema_id__n
		var qrLocalContextSchemaIDn string

		if o.LocalContextSchemaIDn != nil {
			qrLocalContextSchemaIDn = *o.LocalContextSchemaIDn
		}
		qLocalContextSchemaIDn := qrLocalContextSchemaIDn
		if qLocalContextSchemaIDn != "" {

			if err := r.SetQueryParam("local_context_schema_id__n", qLocalContextSchemaIDn); err != nil {
				return err
			}
		}
	}

	if o.MacAddress != nil {

		// query param mac_address
		var qrMacAddress string

		if o.MacAddress != nil {
			qrMacAddress = *o.MacAddress
		}
		qMacAddress := qrMacAddress
		if qMacAddress != "" {

			if err := r.SetQueryParam("mac_address", qMacAddress); err != nil {
				return err
			}
		}
	}

	if o.MacAddressIc != nil {

		// query param mac_address__ic
		var qrMacAddressIc string

		if o.MacAddressIc != nil {
			qrMacAddressIc = *o.MacAddressIc
		}
		qMacAddressIc := qrMacAddressIc
		if qMacAddressIc != "" {

			if err := r.SetQueryParam("mac_address__ic", qMacAddressIc); err != nil {
				return err
			}
		}
	}

	if o.MacAddressIe != nil {

		// query param mac_address__ie
		var qrMacAddressIe string

		if o.MacAddressIe != nil {
			qrMacAddressIe = *o.MacAddressIe
		}
		qMacAddressIe := qrMacAddressIe
		if qMacAddressIe != "" {

			if err := r.SetQueryParam("mac_address__ie", qMacAddressIe); err != nil {
				return err
			}
		}
	}

	if o.MacAddressIew != nil {

		// query param mac_address__iew
		var qrMacAddressIew string

		if o.MacAddressIew != nil {
			qrMacAddressIew = *o.MacAddressIew
		}
		qMacAddressIew := qrMacAddressIew
		if qMacAddressIew != "" {

			if err := r.SetQueryParam("mac_address__iew", qMacAddressIew); err != nil {
				return err
			}
		}
	}

	if o.MacAddressIsw != nil {

		// query param mac_address__isw
		var qrMacAddressIsw string

		if o.MacAddressIsw != nil {
			qrMacAddressIsw = *o.MacAddressIsw
		}
		qMacAddressIsw := qrMacAddressIsw
		if qMacAddressIsw != "" {

			if err := r.SetQueryParam("mac_address__isw", qMacAddressIsw); err != nil {
				return err
			}
		}
	}

	if o.MacAddressn != nil {

		// query param mac_address__n
		var qrMacAddressn string

		if o.MacAddressn != nil {
			qrMacAddressn = *o.MacAddressn
		}
		qMacAddressn := qrMacAddressn
		if qMacAddressn != "" {

			if err := r.SetQueryParam("mac_address__n", qMacAddressn); err != nil {
				return err
			}
		}
	}

	if o.MacAddressNic != nil {

		// query param mac_address__nic
		var qrMacAddressNic string

		if o.MacAddressNic != nil {
			qrMacAddressNic = *o.MacAddressNic
		}
		qMacAddressNic := qrMacAddressNic
		if qMacAddressNic != "" {

			if err := r.SetQueryParam("mac_address__nic", qMacAddressNic); err != nil {
				return err
			}
		}
	}

	if o.MacAddressNie != nil {

		// query param mac_address__nie
		var qrMacAddressNie string

		if o.MacAddressNie != nil {
			qrMacAddressNie = *o.MacAddressNie
		}
		qMacAddressNie := qrMacAddressNie
		if qMacAddressNie != "" {

			if err := r.SetQueryParam("mac_address__nie", qMacAddressNie); err != nil {
				return err
			}
		}
	}

	if o.MacAddressNiew != nil {

		// query param mac_address__niew
		var qrMacAddressNiew string

		if o.MacAddressNiew != nil {
			qrMacAddressNiew = *o.MacAddressNiew
		}
		qMacAddressNiew := qrMacAddressNiew
		if qMacAddressNiew != "" {

			if err := r.SetQueryParam("mac_address__niew", qMacAddressNiew); err != nil {
				return err
			}
		}
	}

	if o.MacAddressNisw != nil {

		// query param mac_address__nisw
		var qrMacAddressNisw string

		if o.MacAddressNisw != nil {
			qrMacAddressNisw = *o.MacAddressNisw
		}
		qMacAddressNisw := qrMacAddressNisw
		if qMacAddressNisw != "" {

			if err := r.SetQueryParam("mac_address__nisw", qMacAddressNisw); err != nil {
				return err
			}
		}
	}

	if o.Memory != nil {

		// query param memory
		var qrMemory string

		if o.Memory != nil {
			qrMemory = *o.Memory
		}
		qMemory := qrMemory
		if qMemory != "" {

			if err := r.SetQueryParam("memory", qMemory); err != nil {
				return err
			}
		}
	}

	if o.MemoryGt != nil {

		// query param memory__gt
		var qrMemoryGt string

		if o.MemoryGt != nil {
			qrMemoryGt = *o.MemoryGt
		}
		qMemoryGt := qrMemoryGt
		if qMemoryGt != "" {

			if err := r.SetQueryParam("memory__gt", qMemoryGt); err != nil {
				return err
			}
		}
	}

	if o.MemoryGte != nil {

		// query param memory__gte
		var qrMemoryGte string

		if o.MemoryGte != nil {
			qrMemoryGte = *o.MemoryGte
		}
		qMemoryGte := qrMemoryGte
		if qMemoryGte != "" {

			if err := r.SetQueryParam("memory__gte", qMemoryGte); err != nil {
				return err
			}
		}
	}

	if o.MemoryLt != nil {

		// query param memory__lt
		var qrMemoryLt string

		if o.MemoryLt != nil {
			qrMemoryLt = *o.MemoryLt
		}
		qMemoryLt := qrMemoryLt
		if qMemoryLt != "" {

			if err := r.SetQueryParam("memory__lt", qMemoryLt); err != nil {
				return err
			}
		}
	}

	if o.MemoryLte != nil {

		// query param memory__lte
		var qrMemoryLte string

		if o.MemoryLte != nil {
			qrMemoryLte = *o.MemoryLte
		}
		qMemoryLte := qrMemoryLte
		if qMemoryLte != "" {

			if err := r.SetQueryParam("memory__lte", qMemoryLte); err != nil {
				return err
			}
		}
	}

	if o.Memoryn != nil {

		// query param memory__n
		var qrMemoryn string

		if o.Memoryn != nil {
			qrMemoryn = *o.Memoryn
		}
		qMemoryn := qrMemoryn
		if qMemoryn != "" {

			if err := r.SetQueryParam("memory__n", qMemoryn); err != nil {
				return err
			}
		}
	}

	if o.Name != nil {

		// query param name
		var qrName string

		if o.Name != nil {
			qrName = *o.Name
		}
		qName := qrName
		if qName != "" {

			if err := r.SetQueryParam("name", qName); err != nil {
				return err
			}
		}
	}

	if o.NameIc != nil {

		// query param name__ic
		var qrNameIc string

		if o.NameIc != nil {
			qrNameIc = *o.NameIc
		}
		qNameIc := qrNameIc
		if qNameIc != "" {

			if err := r.SetQueryParam("name__ic", qNameIc); err != nil {
				return err
			}
		}
	}

	if o.NameIe != nil {

		// query param name__ie
		var qrNameIe string

		if o.NameIe != nil {
			qrNameIe = *o.NameIe
		}
		qNameIe := qrNameIe
		if qNameIe != "" {

			if err := r.SetQueryParam("name__ie", qNameIe); err != nil {
				return err
			}
		}
	}

	if o.NameIew != nil {

		// query param name__iew
		var qrNameIew string

		if o.NameIew != nil {
			qrNameIew = *o.NameIew
		}
		qNameIew := qrNameIew
		if qNameIew != "" {

			if err := r.SetQueryParam("name__iew", qNameIew); err != nil {
				return err
			}
		}
	}

	if o.NameIsw != nil {

		// query param name__isw
		var qrNameIsw string

		if o.NameIsw != nil {
			qrNameIsw = *o.NameIsw
		}
		qNameIsw := qrNameIsw
		if qNameIsw != "" {

			if err := r.SetQueryParam("name__isw", qNameIsw); err != nil {
				return err
			}
		}
	}

	if o.Namen != nil {

		// query param name__n
		var qrNamen string

		if o.Namen != nil {
			qrNamen = *o.Namen
		}
		qNamen := qrNamen
		if qNamen != "" {

			if err := r.SetQueryParam("name__n", qNamen); err != nil {
				return err
			}
		}
	}

	if o.NameNic != nil {

		// query param name__nic
		var qrNameNic string

		if o.NameNic != nil {
			qrNameNic = *o.NameNic
		}
		qNameNic := qrNameNic
		if qNameNic != "" {

			if err := r.SetQueryParam("name__nic", qNameNic); err != nil {
				return err
			}
		}
	}

	if o.NameNie != nil {

		// query param name__nie
		var qrNameNie string

		if o.NameNie != nil {
			qrNameNie = *o.NameNie
		}
		qNameNie := qrNameNie
		if qNameNie != "" {

			if err := r.SetQueryParam("name__nie", qNameNie); err != nil {
				return err
			}
		}
	}

	if o.NameNiew != nil {

		// query param name__niew
		var qrNameNiew string

		if o.NameNiew != nil {
			qrNameNiew = *o.NameNiew
		}
		qNameNiew := qrNameNiew
		if qNameNiew != "" {

			if err := r.SetQueryParam("name__niew", qNameNiew); err != nil {
				return err
			}
		}
	}

	if o.NameNisw != nil {

		// query param name__nisw
		var qrNameNisw string

		if o.NameNisw != nil {
			qrNameNisw = *o.NameNisw
		}
		qNameNisw := qrNameNisw
		if qNameNisw != "" {

			if err := r.SetQueryParam("name__nisw", qNameNisw); err != nil {
				return err
			}
		}
	}

	if o.Offset != nil {

		// query param offset
		var qrOffset int64

		if o.Offset != nil {
			qrOffset = *o.Offset
		}
		qOffset := swag.FormatInt64(qrOffset)
		if qOffset != "" {

			if err := r.SetQueryParam("offset", qOffset); err != nil {
				return err
			}
		}
	}

	if o.Platform != nil {

		// query param platform
		var qrPlatform string

		if o.Platform != nil {
			qrPlatform = *o.Platform
		}
		qPlatform := qrPlatform
		if qPlatform != "" {

			if err := r.SetQueryParam("platform", qPlatform); err != nil {
				return err
			}
		}
	}

	if o.Platformn != nil {

		// query param platform__n
		var qrPlatformn string

		if o.Platformn != nil {
			qrPlatformn = *o.Platformn
		}
		qPlatformn := qrPlatformn
		if qPlatformn != "" {

			if err := r.SetQueryParam("platform__n", qPlatformn); err != nil {
				return err
			}
		}
	}

	if o.PlatformID != nil {

		// query param platform_id
		var qrPlatformID string

		if o.PlatformID != nil {
			qrPlatformID = *o.PlatformID
		}
		qPlatformID := qrPlatformID
		if qPlatformID != "" {

			if err := r.SetQueryParam("platform_id", qPlatformID); err != nil {
				return err
			}
		}
	}

	if o.PlatformIDn != nil {

		// query param platform_id__n
		var qrPlatformIDn string

		if o.PlatformIDn != nil {
			qrPlatformIDn = *o.PlatformIDn
		}
		qPlatformIDn := qrPlatformIDn
		if qPlatformIDn != "" {

			if err := r.SetQueryParam("platform_id__n", qPlatformIDn); err != nil {
				return err
			}
		}
	}

	if o.Q != nil {

		// query param q
		var qrQ string

		if o.Q != nil {
			qrQ = *o.Q
		}
		qQ := qrQ
		if qQ != "" {

			if err := r.SetQueryParam("q", qQ); err != nil {
				return err
			}
		}
	}

	if o.Region != nil {

		// query param region
		var qrRegion string

		if o.Region != nil {
			qrRegion = *o.Region
		}
		qRegion := qrRegion
		if qRegion != "" {

			if err := r.SetQueryParam("region", qRegion); err != nil {
				return err
			}
		}
	}

	if o.Regionn != nil {

		// query param region__n
		var qrRegionn string

		if o.Regionn != nil {
			qrRegionn = *o.Regionn
		}
		qRegionn := qrRegionn
		if qRegionn != "" {

			if err := r.SetQueryParam("region__n", qRegionn); err != nil {
				return err
			}
		}
	}

	if o.RegionID != nil {

		// query param region_id
		var qrRegionID string

		if o.RegionID != nil {
			qrRegionID = *o.RegionID
		}
		qRegionID := qrRegionID
		if qRegionID != "" {

			if err := r.SetQueryParam("region_id", qRegionID); err != nil {
				return err
			}
		}
	}

	if o.RegionIDn != nil {

		// query param region_id__n
		var qrRegionIDn string

		if o.RegionIDn != nil {
			qrRegionIDn = *o.RegionIDn
		}
		qRegionIDn := qrRegionIDn
		if qRegionIDn != "" {

			if err := r.SetQueryParam("region_id__n", qRegionIDn); err != nil {
				return err
			}
		}
	}

	if o.Role != nil {

		// query param role
		var qrRole string

		if o.Role != nil {
			qrRole = *o.Role
		}
		qRole := qrRole
		if qRole != "" {

			if err := r.SetQueryParam("role", qRole); err != nil {
				return err
			}
		}
	}

	if o.Rolen != nil {

		// query param role__n
		var qrRolen string

		if o.Rolen != nil {
			qrRolen = *o.Rolen
		}
		qRolen := qrRolen
		if qRolen != "" {

			if err := r.SetQueryParam("role__n", qRolen); err != nil {
				return err
			}
		}
	}

	if o.RoleID != nil {

		// query param role_id
		var qrRoleID string

		if o.RoleID != nil {
			qrRoleID = *o.RoleID
		}
		qRoleID := qrRoleID
		if qRoleID != "" {

			if err := r.SetQueryParam("role_id", qRoleID); err != nil {
				return err
			}
		}
	}

	if o.RoleIDn != nil {

		// query param role_id__n
		var qrRoleIDn string

		if o.RoleIDn != nil {
			qrRoleIDn = *o.RoleIDn
		}
		qRoleIDn := qrRoleIDn
		if qRoleIDn != "" {

			if err := r.SetQueryParam("role_id__n", qRoleIDn); err != nil {
				return err
			}
		}
	}

	if o.Site != nil {

		// query param site
		var qrSite string

		if o.Site != nil {
			qrSite = *o.Site
		}
		qSite := qrSite
		if qSite != "" {

			if err := r.SetQueryParam("site", qSite); err != nil {
				return err
			}
		}
	}

	if o.Siten != nil {

		// query param site__n
		var qrSiten string

		if o.Siten != nil {
			qrSiten = *o.Siten
		}
		qSiten := qrSiten
		if qSiten != "" {

			if err := r.SetQueryParam("site__n", qSiten); err != nil {
				return err
			}
		}
	}

	if o.SiteID != nil {

		// query param site_id
		var qrSiteID string

		if o.SiteID != nil {
			qrSiteID = *o.SiteID
		}
		qSiteID := qrSiteID
		if qSiteID != "" {

			if err := r.SetQueryParam("site_id", qSiteID); err != nil {
				return err
			}
		}
	}

	if o.SiteIDn != nil {

		// query param site_id__n
		var qrSiteIDn string

		if o.SiteIDn != nil {
			qrSiteIDn = *o.SiteIDn
		}
		qSiteIDn := qrSiteIDn
		if qSiteIDn != "" {

			if err := r.SetQueryParam("site_id__n", qSiteIDn); err != nil {
				return err
			}
		}
	}

	if o.Status != nil {

		// query param status
		var qrStatus string

		if o.Status != nil {
			qrStatus = *o.Status
		}
		qStatus := qrStatus
		if qStatus != "" {

			if err := r.SetQueryParam("status", qStatus); err != nil {
				return err
			}
		}
	}

	if o.Statusn != nil {

		// query param status__n
		var qrStatusn string

		if o.Statusn != nil {
			qrStatusn = *o.Statusn
		}
		qStatusn := qrStatusn
		if qStatusn != "" {

			if err := r.SetQueryParam("status__n", qStatusn); err != nil {
				return err
			}
		}
	}

	if o.Tag != nil {

		// query param tag
		var qrTag string

		if o.Tag != nil {
			qrTag = *o.Tag
		}
		qTag := qrTag
		if qTag != "" {

			if err := r.SetQueryParam("tag", qTag); err != nil {
				return err
			}
		}
	}

	if o.Tagn != nil {

		// query param tag__n
		var qrTagn string

		if o.Tagn != nil {
			qrTagn = *o.Tagn
		}
		qTagn := qrTagn
		if qTagn != "" {

			if err := r.SetQueryParam("tag__n", qTagn); err != nil {
				return err
			}
		}
	}

	if o.Tenant != nil {

		// query param tenant
		var qrTenant string

		if o.Tenant != nil {
			qrTenant = *o.Tenant
		}
		qTenant := qrTenant
		if qTenant != "" {

			if err := r.SetQueryParam("tenant", qTenant); err != nil {
				return err
			}
		}
	}

	if o.Tenantn != nil {

		// query param tenant__n
		var qrTenantn string

		if o.Tenantn != nil {
			qrTenantn = *o.Tenantn
		}
		qTenantn := qrTenantn
		if qTenantn != "" {

			if err := r.SetQueryParam("tenant__n", qTenantn); err != nil {
				return err
			}
		}
	}

	if o.TenantGroup != nil {

		// query param tenant_group
		var qrTenantGroup string

		if o.TenantGroup != nil {
			qrTenantGroup = *o.TenantGroup
		}
		qTenantGroup := qrTenantGroup
		if qTenantGroup != "" {

			if err := r.SetQueryParam("tenant_group", qTenantGroup); err != nil {
				return err
			}
		}
	}

	if o.TenantGroupn != nil {

		// query param tenant_group__n
		var qrTenantGroupn string

		if o.TenantGroupn != nil {
			qrTenantGroupn = *o.TenantGroupn
		}
		qTenantGroupn := qrTenantGroupn
		if qTenantGroupn != "" {

			if err := r.SetQueryParam("tenant_group__n", qTenantGroupn); err != nil {
				return err
			}
		}
	}

	if o.TenantGroupID != nil {

		// query param tenant_group_id
		var qrTenantGroupID string

		if o.TenantGroupID != nil {
			qrTenantGroupID = *o.TenantGroupID
		}
		qTenantGroupID := qrTenantGroupID
		if qTenantGroupID != "" {

			if err := r.SetQueryParam("tenant_group_id", qTenantGroupID); err != nil {
				return err
			}
		}
	}

	if o.TenantGroupIDn != nil {

		// query param tenant_group_id__n
		var qrTenantGroupIDn string

		if o.TenantGroupIDn != nil {
			qrTenantGroupIDn = *o.TenantGroupIDn
		}
		qTenantGroupIDn := qrTenantGroupIDn
		if qTenantGroupIDn != "" {

			if err := r.SetQueryParam("tenant_group_id__n", qTenantGroupIDn); err != nil {
				return err
			}
		}
	}

	if o.TenantID != nil {

		// query param tenant_id
		var qrTenantID string

		if o.TenantID != nil {
			qrTenantID = *o.TenantID
		}
		qTenantID := qrTenantID
		if qTenantID != "" {

			if err := r.SetQueryParam("tenant_id", qTenantID); err != nil {
				return err
			}
		}
	}

	if o.TenantIDn != nil {

		// query param tenant_id__n
		var qrTenantIDn string

		if o.TenantIDn != nil {
			qrTenantIDn = *o.TenantIDn
		}
		qTenantIDn := qrTenantIDn
		if qTenantIDn != "" {

			if err := r.SetQueryParam("tenant_id__n", qTenantIDn); err != nil {
				return err
			}
		}
	}

	if o.Vcpus != nil {

		// query param vcpus
		var qrVcpus string

		if o.Vcpus != nil {
			qrVcpus = *o.Vcpus
		}
		qVcpus := qrVcpus
		if qVcpus != "" {

			if err := r.SetQueryParam("vcpus", qVcpus); err != nil {
				return err
			}
		}
	}

	if o.VcpusGt != nil {

		// query param vcpus__gt
		var qrVcpusGt string

		if o.VcpusGt != nil {
			qrVcpusGt = *o.VcpusGt
		}
		qVcpusGt := qrVcpusGt
		if qVcpusGt != "" {

			if err := r.SetQueryParam("vcpus__gt", qVcpusGt); err != nil {
				return err
			}
		}
	}

	if o.VcpusGte != nil {

		// query param vcpus__gte
		var qrVcpusGte string

		if o.VcpusGte != nil {
			qrVcpusGte = *o.VcpusGte
		}
		qVcpusGte := qrVcpusGte
		if qVcpusGte != "" {

			if err := r.SetQueryParam("vcpus__gte", qVcpusGte); err != nil {
				return err
			}
		}
	}

	if o.VcpusLt != nil {

		// query param vcpus__lt
		var qrVcpusLt string

		if o.VcpusLt != nil {
			qrVcpusLt = *o.VcpusLt
		}
		qVcpusLt := qrVcpusLt
		if qVcpusLt != "" {

			if err := r.SetQueryParam("vcpus__lt", qVcpusLt); err != nil {
				return err
			}
		}
	}

	if o.VcpusLte != nil {

		// query param vcpus__lte
		var qrVcpusLte string

		if o.VcpusLte != nil {
			qrVcpusLte = *o.VcpusLte
		}
		qVcpusLte := qrVcpusLte
		if qVcpusLte != "" {

			if err := r.SetQueryParam("vcpus__lte", qVcpusLte); err != nil {
				return err
			}
		}
	}

	if o.Vcpusn != nil {

		// query param vcpus__n
		var qrVcpusn string

		if o.Vcpusn != nil {
			qrVcpusn = *o.Vcpusn
		}
		qVcpusn := qrVcpusn
		if qVcpusn != "" {

			if err := r.SetQueryParam("vcpus__n", qVcpusn); err != nil {
				return err
			}
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
