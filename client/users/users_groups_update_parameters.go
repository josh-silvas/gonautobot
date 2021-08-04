package users

import (
	"context"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"

	"github.com/josh5276/gonautobot/gonautobot/models"
)

// NewUsersGroupsUpdateParams creates a new UsersGroupsUpdateParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewUsersGroupsUpdateParams() *UsersGroupsUpdateParams {
	return &UsersGroupsUpdateParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewUsersGroupsUpdateParamsWithTimeout creates a new UsersGroupsUpdateParams object
// with the ability to set a timeout on a request.
func NewUsersGroupsUpdateParamsWithTimeout(timeout time.Duration) *UsersGroupsUpdateParams {
	return &UsersGroupsUpdateParams{
		timeout: timeout,
	}
}

// NewUsersGroupsUpdateParamsWithContext creates a new UsersGroupsUpdateParams object
// with the ability to set a context for a request.
func NewUsersGroupsUpdateParamsWithContext(ctx context.Context) *UsersGroupsUpdateParams {
	return &UsersGroupsUpdateParams{
		Context: ctx,
	}
}

// NewUsersGroupsUpdateParamsWithHTTPClient creates a new UsersGroupsUpdateParams object
// with the ability to set a custom HTTPClient for a request.
func NewUsersGroupsUpdateParamsWithHTTPClient(client *http.Client) *UsersGroupsUpdateParams {
	return &UsersGroupsUpdateParams{
		HTTPClient: client,
	}
}

/* UsersGroupsUpdateParams contains all the parameters to send to the API endpoint
   for the users groups update operation.

   Typically these are written to a http.Request.
*/
type UsersGroupsUpdateParams struct {

	// Data.
	Data *models.Group

	/* ID.

	   A unique integer value identifying this group.
	*/
	ID int64

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the users groups update params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *UsersGroupsUpdateParams) WithDefaults() *UsersGroupsUpdateParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the users groups update params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *UsersGroupsUpdateParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the users groups update params
func (o *UsersGroupsUpdateParams) WithTimeout(timeout time.Duration) *UsersGroupsUpdateParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the users groups update params
func (o *UsersGroupsUpdateParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the users groups update params
func (o *UsersGroupsUpdateParams) WithContext(ctx context.Context) *UsersGroupsUpdateParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the users groups update params
func (o *UsersGroupsUpdateParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the users groups update params
func (o *UsersGroupsUpdateParams) WithHTTPClient(client *http.Client) *UsersGroupsUpdateParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the users groups update params
func (o *UsersGroupsUpdateParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithData adds the data to the users groups update params
func (o *UsersGroupsUpdateParams) WithData(data *models.Group) *UsersGroupsUpdateParams {
	o.SetData(data)
	return o
}

// SetData adds the data to the users groups update params
func (o *UsersGroupsUpdateParams) SetData(data *models.Group) {
	o.Data = data
}

// WithID adds the id to the users groups update params
func (o *UsersGroupsUpdateParams) WithID(id int64) *UsersGroupsUpdateParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the users groups update params
func (o *UsersGroupsUpdateParams) SetID(id int64) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *UsersGroupsUpdateParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
	if err := r.SetPathParam("id", swag.FormatInt64(o.ID)); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
