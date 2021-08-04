package models

import (
	"context"
	"encoding/json"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// RearPortTemplate rear port template
//
// swagger:model RearPortTemplate
type RearPortTemplate struct {

	// Custom fields
	CustomFields interface{} `json:"custom_fields,omitempty"`

	// Description
	// Max Length: 200
	Description string `json:"description,omitempty"`

	// device type
	// Required: true
	DeviceType *NestedDeviceType `json:"device_type"`

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

	// type
	// Required: true
	Type *RearPortTemplateType `json:"type"`

	// Url
	// Read Only: true
	// Format: uri
	URL strfmt.URI `json:"url,omitempty"`
}

// Validate validates this rear port template
func (m *RearPortTemplate) Validate(formats strfmt.Registry) error {
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

func (m *RearPortTemplate) validateDescription(formats strfmt.Registry) error {
	if swag.IsZero(m.Description) { // not required
		return nil
	}

	if err := validate.MaxLength("description", "body", m.Description, 200); err != nil {
		return err
	}

	return nil
}

func (m *RearPortTemplate) validateDeviceType(formats strfmt.Registry) error {

	if err := validate.Required("device_type", "body", m.DeviceType); err != nil {
		return err
	}

	if m.DeviceType != nil {
		if err := m.DeviceType.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("device_type")
			}
			return err
		}
	}

	return nil
}

func (m *RearPortTemplate) validateDisplay(formats strfmt.Registry) error {
	if swag.IsZero(m.Display) { // not required
		return nil
	}

	if err := validate.MinLength("display", "body", m.Display, 1); err != nil {
		return err
	}

	return nil
}

func (m *RearPortTemplate) validateID(formats strfmt.Registry) error {
	if swag.IsZero(m.ID) { // not required
		return nil
	}

	if err := validate.FormatOf("id", "body", "uuid", m.ID.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *RearPortTemplate) validateLabel(formats strfmt.Registry) error {
	if swag.IsZero(m.Label) { // not required
		return nil
	}

	if err := validate.MaxLength("label", "body", m.Label, 64); err != nil {
		return err
	}

	return nil
}

func (m *RearPortTemplate) validateName(formats strfmt.Registry) error {

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

func (m *RearPortTemplate) validatePositions(formats strfmt.Registry) error {
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

func (m *RearPortTemplate) validateType(formats strfmt.Registry) error {

	if err := validate.Required("type", "body", m.Type); err != nil {
		return err
	}

	if m.Type != nil {
		if err := m.Type.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("type")
			}
			return err
		}
	}

	return nil
}

func (m *RearPortTemplate) validateURL(formats strfmt.Registry) error {
	if swag.IsZero(m.URL) { // not required
		return nil
	}

	if err := validate.FormatOf("url", "body", "uri", m.URL.String(), formats); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this rear port template based on the context it is used
func (m *RearPortTemplate) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateDeviceType(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateDisplay(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateID(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateType(ctx, formats); err != nil {
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

func (m *RearPortTemplate) contextValidateDeviceType(ctx context.Context, formats strfmt.Registry) error {

	if m.DeviceType != nil {
		if err := m.DeviceType.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("device_type")
			}
			return err
		}
	}

	return nil
}

func (m *RearPortTemplate) contextValidateDisplay(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "display", "body", string(m.Display)); err != nil {
		return err
	}

	return nil
}

func (m *RearPortTemplate) contextValidateID(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "id", "body", strfmt.UUID(m.ID)); err != nil {
		return err
	}

	return nil
}

func (m *RearPortTemplate) contextValidateType(ctx context.Context, formats strfmt.Registry) error {

	if m.Type != nil {
		if err := m.Type.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("type")
			}
			return err
		}
	}

	return nil
}

