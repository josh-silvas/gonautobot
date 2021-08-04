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

// NewDcimRearPortsPathsParams creates a new DcimRearPortsPathsParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewDcimRearPortsPathsParams() *DcimRearPortsPathsParams {
	return &DcimRearPortsPathsParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewDcimRearPortsPathsParamsWithTimeout creates a new DcimRearPortsPathsParams object
// with the ability to set a timeout on a request.
func NewDcimRearPortsPathsParamsWithTimeout(timeout time.Duration) *DcimRearPortsPathsParams {
	return &DcimRearPortsPathsParams{
		timeout: timeout,
	}
}

// NewDcimRearPortsPathsParamsWithContext creates a new DcimRearPortsPathsParams object
// with the ability to set a context for a request.
func NewDcimRearPortsPathsParamsWithContext(ctx context.Context) *DcimRearPortsPathsParams {
	return &DcimRearPortsPathsParams{
		Context: ctx,
	}
}

// NewDcimRearPortsPathsParamsWithHTTPClient creates a new DcimRearPortsPathsParams object
// with the ability to set a custom HTTPClient for a request.
func NewDcimRearPortsPathsParamsWithHTTPClient(client *http.Client) *DcimRearPortsPathsParams {
	return &DcimRearPortsPathsParams{
		HTTPClient: client,
	}
}

/* DcimRearPortsPathsParams contains all the parameters to send to the API endpoint
   for the dcim rear ports paths operation.

   Typically these are written to a http.Request.
*/
type DcimRearPortsPathsParams struct {

	/* ID.

	   A UUID string identifying this rear port.

	   Format: uuid
	*/
	ID strfmt.UUID

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the dcim rear ports paths params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DcimRearPortsPathsParams) WithDefaults() *DcimRearPortsPathsParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the dcim rear ports paths params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DcimRearPortsPathsParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the dcim rear ports paths params
func (o *DcimRearPortsPathsParams) WithTimeout(timeout time.Duration) *DcimRearPortsPathsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the dcim rear ports paths params
func (o *DcimRearPortsPathsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the dcim rear ports paths params
func (o *DcimRearPortsPathsParams) WithContext(ctx context.Context) *DcimRearPortsPathsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the dcim rear ports paths params
func (o *DcimRearPortsPathsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the dcim rear ports paths params
func (o *DcimRearPortsPathsParams) WithHTTPClient(client *http.Client) *DcimRearPortsPathsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the dcim rear ports paths params
func (o *DcimRearPortsPathsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the dcim rear ports paths params
func (o *DcimRearPortsPathsParams) WithID(id strfmt.UUID) *DcimRearPortsPathsParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the dcim rear ports paths params
func (o *DcimRearPortsPathsParams) SetID(id strfmt.UUID) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *DcimRearPortsPathsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
