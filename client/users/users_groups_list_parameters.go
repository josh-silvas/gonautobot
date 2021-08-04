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
)

// NewUsersGroupsListParams creates a new UsersGroupsListParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewUsersGroupsListParams() *UsersGroupsListParams {
	return &UsersGroupsListParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewUsersGroupsListParamsWithTimeout creates a new UsersGroupsListParams object
// with the ability to set a timeout on a request.
func NewUsersGroupsListParamsWithTimeout(timeout time.Duration) *UsersGroupsListParams {
	return &UsersGroupsListParams{
		timeout: timeout,
	}
}

// NewUsersGroupsListParamsWithContext creates a new UsersGroupsListParams object
// with the ability to set a context for a request.
func NewUsersGroupsListParamsWithContext(ctx context.Context) *UsersGroupsListParams {
	return &UsersGroupsListParams{
		Context: ctx,
	}
}

// NewUsersGroupsListParamsWithHTTPClient creates a new UsersGroupsListParams object
// with the ability to set a custom HTTPClient for a request.
func NewUsersGroupsListParamsWithHTTPClient(client *http.Client) *UsersGroupsListParams {
	return &UsersGroupsListParams{
		HTTPClient: client,
	}
}

/* UsersGroupsListParams contains all the parameters to send to the API endpoint
   for the users groups list operation.

   Typically these are written to a http.Request.
*/
type UsersGroupsListParams struct {

	// ID.
	ID *string

	// IDGt.
	IDGt *string

	// IDGte.
	IDGte *string

	// IDLt.
	IDLt *string

	// IDLte.
	IDLte *string

	// IDn.
	IDn *string

	/* Limit.

	   Number of results to return per page.
	*/
	Limit *int64

	// Name.
	Name *string

	// NameIc.
	NameIc *string

	// NameIe.
	NameIe *string

	// NameIew.
	NameIew *string

	// NameIsw.
	NameIsw *string

	// Namen.
	Namen *string

	// NameNic.
	NameNic *string

	// NameNie.
	NameNie *string

	// NameNiew.
	NameNiew *string

	// NameNisw.
	NameNisw *string

	/* Offset.

	   The initial index from which to return the results.
	*/
	Offset *int64

	// Q.
	Q *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the users groups list params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *UsersGroupsListParams) WithDefaults() *UsersGroupsListParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the users groups list params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *UsersGroupsListParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the users groups list params