func (m *RearPortTemplate) contextValidateURL(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "url", "body", strfmt.URI(m.URL)); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *RearPortTemplate) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *RearPortTemplate) UnmarshalBinary(b []byte) error {
	var res RearPortTemplate
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// RearPortTemplateType Type
//
// swagger:model RearPortTemplateType
type RearPortTemplateType struct {

	// label
	// Required: true
	// Enum: [8P8C 8P6C 8P4C 8P2C GG45 TERA 4P TERA 2P TERA 1P 110 Punch BNC MRJ21 FC LC LC/APC LSH LSH/APC MPO MTRJ SC SC/APC ST CS SN Splice]
	Label *string `json:"label"`

	// value
	// Required: true
	// Enum: [8p8c 8p6c 8p4c 8p2c gg45 tera-4p tera-2p tera-1p 110-punch bnc mrj21 fc lc lc-apc lsh lsh-apc mpo mtrj sc sc-apc st cs sn splice]
	Value *string `json:"value"`
}

// Validate validates this rear port template type
func (m *RearPortTemplateType) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateLabel(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateValue(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

var rearPortTemplateTypeTypeLabelPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["8P8C","8P6C","8P4C","8P2C","GG45","TERA 4P","TERA 2P","TERA 1P","110 Punch","BNC","MRJ21","FC","LC","LC/APC","LSH","LSH/APC","MPO","MTRJ","SC","SC/APC","ST","CS","SN","Splice"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		rearPortTemplateTypeTypeLabelPropEnum = append(rearPortTemplateTypeTypeLabelPropEnum, v)
	}
}

const (

	// RearPortTemplateTypeLabelNr8P8C captures enum value "8P8C"
	RearPortTemplateTypeLabelNr8P8C string = "8P8C"

	// RearPortTemplateTypeLabelNr8P6C captures enum value "8P6C"
	RearPortTemplateTypeLabelNr8P6C string = "8P6C"

	// RearPortTemplateTypeLabelNr8P4C captures enum value "8P4C"
	RearPortTemplateTypeLabelNr8P4C string = "8P4C"

	// RearPortTemplateTypeLabelNr8P2C captures enum value "8P2C"
	RearPortTemplateTypeLabelNr8P2C string = "8P2C"

	// RearPortTemplateTypeLabelGG45 captures enum value "GG45"
	RearPortTemplateTypeLabelGG45 string = "GG45"

	// RearPortTemplateTypeLabelTERA4P captures enum value "TERA 4P"
	RearPortTemplateTypeLabelTERA4P string = "TERA 4P"

	// RearPortTemplateTypeLabelTERA2P captures enum value "TERA 2P"
	RearPortTemplateTypeLabelTERA2P string = "TERA 2P"

	// RearPortTemplateTypeLabelTERA1P captures enum value "TERA 1P"
	RearPortTemplateTypeLabelTERA1P string = "TERA 1P"

	// RearPortTemplateTypeLabelNr110Punch captures enum value "110 Punch"
	RearPortTemplateTypeLabelNr110Punch string = "110 Punch"

	// RearPortTemplateTypeLabelBNC captures enum value "BNC"
	RearPortTemplateTypeLabelBNC string = "BNC"

	// RearPortTemplateTypeLabelMRJ21 captures enum value "MRJ21"
	RearPortTemplateTypeLabelMRJ21 string = "MRJ21"

	// RearPortTemplateTypeLabelFC captures enum value "FC"
	RearPortTemplateTypeLabelFC string = "FC"

	// RearPortTemplateTypeLabelLC captures enum value "LC"
	RearPortTemplateTypeLabelLC string = "LC"

	// RearPortTemplateTypeLabelLCAPC captures enum value "LC/APC"
	RearPortTemplateTypeLabelLCAPC string = "LC/APC"

	// RearPortTemplateTypeLabelLSH captures enum value "LSH"
	RearPortTemplateTypeLabelLSH string = "LSH"

	// RearPortTemplateTypeLabelLSHAPC captures enum value "LSH/APC"
	RearPortTemplateTypeLabelLSHAPC string = "LSH/APC"

	// RearPortTemplateTypeLabelMPO captures enum value "MPO"
	RearPortTemplateTypeLabelMPO string = "MPO"

	// RearPortTemplateTypeLabelMTRJ captures enum value "MTRJ"
	RearPortTemplateTypeLabelMTRJ string = "MTRJ"

	// RearPortTemplateTypeLabelSC captures enum value "SC"
	RearPortTemplateTypeLabelSC string = "SC"

	// RearPortTemplateTypeLabelSCAPC captures enum value "SC/APC"
	RearPortTemplateTypeLabelSCAPC string = "SC/APC"

	// RearPortTemplateTypeLabelST captures enum value "ST"
	RearPortTemplateTypeLabelST string = "ST"

	// RearPortTemplateTypeLabelCS captures enum value "CS"
	RearPortTemplateTypeLabelCS string = "CS"

	// RearPortTemplateTypeLabelSN captures enum value "SN"
	RearPortTemplateTypeLabelSN string = "SN"

	// RearPortTemplateTypeLabelSplice captures enum value "Splice"
	RearPortTemplateTypeLabelSplice string = "Splice"
)

// prop value enum
func (m *RearPortTemplateType) validateLabelEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, rearPortTemplateTypeTypeLabelPropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *RearPortTemplateType) validateLabel(formats strfmt.Registry) error {

	if err := validate.Required("type"+"."+"label", "body", m.Label); err != nil {
		return err
	}

	// value enum
	if err := m.validateLabelEnum("type"+"."+"label", "body", *m.Label); err != nil {
		return err
	}

	return nil
}

var rearPortTemplateTypeTypeValuePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["8p8c","8p6c","8p4c","8p2c","gg45","tera-4p","tera-2p","tera-1p","110-punch","bnc","mrj21","fc","lc","lc-apc","lsh","lsh-apc","mpo","mtrj","sc","sc-apc","st","cs","sn","splice"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		rearPortTemplateTypeTypeValuePropEnum = append(rearPortTemplateTypeTypeValuePropEnum, v)
	}
}

