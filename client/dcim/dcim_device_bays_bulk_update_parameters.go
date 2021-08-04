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

// NewDcimDeviceBaysBulkUpdateParams creates a new DcimDeviceBaysBulkUpdateParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewDcimDeviceBaysBulkUpdateParams() *DcimDeviceBaysBulkUpdateParams {
	return &DcimDeviceBaysBulkUpdateParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewDcimDeviceBaysBulkUpdateParamsWithTimeout creates a new DcimDeviceBaysBulkUpdateParams object
// with the ability to set a timeout on a request.
func NewDcimDeviceBaysBulkUpdateParamsWithTimeout(timeout time.Duration) *DcimDeviceBaysBulkUpdateParams {
	return &DcimDeviceBaysBulkUpdateParams{
		timeout: timeout,
	}
}

// NewDcimDeviceBaysBulkUpdateParamsWithContext creates a new DcimDeviceBaysBulkUpdateParams object
// with the ability to set a context for a request.
func NewDcimDeviceBaysBulkUpdateParamsWithContext(ctx context.Context) *DcimDeviceBaysBulkUpdateParams {
	return &DcimDeviceBaysBulkUpdateParams{
		Context: ctx,
	}
}

// NewDcimDeviceBaysBulkUpdateParamsWithHTTPClient creates a new DcimDeviceBaysBulkUpdateParams object
// with the ability to set a custom HTTPClient for a request.
func NewDcimDeviceBaysBulkUpdateParamsWithHTTPClient(client *http.Client) *DcimDeviceBaysBulkUpdateParams {
	return &DcimDeviceBaysBulkUpdateParams{
		HTTPClient: client,
	}
}

/* DcimDeviceBaysBulkUpdateParams contains all the parameters to send to the API endpoint
   for the dcim device bays bulk update operation.

   Typically these are written to a http.Request.
*/
type DcimDeviceBaysBulkUpdateParams struct {

	// Data.
	Data *models.WritableDeviceBay

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the dcim device bays bulk update params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DcimDeviceBaysBulkUpdateParams) WithDefaults() *DcimDeviceBaysBulkUpdateParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the dcim device bays bulk update params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DcimDeviceBaysBulkUpdateParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the dcim device bays bulk update params
func (o *DcimDeviceBaysBulkUpdateParams) WithTimeout(timeout time.Duration) *DcimDeviceBaysBulkUpdateParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the dcim device bays bulk update params
func (o *DcimDeviceBaysBulkUpdateParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the dcim device bays bulk update params
func (o *DcimDeviceBaysBulkUpdateParams) WithContext(ctx context.Context) *DcimDeviceBaysBulkUpdateParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the dcim device bays bulk update params
func (o *DcimDeviceBaysBulkUpdateParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the dcim device bays bulk update params
func (o *DcimDeviceBaysBulkUpdateParams) WithHTTPClient(client *http.Client) *DcimDeviceBaysBulkUpdateParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the dcim device bays bulk update params
func (o *DcimDeviceBaysBulkUpdateParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithData adds the data to the dcim device bays bulk update params
func (o *DcimDeviceBaysBulkUpdateParams) WithData(data *models.WritableDeviceBay) *DcimDeviceBaysBulkUpdateParams {
	o.SetData(data)
	return o
}

// SetData adds the data to the dcim device bays bulk update params
func (o *DcimDeviceBaysBulkUpdateParams) SetData(data *models.WritableDeviceBay) {
	o.Data = data
}

// WriteToRequest writes these params to a swagger request
func (o *DcimDeviceBaysBulkUpdateParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
