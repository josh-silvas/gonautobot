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

// NewIpamPrefixesUpdateParams creates a new IpamPrefixesUpdateParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewIpamPrefixesUpdateParams() *IpamPrefixesUpdateParams {
	return &IpamPrefixesUpdateParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewIpamPrefixesUpdateParamsWithTimeout creates a new IpamPrefixesUpdateParams object
// with the ability to set a timeout on a request.
func NewIpamPrefixesUpdateParamsWithTimeout(timeout time.Duration) *IpamPrefixesUpdateParams {
	return &IpamPrefixesUpdateParams{
		timeout: timeout,
	}
}

// NewIpamPrefixesUpdateParamsWithContext creates a new IpamPrefixesUpdateParams object
// with the ability to set a context for a request.
func NewIpamPrefixesUpdateParamsWithContext(ctx context.Context) *IpamPrefixesUpdateParams {
	return &IpamPrefixesUpdateParams{
		Context: ctx,
	}
}

// NewIpamPrefixesUpdateParamsWithHTTPClient creates a new IpamPrefixesUpdateParams object
// with the ability to set a custom HTTPClient for a request.
func NewIpamPrefixesUpdateParamsWithHTTPClient(client *http.Client) *IpamPrefixesUpdateParams {
	return &IpamPrefixesUpdateParams{
		HTTPClient: client,
	}
}

/* IpamPrefixesUpdateParams contains all the parameters to send to the API endpoint
   for the ipam prefixes update operation.

   Typically these are written to a http.Request.
*/
type IpamPrefixesUpdateParams struct {

	// Data.
	Data *models.WritablePrefix

	/* ID.

	   A UUID string identifying this prefix.

	   Format: uuid
	*/
	ID strfmt.UUID

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the ipam prefixes update params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *IpamPrefixesUpdateParams) WithDefaults() *IpamPrefixesUpdateParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the ipam prefixes update params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *IpamPrefixesUpdateParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the ipam prefixes update params
func (o *IpamPrefixesUpdateParams) WithTimeout(timeout time.Duration) *IpamPrefixesUpdateParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the ipam prefixes update params
func (o *IpamPrefixesUpdateParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the ipam prefixes update params
func (o *IpamPrefixesUpdateParams) WithContext(ctx context.Context) *IpamPrefixesUpdateParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the ipam prefixes update params
func (o *IpamPrefixesUpdateParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the ipam prefixes update params
func (o *IpamPrefixesUpdateParams) WithHTTPClient(client *http.Client) *IpamPrefixesUpdateParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the ipam prefixes update params
func (o *IpamPrefixesUpdateParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithData adds the data to the ipam prefixes update params
func (o *IpamPrefixesUpdateParams) WithData(data *models.WritablePrefix) *IpamPrefixesUpdateParams {
	o.SetData(data)
	return o
}

// SetData adds the data to the ipam prefixes update params
func (o *IpamPrefixesUpdateParams) SetData(data *models.WritablePrefix) {
	o.Data = data
}

// WithID adds the id to the ipam prefixes update params
func (o *IpamPrefixesUpdateParams) WithID(id strfmt.UUID) *IpamPrefixesUpdateParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the ipam prefixes update params
func (o *IpamPrefixesUpdateParams) SetID(id strfmt.UUID) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *IpamPrefixesUpdateParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
