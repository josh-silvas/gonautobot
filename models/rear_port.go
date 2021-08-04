package models

import (
	"context"
	"encoding/json"
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// RearPort rear port
//
// swagger:model RearPort
type RearPort struct {

	// cable
	Cable *NestedCable `json:"cable,omitempty"`

	// Cable peer
	//
	//
	// Return the appropriate serializer for the cable termination model.
	//
	// Read Only: true
	CablePeer map[string]*string `json:"cable_peer,omitempty"`

	// Cable peer type
	// Read Only: true
	CablePeerType string `json:"cable_peer_type,omitempty"`

	// Custom fields
	CustomFields interface{} `json:"custom_fields,omitempty"`

	// Description
	// Max Length: 200
	Description string `json:"description,omitempty"`

	// device
	// Required: true
	Device *NestedDevice `json:"device"`

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

	// tags
	Tags []*NestedTag `json:"tags"`

	// type
	// Required: true
	Type *RearPortType `json:"type"`

	// Url
	// Read Only: true
	// Format: uri
	URL strfmt.URI `json:"url,omitempty"`
}

// Validate validates this rear port
func (m *RearPort) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCable(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDescription(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDevice(formats); err != nil {
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

	if err := m.validateTags(formats); err != nil {
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

func (m *RearPort) validateCable(formats strfmt.Registry) error {
	if swag.IsZero(m.Cable) { // not required
		return nil
	}

	if m.Cable != nil {
		if err := m.Cable.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("cable")
			}
			return err
		}
	}

	return nil
}

func (m *RearPort) validateDescription(formats strfmt.Registry) error {
	if swag.IsZero(m.Description) { // not required
		return nil
	}

	if err := validate.MaxLength("description", "body", m.Description, 200); err != nil {
		return err
	}

	return nil
}

func (m *RearPort) validateDevice(formats strfmt.Registry) error {

	if err := validate.Required("device", "body", m.Device); err != nil {
		return err
	}

	if m.Device != nil {
		if err := m.Device.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("device")
			}
			return err
		}
	}

	return nil
}

func (m *RearPort) validateDisplay(formats strfmt.Registry) error {
	if swag.IsZero(m.Display) { // not required
		return nil
	}

	if err := validate.MinLength("display", "body", m.Display, 1); err != nil {
		return err
	}

	return nil
}

func (m *RearPort) validateID(formats strfmt.Registry) error {
	if swag.IsZero(m.ID) { // not required
		return nil
	}

	if err := validate.FormatOf("id", "body", "uuid", m.ID.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *RearPort) validateLabel(formats strfmt.Registry) error {
	if swag.IsZero(m.Label) { // not required
		return nil
	}

	if err := validate.MaxLength("label", "body", m.Label, 64); err != nil {
		return err
	}

	return nil
}

func (m *RearPort) validateName(formats strfmt.Registry) error {

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

func (m *RearPort) validatePositions(formats strfmt.Registry) error {
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

func (m *RearPort) validateTags(formats strfmt.Registry) error {
	if swag.IsZero(m.Tags) { // not required
		return nil
	}

	for i := 0; i < len(m.Tags); i++ {
		if swag.IsZero(m.Tags[i]) { // not required
			continue
		}

		if m.Tags[i] != nil {
			if err := m.Tags[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("tags" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *RearPort) validateType(formats strfmt.Registry) error {

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

func (m *RearPort) validateURL(formats strfmt.Registry) error {
	if swag.IsZero(m.URL) { // not required
		return nil
	}

	if err := validate.FormatOf("url", "body", "uri", m.URL.String(), formats); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this rear port based on the context it is used
func (m *RearPort) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateCable(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateCablePeer(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateCablePeerType(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateDevice(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateDisplay(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateID(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateTags(ctx, formats); err != nil {
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

func (m *RearPort) contextValidateCable(ctx context.Context, formats strfmt.Registry) error {

	if m.Cable != nil {
		if err := m.Cable.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("cable")
			}
			return err
		}
	}

	return nil
}

func (m *RearPort) contextValidateCablePeer(ctx context.Context, formats strfmt.Registry) error {

	return nil
}

func (m *RearPort) contextValidateCablePeerType(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "cable_peer_type", "body", string(m.CablePeerType)); err != nil {
		return err
	}

	return nil
}

func (m *RearPort) contextValidateDevice(ctx context.Context, formats strfmt.Registry) error {

	if m.Device != nil {
		if err := m.Device.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("device")
			}
			return err
		}
	}

	return nil
}

func (m *RearPort) contextValidateDisplay(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "display", "body", string(m.Display)); err != nil {
		return err
	}

	return nil
}

func (m *RearPort) contextValidateID(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "id", "body", strfmt.UUID(m.ID)); err != nil {
		return err
	}

	return nil
}

func (m *RearPort) contextValidateTags(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.Tags); i++ {

		if m.Tags[i] != nil {
			if err := m.Tags[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("tags" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *RearPort) contextValidateType(ctx context.Context, formats strfmt.Registry) error {

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

func (m *RearPort) contextValidateURL(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "url", "body", strfmt.URI(m.URL)); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *RearPort) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *RearPort) UnmarshalBinary(b []byte) error {
	var res RearPort
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// RearPortType Type
//
// swagger:model RearPortType
type RearPortType struct {

	// label
	// Required: true
	// Enum: [8P8C 8P6C 8P4C 8P2C GG45 TERA 4P TERA 2P TERA 1P 110 Punch BNC MRJ21 FC LC LC/APC LSH LSH/APC MPO MTRJ SC SC/APC ST CS SN Splice]
	Label *string `json:"label"`

	// value
	// Required: true
	// Enum: [8p8c 8p6c 8p4c 8p2c gg45 tera-4p tera-2p tera-1p 110-punch bnc mrj21 fc lc lc-apc lsh lsh-apc mpo mtrj sc sc-apc st cs sn splice]
	Value *string `json:"value"`
}

// Validate validates this rear port type
func (m *RearPortType) Validate(formats strfmt.Registry) error {
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

var rearPortTypeTypeLabelPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["8P8C","8P6C","8P4C","8P2C","GG45","TERA 4P","TERA 2P","TERA 1P","110 Punch","BNC","MRJ21","FC","LC","LC/APC","LSH","LSH/APC","MPO","MTRJ","SC","SC/APC","ST","CS","SN","Splice"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		rearPortTypeTypeLabelPropEnum = append(rearPortTypeTypeLabelPropEnum, v)
	}
}

const (

	// RearPortTypeLabelNr8P8C captures enum value "8P8C"
	RearPortTypeLabelNr8P8C string = "8P8C"

	// RearPortTypeLabelNr8P6C captures enum value "8P6C"
	RearPortTypeLabelNr8P6C string = "8P6C"

	// RearPortTypeLabelNr8P4C captures enum value "8P4C"
	RearPortTypeLabelNr8P4C string = "8P4C"

	// RearPortTypeLabelNr8P2C captures enum value "8P2C"
	RearPortTypeLabelNr8P2C string = "8P2C"

	// RearPortTypeLabelGG45 captures enum value "GG45"
	RearPortTypeLabelGG45 string = "GG45"

	// RearPortTypeLabelTERA4P captures enum value "TERA 4P"
	RearPortTypeLabelTERA4P string = "TERA 4P"

	// RearPortTypeLabelTERA2P captures enum value "TERA 2P"
	RearPortTypeLabelTERA2P string = "TERA 2P"

	// RearPortTypeLabelTERA1P captures enum value "TERA 1P"
	RearPortTypeLabelTERA1P string = "TERA 1P"

	// RearPortTypeLabelNr110Punch captures enum value "110 Punch"
	RearPortTypeLabelNr110Punch string = "110 Punch"

	// RearPortTypeLabelBNC captures enum value "BNC"
	RearPortTypeLabelBNC string = "BNC"

	// RearPortTypeLabelMRJ21 captures enum value "MRJ21"
	RearPortTypeLabelMRJ21 string = "MRJ21"

	// RearPortTypeLabelFC captures enum value "FC"
	RearPortTypeLabelFC string = "FC"

	// RearPortTypeLabelLC captures enum value "LC"
	RearPortTypeLabelLC string = "LC"

	// RearPortTypeLabelLCAPC captures enum value "LC/APC"
	RearPortTypeLabelLCAPC string = "LC/APC"

	// RearPortTypeLabelLSH captures enum value "LSH"
	RearPortTypeLabelLSH string = "LSH"

	// RearPortTypeLabelLSHAPC captures enum value "LSH/APC"
	RearPortTypeLabelLSHAPC string = "LSH/APC"

	// RearPortTypeLabelMPO captures enum value "MPO"
	RearPortTypeLabelMPO string = "MPO"

	// RearPortTypeLabelMTRJ captures enum value "MTRJ"
	RearPortTypeLabelMTRJ string = "MTRJ"

	// RearPortTypeLabelSC captures enum value "SC"
	RearPortTypeLabelSC string = "SC"

	// RearPortTypeLabelSCAPC captures enum value "SC/APC"
	RearPortTypeLabelSCAPC string = "SC/APC"

	// RearPortTypeLabelST captures enum value "ST"
	RearPortTypeLabelST string = "ST"

	// RearPortTypeLabelCS captures enum value "CS"
	RearPortTypeLabelCS string = "CS"

	// RearPortTypeLabelSN captures enum value "SN"
	RearPortTypeLabelSN string = "SN"

	// RearPortTypeLabelSplice captures enum value "Splice"
	RearPortTypeLabelSplice string = "Splice"
)

// prop value enum
func (m *RearPortType) validateLabelEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, rearPortTypeTypeLabelPropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *RearPortType) validateLabel(formats strfmt.Registry) error {

	if err := validate.Required("type"+"."+"label", "body", m.Label); err != nil {
		return err
	}

	// value enum
	if err := m.validateLabelEnum("type"+"."+"label", "body", *m.Label); err != nil {
		return err
	}

	return nil
}

var rearPortTypeTypeValuePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["8p8c","8p6c","8p4c","8p2c","gg45","tera-4p","tera-2p","tera-1p","110-punch","bnc","mrj21","fc","lc","lc-apc","lsh","lsh-apc","mpo","mtrj","sc","sc-apc","st","cs","sn","splice"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		rearPortTypeTypeValuePropEnum = append(rearPortTypeTypeValuePropEnum, v)
	}
}

const (

	// RearPortTypeValueNr8p8c captures enum value "8p8c"
	RearPortTypeValueNr8p8c string = "8p8c"

	// RearPortTypeValueNr8p6c captures enum value "8p6c"
	RearPortTypeValueNr8p6c string = "8p6c"

	// RearPortTypeValueNr8p4c captures enum value "8p4c"
	RearPortTypeValueNr8p4c string = "8p4c"

	// RearPortTypeValueNr8p2c captures enum value "8p2c"
	RearPortTypeValueNr8p2c string = "8p2c"

	// RearPortTypeValueGg45 captures enum value "gg45"
	RearPortTypeValueGg45 string = "gg45"

	// RearPortTypeValueTeraDash4p captures enum value "tera-4p"
	RearPortTypeValueTeraDash4p string = "tera-4p"

	// RearPortTypeValueTeraDash2p captures enum value "tera-2p"
	RearPortTypeValueTeraDash2p string = "tera-2p"

	// RearPortTypeValueTeraDash1p captures enum value "tera-1p"
	RearPortTypeValueTeraDash1p string = "tera-1p"

	// RearPortTypeValueNr110DashPunch captures enum value "110-punch"
	RearPortTypeValueNr110DashPunch string = "110-punch"

	// RearPortTypeValueBnc captures enum value "bnc"
	RearPortTypeValueBnc string = "bnc"

	// RearPortTypeValueMrj21 captures enum value "mrj21"
	RearPortTypeValueMrj21 string = "mrj21"

	// RearPortTypeValueFc captures enum value "fc"
	RearPortTypeValueFc string = "fc"

	// RearPortTypeValueLc captures enum value "lc"
	RearPortTypeValueLc string = "lc"

	// RearPortTypeValueLcDashApc captures enum value "lc-apc"
	RearPortTypeValueLcDashApc string = "lc-apc"

	// RearPortTypeValueLsh captures enum value "lsh"
	RearPortTypeValueLsh string = "lsh"

	// RearPortTypeValueLshDashApc captures enum value "lsh-apc"
	RearPortTypeValueLshDashApc string = "lsh-apc"

	// RearPortTypeValueMpo captures enum value "mpo"
	RearPortTypeValueMpo string = "mpo"

	// RearPortTypeValueMtrj captures enum value "mtrj"
	RearPortTypeValueMtrj string = "mtrj"

	// RearPortTypeValueSc captures enum value "sc"
	RearPortTypeValueSc string = "sc"

	// RearPortTypeValueScDashApc captures enum value "sc-apc"
	RearPortTypeValueScDashApc string = "sc-apc"

	// RearPortTypeValueSt captures enum value "st"
	RearPortTypeValueSt string = "st"

	// RearPortTypeValueCs captures enum value "cs"
	RearPortTypeValueCs string = "cs"

	// RearPortTypeValueSn captures enum value "sn"
	RearPortTypeValueSn string = "sn"

	// RearPortTypeValueSplice captures enum value "splice"
	RearPortTypeValueSplice string = "splice"
)

// prop value enum
func (m *RearPortType) validateValueEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, rearPortTypeTypeValuePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *RearPortType) validateValue(formats strfmt.Registry) error {

	if err := validate.Required("type"+"."+"value", "body", m.Value); err != nil {
		return err
	}

	// value enum
	if err := m.validateValueEnum("type"+"."+"value", "body", *m.Value); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this rear port type based on context it is used
func (m *RearPortType) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *RearPortType) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *RearPortType) UnmarshalBinary(b []byte) error {
	var res RearPortType
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
