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

// NewUsersUsersBulkUpdateParams creates a new UsersUsersBulkUpdateParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewUsersUsersBulkUpdateParams() *UsersUsersBulkUpdateParams {
	return &UsersUsersBulkUpdateParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewUsersUsersBulkUpdateParamsWithTimeout creates a new UsersUsersBulkUpdateParams object
// with the ability to set a timeout on a request.
func NewUsersUsersBulkUpdateParamsWithTimeout(timeout time.Duration) *UsersUsersBulkUpdateParams {
	return &UsersUsersBulkUpdateParams{
		timeout: timeout,
	}
}

// NewUsersUsersBulkUpdateParamsWithContext creates a new UsersUsersBulkUpdateParams object
// with the ability to set a context for a request.
func NewUsersUsersBulkUpdateParamsWithContext(ctx context.Context) *UsersUsersBulkUpdateParams {
	return &UsersUsersBulkUpdateParams{
		Context: ctx,
	}
}

// NewUsersUsersBulkUpdateParamsWithHTTPClient creates a new UsersUsersBulkUpdateParams object
// with the ability to set a custom HTTPClient for a request.
func NewUsersUsersBulkUpdateParamsWithHTTPClient(client *http.Client) *UsersUsersBulkUpdateParams {
	return &UsersUsersBulkUpdateParams{
		HTTPClient: client,
	}
}

/* UsersUsersBulkUpdateParams contains all the parameters to send to the API endpoint
   for the users users bulk update operation.

   Typically these are written to a http.Request.
*/
type UsersUsersBulkUpdateParams struct {

	// Data.
	Data *models.WritableUser

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the users users bulk update params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *UsersUsersBulkUpdateParams) WithDefaults() *UsersUsersBulkUpdateParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the users users bulk update params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *UsersUsersBulkUpdateParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the users users bulk update params
func (o *UsersUsersBulkUpdateParams) WithTimeout(timeout time.Duration) *UsersUsersBulkUpdateParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the users users bulk update params
func (o *UsersUsersBulkUpdateParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the users users bulk update params
func (o *UsersUsersBulkUpdateParams) WithContext(ctx context.Context) *UsersUsersBulkUpdateParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the users users bulk update params
func (o *UsersUsersBulkUpdateParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the users users bulk update params
func (o *UsersUsersBulkUpdateParams) WithHTTPClient(client *http.Client) *UsersUsersBulkUpdateParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the users users bulk update params
func (o *UsersUsersBulkUpdateParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithData adds the data to the users users bulk update params
func (o *UsersUsersBulkUpdateParams) WithData(data *models.WritableUser) *UsersUsersBulkUpdateParams {
	o.SetData(data)
	return o
}

// SetData adds the data to the users users bulk update params
func (o *UsersUsersBulkUpdateParams) SetData(data *models.WritableUser) {
	o.Data = data
}

// WriteToRequest writes these params to a swagger request
func (o *UsersUsersBulkUpdateParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
