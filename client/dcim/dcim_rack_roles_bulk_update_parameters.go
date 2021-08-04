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

// NewDcimRackRolesBulkUpdateParams creates a new DcimRackRolesBulkUpdateParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewDcimRackRolesBulkUpdateParams() *DcimRackRolesBulkUpdateParams {
	return &DcimRackRolesBulkUpdateParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewDcimRackRolesBulkUpdateParamsWithTimeout creates a new DcimRackRolesBulkUpdateParams object
// with the ability to set a timeout on a request.
func NewDcimRackRolesBulkUpdateParamsWithTimeout(timeout time.Duration) *DcimRackRolesBulkUpdateParams {
	return &DcimRackRolesBulkUpdateParams{
		timeout: timeout,
	}
}

// NewDcimRackRolesBulkUpdateParamsWithContext creates a new DcimRackRolesBulkUpdateParams object
// with the ability to set a context for a request.
func NewDcimRackRolesBulkUpdateParamsWithContext(ctx context.Context) *DcimRackRolesBulkUpdateParams {
	return &DcimRackRolesBulkUpdateParams{
		Context: ctx,
	}
}

// NewDcimRackRolesBulkUpdateParamsWithHTTPClient creates a new DcimRackRolesBulkUpdateParams object
// with the ability to set a custom HTTPClient for a request.
func NewDcimRackRolesBulkUpdateParamsWithHTTPClient(client *http.Client) *DcimRackRolesBulkUpdateParams {
	return &DcimRackRolesBulkUpdateParams{
		HTTPClient: client,
	}
}

/* DcimRackRolesBulkUpdateParams contains all the parameters to send to the API endpoint
   for the dcim rack roles bulk update operation.

   Typically these are written to a http.Request.
*/
type DcimRackRolesBulkUpdateParams struct {

	// Data.
	Data *models.RackRole

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the dcim rack roles bulk update params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DcimRackRolesBulkUpdateParams) WithDefaults() *DcimRackRolesBulkUpdateParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the dcim rack roles bulk update params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DcimRackRolesBulkUpdateParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the dcim rack roles bulk update params
func (o *DcimRackRolesBulkUpdateParams) WithTimeout(timeout time.Duration) *DcimRackRolesBulkUpdateParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the dcim rack roles bulk update params
func (o *DcimRackRolesBulkUpdateParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the dcim rack roles bulk update params
func (o *DcimRackRolesBulkUpdateParams) WithContext(ctx context.Context) *DcimRackRolesBulkUpdateParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the dcim rack roles bulk update params
func (o *DcimRackRolesBulkUpdateParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the dcim rack roles bulk update params
func (o *DcimRackRolesBulkUpdateParams) WithHTTPClient(client *http.Client) *DcimRackRolesBulkUpdateParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the dcim rack roles bulk update params
func (o *DcimRackRolesBulkUpdateParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithData adds the data to the dcim rack roles bulk update params
func (o *DcimRackRolesBulkUpdateParams) WithData(data *models.RackRole) *DcimRackRolesBulkUpdateParams {
	o.SetData(data)
	return o
}

// SetData adds the data to the dcim rack roles bulk update params
func (o *DcimRackRolesBulkUpdateParams) SetData(data *models.RackRole) {
	o.Data = data
}

// WriteToRequest writes these params to a swagger request
func (o *DcimRackRolesBulkUpdateParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
