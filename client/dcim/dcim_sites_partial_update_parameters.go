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

// NewDcimSitesPartialUpdateParams creates a new DcimSitesPartialUpdateParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewDcimSitesPartialUpdateParams() *DcimSitesPartialUpdateParams {
	return &DcimSitesPartialUpdateParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewDcimSitesPartialUpdateParamsWithTimeout creates a new DcimSitesPartialUpdateParams object
// with the ability to set a timeout on a request.
func NewDcimSitesPartialUpdateParamsWithTimeout(timeout time.Duration) *DcimSitesPartialUpdateParams {
	return &DcimSitesPartialUpdateParams{
		timeout: timeout,
	}
}

// NewDcimSitesPartialUpdateParamsWithContext creates a new DcimSitesPartialUpdateParams object
// with the ability to set a context for a request.
func NewDcimSitesPartialUpdateParamsWithContext(ctx context.Context) *DcimSitesPartialUpdateParams {
	return &DcimSitesPartialUpdateParams{
		Context: ctx,
	}
}

// NewDcimSitesPartialUpdateParamsWithHTTPClient creates a new DcimSitesPartialUpdateParams object
// with the ability to set a custom HTTPClient for a request.
func NewDcimSitesPartialUpdateParamsWithHTTPClient(client *http.Client) *DcimSitesPartialUpdateParams {
	return &DcimSitesPartialUpdateParams{
		HTTPClient: client,
	}
}

/* DcimSitesPartialUpdateParams contains all the parameters to send to the API endpoint
   for the dcim sites partial update operation.

   Typically these are written to a http.Request.
*/
type DcimSitesPartialUpdateParams struct {

	// Data.
	Data *models.WritableSite

	/* ID.

	   A UUID string identifying this site.

	   Format: uuid
	*/
	ID strfmt.UUID

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the dcim sites partial update params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DcimSitesPartialUpdateParams) WithDefaults() *DcimSitesPartialUpdateParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the dcim sites partial update params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DcimSitesPartialUpdateParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the dcim sites partial update params
func (o *DcimSitesPartialUpdateParams) WithTimeout(timeout time.Duration) *DcimSitesPartialUpdateParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the dcim sites partial update params
func (o *DcimSitesPartialUpdateParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the dcim sites partial update params
func (o *DcimSitesPartialUpdateParams) WithContext(ctx context.Context) *DcimSitesPartialUpdateParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the dcim sites partial update params
func (o *DcimSitesPartialUpdateParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the dcim sites partial update params
func (o *DcimSitesPartialUpdateParams) WithHTTPClient(client *http.Client) *DcimSitesPartialUpdateParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the dcim sites partial update params
func (o *DcimSitesPartialUpdateParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithData adds the data to the dcim sites partial update params
func (o *DcimSitesPartialUpdateParams) WithData(data *models.WritableSite) *DcimSitesPartialUpdateParams {
	o.SetData(data)
	return o
}

// SetData adds the data to the dcim sites partial update params
func (o *DcimSitesPartialUpdateParams) SetData(data *models.WritableSite) {
	o.Data = data
}

// WithID adds the id to the dcim sites partial update params
func (o *DcimSitesPartialUpdateParams) WithID(id strfmt.UUID) *DcimSitesPartialUpdateParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the dcim sites partial update params
func (o *DcimSitesPartialUpdateParams) SetID(id strfmt.UUID) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *DcimSitesPartialUpdateParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
