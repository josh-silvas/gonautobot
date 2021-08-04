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

// NewIpamVrfsPartialUpdateParams creates a new IpamVrfsPartialUpdateParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewIpamVrfsPartialUpdateParams() *IpamVrfsPartialUpdateParams {
	return &IpamVrfsPartialUpdateParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewIpamVrfsPartialUpdateParamsWithTimeout creates a new IpamVrfsPartialUpdateParams object
// with the ability to set a timeout on a request.
func NewIpamVrfsPartialUpdateParamsWithTimeout(timeout time.Duration) *IpamVrfsPartialUpdateParams {
	return &IpamVrfsPartialUpdateParams{
		timeout: timeout,
	}
}

// NewIpamVrfsPartialUpdateParamsWithContext creates a new IpamVrfsPartialUpdateParams object
// with the ability to set a context for a request.
func NewIpamVrfsPartialUpdateParamsWithContext(ctx context.Context) *IpamVrfsPartialUpdateParams {
	return &IpamVrfsPartialUpdateParams{
		Context: ctx,
	}
}

// NewIpamVrfsPartialUpdateParamsWithHTTPClient creates a new IpamVrfsPartialUpdateParams object
// with the ability to set a custom HTTPClient for a request.
func NewIpamVrfsPartialUpdateParamsWithHTTPClient(client *http.Client) *IpamVrfsPartialUpdateParams {
	return &IpamVrfsPartialUpdateParams{
		HTTPClient: client,
	}
}

/* IpamVrfsPartialUpdateParams contains all the parameters to send to the API endpoint
   for the ipam vrfs partial update operation.

   Typically these are written to a http.Request.
*/
type IpamVrfsPartialUpdateParams struct {

	// Data.
	Data *models.WritableVRF

	/* ID.

	   A UUID string identifying this VRF.

	   Format: uuid
	*/
	ID strfmt.UUID

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the ipam vrfs partial update params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *IpamVrfsPartialUpdateParams) WithDefaults() *IpamVrfsPartialUpdateParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the ipam vrfs partial update params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *IpamVrfsPartialUpdateParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the ipam vrfs partial update params
func (o *IpamVrfsPartialUpdateParams) WithTimeout(timeout time.Duration) *IpamVrfsPartialUpdateParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the ipam vrfs partial update params
func (o *IpamVrfsPartialUpdateParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the ipam vrfs partial update params
func (o *IpamVrfsPartialUpdateParams) WithContext(ctx context.Context) *IpamVrfsPartialUpdateParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the ipam vrfs partial update params
func (o *IpamVrfsPartialUpdateParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the ipam vrfs partial update params
func (o *IpamVrfsPartialUpdateParams) WithHTTPClient(client *http.Client) *IpamVrfsPartialUpdateParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the ipam vrfs partial update params
func (o *IpamVrfsPartialUpdateParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithData adds the data to the ipam vrfs partial update params
func (o *IpamVrfsPartialUpdateParams) WithData(data *models.WritableVRF) *IpamVrfsPartialUpdateParams {
	o.SetData(data)
	return o
}

// SetData adds the data to the ipam vrfs partial update params
func (o *IpamVrfsPartialUpdateParams) SetData(data *models.WritableVRF) {
	o.Data = data
}

// WithID adds the id to the ipam vrfs partial update params
func (o *IpamVrfsPartialUpdateParams) WithID(id strfmt.UUID) *IpamVrfsPartialUpdateParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the ipam vrfs partial update params
func (o *IpamVrfsPartialUpdateParams) SetID(id strfmt.UUID) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *IpamVrfsPartialUpdateParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
