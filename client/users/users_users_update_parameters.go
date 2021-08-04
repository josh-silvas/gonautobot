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

// NewUsersUsersUpdateParams creates a new UsersUsersUpdateParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewUsersUsersUpdateParams() *UsersUsersUpdateParams {
	return &UsersUsersUpdateParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewUsersUsersUpdateParamsWithTimeout creates a new UsersUsersUpdateParams object
// with the ability to set a timeout on a request.
func NewUsersUsersUpdateParamsWithTimeout(timeout time.Duration) *UsersUsersUpdateParams {
	return &UsersUsersUpdateParams{
		timeout: timeout,
	}
}

// NewUsersUsersUpdateParamsWithContext creates a new UsersUsersUpdateParams object
// with the ability to set a context for a request.
func NewUsersUsersUpdateParamsWithContext(ctx context.Context) *UsersUsersUpdateParams {
	return &UsersUsersUpdateParams{
		Context: ctx,
	}
}

// NewUsersUsersUpdateParamsWithHTTPClient creates a new UsersUsersUpdateParams object
// with the ability to set a custom HTTPClient for a request.
func NewUsersUsersUpdateParamsWithHTTPClient(client *http.Client) *UsersUsersUpdateParams {
	return &UsersUsersUpdateParams{
		HTTPClient: client,
	}
}

/* UsersUsersUpdateParams contains all the parameters to send to the API endpoint
   for the users users update operation.

   Typically these are written to a http.Request.
*/
type UsersUsersUpdateParams struct {

	// Data.
	Data *models.WritableUser

	/* ID.

	   A UUID string identifying this user.

	   Format: uuid
	*/
	ID strfmt.UUID

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the users users update params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *UsersUsersUpdateParams) WithDefaults() *UsersUsersUpdateParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the users users update params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *UsersUsersUpdateParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the users users update params
func (o *UsersUsersUpdateParams) WithTimeout(timeout time.Duration) *UsersUsersUpdateParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the users users update params
func (o *UsersUsersUpdateParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the users users update params
func (o *UsersUsersUpdateParams) WithContext(ctx context.Context) *UsersUsersUpdateParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the users users update params
func (o *UsersUsersUpdateParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the users users update params
func (o *UsersUsersUpdateParams) WithHTTPClient(client *http.Client) *UsersUsersUpdateParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the users users update params
func (o *UsersUsersUpdateParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithData adds the data to the users users update params
func (o *UsersUsersUpdateParams) WithData(data *models.WritableUser) *UsersUsersUpdateParams {
	o.SetData(data)
	return o
}

// SetData adds the data to the users users update params
func (o *UsersUsersUpdateParams) SetData(data *models.WritableUser) {
	o.Data = data
}

// WithID adds the id to the users users update params
func (o *UsersUsersUpdateParams) WithID(id strfmt.UUID) *UsersUsersUpdateParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the users users update params
func (o *UsersUsersUpdateParams) SetID(id strfmt.UUID) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *UsersUsersUpdateParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error
	if o.Data != nil {
		if err := r.SetBodyParam(o.Data); err != nil {
			return err
		}
	}

	// path param id
	if err := r.SetPathParam("id", o.ID.String()); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