func (o *UsersGroupsListParams) WithTimeout(timeout time.Duration) *UsersGroupsListParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the users groups list params
func (o *UsersGroupsListParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the users groups list params
func (o *UsersGroupsListParams) WithContext(ctx context.Context) *UsersGroupsListParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the users groups list params
func (o *UsersGroupsListParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the users groups list params
func (o *UsersGroupsListParams) WithHTTPClient(client *http.Client) *UsersGroupsListParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the users groups list params
func (o *UsersGroupsListParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the users groups list params
func (o *UsersGroupsListParams) WithID(id *string) *UsersGroupsListParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the users groups list params
func (o *UsersGroupsListParams) SetID(id *string) {
	o.ID = id
}

// WithIDGt adds the iDGt to the users groups list params
func (o *UsersGroupsListParams) WithIDGt(iDGt *string) *UsersGroupsListParams {
	o.SetIDGt(iDGt)
	return o
}

// SetIDGt adds the idGt to the users groups list params
func (o *UsersGroupsListParams) SetIDGt(iDGt *string) {
	o.IDGt = iDGt
}

// WithIDGte adds the iDGte to the users groups list params
func (o *UsersGroupsListParams) WithIDGte(iDGte *string) *UsersGroupsListParams {
	o.SetIDGte(iDGte)
	return o
}

// SetIDGte adds the idGte to the users groups list params
func (o *UsersGroupsListParams) SetIDGte(iDGte *string) {
	o.IDGte = iDGte
}

// WithIDLt adds the iDLt to the users groups list params
func (o *UsersGroupsListParams) WithIDLt(iDLt *string) *UsersGroupsListParams {
	o.SetIDLt(iDLt)
	return o
}

// SetIDLt adds the idLt to the users groups list params
func (o *UsersGroupsListParams) SetIDLt(iDLt *string) {
	o.IDLt = iDLt
}

// WithIDLte adds the iDLte to the users groups list params
func (o *UsersGroupsListParams) WithIDLte(iDLte *string) *UsersGroupsListParams {
	o.SetIDLte(iDLte)
	return o
}

// SetIDLte adds the idLte to the users groups list params
func (o *UsersGroupsListParams) SetIDLte(iDLte *string) {
	o.IDLte = iDLte
}

// WithIDn adds the iDn to the users groups list params
func (o *UsersGroupsListParams) WithIDn(iDn *string) *UsersGroupsListParams {
	o.SetIDn(iDn)
	return o
}

// SetIDn adds the idN to the users groups list params
func (o *UsersGroupsListParams) SetIDn(iDn *string) {
	o.IDn = iDn
}

// WithLimit adds the limit to the users groups list params
func (o *UsersGroupsListParams) WithLimit(limit *int64) *UsersGroupsListParams {
	o.SetLimit(limit)
	return o
}

// SetLimit adds the limit to the users groups list params
func (o *UsersGroupsListParams) SetLimit(limit *int64) {
	o.Limit = limit
}

// WithName adds the name to the users groups list params
func (o *UsersGroupsListParams) WithName(name *string) *UsersGroupsListParams {
	o.SetName(name)
	return o
}

// SetName adds the name to the users groups list params
func (o *UsersGroupsListParams) SetName(name *string) {
	o.Name = name
}

// WithNameIc adds the nameIc to the users groups list params
func (o *UsersGroupsListParams) WithNameIc(nameIc *string) *UsersGroupsListParams {
	o.SetNameIc(nameIc)
	return o
}

// SetNameIc adds the nameIc to the users groups list params
func (o *UsersGroupsListParams) SetNameIc(nameIc *string) {
	o.NameIc = nameIc
}

// WithNameIe adds the nameIe to the users groups list params
func (o *UsersGroupsListParams) WithNameIe(nameIe *string) *UsersGroupsListParams {
	o.SetNameIe(nameIe)
	return o
}

// SetNameIe adds the nameIe to the users groups list params
func (o *UsersGroupsListParams) SetNameIe(nameIe *string) {
	o.NameIe = nameIe
}

// WithNameIew adds the nameIew to the users groups list params
func (o *UsersGroupsListParams) WithNameIew(nameIew *string) *UsersGroupsListParams {
	o.SetNameIew(nameIew)
	return o
}

// SetNameIew adds the nameIew to the users groups list params
func (o *UsersGroupsListParams) SetNameIew(nameIew *string) {
	o.NameIew = nameIew
}

// WithNameIsw adds the nameIsw to the users groups list params
func (o *UsersGroupsListParams) WithNameIsw(nameIsw *string) *UsersGroupsListParams {
	o.SetNameIsw(nameIsw)
	return o
}

// SetNameIsw adds the nameIsw to the users groups list params
func (o *UsersGroupsListParams) SetNameIsw(nameIsw *string) {
	o.NameIsw = nameIsw
}

// WithNamen adds the namen to the users groups list params
func (o *UsersGroupsListParams) WithNamen(namen *string) *UsersGroupsListParams {
	o.SetNamen(namen)
	return o
}

// SetNamen adds the nameN to the users groups list params
func (o *UsersGroupsListParams) SetNamen(namen *string) {
	o.Namen = namen
}

// WithNameNic adds the nameNic to the users groups list params
func (o *UsersGroupsListParams) WithNameNic(nameNic *string) *UsersGroupsListParams {
	o.SetNameNic(nameNic)
	return o
}

// SetNameNic adds the nameNic to the users groups list params
func (o *UsersGroupsListParams) SetNameNic(nameNic *string) {
	o.NameNic = nameNic
}

// WithNameNie adds the nameNie to the users groups list params
func (o *UsersGroupsListParams) WithNameNie(nameNie *string) *UsersGroupsListParams {
	o.SetNameNie(nameNie)
	return o
}

// SetNameNie adds the nameNie to the users groups list params
func (o *UsersGroupsListParams) SetNameNie(nameNie *string) {
	o.NameNie = nameNie
}

// WithNameNiew adds the nameNiew to the users groups list params
func (o *UsersGroupsListParams) WithNameNiew(nameNiew *string) *UsersGroupsListParams {
	o.SetNameNiew(nameNiew)
	return o
}

// SetNameNiew adds the nameNiew to the users groups list params
func (o *UsersGroupsListParams) SetNameNiew(nameNiew *string) {
	o.NameNiew = nameNiew
}

// WithNameNisw adds the nameNisw to the users groups list params
func (o *UsersGroupsListParams) WithNameNisw(nameNisw *string) *UsersGroupsListParams {
	o.SetNameNisw(nameNisw)
	return o
}

// SetNameNisw adds the nameNisw to the users groups list params
func (o *UsersGroupsListParams) SetNameNisw(nameNisw *string) {
	o.NameNisw = nameNisw
}

// WithOffset adds the offset to the users groups list params
func (o *UsersGroupsListParams) WithOffset(offset *int64) *UsersGroupsListParams {
	o.SetOffset(offset)
	return o
}

// SetOffset adds the offset to the users groups list params
func (o *UsersGroupsListParams) SetOffset(offset *int64) {
	o.Offset = offset
}

// WithQ adds the q to the users groups list params
func (o *UsersGroupsListParams) WithQ(q *string) *UsersGroupsListParams {
	o.SetQ(q)
	return o
}

// SetQ adds the q to the users groups list params
func (o *UsersGroupsListParams) SetQ(q *string) {
	o.Q = q
}

// WriteToRequest writes these params to a swagger request
func (o *UsersGroupsListParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.ID != nil {

		// query param id
		var qrID string

		if o.ID != nil {
			qrID = *o.ID
		}
		qID := qrID
		if qID != "" {

			if err := r.SetQueryParam("id", qID); err != nil {
				return err
			}
		}
	}

	if o.IDGt != nil {

		// query param id__gt
		var qrIDGt string

		if o.IDGt != nil {
			qrIDGt = *o.IDGt
		}
		qIDGt := qrIDGt
		if qIDGt != "" {

			if err := r.SetQueryParam("id__gt", qIDGt); err != nil {
				return err
			}
		}
	}

	if o.IDGte != nil {

		// query param id__gte
		var qrIDGte string

		if o.IDGte != nil {
			qrIDGte = *o.IDGte
		}
		qIDGte := qrIDGte
		if qIDGte != "" {

			if err := r.SetQueryParam("id__gte", qIDGte); err != nil {
				return err
			}
		}
	}

	if o.IDLt != nil {

		// query param id__lt
		var qrIDLt string

		if o.IDLt != nil {
			qrIDLt = *o.IDLt
		}
		qIDLt := qrIDLt
		if qIDLt != "" {

			if err := r.SetQueryParam("id__lt", qIDLt); err != nil {
				return err
			}
		}
	}

	if o.IDLte != nil {

		// query param id__lte
		var qrIDLte string

		if o.IDLte != nil {
			qrIDLte = *o.IDLte
		}
		qIDLte := qrIDLte
		if qIDLte != "" {

			if err := r.SetQueryParam("id__lte", qIDLte); err != nil {
				return err
			}
		}
	}

	if o.IDn != nil {

		// query param id__n
		var qrIDn string

		if o.IDn != nil {
			qrIDn = *o.IDn
		}
		qIDn := qrIDn
		if qIDn != "" {

			if err := r.SetQueryParam("id__n", qIDn); err != nil {
				return err
			}
		}
	}

	if o.Limit != nil {

		// query param limit
		var qrLimit int64

		if o.Limit != nil {
			qrLimit = *o.Limit
		}
		qLimit := swag.FormatInt64(qrLimit)
		if qLimit != "" {

			if err := r.SetQueryParam("limit", qLimit); err != nil {
				return err
			}
		}
	}

	if o.Name != nil {

		// query param name
		var qrName string

		if o.Name != nil {
			qrName = *o.Name
		}
		qName := qrName
		if qName != "" {

			if err := r.SetQueryParam("name", qName); err != nil {
				return err
			}
		}
	}

	if o.NameIc != nil {

		// query param name__ic
		var qrNameIc string

		if o.NameIc != nil {
			qrNameIc = *o.NameIc
		}
		qNameIc := qrNameIc
		if qNameIc != "" {

			if err := r.SetQueryParam("name__ic", qNameIc); err != nil {
				return err
			}
		}
	}

	if o.NameIe != nil {

		// query param name__ie
		var qrNameIe string

		if o.NameIe != nil {
			qrNameIe = *o.NameIe
		}
		qNameIe := qrNameIe
		if qNameIe != "" {

			if err := r.SetQueryParam("name__ie", qNameIe); err != nil {
				return err
			}
		}
	}

	if o.NameIew != nil {

		// query param name__iew
		var qrNameIew string

		if o.NameIew != nil {
			qrNameIew = *o.NameIew
		}
		qNameIew := qrNameIew
		if qNameIew != "" {

			if err := r.SetQueryParam("name__iew", qNameIew); err != nil {
				return err
			}
		}
	}

	if o.NameIsw != nil {

		// query param name__isw
		var qrNameIsw string

		if o.NameIsw != nil {
			qrNameIsw = *o.NameIsw
		}
		qNameIsw := qrNameIsw
		if qNameIsw != "" {

			if err := r.SetQueryParam("name__isw", qNameIsw); err != nil {
				return err
			}
		}
	}

	if o.Namen != nil {

		// query param name__n
		var qrNamen string

		if o.Namen != nil {
			qrNamen = *o.Namen
		}
		qNamen := qrNamen
		if qNamen != "" {

			if err := r.SetQueryParam("name__n", qNamen); err != nil {
				return err
			}
		}
	}

	if o.NameNic != nil {

		// query param name__nic
		var qrNameNic string

		if o.NameNic != nil {
			qrNameNic = *o.NameNic
		}
		qNameNic := qrNameNic
		if qNameNic != "" {

			if err := r.SetQueryParam("name__nic", qNameNic); err != nil {
				return err
			}
		}
	}

	if o.NameNie != nil {

		// query param name__nie
		var qrNameNie string

		if o.NameNie != nil {
			qrNameNie = *o.NameNie
		}
		qNameNie := qrNameNie
		if qNameNie != "" {

			if err := r.SetQueryParam("name__nie", qNameNie); err != nil {
				return err
			}
		}
	}

	if o.NameNiew != nil {

		// query param name__niew
		var qrNameNiew string

		if o.NameNiew != nil {
			qrNameNiew = *o.NameNiew
		}
		qNameNiew := qrNameNiew
		if qNameNiew != "" {

			if err := r.SetQueryParam("name__niew", qNameNiew); err != nil {
				return err
			}
		}
	}

	if o.NameNisw != nil {

		// query param name__nisw
		var qrNameNisw string

		if o.NameNisw != nil {
			qrNameNisw = *o.NameNisw
		}
		qNameNisw := qrNameNisw
		if qNameNisw != "" {

			if err := r.SetQueryParam("name__nisw", qNameNisw); err != nil {
				return err
			}
		}
	}

	if o.Offset != nil {

		// query param offset
		var qrOffset int64

		if o.Offset != nil {
			qrOffset = *o.Offset
		}
		qOffset := swag.FormatInt64(qrOffset)
		if qOffset != "" {

			if err := r.SetQueryParam("offset", qOffset); err != nil {
				return err
			}
		}
	}

	if o.Q != nil {

		// query param q
		var qrQ string

		if o.Q != nil {
			qrQ = *o.Q
		}
		qQ := qrQ
		if qQ != "" {

			if err := r.SetQueryParam("q", qQ); err != nil {
				return err
			}
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
