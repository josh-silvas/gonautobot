package models

import (
	"context"
	"encoding/json"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// ConsoleServerPortTemplate console server port template
//
// swagger:model ConsoleServerPortTemplate
type ConsoleServerPortTemplate struct {

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

	// type
	Type *ConsoleServerPortTemplateType `json:"type,omitempty"`

	// Url
	// Read Only: true
	// Format: uri
	URL strfmt.URI `json:"url,omitempty"`
}

// Validate validates this console server port template
func (m *ConsoleServerPortTemplate) Validate(formats strfmt.Registry) error {
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

func (m *ConsoleServerPortTemplate) validateDescription(formats strfmt.Registry) error {
	if swag.IsZero(m.Description) { // not required
		return nil
	}

	if err := validate.MaxLength("description", "body", m.Description, 200); err != nil {
		return err
	}

	return nil
}

func (m *ConsoleServerPortTemplate) validateDeviceType(formats strfmt.Registry) error {

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

func (m *ConsoleServerPortTemplate) validateDisplay(formats strfmt.Registry) error {
	if swag.IsZero(m.Display) { // not required
		return nil
	}

	if err := validate.MinLength("display", "body", m.Display, 1); err != nil {
		return err
	}

	return nil
}

func (m *ConsoleServerPortTemplate) validateID(formats strfmt.Registry) error {
	if swag.IsZero(m.ID) { // not required
		return nil
	}

	if err := validate.FormatOf("id", "body", "uuid", m.ID.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *ConsoleServerPortTemplate) validateLabel(formats strfmt.Registry) error {
	if swag.IsZero(m.Label) { // not required
		return nil
	}

	if err := validate.MaxLength("label", "body", m.Label, 64); err != nil {
		return err
	}

	return nil
}

func (m *ConsoleServerPortTemplate) validateName(formats strfmt.Registry) error {

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

func (m *ConsoleServerPortTemplate) validateType(formats strfmt.Registry) error {
	if swag.IsZero(m.Type) { // not required
		return nil
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

func (m *ConsoleServerPortTemplate) validateURL(formats strfmt.Registry) error {
	if swag.IsZero(m.URL) { // not required
		return nil
	}

	if err := validate.FormatOf("url", "body", "uri", m.URL.String(), formats); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this console server port template based on the context it is used
func (m *ConsoleServerPortTemplate) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
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

func (m *ConsoleServerPortTemplate) contextValidateDeviceType(ctx context.Context, formats strfmt.Registry) error {

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

func (m *ConsoleServerPortTemplate) contextValidateDisplay(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "display", "body", string(m.Display)); err != nil {
		return err
	}

	return nil
}

func (m *ConsoleServerPortTemplate) contextValidateID(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "id", "body", strfmt.UUID(m.ID)); err != nil {
		return err
	}

	return nil
}

func (m *ConsoleServerPortTemplate) contextValidateType(ctx context.Context, formats strfmt.Registry) error {

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

func (m *ConsoleServerPortTemplate) contextValidateURL(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "url", "body", strfmt.URI(m.URL)); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *ConsoleServerPortTemplate) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ConsoleServerPortTemplate) UnmarshalBinary(b []byte) error {
	var res ConsoleServerPortTemplate
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// ConsoleServerPortTemplateType Type
//
// swagger:model ConsoleServerPortTemplateType
type ConsoleServerPortTemplateType struct {

	// label
	// Required: true
	// Enum: [DE-9 DB-25 RJ-11 RJ-12 RJ-45 USB Type A USB Type B USB Type C USB Mini A USB Mini B USB Micro A USB Micro B Other]
	Label *string `json:"label"`

	// value
	// Required: true
	// Enum: [de-9 db-25 rj-11 rj-12 rj-45 usb-a usb-b usb-c usb-mini-a usb-mini-b usb-micro-a usb-micro-b other]
	Value *string `json:"value"`
}

// Validate validates this console server port template type
func (m *ConsoleServerPortTemplateType) Validate(formats strfmt.Registry) error {
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

var consoleServerPortTemplateTypeTypeLabelPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["DE-9","DB-25","RJ-11","RJ-12","RJ-45","USB Type A","USB Type B","USB Type C","USB Mini A","USB Mini B","USB Micro A","USB Micro B","Other"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		consoleServerPortTemplateTypeTypeLabelPropEnum = append(consoleServerPortTemplateTypeTypeLabelPropEnum, v)
	}
}

const (

	// ConsoleServerPortTemplateTypeLabelDEDash9 captures enum value "DE-9"
	ConsoleServerPortTemplateTypeLabelDEDash9 string = "DE-9"

	// ConsoleServerPortTemplateTypeLabelDBDash25 captures enum value "DB-25"
	ConsoleServerPortTemplateTypeLabelDBDash25 string = "DB-25"

	// ConsoleServerPortTemplateTypeLabelRJDash11 captures enum value "RJ-11"
	ConsoleServerPortTemplateTypeLabelRJDash11 string = "RJ-11"

	// ConsoleServerPortTemplateTypeLabelRJDash12 captures enum value "RJ-12"
	ConsoleServerPortTemplateTypeLabelRJDash12 string = "RJ-12"

	// ConsoleServerPortTemplateTypeLabelRJDash45 captures enum value "RJ-45"
	ConsoleServerPortTemplateTypeLabelRJDash45 string = "RJ-45"

	// ConsoleServerPortTemplateTypeLabelUSBTypeA captures enum value "USB Type A"
	ConsoleServerPortTemplateTypeLabelUSBTypeA string = "USB Type A"

	// ConsoleServerPortTemplateTypeLabelUSBTypeB captures enum value "USB Type B"
	ConsoleServerPortTemplateTypeLabelUSBTypeB string = "USB Type B"

	// ConsoleServerPortTemplateTypeLabelUSBTypeC captures enum value "USB Type C"
	ConsoleServerPortTemplateTypeLabelUSBTypeC string = "USB Type C"

	// ConsoleServerPortTemplateTypeLabelUSBMiniA captures enum value "USB Mini A"
	ConsoleServerPortTemplateTypeLabelUSBMiniA string = "USB Mini A"

	// ConsoleServerPortTemplateTypeLabelUSBMiniB captures enum value "USB Mini B"
	ConsoleServerPortTemplateTypeLabelUSBMiniB string = "USB Mini B"

	// ConsoleServerPortTemplateTypeLabelUSBMicroA captures enum value "USB Micro A"
	ConsoleServerPortTemplateTypeLabelUSBMicroA string = "USB Micro A"

	// ConsoleServerPortTemplateTypeLabelUSBMicroB captures enum value "USB Micro B"
	ConsoleServerPortTemplateTypeLabelUSBMicroB string = "USB Micro B"

	// ConsoleServerPortTemplateTypeLabelOther captures enum value "Other"
	ConsoleServerPortTemplateTypeLabelOther string = "Other"
)

// prop value enum
func (m *ConsoleServerPortTemplateType) validateLabelEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, consoleServerPortTemplateTypeTypeLabelPropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *ConsoleServerPortTemplateType) validateLabel(formats strfmt.Registry) error {

	if err := validate.Required("type"+"."+"label", "body", m.Label); err != nil {
		return err
	}

	// value enum
	if err := m.validateLabelEnum("type"+"."+"label", "body", *m.Label); err != nil {
		return err
	}

	return nil
}

var consoleServerPortTemplateTypeTypeValuePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["de-9","db-25","rj-11","rj-12","rj-45","usb-a","usb-b","usb-c","usb-mini-a","usb-mini-b","usb-micro-a","usb-micro-b","other"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		consoleServerPortTemplateTypeTypeValuePropEnum = append(consoleServerPortTemplateTypeTypeValuePropEnum, v)
	}
}

const (

	// ConsoleServerPortTemplateTypeValueDeDash9 captures enum value "de-9"
	ConsoleServerPortTemplateTypeValueDeDash9 string = "de-9"

	// ConsoleServerPortTemplateTypeValueDbDash25 captures enum value "db-25"
	ConsoleServerPortTemplateTypeValueDbDash25 string = "db-25"

	// ConsoleServerPortTemplateTypeValueRjDash11 captures enum value "rj-11"
	ConsoleServerPortTemplateTypeValueRjDash11 string = "rj-11"

	// ConsoleServerPortTemplateTypeValueRjDash12 captures enum value "rj-12"
	ConsoleServerPortTemplateTypeValueRjDash12 string = "rj-12"

	// ConsoleServerPortTemplateTypeValueRjDash45 captures enum value "rj-45"
	ConsoleServerPortTemplateTypeValueRjDash45 string = "rj-45"

	// ConsoleServerPortTemplateTypeValueUsbDasha captures enum value "usb-a"
	ConsoleServerPortTemplateTypeValueUsbDasha string = "usb-a"

	// ConsoleServerPortTemplateTypeValueUsbDashb captures enum value "usb-b"
	ConsoleServerPortTemplateTypeValueUsbDashb string = "usb-b"

	// ConsoleServerPortTemplateTypeValueUsbDashc captures enum value "usb-c"
	ConsoleServerPortTemplateTypeValueUsbDashc string = "usb-c"

	// ConsoleServerPortTemplateTypeValueUsbDashMiniDasha captures enum value "usb-mini-a"
	ConsoleServerPortTemplateTypeValueUsbDashMiniDasha string = "usb-mini-a"

	// ConsoleServerPortTemplateTypeValueUsbDashMiniDashb captures enum value "usb-mini-b"
	ConsoleServerPortTemplateTypeValueUsbDashMiniDashb string = "usb-mini-b"

	// ConsoleServerPortTemplateTypeValueUsbDashMicroDasha captures enum value "usb-micro-a"
	ConsoleServerPortTemplateTypeValueUsbDashMicroDasha string = "usb-micro-a"

	// ConsoleServerPortTemplateTypeValueUsbDashMicroDashb captures enum value "usb-micro-b"
	ConsoleServerPortTemplateTypeValueUsbDashMicroDashb string = "usb-micro-b"

	// ConsoleServerPortTemplateTypeValueOther captures enum value "other"
	ConsoleServerPortTemplateTypeValueOther string = "other"
)

// prop value enum
func (m *ConsoleServerPortTemplateType) validateValueEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, consoleServerPortTemplateTypeTypeValuePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *ConsoleServerPortTemplateType) validateValue(formats strfmt.Registry) error {

	if err := validate.Required("type"+"."+"value", "body", m.Value); err != nil {
		return err
	}

	// value enum
	if err := m.validateValueEnum("type"+"."+"value", "body", *m.Value); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this console server port template type based on context it is used
func (m *ConsoleServerPortTemplateType) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *ConsoleServerPortTemplateType) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ConsoleServerPortTemplateType) UnmarshalBinary(b []byte) error {
	var res ConsoleServerPortTemplateType
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
