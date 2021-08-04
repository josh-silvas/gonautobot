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

// NewDcimInterfaceTemplatesBulkUpdateParams creates a new DcimInterfaceTemplatesBulkUpdateParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewDcimInterfaceTemplatesBulkUpdateParams() *DcimInterfaceTemplatesBulkUpdateParams {
	return &DcimInterfaceTemplatesBulkUpdateParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewDcimInterfaceTemplatesBulkUpdateParamsWithTimeout creates a new DcimInterfaceTemplatesBulkUpdateParams object
// with the ability to set a timeout on a request.
func NewDcimInterfaceTemplatesBulkUpdateParamsWithTimeout(timeout time.Duration) *DcimInterfaceTemplatesBulkUpdateParams {
	return &DcimInterfaceTemplatesBulkUpdateParams{
		timeout: timeout,
	}
}

// NewDcimInterfaceTemplatesBulkUpdateParamsWithContext creates a new DcimInterfaceTemplatesBulkUpdateParams object
// with the ability to set a context for a request.
func NewDcimInterfaceTemplatesBulkUpdateParamsWithContext(ctx context.Context) *DcimInterfaceTemplatesBulkUpdateParams {
	return &DcimInterfaceTemplatesBulkUpdateParams{
		Context: ctx,
	}
}

// NewDcimInterfaceTemplatesBulkUpdateParamsWithHTTPClient creates a new DcimInterfaceTemplatesBulkUpdateParams object
// with the ability to set a custom HTTPClient for a request.
func NewDcimInterfaceTemplatesBulkUpdateParamsWithHTTPClient(client *http.Client) *DcimInterfaceTemplatesBulkUpdateParams {
	return &DcimInterfaceTemplatesBulkUpdateParams{
		HTTPClient: client,
	}
}

/* DcimInterfaceTemplatesBulkUpdateParams contains all the parameters to send to the API endpoint
   for the dcim interface templates bulk update operation.

   Typically these are written to a http.Request.
*/
type DcimInterfaceTemplatesBulkUpdateParams struct {

	// Data.
	Data *models.WritableInterfaceTemplate

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the dcim interface templates bulk update params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DcimInterfaceTemplatesBulkUpdateParams) WithDefaults() *DcimInterfaceTemplatesBulkUpdateParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the dcim interface templates bulk update params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DcimInterfaceTemplatesBulkUpdateParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the dcim interface templates bulk update params
func (o *DcimInterfaceTemplatesBulkUpdateParams) WithTimeout(timeout time.Duration) *DcimInterfaceTemplatesBulkUpdateParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the dcim interface templates bulk update params
func (o *DcimInterfaceTemplatesBulkUpdateParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the dcim interface templates bulk update params
func (o *DcimInterfaceTemplatesBulkUpdateParams) WithContext(ctx context.Context) *DcimInterfaceTemplatesBulkUpdateParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the dcim interface templates bulk update params
func (o *DcimInterfaceTemplatesBulkUpdateParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the dcim interface templates bulk update params
func (o *DcimInterfaceTemplatesBulkUpdateParams) WithHTTPClient(client *http.Client) *DcimInterfaceTemplatesBulkUpdateParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the dcim interface templates bulk update params
func (o *DcimInterfaceTemplatesBulkUpdateParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithData adds the data to the dcim interface templates bulk update params
func (o *DcimInterfaceTemplatesBulkUpdateParams) WithData(data *models.WritableInterfaceTemplate) *DcimInterfaceTemplatesBulkUpdateParams {
	o.SetData(data)
	return o
}

// SetData adds the data to the dcim interface templates bulk update params
func (o *DcimInterfaceTemplatesBulkUpdateParams) SetData(data *models.WritableInterfaceTemplate) {
	o.Data = data
}

// WriteToRequest writes these params to a swagger request
func (o *DcimInterfaceTemplatesBulkUpdateParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
