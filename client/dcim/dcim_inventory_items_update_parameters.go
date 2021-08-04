package dcim

import (
	"context"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"

	"github.com/josh5276/gonautobot/gonautobot/models"
)

// NewDcimInventoryItemsUpdateParams creates a new DcimInventoryItemsUpdateParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewDcimInventoryItemsUpdateParams() *DcimInventoryItemsUpdateParams {
	return &DcimInventoryItemsUpdateParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewDcimInventoryItemsUpdateParamsWithTimeout creates a new DcimInventoryItemsUpdateParams object
// with the ability to set a timeout on a request.
func NewDcimInventoryItemsUpdateParamsWithTimeout(timeout time.Duration) *DcimInventoryItemsUpdateParams {
	return &DcimInventoryItemsUpdateParams{
		timeout: timeout,
	}
}

// NewDcimInventoryItemsUpdateParamsWithContext creates a new DcimInventoryItemsUpdateParams object
// with the ability to set a context for a request.
func NewDcimInventoryItemsUpdateParamsWithContext(ctx context.Context) *DcimInventoryItemsUpdateParams {
	return &DcimInventoryItemsUpdateParams{
		Context: ctx,
	}
}

// NewDcimInventoryItemsUpdateParamsWithHTTPClient creates a new DcimInventoryItemsUpdateParams object
// with the ability to set a custom HTTPClient for a request.
func NewDcimInventoryItemsUpdateParamsWithHTTPClient(client *http.Client) *DcimInventoryItemsUpdateParams {
	return &DcimInventoryItemsUpdateParams{
		HTTPClient: client,
	}
}

/* DcimInventoryItemsUpdateParams contains all the parameters to send to the API endpoint
   for the dcim inventory items update operation.

   Typically these are written to a http.Request.
*/
type DcimInventoryItemsUpdateParams struct {

	// Data.
	Data *models.WritableInventoryItem

	/* ID.

	   A UUID string identifying this inventory item.

	   Format: uuid
	*/
	ID strfmt.UUID

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the dcim inventory items update params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DcimInventoryItemsUpdateParams) WithDefaults() *DcimInventoryItemsUpdateParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the dcim inventory items update params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DcimInventoryItemsUpdateParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the dcim inventory items update params
func (o *DcimInventoryItemsUpdateParams) WithTimeout(timeout time.Duration) *DcimInventoryItemsUpdateParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the dcim inventory items update params
func (o *DcimInventoryItemsUpdateParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the dcim inventory items update params
func (o *DcimInventoryItemsUpdateParams) WithContext(ctx context.Context) *DcimInventoryItemsUpdateParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the dcim inventory items update params
func (o *DcimInventoryItemsUpdateParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the dcim inventory items update params
func (o *DcimInventoryItemsUpdateParams) WithHTTPClient(client *http.Client) *DcimInventoryItemsUpdateParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the dcim inventory items update params
func (o *DcimInventoryItemsUpdateParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithData adds the data to the dcim inventory items update params
func (o *DcimInventoryItemsUpdateParams) WithData(data *models.WritableInventoryItem) *DcimInventoryItemsUpdateParams {
	o.SetData(data)
	return o
}

// SetData adds the data to the dcim inventory items update params
func (o *DcimInventoryItemsUpdateParams) SetData(data *models.WritableInventoryItem) {
	o.Data = data
}

// WithID adds the id to the dcim inventory items update params
func (o *DcimInventoryItemsUpdateParams) WithID(id strfmt.UUID) *DcimInventoryItemsUpdateParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the dcim inventory items update params
func (o *DcimInventoryItemsUpdateParams) SetID(id strfmt.UUID) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *DcimInventoryItemsUpdateParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error
	if o.Data != nil {
		if err := r.SetBodyParam(o.Data); err != nil {
			return err
		}
	}

	// path param id
	if err := r.SetPathParam("id", o.ID.String()); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
