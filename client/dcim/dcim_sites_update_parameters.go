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

// NewDcimSitesUpdateParams creates a new DcimSitesUpdateParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewDcimSitesUpdateParams() *DcimSitesUpdateParams {
	return &DcimSitesUpdateParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewDcimSitesUpdateParamsWithTimeout creates a new DcimSitesUpdateParams object
// with the ability to set a timeout on a request.
func NewDcimSitesUpdateParamsWithTimeout(timeout time.Duration) *DcimSitesUpdateParams {
	return &DcimSitesUpdateParams{
		timeout: timeout,
	}
}

// NewDcimSitesUpdateParamsWithContext creates a new DcimSitesUpdateParams object
// with the ability to set a context for a request.
func NewDcimSitesUpdateParamsWithContext(ctx context.Context) *DcimSitesUpdateParams {
	return &DcimSitesUpdateParams{
		Context: ctx,
	}
}

// NewDcimSitesUpdateParamsWithHTTPClient creates a new DcimSitesUpdateParams object
// with the ability to set a custom HTTPClient for a request.
func NewDcimSitesUpdateParamsWithHTTPClient(client *http.Client) *DcimSitesUpdateParams {
	return &DcimSitesUpdateParams{
		HTTPClient: client,
	}
}

/* DcimSitesUpdateParams contains all the parameters to send to the API endpoint
   for the dcim sites update operation.

   Typically these are written to a http.Request.
*/
type DcimSitesUpdateParams struct {

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

// WithDefaults hydrates default values in the dcim sites update params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DcimSitesUpdateParams) WithDefaults() *DcimSitesUpdateParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the dcim sites update params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DcimSitesUpdateParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the dcim sites update params
func (o *DcimSitesUpdateParams) WithTimeout(timeout time.Duration) *DcimSitesUpdateParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the dcim sites update params
func (o *DcimSitesUpdateParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the dcim sites update params
func (o *DcimSitesUpdateParams) WithContext(ctx context.Context) *DcimSitesUpdateParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the dcim sites update params
func (o *DcimSitesUpdateParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the dcim sites update params
func (o *DcimSitesUpdateParams) WithHTTPClient(client *http.Client) *DcimSitesUpdateParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the dcim sites update params
func (o *DcimSitesUpdateParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithData adds the data to the dcim sites update params
func (o *DcimSitesUpdateParams) WithData(data *models.WritableSite) *DcimSitesUpdateParams {
	o.SetData(data)
	return o
}

// SetData adds the data to the dcim sites update params
func (o *DcimSitesUpdateParams) SetData(data *models.WritableSite) {
	o.Data = data
}

// WithID adds the id to the dcim sites update params
func (o *DcimSitesUpdateParams) WithID(id strfmt.UUID) *DcimSitesUpdateParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the dcim sites update params
func (o *DcimSitesUpdateParams) SetID(id strfmt.UUID) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *DcimSitesUpdateParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
