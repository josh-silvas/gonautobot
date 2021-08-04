package dcim

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

// NewDcimConsolePortTemplatesListParams creates a new DcimConsolePortTemplatesListParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewDcimConsolePortTemplatesListParams() *DcimConsolePortTemplatesListParams {
	return &DcimConsolePortTemplatesListParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewDcimConsolePortTemplatesListParamsWithTimeout creates a new DcimConsolePortTemplatesListParams object
// with the ability to set a timeout on a request.
func NewDcimConsolePortTemplatesListParamsWithTimeout(timeout time.Duration) *DcimConsolePortTemplatesListParams {
	return &DcimConsolePortTemplatesListParams{
		timeout: timeout,
	}
}

// NewDcimConsolePortTemplatesListParamsWithContext creates a new DcimConsolePortTemplatesListParams object
// with the ability to set a context for a request.
func NewDcimConsolePortTemplatesListParamsWithContext(ctx context.Context) *DcimConsolePortTemplatesListParams {
	return &DcimConsolePortTemplatesListParams{
		Context: ctx,
	}
}

// NewDcimConsolePortTemplatesListParamsWithHTTPClient creates a new DcimConsolePortTemplatesListParams object
// with the ability to set a custom HTTPClient for a request.
func NewDcimConsolePortTemplatesListParamsWithHTTPClient(client *http.Client) *DcimConsolePortTemplatesListParams {
	return &DcimConsolePortTemplatesListParams{
		HTTPClient: client,
	}
}

/* DcimConsolePortTemplatesListParams contains all the parameters to send to the API endpoint
   for the dcim console port templates list operation.

   Typically these are written to a http.Request.
*/
type DcimConsolePortTemplatesListParams struct {

	// DevicetypeID.
	DevicetypeID *string

	// DevicetypeIDn.
	DevicetypeIDn *string

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

	/* Offset.

	   The initial index from which to return the results.
	*/
	Offset *int64

	// Q.
	Q *string

	// Type.
	Type *string

	// Typen.
	Typen *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the dcim console port templates list params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DcimConsolePortTemplatesListParams) WithDefaults() *DcimConsolePortTemplatesListParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the dcim console port templates list params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DcimConsolePortTemplatesListParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the dcim console port templates list params
