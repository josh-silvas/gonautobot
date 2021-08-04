package dcim

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

// NewDcimRegionsListParams creates a new DcimRegionsListParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewDcimRegionsListParams() *DcimRegionsListParams {
	return &DcimRegionsListParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewDcimRegionsListParamsWithTimeout creates a new DcimRegionsListParams object
// with the ability to set a timeout on a request.
func NewDcimRegionsListParamsWithTimeout(timeout time.Duration) *DcimRegionsListParams {
	return &DcimRegionsListParams{
		timeout: timeout,
	}
}

// NewDcimRegionsListParamsWithContext creates a new DcimRegionsListParams object
// with the ability to set a context for a request.
func NewDcimRegionsListParamsWithContext(ctx context.Context) *DcimRegionsListParams {
	return &DcimRegionsListParams{
		Context: ctx,
	}
}

// NewDcimRegionsListParamsWithHTTPClient creates a new DcimRegionsListParams object
// with the ability to set a custom HTTPClient for a request.
func NewDcimRegionsListParamsWithHTTPClient(client *http.Client) *DcimRegionsListParams {
	return &DcimRegionsListParams{
		HTTPClient: client,
	}
}

/* DcimRegionsListParams contains all the parameters to send to the API endpoint
   for the dcim regions list operation.

   Typically these are written to a http.Request.
*/
type DcimRegionsListParams struct {

	// Created.
	Created *string

	// CreatedGte.
	CreatedGte *string

	// CreatedLte.
	CreatedLte *string

	// Description.
	Description *string

	// DescriptionIc.
	DescriptionIc *string

	// DescriptionIe.
	DescriptionIe *string

	// DescriptionIew.
	DescriptionIew *string

	// DescriptionIsw.
	DescriptionIsw *string

	// Descriptionn.
	Descriptionn *string

	// DescriptionNic.
	DescriptionNic *string

	// DescriptionNie.
	DescriptionNie *string

	// DescriptionNiew.
	DescriptionNiew *string

	// DescriptionNisw.
	DescriptionNisw *string

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

	// Parent.
	Parent *string

	// Parentn.
	Parentn *string

	// ParentID.
	ParentID *string

	// ParentIDn.
	ParentIDn *string

	// Q.
	Q *string

	// Slug.
	Slug *string

	// SlugIc.
	SlugIc *string

	// SlugIe.
	SlugIe *string

	// SlugIew.
	SlugIew *string

	// SlugIsw.
	SlugIsw *string

	// Slugn.
	Slugn *string

	// SlugNic.
	SlugNic *string

	// SlugNie.
	SlugNie *string

	// SlugNiew.
	SlugNiew *string

	// SlugNisw.
	SlugNisw *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the dcim regions list params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DcimRegionsListParams) WithDefaults() *DcimRegionsListParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the dcim regions list params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DcimRegionsListParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the dcim regions list params
