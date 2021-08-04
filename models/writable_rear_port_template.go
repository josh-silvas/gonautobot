package models

import (
	"context"
	"encoding/json"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// WritableRearPortTemplate writable rear port template
//
// swagger:model WritableRearPortTemplate
type WritableRearPortTemplate struct {

	// Custom fields
	CustomFields interface{} `json:"custom_fields,omitempty"`

	// Description
	// Max Length: 200
	Description string `json:"description,omitempty"`

	// Device type
	// Required: true
	// Format: uuid
	DeviceType *strfmt.UUID `json:"device_type"`

	// Display
	//
	// Human friendly display value
	// Read Only: true
	// Min Length: 1
	Display string `json:"display,omitempty"`

	// Id
	// Read Only: true
	// Format: uuid
	ID strfmt.UUID `json:"id,omitempty"`

	// Label
	//
	// Physical label
	// Max Length: 64
	Label string `json:"label,omitempty"`

	// Name
	// Required: true
	// Max Length: 64
	// Min Length: 1
	Name *string `json:"name"`

	// Positions
	// Maximum: 1024
	// Minimum: 1
	Positions int64 `json:"positions,omitempty"`

	// Type
	// Required: true
	// Enum: [8p8c 8p6c 8p4c 8p2c gg45 tera-4p tera-2p tera-1p 110-punch bnc mrj21 fc lc lc-apc lsh lsh-apc mpo mtrj sc sc-apc st cs sn splice]
	Type *string `json:"type"`

	// Url
	// Read Only: true
	// Format: uri
	URL strfmt.URI `json:"url,omitempty"`
}

// Validate validates this writable rear port template
func (m *WritableRearPortTemplate) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateDescription(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDeviceType(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDisplay(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateLabel(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateName(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePositions(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateType(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateURL(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *WritableRearPortTemplate) validateDescription(formats strfmt.Registry) error {
	if swag.IsZero(m.Description) { // not required
		return nil
	}

	if err := validate.MaxLength("description", "body", m.Description, 200); err != nil {
		return err
	}

	return nil
}

func (m *WritableRearPortTemplate) validateDeviceType(formats strfmt.Registry) error {

	if err := validate.Required("device_type", "body", m.DeviceType); err != nil {
		return err
	}

	if err := validate.FormatOf("device_type", "body", "uuid", m.DeviceType.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *WritableRearPortTemplate) validateDisplay(formats strfmt.Registry) error {
	if swag.IsZero(m.Display) { // not required
		return nil
	}

	if err := validate.MinLength("display", "body", m.Display, 1); err != nil {
		return err
	}

	return nil
}

func (m *WritableRearPortTemplate) validateID(formats strfmt.Registry) error {
	if swag.IsZero(m.ID) { // not required
		return nil
	}

	if err := validate.FormatOf("id", "body", "uuid", m.ID.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *WritableRearPortTemplate) validateLabel(formats strfmt.Registry) error {
	if swag.IsZero(m.Label) { // not required
		return nil
	}

	if err := validate.MaxLength("label", "body", m.Label, 64); err != nil {
		return err
	}

	return nil
}

func (m *WritableRearPortTemplate) validateName(formats strfmt.Registry) error {

	if err := validate.Required("name", "body", m.Name); err != nil {
		return err
	}

	if err := validate.MinLength("name", "body", *m.Name, 1); err != nil {
		return err
	}

	if err := validate.MaxLength("name", "body", *m.Name, 64); err != nil {
		return err
	}

	return nil
}

func (m *WritableRearPortTemplate) validatePositions(formats strfmt.Registry) error {
	if swag.IsZero(m.Positions) { // not required
		return nil
	}

	if err := validate.MinimumInt("positions", "body", m.Positions, 1, false); err != nil {
		return err
	}

	if err := validate.MaximumInt("positions", "body", m.Positions, 1024, false); err != nil {
		return err
	}

	return nil
}

var writableRearPortTemplateTypeTypePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["8p8c","8p6c","8p4c","8p2c","gg45","tera-4p","tera-2p","tera-1p","110-punch","bnc","mrj21","fc","lc","lc-apc","lsh","lsh-apc","mpo","mtrj","sc","sc-apc","st","cs","sn","splice"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		writableRearPortTemplateTypeTypePropEnum = append(writableRearPortTemplateTypeTypePropEnum, v)
	}
}

const (

	// WritableRearPortTemplateTypeNr8p8c captures enum value "8p8c"
	WritableRearPortTemplateTypeNr8p8c string = "8p8c"

	// WritableRearPortTemplateTypeNr8p6c captures enum value "8p6c"
	WritableRearPortTemplateTypeNr8p6c string = "8p6c"

	// WritableRearPortTemplateTypeNr8p4c captures enum value "8p4c"
	WritableRearPortTemplateTypeNr8p4c string = "8p4c"

	// WritableRearPortTemplateTypeNr8p2c captures enum value "8p2c"
	WritableRearPortTemplateTypeNr8p2c string = "8p2c"

	// WritableRearPortTemplateTypeGg45 captures enum value "gg45"
	WritableRearPortTemplateTypeGg45 string = "gg45"

	// WritableRearPortTemplateTypeTeraDash4p captures enum value "tera-4p"
	WritableRearPortTemplateTypeTeraDash4p string = "tera-4p"

	// WritableRearPortTemplateTypeTeraDash2p captures enum value "tera-2p"
	WritableRearPortTemplateTypeTeraDash2p string = "tera-2p"

	// WritableRearPortTemplateTypeTeraDash1p captures enum value "tera-1p"
	WritableRearPortTemplateTypeTeraDash1p string = "tera-1p"

	// WritableRearPortTemplateTypeNr110DashPunch captures enum value "110-punch"
	WritableRearPortTemplateTypeNr110DashPunch string = "110-punch"

	// WritableRearPortTemplateTypeBnc captures enum value "bnc"
	WritableRearPortTemplateTypeBnc string = "bnc"

	// WritableRearPortTemplateTypeMrj21 captures enum value "mrj21"
	WritableRearPortTemplateTypeMrj21 string = "mrj21"

	// WritableRearPortTemplateTypeFc captures enum value "fc"
	WritableRearPortTemplateTypeFc string = "fc"

	// WritableRearPortTemplateTypeLc captures enum value "lc"
	WritableRearPortTemplateTypeLc string = "lc"

	// WritableRearPortTemplateTypeLcDashApc captures enum value "lc-apc"
	WritableRearPortTemplateTypeLcDashApc string = "lc-apc"

	// WritableRearPortTemplateTypeLsh captures enum value "lsh"
	WritableRearPortTemplateTypeLsh string = "lsh"

	// WritableRearPortTemplateTypeLshDashApc captures enum value "lsh-apc"
	WritableRearPortTemplateTypeLshDashApc string = "lsh-apc"

	// WritableRearPortTemplateTypeMpo captures enum value "mpo"
	WritableRearPortTemplateTypeMpo string = "mpo"

	// WritableRearPortTemplateTypeMtrj captures enum value "mtrj"
	WritableRearPortTemplateTypeMtrj string = "mtrj"

	// WritableRearPortTemplateTypeSc captures enum value "sc"
	WritableRearPortTemplateTypeSc string = "sc"

	// WritableRearPortTemplateTypeScDashApc captures enum value "sc-apc"
	WritableRearPortTemplateTypeScDashApc string = "sc-apc"

	// WritableRearPortTemplateTypeSt captures enum value "st"
	WritableRearPortTemplateTypeSt string = "st"

	// WritableRearPortTemplateTypeCs captures enum value "cs"
	WritableRearPortTemplateTypeCs string = "cs"

	// WritableRearPortTemplateTypeSn captures enum value "sn"
	WritableRearPortTemplateTypeSn string = "sn"

	// WritableRearPortTemplateTypeSplice captures enum value "splice"
	WritableRearPortTemplateTypeSplice string = "splice"
)

// prop value enum
func (m *WritableRearPortTemplate) validateTypeEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, writableRearPortTemplateTypeTypePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *WritableRearPortTemplate) validateType(formats strfmt.Registry) error {

	if err := validate.Required("type", "body", m.Type); err != nil {
		return err
	}

	// value enum
	if err := m.validateTypeEnum("type", "body", *m.Type); err != nil {
		return err
	}

	return nil
}

func (m *WritableRearPortTemplate) validateURL(formats strfmt.Registry) error {
	if swag.IsZero(m.URL) { // not required
		return nil
	}

	if err := validate.FormatOf("url", "body", "uri", m.URL.String(), formats); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this writable rear port template based on the context it is used
func (m *WritableRearPortTemplate) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateDisplay(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateID(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateURL(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *WritableRearPortTemplate) contextValidateDisplay(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "display", "body", string(m.Display)); err != nil {
		return err
	}

	return nil
}

func (m *WritableRearPortTemplate) contextValidateID(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "id", "body", strfmt.UUID(m.ID)); err != nil {
		return err
	}

	return nil
}

func (m *WritableRearPortTemplate) contextValidateURL(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "url", "body", strfmt.URI(m.URL)); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *WritableRearPortTemplate) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *WritableRearPortTemplate) UnmarshalBinary(b []byte) error {
	var res WritableRearPortTemplate
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
