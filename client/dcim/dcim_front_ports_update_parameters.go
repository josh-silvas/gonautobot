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

// NewDcimFrontPortsUpdateParams creates a new DcimFrontPortsUpdateParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewDcimFrontPortsUpdateParams() *DcimFrontPortsUpdateParams {
	return &DcimFrontPortsUpdateParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewDcimFrontPortsUpdateParamsWithTimeout creates a new DcimFrontPortsUpdateParams object
// with the ability to set a timeout on a request.
func NewDcimFrontPortsUpdateParamsWithTimeout(timeout time.Duration) *DcimFrontPortsUpdateParams {
	return &DcimFrontPortsUpdateParams{
		timeout: timeout,
	}
}

// NewDcimFrontPortsUpdateParamsWithContext creates a new DcimFrontPortsUpdateParams object
// with the ability to set a context for a request.
func NewDcimFrontPortsUpdateParamsWithContext(ctx context.Context) *DcimFrontPortsUpdateParams {
	return &DcimFrontPortsUpdateParams{
		Context: ctx,
	}
}

// NewDcimFrontPortsUpdateParamsWithHTTPClient creates a new DcimFrontPortsUpdateParams object
// with the ability to set a custom HTTPClient for a request.
func NewDcimFrontPortsUpdateParamsWithHTTPClient(client *http.Client) *DcimFrontPortsUpdateParams {
	return &DcimFrontPortsUpdateParams{
		HTTPClient: client,
	}
}

/* DcimFrontPortsUpdateParams contains all the parameters to send to the API endpoint
   for the dcim front ports update operation.

   Typically these are written to a http.Request.
*/
type DcimFrontPortsUpdateParams struct {

	// Data.
	Data *models.WritableFrontPort

	/* ID.

	   A UUID string identifying this front port.

	   Format: uuid
	*/
	ID strfmt.UUID

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the dcim front ports update params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DcimFrontPortsUpdateParams) WithDefaults() *DcimFrontPortsUpdateParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the dcim front ports update params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DcimFrontPortsUpdateParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the dcim front ports update params
func (o *DcimFrontPortsUpdateParams) WithTimeout(timeout time.Duration) *DcimFrontPortsUpdateParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the dcim front ports update params
func (o *DcimFrontPortsUpdateParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the dcim front ports update params
func (o *DcimFrontPortsUpdateParams) WithContext(ctx context.Context) *DcimFrontPortsUpdateParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the dcim front ports update params
func (o *DcimFrontPortsUpdateParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the dcim front ports update params
func (o *DcimFrontPortsUpdateParams) WithHTTPClient(client *http.Client) *DcimFrontPortsUpdateParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the dcim front ports update params
func (o *DcimFrontPortsUpdateParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithData adds the data to the dcim front ports update params
func (o *DcimFrontPortsUpdateParams) WithData(data *models.WritableFrontPort) *DcimFrontPortsUpdateParams {
	o.SetData(data)
	return o
}

// SetData adds the data to the dcim front ports update params
func (o *DcimFrontPortsUpdateParams) SetData(data *models.WritableFrontPort) {
	o.Data = data
}

// WithID adds the id to the dcim front ports update params
func (o *DcimFrontPortsUpdateParams) WithID(id strfmt.UUID) *DcimFrontPortsUpdateParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the dcim front ports update params
func (o *DcimFrontPortsUpdateParams) SetID(id strfmt.UUID) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *DcimFrontPortsUpdateParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error
	if o.Data != nil {
		if err := r.SetBodyParam(o.Data); err != nil {
			return err
		}
	}

	// path param id
	if err := r.SetPathParam("id", o.ID.String()); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
