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

// NewDcimSitesBulkPartialUpdateParams creates a new DcimSitesBulkPartialUpdateParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewDcimSitesBulkPartialUpdateParams() *DcimSitesBulkPartialUpdateParams {
	return &DcimSitesBulkPartialUpdateParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewDcimSitesBulkPartialUpdateParamsWithTimeout creates a new DcimSitesBulkPartialUpdateParams object
// with the ability to set a timeout on a request.
func NewDcimSitesBulkPartialUpdateParamsWithTimeout(timeout time.Duration) *DcimSitesBulkPartialUpdateParams {
	return &DcimSitesBulkPartialUpdateParams{
		timeout: timeout,
	}
}

// NewDcimSitesBulkPartialUpdateParamsWithContext creates a new DcimSitesBulkPartialUpdateParams object
// with the ability to set a context for a request.
func NewDcimSitesBulkPartialUpdateParamsWithContext(ctx context.Context) *DcimSitesBulkPartialUpdateParams {
	return &DcimSitesBulkPartialUpdateParams{
		Context: ctx,
	}
}

// NewDcimSitesBulkPartialUpdateParamsWithHTTPClient creates a new DcimSitesBulkPartialUpdateParams object
// with the ability to set a custom HTTPClient for a request.
func NewDcimSitesBulkPartialUpdateParamsWithHTTPClient(client *http.Client) *DcimSitesBulkPartialUpdateParams {
	return &DcimSitesBulkPartialUpdateParams{
		HTTPClient: client,
	}
}

/* DcimSitesBulkPartialUpdateParams contains all the parameters to send to the API endpoint
   for the dcim sites bulk partial update operation.

   Typically these are written to a http.Request.
*/
type DcimSitesBulkPartialUpdateParams struct {

	// Data.
	Data *models.WritableSite

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the dcim sites bulk partial update params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DcimSitesBulkPartialUpdateParams) WithDefaults() *DcimSitesBulkPartialUpdateParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the dcim sites bulk partial update params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DcimSitesBulkPartialUpdateParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the dcim sites bulk partial update params
func (o *DcimSitesBulkPartialUpdateParams) WithTimeout(timeout time.Duration) *DcimSitesBulkPartialUpdateParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the dcim sites bulk partial update params
func (o *DcimSitesBulkPartialUpdateParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the dcim sites bulk partial update params
func (o *DcimSitesBulkPartialUpdateParams) WithContext(ctx context.Context) *DcimSitesBulkPartialUpdateParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the dcim sites bulk partial update params
func (o *DcimSitesBulkPartialUpdateParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the dcim sites bulk partial update params
func (o *DcimSitesBulkPartialUpdateParams) WithHTTPClient(client *http.Client) *DcimSitesBulkPartialUpdateParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the dcim sites bulk partial update params
func (o *DcimSitesBulkPartialUpdateParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithData adds the data to the dcim sites bulk partial update params
func (o *DcimSitesBulkPartialUpdateParams) WithData(data *models.WritableSite) *DcimSitesBulkPartialUpdateParams {
	o.SetData(data)
	return o
}

// SetData adds the data to the dcim sites bulk partial update params
func (o *DcimSitesBulkPartialUpdateParams) SetData(data *models.WritableSite) {
	o.Data = data
}

// WriteToRequest writes these params to a swagger request
func (o *DcimSitesBulkPartialUpdateParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
