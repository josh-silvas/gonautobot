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

// NewDcimRearPortsUpdateParams creates a new DcimRearPortsUpdateParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewDcimRearPortsUpdateParams() *DcimRearPortsUpdateParams {
	return &DcimRearPortsUpdateParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewDcimRearPortsUpdateParamsWithTimeout creates a new DcimRearPortsUpdateParams object
// with the ability to set a timeout on a request.
func NewDcimRearPortsUpdateParamsWithTimeout(timeout time.Duration) *DcimRearPortsUpdateParams {
	return &DcimRearPortsUpdateParams{
		timeout: timeout,
	}
}

// NewDcimRearPortsUpdateParamsWithContext creates a new DcimRearPortsUpdateParams object
// with the ability to set a context for a request.
func NewDcimRearPortsUpdateParamsWithContext(ctx context.Context) *DcimRearPortsUpdateParams {
	return &DcimRearPortsUpdateParams{
		Context: ctx,
	}
}

// NewDcimRearPortsUpdateParamsWithHTTPClient creates a new DcimRearPortsUpdateParams object
// with the ability to set a custom HTTPClient for a request.
func NewDcimRearPortsUpdateParamsWithHTTPClient(client *http.Client) *DcimRearPortsUpdateParams {
	return &DcimRearPortsUpdateParams{
		HTTPClient: client,
	}
}

/* DcimRearPortsUpdateParams contains all the parameters to send to the API endpoint
   for the dcim rear ports update operation.

   Typically these are written to a http.Request.
*/
type DcimRearPortsUpdateParams struct {

	// Data.
	Data *models.WritableRearPort

	/* ID.

	   A UUID string identifying this rear port.

	   Format: uuid
	*/
	ID strfmt.UUID

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the dcim rear ports update params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DcimRearPortsUpdateParams) WithDefaults() *DcimRearPortsUpdateParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the dcim rear ports update params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DcimRearPortsUpdateParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the dcim rear ports update params
func (o *DcimRearPortsUpdateParams) WithTimeout(timeout time.Duration) *DcimRearPortsUpdateParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the dcim rear ports update params
func (o *DcimRearPortsUpdateParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the dcim rear ports update params
func (o *DcimRearPortsUpdateParams) WithContext(ctx context.Context) *DcimRearPortsUpdateParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the dcim rear ports update params
func (o *DcimRearPortsUpdateParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the dcim rear ports update params
func (o *DcimRearPortsUpdateParams) WithHTTPClient(client *http.Client) *DcimRearPortsUpdateParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the dcim rear ports update params
func (o *DcimRearPortsUpdateParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithData adds the data to the dcim rear ports update params
func (o *DcimRearPortsUpdateParams) WithData(data *models.WritableRearPort) *DcimRearPortsUpdateParams {
	o.SetData(data)
	return o
}

// SetData adds the data to the dcim rear ports update params
func (o *DcimRearPortsUpdateParams) SetData(data *models.WritableRearPort) {
	o.Data = data
}

// WithID adds the id to the dcim rear ports update params
func (o *DcimRearPortsUpdateParams) WithID(id strfmt.UUID) *DcimRearPortsUpdateParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the dcim rear ports update params
func (o *DcimRearPortsUpdateParams) SetID(id strfmt.UUID) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *DcimRearPortsUpdateParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
