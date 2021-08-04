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

// NewDcimSitesReadParams creates a new DcimSitesReadParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewDcimSitesReadParams() *DcimSitesReadParams {
	return &DcimSitesReadParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewDcimSitesReadParamsWithTimeout creates a new DcimSitesReadParams object
// with the ability to set a timeout on a request.
func NewDcimSitesReadParamsWithTimeout(timeout time.Duration) *DcimSitesReadParams {
	return &DcimSitesReadParams{
		timeout: timeout,
	}
}

// NewDcimSitesReadParamsWithContext creates a new DcimSitesReadParams object
// with the ability to set a context for a request.
func NewDcimSitesReadParamsWithContext(ctx context.Context) *DcimSitesReadParams {
	return &DcimSitesReadParams{
		Context: ctx,
	}
}

// NewDcimSitesReadParamsWithHTTPClient creates a new DcimSitesReadParams object
// with the ability to set a custom HTTPClient for a request.
func NewDcimSitesReadParamsWithHTTPClient(client *http.Client) *DcimSitesReadParams {
	return &DcimSitesReadParams{
		HTTPClient: client,
	}
}

/* DcimSitesReadParams contains all the parameters to send to the API endpoint
   for the dcim sites read operation.

   Typically these are written to a http.Request.
*/
type DcimSitesReadParams struct {

	/* ID.

	   A UUID string identifying this site.

	   Format: uuid
	*/
	ID strfmt.UUID

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the dcim sites read params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DcimSitesReadParams) WithDefaults() *DcimSitesReadParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the dcim sites read params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DcimSitesReadParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the dcim sites read params
func (o *DcimSitesReadParams) WithTimeout(timeout time.Duration) *DcimSitesReadParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the dcim sites read params
func (o *DcimSitesReadParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the dcim sites read params
func (o *DcimSitesReadParams) WithContext(ctx context.Context) *DcimSitesReadParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the dcim sites read params
func (o *DcimSitesReadParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the dcim sites read params
func (o *DcimSitesReadParams) WithHTTPClient(client *http.Client) *DcimSitesReadParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the dcim sites read params
func (o *DcimSitesReadParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the dcim sites read params
func (o *DcimSitesReadParams) WithID(id strfmt.UUID) *DcimSitesReadParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the dcim sites read params
func (o *DcimSitesReadParams) SetID(id strfmt.UUID) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *DcimSitesReadParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
