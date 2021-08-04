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

// NewUsersUsersListParams creates a new UsersUsersListParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewUsersUsersListParams() *UsersUsersListParams {
	return &UsersUsersListParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewUsersUsersListParamsWithTimeout creates a new UsersUsersListParams object
// with the ability to set a timeout on a request.
func NewUsersUsersListParamsWithTimeout(timeout time.Duration) *UsersUsersListParams {
	return &UsersUsersListParams{
		timeout: timeout,
	}
}

// NewUsersUsersListParamsWithContext creates a new UsersUsersListParams object
// with the ability to set a context for a request.
func NewUsersUsersListParamsWithContext(ctx context.Context) *UsersUsersListParams {
	return &UsersUsersListParams{
		Context: ctx,
	}
}

// NewUsersUsersListParamsWithHTTPClient creates a new UsersUsersListParams object
// with the ability to set a custom HTTPClient for a request.
func NewUsersUsersListParamsWithHTTPClient(client *http.Client) *UsersUsersListParams {
	return &UsersUsersListParams{
		HTTPClient: client,
	}
}

/* UsersUsersListParams contains all the parameters to send to the API endpoint
   for the users users list operation.

   Typically these are written to a http.Request.
*/
type UsersUsersListParams struct {

	// Email.
	Email *string

	// EmailIc.
	EmailIc *string

	// EmailIe.
	EmailIe *string

	// EmailIew.
	EmailIew *string

	// EmailIsw.
	EmailIsw *string

	// Emailn.
	Emailn *string

	// EmailNic.
	EmailNic *string

	// EmailNie.
	EmailNie *string

	// EmailNiew.
	EmailNiew *string

	// EmailNisw.
	EmailNisw *string

	// FirstName.
	FirstName *string

	// FirstNameIc.
	FirstNameIc *string

	// FirstNameIe.
	FirstNameIe *string

	// FirstNameIew.
	FirstNameIew *string

	// FirstNameIsw.
	FirstNameIsw *string

	// FirstNamen.
	FirstNamen *string

	// FirstNameNic.
	FirstNameNic *string

	// FirstNameNie.
	FirstNameNie *string

	// FirstNameNiew.
	FirstNameNiew *string

	// FirstNameNisw.
	FirstNameNisw *string

	// Group.
	Group *string

	// Groupn.
	Groupn *string

	// GroupID.
	GroupID *string

	// GroupIDn.
	GroupIDn *string

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

	// IsActive.
	IsActive *string

	// IsStaff.
	IsStaff *string

	// LastName.
	LastName *string

	// LastNameIc.
	LastNameIc *string

	// LastNameIe.
	LastNameIe *string

	// LastNameIew.
	LastNameIew *string

	// LastNameIsw.
	LastNameIsw *string

	// LastNamen.
	LastNamen *string

	// LastNameNic.
	LastNameNic *string

	// LastNameNie.
	LastNameNie *string

	// LastNameNiew.
	LastNameNiew *string

	// LastNameNisw.
	LastNameNisw *string

	/* Limit.

	   Number of results to return per page.
	*/
	Limit *int64

	/* Offset.

	   The initial index from which to return the results.
	*/
	Offset *int64

	// Q.
	Q *string

	// Username.
	Username *string

	// UsernameIc.
	UsernameIc *string

	// UsernameIe.
	UsernameIe *string

	// UsernameIew.
	UsernameIew *string

	// UsernameIsw.
	UsernameIsw *string

	// Usernamen.
	Usernamen *string

	// UsernameNic.
	UsernameNic *string

	// UsernameNie.
	UsernameNie *string

	// UsernameNiew.
	UsernameNiew *string

	// UsernameNisw.
	UsernameNisw *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the users users list params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *UsersUsersListParams) WithDefaults() *UsersUsersListParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the users users list params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *UsersUsersListParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the users users list params
