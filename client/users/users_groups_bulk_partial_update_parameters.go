package users

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

// NewUsersGroupsBulkPartialUpdateParams creates a new UsersGroupsBulkPartialUpdateParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewUsersGroupsBulkPartialUpdateParams() *UsersGroupsBulkPartialUpdateParams {
	return &UsersGroupsBulkPartialUpdateParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewUsersGroupsBulkPartialUpdateParamsWithTimeout creates a new UsersGroupsBulkPartialUpdateParams object
// with the ability to set a timeout on a request.
func NewUsersGroupsBulkPartialUpdateParamsWithTimeout(timeout time.Duration) *UsersGroupsBulkPartialUpdateParams {
	return &UsersGroupsBulkPartialUpdateParams{
		timeout: timeout,
	}
}

// NewUsersGroupsBulkPartialUpdateParamsWithContext creates a new UsersGroupsBulkPartialUpdateParams object
// with the ability to set a context for a request.
func NewUsersGroupsBulkPartialUpdateParamsWithContext(ctx context.Context) *UsersGroupsBulkPartialUpdateParams {
	return &UsersGroupsBulkPartialUpdateParams{
		Context: ctx,
	}
}

// NewUsersGroupsBulkPartialUpdateParamsWithHTTPClient creates a new UsersGroupsBulkPartialUpdateParams object
// with the ability to set a custom HTTPClient for a request.
func NewUsersGroupsBulkPartialUpdateParamsWithHTTPClient(client *http.Client) *UsersGroupsBulkPartialUpdateParams {
	return &UsersGroupsBulkPartialUpdateParams{
		HTTPClient: client,
	}
}

/* UsersGroupsBulkPartialUpdateParams contains all the parameters to send to the API endpoint
   for the users groups bulk partial update operation.

   Typically these are written to a http.Request.
*/
type UsersGroupsBulkPartialUpdateParams struct {

	// Data.
	Data *models.Group

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the users groups bulk partial update params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *UsersGroupsBulkPartialUpdateParams) WithDefaults() *UsersGroupsBulkPartialUpdateParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the users groups bulk partial update params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *UsersGroupsBulkPartialUpdateParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the users groups bulk partial update params
func (o *UsersGroupsBulkPartialUpdateParams) WithTimeout(timeout time.Duration) *UsersGroupsBulkPartialUpdateParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the users groups bulk partial update params
func (o *UsersGroupsBulkPartialUpdateParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the users groups bulk partial update params
func (o *UsersGroupsBulkPartialUpdateParams) WithContext(ctx context.Context) *UsersGroupsBulkPartialUpdateParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the users groups bulk partial update params
func (o *UsersGroupsBulkPartialUpdateParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the users groups bulk partial update params
func (o *UsersGroupsBulkPartialUpdateParams) WithHTTPClient(client *http.Client) *UsersGroupsBulkPartialUpdateParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the users groups bulk partial update params
func (o *UsersGroupsBulkPartialUpdateParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithData adds the data to the users groups bulk partial update params
func (o *UsersGroupsBulkPartialUpdateParams) WithData(data *models.Group) *UsersGroupsBulkPartialUpdateParams {
	o.SetData(data)
	return o
}

// SetData adds the data to the users groups bulk partial update params
func (o *UsersGroupsBulkPartialUpdateParams) SetData(data *models.Group) {
	o.Data = data
}

// WriteToRequest writes these params to a swagger request
func (o *UsersGroupsBulkPartialUpdateParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
