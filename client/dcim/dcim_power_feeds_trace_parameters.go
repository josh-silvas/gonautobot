package dcim

import (
	"context"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"
)

// NewDcimPowerFeedsTraceParams creates a new DcimPowerFeedsTraceParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewDcimPowerFeedsTraceParams() *DcimPowerFeedsTraceParams {
	return &DcimPowerFeedsTraceParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewDcimPowerFeedsTraceParamsWithTimeout creates a new DcimPowerFeedsTraceParams object
// with the ability to set a timeout on a request.
func NewDcimPowerFeedsTraceParamsWithTimeout(timeout time.Duration) *DcimPowerFeedsTraceParams {
	return &DcimPowerFeedsTraceParams{
		timeout: timeout,
	}
}

// NewDcimPowerFeedsTraceParamsWithContext creates a new DcimPowerFeedsTraceParams object
// with the ability to set a context for a request.
func NewDcimPowerFeedsTraceParamsWithContext(ctx context.Context) *DcimPowerFeedsTraceParams {
	return &DcimPowerFeedsTraceParams{
		Context: ctx,
	}
}

// NewDcimPowerFeedsTraceParamsWithHTTPClient creates a new DcimPowerFeedsTraceParams object
// with the ability to set a custom HTTPClient for a request.
func NewDcimPowerFeedsTraceParamsWithHTTPClient(client *http.Client) *DcimPowerFeedsTraceParams {
	return &DcimPowerFeedsTraceParams{
		HTTPClient: client,
	}
}

/* DcimPowerFeedsTraceParams contains all the parameters to send to the API endpoint
   for the dcim power feeds trace operation.

   Typically these are written to a http.Request.
*/
type DcimPowerFeedsTraceParams struct {

	/* ID.

	   A UUID string identifying this power feed.

	   Format: uuid
	*/
	ID strfmt.UUID

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the dcim power feeds trace params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DcimPowerFeedsTraceParams) WithDefaults() *DcimPowerFeedsTraceParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the dcim power feeds trace params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DcimPowerFeedsTraceParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the dcim power feeds trace params
func (o *DcimPowerFeedsTraceParams) WithTimeout(timeout time.Duration) *DcimPowerFeedsTraceParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the dcim power feeds trace params
func (o *DcimPowerFeedsTraceParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the dcim power feeds trace params
func (o *DcimPowerFeedsTraceParams) WithContext(ctx context.Context) *DcimPowerFeedsTraceParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the dcim power feeds trace params
func (o *DcimPowerFeedsTraceParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the dcim power feeds trace params
func (o *DcimPowerFeedsTraceParams) WithHTTPClient(client *http.Client) *DcimPowerFeedsTraceParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the dcim power feeds trace params
func (o *DcimPowerFeedsTraceParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the dcim power feeds trace params
func (o *DcimPowerFeedsTraceParams) WithID(id strfmt.UUID) *DcimPowerFeedsTraceParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the dcim power feeds trace params
func (o *DcimPowerFeedsTraceParams) SetID(id strfmt.UUID) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *DcimPowerFeedsTraceParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param id
	if err := r.SetPathParam("id", o.ID.String()); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
