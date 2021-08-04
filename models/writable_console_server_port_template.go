package models

import (
	"context"
	"encoding/json"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// WritableConsoleServerPortTemplate writable console server port template
//
// swagger:model WritableConsoleServerPortTemplate
type WritableConsoleServerPortTemplate struct {

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

	// Type
	// Enum: [de-9 db-25 rj-11 rj-12 rj-45 usb-a usb-b usb-c usb-mini-a usb-mini-b usb-micro-a usb-micro-b other]
	Type string `json:"type,omitempty"`

	// Url
	// Read Only: true
	// Format: uri
	URL strfmt.URI `json:"url,omitempty"`
}

// Validate validates this writable console server port template
func (m *WritableConsoleServerPortTemplate) Validate(formats strfmt.Registry) error {
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

func (m *WritableConsoleServerPortTemplate) validateDescription(formats strfmt.Registry) error {
	if swag.IsZero(m.Description) { // not required
		return nil
	}

	if err := validate.MaxLength("description", "body", m.Description, 200); err != nil {
		return err
	}

	return nil
}

func (m *WritableConsoleServerPortTemplate) validateDeviceType(formats strfmt.Registry) error {

	if err := validate.Required("device_type", "body", m.DeviceType); err != nil {
		return err
	}

	if err := validate.FormatOf("device_type", "body", "uuid", m.DeviceType.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *WritableConsoleServerPortTemplate) validateDisplay(formats strfmt.Registry) error {
	if swag.IsZero(m.Display) { // not required
		return nil
	}

	if err := validate.MinLength("display", "body", m.Display, 1); err != nil {
		return err
	}

	return nil
}

func (m *WritableConsoleServerPortTemplate) validateID(formats strfmt.Registry) error {
	if swag.IsZero(m.ID) { // not required
		return nil
	}

	if err := validate.FormatOf("id", "body", "uuid", m.ID.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *WritableConsoleServerPortTemplate) validateLabel(formats strfmt.Registry) error {
	if swag.IsZero(m.Label) { // not required
		return nil
	}

	if err := validate.MaxLength("label", "body", m.Label, 64); err != nil {
		return err
	}

	return nil
}

func (m *WritableConsoleServerPortTemplate) validateName(formats strfmt.Registry) error {

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

var writableConsoleServerPortTemplateTypeTypePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["de-9","db-25","rj-11","rj-12","rj-45","usb-a","usb-b","usb-c","usb-mini-a","usb-mini-b","usb-micro-a","usb-micro-b","other"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		writableConsoleServerPortTemplateTypeTypePropEnum = append(writableConsoleServerPortTemplateTypeTypePropEnum, v)
	}
}

const (

	// WritableConsoleServerPortTemplateTypeDeDash9 captures enum value "de-9"
	WritableConsoleServerPortTemplateTypeDeDash9 string = "de-9"

	// WritableConsoleServerPortTemplateTypeDbDash25 captures enum value "db-25"
	WritableConsoleServerPortTemplateTypeDbDash25 string = "db-25"

	// WritableConsoleServerPortTemplateTypeRjDash11 captures enum value "rj-11"
	WritableConsoleServerPortTemplateTypeRjDash11 string = "rj-11"

	// WritableConsoleServerPortTemplateTypeRjDash12 captures enum value "rj-12"
	WritableConsoleServerPortTemplateTypeRjDash12 string = "rj-12"

	// WritableConsoleServerPortTemplateTypeRjDash45 captures enum value "rj-45"
	WritableConsoleServerPortTemplateTypeRjDash45 string = "rj-45"

	// WritableConsoleServerPortTemplateTypeUsbDasha captures enum value "usb-a"
	WritableConsoleServerPortTemplateTypeUsbDasha string = "usb-a"

	// WritableConsoleServerPortTemplateTypeUsbDashb captures enum value "usb-b"
	WritableConsoleServerPortTemplateTypeUsbDashb string = "usb-b"

	// WritableConsoleServerPortTemplateTypeUsbDashc captures enum value "usb-c"
	WritableConsoleServerPortTemplateTypeUsbDashc string = "usb-c"

	// WritableConsoleServerPortTemplateTypeUsbDashMiniDasha captures enum value "usb-mini-a"
	WritableConsoleServerPortTemplateTypeUsbDashMiniDasha string = "usb-mini-a"

	// WritableConsoleServerPortTemplateTypeUsbDashMiniDashb captures enum value "usb-mini-b"
	WritableConsoleServerPortTemplateTypeUsbDashMiniDashb string = "usb-mini-b"

	// WritableConsoleServerPortTemplateTypeUsbDashMicroDasha captures enum value "usb-micro-a"
	WritableConsoleServerPortTemplateTypeUsbDashMicroDasha string = "usb-micro-a"

	// WritableConsoleServerPortTemplateTypeUsbDashMicroDashb captures enum value "usb-micro-b"
	WritableConsoleServerPortTemplateTypeUsbDashMicroDashb string = "usb-micro-b"

	// WritableConsoleServerPortTemplateTypeOther captures enum value "other"
	WritableConsoleServerPortTemplateTypeOther string = "other"
)

// prop value enum
func (m *WritableConsoleServerPortTemplate) validateTypeEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, writableConsoleServerPortTemplateTypeTypePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *WritableConsoleServerPortTemplate) validateType(formats strfmt.Registry) error {
	if swag.IsZero(m.Type) { // not required
		return nil
	}

	// value enum
	if err := m.validateTypeEnum("type", "body", m.Type); err != nil {
		return err
	}

	return nil
}

func (m *WritableConsoleServerPortTemplate) validateURL(formats strfmt.Registry) error {
	if swag.IsZero(m.URL) { // not required
		return nil
	}

	if err := validate.FormatOf("url", "body", "uri", m.URL.String(), formats); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this writable console server port template based on the context it is used
func (m *WritableConsoleServerPortTemplate) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
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

func (m *WritableConsoleServerPortTemplate) contextValidateDisplay(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "display", "body", string(m.Display)); err != nil {
		return err
	}

	return nil
}

func (m *WritableConsoleServerPortTemplate) contextValidateID(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "id", "body", strfmt.UUID(m.ID)); err != nil {
		return err
	}

	return nil
}

func (m *WritableConsoleServerPortTemplate) contextValidateURL(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "url", "body", strfmt.URI(m.URL)); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *WritableConsoleServerPortTemplate) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *WritableConsoleServerPortTemplate) UnmarshalBinary(b []byte) error {
	var res WritableConsoleServerPortTemplate
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
