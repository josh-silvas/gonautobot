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

// NewDcimRearPortsBulkPartialUpdateParams creates a new DcimRearPortsBulkPartialUpdateParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewDcimRearPortsBulkPartialUpdateParams() *DcimRearPortsBulkPartialUpdateParams {
	return &DcimRearPortsBulkPartialUpdateParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewDcimRearPortsBulkPartialUpdateParamsWithTimeout creates a new DcimRearPortsBulkPartialUpdateParams object
// with the ability to set a timeout on a request.
func NewDcimRearPortsBulkPartialUpdateParamsWithTimeout(timeout time.Duration) *DcimRearPortsBulkPartialUpdateParams {
	return &DcimRearPortsBulkPartialUpdateParams{
		timeout: timeout,
	}
}

// NewDcimRearPortsBulkPartialUpdateParamsWithContext creates a new DcimRearPortsBulkPartialUpdateParams object
// with the ability to set a context for a request.
func NewDcimRearPortsBulkPartialUpdateParamsWithContext(ctx context.Context) *DcimRearPortsBulkPartialUpdateParams {
	return &DcimRearPortsBulkPartialUpdateParams{
		Context: ctx,
	}
}

// NewDcimRearPortsBulkPartialUpdateParamsWithHTTPClient creates a new DcimRearPortsBulkPartialUpdateParams object
// with the ability to set a custom HTTPClient for a request.
func NewDcimRearPortsBulkPartialUpdateParamsWithHTTPClient(client *http.Client) *DcimRearPortsBulkPartialUpdateParams {
	return &DcimRearPortsBulkPartialUpdateParams{
		HTTPClient: client,
	}
}

/* DcimRearPortsBulkPartialUpdateParams contains all the parameters to send to the API endpoint
   for the dcim rear ports bulk partial update operation.

   Typically these are written to a http.Request.
*/
type DcimRearPortsBulkPartialUpdateParams struct {

	// Data.
	Data *models.WritableRearPort

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the dcim rear ports bulk partial update params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DcimRearPortsBulkPartialUpdateParams) WithDefaults() *DcimRearPortsBulkPartialUpdateParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the dcim rear ports bulk partial update params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DcimRearPortsBulkPartialUpdateParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the dcim rear ports bulk partial update params
func (o *DcimRearPortsBulkPartialUpdateParams) WithTimeout(timeout time.Duration) *DcimRearPortsBulkPartialUpdateParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the dcim rear ports bulk partial update params
func (o *DcimRearPortsBulkPartialUpdateParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the dcim rear ports bulk partial update params
func (o *DcimRearPortsBulkPartialUpdateParams) WithContext(ctx context.Context) *DcimRearPortsBulkPartialUpdateParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the dcim rear ports bulk partial update params
func (o *DcimRearPortsBulkPartialUpdateParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the dcim rear ports bulk partial update params
func (o *DcimRearPortsBulkPartialUpdateParams) WithHTTPClient(client *http.Client) *DcimRearPortsBulkPartialUpdateParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the dcim rear ports bulk partial update params
func (o *DcimRearPortsBulkPartialUpdateParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithData adds the data to the dcim rear ports bulk partial update params
func (o *DcimRearPortsBulkPartialUpdateParams) WithData(data *models.WritableRearPort) *DcimRearPortsBulkPartialUpdateParams {
	o.SetData(data)
	return o
}

// SetData adds the data to the dcim rear ports bulk partial update params
func (o *DcimRearPortsBulkPartialUpdateParams) SetData(data *models.WritableRearPort) {
	o.Data = data
}

// WriteToRequest writes these params to a swagger request
func (o *DcimRearPortsBulkPartialUpdateParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
