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

// NewDcimRackRolesUpdateParams creates a new DcimRackRolesUpdateParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewDcimRackRolesUpdateParams() *DcimRackRolesUpdateParams {
	return &DcimRackRolesUpdateParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewDcimRackRolesUpdateParamsWithTimeout creates a new DcimRackRolesUpdateParams object
// with the ability to set a timeout on a request.
func NewDcimRackRolesUpdateParamsWithTimeout(timeout time.Duration) *DcimRackRolesUpdateParams {
	return &DcimRackRolesUpdateParams{
		timeout: timeout,
	}
}

// NewDcimRackRolesUpdateParamsWithContext creates a new DcimRackRolesUpdateParams object
// with the ability to set a context for a request.
func NewDcimRackRolesUpdateParamsWithContext(ctx context.Context) *DcimRackRolesUpdateParams {
	return &DcimRackRolesUpdateParams{
		Context: ctx,
	}
}

// NewDcimRackRolesUpdateParamsWithHTTPClient creates a new DcimRackRolesUpdateParams object
// with the ability to set a custom HTTPClient for a request.
func NewDcimRackRolesUpdateParamsWithHTTPClient(client *http.Client) *DcimRackRolesUpdateParams {
	return &DcimRackRolesUpdateParams{
		HTTPClient: client,
	}
}

/* DcimRackRolesUpdateParams contains all the parameters to send to the API endpoint
   for the dcim rack roles update operation.

   Typically these are written to a http.Request.
*/
type DcimRackRolesUpdateParams struct {

	// Data.
	Data *models.RackRole

	/* ID.

	   A UUID string identifying this rack role.

	   Format: uuid
	*/
	ID strfmt.UUID

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the dcim rack roles update params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DcimRackRolesUpdateParams) WithDefaults() *DcimRackRolesUpdateParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the dcim rack roles update params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DcimRackRolesUpdateParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the dcim rack roles update params
func (o *DcimRackRolesUpdateParams) WithTimeout(timeout time.Duration) *DcimRackRolesUpdateParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the dcim rack roles update params
func (o *DcimRackRolesUpdateParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the dcim rack roles update params
func (o *DcimRackRolesUpdateParams) WithContext(ctx context.Context) *DcimRackRolesUpdateParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the dcim rack roles update params
func (o *DcimRackRolesUpdateParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the dcim rack roles update params
func (o *DcimRackRolesUpdateParams) WithHTTPClient(client *http.Client) *DcimRackRolesUpdateParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the dcim rack roles update params
func (o *DcimRackRolesUpdateParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithData adds the data to the dcim rack roles update params
func (o *DcimRackRolesUpdateParams) WithData(data *models.RackRole) *DcimRackRolesUpdateParams {
	o.SetData(data)
	return o
}

// SetData adds the data to the dcim rack roles update params
func (o *DcimRackRolesUpdateParams) SetData(data *models.RackRole) {
	o.Data = data
}

// WithID adds the id to the dcim rack roles update params
func (o *DcimRackRolesUpdateParams) WithID(id strfmt.UUID) *DcimRackRolesUpdateParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the dcim rack roles update params
func (o *DcimRackRolesUpdateParams) SetID(id strfmt.UUID) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *DcimRackRolesUpdateParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
