package extras

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

// NewExtrasJobResultsListParams creates a new ExtrasJobResultsListParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewExtrasJobResultsListParams() *ExtrasJobResultsListParams {
	return &ExtrasJobResultsListParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewExtrasJobResultsListParamsWithTimeout creates a new ExtrasJobResultsListParams object
// with the ability to set a timeout on a request.
func NewExtrasJobResultsListParamsWithTimeout(timeout time.Duration) *ExtrasJobResultsListParams {
	return &ExtrasJobResultsListParams{
		timeout: timeout,
	}
}

// NewExtrasJobResultsListParamsWithContext creates a new ExtrasJobResultsListParams object
// with the ability to set a context for a request.
func NewExtrasJobResultsListParamsWithContext(ctx context.Context) *ExtrasJobResultsListParams {
	return &ExtrasJobResultsListParams{
		Context: ctx,
	}
}

// NewExtrasJobResultsListParamsWithHTTPClient creates a new ExtrasJobResultsListParams object
// with the ability to set a custom HTTPClient for a request.
func NewExtrasJobResultsListParamsWithHTTPClient(client *http.Client) *ExtrasJobResultsListParams {
	return &ExtrasJobResultsListParams{
		HTTPClient: client,
	}
}

/* ExtrasJobResultsListParams contains all the parameters to send to the API endpoint
   for the extras job results list operation.

   Typically these are written to a http.Request.
*/
type ExtrasJobResultsListParams struct {

	// Completed.
	Completed *string

	// Created.
	Created *string

	// ID.
	ID *string

	// IDIc.
	IDIc *string

	// IDIe.
	IDIe *string

	// IDIew.
	IDIew *string

	// IDIsw.
	IDIsw *string

	// IDn.
	IDn *string

	// IDNic.
	IDNic *string

	// IDNie.
	IDNie *string

	// IDNiew.
	IDNiew *string

	// IDNisw.
	IDNisw *string

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

	// ObjType.
	ObjType *string

	// ObjTypen.
	ObjTypen *string

	/* Offset.

	   The initial index from which to return the results.
	*/
	Offset *int64

	// Q.
	Q *string

	// Status.
	Status *string

	// Statusn.
	Statusn *string

	// User.
	User *string

	// Usern.
	Usern *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the extras job results list params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ExtrasJobResultsListParams) WithDefaults() *ExtrasJobResultsListParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the extras job results list params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ExtrasJobResultsListParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the extras job results list params
