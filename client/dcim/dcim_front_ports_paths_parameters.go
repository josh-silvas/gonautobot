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

// NewDcimFrontPortsPathsParams creates a new DcimFrontPortsPathsParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewDcimFrontPortsPathsParams() *DcimFrontPortsPathsParams {
	return &DcimFrontPortsPathsParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewDcimFrontPortsPathsParamsWithTimeout creates a new DcimFrontPortsPathsParams object
// with the ability to set a timeout on a request.
func NewDcimFrontPortsPathsParamsWithTimeout(timeout time.Duration) *DcimFrontPortsPathsParams {
	return &DcimFrontPortsPathsParams{
		timeout: timeout,
	}
}

// NewDcimFrontPortsPathsParamsWithContext creates a new DcimFrontPortsPathsParams object
// with the ability to set a context for a request.
func NewDcimFrontPortsPathsParamsWithContext(ctx context.Context) *DcimFrontPortsPathsParams {
	return &DcimFrontPortsPathsParams{
		Context: ctx,
	}
}

// NewDcimFrontPortsPathsParamsWithHTTPClient creates a new DcimFrontPortsPathsParams object
// with the ability to set a custom HTTPClient for a request.
func NewDcimFrontPortsPathsParamsWithHTTPClient(client *http.Client) *DcimFrontPortsPathsParams {
	return &DcimFrontPortsPathsParams{
		HTTPClient: client,
	}
}

/* DcimFrontPortsPathsParams contains all the parameters to send to the API endpoint
   for the dcim front ports paths operation.

   Typically these are written to a http.Request.
*/
type DcimFrontPortsPathsParams struct {

	/* ID.

	   A UUID string identifying this front port.

	   Format: uuid
	*/
	ID strfmt.UUID

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the dcim front ports paths params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DcimFrontPortsPathsParams) WithDefaults() *DcimFrontPortsPathsParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the dcim front ports paths params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DcimFrontPortsPathsParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the dcim front ports paths params
func (o *DcimFrontPortsPathsParams) WithTimeout(timeout time.Duration) *DcimFrontPortsPathsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the dcim front ports paths params
func (o *DcimFrontPortsPathsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the dcim front ports paths params
func (o *DcimFrontPortsPathsParams) WithContext(ctx context.Context) *DcimFrontPortsPathsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the dcim front ports paths params
func (o *DcimFrontPortsPathsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the dcim front ports paths params
func (o *DcimFrontPortsPathsParams) WithHTTPClient(client *http.Client) *DcimFrontPortsPathsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the dcim front ports paths params
func (o *DcimFrontPortsPathsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the dcim front ports paths params
func (o *DcimFrontPortsPathsParams) WithID(id strfmt.UUID) *DcimFrontPortsPathsParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the dcim front ports paths params
func (o *DcimFrontPortsPathsParams) SetID(id strfmt.UUID) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *DcimFrontPortsPathsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
