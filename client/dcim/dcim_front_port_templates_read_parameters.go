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

// NewDcimFrontPortTemplatesReadParams creates a new DcimFrontPortTemplatesReadParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewDcimFrontPortTemplatesReadParams() *DcimFrontPortTemplatesReadParams {
	return &DcimFrontPortTemplatesReadParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewDcimFrontPortTemplatesReadParamsWithTimeout creates a new DcimFrontPortTemplatesReadParams object
// with the ability to set a timeout on a request.
func NewDcimFrontPortTemplatesReadParamsWithTimeout(timeout time.Duration) *DcimFrontPortTemplatesReadParams {
	return &DcimFrontPortTemplatesReadParams{
		timeout: timeout,
	}
}

// NewDcimFrontPortTemplatesReadParamsWithContext creates a new DcimFrontPortTemplatesReadParams object
// with the ability to set a context for a request.
func NewDcimFrontPortTemplatesReadParamsWithContext(ctx context.Context) *DcimFrontPortTemplatesReadParams {
	return &DcimFrontPortTemplatesReadParams{
		Context: ctx,
	}
}

// NewDcimFrontPortTemplatesReadParamsWithHTTPClient creates a new DcimFrontPortTemplatesReadParams object
// with the ability to set a custom HTTPClient for a request.
func NewDcimFrontPortTemplatesReadParamsWithHTTPClient(client *http.Client) *DcimFrontPortTemplatesReadParams {
	return &DcimFrontPortTemplatesReadParams{
		HTTPClient: client,
	}
}

/* DcimFrontPortTemplatesReadParams contains all the parameters to send to the API endpoint
   for the dcim front port templates read operation.

   Typically these are written to a http.Request.
*/
type DcimFrontPortTemplatesReadParams struct {

	/* ID.

	   A UUID string identifying this front port template.

	   Format: uuid
	*/
	ID strfmt.UUID

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the dcim front port templates read params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DcimFrontPortTemplatesReadParams) WithDefaults() *DcimFrontPortTemplatesReadParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the dcim front port templates read params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DcimFrontPortTemplatesReadParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the dcim front port templates read params
func (o *DcimFrontPortTemplatesReadParams) WithTimeout(timeout time.Duration) *DcimFrontPortTemplatesReadParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the dcim front port templates read params
func (o *DcimFrontPortTemplatesReadParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the dcim front port templates read params
func (o *DcimFrontPortTemplatesReadParams) WithContext(ctx context.Context) *DcimFrontPortTemplatesReadParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the dcim front port templates read params
func (o *DcimFrontPortTemplatesReadParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the dcim front port templates read params
func (o *DcimFrontPortTemplatesReadParams) WithHTTPClient(client *http.Client) *DcimFrontPortTemplatesReadParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the dcim front port templates read params
func (o *DcimFrontPortTemplatesReadParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the dcim front port templates read params
func (o *DcimFrontPortTemplatesReadParams) WithID(id strfmt.UUID) *DcimFrontPortTemplatesReadParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the dcim front port templates read params
func (o *DcimFrontPortTemplatesReadParams) SetID(id strfmt.UUID) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *DcimFrontPortTemplatesReadParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
