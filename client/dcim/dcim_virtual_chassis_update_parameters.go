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

// NewDcimVirtualChassisUpdateParams creates a new DcimVirtualChassisUpdateParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewDcimVirtualChassisUpdateParams() *DcimVirtualChassisUpdateParams {
	return &DcimVirtualChassisUpdateParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewDcimVirtualChassisUpdateParamsWithTimeout creates a new DcimVirtualChassisUpdateParams object
// with the ability to set a timeout on a request.
func NewDcimVirtualChassisUpdateParamsWithTimeout(timeout time.Duration) *DcimVirtualChassisUpdateParams {
	return &DcimVirtualChassisUpdateParams{
		timeout: timeout,
	}
}

// NewDcimVirtualChassisUpdateParamsWithContext creates a new DcimVirtualChassisUpdateParams object
// with the ability to set a context for a request.
func NewDcimVirtualChassisUpdateParamsWithContext(ctx context.Context) *DcimVirtualChassisUpdateParams {
	return &DcimVirtualChassisUpdateParams{
		Context: ctx,
	}
}

// NewDcimVirtualChassisUpdateParamsWithHTTPClient creates a new DcimVirtualChassisUpdateParams object
// with the ability to set a custom HTTPClient for a request.
func NewDcimVirtualChassisUpdateParamsWithHTTPClient(client *http.Client) *DcimVirtualChassisUpdateParams {
	return &DcimVirtualChassisUpdateParams{
		HTTPClient: client,
	}
}

/* DcimVirtualChassisUpdateParams contains all the parameters to send to the API endpoint
   for the dcim virtual chassis update operation.

   Typically these are written to a http.Request.
*/
type DcimVirtualChassisUpdateParams struct {

	// Data.
	Data *models.WritableVirtualChassis

	/* ID.

	   A UUID string identifying this virtual chassis.

	   Format: uuid
	*/
	ID strfmt.UUID

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the dcim virtual chassis update params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DcimVirtualChassisUpdateParams) WithDefaults() *DcimVirtualChassisUpdateParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the dcim virtual chassis update params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DcimVirtualChassisUpdateParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the dcim virtual chassis update params
func (o *DcimVirtualChassisUpdateParams) WithTimeout(timeout time.Duration) *DcimVirtualChassisUpdateParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the dcim virtual chassis update params
func (o *DcimVirtualChassisUpdateParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the dcim virtual chassis update params
func (o *DcimVirtualChassisUpdateParams) WithContext(ctx context.Context) *DcimVirtualChassisUpdateParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the dcim virtual chassis update params
func (o *DcimVirtualChassisUpdateParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the dcim virtual chassis update params
func (o *DcimVirtualChassisUpdateParams) WithHTTPClient(client *http.Client) *DcimVirtualChassisUpdateParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the dcim virtual chassis update params
func (o *DcimVirtualChassisUpdateParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithData adds the data to the dcim virtual chassis update params
func (o *DcimVirtualChassisUpdateParams) WithData(data *models.WritableVirtualChassis) *DcimVirtualChassisUpdateParams {
	o.SetData(data)
	return o
}

// SetData adds the data to the dcim virtual chassis update params
func (o *DcimVirtualChassisUpdateParams) SetData(data *models.WritableVirtualChassis) {
	o.Data = data
}

// WithID adds the id to the dcim virtual chassis update params
func (o *DcimVirtualChassisUpdateParams) WithID(id strfmt.UUID) *DcimVirtualChassisUpdateParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the dcim virtual chassis update params
func (o *DcimVirtualChassisUpdateParams) SetID(id strfmt.UUID) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *DcimVirtualChassisUpdateParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
