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

// NewDcimRearPortTemplatesReadParams creates a new DcimRearPortTemplatesReadParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewDcimRearPortTemplatesReadParams() *DcimRearPortTemplatesReadParams {
	return &DcimRearPortTemplatesReadParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewDcimRearPortTemplatesReadParamsWithTimeout creates a new DcimRearPortTemplatesReadParams object
// with the ability to set a timeout on a request.
func NewDcimRearPortTemplatesReadParamsWithTimeout(timeout time.Duration) *DcimRearPortTemplatesReadParams {
	return &DcimRearPortTemplatesReadParams{
		timeout: timeout,
	}
}

// NewDcimRearPortTemplatesReadParamsWithContext creates a new DcimRearPortTemplatesReadParams object
// with the ability to set a context for a request.
func NewDcimRearPortTemplatesReadParamsWithContext(ctx context.Context) *DcimRearPortTemplatesReadParams {
	return &DcimRearPortTemplatesReadParams{
		Context: ctx,
	}
}

// NewDcimRearPortTemplatesReadParamsWithHTTPClient creates a new DcimRearPortTemplatesReadParams object
// with the ability to set a custom HTTPClient for a request.
func NewDcimRearPortTemplatesReadParamsWithHTTPClient(client *http.Client) *DcimRearPortTemplatesReadParams {
	return &DcimRearPortTemplatesReadParams{
		HTTPClient: client,
	}
}

/* DcimRearPortTemplatesReadParams contains all the parameters to send to the API endpoint
   for the dcim rear port templates read operation.

   Typically these are written to a http.Request.
*/
type DcimRearPortTemplatesReadParams struct {

	/* ID.

	   A UUID string identifying this rear port template.

	   Format: uuid
	*/
	ID strfmt.UUID

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the dcim rear port templates read params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DcimRearPortTemplatesReadParams) WithDefaults() *DcimRearPortTemplatesReadParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the dcim rear port templates read params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DcimRearPortTemplatesReadParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the dcim rear port templates read params
func (o *DcimRearPortTemplatesReadParams) WithTimeout(timeout time.Duration) *DcimRearPortTemplatesReadParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the dcim rear port templates read params
func (o *DcimRearPortTemplatesReadParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the dcim rear port templates read params
func (o *DcimRearPortTemplatesReadParams) WithContext(ctx context.Context) *DcimRearPortTemplatesReadParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the dcim rear port templates read params
func (o *DcimRearPortTemplatesReadParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the dcim rear port templates read params
func (o *DcimRearPortTemplatesReadParams) WithHTTPClient(client *http.Client) *DcimRearPortTemplatesReadParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the dcim rear port templates read params
func (o *DcimRearPortTemplatesReadParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the dcim rear port templates read params
func (o *DcimRearPortTemplatesReadParams) WithID(id strfmt.UUID) *DcimRearPortTemplatesReadParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the dcim rear port templates read params
func (o *DcimRearPortTemplatesReadParams) SetID(id strfmt.UUID) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *DcimRearPortTemplatesReadParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