func (o *DcimRegionsListParams) WithTimeout(timeout time.Duration) *DcimRegionsListParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the dcim regions list params
func (o *DcimRegionsListParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the dcim regions list params
func (o *DcimRegionsListParams) WithContext(ctx context.Context) *DcimRegionsListParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the dcim regions list params
func (o *DcimRegionsListParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the dcim regions list params
func (o *DcimRegionsListParams) WithHTTPClient(client *http.Client) *DcimRegionsListParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the dcim regions list params
func (o *DcimRegionsListParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithCreated adds the created to the dcim regions list params
func (o *DcimRegionsListParams) WithCreated(created *string) *DcimRegionsListParams {
	o.SetCreated(created)
	return o
}

// SetCreated adds the created to the dcim regions list params
func (o *DcimRegionsListParams) SetCreated(created *string) {
	o.Created = created
}

// WithCreatedGte adds the createdGte to the dcim regions list params
func (o *DcimRegionsListParams) WithCreatedGte(createdGte *string) *DcimRegionsListParams {
	o.SetCreatedGte(createdGte)
	return o
}

// SetCreatedGte adds the createdGte to the dcim regions list params
func (o *DcimRegionsListParams) SetCreatedGte(createdGte *string) {
	o.CreatedGte = createdGte
}

// WithCreatedLte adds the createdLte to the dcim regions list params
func (o *DcimRegionsListParams) WithCreatedLte(createdLte *string) *DcimRegionsListParams {
	o.SetCreatedLte(createdLte)
	return o
}

// SetCreatedLte adds the createdLte to the dcim regions list params
func (o *DcimRegionsListParams) SetCreatedLte(createdLte *string) {
	o.CreatedLte = createdLte
}

// WithDescription adds the description to the dcim regions list params
func (o *DcimRegionsListParams) WithDescription(description *string) *DcimRegionsListParams {
	o.SetDescription(description)
	return o
}

// SetDescription adds the description to the dcim regions list params
func (o *DcimRegionsListParams) SetDescription(description *string) {
	o.Description = description
}

// WithDescriptionIc adds the descriptionIc to the dcim regions list params
func (o *DcimRegionsListParams) WithDescriptionIc(descriptionIc *string) *DcimRegionsListParams {
	o.SetDescriptionIc(descriptionIc)
	return o
}

// SetDescriptionIc adds the descriptionIc to the dcim regions list params
func (o *DcimRegionsListParams) SetDescriptionIc(descriptionIc *string) {
	o.DescriptionIc = descriptionIc
}

// WithDescriptionIe adds the descriptionIe to the dcim regions list params
func (o *DcimRegionsListParams) WithDescriptionIe(descriptionIe *string) *DcimRegionsListParams {
	o.SetDescriptionIe(descriptionIe)
	return o
}

// SetDescriptionIe adds the descriptionIe to the dcim regions list params
func (o *DcimRegionsListParams) SetDescriptionIe(descriptionIe *string) {
	o.DescriptionIe = descriptionIe
}

// WithDescriptionIew adds the descriptionIew to the dcim regions list params
func (o *DcimRegionsListParams) WithDescriptionIew(descriptionIew *string) *DcimRegionsListParams {
	o.SetDescriptionIew(descriptionIew)
	return o
}

// SetDescriptionIew adds the descriptionIew to the dcim regions list params
func (o *DcimRegionsListParams) SetDescriptionIew(descriptionIew *string) {
	o.DescriptionIew = descriptionIew
}

// WithDescriptionIsw adds the descriptionIsw to the dcim regions list params
func (o *DcimRegionsListParams) WithDescriptionIsw(descriptionIsw *string) *DcimRegionsListParams {
	o.SetDescriptionIsw(descriptionIsw)
	return o
}

// SetDescriptionIsw adds the descriptionIsw to the dcim regions list params
func (o *DcimRegionsListParams) SetDescriptionIsw(descriptionIsw *string) {
	o.DescriptionIsw = descriptionIsw
}

// WithDescriptionn adds the descriptionn to the dcim regions list params
func (o *DcimRegionsListParams) WithDescriptionn(descriptionn *string) *DcimRegionsListParams {
	o.SetDescriptionn(descriptionn)
	return o
}

// SetDescriptionn adds the descriptionN to the dcim regions list params
func (o *DcimRegionsListParams) SetDescriptionn(descriptionn *string) {
	o.Descriptionn = descriptionn
}

// WithDescriptionNic adds the descriptionNic to the dcim regions list params
func (o *DcimRegionsListParams) WithDescriptionNic(descriptionNic *string) *DcimRegionsListParams {
	o.SetDescriptionNic(descriptionNic)
	return o
}

// SetDescriptionNic adds the descriptionNic to the dcim regions list params
func (o *DcimRegionsListParams) SetDescriptionNic(descriptionNic *string) {
	o.DescriptionNic = descriptionNic
}

// WithDescriptionNie adds the descriptionNie to the dcim regions list params
func (o *DcimRegionsListParams) WithDescriptionNie(descriptionNie *string) *DcimRegionsListParams {
	o.SetDescriptionNie(descriptionNie)
	return o
}

// SetDescriptionNie adds the descriptionNie to the dcim regions list params
func (o *DcimRegionsListParams) SetDescriptionNie(descriptionNie *string) {
	o.DescriptionNie = descriptionNie
}

// WithDescriptionNiew adds the descriptionNiew to the dcim regions list params
func (o *DcimRegionsListParams) WithDescriptionNiew(descriptionNiew *string) *DcimRegionsListParams {
	o.SetDescriptionNiew(descriptionNiew)
	return o
}

// SetDescriptionNiew adds the descriptionNiew to the dcim regions list params
func (o *DcimRegionsListParams) SetDescriptionNiew(descriptionNiew *string) {
	o.DescriptionNiew = descriptionNiew
}

// WithDescriptionNisw adds the descriptionNisw to the dcim regions list params
func (o *DcimRegionsListParams) WithDescriptionNisw(descriptionNisw *string) *DcimRegionsListParams {
	o.SetDescriptionNisw(descriptionNisw)
	return o
}

// SetDescriptionNisw adds the descriptionNisw to the dcim regions list params
func (o *DcimRegionsListParams) SetDescriptionNisw(descriptionNisw *string) {
	o.DescriptionNisw = descriptionNisw
}

// WithID adds the id to the dcim regions list params
func (o *DcimRegionsListParams) WithID(id *string) *DcimRegionsListParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the dcim regions list params
func (o *DcimRegionsListParams) SetID(id *string) {
	o.ID = id
}

// WithIDIc adds the iDIc to the dcim regions list params
func (o *DcimRegionsListParams) WithIDIc(iDIc *string) *DcimRegionsListParams {
	o.SetIDIc(iDIc)
	return o
}

// SetIDIc adds the idIc to the dcim regions list params
func (o *DcimRegionsListParams) SetIDIc(iDIc *string) {
	o.IDIc = iDIc
}

// WithIDIe adds the iDIe to the dcim regions list params
func (o *DcimRegionsListParams) WithIDIe(iDIe *string) *DcimRegionsListParams {
	o.SetIDIe(iDIe)
	return o
}

// SetIDIe adds the idIe to the dcim regions list params
func (o *DcimRegionsListParams) SetIDIe(iDIe *string) {
	o.IDIe = iDIe
}

// WithIDIew adds the iDIew to the dcim regions list params
func (o *DcimRegionsListParams) WithIDIew(iDIew *string) *DcimRegionsListParams {
	o.SetIDIew(iDIew)
	return o
}

// SetIDIew adds the idIew to the dcim regions list params
func (o *DcimRegionsListParams) SetIDIew(iDIew *string) {
	o.IDIew = iDIew
}

// WithIDIsw adds the iDIsw to the dcim regions list params
func (o *DcimRegionsListParams) WithIDIsw(iDIsw *string) *DcimRegionsListParams {
	o.SetIDIsw(iDIsw)
	return o
}

// SetIDIsw adds the idIsw to the dcim regions list params
func (o *DcimRegionsListParams) SetIDIsw(iDIsw *string) {
	o.IDIsw = iDIsw
}

// WithIDn adds the iDn to the dcim regions list params
func (o *DcimRegionsListParams) WithIDn(iDn *string) *DcimRegionsListParams {
	o.SetIDn(iDn)
	return o
}

// SetIDn adds the idN to the dcim regions list params
func (o *DcimRegionsListParams) SetIDn(iDn *string) {
	o.IDn = iDn
}

// WithIDNic adds the iDNic to the dcim regions list params
func (o *DcimRegionsListParams) WithIDNic(iDNic *string) *DcimRegionsListParams {
	o.SetIDNic(iDNic)
	return o
}

// SetIDNic adds the idNic to the dcim regions list params
func (o *DcimRegionsListParams) SetIDNic(iDNic *string) {
	o.IDNic = iDNic
}

// WithIDNie adds the iDNie to the dcim regions list params
func (o *DcimRegionsListParams) WithIDNie(iDNie *string) *DcimRegionsListParams {
	o.SetIDNie(iDNie)
	return o
}

// SetIDNie adds the idNie to the dcim regions list params
func (o *DcimRegionsListParams) SetIDNie(iDNie *string) {
	o.IDNie = iDNie
}

// WithIDNiew adds the iDNiew to the dcim regions list params
func (o *DcimRegionsListParams) WithIDNiew(iDNiew *string) *DcimRegionsListParams {
	o.SetIDNiew(iDNiew)
	return o
}

// SetIDNiew adds the idNiew to the dcim regions list params
func (o *DcimRegionsListParams) SetIDNiew(iDNiew *string) {
	o.IDNiew = iDNiew
}

// WithIDNisw adds the iDNisw to the dcim regions list params
func (o *DcimRegionsListParams) WithIDNisw(iDNisw *string) *DcimRegionsListParams {
	o.SetIDNisw(iDNisw)
	return o
}

// SetIDNisw adds the idNisw to the dcim regions list params
func (o *DcimRegionsListParams) SetIDNisw(iDNisw *string) {
	o.IDNisw = iDNisw
}

// WithLastUpdated adds the lastUpdated to the dcim regions list params
func (o *DcimRegionsListParams) WithLastUpdated(lastUpdated *string) *DcimRegionsListParams {
	o.SetLastUpdated(lastUpdated)
	return o
}

// SetLastUpdated adds the lastUpdated to the dcim regions list params
func (o *DcimRegionsListParams) SetLastUpdated(lastUpdated *string) {
	o.LastUpdated = lastUpdated
}

// WithLastUpdatedGte adds the lastUpdatedGte to the dcim regions list params
func (o *DcimRegionsListParams) WithLastUpdatedGte(lastUpdatedGte *string) *DcimRegionsListParams {
	o.SetLastUpdatedGte(lastUpdatedGte)
	return o
}

// SetLastUpdatedGte adds the lastUpdatedGte to the dcim regions list params
func (o *DcimRegionsListParams) SetLastUpdatedGte(lastUpdatedGte *string) {
	o.LastUpdatedGte = lastUpdatedGte
}

// WithLastUpdatedLte adds the lastUpdatedLte to the dcim regions list params
func (o *DcimRegionsListParams) WithLastUpdatedLte(lastUpdatedLte *string) *DcimRegionsListParams {
	o.SetLastUpdatedLte(lastUpdatedLte)
	return o
}

// SetLastUpdatedLte adds the lastUpdatedLte to the dcim regions list params
func (o *DcimRegionsListParams) SetLastUpdatedLte(lastUpdatedLte *string) {
	o.LastUpdatedLte = lastUpdatedLte
}

// WithLimit adds the limit to the dcim regions list params
func (o *DcimRegionsListParams) WithLimit(limit *int64) *DcimRegionsListParams {
	o.SetLimit(limit)
	return o
}

// SetLimit adds the limit to the dcim regions list params
func (o *DcimRegionsListParams) SetLimit(limit *int64) {
	o.Limit = limit
}

// WithName adds the name to the dcim regions list params
func (o *DcimRegionsListParams) WithName(name *string) *DcimRegionsListParams {
	o.SetName(name)
	return o
}

// SetName adds the name to the dcim regions list params
func (o *DcimRegionsListParams) SetName(name *string) {
	o.Name = name
}

// WithNameIc adds the nameIc to the dcim regions list params
func (o *DcimRegionsListParams) WithNameIc(nameIc *string) *DcimRegionsListParams {
	o.SetNameIc(nameIc)
	return o
}

// SetNameIc adds the nameIc to the dcim regions list params
func (o *DcimRegionsListParams) SetNameIc(nameIc *string) {
	o.NameIc = nameIc
}

// WithNameIe adds the nameIe to the dcim regions list params
func (o *DcimRegionsListParams) WithNameIe(nameIe *string) *DcimRegionsListParams {
	o.SetNameIe(nameIe)
	return o
}

// SetNameIe adds the nameIe to the dcim regions list params
func (o *DcimRegionsListParams) SetNameIe(nameIe *string) {
	o.NameIe = nameIe
}

// WithNameIew adds the nameIew to the dcim regions list params
func (o *DcimRegionsListParams) WithNameIew(nameIew *string) *DcimRegionsListParams {
	o.SetNameIew(nameIew)
	return o
}

// SetNameIew adds the nameIew to the dcim regions list params
func (o *DcimRegionsListParams) SetNameIew(nameIew *string) {
	o.NameIew = nameIew
}

// WithNameIsw adds the nameIsw to the dcim regions list params
func (o *DcimRegionsListParams) WithNameIsw(nameIsw *string) *DcimRegionsListParams {
	o.SetNameIsw(nameIsw)
	return o
}

// SetNameIsw adds the nameIsw to the dcim regions list params
func (o *DcimRegionsListParams) SetNameIsw(nameIsw *string) {
	o.NameIsw = nameIsw
}

// WithNamen adds the namen to the dcim regions list params
func (o *DcimRegionsListParams) WithNamen(namen *string) *DcimRegionsListParams {
	o.SetNamen(namen)
	return o
}

// SetNamen adds the nameN to the dcim regions list params
func (o *DcimRegionsListParams) SetNamen(namen *string) {
	o.Namen = namen
}

// WithNameNic adds the nameNic to the dcim regions list params
func (o *DcimRegionsListParams) WithNameNic(nameNic *string) *DcimRegionsListParams {
	o.SetNameNic(nameNic)
	return o
}

// SetNameNic adds the nameNic to the dcim regions list params
func (o *DcimRegionsListParams) SetNameNic(nameNic *string) {
	o.NameNic = nameNic
}

// WithNameNie adds the nameNie to the dcim regions list params
func (o *DcimRegionsListParams) WithNameNie(nameNie *string) *DcimRegionsListParams {
	o.SetNameNie(nameNie)
	return o
}

// SetNameNie adds the nameNie to the dcim regions list params
func (o *DcimRegionsListParams) SetNameNie(nameNie *string) {
	o.NameNie = nameNie
}

// WithNameNiew adds the nameNiew to the dcim regions list params
func (o *DcimRegionsListParams) WithNameNiew(nameNiew *string) *DcimRegionsListParams {
	o.SetNameNiew(nameNiew)
	return o
}

// SetNameNiew adds the nameNiew to the dcim regions list params
func (o *DcimRegionsListParams) SetNameNiew(nameNiew *string) {
	o.NameNiew = nameNiew
}

// WithNameNisw adds the nameNisw to the dcim regions list params
func (o *DcimRegionsListParams) WithNameNisw(nameNisw *string) *DcimRegionsListParams {
	o.SetNameNisw(nameNisw)
	return o
}

// SetNameNisw adds the nameNisw to the dcim regions list params
func (o *DcimRegionsListParams) SetNameNisw(nameNisw *string) {
	o.NameNisw = nameNisw
}

// WithOffset adds the offset to the dcim regions list params
func (o *DcimRegionsListParams) WithOffset(offset *int64) *DcimRegionsListParams {
	o.SetOffset(offset)
	return o
}

// SetOffset adds the offset to the dcim regions list params
func (o *DcimRegionsListParams) SetOffset(offset *int64) {
	o.Offset = offset
}

// WithParent adds the parent to the dcim regions list params
func (o *DcimRegionsListParams) WithParent(parent *string) *DcimRegionsListParams {
	o.SetParent(parent)
	return o
}

// SetParent adds the parent to the dcim regions list params
func (o *DcimRegionsListParams) SetParent(parent *string) {
	o.Parent = parent
}

// WithParentn adds the parentn to the dcim regions list params
func (o *DcimRegionsListParams) WithParentn(parentn *string) *DcimRegionsListParams {
	o.SetParentn(parentn)
	return o
}

// SetParentn adds the parentN to the dcim regions list params
func (o *DcimRegionsListParams) SetParentn(parentn *string) {
	o.Parentn = parentn
}

// WithParentID adds the parentID to the dcim regions list params
func (o *DcimRegionsListParams) WithParentID(parentID *string) *DcimRegionsListParams {
	o.SetParentID(parentID)
	return o
}

// SetParentID adds the parentId to the dcim regions list params
func (o *DcimRegionsListParams) SetParentID(parentID *string) {
	o.ParentID = parentID
}

// WithParentIDn adds the parentIDn to the dcim regions list params
func (o *DcimRegionsListParams) WithParentIDn(parentIDn *string) *DcimRegionsListParams {
	o.SetParentIDn(parentIDn)
	return o
}

// SetParentIDn adds the parentIdN to the dcim regions list params
func (o *DcimRegionsListParams) SetParentIDn(parentIDn *string) {
	o.ParentIDn = parentIDn
}

// WithQ adds the q to the dcim regions list params
func (o *DcimRegionsListParams) WithQ(q *string) *DcimRegionsListParams {
	o.SetQ(q)
	return o
}

// SetQ adds the q to the dcim regions list params
func (o *DcimRegionsListParams) SetQ(q *string) {
	o.Q = q
}

// WithSlug adds the slug to the dcim regions list params
func (o *DcimRegionsListParams) WithSlug(slug *string) *DcimRegionsListParams {
	o.SetSlug(slug)
	return o
}

// SetSlug adds the slug to the dcim regions list params
func (o *DcimRegionsListParams) SetSlug(slug *string) {
	o.Slug = slug
}

// WithSlugIc adds the slugIc to the dcim regions list params
func (o *DcimRegionsListParams) WithSlugIc(slugIc *string) *DcimRegionsListParams {
	o.SetSlugIc(slugIc)
	return o
}

// SetSlugIc adds the slugIc to the dcim regions list params
func (o *DcimRegionsListParams) SetSlugIc(slugIc *string) {
	o.SlugIc = slugIc
}

// WithSlugIe adds the slugIe to the dcim regions list params
func (o *DcimRegionsListParams) WithSlugIe(slugIe *string) *DcimRegionsListParams {
	o.SetSlugIe(slugIe)
	return o
}

// SetSlugIe adds the slugIe to the dcim regions list params
func (o *DcimRegionsListParams) SetSlugIe(slugIe *string) {
	o.SlugIe = slugIe
}

// WithSlugIew adds the slugIew to the dcim regions list params
func (o *DcimRegionsListParams) WithSlugIew(slugIew *string) *DcimRegionsListParams {
	o.SetSlugIew(slugIew)
	return o
}

// SetSlugIew adds the slugIew to the dcim regions list params
func (o *DcimRegionsListParams) SetSlugIew(slugIew *string) {
	o.SlugIew = slugIew
}

// WithSlugIsw adds the slugIsw to the dcim regions list params
func (o *DcimRegionsListParams) WithSlugIsw(slugIsw *string) *DcimRegionsListParams {
	o.SetSlugIsw(slugIsw)
	return o
}

// SetSlugIsw adds the slugIsw to the dcim regions list params
func (o *DcimRegionsListParams) SetSlugIsw(slugIsw *string) {
	o.SlugIsw = slugIsw
}

// WithSlugn adds the slugn to the dcim regions list params
func (o *DcimRegionsListParams) WithSlugn(slugn *string) *DcimRegionsListParams {
	o.SetSlugn(slugn)
	return o
}

// SetSlugn adds the slugN to the dcim regions list params
func (o *DcimRegionsListParams) SetSlugn(slugn *string) {
	o.Slugn = slugn
}

// WithSlugNic adds the slugNic to the dcim regions list params
func (o *DcimRegionsListParams) WithSlugNic(slugNic *string) *DcimRegionsListParams {
	o.SetSlugNic(slugNic)
	return o
}

// SetSlugNic adds the slugNic to the dcim regions list params
func (o *DcimRegionsListParams) SetSlugNic(slugNic *string) {
	o.SlugNic = slugNic
}

// WithSlugNie adds the slugNie to the dcim regions list params
func (o *DcimRegionsListParams) WithSlugNie(slugNie *string) *DcimRegionsListParams {
	o.SetSlugNie(slugNie)
	return o
}

// SetSlugNie adds the slugNie to the dcim regions list params
func (o *DcimRegionsListParams) SetSlugNie(slugNie *string) {
	o.SlugNie = slugNie
}

// WithSlugNiew adds the slugNiew to the dcim regions list params
func (o *DcimRegionsListParams) WithSlugNiew(slugNiew *string) *DcimRegionsListParams {
	o.SetSlugNiew(slugNiew)
	return o
}

// SetSlugNiew adds the slugNiew to the dcim regions list params
func (o *DcimRegionsListParams) SetSlugNiew(slugNiew *string) {
	o.SlugNiew = slugNiew
}

// WithSlugNisw adds the slugNisw to the dcim regions list params
func (o *DcimRegionsListParams) WithSlugNisw(slugNisw *string) *DcimRegionsListParams {
	o.SetSlugNisw(slugNisw)
	return o
}

// SetSlugNisw adds the slugNisw to the dcim regions list params
func (o *DcimRegionsListParams) SetSlugNisw(slugNisw *string) {
	o.SlugNisw = slugNisw
}

// WriteToRequest writes these params to a swagger request
func (o *DcimRegionsListParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

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

	if o.Description != nil {

		// query param description
		var qrDescription string

		if o.Description != nil {
			qrDescription = *o.Description
		}
		qDescription := qrDescription
		if qDescription != "" {

			if err := r.SetQueryParam("description", qDescription); err != nil {
				return err
			}
		}
	}

	if o.DescriptionIc != nil {

		// query param description__ic
		var qrDescriptionIc string

		if o.DescriptionIc != nil {
			qrDescriptionIc = *o.DescriptionIc
		}
		qDescriptionIc := qrDescriptionIc
		if qDescriptionIc != "" {

			if err := r.SetQueryParam("description__ic", qDescriptionIc); err != nil {
				return err
			}
		}
	}

	if o.DescriptionIe != nil {

		// query param description__ie
		var qrDescriptionIe string

		if o.DescriptionIe != nil {
			qrDescriptionIe = *o.DescriptionIe
		}
		qDescriptionIe := qrDescriptionIe
		if qDescriptionIe != "" {

			if err := r.SetQueryParam("description__ie", qDescriptionIe); err != nil {
				return err
			}
		}
	}

	if o.DescriptionIew != nil {

		// query param description__iew
		var qrDescriptionIew string

		if o.DescriptionIew != nil {
			qrDescriptionIew = *o.DescriptionIew
		}
		qDescriptionIew := qrDescriptionIew
		if qDescriptionIew != "" {

			if err := r.SetQueryParam("description__iew", qDescriptionIew); err != nil {
				return err
			}
		}
	}

	if o.DescriptionIsw != nil {

		// query param description__isw
		var qrDescriptionIsw string

		if o.DescriptionIsw != nil {
			qrDescriptionIsw = *o.DescriptionIsw
		}
		qDescriptionIsw := qrDescriptionIsw
		if qDescriptionIsw != "" {

			if err := r.SetQueryParam("description__isw", qDescriptionIsw); err != nil {
				return err
			}
		}
	}

	if o.Descriptionn != nil {

		// query param description__n
		var qrDescriptionn string

		if o.Descriptionn != nil {
			qrDescriptionn = *o.Descriptionn
		}
		qDescriptionn := qrDescriptionn
		if qDescriptionn != "" {

			if err := r.SetQueryParam("description__n", qDescriptionn); err != nil {
				return err
			}
		}
	}

	if o.DescriptionNic != nil {

		// query param description__nic
		var qrDescriptionNic string

		if o.DescriptionNic != nil {
			qrDescriptionNic = *o.DescriptionNic
		}
		qDescriptionNic := qrDescriptionNic
		if qDescriptionNic != "" {

			if err := r.SetQueryParam("description__nic", qDescriptionNic); err != nil {
				return err
			}
		}
	}

	if o.DescriptionNie != nil {

		// query param description__nie
		var qrDescriptionNie string

		if o.DescriptionNie != nil {
			qrDescriptionNie = *o.DescriptionNie
		}
		qDescriptionNie := qrDescriptionNie
		if qDescriptionNie != "" {

			if err := r.SetQueryParam("description__nie", qDescriptionNie); err != nil {
				return err
			}
		}
	}

	if o.DescriptionNiew != nil {

		// query param description__niew
		var qrDescriptionNiew string

		if o.DescriptionNiew != nil {
			qrDescriptionNiew = *o.DescriptionNiew
		}
		qDescriptionNiew := qrDescriptionNiew
		if qDescriptionNiew != "" {

			if err := r.SetQueryParam("description__niew", qDescriptionNiew); err != nil {
				return err
			}
		}
	}

	if o.DescriptionNisw != nil {

		// query param description__nisw
		var qrDescriptionNisw string

		if o.DescriptionNisw != nil {
			qrDescriptionNisw = *o.DescriptionNisw
		}
		qDescriptionNisw := qrDescriptionNisw
		if qDescriptionNisw != "" {

			if err := r.SetQueryParam("description__nisw", qDescriptionNisw); err != nil {
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

	if o.Parent != nil {

		// query param parent
		var qrParent string

		if o.Parent != nil {
			qrParent = *o.Parent
		}
		qParent := qrParent
		if qParent != "" {

			if err := r.SetQueryParam("parent", qParent); err != nil {
				return err
			}
		}
	}

	if o.Parentn != nil {

		// query param parent__n
		var qrParentn string

		if o.Parentn != nil {
			qrParentn = *o.Parentn
		}
		qParentn := qrParentn
		if qParentn != "" {

			if err := r.SetQueryParam("parent__n", qParentn); err != nil {
				return err
			}
		}
	}

	if o.ParentID != nil {

		// query param parent_id
		var qrParentID string

		if o.ParentID != nil {
			qrParentID = *o.ParentID
		}
		qParentID := qrParentID
		if qParentID != "" {

			if err := r.SetQueryParam("parent_id", qParentID); err != nil {
				return err
			}
		}
	}

	if o.ParentIDn != nil {

		// query param parent_id__n
		var qrParentIDn string

		if o.ParentIDn != nil {
			qrParentIDn = *o.ParentIDn
		}
		qParentIDn := qrParentIDn
		if qParentIDn != "" {

			if err := r.SetQueryParam("parent_id__n", qParentIDn); err != nil {
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

	if o.Slug != nil {

		// query param slug
		var qrSlug string

		if o.Slug != nil {
			qrSlug = *o.Slug
		}
		qSlug := qrSlug
		if qSlug != "" {

			if err := r.SetQueryParam("slug", qSlug); err != nil {
				return err
			}
		}
	}

	if o.SlugIc != nil {

		// query param slug__ic
		var qrSlugIc string

		if o.SlugIc != nil {
			qrSlugIc = *o.SlugIc
		}
		qSlugIc := qrSlugIc
		if qSlugIc != "" {

			if err := r.SetQueryParam("slug__ic", qSlugIc); err != nil {
				return err
			}
		}
	}

	if o.SlugIe != nil {

		// query param slug__ie
		var qrSlugIe string

		if o.SlugIe != nil {
			qrSlugIe = *o.SlugIe
		}
		qSlugIe := qrSlugIe
		if qSlugIe != "" {

			if err := r.SetQueryParam("slug__ie", qSlugIe); err != nil {
				return err
			}
		}
	}

	if o.SlugIew != nil {

		// query param slug__iew
		var qrSlugIew string

		if o.SlugIew != nil {
			qrSlugIew = *o.SlugIew
		}
		qSlugIew := qrSlugIew
		if qSlugIew != "" {

			if err := r.SetQueryParam("slug__iew", qSlugIew); err != nil {
				return err
			}
		}
	}

	if o.SlugIsw != nil {

		// query param slug__isw
		var qrSlugIsw string

		if o.SlugIsw != nil {
			qrSlugIsw = *o.SlugIsw
		}
		qSlugIsw := qrSlugIsw
		if qSlugIsw != "" {

			if err := r.SetQueryParam("slug__isw", qSlugIsw); err != nil {
				return err
			}
		}
	}

	if o.Slugn != nil {

		// query param slug__n
		var qrSlugn string

		if o.Slugn != nil {
			qrSlugn = *o.Slugn
		}
		qSlugn := qrSlugn
		if qSlugn != "" {

			if err := r.SetQueryParam("slug__n", qSlugn); err != nil {
				return err
			}
		}
	}

	if o.SlugNic != nil {

		// query param slug__nic
		var qrSlugNic string

		if o.SlugNic != nil {
			qrSlugNic = *o.SlugNic
		}
		qSlugNic := qrSlugNic
		if qSlugNic != "" {

			if err := r.SetQueryParam("slug__nic", qSlugNic); err != nil {
				return err
			}
		}
	}

	if o.SlugNie != nil {

		// query param slug__nie
		var qrSlugNie string

		if o.SlugNie != nil {
			qrSlugNie = *o.SlugNie
		}
		qSlugNie := qrSlugNie
		if qSlugNie != "" {

			if err := r.SetQueryParam("slug__nie", qSlugNie); err != nil {
				return err
			}
		}
	}

	if o.SlugNiew != nil {

		// query param slug__niew
		var qrSlugNiew string

		if o.SlugNiew != nil {
			qrSlugNiew = *o.SlugNiew
		}
		qSlugNiew := qrSlugNiew
		if qSlugNiew != "" {

			if err := r.SetQueryParam("slug__niew", qSlugNiew); err != nil {
				return err
			}
		}
	}

	if o.SlugNisw != nil {

		// query param slug__nisw
		var qrSlugNisw string

		if o.SlugNisw != nil {
			qrSlugNisw = *o.SlugNisw
		}
		qSlugNisw := qrSlugNisw
		if qSlugNisw != "" {

			if err := r.SetQueryParam("slug__nisw", qSlugNisw); err != nil {
				return err
			}
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
