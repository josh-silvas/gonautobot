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

// WritableFrontPort writable front port
//
// swagger:model WritableFrontPort
type WritableFrontPort struct {

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

	// Computed fields
	// Read Only: true
	ComputedFields string `json:"computed_fields,omitempty"`

	// Custom fields
	CustomFields interface{} `json:"custom_fields,omitempty"`

	// Description
	// Max Length: 200
	Description string `json:"description,omitempty"`

	// Device
	// Required: true
	// Format: uuid
	Device *strfmt.UUID `json:"device"`

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

	// Rear port
	// Required: true
	// Format: uuid
	RearPort *strfmt.UUID `json:"rear_port"`

	// Rear port position
	// Maximum: 1024
	// Minimum: 1
	RearPortPosition int64 `json:"rear_port_position,omitempty"`

	// tags
	Tags []*NestedTag `json:"tags"`

	// Type
	// Required: true
	// Enum: [8p8c 8p6c 8p4c 8p2c gg45 tera-4p tera-2p tera-1p 110-punch bnc mrj21 fc lc lc-apc lsh lsh-apc mpo mtrj sc sc-apc st cs sn splice]
	Type *string `json:"type"`

	// Url
	// Read Only: true
	// Format: uri
	URL strfmt.URI `json:"url,omitempty"`
}