const (

	// RearPortTemplateTypeValueNr8p8c captures enum value "8p8c"
	RearPortTemplateTypeValueNr8p8c string = "8p8c"

	// RearPortTemplateTypeValueNr8p6c captures enum value "8p6c"
	RearPortTemplateTypeValueNr8p6c string = "8p6c"

	// RearPortTemplateTypeValueNr8p4c captures enum value "8p4c"
	RearPortTemplateTypeValueNr8p4c string = "8p4c"

	// RearPortTemplateTypeValueNr8p2c captures enum value "8p2c"
	RearPortTemplateTypeValueNr8p2c string = "8p2c"

	// RearPortTemplateTypeValueGg45 captures enum value "gg45"
	RearPortTemplateTypeValueGg45 string = "gg45"

	// RearPortTemplateTypeValueTeraDash4p captures enum value "tera-4p"
	RearPortTemplateTypeValueTeraDash4p string = "tera-4p"

	// RearPortTemplateTypeValueTeraDash2p captures enum value "tera-2p"
	RearPortTemplateTypeValueTeraDash2p string = "tera-2p"

	// RearPortTemplateTypeValueTeraDash1p captures enum value "tera-1p"
	RearPortTemplateTypeValueTeraDash1p string = "tera-1p"

	// RearPortTemplateTypeValueNr110DashPunch captures enum value "110-punch"
	RearPortTemplateTypeValueNr110DashPunch string = "110-punch"

	// RearPortTemplateTypeValueBnc captures enum value "bnc"
	RearPortTemplateTypeValueBnc string = "bnc"

	// RearPortTemplateTypeValueMrj21 captures enum value "mrj21"
	RearPortTemplateTypeValueMrj21 string = "mrj21"

	// RearPortTemplateTypeValueFc captures enum value "fc"
	RearPortTemplateTypeValueFc string = "fc"

	// RearPortTemplateTypeValueLc captures enum value "lc"
	RearPortTemplateTypeValueLc string = "lc"

	// RearPortTemplateTypeValueLcDashApc captures enum value "lc-apc"
	RearPortTemplateTypeValueLcDashApc string = "lc-apc"

	// RearPortTemplateTypeValueLsh captures enum value "lsh"
	RearPortTemplateTypeValueLsh string = "lsh"

	// RearPortTemplateTypeValueLshDashApc captures enum value "lsh-apc"
	RearPortTemplateTypeValueLshDashApc string = "lsh-apc"

	// RearPortTemplateTypeValueMpo captures enum value "mpo"
	RearPortTemplateTypeValueMpo string = "mpo"

	// RearPortTemplateTypeValueMtrj captures enum value "mtrj"
	RearPortTemplateTypeValueMtrj string = "mtrj"

	// RearPortTemplateTypeValueSc captures enum value "sc"
	RearPortTemplateTypeValueSc string = "sc"

	// RearPortTemplateTypeValueScDashApc captures enum value "sc-apc"
	RearPortTemplateTypeValueScDashApc string = "sc-apc"

	// RearPortTemplateTypeValueSt captures enum value "st"
	RearPortTemplateTypeValueSt string = "st"

	// RearPortTemplateTypeValueCs captures enum value "cs"
	RearPortTemplateTypeValueCs string = "cs"

	// RearPortTemplateTypeValueSn captures enum value "sn"
	RearPortTemplateTypeValueSn string = "sn"

	// RearPortTemplateTypeValueSplice captures enum value "splice"
	RearPortTemplateTypeValueSplice string = "splice"
)

// prop value enum
func (m *RearPortTemplateType) validateValueEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, rearPortTemplateTypeTypeValuePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *RearPortTemplateType) validateValue(formats strfmt.Registry) error {

	if err := validate.Required("type"+"."+"value", "body", m.Value); err != nil {
		return err
	}

	// value enum
	if err := m.validateValueEnum("type"+"."+"value", "body", *m.Value); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this rear port template type based on context it is used
func (m *RearPortTemplateType) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *RearPortTemplateType) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *RearPortTemplateType) UnmarshalBinary(b []byte) error {
	var res RearPortTemplateType
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
