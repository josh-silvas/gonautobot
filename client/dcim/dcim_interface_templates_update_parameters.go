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

// NewDcimInterfaceTemplatesUpdateParams creates a new DcimInterfaceTemplatesUpdateParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewDcimInterfaceTemplatesUpdateParams() *DcimInterfaceTemplatesUpdateParams {
	return &DcimInterfaceTemplatesUpdateParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewDcimInterfaceTemplatesUpdateParamsWithTimeout creates a new DcimInterfaceTemplatesUpdateParams object
// with the ability to set a timeout on a request.
func NewDcimInterfaceTemplatesUpdateParamsWithTimeout(timeout time.Duration) *DcimInterfaceTemplatesUpdateParams {
	return &DcimInterfaceTemplatesUpdateParams{
		timeout: timeout,
	}
}

// NewDcimInterfaceTemplatesUpdateParamsWithContext creates a new DcimInterfaceTemplatesUpdateParams object
// with the ability to set a context for a request.
func NewDcimInterfaceTemplatesUpdateParamsWithContext(ctx context.Context) *DcimInterfaceTemplatesUpdateParams {
	return &DcimInterfaceTemplatesUpdateParams{
		Context: ctx,
	}
}

// NewDcimInterfaceTemplatesUpdateParamsWithHTTPClient creates a new DcimInterfaceTemplatesUpdateParams object
// with the ability to set a custom HTTPClient for a request.
func NewDcimInterfaceTemplatesUpdateParamsWithHTTPClient(client *http.Client) *DcimInterfaceTemplatesUpdateParams {
	return &DcimInterfaceTemplatesUpdateParams{
		HTTPClient: client,
	}
}

/* DcimInterfaceTemplatesUpdateParams contains all the parameters to send to the API endpoint
   for the dcim interface templates update operation.

   Typically these are written to a http.Request.
*/
type DcimInterfaceTemplatesUpdateParams struct {

	// Data.
	Data *models.WritableInterfaceTemplate

	/* ID.

	   A UUID string identifying this interface template.

	   Format: uuid
	*/
	ID strfmt.UUID

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the dcim interface templates update params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DcimInterfaceTemplatesUpdateParams) WithDefaults() *DcimInterfaceTemplatesUpdateParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the dcim interface templates update params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DcimInterfaceTemplatesUpdateParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the dcim interface templates update params
func (o *DcimInterfaceTemplatesUpdateParams) WithTimeout(timeout time.Duration) *DcimInterfaceTemplatesUpdateParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the dcim interface templates update params
func (o *DcimInterfaceTemplatesUpdateParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the dcim interface templates update params
func (o *DcimInterfaceTemplatesUpdateParams) WithContext(ctx context.Context) *DcimInterfaceTemplatesUpdateParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the dcim interface templates update params
func (o *DcimInterfaceTemplatesUpdateParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the dcim interface templates update params
func (o *DcimInterfaceTemplatesUpdateParams) WithHTTPClient(client *http.Client) *DcimInterfaceTemplatesUpdateParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the dcim interface templates update params
func (o *DcimInterfaceTemplatesUpdateParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithData adds the data to the dcim interface templates update params
func (o *DcimInterfaceTemplatesUpdateParams) WithData(data *models.WritableInterfaceTemplate) *DcimInterfaceTemplatesUpdateParams {
	o.SetData(data)
	return o
}

// SetData adds the data to the dcim interface templates update params
func (o *DcimInterfaceTemplatesUpdateParams) SetData(data *models.WritableInterfaceTemplate) {
	o.Data = data
}

// WithID adds the id to the dcim interface templates update params
func (o *DcimInterfaceTemplatesUpdateParams) WithID(id strfmt.UUID) *DcimInterfaceTemplatesUpdateParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the dcim interface templates update params
func (o *DcimInterfaceTemplatesUpdateParams) SetID(id strfmt.UUID) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *DcimInterfaceTemplatesUpdateParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
