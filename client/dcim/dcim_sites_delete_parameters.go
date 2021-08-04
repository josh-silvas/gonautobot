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

// NewDcimSitesDeleteParams creates a new DcimSitesDeleteParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewDcimSitesDeleteParams() *DcimSitesDeleteParams {
	return &DcimSitesDeleteParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewDcimSitesDeleteParamsWithTimeout creates a new DcimSitesDeleteParams object
// with the ability to set a timeout on a request.
func NewDcimSitesDeleteParamsWithTimeout(timeout time.Duration) *DcimSitesDeleteParams {
	return &DcimSitesDeleteParams{
		timeout: timeout,
	}
}

// NewDcimSitesDeleteParamsWithContext creates a new DcimSitesDeleteParams object
// with the ability to set a context for a request.
func NewDcimSitesDeleteParamsWithContext(ctx context.Context) *DcimSitesDeleteParams {
	return &DcimSitesDeleteParams{
		Context: ctx,
	}
}

// NewDcimSitesDeleteParamsWithHTTPClient creates a new DcimSitesDeleteParams object
// with the ability to set a custom HTTPClient for a request.
func NewDcimSitesDeleteParamsWithHTTPClient(client *http.Client) *DcimSitesDeleteParams {
	return &DcimSitesDeleteParams{
		HTTPClient: client,
	}
}

/* DcimSitesDeleteParams contains all the parameters to send to the API endpoint
   for the dcim sites delete operation.

   Typically these are written to a http.Request.
*/
type DcimSitesDeleteParams struct {

	/* ID.

	   A UUID string identifying this site.

	   Format: uuid
	*/
	ID strfmt.UUID

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the dcim sites delete params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DcimSitesDeleteParams) WithDefaults() *DcimSitesDeleteParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the dcim sites delete params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DcimSitesDeleteParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the dcim sites delete params
func (o *DcimSitesDeleteParams) WithTimeout(timeout time.Duration) *DcimSitesDeleteParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the dcim sites delete params
func (o *DcimSitesDeleteParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the dcim sites delete params
func (o *DcimSitesDeleteParams) WithContext(ctx context.Context) *DcimSitesDeleteParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the dcim sites delete params
func (o *DcimSitesDeleteParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the dcim sites delete params
func (o *DcimSitesDeleteParams) WithHTTPClient(client *http.Client) *DcimSitesDeleteParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the dcim sites delete params
func (o *DcimSitesDeleteParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the dcim sites delete params
func (o *DcimSitesDeleteParams) WithID(id strfmt.UUID) *DcimSitesDeleteParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the dcim sites delete params
func (o *DcimSitesDeleteParams) SetID(id strfmt.UUID) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *DcimSitesDeleteParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
