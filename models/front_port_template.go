package models

import (
	"context"
	"encoding/json"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// FrontPortTemplate front port template
//
// swagger:model FrontPortTemplate
type FrontPortTemplate struct {

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

	// rear port
	// Required: true
	RearPort *NestedRearPortTemplate `json:"rear_port"`

	// Rear port position
	// Maximum: 1024
	// Minimum: 1
	RearPortPosition int64 `json:"rear_port_position,omitempty"`

	// type
	// Required: true
	Type *FrontPortTemplateType `json:"type"`

	// Url
	// Read Only: true
	// Format: uri
	URL strfmt.URI `json:"url,omitempty"`
}

// Validate validates this front port template
func (m *FrontPortTemplate) Validate(formats strfmt.Registry) error {
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

	if err := m.validateRearPort(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateRearPortPosition(formats); err != nil {
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

func (m *FrontPortTemplate) validateDescription(formats strfmt.Registry) error {
	if swag.IsZero(m.Description) { // not required
		return nil
	}

	if err := validate.MaxLength("description", "body", m.Description, 200); err != nil {
		return err
	}

	return nil
}

func (m *FrontPortTemplate) validateDeviceType(formats strfmt.Registry) error {

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

func (m *FrontPortTemplate) validateDisplay(formats strfmt.Registry) error {
	if swag.IsZero(m.Display) { // not required
		return nil
	}

	if err := validate.MinLength("display", "body", m.Display, 1); err != nil {
		return err
	}

	return nil
}

func (m *FrontPortTemplate) validateID(formats strfmt.Registry) error {
	if swag.IsZero(m.ID) { // not required
		return nil
	}

	if err := validate.FormatOf("id", "body", "uuid", m.ID.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *FrontPortTemplate) validateLabel(formats strfmt.Registry) error {
	if swag.IsZero(m.Label) { // not required
		return nil
	}

	if err := validate.MaxLength("label", "body", m.Label, 64); err != nil {
		return err
	}

	return nil
}

func (m *FrontPortTemplate) validateName(formats strfmt.Registry) error {

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

func (m *FrontPortTemplate) validateRearPort(formats strfmt.Registry) error {

	if err := validate.Required("rear_port", "body", m.RearPort); err != nil {
		return err
	}

	if m.RearPort != nil {
		if err := m.RearPort.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("rear_port")
			}
			return err
		}
	}

	return nil
}

func (m *FrontPortTemplate) validateRearPortPosition(formats strfmt.Registry) error {
	if swag.IsZero(m.RearPortPosition) { // not required
		return nil
	}

	if err := validate.MinimumInt("rear_port_position", "body", m.RearPortPosition, 1, false); err != nil {
		return err
	}

	if err := validate.MaximumInt("rear_port_position", "body", m.RearPortPosition, 1024, false); err != nil {
		return err
	}

	return nil
}

func (m *FrontPortTemplate) validateType(formats strfmt.Registry) error {

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

func (m *FrontPortTemplate) validateURL(formats strfmt.Registry) error {
	if swag.IsZero(m.URL) { // not required
		return nil
	}

	if err := validate.FormatOf("url", "body", "uri", m.URL.String(), formats); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this front port template based on the context it is used
func (m *FrontPortTemplate) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
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

	if err := m.contextValidateRearPort(ctx, formats); err != nil {
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

func (m *FrontPortTemplate) contextValidateDeviceType(ctx context.Context, formats strfmt.Registry) error {

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

func (m *FrontPortTemplate) contextValidateDisplay(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "display", "body", string(m.Display)); err != nil {
		return err
	}

	return nil
}

func (m *FrontPortTemplate) contextValidateID(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "id", "body", strfmt.UUID(m.ID)); err != nil {
		return err
	}

	return nil
}

func (m *FrontPortTemplate) contextValidateRearPort(ctx context.Context, formats strfmt.Registry) error {

	if m.RearPort != nil {
		if err := m.RearPort.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("rear_port")
			}
			return err
		}
	}

	return nil
}

func (m *FrontPortTemplate) contextValidateType(ctx context.Context, formats strfmt.Registry) error {

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

func (m *FrontPortTemplate) contextValidateURL(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "url", "body", strfmt.URI(m.URL)); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *FrontPortTemplate) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *FrontPortTemplate) UnmarshalBinary(b []byte) error {
	var res FrontPortTemplate
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// FrontPortTemplateType Type
//
// swagger:model FrontPortTemplateType
type FrontPortTemplateType struct {

	// label
	// Required: true
	// Enum: [8P8C 8P6C 8P4C 8P2C GG45 TERA 4P TERA 2P TERA 1P 110 Punch BNC MRJ21 FC LC LC/APC LSH LSH/APC MPO MTRJ SC SC/APC ST CS SN Splice]
	Label *string `json:"label"`

	// value
	// Required: true
	// Enum: [8p8c 8p6c 8p4c 8p2c gg45 tera-4p tera-2p tera-1p 110-punch bnc mrj21 fc lc lc-apc lsh lsh-apc mpo mtrj sc sc-apc st cs sn splice]
	Value *string `json:"value"`
}

// Validate validates this front port template type
func (m *FrontPortTemplateType) Validate(formats strfmt.Registry) error {
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

var frontPortTemplateTypeTypeLabelPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["8P8C","8P6C","8P4C","8P2C","GG45","TERA 4P","TERA 2P","TERA 1P","110 Punch","BNC","MRJ21","FC","LC","LC/APC","LSH","LSH/APC","MPO","MTRJ","SC","SC/APC","ST","CS","SN","Splice"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		frontPortTemplateTypeTypeLabelPropEnum = append(frontPortTemplateTypeTypeLabelPropEnum, v)
	}
}

const (

	// FrontPortTemplateTypeLabelNr8P8C captures enum value "8P8C"
	FrontPortTemplateTypeLabelNr8P8C string = "8P8C"

	// FrontPortTemplateTypeLabelNr8P6C captures enum value "8P6C"
	FrontPortTemplateTypeLabelNr8P6C string = "8P6C"

	// FrontPortTemplateTypeLabelNr8P4C captures enum value "8P4C"
	FrontPortTemplateTypeLabelNr8P4C string = "8P4C"

	// FrontPortTemplateTypeLabelNr8P2C captures enum value "8P2C"
	FrontPortTemplateTypeLabelNr8P2C string = "8P2C"

	// FrontPortTemplateTypeLabelGG45 captures enum value "GG45"
	FrontPortTemplateTypeLabelGG45 string = "GG45"

	// FrontPortTemplateTypeLabelTERA4P captures enum value "TERA 4P"
	FrontPortTemplateTypeLabelTERA4P string = "TERA 4P"

	// FrontPortTemplateTypeLabelTERA2P captures enum value "TERA 2P"
	FrontPortTemplateTypeLabelTERA2P string = "TERA 2P"

	// FrontPortTemplateTypeLabelTERA1P captures enum value "TERA 1P"
	FrontPortTemplateTypeLabelTERA1P string = "TERA 1P"

	// FrontPortTemplateTypeLabelNr110Punch captures enum value "110 Punch"
	FrontPortTemplateTypeLabelNr110Punch string = "110 Punch"

	// FrontPortTemplateTypeLabelBNC captures enum value "BNC"
	FrontPortTemplateTypeLabelBNC string = "BNC"

	// FrontPortTemplateTypeLabelMRJ21 captures enum value "MRJ21"
	FrontPortTemplateTypeLabelMRJ21 string = "MRJ21"

	// FrontPortTemplateTypeLabelFC captures enum value "FC"
	FrontPortTemplateTypeLabelFC string = "FC"

	// FrontPortTemplateTypeLabelLC captures enum value "LC"
	FrontPortTemplateTypeLabelLC string = "LC"

	// FrontPortTemplateTypeLabelLCAPC captures enum value "LC/APC"
	FrontPortTemplateTypeLabelLCAPC string = "LC/APC"

	// FrontPortTemplateTypeLabelLSH captures enum value "LSH"
	FrontPortTemplateTypeLabelLSH string = "LSH"

	// FrontPortTemplateTypeLabelLSHAPC captures enum value "LSH/APC"
	FrontPortTemplateTypeLabelLSHAPC string = "LSH/APC"

	// FrontPortTemplateTypeLabelMPO captures enum value "MPO"
	FrontPortTemplateTypeLabelMPO string = "MPO"

	// FrontPortTemplateTypeLabelMTRJ captures enum value "MTRJ"
	FrontPortTemplateTypeLabelMTRJ string = "MTRJ"

	// FrontPortTemplateTypeLabelSC captures enum value "SC"
	FrontPortTemplateTypeLabelSC string = "SC"

	// FrontPortTemplateTypeLabelSCAPC captures enum value "SC/APC"
	FrontPortTemplateTypeLabelSCAPC string = "SC/APC"

	// FrontPortTemplateTypeLabelST captures enum value "ST"
	FrontPortTemplateTypeLabelST string = "ST"

	// FrontPortTemplateTypeLabelCS captures enum value "CS"
	FrontPortTemplateTypeLabelCS string = "CS"

	// FrontPortTemplateTypeLabelSN captures enum value "SN"
	FrontPortTemplateTypeLabelSN string = "SN"

	// FrontPortTemplateTypeLabelSplice captures enum value "Splice"
	FrontPortTemplateTypeLabelSplice string = "Splice"
)

// prop value enum
func (m *FrontPortTemplateType) validateLabelEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, frontPortTemplateTypeTypeLabelPropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *FrontPortTemplateType) validateLabel(formats strfmt.Registry) error {

	if err := validate.Required("type"+"."+"label", "body", m.Label); err != nil {
		return err
	}

	// value enum
	if err := m.validateLabelEnum("type"+"."+"label", "body", *m.Label); err != nil {
		return err
	}

	return nil
}

var frontPortTemplateTypeTypeValuePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["8p8c","8p6c","8p4c","8p2c","gg45","tera-4p","tera-2p","tera-1p","110-punch","bnc","mrj21","fc","lc","lc-apc","lsh","lsh-apc","mpo","mtrj","sc","sc-apc","st","cs","sn","splice"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		frontPortTemplateTypeTypeValuePropEnum = append(frontPortTemplateTypeTypeValuePropEnum, v)
	}
}

const (

	// FrontPortTemplateTypeValueNr8p8c captures enum value "8p8c"
	FrontPortTemplateTypeValueNr8p8c string = "8p8c"

	// FrontPortTemplateTypeValueNr8p6c captures enum value "8p6c"
	FrontPortTemplateTypeValueNr8p6c string = "8p6c"

	// FrontPortTemplateTypeValueNr8p4c captures enum value "8p4c"
	FrontPortTemplateTypeValueNr8p4c string = "8p4c"

	// FrontPortTemplateTypeValueNr8p2c captures enum value "8p2c"
	FrontPortTemplateTypeValueNr8p2c string = "8p2c"

	// FrontPortTemplateTypeValueGg45 captures enum value "gg45"
	FrontPortTemplateTypeValueGg45 string = "gg45"

	// FrontPortTemplateTypeValueTeraDash4p captures enum value "tera-4p"
	FrontPortTemplateTypeValueTeraDash4p string = "tera-4p"

	// FrontPortTemplateTypeValueTeraDash2p captures enum value "tera-2p"
	FrontPortTemplateTypeValueTeraDash2p string = "tera-2p"

	// FrontPortTemplateTypeValueTeraDash1p captures enum value "tera-1p"
	FrontPortTemplateTypeValueTeraDash1p string = "tera-1p"

	// FrontPortTemplateTypeValueNr110DashPunch captures enum value "110-punch"
	FrontPortTemplateTypeValueNr110DashPunch string = "110-punch"

	// FrontPortTemplateTypeValueBnc captures enum value "bnc"
	FrontPortTemplateTypeValueBnc string = "bnc"

	// FrontPortTemplateTypeValueMrj21 captures enum value "mrj21"
	FrontPortTemplateTypeValueMrj21 string = "mrj21"

	// FrontPortTemplateTypeValueFc captures enum value "fc"
	FrontPortTemplateTypeValueFc string = "fc"

	// FrontPortTemplateTypeValueLc captures enum value "lc"
	FrontPortTemplateTypeValueLc string = "lc"

	// FrontPortTemplateTypeValueLcDashApc captures enum value "lc-apc"
	FrontPortTemplateTypeValueLcDashApc string = "lc-apc"

	// FrontPortTemplateTypeValueLsh captures enum value "lsh"
	FrontPortTemplateTypeValueLsh string = "lsh"

	// FrontPortTemplateTypeValueLshDashApc captures enum value "lsh-apc"
	FrontPortTemplateTypeValueLshDashApc string = "lsh-apc"

	// FrontPortTemplateTypeValueMpo captures enum value "mpo"
	FrontPortTemplateTypeValueMpo string = "mpo"

	// FrontPortTemplateTypeValueMtrj captures enum value "mtrj"
	FrontPortTemplateTypeValueMtrj string = "mtrj"

	// FrontPortTemplateTypeValueSc captures enum value "sc"
	FrontPortTemplateTypeValueSc string = "sc"

	// FrontPortTemplateTypeValueScDashApc captures enum value "sc-apc"
	FrontPortTemplateTypeValueScDashApc string = "sc-apc"

	// FrontPortTemplateTypeValueSt captures enum value "st"
	FrontPortTemplateTypeValueSt string = "st"

	// FrontPortTemplateTypeValueCs captures enum value "cs"
	FrontPortTemplateTypeValueCs string = "cs"

	// FrontPortTemplateTypeValueSn captures enum value "sn"
	FrontPortTemplateTypeValueSn string = "sn"

	// FrontPortTemplateTypeValueSplice captures enum value "splice"
	FrontPortTemplateTypeValueSplice string = "splice"
)

// prop value enum
func (m *FrontPortTemplateType) validateValueEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, frontPortTemplateTypeTypeValuePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *FrontPortTemplateType) validateValue(formats strfmt.Registry) error {

	if err := validate.Required("type"+"."+"value", "body", m.Value); err != nil {
		return err
	}

	// value enum
	if err := m.validateValueEnum("type"+"."+"value", "body", *m.Value); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this front port template type based on context it is used
func (m *FrontPortTemplateType) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *FrontPortTemplateType) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *FrontPortTemplateType) UnmarshalBinary(b []byte) error {
	var res FrontPortTemplateType
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
