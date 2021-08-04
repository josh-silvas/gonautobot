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

// NewDcimCablesBulkUpdateParams creates a new DcimCablesBulkUpdateParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewDcimCablesBulkUpdateParams() *DcimCablesBulkUpdateParams {
	return &DcimCablesBulkUpdateParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewDcimCablesBulkUpdateParamsWithTimeout creates a new DcimCablesBulkUpdateParams object
// with the ability to set a timeout on a request.
func NewDcimCablesBulkUpdateParamsWithTimeout(timeout time.Duration) *DcimCablesBulkUpdateParams {
	return &DcimCablesBulkUpdateParams{
		timeout: timeout,
	}
}

// NewDcimCablesBulkUpdateParamsWithContext creates a new DcimCablesBulkUpdateParams object
// with the ability to set a context for a request.
func NewDcimCablesBulkUpdateParamsWithContext(ctx context.Context) *DcimCablesBulkUpdateParams {
	return &DcimCablesBulkUpdateParams{
		Context: ctx,
	}
}

// NewDcimCablesBulkUpdateParamsWithHTTPClient creates a new DcimCablesBulkUpdateParams object
// with the ability to set a custom HTTPClient for a request.
func NewDcimCablesBulkUpdateParamsWithHTTPClient(client *http.Client) *DcimCablesBulkUpdateParams {
	return &DcimCablesBulkUpdateParams{
		HTTPClient: client,
	}
}

/* DcimCablesBulkUpdateParams contains all the parameters to send to the API endpoint
   for the dcim cables bulk update operation.

   Typically these are written to a http.Request.
*/
type DcimCablesBulkUpdateParams struct {

	// Data.
	Data *models.WritableCable

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the dcim cables bulk update params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DcimCablesBulkUpdateParams) WithDefaults() *DcimCablesBulkUpdateParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the dcim cables bulk update params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DcimCablesBulkUpdateParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the dcim cables bulk update params
func (o *DcimCablesBulkUpdateParams) WithTimeout(timeout time.Duration) *DcimCablesBulkUpdateParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the dcim cables bulk update params
func (o *DcimCablesBulkUpdateParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the dcim cables bulk update params
func (o *DcimCablesBulkUpdateParams) WithContext(ctx context.Context) *DcimCablesBulkUpdateParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the dcim cables bulk update params
func (o *DcimCablesBulkUpdateParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the dcim cables bulk update params
func (o *DcimCablesBulkUpdateParams) WithHTTPClient(client *http.Client) *DcimCablesBulkUpdateParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the dcim cables bulk update params
func (o *DcimCablesBulkUpdateParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithData adds the data to the dcim cables bulk update params
func (o *DcimCablesBulkUpdateParams) WithData(data *models.WritableCable) *DcimCablesBulkUpdateParams {
	o.SetData(data)
	return o
}

// SetData adds the data to the dcim cables bulk update params
func (o *DcimCablesBulkUpdateParams) SetData(data *models.WritableCable) {
	o.Data = data
}

// WriteToRequest writes these params to a swagger request
func (o *DcimCablesBulkUpdateParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