// Validate validates this writable front port
func (m *WritableFrontPort) Validate(formats strfmt.Registry) error {
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

	if err := m.validateRearPort(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateRearPortPosition(formats); err != nil {
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

func (m *WritableFrontPort) validateCable(formats strfmt.Registry) error {
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

func (m *WritableFrontPort) validateDescription(formats strfmt.Registry) error {
	if swag.IsZero(m.Description) { // not required
		return nil
	}

	if err := validate.MaxLength("description", "body", m.Description, 200); err != nil {
		return err
	}

	return nil
}

func (m *WritableFrontPort) validateDevice(formats strfmt.Registry) error {

	if err := validate.Required("device", "body", m.Device); err != nil {
		return err
	}

	if err := validate.FormatOf("device", "body", "uuid", m.Device.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *WritableFrontPort) validateDisplay(formats strfmt.Registry) error {
	if swag.IsZero(m.Display) { // not required
		return nil
	}

	if err := validate.MinLength("display", "body", m.Display, 1); err != nil {
		return err
	}

	return nil
}

func (m *WritableFrontPort) validateID(formats strfmt.Registry) error {
	if swag.IsZero(m.ID) { // not required
		return nil
	}

	if err := validate.FormatOf("id", "body", "uuid", m.ID.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *WritableFrontPort) validateLabel(formats strfmt.Registry) error {
	if swag.IsZero(m.Label) { // not required
		return nil
	}

	if err := validate.MaxLength("label", "body", m.Label, 64); err != nil {
		return err
	}

	return nil
}

func (m *WritableFrontPort) validateName(formats strfmt.Registry) error {

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

func (m *WritableFrontPort) validateRearPort(formats strfmt.Registry) error {

	if err := validate.Required("rear_port", "body", m.RearPort); err != nil {
		return err
	}

	if err := validate.FormatOf("rear_port", "body", "uuid", m.RearPort.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *WritableFrontPort) validateRearPortPosition(formats strfmt.Registry) error {
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

func (m *WritableFrontPort) validateTags(formats strfmt.Registry) error {
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

var writableFrontPortTypeTypePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["8p8c","8p6c","8p4c","8p2c","gg45","tera-4p","tera-2p","tera-1p","110-punch","bnc","mrj21","fc","lc","lc-apc","lsh","lsh-apc","mpo","mtrj","sc","sc-apc","st","cs","sn","splice"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		writableFrontPortTypeTypePropEnum = append(writableFrontPortTypeTypePropEnum, v)
	}
}

const (

	// WritableFrontPortTypeNr8p8c captures enum value "8p8c"
	WritableFrontPortTypeNr8p8c string = "8p8c"

	// WritableFrontPortTypeNr8p6c captures enum value "8p6c"
	WritableFrontPortTypeNr8p6c string = "8p6c"

	// WritableFrontPortTypeNr8p4c captures enum value "8p4c"
	WritableFrontPortTypeNr8p4c string = "8p4c"

	// WritableFrontPortTypeNr8p2c captures enum value "8p2c"
	WritableFrontPortTypeNr8p2c string = "8p2c"

	// WritableFrontPortTypeGg45 captures enum value "gg45"
	WritableFrontPortTypeGg45 string = "gg45"

	// WritableFrontPortTypeTeraDash4p captures enum value "tera-4p"
	WritableFrontPortTypeTeraDash4p string = "tera-4p"

	// WritableFrontPortTypeTeraDash2p captures enum value "tera-2p"
	WritableFrontPortTypeTeraDash2p string = "tera-2p"

	// WritableFrontPortTypeTeraDash1p captures enum value "tera-1p"
	WritableFrontPortTypeTeraDash1p string = "tera-1p"

	// WritableFrontPortTypeNr110DashPunch captures enum value "110-punch"
	WritableFrontPortTypeNr110DashPunch string = "110-punch"

	// WritableFrontPortTypeBnc captures enum value "bnc"
	WritableFrontPortTypeBnc string = "bnc"

	// WritableFrontPortTypeMrj21 captures enum value "mrj21"
	WritableFrontPortTypeMrj21 string = "mrj21"

	// WritableFrontPortTypeFc captures enum value "fc"
	WritableFrontPortTypeFc string = "fc"

	// WritableFrontPortTypeLc captures enum value "lc"
	WritableFrontPortTypeLc string = "lc"

	// WritableFrontPortTypeLcDashApc captures enum value "lc-apc"
	WritableFrontPortTypeLcDashApc string = "lc-apc"

	// WritableFrontPortTypeLsh captures enum value "lsh"
	WritableFrontPortTypeLsh string = "lsh"

	// WritableFrontPortTypeLshDashApc captures enum value "lsh-apc"
	WritableFrontPortTypeLshDashApc string = "lsh-apc"

	// WritableFrontPortTypeMpo captures enum value "mpo"
	WritableFrontPortTypeMpo string = "mpo"

	// WritableFrontPortTypeMtrj captures enum value "mtrj"
	WritableFrontPortTypeMtrj string = "mtrj"

	// WritableFrontPortTypeSc captures enum value "sc"
	WritableFrontPortTypeSc string = "sc"

	// WritableFrontPortTypeScDashApc captures enum value "sc-apc"
	WritableFrontPortTypeScDashApc string = "sc-apc"

	// WritableFrontPortTypeSt captures enum value "st"
	WritableFrontPortTypeSt string = "st"

	// WritableFrontPortTypeCs captures enum value "cs"
	WritableFrontPortTypeCs string = "cs"

	// WritableFrontPortTypeSn captures enum value "sn"
	WritableFrontPortTypeSn string = "sn"

	// WritableFrontPortTypeSplice captures enum value "splice"
	WritableFrontPortTypeSplice string = "splice"
)

// prop value enum
func (m *WritableFrontPort) validateTypeEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, writableFrontPortTypeTypePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *WritableFrontPort) validateType(formats strfmt.Registry) error {

	if err := validate.Required("type", "body", m.Type); err != nil {
		return err
	}

	// value enum
	if err := m.validateTypeEnum("type", "body", *m.Type); err != nil {
		return err
	}

	return nil
}

func (m *WritableFrontPort) validateURL(formats strfmt.Registry) error {
	if swag.IsZero(m.URL) { // not required
		return nil
	}

	if err := validate.FormatOf("url", "body", "uri", m.URL.String(), formats); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this writable front port based on the context it is used
func (m *WritableFrontPort) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
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

	if err := m.contextValidateComputedFields(ctx, formats); err != nil {
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

	if err := m.contextValidateURL(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *WritableFrontPort) contextValidateCable(ctx context.Context, formats strfmt.Registry) error {

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

func (m *WritableFrontPort) contextValidateCablePeer(ctx context.Context, formats strfmt.Registry) error {

	return nil
}

func (m *WritableFrontPort) contextValidateCablePeerType(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "cable_peer_type", "body", string(m.CablePeerType)); err != nil {
		return err
	}

	return nil
}

func (m *WritableFrontPort) contextValidateComputedFields(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "computed_fields", "body", string(m.ComputedFields)); err != nil {
		return err
	}

	return nil
}

func (m *WritableFrontPort) contextValidateDisplay(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "display", "body", string(m.Display)); err != nil {
		return err
	}

	return nil
}

func (m *WritableFrontPort) contextValidateID(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "id", "body", strfmt.UUID(m.ID)); err != nil {
		return err
	}

	return nil
}

func (m *WritableFrontPort) contextValidateTags(ctx context.Context, formats strfmt.Registry) error {

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

func (m *WritableFrontPort) contextValidateURL(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "url", "body", strfmt.URI(m.URL)); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *WritableFrontPort) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *WritableFrontPort) UnmarshalBinary(b []byte) error {
	var res WritableFrontPort
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
