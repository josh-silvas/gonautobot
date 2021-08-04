package ipam

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

// NewIpamRirsUpdateParams creates a new IpamRirsUpdateParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewIpamRirsUpdateParams() *IpamRirsUpdateParams {
	return &IpamRirsUpdateParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewIpamRirsUpdateParamsWithTimeout creates a new IpamRirsUpdateParams object
// with the ability to set a timeout on a request.
func NewIpamRirsUpdateParamsWithTimeout(timeout time.Duration) *IpamRirsUpdateParams {
	return &IpamRirsUpdateParams{
		timeout: timeout,
	}
}

// NewIpamRirsUpdateParamsWithContext creates a new IpamRirsUpdateParams object
// with the ability to set a context for a request.
func NewIpamRirsUpdateParamsWithContext(ctx context.Context) *IpamRirsUpdateParams {
	return &IpamRirsUpdateParams{
		Context: ctx,
	}
}

// NewIpamRirsUpdateParamsWithHTTPClient creates a new IpamRirsUpdateParams object
// with the ability to set a custom HTTPClient for a request.
func NewIpamRirsUpdateParamsWithHTTPClient(client *http.Client) *IpamRirsUpdateParams {
	return &IpamRirsUpdateParams{
		HTTPClient: client,
	}
}

/* IpamRirsUpdateParams contains all the parameters to send to the API endpoint
   for the ipam rirs update operation.

   Typically these are written to a http.Request.
*/
type IpamRirsUpdateParams struct {

	// Data.
	Data *models.RIR

	/* ID.

	   A UUID string identifying this RIR.

	   Format: uuid
	*/
	ID strfmt.UUID

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the ipam rirs update params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *IpamRirsUpdateParams) WithDefaults() *IpamRirsUpdateParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the ipam rirs update params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *IpamRirsUpdateParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the ipam rirs update params
func (o *IpamRirsUpdateParams) WithTimeout(timeout time.Duration) *IpamRirsUpdateParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the ipam rirs update params
func (o *IpamRirsUpdateParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the ipam rirs update params
func (o *IpamRirsUpdateParams) WithContext(ctx context.Context) *IpamRirsUpdateParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the ipam rirs update params
func (o *IpamRirsUpdateParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the ipam rirs update params
func (o *IpamRirsUpdateParams) WithHTTPClient(client *http.Client) *IpamRirsUpdateParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the ipam rirs update params
func (o *IpamRirsUpdateParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithData adds the data to the ipam rirs update params
func (o *IpamRirsUpdateParams) WithData(data *models.RIR) *IpamRirsUpdateParams {
	o.SetData(data)
	return o
}

// SetData adds the data to the ipam rirs update params
func (o *IpamRirsUpdateParams) SetData(data *models.RIR) {
	o.Data = data
}

// WithID adds the id to the ipam rirs update params
func (o *IpamRirsUpdateParams) WithID(id strfmt.UUID) *IpamRirsUpdateParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the ipam rirs update params
func (o *IpamRirsUpdateParams) SetID(id strfmt.UUID) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *IpamRirsUpdateParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