func (o *ExtrasJobResultsListParams) WithTimeout(timeout time.Duration) *ExtrasJobResultsListParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the extras job results list params
func (o *ExtrasJobResultsListParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the extras job results list params
func (o *ExtrasJobResultsListParams) WithContext(ctx context.Context) *ExtrasJobResultsListParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the extras job results list params
func (o *ExtrasJobResultsListParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the extras job results list params
func (o *ExtrasJobResultsListParams) WithHTTPClient(client *http.Client) *ExtrasJobResultsListParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the extras job results list params
func (o *ExtrasJobResultsListParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithCompleted adds the completed to the extras job results list params
func (o *ExtrasJobResultsListParams) WithCompleted(completed *string) *ExtrasJobResultsListParams {
	o.SetCompleted(completed)
	return o
}

// SetCompleted adds the completed to the extras job results list params
func (o *ExtrasJobResultsListParams) SetCompleted(completed *string) {
	o.Completed = completed
}

// WithCreated adds the created to the extras job results list params
func (o *ExtrasJobResultsListParams) WithCreated(created *string) *ExtrasJobResultsListParams {
	o.SetCreated(created)
	return o
}

// SetCreated adds the created to the extras job results list params
func (o *ExtrasJobResultsListParams) SetCreated(created *string) {
	o.Created = created
}

// WithID adds the id to the extras job results list params
func (o *ExtrasJobResultsListParams) WithID(id *string) *ExtrasJobResultsListParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the extras job results list params
func (o *ExtrasJobResultsListParams) SetID(id *string) {
	o.ID = id
}

// WithIDIc adds the iDIc to the extras job results list params
func (o *ExtrasJobResultsListParams) WithIDIc(iDIc *string) *ExtrasJobResultsListParams {
	o.SetIDIc(iDIc)
	return o
}

// SetIDIc adds the idIc to the extras job results list params
func (o *ExtrasJobResultsListParams) SetIDIc(iDIc *string) {
	o.IDIc = iDIc
}

// WithIDIe adds the iDIe to the extras job results list params
func (o *ExtrasJobResultsListParams) WithIDIe(iDIe *string) *ExtrasJobResultsListParams {
	o.SetIDIe(iDIe)
	return o
}

// SetIDIe adds the idIe to the extras job results list params
func (o *ExtrasJobResultsListParams) SetIDIe(iDIe *string) {
	o.IDIe = iDIe
}

// WithIDIew adds the iDIew to the extras job results list params
func (o *ExtrasJobResultsListParams) WithIDIew(iDIew *string) *ExtrasJobResultsListParams {
	o.SetIDIew(iDIew)
	return o
}

// SetIDIew adds the idIew to the extras job results list params
func (o *ExtrasJobResultsListParams) SetIDIew(iDIew *string) {
	o.IDIew = iDIew
}

// WithIDIsw adds the iDIsw to the extras job results list params
func (o *ExtrasJobResultsListParams) WithIDIsw(iDIsw *string) *ExtrasJobResultsListParams {
	o.SetIDIsw(iDIsw)
	return o
}

// SetIDIsw adds the idIsw to the extras job results list params
func (o *ExtrasJobResultsListParams) SetIDIsw(iDIsw *string) {
	o.IDIsw = iDIsw
}

// WithIDn adds the iDn to the extras job results list params
func (o *ExtrasJobResultsListParams) WithIDn(iDn *string) *ExtrasJobResultsListParams {
	o.SetIDn(iDn)
	return o
}

// SetIDn adds the idN to the extras job results list params
func (o *ExtrasJobResultsListParams) SetIDn(iDn *string) {
	o.IDn = iDn
}

// WithIDNic adds the iDNic to the extras job results list params
func (o *ExtrasJobResultsListParams) WithIDNic(iDNic *string) *ExtrasJobResultsListParams {
	o.SetIDNic(iDNic)
	return o
}

// SetIDNic adds the idNic to the extras job results list params
func (o *ExtrasJobResultsListParams) SetIDNic(iDNic *string) {
	o.IDNic = iDNic
}

// WithIDNie adds the iDNie to the extras job results list params
func (o *ExtrasJobResultsListParams) WithIDNie(iDNie *string) *ExtrasJobResultsListParams {
	o.SetIDNie(iDNie)
	return o
}

// SetIDNie adds the idNie to the extras job results list params
func (o *ExtrasJobResultsListParams) SetIDNie(iDNie *string) {
	o.IDNie = iDNie
}

// WithIDNiew adds the iDNiew to the extras job results list params
func (o *ExtrasJobResultsListParams) WithIDNiew(iDNiew *string) *ExtrasJobResultsListParams {
	o.SetIDNiew(iDNiew)
	return o
}

// SetIDNiew adds the idNiew to the extras job results list params
func (o *ExtrasJobResultsListParams) SetIDNiew(iDNiew *string) {
	o.IDNiew = iDNiew
}

// WithIDNisw adds the iDNisw to the extras job results list params
func (o *ExtrasJobResultsListParams) WithIDNisw(iDNisw *string) *ExtrasJobResultsListParams {
	o.SetIDNisw(iDNisw)
	return o
}

// SetIDNisw adds the idNisw to the extras job results list params
func (o *ExtrasJobResultsListParams) SetIDNisw(iDNisw *string) {
	o.IDNisw = iDNisw
}

// WithLimit adds the limit to the extras job results list params
func (o *ExtrasJobResultsListParams) WithLimit(limit *int64) *ExtrasJobResultsListParams {
	o.SetLimit(limit)
	return o
}

// SetLimit adds the limit to the extras job results list params
func (o *ExtrasJobResultsListParams) SetLimit(limit *int64) {
	o.Limit = limit
}

// WithName adds the name to the extras job results list params
func (o *ExtrasJobResultsListParams) WithName(name *string) *ExtrasJobResultsListParams {
	o.SetName(name)
	return o
}

// SetName adds the name to the extras job results list params
func (o *ExtrasJobResultsListParams) SetName(name *string) {
	o.Name = name
}

// WithNameIc adds the nameIc to the extras job results list params
func (o *ExtrasJobResultsListParams) WithNameIc(nameIc *string) *ExtrasJobResultsListParams {
	o.SetNameIc(nameIc)
	return o
}

// SetNameIc adds the nameIc to the extras job results list params
func (o *ExtrasJobResultsListParams) SetNameIc(nameIc *string) {
	o.NameIc = nameIc
}

// WithNameIe adds the nameIe to the extras job results list params
func (o *ExtrasJobResultsListParams) WithNameIe(nameIe *string) *ExtrasJobResultsListParams {
	o.SetNameIe(nameIe)
	return o
}

// SetNameIe adds the nameIe to the extras job results list params
func (o *ExtrasJobResultsListParams) SetNameIe(nameIe *string) {
	o.NameIe = nameIe
}

// WithNameIew adds the nameIew to the extras job results list params
func (o *ExtrasJobResultsListParams) WithNameIew(nameIew *string) *ExtrasJobResultsListParams {
	o.SetNameIew(nameIew)
	return o
}

// SetNameIew adds the nameIew to the extras job results list params
func (o *ExtrasJobResultsListParams) SetNameIew(nameIew *string) {
	o.NameIew = nameIew
}

// WithNameIsw adds the nameIsw to the extras job results list params
func (o *ExtrasJobResultsListParams) WithNameIsw(nameIsw *string) *ExtrasJobResultsListParams {
	o.SetNameIsw(nameIsw)
	return o
}

// SetNameIsw adds the nameIsw to the extras job results list params
func (o *ExtrasJobResultsListParams) SetNameIsw(nameIsw *string) {
	o.NameIsw = nameIsw
}

// WithNamen adds the namen to the extras job results list params
func (o *ExtrasJobResultsListParams) WithNamen(namen *string) *ExtrasJobResultsListParams {
	o.SetNamen(namen)
	return o
}

// SetNamen adds the nameN to the extras job results list params
func (o *ExtrasJobResultsListParams) SetNamen(namen *string) {
	o.Namen = namen
}

// WithNameNic adds the nameNic to the extras job results list params
func (o *ExtrasJobResultsListParams) WithNameNic(nameNic *string) *ExtrasJobResultsListParams {
	o.SetNameNic(nameNic)
	return o
}

// SetNameNic adds the nameNic to the extras job results list params
func (o *ExtrasJobResultsListParams) SetNameNic(nameNic *string) {
	o.NameNic = nameNic
}

// WithNameNie adds the nameNie to the extras job results list params
func (o *ExtrasJobResultsListParams) WithNameNie(nameNie *string) *ExtrasJobResultsListParams {
	o.SetNameNie(nameNie)
	return o
}

// SetNameNie adds the nameNie to the extras job results list params
func (o *ExtrasJobResultsListParams) SetNameNie(nameNie *string) {
	o.NameNie = nameNie
}

// WithNameNiew adds the nameNiew to the extras job results list params
func (o *ExtrasJobResultsListParams) WithNameNiew(nameNiew *string) *ExtrasJobResultsListParams {
	o.SetNameNiew(nameNiew)
	return o
}

// SetNameNiew adds the nameNiew to the extras job results list params
func (o *ExtrasJobResultsListParams) SetNameNiew(nameNiew *string) {
	o.NameNiew = nameNiew
}

// WithNameNisw adds the nameNisw to the extras job results list params
func (o *ExtrasJobResultsListParams) WithNameNisw(nameNisw *string) *ExtrasJobResultsListParams {
	o.SetNameNisw(nameNisw)
	return o
}

// SetNameNisw adds the nameNisw to the extras job results list params
func (o *ExtrasJobResultsListParams) SetNameNisw(nameNisw *string) {
	o.NameNisw = nameNisw
}

// WithObjType adds the objType to the extras job results list params
func (o *ExtrasJobResultsListParams) WithObjType(objType *string) *ExtrasJobResultsListParams {
	o.SetObjType(objType)
	return o
}

// SetObjType adds the objType to the extras job results list params
func (o *ExtrasJobResultsListParams) SetObjType(objType *string) {
	o.ObjType = objType
}

// WithObjTypen adds the objTypen to the extras job results list params
func (o *ExtrasJobResultsListParams) WithObjTypen(objTypen *string) *ExtrasJobResultsListParams {
	o.SetObjTypen(objTypen)
	return o
}

// SetObjTypen adds the objTypeN to the extras job results list params
func (o *ExtrasJobResultsListParams) SetObjTypen(objTypen *string) {
	o.ObjTypen = objTypen
}

// WithOffset adds the offset to the extras job results list params
func (o *ExtrasJobResultsListParams) WithOffset(offset *int64) *ExtrasJobResultsListParams {
	o.SetOffset(offset)
	return o
}

// SetOffset adds the offset to the extras job results list params
func (o *ExtrasJobResultsListParams) SetOffset(offset *int64) {
	o.Offset = offset
}

// WithQ adds the q to the extras job results list params
func (o *ExtrasJobResultsListParams) WithQ(q *string) *ExtrasJobResultsListParams {
	o.SetQ(q)
	return o
}

// SetQ adds the q to the extras job results list params
func (o *ExtrasJobResultsListParams) SetQ(q *string) {
	o.Q = q
}

// WithStatus adds the status to the extras job results list params
func (o *ExtrasJobResultsListParams) WithStatus(status *string) *ExtrasJobResultsListParams {
	o.SetStatus(status)
	return o
}

// SetStatus adds the status to the extras job results list params
func (o *ExtrasJobResultsListParams) SetStatus(status *string) {
	o.Status = status
}

// WithStatusn adds the statusn to the extras job results list params
func (o *ExtrasJobResultsListParams) WithStatusn(statusn *string) *ExtrasJobResultsListParams {
	o.SetStatusn(statusn)
	return o
}

// SetStatusn adds the statusN to the extras job results list params
func (o *ExtrasJobResultsListParams) SetStatusn(statusn *string) {
	o.Statusn = statusn
}

// WithUser adds the user to the extras job results list params
func (o *ExtrasJobResultsListParams) WithUser(user *string) *ExtrasJobResultsListParams {
	o.SetUser(user)
	return o
}

// SetUser adds the user to the extras job results list params
func (o *ExtrasJobResultsListParams) SetUser(user *string) {
	o.User = user
}

// WithUsern adds the usern to the extras job results list params
func (o *ExtrasJobResultsListParams) WithUsern(usern *string) *ExtrasJobResultsListParams {
	o.SetUsern(usern)
	return o
}

// SetUsern adds the userN to the extras job results list params
func (o *ExtrasJobResultsListParams) SetUsern(usern *string) {
	o.Usern = usern
}

// WriteToRequest writes these params to a swagger request
func (o *ExtrasJobResultsListParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Completed != nil {

		// query param completed
		var qrCompleted string

		if o.Completed != nil {
			qrCompleted = *o.Completed
		}
		qCompleted := qrCompleted
		if qCompleted != "" {

			if err := r.SetQueryParam("completed", qCompleted); err != nil {
				return err
			}
		}
	}

	if o.Created != nil {

		// query param created
		var qrCreated string

		if o.Created != nil {
			qrCreated = *o.Created
		}
		qCreated := qrCreated
		if qCreated != "" {

			if err := r.SetQueryParam("created", qCreated); err != nil {
				return err
			}
		}
	}

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

	if o.IDIc != nil {

		// query param id__ic
		var qrIDIc string

		if o.IDIc != nil {
			qrIDIc = *o.IDIc
		}
		qIDIc := qrIDIc
		if qIDIc != "" {

			if err := r.SetQueryParam("id__ic", qIDIc); err != nil {
				return err
			}
		}
	}

	if o.IDIe != nil {

		// query param id__ie
		var qrIDIe string

		if o.IDIe != nil {
			qrIDIe = *o.IDIe
		}
		qIDIe := qrIDIe
		if qIDIe != "" {

			if err := r.SetQueryParam("id__ie", qIDIe); err != nil {
				return err
			}
		}
	}

	if o.IDIew != nil {

		// query param id__iew
		var qrIDIew string

		if o.IDIew != nil {
			qrIDIew = *o.IDIew
		}
		qIDIew := qrIDIew
		if qIDIew != "" {

			if err := r.SetQueryParam("id__iew", qIDIew); err != nil {
				return err
			}
		}
	}

	if o.IDIsw != nil {

		// query param id__isw
		var qrIDIsw string

		if o.IDIsw != nil {
			qrIDIsw = *o.IDIsw
		}
		qIDIsw := qrIDIsw
		if qIDIsw != "" {

			if err := r.SetQueryParam("id__isw", qIDIsw); err != nil {
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

	if o.IDNic != nil {

		// query param id__nic
		var qrIDNic string

		if o.IDNic != nil {
			qrIDNic = *o.IDNic
		}
		qIDNic := qrIDNic
		if qIDNic != "" {

			if err := r.SetQueryParam("id__nic", qIDNic); err != nil {
				return err
			}
		}
	}

	if o.IDNie != nil {

		// query param id__nie
		var qrIDNie string

		if o.IDNie != nil {
			qrIDNie = *o.IDNie
		}
		qIDNie := qrIDNie
		if qIDNie != "" {

			if err := r.SetQueryParam("id__nie", qIDNie); err != nil {
				return err
			}
		}
	}

	if o.IDNiew != nil {

		// query param id__niew
		var qrIDNiew string

		if o.IDNiew != nil {
			qrIDNiew = *o.IDNiew
		}
		qIDNiew := qrIDNiew
		if qIDNiew != "" {

			if err := r.SetQueryParam("id__niew", qIDNiew); err != nil {
				return err
			}
		}
	}

	if o.IDNisw != nil {

		// query param id__nisw
		var qrIDNisw string

		if o.IDNisw != nil {
			qrIDNisw = *o.IDNisw
		}
		qIDNisw := qrIDNisw
		if qIDNisw != "" {

			if err := r.SetQueryParam("id__nisw", qIDNisw); err != nil {
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

	if o.ObjType != nil {

		// query param obj_type
		var qrObjType string

		if o.ObjType != nil {
			qrObjType = *o.ObjType
		}
		qObjType := qrObjType
		if qObjType != "" {

			if err := r.SetQueryParam("obj_type", qObjType); err != nil {
				return err
			}
		}
	}

	if o.ObjTypen != nil {

		// query param obj_type__n
		var qrObjTypen string

		if o.ObjTypen != nil {
			qrObjTypen = *o.ObjTypen
		}
		qObjTypen := qrObjTypen
		if qObjTypen != "" {

			if err := r.SetQueryParam("obj_type__n", qObjTypen); err != nil {
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

	if o.Status != nil {

		// query param status
		var qrStatus string

		if o.Status != nil {
			qrStatus = *o.Status
		}
		qStatus := qrStatus
		if qStatus != "" {

			if err := r.SetQueryParam("status", qStatus); err != nil {
				return err
			}
		}
	}

	if o.Statusn != nil {

		// query param status__n
		var qrStatusn string

		if o.Statusn != nil {
			qrStatusn = *o.Statusn
		}
		qStatusn := qrStatusn
		if qStatusn != "" {

			if err := r.SetQueryParam("status__n", qStatusn); err != nil {
				return err
			}
		}
	}

	if o.User != nil {

		// query param user
		var qrUser string

		if o.User != nil {
			qrUser = *o.User
		}
		qUser := qrUser
		if qUser != "" {

			if err := r.SetQueryParam("user", qUser); err != nil {
				return err
			}
		}
	}

	if o.Usern != nil {

		// query param user__n
		var qrUsern string

		if o.Usern != nil {
			qrUsern = *o.Usern
		}
		qUsern := qrUsern
		if qUsern != "" {

			if err := r.SetQueryParam("user__n", qUsern); err != nil {
				return err
			}
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
