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

// NewDcimConsoleServerPortTemplatesBulkPartialUpdateParams creates a new DcimConsoleServerPortTemplatesBulkPartialUpdateParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewDcimConsoleServerPortTemplatesBulkPartialUpdateParams() *DcimConsoleServerPortTemplatesBulkPartialUpdateParams {
	return &DcimConsoleServerPortTemplatesBulkPartialUpdateParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewDcimConsoleServerPortTemplatesBulkPartialUpdateParamsWithTimeout creates a new DcimConsoleServerPortTemplatesBulkPartialUpdateParams object
// with the ability to set a timeout on a request.
func NewDcimConsoleServerPortTemplatesBulkPartialUpdateParamsWithTimeout(timeout time.Duration) *DcimConsoleServerPortTemplatesBulkPartialUpdateParams {
	return &DcimConsoleServerPortTemplatesBulkPartialUpdateParams{
		timeout: timeout,
	}
}

// NewDcimConsoleServerPortTemplatesBulkPartialUpdateParamsWithContext creates a new DcimConsoleServerPortTemplatesBulkPartialUpdateParams object
// with the ability to set a context for a request.
func NewDcimConsoleServerPortTemplatesBulkPartialUpdateParamsWithContext(ctx context.Context) *DcimConsoleServerPortTemplatesBulkPartialUpdateParams {
	return &DcimConsoleServerPortTemplatesBulkPartialUpdateParams{
		Context: ctx,
	}
}

// NewDcimConsoleServerPortTemplatesBulkPartialUpdateParamsWithHTTPClient creates a new DcimConsoleServerPortTemplatesBulkPartialUpdateParams object
// with the ability to set a custom HTTPClient for a request.
func NewDcimConsoleServerPortTemplatesBulkPartialUpdateParamsWithHTTPClient(client *http.Client) *DcimConsoleServerPortTemplatesBulkPartialUpdateParams {
	return &DcimConsoleServerPortTemplatesBulkPartialUpdateParams{
		HTTPClient: client,
	}
}

/* DcimConsoleServerPortTemplatesBulkPartialUpdateParams contains all the parameters to send to the API endpoint
   for the dcim console server port templates bulk partial update operation.

   Typically these are written to a http.Request.
*/
type DcimConsoleServerPortTemplatesBulkPartialUpdateParams struct {

	// Data.
	Data *models.WritableConsoleServerPortTemplate

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the dcim console server port templates bulk partial update params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DcimConsoleServerPortTemplatesBulkPartialUpdateParams) WithDefaults() *DcimConsoleServerPortTemplatesBulkPartialUpdateParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the dcim console server port templates bulk partial update params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DcimConsoleServerPortTemplatesBulkPartialUpdateParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the dcim console server port templates bulk partial update params
func (o *DcimConsoleServerPortTemplatesBulkPartialUpdateParams) WithTimeout(timeout time.Duration) *DcimConsoleServerPortTemplatesBulkPartialUpdateParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the dcim console server port templates bulk partial update params
func (o *DcimConsoleServerPortTemplatesBulkPartialUpdateParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the dcim console server port templates bulk partial update params
func (o *DcimConsoleServerPortTemplatesBulkPartialUpdateParams) WithContext(ctx context.Context) *DcimConsoleServerPortTemplatesBulkPartialUpdateParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the dcim console server port templates bulk partial update params
func (o *DcimConsoleServerPortTemplatesBulkPartialUpdateParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the dcim console server port templates bulk partial update params
func (o *DcimConsoleServerPortTemplatesBulkPartialUpdateParams) WithHTTPClient(client *http.Client) *DcimConsoleServerPortTemplatesBulkPartialUpdateParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the dcim console server port templates bulk partial update params
func (o *DcimConsoleServerPortTemplatesBulkPartialUpdateParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithData adds the data to the dcim console server port templates bulk partial update params
func (o *DcimConsoleServerPortTemplatesBulkPartialUpdateParams) WithData(data *models.WritableConsoleServerPortTemplate) *DcimConsoleServerPortTemplatesBulkPartialUpdateParams {
	o.SetData(data)
	return o
}

// SetData adds the data to the dcim console server port templates bulk partial update params
func (o *DcimConsoleServerPortTemplatesBulkPartialUpdateParams) SetData(data *models.WritableConsoleServerPortTemplate) {
	o.Data = data
}

// WriteToRequest writes these params to a swagger request
func (o *DcimConsoleServerPortTemplatesBulkPartialUpdateParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