func (o *UsersUsersListParams) WithTimeout(timeout time.Duration) *UsersUsersListParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the users users list params
func (o *UsersUsersListParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the users users list params
func (o *UsersUsersListParams) WithContext(ctx context.Context) *UsersUsersListParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the users users list params
func (o *UsersUsersListParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the users users list params
func (o *UsersUsersListParams) WithHTTPClient(client *http.Client) *UsersUsersListParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the users users list params
func (o *UsersUsersListParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithEmail adds the email to the users users list params
func (o *UsersUsersListParams) WithEmail(email *string) *UsersUsersListParams {
	o.SetEmail(email)
	return o
}

// SetEmail adds the email to the users users list params
func (o *UsersUsersListParams) SetEmail(email *string) {
	o.Email = email
}

// WithEmailIc adds the emailIc to the users users list params
func (o *UsersUsersListParams) WithEmailIc(emailIc *string) *UsersUsersListParams {
	o.SetEmailIc(emailIc)
	return o
}

// SetEmailIc adds the emailIc to the users users list params
func (o *UsersUsersListParams) SetEmailIc(emailIc *string) {
	o.EmailIc = emailIc
}

// WithEmailIe adds the emailIe to the users users list params
func (o *UsersUsersListParams) WithEmailIe(emailIe *string) *UsersUsersListParams {
	o.SetEmailIe(emailIe)
	return o
}

// SetEmailIe adds the emailIe to the users users list params
func (o *UsersUsersListParams) SetEmailIe(emailIe *string) {
	o.EmailIe = emailIe
}

// WithEmailIew adds the emailIew to the users users list params
func (o *UsersUsersListParams) WithEmailIew(emailIew *string) *UsersUsersListParams {
	o.SetEmailIew(emailIew)
	return o
}

// SetEmailIew adds the emailIew to the users users list params
func (o *UsersUsersListParams) SetEmailIew(emailIew *string) {
	o.EmailIew = emailIew
}

// WithEmailIsw adds the emailIsw to the users users list params
func (o *UsersUsersListParams) WithEmailIsw(emailIsw *string) *UsersUsersListParams {
	o.SetEmailIsw(emailIsw)
	return o
}

// SetEmailIsw adds the emailIsw to the users users list params
func (o *UsersUsersListParams) SetEmailIsw(emailIsw *string) {
	o.EmailIsw = emailIsw
}

// WithEmailn adds the emailn to the users users list params
func (o *UsersUsersListParams) WithEmailn(emailn *string) *UsersUsersListParams {
	o.SetEmailn(emailn)
	return o
}

// SetEmailn adds the emailN to the users users list params
func (o *UsersUsersListParams) SetEmailn(emailn *string) {
	o.Emailn = emailn
}

// WithEmailNic adds the emailNic to the users users list params
func (o *UsersUsersListParams) WithEmailNic(emailNic *string) *UsersUsersListParams {
	o.SetEmailNic(emailNic)
	return o
}

// SetEmailNic adds the emailNic to the users users list params
func (o *UsersUsersListParams) SetEmailNic(emailNic *string) {
	o.EmailNic = emailNic
}

// WithEmailNie adds the emailNie to the users users list params
func (o *UsersUsersListParams) WithEmailNie(emailNie *string) *UsersUsersListParams {
	o.SetEmailNie(emailNie)
	return o
}

// SetEmailNie adds the emailNie to the users users list params
func (o *UsersUsersListParams) SetEmailNie(emailNie *string) {
	o.EmailNie = emailNie
}

// WithEmailNiew adds the emailNiew to the users users list params
func (o *UsersUsersListParams) WithEmailNiew(emailNiew *string) *UsersUsersListParams {
	o.SetEmailNiew(emailNiew)
	return o
}

// SetEmailNiew adds the emailNiew to the users users list params
func (o *UsersUsersListParams) SetEmailNiew(emailNiew *string) {
	o.EmailNiew = emailNiew
}

// WithEmailNisw adds the emailNisw to the users users list params
func (o *UsersUsersListParams) WithEmailNisw(emailNisw *string) *UsersUsersListParams {
	o.SetEmailNisw(emailNisw)
	return o
}

// SetEmailNisw adds the emailNisw to the users users list params
func (o *UsersUsersListParams) SetEmailNisw(emailNisw *string) {
	o.EmailNisw = emailNisw
}

// WithFirstName adds the firstName to the users users list params
func (o *UsersUsersListParams) WithFirstName(firstName *string) *UsersUsersListParams {
	o.SetFirstName(firstName)
	return o
}

// SetFirstName adds the firstName to the users users list params
func (o *UsersUsersListParams) SetFirstName(firstName *string) {
	o.FirstName = firstName
}

// WithFirstNameIc adds the firstNameIc to the users users list params
func (o *UsersUsersListParams) WithFirstNameIc(firstNameIc *string) *UsersUsersListParams {
	o.SetFirstNameIc(firstNameIc)
	return o
}

// SetFirstNameIc adds the firstNameIc to the users users list params
func (o *UsersUsersListParams) SetFirstNameIc(firstNameIc *string) {
	o.FirstNameIc = firstNameIc
}

// WithFirstNameIe adds the firstNameIe to the users users list params
func (o *UsersUsersListParams) WithFirstNameIe(firstNameIe *string) *UsersUsersListParams {
	o.SetFirstNameIe(firstNameIe)
	return o
}

// SetFirstNameIe adds the firstNameIe to the users users list params
func (o *UsersUsersListParams) SetFirstNameIe(firstNameIe *string) {
	o.FirstNameIe = firstNameIe
}

// WithFirstNameIew adds the firstNameIew to the users users list params
func (o *UsersUsersListParams) WithFirstNameIew(firstNameIew *string) *UsersUsersListParams {
	o.SetFirstNameIew(firstNameIew)
	return o
}

// SetFirstNameIew adds the firstNameIew to the users users list params
func (o *UsersUsersListParams) SetFirstNameIew(firstNameIew *string) {
	o.FirstNameIew = firstNameIew
}

// WithFirstNameIsw adds the firstNameIsw to the users users list params
func (o *UsersUsersListParams) WithFirstNameIsw(firstNameIsw *string) *UsersUsersListParams {
	o.SetFirstNameIsw(firstNameIsw)
	return o
}

// SetFirstNameIsw adds the firstNameIsw to the users users list params
func (o *UsersUsersListParams) SetFirstNameIsw(firstNameIsw *string) {
	o.FirstNameIsw = firstNameIsw
}

// WithFirstNamen adds the firstNamen to the users users list params
func (o *UsersUsersListParams) WithFirstNamen(firstNamen *string) *UsersUsersListParams {
	o.SetFirstNamen(firstNamen)
	return o
}

// SetFirstNamen adds the firstNameN to the users users list params
func (o *UsersUsersListParams) SetFirstNamen(firstNamen *string) {
	o.FirstNamen = firstNamen
}

// WithFirstNameNic adds the firstNameNic to the users users list params
func (o *UsersUsersListParams) WithFirstNameNic(firstNameNic *string) *UsersUsersListParams {
	o.SetFirstNameNic(firstNameNic)
	return o
}

// SetFirstNameNic adds the firstNameNic to the users users list params
func (o *UsersUsersListParams) SetFirstNameNic(firstNameNic *string) {
	o.FirstNameNic = firstNameNic
}

// WithFirstNameNie adds the firstNameNie to the users users list params
func (o *UsersUsersListParams) WithFirstNameNie(firstNameNie *string) *UsersUsersListParams {
	o.SetFirstNameNie(firstNameNie)
	return o
}

// SetFirstNameNie adds the firstNameNie to the users users list params
func (o *UsersUsersListParams) SetFirstNameNie(firstNameNie *string) {
	o.FirstNameNie = firstNameNie
}

// WithFirstNameNiew adds the firstNameNiew to the users users list params
func (o *UsersUsersListParams) WithFirstNameNiew(firstNameNiew *string) *UsersUsersListParams {
	o.SetFirstNameNiew(firstNameNiew)
	return o
}

// SetFirstNameNiew adds the firstNameNiew to the users users list params
func (o *UsersUsersListParams) SetFirstNameNiew(firstNameNiew *string) {
	o.FirstNameNiew = firstNameNiew
}

// WithFirstNameNisw adds the firstNameNisw to the users users list params
func (o *UsersUsersListParams) WithFirstNameNisw(firstNameNisw *string) *UsersUsersListParams {
	o.SetFirstNameNisw(firstNameNisw)
	return o
}

// SetFirstNameNisw adds the firstNameNisw to the users users list params
func (o *UsersUsersListParams) SetFirstNameNisw(firstNameNisw *string) {
	o.FirstNameNisw = firstNameNisw
}

// WithGroup adds the group to the users users list params
func (o *UsersUsersListParams) WithGroup(group *string) *UsersUsersListParams {
	o.SetGroup(group)
	return o
}

// SetGroup adds the group to the users users list params
func (o *UsersUsersListParams) SetGroup(group *string) {
	o.Group = group
}

// WithGroupn adds the groupn to the users users list params
func (o *UsersUsersListParams) WithGroupn(groupn *string) *UsersUsersListParams {
	o.SetGroupn(groupn)
	return o
}

// SetGroupn adds the groupN to the users users list params
func (o *UsersUsersListParams) SetGroupn(groupn *string) {
	o.Groupn = groupn
}

// WithGroupID adds the groupID to the users users list params
func (o *UsersUsersListParams) WithGroupID(groupID *string) *UsersUsersListParams {
	o.SetGroupID(groupID)
	return o
}

// SetGroupID adds the groupId to the users users list params
func (o *UsersUsersListParams) SetGroupID(groupID *string) {
	o.GroupID = groupID
}

// WithGroupIDn adds the groupIDn to the users users list params
func (o *UsersUsersListParams) WithGroupIDn(groupIDn *string) *UsersUsersListParams {
	o.SetGroupIDn(groupIDn)
	return o
}

// SetGroupIDn adds the groupIdN to the users users list params
func (o *UsersUsersListParams) SetGroupIDn(groupIDn *string) {
	o.GroupIDn = groupIDn
}

// WithID adds the id to the users users list params
func (o *UsersUsersListParams) WithID(id *string) *UsersUsersListParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the users users list params
func (o *UsersUsersListParams) SetID(id *string) {
	o.ID = id
}

// WithIDIc adds the iDIc to the users users list params
func (o *UsersUsersListParams) WithIDIc(iDIc *string) *UsersUsersListParams {
	o.SetIDIc(iDIc)
	return o
}

// SetIDIc adds the idIc to the users users list params
func (o *UsersUsersListParams) SetIDIc(iDIc *string) {
	o.IDIc = iDIc
}

// WithIDIe adds the iDIe to the users users list params
func (o *UsersUsersListParams) WithIDIe(iDIe *string) *UsersUsersListParams {
	o.SetIDIe(iDIe)
	return o
}

// SetIDIe adds the idIe to the users users list params
func (o *UsersUsersListParams) SetIDIe(iDIe *string) {
	o.IDIe = iDIe
}

// WithIDIew adds the iDIew to the users users list params
func (o *UsersUsersListParams) WithIDIew(iDIew *string) *UsersUsersListParams {
	o.SetIDIew(iDIew)
	return o
}

// SetIDIew adds the idIew to the users users list params
func (o *UsersUsersListParams) SetIDIew(iDIew *string) {
	o.IDIew = iDIew
}

// WithIDIsw adds the iDIsw to the users users list params
func (o *UsersUsersListParams) WithIDIsw(iDIsw *string) *UsersUsersListParams {
	o.SetIDIsw(iDIsw)
	return o
}

// SetIDIsw adds the idIsw to the users users list params
func (o *UsersUsersListParams) SetIDIsw(iDIsw *string) {
	o.IDIsw = iDIsw
}

// WithIDn adds the iDn to the users users list params
func (o *UsersUsersListParams) WithIDn(iDn *string) *UsersUsersListParams {
	o.SetIDn(iDn)
	return o
}

// SetIDn adds the idN to the users users list params
func (o *UsersUsersListParams) SetIDn(iDn *string) {
	o.IDn = iDn
}

// WithIDNic adds the iDNic to the users users list params
func (o *UsersUsersListParams) WithIDNic(iDNic *string) *UsersUsersListParams {
	o.SetIDNic(iDNic)
	return o
}

// SetIDNic adds the idNic to the users users list params
func (o *UsersUsersListParams) SetIDNic(iDNic *string) {
	o.IDNic = iDNic
}

// WithIDNie adds the iDNie to the users users list params
func (o *UsersUsersListParams) WithIDNie(iDNie *string) *UsersUsersListParams {
	o.SetIDNie(iDNie)
	return o
}

// SetIDNie adds the idNie to the users users list params
func (o *UsersUsersListParams) SetIDNie(iDNie *string) {
	o.IDNie = iDNie
}

// WithIDNiew adds the iDNiew to the users users list params
func (o *UsersUsersListParams) WithIDNiew(iDNiew *string) *UsersUsersListParams {
	o.SetIDNiew(iDNiew)
	return o
}

// SetIDNiew adds the idNiew to the users users list params
func (o *UsersUsersListParams) SetIDNiew(iDNiew *string) {
	o.IDNiew = iDNiew
}

// WithIDNisw adds the iDNisw to the users users list params
func (o *UsersUsersListParams) WithIDNisw(iDNisw *string) *UsersUsersListParams {
	o.SetIDNisw(iDNisw)
	return o
}

// SetIDNisw adds the idNisw to the users users list params
func (o *UsersUsersListParams) SetIDNisw(iDNisw *string) {
	o.IDNisw = iDNisw
}

// WithIsActive adds the isActive to the users users list params
func (o *UsersUsersListParams) WithIsActive(isActive *string) *UsersUsersListParams {
	o.SetIsActive(isActive)
	return o
}

// SetIsActive adds the isActive to the users users list params
func (o *UsersUsersListParams) SetIsActive(isActive *string) {
	o.IsActive = isActive
}

// WithIsStaff adds the isStaff to the users users list params
func (o *UsersUsersListParams) WithIsStaff(isStaff *string) *UsersUsersListParams {
	o.SetIsStaff(isStaff)
	return o
}

// SetIsStaff adds the isStaff to the users users list params
func (o *UsersUsersListParams) SetIsStaff(isStaff *string) {
	o.IsStaff = isStaff
}

// WithLastName adds the lastName to the users users list params
func (o *UsersUsersListParams) WithLastName(lastName *string) *UsersUsersListParams {
	o.SetLastName(lastName)
	return o
}

// SetLastName adds the lastName to the users users list params
func (o *UsersUsersListParams) SetLastName(lastName *string) {
	o.LastName = lastName
}

// WithLastNameIc adds the lastNameIc to the users users list params
func (o *UsersUsersListParams) WithLastNameIc(lastNameIc *string) *UsersUsersListParams {
	o.SetLastNameIc(lastNameIc)
	return o
}

// SetLastNameIc adds the lastNameIc to the users users list params
func (o *UsersUsersListParams) SetLastNameIc(lastNameIc *string) {
	o.LastNameIc = lastNameIc
}

// WithLastNameIe adds the lastNameIe to the users users list params
func (o *UsersUsersListParams) WithLastNameIe(lastNameIe *string) *UsersUsersListParams {
	o.SetLastNameIe(lastNameIe)
	return o
}

// SetLastNameIe adds the lastNameIe to the users users list params
func (o *UsersUsersListParams) SetLastNameIe(lastNameIe *string) {
	o.LastNameIe = lastNameIe
}

// WithLastNameIew adds the lastNameIew to the users users list params
func (o *UsersUsersListParams) WithLastNameIew(lastNameIew *string) *UsersUsersListParams {
	o.SetLastNameIew(lastNameIew)
	return o
}

// SetLastNameIew adds the lastNameIew to the users users list params
func (o *UsersUsersListParams) SetLastNameIew(lastNameIew *string) {
	o.LastNameIew = lastNameIew
}

// WithLastNameIsw adds the lastNameIsw to the users users list params
func (o *UsersUsersListParams) WithLastNameIsw(lastNameIsw *string) *UsersUsersListParams {
	o.SetLastNameIsw(lastNameIsw)
	return o
}

// SetLastNameIsw adds the lastNameIsw to the users users list params
func (o *UsersUsersListParams) SetLastNameIsw(lastNameIsw *string) {
	o.LastNameIsw = lastNameIsw
}

// WithLastNamen adds the lastNamen to the users users list params
func (o *UsersUsersListParams) WithLastNamen(lastNamen *string) *UsersUsersListParams {
	o.SetLastNamen(lastNamen)
	return o
}

// SetLastNamen adds the lastNameN to the users users list params
func (o *UsersUsersListParams) SetLastNamen(lastNamen *string) {
	o.LastNamen = lastNamen
}

// WithLastNameNic adds the lastNameNic to the users users list params
func (o *UsersUsersListParams) WithLastNameNic(lastNameNic *string) *UsersUsersListParams {
	o.SetLastNameNic(lastNameNic)
	return o
}

// SetLastNameNic adds the lastNameNic to the users users list params
func (o *UsersUsersListParams) SetLastNameNic(lastNameNic *string) {
	o.LastNameNic = lastNameNic
}

// WithLastNameNie adds the lastNameNie to the users users list params
func (o *UsersUsersListParams) WithLastNameNie(lastNameNie *string) *UsersUsersListParams {
	o.SetLastNameNie(lastNameNie)
	return o
}

// SetLastNameNie adds the lastNameNie to the users users list params
func (o *UsersUsersListParams) SetLastNameNie(lastNameNie *string) {
	o.LastNameNie = lastNameNie
}

// WithLastNameNiew adds the lastNameNiew to the users users list params
func (o *UsersUsersListParams) WithLastNameNiew(lastNameNiew *string) *UsersUsersListParams {
	o.SetLastNameNiew(lastNameNiew)
	return o
}

// SetLastNameNiew adds the lastNameNiew to the users users list params
func (o *UsersUsersListParams) SetLastNameNiew(lastNameNiew *string) {
	o.LastNameNiew = lastNameNiew
}

// WithLastNameNisw adds the lastNameNisw to the users users list params
func (o *UsersUsersListParams) WithLastNameNisw(lastNameNisw *string) *UsersUsersListParams {
	o.SetLastNameNisw(lastNameNisw)
	return o
}

// SetLastNameNisw adds the lastNameNisw to the users users list params
func (o *UsersUsersListParams) SetLastNameNisw(lastNameNisw *string) {
	o.LastNameNisw = lastNameNisw
}

// WithLimit adds the limit to the users users list params
func (o *UsersUsersListParams) WithLimit(limit *int64) *UsersUsersListParams {
	o.SetLimit(limit)
	return o
}

// SetLimit adds the limit to the users users list params
func (o *UsersUsersListParams) SetLimit(limit *int64) {
	o.Limit = limit
}

// WithOffset adds the offset to the users users list params
func (o *UsersUsersListParams) WithOffset(offset *int64) *UsersUsersListParams {
	o.SetOffset(offset)
	return o
}

// SetOffset adds the offset to the users users list params
func (o *UsersUsersListParams) SetOffset(offset *int64) {
	o.Offset = offset
}

// WithQ adds the q to the users users list params
func (o *UsersUsersListParams) WithQ(q *string) *UsersUsersListParams {
	o.SetQ(q)
	return o
}

// SetQ adds the q to the users users list params
func (o *UsersUsersListParams) SetQ(q *string) {
	o.Q = q
}

// WithUsername adds the username to the users users list params
func (o *UsersUsersListParams) WithUsername(username *string) *UsersUsersListParams {
	o.SetUsername(username)
	return o
}

// SetUsername adds the username to the users users list params
func (o *UsersUsersListParams) SetUsername(username *string) {
	o.Username = username
}

// WithUsernameIc adds the usernameIc to the users users list params
func (o *UsersUsersListParams) WithUsernameIc(usernameIc *string) *UsersUsersListParams {
	o.SetUsernameIc(usernameIc)
	return o
}

// SetUsernameIc adds the usernameIc to the users users list params
func (o *UsersUsersListParams) SetUsernameIc(usernameIc *string) {
	o.UsernameIc = usernameIc
}

// WithUsernameIe adds the usernameIe to the users users list params
func (o *UsersUsersListParams) WithUsernameIe(usernameIe *string) *UsersUsersListParams {
	o.SetUsernameIe(usernameIe)
	return o
}

// SetUsernameIe adds the usernameIe to the users users list params
func (o *UsersUsersListParams) SetUsernameIe(usernameIe *string) {
	o.UsernameIe = usernameIe
}

// WithUsernameIew adds the usernameIew to the users users list params
func (o *UsersUsersListParams) WithUsernameIew(usernameIew *string) *UsersUsersListParams {
	o.SetUsernameIew(usernameIew)
	return o
}

// SetUsernameIew adds the usernameIew to the users users list params
func (o *UsersUsersListParams) SetUsernameIew(usernameIew *string) {
	o.UsernameIew = usernameIew
}

// WithUsernameIsw adds the usernameIsw to the users users list params
func (o *UsersUsersListParams) WithUsernameIsw(usernameIsw *string) *UsersUsersListParams {
	o.SetUsernameIsw(usernameIsw)
	return o
}

// SetUsernameIsw adds the usernameIsw to the users users list params
func (o *UsersUsersListParams) SetUsernameIsw(usernameIsw *string) {
	o.UsernameIsw = usernameIsw
}

// WithUsernamen adds the usernamen to the users users list params
func (o *UsersUsersListParams) WithUsernamen(usernamen *string) *UsersUsersListParams {
	o.SetUsernamen(usernamen)
	return o
}

// SetUsernamen adds the usernameN to the users users list params
func (o *UsersUsersListParams) SetUsernamen(usernamen *string) {
	o.Usernamen = usernamen
}

// WithUsernameNic adds the usernameNic to the users users list params
func (o *UsersUsersListParams) WithUsernameNic(usernameNic *string) *UsersUsersListParams {
	o.SetUsernameNic(usernameNic)
	return o
}

// SetUsernameNic adds the usernameNic to the users users list params
func (o *UsersUsersListParams) SetUsernameNic(usernameNic *string) {
	o.UsernameNic = usernameNic
}

// WithUsernameNie adds the usernameNie to the users users list params
func (o *UsersUsersListParams) WithUsernameNie(usernameNie *string) *UsersUsersListParams {
	o.SetUsernameNie(usernameNie)
	return o
}

// SetUsernameNie adds the usernameNie to the users users list params
func (o *UsersUsersListParams) SetUsernameNie(usernameNie *string) {
	o.UsernameNie = usernameNie
}

// WithUsernameNiew adds the usernameNiew to the users users list params
func (o *UsersUsersListParams) WithUsernameNiew(usernameNiew *string) *UsersUsersListParams {
	o.SetUsernameNiew(usernameNiew)
	return o
}

// SetUsernameNiew adds the usernameNiew to the users users list params
func (o *UsersUsersListParams) SetUsernameNiew(usernameNiew *string) {
	o.UsernameNiew = usernameNiew
}

// WithUsernameNisw adds the usernameNisw to the users users list params
func (o *UsersUsersListParams) WithUsernameNisw(usernameNisw *string) *UsersUsersListParams {
	o.SetUsernameNisw(usernameNisw)
	return o
}

// SetUsernameNisw adds the usernameNisw to the users users list params
func (o *UsersUsersListParams) SetUsernameNisw(usernameNisw *string) {
	o.UsernameNisw = usernameNisw
}

// WriteToRequest writes these params to a swagger request
func (o *UsersUsersListParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Email != nil {

		// query param email
		var qrEmail string

		if o.Email != nil {
			qrEmail = *o.Email
		}
		qEmail := qrEmail
		if qEmail != "" {

			if err := r.SetQueryParam("email", qEmail); err != nil {
				return err
			}
		}
	}

	if o.EmailIc != nil {

		// query param email__ic
		var qrEmailIc string

		if o.EmailIc != nil {
			qrEmailIc = *o.EmailIc
		}
		qEmailIc := qrEmailIc
		if qEmailIc != "" {

			if err := r.SetQueryParam("email__ic", qEmailIc); err != nil {
				return err
			}
		}
	}

	if o.EmailIe != nil {

		// query param email__ie
		var qrEmailIe string

		if o.EmailIe != nil {
			qrEmailIe = *o.EmailIe
		}
		qEmailIe := qrEmailIe
		if qEmailIe != "" {

			if err := r.SetQueryParam("email__ie", qEmailIe); err != nil {
				return err
			}
		}
	}

	if o.EmailIew != nil {

		// query param email__iew
		var qrEmailIew string

		if o.EmailIew != nil {
			qrEmailIew = *o.EmailIew
		}
		qEmailIew := qrEmailIew
		if qEmailIew != "" {

			if err := r.SetQueryParam("email__iew", qEmailIew); err != nil {
				return err
			}
		}
	}

	if o.EmailIsw != nil {

		// query param email__isw
		var qrEmailIsw string

		if o.EmailIsw != nil {
			qrEmailIsw = *o.EmailIsw
		}
		qEmailIsw := qrEmailIsw
		if qEmailIsw != "" {

			if err := r.SetQueryParam("email__isw", qEmailIsw); err != nil {
				return err
			}
		}
	}

	if o.Emailn != nil {

		// query param email__n
		var qrEmailn string

		if o.Emailn != nil {
			qrEmailn = *o.Emailn
		}
		qEmailn := qrEmailn
		if qEmailn != "" {

			if err := r.SetQueryParam("email__n", qEmailn); err != nil {
				return err
			}
		}
	}

	if o.EmailNic != nil {

		// query param email__nic
		var qrEmailNic string

		if o.EmailNic != nil {
			qrEmailNic = *o.EmailNic
		}
		qEmailNic := qrEmailNic
		if qEmailNic != "" {

			if err := r.SetQueryParam("email__nic", qEmailNic); err != nil {
				return err
			}
		}
	}

	if o.EmailNie != nil {

		// query param email__nie
		var qrEmailNie string

		if o.EmailNie != nil {
			qrEmailNie = *o.EmailNie
		}
		qEmailNie := qrEmailNie
		if qEmailNie != "" {

			if err := r.SetQueryParam("email__nie", qEmailNie); err != nil {
				return err
			}
		}
	}

	if o.EmailNiew != nil {

		// query param email__niew
		var qrEmailNiew string

		if o.EmailNiew != nil {
			qrEmailNiew = *o.EmailNiew
		}
		qEmailNiew := qrEmailNiew
		if qEmailNiew != "" {

			if err := r.SetQueryParam("email__niew", qEmailNiew); err != nil {
				return err
			}
		}
	}

	if o.EmailNisw != nil {

		// query param email__nisw
		var qrEmailNisw string

		if o.EmailNisw != nil {
			qrEmailNisw = *o.EmailNisw
		}
		qEmailNisw := qrEmailNisw
		if qEmailNisw != "" {

			if err := r.SetQueryParam("email__nisw", qEmailNisw); err != nil {
				return err
			}
		}
	}

	if o.FirstName != nil {

		// query param first_name
		var qrFirstName string

		if o.FirstName != nil {
			qrFirstName = *o.FirstName
		}
		qFirstName := qrFirstName
		if qFirstName != "" {

			if err := r.SetQueryParam("first_name", qFirstName); err != nil {
				return err
			}
		}
	}

	if o.FirstNameIc != nil {

		// query param first_name__ic
		var qrFirstNameIc string

		if o.FirstNameIc != nil {
			qrFirstNameIc = *o.FirstNameIc
		}
		qFirstNameIc := qrFirstNameIc
		if qFirstNameIc != "" {

			if err := r.SetQueryParam("first_name__ic", qFirstNameIc); err != nil {
				return err
			}
		}
	}

	if o.FirstNameIe != nil {

		// query param first_name__ie
		var qrFirstNameIe string

		if o.FirstNameIe != nil {
			qrFirstNameIe = *o.FirstNameIe
		}
		qFirstNameIe := qrFirstNameIe
		if qFirstNameIe != "" {

			if err := r.SetQueryParam("first_name__ie", qFirstNameIe); err != nil {
				return err
			}
		}
	}

	if o.FirstNameIew != nil {

		// query param first_name__iew
		var qrFirstNameIew string

		if o.FirstNameIew != nil {
			qrFirstNameIew = *o.FirstNameIew
		}
		qFirstNameIew := qrFirstNameIew
		if qFirstNameIew != "" {

			if err := r.SetQueryParam("first_name__iew", qFirstNameIew); err != nil {
				return err
			}
		}
	}

	if o.FirstNameIsw != nil {

		// query param first_name__isw
		var qrFirstNameIsw string

		if o.FirstNameIsw != nil {
			qrFirstNameIsw = *o.FirstNameIsw
		}
		qFirstNameIsw := qrFirstNameIsw
		if qFirstNameIsw != "" {

			if err := r.SetQueryParam("first_name__isw", qFirstNameIsw); err != nil {
				return err
			}
		}
	}

	if o.FirstNamen != nil {

		// query param first_name__n
		var qrFirstNamen string

		if o.FirstNamen != nil {
			qrFirstNamen = *o.FirstNamen
		}
		qFirstNamen := qrFirstNamen
		if qFirstNamen != "" {

			if err := r.SetQueryParam("first_name__n", qFirstNamen); err != nil {
				return err
			}
		}
	}

	if o.FirstNameNic != nil {

		// query param first_name__nic
		var qrFirstNameNic string

		if o.FirstNameNic != nil {
			qrFirstNameNic = *o.FirstNameNic
		}
		qFirstNameNic := qrFirstNameNic
		if qFirstNameNic != "" {

			if err := r.SetQueryParam("first_name__nic", qFirstNameNic); err != nil {
				return err
			}
		}
	}

	if o.FirstNameNie != nil {

		// query param first_name__nie
		var qrFirstNameNie string

		if o.FirstNameNie != nil {
			qrFirstNameNie = *o.FirstNameNie
		}
		qFirstNameNie := qrFirstNameNie
		if qFirstNameNie != "" {

			if err := r.SetQueryParam("first_name__nie", qFirstNameNie); err != nil {
				return err
			}
		}
	}

	if o.FirstNameNiew != nil {

		// query param first_name__niew
		var qrFirstNameNiew string

		if o.FirstNameNiew != nil {
			qrFirstNameNiew = *o.FirstNameNiew
		}
		qFirstNameNiew := qrFirstNameNiew
		if qFirstNameNiew != "" {

			if err := r.SetQueryParam("first_name__niew", qFirstNameNiew); err != nil {
				return err
			}
		}
	}

	if o.FirstNameNisw != nil {

		// query param first_name__nisw
		var qrFirstNameNisw string

		if o.FirstNameNisw != nil {
			qrFirstNameNisw = *o.FirstNameNisw
		}
		qFirstNameNisw := qrFirstNameNisw
		if qFirstNameNisw != "" {

			if err := r.SetQueryParam("first_name__nisw", qFirstNameNisw); err != nil {
				return err
			}
		}
	}

	if o.Group != nil {

		// query param group
		var qrGroup string

		if o.Group != nil {
			qrGroup = *o.Group
		}
		qGroup := qrGroup
		if qGroup != "" {

			if err := r.SetQueryParam("group", qGroup); err != nil {
				return err
			}
		}
	}

	if o.Groupn != nil {

		// query param group__n
		var qrGroupn string

		if o.Groupn != nil {
			qrGroupn = *o.Groupn
		}
		qGroupn := qrGroupn
		if qGroupn != "" {

			if err := r.SetQueryParam("group__n", qGroupn); err != nil {
				return err
			}
		}
	}

	if o.GroupID != nil {

		// query param group_id
		var qrGroupID string

		if o.GroupID != nil {
			qrGroupID = *o.GroupID
		}
		qGroupID := qrGroupID
		if qGroupID != "" {

			if err := r.SetQueryParam("group_id", qGroupID); err != nil {
				return err
			}
		}
	}

	if o.GroupIDn != nil {

		// query param group_id__n
		var qrGroupIDn string

		if o.GroupIDn != nil {
			qrGroupIDn = *o.GroupIDn
		}
		qGroupIDn := qrGroupIDn
		if qGroupIDn != "" {

			if err := r.SetQueryParam("group_id__n", qGroupIDn); err != nil {
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

	if o.IsActive != nil {

		// query param is_active
		var qrIsActive string

		if o.IsActive != nil {
			qrIsActive = *o.IsActive
		}
		qIsActive := qrIsActive
		if qIsActive != "" {

			if err := r.SetQueryParam("is_active", qIsActive); err != nil {
				return err
			}
		}
	}

	if o.IsStaff != nil {

		// query param is_staff
		var qrIsStaff string

		if o.IsStaff != nil {
			qrIsStaff = *o.IsStaff
		}
		qIsStaff := qrIsStaff
		if qIsStaff != "" {

			if err := r.SetQueryParam("is_staff", qIsStaff); err != nil {
				return err
			}
		}
	}

	if o.LastName != nil {

		// query param last_name
		var qrLastName string

		if o.LastName != nil {
			qrLastName = *o.LastName
		}
		qLastName := qrLastName
		if qLastName != "" {

			if err := r.SetQueryParam("last_name", qLastName); err != nil {
				return err
			}
		}
	}

	if o.LastNameIc != nil {

		// query param last_name__ic
		var qrLastNameIc string

		if o.LastNameIc != nil {
			qrLastNameIc = *o.LastNameIc
		}
		qLastNameIc := qrLastNameIc
		if qLastNameIc != "" {

			if err := r.SetQueryParam("last_name__ic", qLastNameIc); err != nil {
				return err
			}
		}
	}

	if o.LastNameIe != nil {

		// query param last_name__ie
		var qrLastNameIe string

		if o.LastNameIe != nil {
			qrLastNameIe = *o.LastNameIe
		}
		qLastNameIe := qrLastNameIe
		if qLastNameIe != "" {

			if err := r.SetQueryParam("last_name__ie", qLastNameIe); err != nil {
				return err
			}
		}
	}

	if o.LastNameIew != nil {

		// query param last_name__iew
		var qrLastNameIew string

		if o.LastNameIew != nil {
			qrLastNameIew = *o.LastNameIew
		}
		qLastNameIew := qrLastNameIew
		if qLastNameIew != "" {

			if err := r.SetQueryParam("last_name__iew", qLastNameIew); err != nil {
				return err
			}
		}
	}

	if o.LastNameIsw != nil {

		// query param last_name__isw
		var qrLastNameIsw string

		if o.LastNameIsw != nil {
			qrLastNameIsw = *o.LastNameIsw
		}
		qLastNameIsw := qrLastNameIsw
		if qLastNameIsw != "" {

			if err := r.SetQueryParam("last_name__isw", qLastNameIsw); err != nil {
				return err
			}
		}
	}

	if o.LastNamen != nil {

		// query param last_name__n
		var qrLastNamen string

		if o.LastNamen != nil {
			qrLastNamen = *o.LastNamen
		}
		qLastNamen := qrLastNamen
		if qLastNamen != "" {

			if err := r.SetQueryParam("last_name__n", qLastNamen); err != nil {
				return err
			}
		}
	}

	if o.LastNameNic != nil {

		// query param last_name__nic
		var qrLastNameNic string

		if o.LastNameNic != nil {
			qrLastNameNic = *o.LastNameNic
		}
		qLastNameNic := qrLastNameNic
		if qLastNameNic != "" {

			if err := r.SetQueryParam("last_name__nic", qLastNameNic); err != nil {
				return err
			}
		}
	}

	if o.LastNameNie != nil {

		// query param last_name__nie
		var qrLastNameNie string

		if o.LastNameNie != nil {
			qrLastNameNie = *o.LastNameNie
		}
		qLastNameNie := qrLastNameNie
		if qLastNameNie != "" {

			if err := r.SetQueryParam("last_name__nie", qLastNameNie); err != nil {
				return err
			}
		}
	}

	if o.LastNameNiew != nil {

		// query param last_name__niew
		var qrLastNameNiew string

		if o.LastNameNiew != nil {
			qrLastNameNiew = *o.LastNameNiew
		}
		qLastNameNiew := qrLastNameNiew
		if qLastNameNiew != "" {

			if err := r.SetQueryParam("last_name__niew", qLastNameNiew); err != nil {
				return err
			}
		}
	}

	if o.LastNameNisw != nil {

		// query param last_name__nisw
		var qrLastNameNisw string

		if o.LastNameNisw != nil {
			qrLastNameNisw = *o.LastNameNisw
		}
		qLastNameNisw := qrLastNameNisw
		if qLastNameNisw != "" {

			if err := r.SetQueryParam("last_name__nisw", qLastNameNisw); err != nil {
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

	if o.Username != nil {

		// query param username
		var qrUsername string

		if o.Username != nil {
			qrUsername = *o.Username
		}
		qUsername := qrUsername
		if qUsername != "" {

			if err := r.SetQueryParam("username", qUsername); err != nil {
				return err
			}
		}
	}

	if o.UsernameIc != nil {

		// query param username__ic
		var qrUsernameIc string

		if o.UsernameIc != nil {
			qrUsernameIc = *o.UsernameIc
		}
		qUsernameIc := qrUsernameIc
		if qUsernameIc != "" {

			if err := r.SetQueryParam("username__ic", qUsernameIc); err != nil {
				return err
			}
		}
	}

	if o.UsernameIe != nil {

		// query param username__ie
		var qrUsernameIe string

		if o.UsernameIe != nil {
			qrUsernameIe = *o.UsernameIe
		}
		qUsernameIe := qrUsernameIe
		if qUsernameIe != "" {

			if err := r.SetQueryParam("username__ie", qUsernameIe); err != nil {
				return err
			}
		}
	}

	if o.UsernameIew != nil {

		// query param username__iew
		var qrUsernameIew string

		if o.UsernameIew != nil {
			qrUsernameIew = *o.UsernameIew
		}
		qUsernameIew := qrUsernameIew
		if qUsernameIew != "" {

			if err := r.SetQueryParam("username__iew", qUsernameIew); err != nil {
				return err
			}
		}
	}

	if o.UsernameIsw != nil {

		// query param username__isw
		var qrUsernameIsw string

		if o.UsernameIsw != nil {
			qrUsernameIsw = *o.UsernameIsw
		}
		qUsernameIsw := qrUsernameIsw
		if qUsernameIsw != "" {

			if err := r.SetQueryParam("username__isw", qUsernameIsw); err != nil {
				return err
			}
		}
	}

	if o.Usernamen != nil {

		// query param username__n
		var qrUsernamen string

		if o.Usernamen != nil {
			qrUsernamen = *o.Usernamen
		}
		qUsernamen := qrUsernamen
		if qUsernamen != "" {

			if err := r.SetQueryParam("username__n", qUsernamen); err != nil {
				return err
			}
		}
	}

	if o.UsernameNic != nil {

		// query param username__nic
		var qrUsernameNic string

		if o.UsernameNic != nil {
			qrUsernameNic = *o.UsernameNic
		}
		qUsernameNic := qrUsernameNic
		if qUsernameNic != "" {

			if err := r.SetQueryParam("username__nic", qUsernameNic); err != nil {
				return err
			}
		}
	}

	if o.UsernameNie != nil {

		// query param username__nie
		var qrUsernameNie string

		if o.UsernameNie != nil {
			qrUsernameNie = *o.UsernameNie
		}
		qUsernameNie := qrUsernameNie
		if qUsernameNie != "" {

			if err := r.SetQueryParam("username__nie", qUsernameNie); err != nil {
				return err
			}
		}
	}

	if o.UsernameNiew != nil {

		// query param username__niew
		var qrUsernameNiew string

		if o.UsernameNiew != nil {
			qrUsernameNiew = *o.UsernameNiew
		}
		qUsernameNiew := qrUsernameNiew
		if qUsernameNiew != "" {

			if err := r.SetQueryParam("username__niew", qUsernameNiew); err != nil {
				return err
			}
		}
	}

	if o.UsernameNisw != nil {

		// query param username__nisw
		var qrUsernameNisw string

		if o.UsernameNisw != nil {
			qrUsernameNisw = *o.UsernameNisw
		}
		qUsernameNisw := qrUsernameNisw
		if qUsernameNisw != "" {

			if err := r.SetQueryParam("username__nisw", qUsernameNisw); err != nil {
				return err
			}
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
