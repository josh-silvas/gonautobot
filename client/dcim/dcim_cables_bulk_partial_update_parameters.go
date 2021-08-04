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

// NewDcimCablesBulkPartialUpdateParams creates a new DcimCablesBulkPartialUpdateParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewDcimCablesBulkPartialUpdateParams() *DcimCablesBulkPartialUpdateParams {
	return &DcimCablesBulkPartialUpdateParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewDcimCablesBulkPartialUpdateParamsWithTimeout creates a new DcimCablesBulkPartialUpdateParams object
// with the ability to set a timeout on a request.
func NewDcimCablesBulkPartialUpdateParamsWithTimeout(timeout time.Duration) *DcimCablesBulkPartialUpdateParams {
	return &DcimCablesBulkPartialUpdateParams{
		timeout: timeout,
	}
}

// NewDcimCablesBulkPartialUpdateParamsWithContext creates a new DcimCablesBulkPartialUpdateParams object
// with the ability to set a context for a request.
func NewDcimCablesBulkPartialUpdateParamsWithContext(ctx context.Context) *DcimCablesBulkPartialUpdateParams {
	return &DcimCablesBulkPartialUpdateParams{
		Context: ctx,
	}
}

// NewDcimCablesBulkPartialUpdateParamsWithHTTPClient creates a new DcimCablesBulkPartialUpdateParams object
// with the ability to set a custom HTTPClient for a request.
func NewDcimCablesBulkPartialUpdateParamsWithHTTPClient(client *http.Client) *DcimCablesBulkPartialUpdateParams {
	return &DcimCablesBulkPartialUpdateParams{
		HTTPClient: client,
	}
}

/* DcimCablesBulkPartialUpdateParams contains all the parameters to send to the API endpoint
   for the dcim cables bulk partial update operation.

   Typically these are written to a http.Request.
*/
type DcimCablesBulkPartialUpdateParams struct {

	// Data.
	Data *models.WritableCable

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the dcim cables bulk partial update params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DcimCablesBulkPartialUpdateParams) WithDefaults() *DcimCablesBulkPartialUpdateParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the dcim cables bulk partial update params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DcimCablesBulkPartialUpdateParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the dcim cables bulk partial update params
func (o *DcimCablesBulkPartialUpdateParams) WithTimeout(timeout time.Duration) *DcimCablesBulkPartialUpdateParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the dcim cables bulk partial update params
func (o *DcimCablesBulkPartialUpdateParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the dcim cables bulk partial update params
func (o *DcimCablesBulkPartialUpdateParams) WithContext(ctx context.Context) *DcimCablesBulkPartialUpdateParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the dcim cables bulk partial update params
func (o *DcimCablesBulkPartialUpdateParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the dcim cables bulk partial update params
func (o *DcimCablesBulkPartialUpdateParams) WithHTTPClient(client *http.Client) *DcimCablesBulkPartialUpdateParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the dcim cables bulk partial update params
func (o *DcimCablesBulkPartialUpdateParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithData adds the data to the dcim cables bulk partial update params
func (o *DcimCablesBulkPartialUpdateParams) WithData(data *models.WritableCable) *DcimCablesBulkPartialUpdateParams {
	o.SetData(data)
	return o
}

// SetData adds the data to the dcim cables bulk partial update params
func (o *DcimCablesBulkPartialUpdateParams) SetData(data *models.WritableCable) {
	o.Data = data
}

// WriteToRequest writes these params to a swagger request
func (o *DcimCablesBulkPartialUpdateParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
