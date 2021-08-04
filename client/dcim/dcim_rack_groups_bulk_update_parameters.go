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

// NewDcimRackGroupsBulkUpdateParams creates a new DcimRackGroupsBulkUpdateParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewDcimRackGroupsBulkUpdateParams() *DcimRackGroupsBulkUpdateParams {
	return &DcimRackGroupsBulkUpdateParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewDcimRackGroupsBulkUpdateParamsWithTimeout creates a new DcimRackGroupsBulkUpdateParams object
// with the ability to set a timeout on a request.
func NewDcimRackGroupsBulkUpdateParamsWithTimeout(timeout time.Duration) *DcimRackGroupsBulkUpdateParams {
	return &DcimRackGroupsBulkUpdateParams{
		timeout: timeout,
	}
}

// NewDcimRackGroupsBulkUpdateParamsWithContext creates a new DcimRackGroupsBulkUpdateParams object
// with the ability to set a context for a request.
func NewDcimRackGroupsBulkUpdateParamsWithContext(ctx context.Context) *DcimRackGroupsBulkUpdateParams {
	return &DcimRackGroupsBulkUpdateParams{
		Context: ctx,
	}
}

// NewDcimRackGroupsBulkUpdateParamsWithHTTPClient creates a new DcimRackGroupsBulkUpdateParams object
// with the ability to set a custom HTTPClient for a request.
func NewDcimRackGroupsBulkUpdateParamsWithHTTPClient(client *http.Client) *DcimRackGroupsBulkUpdateParams {
	return &DcimRackGroupsBulkUpdateParams{
		HTTPClient: client,
	}
}

/* DcimRackGroupsBulkUpdateParams contains all the parameters to send to the API endpoint
   for the dcim rack groups bulk update operation.

   Typically these are written to a http.Request.
*/
type DcimRackGroupsBulkUpdateParams struct {

	// Data.
	Data *models.WritableRackGroup

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the dcim rack groups bulk update params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DcimRackGroupsBulkUpdateParams) WithDefaults() *DcimRackGroupsBulkUpdateParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the dcim rack groups bulk update params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DcimRackGroupsBulkUpdateParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the dcim rack groups bulk update params
func (o *DcimRackGroupsBulkUpdateParams) WithTimeout(timeout time.Duration) *DcimRackGroupsBulkUpdateParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the dcim rack groups bulk update params
func (o *DcimRackGroupsBulkUpdateParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the dcim rack groups bulk update params
func (o *DcimRackGroupsBulkUpdateParams) WithContext(ctx context.Context) *DcimRackGroupsBulkUpdateParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the dcim rack groups bulk update params
func (o *DcimRackGroupsBulkUpdateParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the dcim rack groups bulk update params
func (o *DcimRackGroupsBulkUpdateParams) WithHTTPClient(client *http.Client) *DcimRackGroupsBulkUpdateParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the dcim rack groups bulk update params
func (o *DcimRackGroupsBulkUpdateParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithData adds the data to the dcim rack groups bulk update params
func (o *DcimRackGroupsBulkUpdateParams) WithData(data *models.WritableRackGroup) *DcimRackGroupsBulkUpdateParams {
	o.SetData(data)
	return o
}

// SetData adds the data to the dcim rack groups bulk update params
func (o *DcimRackGroupsBulkUpdateParams) SetData(data *models.WritableRackGroup) {
	o.Data = data
}

// WriteToRequest writes these params to a swagger request
func (o *DcimRackGroupsBulkUpdateParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
