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

// NewDcimRacksBulkUpdateParams creates a new DcimRacksBulkUpdateParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewDcimRacksBulkUpdateParams() *DcimRacksBulkUpdateParams {
	return &DcimRacksBulkUpdateParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewDcimRacksBulkUpdateParamsWithTimeout creates a new DcimRacksBulkUpdateParams object
// with the ability to set a timeout on a request.
func NewDcimRacksBulkUpdateParamsWithTimeout(timeout time.Duration) *DcimRacksBulkUpdateParams {
	return &DcimRacksBulkUpdateParams{
		timeout: timeout,
	}
}

// NewDcimRacksBulkUpdateParamsWithContext creates a new DcimRacksBulkUpdateParams object
// with the ability to set a context for a request.
func NewDcimRacksBulkUpdateParamsWithContext(ctx context.Context) *DcimRacksBulkUpdateParams {
	return &DcimRacksBulkUpdateParams{
		Context: ctx,
	}
}

// NewDcimRacksBulkUpdateParamsWithHTTPClient creates a new DcimRacksBulkUpdateParams object
// with the ability to set a custom HTTPClient for a request.
func NewDcimRacksBulkUpdateParamsWithHTTPClient(client *http.Client) *DcimRacksBulkUpdateParams {
	return &DcimRacksBulkUpdateParams{
		HTTPClient: client,
	}
}

/* DcimRacksBulkUpdateParams contains all the parameters to send to the API endpoint
   for the dcim racks bulk update operation.

   Typically these are written to a http.Request.
*/
type DcimRacksBulkUpdateParams struct {

	// Data.
	Data *models.WritableRack

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the dcim racks bulk update params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DcimRacksBulkUpdateParams) WithDefaults() *DcimRacksBulkUpdateParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the dcim racks bulk update params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DcimRacksBulkUpdateParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the dcim racks bulk update params
func (o *DcimRacksBulkUpdateParams) WithTimeout(timeout time.Duration) *DcimRacksBulkUpdateParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the dcim racks bulk update params
func (o *DcimRacksBulkUpdateParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the dcim racks bulk update params
func (o *DcimRacksBulkUpdateParams) WithContext(ctx context.Context) *DcimRacksBulkUpdateParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the dcim racks bulk update params
func (o *DcimRacksBulkUpdateParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the dcim racks bulk update params
func (o *DcimRacksBulkUpdateParams) WithHTTPClient(client *http.Client) *DcimRacksBulkUpdateParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the dcim racks bulk update params
func (o *DcimRacksBulkUpdateParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithData adds the data to the dcim racks bulk update params
func (o *DcimRacksBulkUpdateParams) WithData(data *models.WritableRack) *DcimRacksBulkUpdateParams {
	o.SetData(data)
	return o
}

// SetData adds the data to the dcim racks bulk update params
func (o *DcimRacksBulkUpdateParams) SetData(data *models.WritableRack) {
	o.Data = data
}

// WriteToRequest writes these params to a swagger request
func (o *DcimRacksBulkUpdateParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
