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

// NewDcimRackGroupsBulkPartialUpdateParams creates a new DcimRackGroupsBulkPartialUpdateParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewDcimRackGroupsBulkPartialUpdateParams() *DcimRackGroupsBulkPartialUpdateParams {
	return &DcimRackGroupsBulkPartialUpdateParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewDcimRackGroupsBulkPartialUpdateParamsWithTimeout creates a new DcimRackGroupsBulkPartialUpdateParams object
// with the ability to set a timeout on a request.
func NewDcimRackGroupsBulkPartialUpdateParamsWithTimeout(timeout time.Duration) *DcimRackGroupsBulkPartialUpdateParams {
	return &DcimRackGroupsBulkPartialUpdateParams{
		timeout: timeout,
	}
}

// NewDcimRackGroupsBulkPartialUpdateParamsWithContext creates a new DcimRackGroupsBulkPartialUpdateParams object
// with the ability to set a context for a request.
func NewDcimRackGroupsBulkPartialUpdateParamsWithContext(ctx context.Context) *DcimRackGroupsBulkPartialUpdateParams {
	return &DcimRackGroupsBulkPartialUpdateParams{
		Context: ctx,
	}
}

// NewDcimRackGroupsBulkPartialUpdateParamsWithHTTPClient creates a new DcimRackGroupsBulkPartialUpdateParams object
// with the ability to set a custom HTTPClient for a request.
func NewDcimRackGroupsBulkPartialUpdateParamsWithHTTPClient(client *http.Client) *DcimRackGroupsBulkPartialUpdateParams {
	return &DcimRackGroupsBulkPartialUpdateParams{
		HTTPClient: client,
	}
}

/* DcimRackGroupsBulkPartialUpdateParams contains all the parameters to send to the API endpoint
   for the dcim rack groups bulk partial update operation.

   Typically these are written to a http.Request.
*/
type DcimRackGroupsBulkPartialUpdateParams struct {

	// Data.
	Data *models.WritableRackGroup

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the dcim rack groups bulk partial update params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DcimRackGroupsBulkPartialUpdateParams) WithDefaults() *DcimRackGroupsBulkPartialUpdateParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the dcim rack groups bulk partial update params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DcimRackGroupsBulkPartialUpdateParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the dcim rack groups bulk partial update params
func (o *DcimRackGroupsBulkPartialUpdateParams) WithTimeout(timeout time.Duration) *DcimRackGroupsBulkPartialUpdateParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the dcim rack groups bulk partial update params
func (o *DcimRackGroupsBulkPartialUpdateParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the dcim rack groups bulk partial update params
func (o *DcimRackGroupsBulkPartialUpdateParams) WithContext(ctx context.Context) *DcimRackGroupsBulkPartialUpdateParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the dcim rack groups bulk partial update params
func (o *DcimRackGroupsBulkPartialUpdateParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the dcim rack groups bulk partial update params
func (o *DcimRackGroupsBulkPartialUpdateParams) WithHTTPClient(client *http.Client) *DcimRackGroupsBulkPartialUpdateParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the dcim rack groups bulk partial update params
func (o *DcimRackGroupsBulkPartialUpdateParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithData adds the data to the dcim rack groups bulk partial update params
func (o *DcimRackGroupsBulkPartialUpdateParams) WithData(data *models.WritableRackGroup) *DcimRackGroupsBulkPartialUpdateParams {
	o.SetData(data)
	return o
}

// SetData adds the data to the dcim rack groups bulk partial update params
func (o *DcimRackGroupsBulkPartialUpdateParams) SetData(data *models.WritableRackGroup) {
	o.Data = data
}

// WriteToRequest writes these params to a swagger request
func (o *DcimRackGroupsBulkPartialUpdateParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