func (o *DcimConsolePortTemplatesListParams) WithTimeout(timeout time.Duration) *DcimConsolePortTemplatesListParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the dcim console port templates list params
func (o *DcimConsolePortTemplatesListParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the dcim console port templates list params
func (o *DcimConsolePortTemplatesListParams) WithContext(ctx context.Context) *DcimConsolePortTemplatesListParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the dcim console port templates list params
func (o *DcimConsolePortTemplatesListParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the dcim console port templates list params
func (o *DcimConsolePortTemplatesListParams) WithHTTPClient(client *http.Client) *DcimConsolePortTemplatesListParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the dcim console port templates list params
func (o *DcimConsolePortTemplatesListParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithDevicetypeID adds the devicetypeID to the dcim console port templates list params
func (o *DcimConsolePortTemplatesListParams) WithDevicetypeID(devicetypeID *string) *DcimConsolePortTemplatesListParams {
	o.SetDevicetypeID(devicetypeID)
	return o
}

// SetDevicetypeID adds the devicetypeId to the dcim console port templates list params
func (o *DcimConsolePortTemplatesListParams) SetDevicetypeID(devicetypeID *string) {
	o.DevicetypeID = devicetypeID
}

// WithDevicetypeIDn adds the devicetypeIDn to the dcim console port templates list params
func (o *DcimConsolePortTemplatesListParams) WithDevicetypeIDn(devicetypeIDn *string) *DcimConsolePortTemplatesListParams {
	o.SetDevicetypeIDn(devicetypeIDn)
	return o
}

// SetDevicetypeIDn adds the devicetypeIdN to the dcim console port templates list params
func (o *DcimConsolePortTemplatesListParams) SetDevicetypeIDn(devicetypeIDn *string) {
	o.DevicetypeIDn = devicetypeIDn
}

// WithID adds the id to the dcim console port templates list params
func (o *DcimConsolePortTemplatesListParams) WithID(id *string) *DcimConsolePortTemplatesListParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the dcim console port templates list params
func (o *DcimConsolePortTemplatesListParams) SetID(id *string) {
	o.ID = id
}

// WithIDIc adds the iDIc to the dcim console port templates list params
func (o *DcimConsolePortTemplatesListParams) WithIDIc(iDIc *string) *DcimConsolePortTemplatesListParams {
	o.SetIDIc(iDIc)
	return o
}

// SetIDIc adds the idIc to the dcim console port templates list params
func (o *DcimConsolePortTemplatesListParams) SetIDIc(iDIc *string) {
	o.IDIc = iDIc
}

// WithIDIe adds the iDIe to the dcim console port templates list params
func (o *DcimConsolePortTemplatesListParams) WithIDIe(iDIe *string) *DcimConsolePortTemplatesListParams {
	o.SetIDIe(iDIe)
	return o
}

// SetIDIe adds the idIe to the dcim console port templates list params
func (o *DcimConsolePortTemplatesListParams) SetIDIe(iDIe *string) {
	o.IDIe = iDIe
}

// WithIDIew adds the iDIew to the dcim console port templates list params
func (o *DcimConsolePortTemplatesListParams) WithIDIew(iDIew *string) *DcimConsolePortTemplatesListParams {
	o.SetIDIew(iDIew)
	return o
}

// SetIDIew adds the idIew to the dcim console port templates list params
func (o *DcimConsolePortTemplatesListParams) SetIDIew(iDIew *string) {
	o.IDIew = iDIew
}

// WithIDIsw adds the iDIsw to the dcim console port templates list params
func (o *DcimConsolePortTemplatesListParams) WithIDIsw(iDIsw *string) *DcimConsolePortTemplatesListParams {
	o.SetIDIsw(iDIsw)
	return o
}

// SetIDIsw adds the idIsw to the dcim console port templates list params
func (o *DcimConsolePortTemplatesListParams) SetIDIsw(iDIsw *string) {
	o.IDIsw = iDIsw
}

// WithIDn adds the iDn to the dcim console port templates list params
func (o *DcimConsolePortTemplatesListParams) WithIDn(iDn *string) *DcimConsolePortTemplatesListParams {
	o.SetIDn(iDn)
	return o
}

// SetIDn adds the idN to the dcim console port templates list params
func (o *DcimConsolePortTemplatesListParams) SetIDn(iDn *string) {
	o.IDn = iDn
}

// WithIDNic adds the iDNic to the dcim console port templates list params
func (o *DcimConsolePortTemplatesListParams) WithIDNic(iDNic *string) *DcimConsolePortTemplatesListParams {
	o.SetIDNic(iDNic)
	return o
}

// SetIDNic adds the idNic to the dcim console port templates list params
func (o *DcimConsolePortTemplatesListParams) SetIDNic(iDNic *string) {
	o.IDNic = iDNic
}

// WithIDNie adds the iDNie to the dcim console port templates list params
func (o *DcimConsolePortTemplatesListParams) WithIDNie(iDNie *string) *DcimConsolePortTemplatesListParams {
	o.SetIDNie(iDNie)
	return o
}

// SetIDNie adds the idNie to the dcim console port templates list params
func (o *DcimConsolePortTemplatesListParams) SetIDNie(iDNie *string) {
	o.IDNie = iDNie
}

// WithIDNiew adds the iDNiew to the dcim console port templates list params
func (o *DcimConsolePortTemplatesListParams) WithIDNiew(iDNiew *string) *DcimConsolePortTemplatesListParams {
	o.SetIDNiew(iDNiew)
	return o
}

// SetIDNiew adds the idNiew to the dcim console port templates list params
func (o *DcimConsolePortTemplatesListParams) SetIDNiew(iDNiew *string) {
	o.IDNiew = iDNiew
}

// WithIDNisw adds the iDNisw to the dcim console port templates list params
func (o *DcimConsolePortTemplatesListParams) WithIDNisw(iDNisw *string) *DcimConsolePortTemplatesListParams {
	o.SetIDNisw(iDNisw)
	return o
}

// SetIDNisw adds the idNisw to the dcim console port templates list params
func (o *DcimConsolePortTemplatesListParams) SetIDNisw(iDNisw *string) {
	o.IDNisw = iDNisw
}

// WithLimit adds the limit to the dcim console port templates list params
func (o *DcimConsolePortTemplatesListParams) WithLimit(limit *int64) *DcimConsolePortTemplatesListParams {
	o.SetLimit(limit)
	return o
}

// SetLimit adds the limit to the dcim console port templates list params
func (o *DcimConsolePortTemplatesListParams) SetLimit(limit *int64) {
	o.Limit = limit
}

// WithName adds the name to the dcim console port templates list params
func (o *DcimConsolePortTemplatesListParams) WithName(name *string) *DcimConsolePortTemplatesListParams {
	o.SetName(name)
	return o
}

// SetName adds the name to the dcim console port templates list params
func (o *DcimConsolePortTemplatesListParams) SetName(name *string) {
	o.Name = name
}

// WithNameIc adds the nameIc to the dcim console port templates list params
func (o *DcimConsolePortTemplatesListParams) WithNameIc(nameIc *string) *DcimConsolePortTemplatesListParams {
	o.SetNameIc(nameIc)
	return o
}

// SetNameIc adds the nameIc to the dcim console port templates list params
func (o *DcimConsolePortTemplatesListParams) SetNameIc(nameIc *string) {
	o.NameIc = nameIc
}

// WithNameIe adds the nameIe to the dcim console port templates list params
func (o *DcimConsolePortTemplatesListParams) WithNameIe(nameIe *string) *DcimConsolePortTemplatesListParams {
	o.SetNameIe(nameIe)
	return o
}

// SetNameIe adds the nameIe to the dcim console port templates list params
func (o *DcimConsolePortTemplatesListParams) SetNameIe(nameIe *string) {
	o.NameIe = nameIe
}

// WithNameIew adds the nameIew to the dcim console port templates list params
func (o *DcimConsolePortTemplatesListParams) WithNameIew(nameIew *string) *DcimConsolePortTemplatesListParams {
	o.SetNameIew(nameIew)
	return o
}

// SetNameIew adds the nameIew to the dcim console port templates list params
func (o *DcimConsolePortTemplatesListParams) SetNameIew(nameIew *string) {
	o.NameIew = nameIew
}

// WithNameIsw adds the nameIsw to the dcim console port templates list params
func (o *DcimConsolePortTemplatesListParams) WithNameIsw(nameIsw *string) *DcimConsolePortTemplatesListParams {
	o.SetNameIsw(nameIsw)
	return o
}

// SetNameIsw adds the nameIsw to the dcim console port templates list params
func (o *DcimConsolePortTemplatesListParams) SetNameIsw(nameIsw *string) {
	o.NameIsw = nameIsw
}

// WithNamen adds the namen to the dcim console port templates list params
func (o *DcimConsolePortTemplatesListParams) WithNamen(namen *string) *DcimConsolePortTemplatesListParams {
	o.SetNamen(namen)
	return o
}

// SetNamen adds the nameN to the dcim console port templates list params
func (o *DcimConsolePortTemplatesListParams) SetNamen(namen *string) {
	o.Namen = namen
}

// WithNameNic adds the nameNic to the dcim console port templates list params
func (o *DcimConsolePortTemplatesListParams) WithNameNic(nameNic *string) *DcimConsolePortTemplatesListParams {
	o.SetNameNic(nameNic)
	return o
}

// SetNameNic adds the nameNic to the dcim console port templates list params
func (o *DcimConsolePortTemplatesListParams) SetNameNic(nameNic *string) {
	o.NameNic = nameNic
}

// WithNameNie adds the nameNie to the dcim console port templates list params
func (o *DcimConsolePortTemplatesListParams) WithNameNie(nameNie *string) *DcimConsolePortTemplatesListParams {
	o.SetNameNie(nameNie)
	return o
}

// SetNameNie adds the nameNie to the dcim console port templates list params
func (o *DcimConsolePortTemplatesListParams) SetNameNie(nameNie *string) {
	o.NameNie = nameNie
}

// WithNameNiew adds the nameNiew to the dcim console port templates list params
func (o *DcimConsolePortTemplatesListParams) WithNameNiew(nameNiew *string) *DcimConsolePortTemplatesListParams {
	o.SetNameNiew(nameNiew)
	return o
}

// SetNameNiew adds the nameNiew to the dcim console port templates list params
func (o *DcimConsolePortTemplatesListParams) SetNameNiew(nameNiew *string) {
	o.NameNiew = nameNiew
}

// WithNameNisw adds the nameNisw to the dcim console port templates list params
func (o *DcimConsolePortTemplatesListParams) WithNameNisw(nameNisw *string) *DcimConsolePortTemplatesListParams {
	o.SetNameNisw(nameNisw)
	return o
}

// SetNameNisw adds the nameNisw to the dcim console port templates list params
func (o *DcimConsolePortTemplatesListParams) SetNameNisw(nameNisw *string) {
	o.NameNisw = nameNisw
}

// WithOffset adds the offset to the dcim console port templates list params
func (o *DcimConsolePortTemplatesListParams) WithOffset(offset *int64) *DcimConsolePortTemplatesListParams {
	o.SetOffset(offset)
	return o
}

// SetOffset adds the offset to the dcim console port templates list params
func (o *DcimConsolePortTemplatesListParams) SetOffset(offset *int64) {
	o.Offset = offset
}

// WithQ adds the q to the dcim console port templates list params
func (o *DcimConsolePortTemplatesListParams) WithQ(q *string) *DcimConsolePortTemplatesListParams {
	o.SetQ(q)
	return o
}

// SetQ adds the q to the dcim console port templates list params
func (o *DcimConsolePortTemplatesListParams) SetQ(q *string) {
	o.Q = q
}

// WithType adds the typeVar to the dcim console port templates list params
func (o *DcimConsolePortTemplatesListParams) WithType(typeVar *string) *DcimConsolePortTemplatesListParams {
	o.SetType(typeVar)
	return o
}

// SetType adds the type to the dcim console port templates list params
func (o *DcimConsolePortTemplatesListParams) SetType(typeVar *string) {
	o.Type = typeVar
}

// WithTypen adds the typen to the dcim console port templates list params
func (o *DcimConsolePortTemplatesListParams) WithTypen(typen *string) *DcimConsolePortTemplatesListParams {
	o.SetTypen(typen)
	return o
}

// SetTypen adds the typeN to the dcim console port templates list params
func (o *DcimConsolePortTemplatesListParams) SetTypen(typen *string) {
	o.Typen = typen
}

// WriteToRequest writes these params to a swagger request
func (o *DcimConsolePortTemplatesListParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.DevicetypeID != nil {

		// query param devicetype_id
		var qrDevicetypeID string

		if o.DevicetypeID != nil {
			qrDevicetypeID = *o.DevicetypeID
		}
		qDevicetypeID := qrDevicetypeID
		if qDevicetypeID != "" {

			if err := r.SetQueryParam("devicetype_id", qDevicetypeID); err != nil {
				return err
			}
		}
	}

	if o.DevicetypeIDn != nil {

		// query param devicetype_id__n
		var qrDevicetypeIDn string

		if o.DevicetypeIDn != nil {
			qrDevicetypeIDn = *o.DevicetypeIDn
		}
		qDevicetypeIDn := qrDevicetypeIDn
		if qDevicetypeIDn != "" {

			if err := r.SetQueryParam("devicetype_id__n", qDevicetypeIDn); err != nil {
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

	if o.Type != nil {

		// query param type
		var qrType string

		if o.Type != nil {
			qrType = *o.Type
		}
		qType := qrType
		if qType != "" {

			if err := r.SetQueryParam("type", qType); err != nil {
				return err
			}
		}
	}

	if o.Typen != nil {

		// query param type__n
		var qrTypen string

		if o.Typen != nil {
			qrTypen = *o.Typen
		}
		qTypen := qrTypen
		if qTypen != "" {

			if err := r.SetQueryParam("type__n", qTypen); err != nil {
				return err
			}
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
