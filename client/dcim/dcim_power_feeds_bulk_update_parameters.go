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

// NewDcimPowerFeedsBulkUpdateParams creates a new DcimPowerFeedsBulkUpdateParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewDcimPowerFeedsBulkUpdateParams() *DcimPowerFeedsBulkUpdateParams {
	return &DcimPowerFeedsBulkUpdateParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewDcimPowerFeedsBulkUpdateParamsWithTimeout creates a new DcimPowerFeedsBulkUpdateParams object
// with the ability to set a timeout on a request.
func NewDcimPowerFeedsBulkUpdateParamsWithTimeout(timeout time.Duration) *DcimPowerFeedsBulkUpdateParams {
	return &DcimPowerFeedsBulkUpdateParams{
		timeout: timeout,
	}
}

// NewDcimPowerFeedsBulkUpdateParamsWithContext creates a new DcimPowerFeedsBulkUpdateParams object
// with the ability to set a context for a request.
func NewDcimPowerFeedsBulkUpdateParamsWithContext(ctx context.Context) *DcimPowerFeedsBulkUpdateParams {
	return &DcimPowerFeedsBulkUpdateParams{
		Context: ctx,
	}
}

// NewDcimPowerFeedsBulkUpdateParamsWithHTTPClient creates a new DcimPowerFeedsBulkUpdateParams object
// with the ability to set a custom HTTPClient for a request.
func NewDcimPowerFeedsBulkUpdateParamsWithHTTPClient(client *http.Client) *DcimPowerFeedsBulkUpdateParams {
	return &DcimPowerFeedsBulkUpdateParams{
		HTTPClient: client,
	}
}

/* DcimPowerFeedsBulkUpdateParams contains all the parameters to send to the API endpoint
   for the dcim power feeds bulk update operation.

   Typically these are written to a http.Request.
*/
type DcimPowerFeedsBulkUpdateParams struct {

	// Data.
	Data *models.WritablePowerFeed

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the dcim power feeds bulk update params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DcimPowerFeedsBulkUpdateParams) WithDefaults() *DcimPowerFeedsBulkUpdateParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the dcim power feeds bulk update params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DcimPowerFeedsBulkUpdateParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the dcim power feeds bulk update params
func (o *DcimPowerFeedsBulkUpdateParams) WithTimeout(timeout time.Duration) *DcimPowerFeedsBulkUpdateParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the dcim power feeds bulk update params
func (o *DcimPowerFeedsBulkUpdateParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the dcim power feeds bulk update params
func (o *DcimPowerFeedsBulkUpdateParams) WithContext(ctx context.Context) *DcimPowerFeedsBulkUpdateParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the dcim power feeds bulk update params
func (o *DcimPowerFeedsBulkUpdateParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the dcim power feeds bulk update params
func (o *DcimPowerFeedsBulkUpdateParams) WithHTTPClient(client *http.Client) *DcimPowerFeedsBulkUpdateParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the dcim power feeds bulk update params
func (o *DcimPowerFeedsBulkUpdateParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithData adds the data to the dcim power feeds bulk update params
func (o *DcimPowerFeedsBulkUpdateParams) WithData(data *models.WritablePowerFeed) *DcimPowerFeedsBulkUpdateParams {
	o.SetData(data)
	return o
}

// SetData adds the data to the dcim power feeds bulk update params
func (o *DcimPowerFeedsBulkUpdateParams) SetData(data *models.WritablePowerFeed) {
	o.Data = data
}

// WriteToRequest writes these params to a swagger request
func (o *DcimPowerFeedsBulkUpdateParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
