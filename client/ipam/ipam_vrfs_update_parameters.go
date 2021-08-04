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

// NewIpamVrfsUpdateParams creates a new IpamVrfsUpdateParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewIpamVrfsUpdateParams() *IpamVrfsUpdateParams {
	return &IpamVrfsUpdateParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewIpamVrfsUpdateParamsWithTimeout creates a new IpamVrfsUpdateParams object
// with the ability to set a timeout on a request.
func NewIpamVrfsUpdateParamsWithTimeout(timeout time.Duration) *IpamVrfsUpdateParams {
	return &IpamVrfsUpdateParams{
		timeout: timeout,
	}
}

// NewIpamVrfsUpdateParamsWithContext creates a new IpamVrfsUpdateParams object
// with the ability to set a context for a request.
func NewIpamVrfsUpdateParamsWithContext(ctx context.Context) *IpamVrfsUpdateParams {
	return &IpamVrfsUpdateParams{
		Context: ctx,
	}
}

// NewIpamVrfsUpdateParamsWithHTTPClient creates a new IpamVrfsUpdateParams object
// with the ability to set a custom HTTPClient for a request.
func NewIpamVrfsUpdateParamsWithHTTPClient(client *http.Client) *IpamVrfsUpdateParams {
	return &IpamVrfsUpdateParams{
		HTTPClient: client,
	}
}

/* IpamVrfsUpdateParams contains all the parameters to send to the API endpoint
   for the ipam vrfs update operation.

   Typically these are written to a http.Request.
*/
type IpamVrfsUpdateParams struct {

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

// WithDefaults hydrates default values in the ipam vrfs update params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *IpamVrfsUpdateParams) WithDefaults() *IpamVrfsUpdateParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the ipam vrfs update params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *IpamVrfsUpdateParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the ipam vrfs update params
func (o *IpamVrfsUpdateParams) WithTimeout(timeout time.Duration) *IpamVrfsUpdateParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the ipam vrfs update params
func (o *IpamVrfsUpdateParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the ipam vrfs update params
func (o *IpamVrfsUpdateParams) WithContext(ctx context.Context) *IpamVrfsUpdateParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the ipam vrfs update params
func (o *IpamVrfsUpdateParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the ipam vrfs update params
func (o *IpamVrfsUpdateParams) WithHTTPClient(client *http.Client) *IpamVrfsUpdateParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the ipam vrfs update params
func (o *IpamVrfsUpdateParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithData adds the data to the ipam vrfs update params
func (o *IpamVrfsUpdateParams) WithData(data *models.WritableVRF) *IpamVrfsUpdateParams {
	o.SetData(data)
	return o
}

// SetData adds the data to the ipam vrfs update params
func (o *IpamVrfsUpdateParams) SetData(data *models.WritableVRF) {
	o.Data = data
}

// WithID adds the id to the ipam vrfs update params
func (o *IpamVrfsUpdateParams) WithID(id strfmt.UUID) *IpamVrfsUpdateParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the ipam vrfs update params
func (o *IpamVrfsUpdateParams) SetID(id strfmt.UUID) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *IpamVrfsUpdateParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
