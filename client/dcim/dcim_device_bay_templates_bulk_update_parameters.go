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

// NewDcimDeviceBayTemplatesBulkUpdateParams creates a new DcimDeviceBayTemplatesBulkUpdateParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewDcimDeviceBayTemplatesBulkUpdateParams() *DcimDeviceBayTemplatesBulkUpdateParams {
	return &DcimDeviceBayTemplatesBulkUpdateParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewDcimDeviceBayTemplatesBulkUpdateParamsWithTimeout creates a new DcimDeviceBayTemplatesBulkUpdateParams object
// with the ability to set a timeout on a request.
func NewDcimDeviceBayTemplatesBulkUpdateParamsWithTimeout(timeout time.Duration) *DcimDeviceBayTemplatesBulkUpdateParams {
	return &DcimDeviceBayTemplatesBulkUpdateParams{
		timeout: timeout,
	}
}

// NewDcimDeviceBayTemplatesBulkUpdateParamsWithContext creates a new DcimDeviceBayTemplatesBulkUpdateParams object
// with the ability to set a context for a request.
func NewDcimDeviceBayTemplatesBulkUpdateParamsWithContext(ctx context.Context) *DcimDeviceBayTemplatesBulkUpdateParams {
	return &DcimDeviceBayTemplatesBulkUpdateParams{
		Context: ctx,
	}
}

// NewDcimDeviceBayTemplatesBulkUpdateParamsWithHTTPClient creates a new DcimDeviceBayTemplatesBulkUpdateParams object
// with the ability to set a custom HTTPClient for a request.
func NewDcimDeviceBayTemplatesBulkUpdateParamsWithHTTPClient(client *http.Client) *DcimDeviceBayTemplatesBulkUpdateParams {
	return &DcimDeviceBayTemplatesBulkUpdateParams{
		HTTPClient: client,
	}
}

/* DcimDeviceBayTemplatesBulkUpdateParams contains all the parameters to send to the API endpoint
   for the dcim device bay templates bulk update operation.

   Typically these are written to a http.Request.
*/
type DcimDeviceBayTemplatesBulkUpdateParams struct {

	// Data.
	Data *models.WritableDeviceBayTemplate

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the dcim device bay templates bulk update params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DcimDeviceBayTemplatesBulkUpdateParams) WithDefaults() *DcimDeviceBayTemplatesBulkUpdateParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the dcim device bay templates bulk update params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DcimDeviceBayTemplatesBulkUpdateParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the dcim device bay templates bulk update params
func (o *DcimDeviceBayTemplatesBulkUpdateParams) WithTimeout(timeout time.Duration) *DcimDeviceBayTemplatesBulkUpdateParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the dcim device bay templates bulk update params
func (o *DcimDeviceBayTemplatesBulkUpdateParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the dcim device bay templates bulk update params
func (o *DcimDeviceBayTemplatesBulkUpdateParams) WithContext(ctx context.Context) *DcimDeviceBayTemplatesBulkUpdateParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the dcim device bay templates bulk update params
func (o *DcimDeviceBayTemplatesBulkUpdateParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the dcim device bay templates bulk update params
func (o *DcimDeviceBayTemplatesBulkUpdateParams) WithHTTPClient(client *http.Client) *DcimDeviceBayTemplatesBulkUpdateParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the dcim device bay templates bulk update params
func (o *DcimDeviceBayTemplatesBulkUpdateParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithData adds the data to the dcim device bay templates bulk update params
func (o *DcimDeviceBayTemplatesBulkUpdateParams) WithData(data *models.WritableDeviceBayTemplate) *DcimDeviceBayTemplatesBulkUpdateParams {
	o.SetData(data)
	return o
}

// SetData adds the data to the dcim device bay templates bulk update params
func (o *DcimDeviceBayTemplatesBulkUpdateParams) SetData(data *models.WritableDeviceBayTemplate) {
	o.Data = data
}

// WriteToRequest writes these params to a swagger request
func (o *DcimDeviceBayTemplatesBulkUpdateParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
