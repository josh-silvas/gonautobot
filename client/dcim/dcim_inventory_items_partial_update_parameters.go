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

// NewDcimInventoryItemsPartialUpdateParams creates a new DcimInventoryItemsPartialUpdateParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewDcimInventoryItemsPartialUpdateParams() *DcimInventoryItemsPartialUpdateParams {
	return &DcimInventoryItemsPartialUpdateParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewDcimInventoryItemsPartialUpdateParamsWithTimeout creates a new DcimInventoryItemsPartialUpdateParams object
// with the ability to set a timeout on a request.
func NewDcimInventoryItemsPartialUpdateParamsWithTimeout(timeout time.Duration) *DcimInventoryItemsPartialUpdateParams {
	return &DcimInventoryItemsPartialUpdateParams{
		timeout: timeout,
	}
}

// NewDcimInventoryItemsPartialUpdateParamsWithContext creates a new DcimInventoryItemsPartialUpdateParams object
// with the ability to set a context for a request.
func NewDcimInventoryItemsPartialUpdateParamsWithContext(ctx context.Context) *DcimInventoryItemsPartialUpdateParams {
	return &DcimInventoryItemsPartialUpdateParams{
		Context: ctx,
	}
}

// NewDcimInventoryItemsPartialUpdateParamsWithHTTPClient creates a new DcimInventoryItemsPartialUpdateParams object
// with the ability to set a custom HTTPClient for a request.
func NewDcimInventoryItemsPartialUpdateParamsWithHTTPClient(client *http.Client) *DcimInventoryItemsPartialUpdateParams {
	return &DcimInventoryItemsPartialUpdateParams{
		HTTPClient: client,
	}
}

/* DcimInventoryItemsPartialUpdateParams contains all the parameters to send to the API endpoint
   for the dcim inventory items partial update operation.

   Typically these are written to a http.Request.
*/
type DcimInventoryItemsPartialUpdateParams struct {

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

// WithDefaults hydrates default values in the dcim inventory items partial update params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DcimInventoryItemsPartialUpdateParams) WithDefaults() *DcimInventoryItemsPartialUpdateParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the dcim inventory items partial update params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DcimInventoryItemsPartialUpdateParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the dcim inventory items partial update params
func (o *DcimInventoryItemsPartialUpdateParams) WithTimeout(timeout time.Duration) *DcimInventoryItemsPartialUpdateParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the dcim inventory items partial update params
func (o *DcimInventoryItemsPartialUpdateParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the dcim inventory items partial update params
func (o *DcimInventoryItemsPartialUpdateParams) WithContext(ctx context.Context) *DcimInventoryItemsPartialUpdateParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the dcim inventory items partial update params
func (o *DcimInventoryItemsPartialUpdateParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the dcim inventory items partial update params
func (o *DcimInventoryItemsPartialUpdateParams) WithHTTPClient(client *http.Client) *DcimInventoryItemsPartialUpdateParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the dcim inventory items partial update params
func (o *DcimInventoryItemsPartialUpdateParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithData adds the data to the dcim inventory items partial update params
func (o *DcimInventoryItemsPartialUpdateParams) WithData(data *models.WritableInventoryItem) *DcimInventoryItemsPartialUpdateParams {
	o.SetData(data)
	return o
}

// SetData adds the data to the dcim inventory items partial update params
func (o *DcimInventoryItemsPartialUpdateParams) SetData(data *models.WritableInventoryItem) {
	o.Data = data
}

// WithID adds the id to the dcim inventory items partial update params
func (o *DcimInventoryItemsPartialUpdateParams) WithID(id strfmt.UUID) *DcimInventoryItemsPartialUpdateParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the dcim inventory items partial update params
func (o *DcimInventoryItemsPartialUpdateParams) SetID(id strfmt.UUID) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *DcimInventoryItemsPartialUpdateParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
