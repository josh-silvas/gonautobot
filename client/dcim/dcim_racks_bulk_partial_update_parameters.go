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

// NewDcimRacksBulkPartialUpdateParams creates a new DcimRacksBulkPartialUpdateParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewDcimRacksBulkPartialUpdateParams() *DcimRacksBulkPartialUpdateParams {
	return &DcimRacksBulkPartialUpdateParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewDcimRacksBulkPartialUpdateParamsWithTimeout creates a new DcimRacksBulkPartialUpdateParams object
// with the ability to set a timeout on a request.
func NewDcimRacksBulkPartialUpdateParamsWithTimeout(timeout time.Duration) *DcimRacksBulkPartialUpdateParams {
	return &DcimRacksBulkPartialUpdateParams{
		timeout: timeout,
	}
}

// NewDcimRacksBulkPartialUpdateParamsWithContext creates a new DcimRacksBulkPartialUpdateParams object
// with the ability to set a context for a request.
func NewDcimRacksBulkPartialUpdateParamsWithContext(ctx context.Context) *DcimRacksBulkPartialUpdateParams {
	return &DcimRacksBulkPartialUpdateParams{
		Context: ctx,
	}
}

// NewDcimRacksBulkPartialUpdateParamsWithHTTPClient creates a new DcimRacksBulkPartialUpdateParams object
// with the ability to set a custom HTTPClient for a request.
func NewDcimRacksBulkPartialUpdateParamsWithHTTPClient(client *http.Client) *DcimRacksBulkPartialUpdateParams {
	return &DcimRacksBulkPartialUpdateParams{
		HTTPClient: client,
	}
}

/* DcimRacksBulkPartialUpdateParams contains all the parameters to send to the API endpoint
   for the dcim racks bulk partial update operation.

   Typically these are written to a http.Request.
*/
type DcimRacksBulkPartialUpdateParams struct {

	// Data.
	Data *models.WritableRack

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the dcim racks bulk partial update params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DcimRacksBulkPartialUpdateParams) WithDefaults() *DcimRacksBulkPartialUpdateParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the dcim racks bulk partial update params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DcimRacksBulkPartialUpdateParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the dcim racks bulk partial update params
func (o *DcimRacksBulkPartialUpdateParams) WithTimeout(timeout time.Duration) *DcimRacksBulkPartialUpdateParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the dcim racks bulk partial update params
func (o *DcimRacksBulkPartialUpdateParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the dcim racks bulk partial update params
func (o *DcimRacksBulkPartialUpdateParams) WithContext(ctx context.Context) *DcimRacksBulkPartialUpdateParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the dcim racks bulk partial update params
func (o *DcimRacksBulkPartialUpdateParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the dcim racks bulk partial update params
func (o *DcimRacksBulkPartialUpdateParams) WithHTTPClient(client *http.Client) *DcimRacksBulkPartialUpdateParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the dcim racks bulk partial update params
func (o *DcimRacksBulkPartialUpdateParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithData adds the data to the dcim racks bulk partial update params
func (o *DcimRacksBulkPartialUpdateParams) WithData(data *models.WritableRack) *DcimRacksBulkPartialUpdateParams {
	o.SetData(data)
	return o
}

// SetData adds the data to the dcim racks bulk partial update params
func (o *DcimRacksBulkPartialUpdateParams) SetData(data *models.WritableRack) {
	o.Data = data
}

// WriteToRequest writes these params to a swagger request
func (o *DcimRacksBulkPartialUpdateParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
