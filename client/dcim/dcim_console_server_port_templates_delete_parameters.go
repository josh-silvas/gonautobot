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

// NewDcimConsoleServerPortTemplatesDeleteParams creates a new DcimConsoleServerPortTemplatesDeleteParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewDcimConsoleServerPortTemplatesDeleteParams() *DcimConsoleServerPortTemplatesDeleteParams {
	return &DcimConsoleServerPortTemplatesDeleteParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewDcimConsoleServerPortTemplatesDeleteParamsWithTimeout creates a new DcimConsoleServerPortTemplatesDeleteParams object
// with the ability to set a timeout on a request.
func NewDcimConsoleServerPortTemplatesDeleteParamsWithTimeout(timeout time.Duration) *DcimConsoleServerPortTemplatesDeleteParams {
	return &DcimConsoleServerPortTemplatesDeleteParams{
		timeout: timeout,
	}
}

// NewDcimConsoleServerPortTemplatesDeleteParamsWithContext creates a new DcimConsoleServerPortTemplatesDeleteParams object
// with the ability to set a context for a request.
func NewDcimConsoleServerPortTemplatesDeleteParamsWithContext(ctx context.Context) *DcimConsoleServerPortTemplatesDeleteParams {
	return &DcimConsoleServerPortTemplatesDeleteParams{
		Context: ctx,
	}
}

// NewDcimConsoleServerPortTemplatesDeleteParamsWithHTTPClient creates a new DcimConsoleServerPortTemplatesDeleteParams object
// with the ability to set a custom HTTPClient for a request.
func NewDcimConsoleServerPortTemplatesDeleteParamsWithHTTPClient(client *http.Client) *DcimConsoleServerPortTemplatesDeleteParams {
	return &DcimConsoleServerPortTemplatesDeleteParams{
		HTTPClient: client,
	}
}

/* DcimConsoleServerPortTemplatesDeleteParams contains all the parameters to send to the API endpoint
   for the dcim console server port templates delete operation.

   Typically these are written to a http.Request.
*/
type DcimConsoleServerPortTemplatesDeleteParams struct {

	/* ID.

	   A UUID string identifying this console server port template.

	   Format: uuid
	*/
	ID strfmt.UUID

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the dcim console server port templates delete params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DcimConsoleServerPortTemplatesDeleteParams) WithDefaults() *DcimConsoleServerPortTemplatesDeleteParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the dcim console server port templates delete params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DcimConsoleServerPortTemplatesDeleteParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the dcim console server port templates delete params
func (o *DcimConsoleServerPortTemplatesDeleteParams) WithTimeout(timeout time.Duration) *DcimConsoleServerPortTemplatesDeleteParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the dcim console server port templates delete params
func (o *DcimConsoleServerPortTemplatesDeleteParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the dcim console server port templates delete params
func (o *DcimConsoleServerPortTemplatesDeleteParams) WithContext(ctx context.Context) *DcimConsoleServerPortTemplatesDeleteParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the dcim console server port templates delete params
func (o *DcimConsoleServerPortTemplatesDeleteParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the dcim console server port templates delete params
func (o *DcimConsoleServerPortTemplatesDeleteParams) WithHTTPClient(client *http.Client) *DcimConsoleServerPortTemplatesDeleteParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the dcim console server port templates delete params
func (o *DcimConsoleServerPortTemplatesDeleteParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the dcim console server port templates delete params
func (o *DcimConsoleServerPortTemplatesDeleteParams) WithID(id strfmt.UUID) *DcimConsoleServerPortTemplatesDeleteParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the dcim console server port templates delete params
func (o *DcimConsoleServerPortTemplatesDeleteParams) SetID(id strfmt.UUID) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *DcimConsoleServerPortTemplatesDeleteParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
