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

// NewDcimInventoryItemsCreateParams creates a new DcimInventoryItemsCreateParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewDcimInventoryItemsCreateParams() *DcimInventoryItemsCreateParams {
	return &DcimInventoryItemsCreateParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewDcimInventoryItemsCreateParamsWithTimeout creates a new DcimInventoryItemsCreateParams object
// with the ability to set a timeout on a request.
func NewDcimInventoryItemsCreateParamsWithTimeout(timeout time.Duration) *DcimInventoryItemsCreateParams {
	return &DcimInventoryItemsCreateParams{
		timeout: timeout,
	}
}

// NewDcimInventoryItemsCreateParamsWithContext creates a new DcimInventoryItemsCreateParams object
// with the ability to set a context for a request.
func NewDcimInventoryItemsCreateParamsWithContext(ctx context.Context) *DcimInventoryItemsCreateParams {
	return &DcimInventoryItemsCreateParams{
		Context: ctx,
	}
}

// NewDcimInventoryItemsCreateParamsWithHTTPClient creates a new DcimInventoryItemsCreateParams object
// with the ability to set a custom HTTPClient for a request.
func NewDcimInventoryItemsCreateParamsWithHTTPClient(client *http.Client) *DcimInventoryItemsCreateParams {
	return &DcimInventoryItemsCreateParams{
		HTTPClient: client,
	}
}

/* DcimInventoryItemsCreateParams contains all the parameters to send to the API endpoint
   for the dcim inventory items create operation.

   Typically these are written to a http.Request.
*/
type DcimInventoryItemsCreateParams struct {

	// Data.
	Data *models.WritableInventoryItem

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the dcim inventory items create params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DcimInventoryItemsCreateParams) WithDefaults() *DcimInventoryItemsCreateParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the dcim inventory items create params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DcimInventoryItemsCreateParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the dcim inventory items create params
func (o *DcimInventoryItemsCreateParams) WithTimeout(timeout time.Duration) *DcimInventoryItemsCreateParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the dcim inventory items create params
func (o *DcimInventoryItemsCreateParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the dcim inventory items create params
func (o *DcimInventoryItemsCreateParams) WithContext(ctx context.Context) *DcimInventoryItemsCreateParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the dcim inventory items create params
func (o *DcimInventoryItemsCreateParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the dcim inventory items create params
func (o *DcimInventoryItemsCreateParams) WithHTTPClient(client *http.Client) *DcimInventoryItemsCreateParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the dcim inventory items create params
func (o *DcimInventoryItemsCreateParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithData adds the data to the dcim inventory items create params
func (o *DcimInventoryItemsCreateParams) WithData(data *models.WritableInventoryItem) *DcimInventoryItemsCreateParams {
	o.SetData(data)
	return o
}

// SetData adds the data to the dcim inventory items create params
func (o *DcimInventoryItemsCreateParams) SetData(data *models.WritableInventoryItem) {
	o.Data = data
}

// WriteToRequest writes these params to a swagger request
func (o *DcimInventoryItemsCreateParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
