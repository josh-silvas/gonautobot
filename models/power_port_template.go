package models

import (
	"context"
	"encoding/json"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// PowerPortTemplate power port template
//
// swagger:model PowerPortTemplate
type PowerPortTemplate struct {

	// Allocated draw
	//
	// Allocated power draw (watts)
	// Maximum: 32767
	// Minimum: 1
	AllocatedDraw *int64 `json:"allocated_draw,omitempty"`

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

	// Maximum draw
	//
	// Maximum power draw (watts)
	// Maximum: 32767
	// Minimum: 1
	MaximumDraw *int64 `json:"maximum_draw,omitempty"`

	// Name
	// Required: true
	// Max Length: 64
	// Min Length: 1
	Name *string `json:"name"`

	// type
	Type *PowerPortTemplateType `json:"type,omitempty"`

	// Url
	// Read Only: true
	// Format: uri
	URL strfmt.URI `json:"url,omitempty"`
}

// Validate validates this power port template
func (m *PowerPortTemplate) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAllocatedDraw(formats); err != nil {
		res = append(res, err)
	}

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

	if err := m.validateMaximumDraw(formats); err != nil {
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

func (m *PowerPortTemplate) validateAllocatedDraw(formats strfmt.Registry) error {
	if swag.IsZero(m.AllocatedDraw) { // not required
		return nil
	}

	if err := validate.MinimumInt("allocated_draw", "body", *m.AllocatedDraw, 1, false); err != nil {
		return err
	}

	if err := validate.MaximumInt("allocated_draw", "body", *m.AllocatedDraw, 32767, false); err != nil {
		return err
	}

	return nil
}

func (m *PowerPortTemplate) validateDescription(formats strfmt.Registry) error {
	if swag.IsZero(m.Description) { // not required
		return nil
	}

	if err := validate.MaxLength("description", "body", m.Description, 200); err != nil {
		return err
	}

	return nil
}

func (m *PowerPortTemplate) validateDeviceType(formats strfmt.Registry) error {

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

func (m *PowerPortTemplate) validateDisplay(formats strfmt.Registry) error {
	if swag.IsZero(m.Display) { // not required
		return nil
	}

	if err := validate.MinLength("display", "body", m.Display, 1); err != nil {
		return err
	}

	return nil
}

func (m *PowerPortTemplate) validateID(formats strfmt.Registry) error {
	if swag.IsZero(m.ID) { // not required
		return nil
	}

	if err := validate.FormatOf("id", "body", "uuid", m.ID.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *PowerPortTemplate) validateLabel(formats strfmt.Registry) error {
	if swag.IsZero(m.Label) { // not required
		return nil
	}

	if err := validate.MaxLength("label", "body", m.Label, 64); err != nil {
		return err
	}

	return nil
}

func (m *PowerPortTemplate) validateMaximumDraw(formats strfmt.Registry) error {
	if swag.IsZero(m.MaximumDraw) { // not required
		return nil
	}

	if err := validate.MinimumInt("maximum_draw", "body", *m.MaximumDraw, 1, false); err != nil {
		return err
	}

	if err := validate.MaximumInt("maximum_draw", "body", *m.MaximumDraw, 32767, false); err != nil {
		return err
	}

	return nil
}

func (m *PowerPortTemplate) validateName(formats strfmt.Registry) error {

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

func (m *PowerPortTemplate) validateType(formats strfmt.Registry) error {
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

func (m *PowerPortTemplate) validateURL(formats strfmt.Registry) error {
	if swag.IsZero(m.URL) { // not required
		return nil
	}

	if err := validate.FormatOf("url", "body", "uri", m.URL.String(), formats); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this power port template based on the context it is used
func (m *PowerPortTemplate) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
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

func (m *PowerPortTemplate) contextValidateDeviceType(ctx context.Context, formats strfmt.Registry) error {

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

func (m *PowerPortTemplate) contextValidateDisplay(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "display", "body", string(m.Display)); err != nil {
		return err
	}

	return nil
}

func (m *PowerPortTemplate) contextValidateID(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "id", "body", strfmt.UUID(m.ID)); err != nil {
		return err
	}

	return nil
}

func (m *PowerPortTemplate) contextValidateType(ctx context.Context, formats strfmt.Registry) error {

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

func (m *PowerPortTemplate) contextValidateURL(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "url", "body", strfmt.URI(m.URL)); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *PowerPortTemplate) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *PowerPortTemplate) UnmarshalBinary(b []byte) error {
	var res PowerPortTemplate
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// PowerPortTemplateType Type
//
// swagger:model PowerPortTemplateType
type PowerPortTemplateType struct {

	// label
	// Required: true
	// Enum: [C6 C8 C14 C16 C20 P+N+E 4H P+N+E 6H P+N+E 9H 2P+E 4H 2P+E 6H 2P+E 9H 3P+E 4H 3P+E 6H 3P+E 9H 3P+N+E 4H 3P+N+E 6H 3P+N+E 9H NEMA 1-15P NEMA 5-15P NEMA 5-20P NEMA 5-30P NEMA 5-50P NEMA 6-15P NEMA 6-20P NEMA 6-30P NEMA 6-50P NEMA 10-30P NEMA 10-50P NEMA 14-20P NEMA 14-30P NEMA 14-50P NEMA 14-60P NEMA 15-15P NEMA 15-20P NEMA 15-30P NEMA 15-50P NEMA 15-60P NEMA L1-15P NEMA L5-15P NEMA L5-20P NEMA L5-30P NEMA L5-50P NEMA L6-15P NEMA L6-20P NEMA L6-30P NEMA L6-50P NEMA L10-30P NEMA L14-20P NEMA L14-30P NEMA L14-50P NEMA L14-60P NEMA L15-20P NEMA L15-30P NEMA L15-50P NEMA L15-60P NEMA L21-20P NEMA L21-30P CS6361C CS6365C CS8165C CS8265C CS8365C CS8465C ITA Type E (CEE 7/5) ITA Type F (CEE 7/4) ITA Type E/F (CEE 7/7) ITA Type G (BS 1363) ITA Type H ITA Type I ITA Type J ITA Type K ITA Type L (CEI 23-50) ITA Type M (BS 546) ITA Type N ITA Type O USB Type A USB Type B USB Type C USB Mini A USB Mini B USB Micro A USB Micro B USB 3.0 Type B USB 3.0 Micro B]
	Label *string `json:"label"`

	// value
	// Required: true
	// Enum: [iec-60320-c6 iec-60320-c8 iec-60320-c14 iec-60320-c16 iec-60320-c20 iec-60309-p-n-e-4h iec-60309-p-n-e-6h iec-60309-p-n-e-9h iec-60309-2p-e-4h iec-60309-2p-e-6h iec-60309-2p-e-9h iec-60309-3p-e-4h iec-60309-3p-e-6h iec-60309-3p-e-9h iec-60309-3p-n-e-4h iec-60309-3p-n-e-6h iec-60309-3p-n-e-9h nema-1-15p nema-5-15p nema-5-20p nema-5-30p nema-5-50p nema-6-15p nema-6-20p nema-6-30p nema-6-50p nema-10-30p nema-10-50p nema-14-20p nema-14-30p nema-14-50p nema-14-60p nema-15-15p nema-15-20p nema-15-30p nema-15-50p nema-15-60p nema-l1-15p nema-l5-15p nema-l5-20p nema-l5-30p nema-l5-50p nema-l6-15p nema-l6-20p nema-l6-30p nema-l6-50p nema-l10-30p nema-l14-20p nema-l14-30p nema-l14-50p nema-l14-60p nema-l15-20p nema-l15-30p nema-l15-50p nema-l15-60p nema-l21-20p nema-l21-30p cs6361c cs6365c cs8165c cs8265c cs8365c cs8465c ita-e ita-f ita-ef ita-g ita-h ita-i ita-j ita-k ita-l ita-m ita-n ita-o usb-a usb-b usb-c usb-mini-a usb-mini-b usb-micro-a usb-micro-b usb-3-b usb-3-micro-b]
	Value *string `json:"value"`
}

// Validate validates this power port template type
func (m *PowerPortTemplateType) Validate(formats strfmt.Registry) error {
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

var powerPortTemplateTypeTypeLabelPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["C6","C8","C14","C16","C20","P+N+E 4H","P+N+E 6H","P+N+E 9H","2P+E 4H","2P+E 6H","2P+E 9H","3P+E 4H","3P+E 6H","3P+E 9H","3P+N+E 4H","3P+N+E 6H","3P+N+E 9H","NEMA 1-15P","NEMA 5-15P","NEMA 5-20P","NEMA 5-30P","NEMA 5-50P","NEMA 6-15P","NEMA 6-20P","NEMA 6-30P","NEMA 6-50P","NEMA 10-30P","NEMA 10-50P","NEMA 14-20P","NEMA 14-30P","NEMA 14-50P","NEMA 14-60P","NEMA 15-15P","NEMA 15-20P","NEMA 15-30P","NEMA 15-50P","NEMA 15-60P","NEMA L1-15P","NEMA L5-15P","NEMA L5-20P","NEMA L5-30P","NEMA L5-50P","NEMA L6-15P","NEMA L6-20P","NEMA L6-30P","NEMA L6-50P","NEMA L10-30P","NEMA L14-20P","NEMA L14-30P","NEMA L14-50P","NEMA L14-60P","NEMA L15-20P","NEMA L15-30P","NEMA L15-50P","NEMA L15-60P","NEMA L21-20P","NEMA L21-30P","CS6361C","CS6365C","CS8165C","CS8265C","CS8365C","CS8465C","ITA Type E (CEE 7/5)","ITA Type F (CEE 7/4)","ITA Type E/F (CEE 7/7)","ITA Type G (BS 1363)","ITA Type H","ITA Type I","ITA Type J","ITA Type K","ITA Type L (CEI 23-50)","ITA Type M (BS 546)","ITA Type N","ITA Type O","USB Type A","USB Type B","USB Type C","USB Mini A","USB Mini B","USB Micro A","USB Micro B","USB 3.0 Type B","USB 3.0 Micro B"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		powerPortTemplateTypeTypeLabelPropEnum = append(powerPortTemplateTypeTypeLabelPropEnum, v)
	}
}

const (

	// PowerPortTemplateTypeLabelC6 captures enum value "C6"
	PowerPortTemplateTypeLabelC6 string = "C6"

	// PowerPortTemplateTypeLabelC8 captures enum value "C8"
	PowerPortTemplateTypeLabelC8 string = "C8"

	// PowerPortTemplateTypeLabelC14 captures enum value "C14"
	PowerPortTemplateTypeLabelC14 string = "C14"

	// PowerPortTemplateTypeLabelC16 captures enum value "C16"
	PowerPortTemplateTypeLabelC16 string = "C16"

	// PowerPortTemplateTypeLabelC20 captures enum value "C20"
	PowerPortTemplateTypeLabelC20 string = "C20"

	// PowerPortTemplateTypeLabelPPlusNPlusE4H captures enum value "P+N+E 4H"
	PowerPortTemplateTypeLabelPPlusNPlusE4H string = "P+N+E 4H"

	// PowerPortTemplateTypeLabelPPlusNPlusE6H captures enum value "P+N+E 6H"
	PowerPortTemplateTypeLabelPPlusNPlusE6H string = "P+N+E 6H"

	// PowerPortTemplateTypeLabelPPlusNPlusE9H captures enum value "P+N+E 9H"
	PowerPortTemplateTypeLabelPPlusNPlusE9H string = "P+N+E 9H"

	// PowerPortTemplateTypeLabelNr2PPlusE4H captures enum value "2P+E 4H"
	PowerPortTemplateTypeLabelNr2PPlusE4H string = "2P+E 4H"

	// PowerPortTemplateTypeLabelNr2PPlusE6H captures enum value "2P+E 6H"
	PowerPortTemplateTypeLabelNr2PPlusE6H string = "2P+E 6H"

	// PowerPortTemplateTypeLabelNr2PPlusE9H captures enum value "2P+E 9H"
	PowerPortTemplateTypeLabelNr2PPlusE9H string = "2P+E 9H"

	// PowerPortTemplateTypeLabelNr3PPlusE4H captures enum value "3P+E 4H"
	PowerPortTemplateTypeLabelNr3PPlusE4H string = "3P+E 4H"

	// PowerPortTemplateTypeLabelNr3PPlusE6H captures enum value "3P+E 6H"
	PowerPortTemplateTypeLabelNr3PPlusE6H string = "3P+E 6H"

	// PowerPortTemplateTypeLabelNr3PPlusE9H captures enum value "3P+E 9H"
	PowerPortTemplateTypeLabelNr3PPlusE9H string = "3P+E 9H"

	// PowerPortTemplateTypeLabelNr3PPlusNPlusE4H captures enum value "3P+N+E 4H"
	PowerPortTemplateTypeLabelNr3PPlusNPlusE4H string = "3P+N+E 4H"

	// PowerPortTemplateTypeLabelNr3PPlusNPlusE6H captures enum value "3P+N+E 6H"
	PowerPortTemplateTypeLabelNr3PPlusNPlusE6H string = "3P+N+E 6H"

	// PowerPortTemplateTypeLabelNr3PPlusNPlusE9H captures enum value "3P+N+E 9H"
	PowerPortTemplateTypeLabelNr3PPlusNPlusE9H string = "3P+N+E 9H"

	// PowerPortTemplateTypeLabelNEMA1Dash15P captures enum value "NEMA 1-15P"
	PowerPortTemplateTypeLabelNEMA1Dash15P string = "NEMA 1-15P"

	// PowerPortTemplateTypeLabelNEMA5Dash15P captures enum value "NEMA 5-15P"
	PowerPortTemplateTypeLabelNEMA5Dash15P string = "NEMA 5-15P"

	// PowerPortTemplateTypeLabelNEMA5Dash20P captures enum value "NEMA 5-20P"
	PowerPortTemplateTypeLabelNEMA5Dash20P string = "NEMA 5-20P"

	// PowerPortTemplateTypeLabelNEMA5Dash30P captures enum value "NEMA 5-30P"
	PowerPortTemplateTypeLabelNEMA5Dash30P string = "NEMA 5-30P"

	// PowerPortTemplateTypeLabelNEMA5Dash50P captures enum value "NEMA 5-50P"
	PowerPortTemplateTypeLabelNEMA5Dash50P string = "NEMA 5-50P"

	// PowerPortTemplateTypeLabelNEMA6Dash15P captures enum value "NEMA 6-15P"
	PowerPortTemplateTypeLabelNEMA6Dash15P string = "NEMA 6-15P"

	// PowerPortTemplateTypeLabelNEMA6Dash20P captures enum value "NEMA 6-20P"
	PowerPortTemplateTypeLabelNEMA6Dash20P string = "NEMA 6-20P"

	// PowerPortTemplateTypeLabelNEMA6Dash30P captures enum value "NEMA 6-30P"
	PowerPortTemplateTypeLabelNEMA6Dash30P string = "NEMA 6-30P"

	// PowerPortTemplateTypeLabelNEMA6Dash50P captures enum value "NEMA 6-50P"
	PowerPortTemplateTypeLabelNEMA6Dash50P string = "NEMA 6-50P"

	// PowerPortTemplateTypeLabelNEMA10Dash30P captures enum value "NEMA 10-30P"
	PowerPortTemplateTypeLabelNEMA10Dash30P string = "NEMA 10-30P"

	// PowerPortTemplateTypeLabelNEMA10Dash50P captures enum value "NEMA 10-50P"
	PowerPortTemplateTypeLabelNEMA10Dash50P string = "NEMA 10-50P"

	// PowerPortTemplateTypeLabelNEMA14Dash20P captures enum value "NEMA 14-20P"
	PowerPortTemplateTypeLabelNEMA14Dash20P string = "NEMA 14-20P"

	// PowerPortTemplateTypeLabelNEMA14Dash30P captures enum value "NEMA 14-30P"
	PowerPortTemplateTypeLabelNEMA14Dash30P string = "NEMA 14-30P"

	// PowerPortTemplateTypeLabelNEMA14Dash50P captures enum value "NEMA 14-50P"
	PowerPortTemplateTypeLabelNEMA14Dash50P string = "NEMA 14-50P"

	// PowerPortTemplateTypeLabelNEMA14Dash60P captures enum value "NEMA 14-60P"
	PowerPortTemplateTypeLabelNEMA14Dash60P string = "NEMA 14-60P"

	// PowerPortTemplateTypeLabelNEMA15Dash15P captures enum value "NEMA 15-15P"
	PowerPortTemplateTypeLabelNEMA15Dash15P string = "NEMA 15-15P"

	// PowerPortTemplateTypeLabelNEMA15Dash20P captures enum value "NEMA 15-20P"
	PowerPortTemplateTypeLabelNEMA15Dash20P string = "NEMA 15-20P"

	// PowerPortTemplateTypeLabelNEMA15Dash30P captures enum value "NEMA 15-30P"
	PowerPortTemplateTypeLabelNEMA15Dash30P string = "NEMA 15-30P"

	// PowerPortTemplateTypeLabelNEMA15Dash50P captures enum value "NEMA 15-50P"
	PowerPortTemplateTypeLabelNEMA15Dash50P string = "NEMA 15-50P"

	// PowerPortTemplateTypeLabelNEMA15Dash60P captures enum value "NEMA 15-60P"
	PowerPortTemplateTypeLabelNEMA15Dash60P string = "NEMA 15-60P"

	// PowerPortTemplateTypeLabelNEMAL1Dash15P captures enum value "NEMA L1-15P"
	PowerPortTemplateTypeLabelNEMAL1Dash15P string = "NEMA L1-15P"

	// PowerPortTemplateTypeLabelNEMAL5Dash15P captures enum value "NEMA L5-15P"
	PowerPortTemplateTypeLabelNEMAL5Dash15P string = "NEMA L5-15P"

	// PowerPortTemplateTypeLabelNEMAL5Dash20P captures enum value "NEMA L5-20P"
	PowerPortTemplateTypeLabelNEMAL5Dash20P string = "NEMA L5-20P"

	// PowerPortTemplateTypeLabelNEMAL5Dash30P captures enum value "NEMA L5-30P"
	PowerPortTemplateTypeLabelNEMAL5Dash30P string = "NEMA L5-30P"

	// PowerPortTemplateTypeLabelNEMAL5Dash50P captures enum value "NEMA L5-50P"
	PowerPortTemplateTypeLabelNEMAL5Dash50P string = "NEMA L5-50P"

	// PowerPortTemplateTypeLabelNEMAL6Dash15P captures enum value "NEMA L6-15P"
	PowerPortTemplateTypeLabelNEMAL6Dash15P string = "NEMA L6-15P"

	// PowerPortTemplateTypeLabelNEMAL6Dash20P captures enum value "NEMA L6-20P"
	PowerPortTemplateTypeLabelNEMAL6Dash20P string = "NEMA L6-20P"

	// PowerPortTemplateTypeLabelNEMAL6Dash30P captures enum value "NEMA L6-30P"
	PowerPortTemplateTypeLabelNEMAL6Dash30P string = "NEMA L6-30P"

	// PowerPortTemplateTypeLabelNEMAL6Dash50P captures enum value "NEMA L6-50P"
	PowerPortTemplateTypeLabelNEMAL6Dash50P string = "NEMA L6-50P"

	// PowerPortTemplateTypeLabelNEMAL10Dash30P captures enum value "NEMA L10-30P"
	PowerPortTemplateTypeLabelNEMAL10Dash30P string = "NEMA L10-30P"

	// PowerPortTemplateTypeLabelNEMAL14Dash20P captures enum value "NEMA L14-20P"
	PowerPortTemplateTypeLabelNEMAL14Dash20P string = "NEMA L14-20P"

	// PowerPortTemplateTypeLabelNEMAL14Dash30P captures enum value "NEMA L14-30P"
	PowerPortTemplateTypeLabelNEMAL14Dash30P string = "NEMA L14-30P"

	// PowerPortTemplateTypeLabelNEMAL14Dash50P captures enum value "NEMA L14-50P"
	PowerPortTemplateTypeLabelNEMAL14Dash50P string = "NEMA L14-50P"

	// PowerPortTemplateTypeLabelNEMAL14Dash60P captures enum value "NEMA L14-60P"
	PowerPortTemplateTypeLabelNEMAL14Dash60P string = "NEMA L14-60P"

	// PowerPortTemplateTypeLabelNEMAL15Dash20P captures enum value "NEMA L15-20P"
	PowerPortTemplateTypeLabelNEMAL15Dash20P string = "NEMA L15-20P"

	// PowerPortTemplateTypeLabelNEMAL15Dash30P captures enum value "NEMA L15-30P"
	PowerPortTemplateTypeLabelNEMAL15Dash30P string = "NEMA L15-30P"

	// PowerPortTemplateTypeLabelNEMAL15Dash50P captures enum value "NEMA L15-50P"
	PowerPortTemplateTypeLabelNEMAL15Dash50P string = "NEMA L15-50P"

	// PowerPortTemplateTypeLabelNEMAL15Dash60P captures enum value "NEMA L15-60P"
	PowerPortTemplateTypeLabelNEMAL15Dash60P string = "NEMA L15-60P"

	// PowerPortTemplateTypeLabelNEMAL21Dash20P captures enum value "NEMA L21-20P"
	PowerPortTemplateTypeLabelNEMAL21Dash20P string = "NEMA L21-20P"

	// PowerPortTemplateTypeLabelNEMAL21Dash30P captures enum value "NEMA L21-30P"
	PowerPortTemplateTypeLabelNEMAL21Dash30P string = "NEMA L21-30P"

	// PowerPortTemplateTypeLabelCS6361C captures enum value "CS6361C"
	PowerPortTemplateTypeLabelCS6361C string = "CS6361C"

	// PowerPortTemplateTypeLabelCS6365C captures enum value "CS6365C"
	PowerPortTemplateTypeLabelCS6365C string = "CS6365C"

	// PowerPortTemplateTypeLabelCS8165C captures enum value "CS8165C"
	PowerPortTemplateTypeLabelCS8165C string = "CS8165C"

	// PowerPortTemplateTypeLabelCS8265C captures enum value "CS8265C"
	PowerPortTemplateTypeLabelCS8265C string = "CS8265C"

	// PowerPortTemplateTypeLabelCS8365C captures enum value "CS8365C"
	PowerPortTemplateTypeLabelCS8365C string = "CS8365C"

	// PowerPortTemplateTypeLabelCS8465C captures enum value "CS8465C"
	PowerPortTemplateTypeLabelCS8465C string = "CS8465C"

	// PowerPortTemplateTypeLabelITATypeECEE75 captures enum value "ITA Type E (CEE 7/5)"
	PowerPortTemplateTypeLabelITATypeECEE75 string = "ITA Type E (CEE 7/5)"

	// PowerPortTemplateTypeLabelITATypeFCEE74 captures enum value "ITA Type F (CEE 7/4)"
	PowerPortTemplateTypeLabelITATypeFCEE74 string = "ITA Type F (CEE 7/4)"

	// PowerPortTemplateTypeLabelITATypeEFCEE77 captures enum value "ITA Type E/F (CEE 7/7)"
	PowerPortTemplateTypeLabelITATypeEFCEE77 string = "ITA Type E/F (CEE 7/7)"

	// PowerPortTemplateTypeLabelITATypeGBS1363 captures enum value "ITA Type G (BS 1363)"
	PowerPortTemplateTypeLabelITATypeGBS1363 string = "ITA Type G (BS 1363)"

	// PowerPortTemplateTypeLabelITATypeH captures enum value "ITA Type H"
	PowerPortTemplateTypeLabelITATypeH string = "ITA Type H"

	// PowerPortTemplateTypeLabelITATypeI captures enum value "ITA Type I"
	PowerPortTemplateTypeLabelITATypeI string = "ITA Type I"

	// PowerPortTemplateTypeLabelITATypeJ captures enum value "ITA Type J"
	PowerPortTemplateTypeLabelITATypeJ string = "ITA Type J"

	// PowerPortTemplateTypeLabelITATypeK captures enum value "ITA Type K"
	PowerPortTemplateTypeLabelITATypeK string = "ITA Type K"

	// PowerPortTemplateTypeLabelITATypeLCEI23Dash50 captures enum value "ITA Type L (CEI 23-50)"
	PowerPortTemplateTypeLabelITATypeLCEI23Dash50 string = "ITA Type L (CEI 23-50)"

	// PowerPortTemplateTypeLabelITATypeMBS546 captures enum value "ITA Type M (BS 546)"
	PowerPortTemplateTypeLabelITATypeMBS546 string = "ITA Type M (BS 546)"

	// PowerPortTemplateTypeLabelITATypeN captures enum value "ITA Type N"
	PowerPortTemplateTypeLabelITATypeN string = "ITA Type N"

	// PowerPortTemplateTypeLabelITATypeO captures enum value "ITA Type O"
	PowerPortTemplateTypeLabelITATypeO string = "ITA Type O"

	// PowerPortTemplateTypeLabelUSBTypeA captures enum value "USB Type A"
	PowerPortTemplateTypeLabelUSBTypeA string = "USB Type A"

	// PowerPortTemplateTypeLabelUSBTypeB captures enum value "USB Type B"
	PowerPortTemplateTypeLabelUSBTypeB string = "USB Type B"

	// PowerPortTemplateTypeLabelUSBTypeC captures enum value "USB Type C"
	PowerPortTemplateTypeLabelUSBTypeC string = "USB Type C"

	// PowerPortTemplateTypeLabelUSBMiniA captures enum value "USB Mini A"
	PowerPortTemplateTypeLabelUSBMiniA string = "USB Mini A"

	// PowerPortTemplateTypeLabelUSBMiniB captures enum value "USB Mini B"
	PowerPortTemplateTypeLabelUSBMiniB string = "USB Mini B"

	// PowerPortTemplateTypeLabelUSBMicroA captures enum value "USB Micro A"
	PowerPortTemplateTypeLabelUSBMicroA string = "USB Micro A"

	// PowerPortTemplateTypeLabelUSBMicroB captures enum value "USB Micro B"
	PowerPortTemplateTypeLabelUSBMicroB string = "USB Micro B"

	// PowerPortTemplateTypeLabelUSB3Dot0TypeB captures enum value "USB 3.0 Type B"
	PowerPortTemplateTypeLabelUSB3Dot0TypeB string = "USB 3.0 Type B"

	// PowerPortTemplateTypeLabelUSB3Dot0MicroB captures enum value "USB 3.0 Micro B"
	PowerPortTemplateTypeLabelUSB3Dot0MicroB string = "USB 3.0 Micro B"
)

// prop value enum
func (m *PowerPortTemplateType) validateLabelEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, powerPortTemplateTypeTypeLabelPropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *PowerPortTemplateType) validateLabel(formats strfmt.Registry) error {

	if err := validate.Required("type"+"."+"label", "body", m.Label); err != nil {
		return err
	}

	// value enum
	if err := m.validateLabelEnum("type"+"."+"label", "body", *m.Label); err != nil {
		return err
	}

	return nil
}

var powerPortTemplateTypeTypeValuePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["iec-60320-c6","iec-60320-c8","iec-60320-c14","iec-60320-c16","iec-60320-c20","iec-60309-p-n-e-4h","iec-60309-p-n-e-6h","iec-60309-p-n-e-9h","iec-60309-2p-e-4h","iec-60309-2p-e-6h","iec-60309-2p-e-9h","iec-60309-3p-e-4h","iec-60309-3p-e-6h","iec-60309-3p-e-9h","iec-60309-3p-n-e-4h","iec-60309-3p-n-e-6h","iec-60309-3p-n-e-9h","nema-1-15p","nema-5-15p","nema-5-20p","nema-5-30p","nema-5-50p","nema-6-15p","nema-6-20p","nema-6-30p","nema-6-50p","nema-10-30p","nema-10-50p","nema-14-20p","nema-14-30p","nema-14-50p","nema-14-60p","nema-15-15p","nema-15-20p","nema-15-30p","nema-15-50p","nema-15-60p","nema-l1-15p","nema-l5-15p","nema-l5-20p","nema-l5-30p","nema-l5-50p","nema-l6-15p","nema-l6-20p","nema-l6-30p","nema-l6-50p","nema-l10-30p","nema-l14-20p","nema-l14-30p","nema-l14-50p","nema-l14-60p","nema-l15-20p","nema-l15-30p","nema-l15-50p","nema-l15-60p","nema-l21-20p","nema-l21-30p","cs6361c","cs6365c","cs8165c","cs8265c","cs8365c","cs8465c","ita-e","ita-f","ita-ef","ita-g","ita-h","ita-i","ita-j","ita-k","ita-l","ita-m","ita-n","ita-o","usb-a","usb-b","usb-c","usb-mini-a","usb-mini-b","usb-micro-a","usb-micro-b","usb-3-b","usb-3-micro-b"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		powerPortTemplateTypeTypeValuePropEnum = append(powerPortTemplateTypeTypeValuePropEnum, v)
	}
}

const (

	// PowerPortTemplateTypeValueIecDash60320DashC6 captures enum value "iec-60320-c6"
	PowerPortTemplateTypeValueIecDash60320DashC6 string = "iec-60320-c6"

	// PowerPortTemplateTypeValueIecDash60320DashC8 captures enum value "iec-60320-c8"
	PowerPortTemplateTypeValueIecDash60320DashC8 string = "iec-60320-c8"

	// PowerPortTemplateTypeValueIecDash60320DashC14 captures enum value "iec-60320-c14"
	PowerPortTemplateTypeValueIecDash60320DashC14 string = "iec-60320-c14"

	// PowerPortTemplateTypeValueIecDash60320DashC16 captures enum value "iec-60320-c16"
	PowerPortTemplateTypeValueIecDash60320DashC16 string = "iec-60320-c16"

	// PowerPortTemplateTypeValueIecDash60320DashC20 captures enum value "iec-60320-c20"
	PowerPortTemplateTypeValueIecDash60320DashC20 string = "iec-60320-c20"

	// PowerPortTemplateTypeValueIecDash60309DashpDashnDasheDash4h captures enum value "iec-60309-p-n-e-4h"
	PowerPortTemplateTypeValueIecDash60309DashpDashnDasheDash4h string = "iec-60309-p-n-e-4h"

	// PowerPortTemplateTypeValueIecDash60309DashpDashnDasheDash6h captures enum value "iec-60309-p-n-e-6h"
	PowerPortTemplateTypeValueIecDash60309DashpDashnDasheDash6h string = "iec-60309-p-n-e-6h"

	// PowerPortTemplateTypeValueIecDash60309DashpDashnDasheDash9h captures enum value "iec-60309-p-n-e-9h"
	PowerPortTemplateTypeValueIecDash60309DashpDashnDasheDash9h string = "iec-60309-p-n-e-9h"

	// PowerPortTemplateTypeValueIecDash60309Dash2pDasheDash4h captures enum value "iec-60309-2p-e-4h"
	PowerPortTemplateTypeValueIecDash60309Dash2pDasheDash4h string = "iec-60309-2p-e-4h"

	// PowerPortTemplateTypeValueIecDash60309Dash2pDasheDash6h captures enum value "iec-60309-2p-e-6h"
	PowerPortTemplateTypeValueIecDash60309Dash2pDasheDash6h string = "iec-60309-2p-e-6h"

	// PowerPortTemplateTypeValueIecDash60309Dash2pDasheDash9h captures enum value "iec-60309-2p-e-9h"
	PowerPortTemplateTypeValueIecDash60309Dash2pDasheDash9h string = "iec-60309-2p-e-9h"

	// PowerPortTemplateTypeValueIecDash60309Dash3pDasheDash4h captures enum value "iec-60309-3p-e-4h"
	PowerPortTemplateTypeValueIecDash60309Dash3pDasheDash4h string = "iec-60309-3p-e-4h"

	// PowerPortTemplateTypeValueIecDash60309Dash3pDasheDash6h captures enum value "iec-60309-3p-e-6h"
	PowerPortTemplateTypeValueIecDash60309Dash3pDasheDash6h string = "iec-60309-3p-e-6h"

	// PowerPortTemplateTypeValueIecDash60309Dash3pDasheDash9h captures enum value "iec-60309-3p-e-9h"
	PowerPortTemplateTypeValueIecDash60309Dash3pDasheDash9h string = "iec-60309-3p-e-9h"

	// PowerPortTemplateTypeValueIecDash60309Dash3pDashnDasheDash4h captures enum value "iec-60309-3p-n-e-4h"
	PowerPortTemplateTypeValueIecDash60309Dash3pDashnDasheDash4h string = "iec-60309-3p-n-e-4h"

	// PowerPortTemplateTypeValueIecDash60309Dash3pDashnDasheDash6h captures enum value "iec-60309-3p-n-e-6h"
	PowerPortTemplateTypeValueIecDash60309Dash3pDashnDasheDash6h string = "iec-60309-3p-n-e-6h"

	// PowerPortTemplateTypeValueIecDash60309Dash3pDashnDasheDash9h captures enum value "iec-60309-3p-n-e-9h"
	PowerPortTemplateTypeValueIecDash60309Dash3pDashnDasheDash9h string = "iec-60309-3p-n-e-9h"

	// PowerPortTemplateTypeValueNemaDash1Dash15p captures enum value "nema-1-15p"
	PowerPortTemplateTypeValueNemaDash1Dash15p string = "nema-1-15p"

	// PowerPortTemplateTypeValueNemaDash5Dash15p captures enum value "nema-5-15p"
	PowerPortTemplateTypeValueNemaDash5Dash15p string = "nema-5-15p"

	// PowerPortTemplateTypeValueNemaDash5Dash20p captures enum value "nema-5-20p"
	PowerPortTemplateTypeValueNemaDash5Dash20p string = "nema-5-20p"

	// PowerPortTemplateTypeValueNemaDash5Dash30p captures enum value "nema-5-30p"
	PowerPortTemplateTypeValueNemaDash5Dash30p string = "nema-5-30p"

	// PowerPortTemplateTypeValueNemaDash5Dash50p captures enum value "nema-5-50p"
	PowerPortTemplateTypeValueNemaDash5Dash50p string = "nema-5-50p"

	// PowerPortTemplateTypeValueNemaDash6Dash15p captures enum value "nema-6-15p"
	PowerPortTemplateTypeValueNemaDash6Dash15p string = "nema-6-15p"

	// PowerPortTemplateTypeValueNemaDash6Dash20p captures enum value "nema-6-20p"
	PowerPortTemplateTypeValueNemaDash6Dash20p string = "nema-6-20p"

	// PowerPortTemplateTypeValueNemaDash6Dash30p captures enum value "nema-6-30p"
	PowerPortTemplateTypeValueNemaDash6Dash30p string = "nema-6-30p"

	// PowerPortTemplateTypeValueNemaDash6Dash50p captures enum value "nema-6-50p"
	PowerPortTemplateTypeValueNemaDash6Dash50p string = "nema-6-50p"

	// PowerPortTemplateTypeValueNemaDash10Dash30p captures enum value "nema-10-30p"
	PowerPortTemplateTypeValueNemaDash10Dash30p string = "nema-10-30p"

	// PowerPortTemplateTypeValueNemaDash10Dash50p captures enum value "nema-10-50p"
	PowerPortTemplateTypeValueNemaDash10Dash50p string = "nema-10-50p"

	// PowerPortTemplateTypeValueNemaDash14Dash20p captures enum value "nema-14-20p"
	PowerPortTemplateTypeValueNemaDash14Dash20p string = "nema-14-20p"

	// PowerPortTemplateTypeValueNemaDash14Dash30p captures enum value "nema-14-30p"
	PowerPortTemplateTypeValueNemaDash14Dash30p string = "nema-14-30p"

	// PowerPortTemplateTypeValueNemaDash14Dash50p captures enum value "nema-14-50p"
	PowerPortTemplateTypeValueNemaDash14Dash50p string = "nema-14-50p"

	// PowerPortTemplateTypeValueNemaDash14Dash60p captures enum value "nema-14-60p"
	PowerPortTemplateTypeValueNemaDash14Dash60p string = "nema-14-60p"

	// PowerPortTemplateTypeValueNemaDash15Dash15p captures enum value "nema-15-15p"
	PowerPortTemplateTypeValueNemaDash15Dash15p string = "nema-15-15p"

	// PowerPortTemplateTypeValueNemaDash15Dash20p captures enum value "nema-15-20p"
	PowerPortTemplateTypeValueNemaDash15Dash20p string = "nema-15-20p"

	// PowerPortTemplateTypeValueNemaDash15Dash30p captures enum value "nema-15-30p"
	PowerPortTemplateTypeValueNemaDash15Dash30p string = "nema-15-30p"

	// PowerPortTemplateTypeValueNemaDash15Dash50p captures enum value "nema-15-50p"
	PowerPortTemplateTypeValueNemaDash15Dash50p string = "nema-15-50p"

	// PowerPortTemplateTypeValueNemaDash15Dash60p captures enum value "nema-15-60p"
	PowerPortTemplateTypeValueNemaDash15Dash60p string = "nema-15-60p"

	// PowerPortTemplateTypeValueNemaDashL1Dash15p captures enum value "nema-l1-15p"
	PowerPortTemplateTypeValueNemaDashL1Dash15p string = "nema-l1-15p"

	// PowerPortTemplateTypeValueNemaDashL5Dash15p captures enum value "nema-l5-15p"
	PowerPortTemplateTypeValueNemaDashL5Dash15p string = "nema-l5-15p"

	// PowerPortTemplateTypeValueNemaDashL5Dash20p captures enum value "nema-l5-20p"
	PowerPortTemplateTypeValueNemaDashL5Dash20p string = "nema-l5-20p"

	// PowerPortTemplateTypeValueNemaDashL5Dash30p captures enum value "nema-l5-30p"
	PowerPortTemplateTypeValueNemaDashL5Dash30p string = "nema-l5-30p"

	// PowerPortTemplateTypeValueNemaDashL5Dash50p captures enum value "nema-l5-50p"
	PowerPortTemplateTypeValueNemaDashL5Dash50p string = "nema-l5-50p"

	// PowerPortTemplateTypeValueNemaDashL6Dash15p captures enum value "nema-l6-15p"
	PowerPortTemplateTypeValueNemaDashL6Dash15p string = "nema-l6-15p"

	// PowerPortTemplateTypeValueNemaDashL6Dash20p captures enum value "nema-l6-20p"
	PowerPortTemplateTypeValueNemaDashL6Dash20p string = "nema-l6-20p"

	// PowerPortTemplateTypeValueNemaDashL6Dash30p captures enum value "nema-l6-30p"
	PowerPortTemplateTypeValueNemaDashL6Dash30p string = "nema-l6-30p"

	// PowerPortTemplateTypeValueNemaDashL6Dash50p captures enum value "nema-l6-50p"
	PowerPortTemplateTypeValueNemaDashL6Dash50p string = "nema-l6-50p"

	// PowerPortTemplateTypeValueNemaDashL10Dash30p captures enum value "nema-l10-30p"
	PowerPortTemplateTypeValueNemaDashL10Dash30p string = "nema-l10-30p"

	// PowerPortTemplateTypeValueNemaDashL14Dash20p captures enum value "nema-l14-20p"
	PowerPortTemplateTypeValueNemaDashL14Dash20p string = "nema-l14-20p"

	// PowerPortTemplateTypeValueNemaDashL14Dash30p captures enum value "nema-l14-30p"
	PowerPortTemplateTypeValueNemaDashL14Dash30p string = "nema-l14-30p"

	// PowerPortTemplateTypeValueNemaDashL14Dash50p captures enum value "nema-l14-50p"
	PowerPortTemplateTypeValueNemaDashL14Dash50p string = "nema-l14-50p"

	// PowerPortTemplateTypeValueNemaDashL14Dash60p captures enum value "nema-l14-60p"
	PowerPortTemplateTypeValueNemaDashL14Dash60p string = "nema-l14-60p"

	// PowerPortTemplateTypeValueNemaDashL15Dash20p captures enum value "nema-l15-20p"
	PowerPortTemplateTypeValueNemaDashL15Dash20p string = "nema-l15-20p"

	// PowerPortTemplateTypeValueNemaDashL15Dash30p captures enum value "nema-l15-30p"
	PowerPortTemplateTypeValueNemaDashL15Dash30p string = "nema-l15-30p"

	// PowerPortTemplateTypeValueNemaDashL15Dash50p captures enum value "nema-l15-50p"
	PowerPortTemplateTypeValueNemaDashL15Dash50p string = "nema-l15-50p"

	// PowerPortTemplateTypeValueNemaDashL15Dash60p captures enum value "nema-l15-60p"
	PowerPortTemplateTypeValueNemaDashL15Dash60p string = "nema-l15-60p"

	// PowerPortTemplateTypeValueNemaDashL21Dash20p captures enum value "nema-l21-20p"
	PowerPortTemplateTypeValueNemaDashL21Dash20p string = "nema-l21-20p"

	// PowerPortTemplateTypeValueNemaDashL21Dash30p captures enum value "nema-l21-30p"
	PowerPortTemplateTypeValueNemaDashL21Dash30p string = "nema-l21-30p"

	// PowerPortTemplateTypeValueCs6361c captures enum value "cs6361c"
	PowerPortTemplateTypeValueCs6361c string = "cs6361c"

	// PowerPortTemplateTypeValueCs6365c captures enum value "cs6365c"
	PowerPortTemplateTypeValueCs6365c string = "cs6365c"

	// PowerPortTemplateTypeValueCs8165c captures enum value "cs8165c"
	PowerPortTemplateTypeValueCs8165c string = "cs8165c"

	// PowerPortTemplateTypeValueCs8265c captures enum value "cs8265c"
	PowerPortTemplateTypeValueCs8265c string = "cs8265c"

	// PowerPortTemplateTypeValueCs8365c captures enum value "cs8365c"
	PowerPortTemplateTypeValueCs8365c string = "cs8365c"

	// PowerPortTemplateTypeValueCs8465c captures enum value "cs8465c"
	PowerPortTemplateTypeValueCs8465c string = "cs8465c"

	// PowerPortTemplateTypeValueItaDashe captures enum value "ita-e"
	PowerPortTemplateTypeValueItaDashe string = "ita-e"

	// PowerPortTemplateTypeValueItaDashf captures enum value "ita-f"
	PowerPortTemplateTypeValueItaDashf string = "ita-f"

	// PowerPortTemplateTypeValueItaDashEf captures enum value "ita-ef"
	PowerPortTemplateTypeValueItaDashEf string = "ita-ef"

	// PowerPortTemplateTypeValueItaDashg captures enum value "ita-g"
	PowerPortTemplateTypeValueItaDashg string = "ita-g"

	// PowerPortTemplateTypeValueItaDashh captures enum value "ita-h"
	PowerPortTemplateTypeValueItaDashh string = "ita-h"

	// PowerPortTemplateTypeValueItaDashi captures enum value "ita-i"
	PowerPortTemplateTypeValueItaDashi string = "ita-i"

	// PowerPortTemplateTypeValueItaDashj captures enum value "ita-j"
	PowerPortTemplateTypeValueItaDashj string = "ita-j"

	// PowerPortTemplateTypeValueItaDashk captures enum value "ita-k"
	PowerPortTemplateTypeValueItaDashk string = "ita-k"

	// PowerPortTemplateTypeValueItaDashl captures enum value "ita-l"
	PowerPortTemplateTypeValueItaDashl string = "ita-l"

	// PowerPortTemplateTypeValueItaDashm captures enum value "ita-m"
	PowerPortTemplateTypeValueItaDashm string = "ita-m"

	// PowerPortTemplateTypeValueItaDashn captures enum value "ita-n"
	PowerPortTemplateTypeValueItaDashn string = "ita-n"

	// PowerPortTemplateTypeValueItaDasho captures enum value "ita-o"
	PowerPortTemplateTypeValueItaDasho string = "ita-o"

	// PowerPortTemplateTypeValueUsbDasha captures enum value "usb-a"
	PowerPortTemplateTypeValueUsbDasha string = "usb-a"

	// PowerPortTemplateTypeValueUsbDashb captures enum value "usb-b"
	PowerPortTemplateTypeValueUsbDashb string = "usb-b"

	// PowerPortTemplateTypeValueUsbDashc captures enum value "usb-c"
	PowerPortTemplateTypeValueUsbDashc string = "usb-c"

	// PowerPortTemplateTypeValueUsbDashMiniDasha captures enum value "usb-mini-a"
	PowerPortTemplateTypeValueUsbDashMiniDasha string = "usb-mini-a"

	// PowerPortTemplateTypeValueUsbDashMiniDashb captures enum value "usb-mini-b"
	PowerPortTemplateTypeValueUsbDashMiniDashb string = "usb-mini-b"

	// PowerPortTemplateTypeValueUsbDashMicroDasha captures enum value "usb-micro-a"
	PowerPortTemplateTypeValueUsbDashMicroDasha string = "usb-micro-a"

	// PowerPortTemplateTypeValueUsbDashMicroDashb captures enum value "usb-micro-b"
	PowerPortTemplateTypeValueUsbDashMicroDashb string = "usb-micro-b"

	// PowerPortTemplateTypeValueUsbDash3Dashb captures enum value "usb-3-b"
	PowerPortTemplateTypeValueUsbDash3Dashb string = "usb-3-b"

	// PowerPortTemplateTypeValueUsbDash3DashMicroDashb captures enum value "usb-3-micro-b"
	PowerPortTemplateTypeValueUsbDash3DashMicroDashb string = "usb-3-micro-b"
)

// prop value enum
func (m *PowerPortTemplateType) validateValueEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, powerPortTemplateTypeTypeValuePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *PowerPortTemplateType) validateValue(formats strfmt.Registry) error {

	if err := validate.Required("type"+"."+"value", "body", m.Value); err != nil {
		return err
	}

	// value enum
	if err := m.validateValueEnum("type"+"."+"value", "body", *m.Value); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this power port template type based on context it is used
func (m *PowerPortTemplateType) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *PowerPortTemplateType) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *PowerPortTemplateType) UnmarshalBinary(b []byte) error {
	var res PowerPortTemplateType
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
