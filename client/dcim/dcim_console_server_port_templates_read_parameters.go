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

// NewDcimConsoleServerPortTemplatesReadParams creates a new DcimConsoleServerPortTemplatesReadParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewDcimConsoleServerPortTemplatesReadParams() *DcimConsoleServerPortTemplatesReadParams {
	return &DcimConsoleServerPortTemplatesReadParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewDcimConsoleServerPortTemplatesReadParamsWithTimeout creates a new DcimConsoleServerPortTemplatesReadParams object
// with the ability to set a timeout on a request.
func NewDcimConsoleServerPortTemplatesReadParamsWithTimeout(timeout time.Duration) *DcimConsoleServerPortTemplatesReadParams {
	return &DcimConsoleServerPortTemplatesReadParams{
		timeout: timeout,
	}
}

// NewDcimConsoleServerPortTemplatesReadParamsWithContext creates a new DcimConsoleServerPortTemplatesReadParams object
// with the ability to set a context for a request.
func NewDcimConsoleServerPortTemplatesReadParamsWithContext(ctx context.Context) *DcimConsoleServerPortTemplatesReadParams {
	return &DcimConsoleServerPortTemplatesReadParams{
		Context: ctx,
	}
}

// NewDcimConsoleServerPortTemplatesReadParamsWithHTTPClient creates a new DcimConsoleServerPortTemplatesReadParams object
// with the ability to set a custom HTTPClient for a request.
func NewDcimConsoleServerPortTemplatesReadParamsWithHTTPClient(client *http.Client) *DcimConsoleServerPortTemplatesReadParams {
	return &DcimConsoleServerPortTemplatesReadParams{
		HTTPClient: client,
	}
}

/* DcimConsoleServerPortTemplatesReadParams contains all the parameters to send to the API endpoint
   for the dcim console server port templates read operation.

   Typically these are written to a http.Request.
*/
type DcimConsoleServerPortTemplatesReadParams struct {

	/* ID.

	   A UUID string identifying this console server port template.

	   Format: uuid
	*/
	ID strfmt.UUID

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the dcim console server port templates read params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DcimConsoleServerPortTemplatesReadParams) WithDefaults() *DcimConsoleServerPortTemplatesReadParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the dcim console server port templates read params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DcimConsoleServerPortTemplatesReadParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the dcim console server port templates read params
func (o *DcimConsoleServerPortTemplatesReadParams) WithTimeout(timeout time.Duration) *DcimConsoleServerPortTemplatesReadParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the dcim console server port templates read params
func (o *DcimConsoleServerPortTemplatesReadParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the dcim console server port templates read params
func (o *DcimConsoleServerPortTemplatesReadParams) WithContext(ctx context.Context) *DcimConsoleServerPortTemplatesReadParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the dcim console server port templates read params
func (o *DcimConsoleServerPortTemplatesReadParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the dcim console server port templates read params
func (o *DcimConsoleServerPortTemplatesReadParams) WithHTTPClient(client *http.Client) *DcimConsoleServerPortTemplatesReadParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the dcim console server port templates read params
func (o *DcimConsoleServerPortTemplatesReadParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the dcim console server port templates read params
func (o *DcimConsoleServerPortTemplatesReadParams) WithID(id strfmt.UUID) *DcimConsoleServerPortTemplatesReadParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the dcim console server port templates read params
func (o *DcimConsoleServerPortTemplatesReadParams) SetID(id strfmt.UUID) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *DcimConsoleServerPortTemplatesReadParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
