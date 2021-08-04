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

// NewIpamPrefixesBulkPartialUpdateParams creates a new IpamPrefixesBulkPartialUpdateParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewIpamPrefixesBulkPartialUpdateParams() *IpamPrefixesBulkPartialUpdateParams {
	return &IpamPrefixesBulkPartialUpdateParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewIpamPrefixesBulkPartialUpdateParamsWithTimeout creates a new IpamPrefixesBulkPartialUpdateParams object
// with the ability to set a timeout on a request.
func NewIpamPrefixesBulkPartialUpdateParamsWithTimeout(timeout time.Duration) *IpamPrefixesBulkPartialUpdateParams {
	return &IpamPrefixesBulkPartialUpdateParams{
		timeout: timeout,
	}
}

// NewIpamPrefixesBulkPartialUpdateParamsWithContext creates a new IpamPrefixesBulkPartialUpdateParams object
// with the ability to set a context for a request.
func NewIpamPrefixesBulkPartialUpdateParamsWithContext(ctx context.Context) *IpamPrefixesBulkPartialUpdateParams {
	return &IpamPrefixesBulkPartialUpdateParams{
		Context: ctx,
	}
}

// NewIpamPrefixesBulkPartialUpdateParamsWithHTTPClient creates a new IpamPrefixesBulkPartialUpdateParams object
// with the ability to set a custom HTTPClient for a request.
func NewIpamPrefixesBulkPartialUpdateParamsWithHTTPClient(client *http.Client) *IpamPrefixesBulkPartialUpdateParams {
	return &IpamPrefixesBulkPartialUpdateParams{
		HTTPClient: client,
	}
}

/* IpamPrefixesBulkPartialUpdateParams contains all the parameters to send to the API endpoint
   for the ipam prefixes bulk partial update operation.

   Typically these are written to a http.Request.
*/
type IpamPrefixesBulkPartialUpdateParams struct {

	// Data.
	Data *models.WritablePrefix

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the ipam prefixes bulk partial update params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *IpamPrefixesBulkPartialUpdateParams) WithDefaults() *IpamPrefixesBulkPartialUpdateParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the ipam prefixes bulk partial update params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *IpamPrefixesBulkPartialUpdateParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the ipam prefixes bulk partial update params
func (o *IpamPrefixesBulkPartialUpdateParams) WithTimeout(timeout time.Duration) *IpamPrefixesBulkPartialUpdateParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the ipam prefixes bulk partial update params
func (o *IpamPrefixesBulkPartialUpdateParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the ipam prefixes bulk partial update params
func (o *IpamPrefixesBulkPartialUpdateParams) WithContext(ctx context.Context) *IpamPrefixesBulkPartialUpdateParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the ipam prefixes bulk partial update params
func (o *IpamPrefixesBulkPartialUpdateParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the ipam prefixes bulk partial update params
func (o *IpamPrefixesBulkPartialUpdateParams) WithHTTPClient(client *http.Client) *IpamPrefixesBulkPartialUpdateParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the ipam prefixes bulk partial update params
func (o *IpamPrefixesBulkPartialUpdateParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithData adds the data to the ipam prefixes bulk partial update params
func (o *IpamPrefixesBulkPartialUpdateParams) WithData(data *models.WritablePrefix) *IpamPrefixesBulkPartialUpdateParams {
	o.SetData(data)
	return o
}

// SetData adds the data to the ipam prefixes bulk partial update params
func (o *IpamPrefixesBulkPartialUpdateParams) SetData(data *models.WritablePrefix) {
	o.Data = data
}

// WriteToRequest writes these params to a swagger request
func (o *IpamPrefixesBulkPartialUpdateParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error
	if o.Data != nil {
		if err := r.SetBodyParam(o.Data); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
