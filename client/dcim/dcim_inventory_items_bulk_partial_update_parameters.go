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

// NewDcimInventoryItemsBulkPartialUpdateParams creates a new DcimInventoryItemsBulkPartialUpdateParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewDcimInventoryItemsBulkPartialUpdateParams() *DcimInventoryItemsBulkPartialUpdateParams {
	return &DcimInventoryItemsBulkPartialUpdateParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewDcimInventoryItemsBulkPartialUpdateParamsWithTimeout creates a new DcimInventoryItemsBulkPartialUpdateParams object
// with the ability to set a timeout on a request.
func NewDcimInventoryItemsBulkPartialUpdateParamsWithTimeout(timeout time.Duration) *DcimInventoryItemsBulkPartialUpdateParams {
	return &DcimInventoryItemsBulkPartialUpdateParams{
		timeout: timeout,
	}
}

// NewDcimInventoryItemsBulkPartialUpdateParamsWithContext creates a new DcimInventoryItemsBulkPartialUpdateParams object
// with the ability to set a context for a request.
func NewDcimInventoryItemsBulkPartialUpdateParamsWithContext(ctx context.Context) *DcimInventoryItemsBulkPartialUpdateParams {
	return &DcimInventoryItemsBulkPartialUpdateParams{
		Context: ctx,
	}
}

// NewDcimInventoryItemsBulkPartialUpdateParamsWithHTTPClient creates a new DcimInventoryItemsBulkPartialUpdateParams object
// with the ability to set a custom HTTPClient for a request.
func NewDcimInventoryItemsBulkPartialUpdateParamsWithHTTPClient(client *http.Client) *DcimInventoryItemsBulkPartialUpdateParams {
	return &DcimInventoryItemsBulkPartialUpdateParams{
		HTTPClient: client,
	}
}

/* DcimInventoryItemsBulkPartialUpdateParams contains all the parameters to send to the API endpoint
   for the dcim inventory items bulk partial update operation.

   Typically these are written to a http.Request.
*/
type DcimInventoryItemsBulkPartialUpdateParams struct {

	// Data.
	Data *models.WritableInventoryItem

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the dcim inventory items bulk partial update params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DcimInventoryItemsBulkPartialUpdateParams) WithDefaults() *DcimInventoryItemsBulkPartialUpdateParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the dcim inventory items bulk partial update params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DcimInventoryItemsBulkPartialUpdateParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the dcim inventory items bulk partial update params
func (o *DcimInventoryItemsBulkPartialUpdateParams) WithTimeout(timeout time.Duration) *DcimInventoryItemsBulkPartialUpdateParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the dcim inventory items bulk partial update params
func (o *DcimInventoryItemsBulkPartialUpdateParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the dcim inventory items bulk partial update params
func (o *DcimInventoryItemsBulkPartialUpdateParams) WithContext(ctx context.Context) *DcimInventoryItemsBulkPartialUpdateParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the dcim inventory items bulk partial update params
func (o *DcimInventoryItemsBulkPartialUpdateParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the dcim inventory items bulk partial update params
func (o *DcimInventoryItemsBulkPartialUpdateParams) WithHTTPClient(client *http.Client) *DcimInventoryItemsBulkPartialUpdateParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the dcim inventory items bulk partial update params
func (o *DcimInventoryItemsBulkPartialUpdateParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithData adds the data to the dcim inventory items bulk partial update params
func (o *DcimInventoryItemsBulkPartialUpdateParams) WithData(data *models.WritableInventoryItem) *DcimInventoryItemsBulkPartialUpdateParams {
	o.SetData(data)
	return o
}

// SetData adds the data to the dcim inventory items bulk partial update params
func (o *DcimInventoryItemsBulkPartialUpdateParams) SetData(data *models.WritableInventoryItem) {
	o.Data = data
}

// WriteToRequest writes these params to a swagger request
func (o *DcimInventoryItemsBulkPartialUpdateParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error
	if o.Data != nil {
		if err := r.SetBodyParam(o.Data); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
