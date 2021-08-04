package models

import (
	"context"
	"encoding/json"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// PowerOutletTemplate power outlet template
//
// swagger:model PowerOutletTemplate
type PowerOutletTemplate struct {

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

	// feed leg
	FeedLeg *PowerOutletTemplateFeedLeg `json:"feed_leg,omitempty"`

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

	// power port
	PowerPort *NestedPowerPortTemplate `json:"power_port,omitempty"`

	// type
	Type *PowerOutletTemplateType `json:"type,omitempty"`

	// Url
	// Read Only: true
	// Format: uri
	URL strfmt.URI `json:"url,omitempty"`
}

// Validate validates this power outlet template
func (m *PowerOutletTemplate) Validate(formats strfmt.Registry) error {
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

	if err := m.validateFeedLeg(formats); err != nil {
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

	if err := m.validatePowerPort(formats); err != nil {
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

func (m *PowerOutletTemplate) validateDescription(formats strfmt.Registry) error {
	if swag.IsZero(m.Description) { // not required
		return nil
	}

	if err := validate.MaxLength("description", "body", m.Description, 200); err != nil {
		return err
	}

	return nil
}

func (m *PowerOutletTemplate) validateDeviceType(formats strfmt.Registry) error {

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

func (m *PowerOutletTemplate) validateDisplay(formats strfmt.Registry) error {
	if swag.IsZero(m.Display) { // not required
		return nil
	}

	if err := validate.MinLength("display", "body", m.Display, 1); err != nil {
		return err
	}

	return nil
}

func (m *PowerOutletTemplate) validateFeedLeg(formats strfmt.Registry) error {
	if swag.IsZero(m.FeedLeg) { // not required
		return nil
	}

	if m.FeedLeg != nil {
		if err := m.FeedLeg.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("feed_leg")
			}
			return err
		}
	}

	return nil
}

func (m *PowerOutletTemplate) validateID(formats strfmt.Registry) error {
	if swag.IsZero(m.ID) { // not required
		return nil
	}

	if err := validate.FormatOf("id", "body", "uuid", m.ID.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *PowerOutletTemplate) validateLabel(formats strfmt.Registry) error {
	if swag.IsZero(m.Label) { // not required
		return nil
	}

	if err := validate.MaxLength("label", "body", m.Label, 64); err != nil {
		return err
	}

	return nil
}

func (m *PowerOutletTemplate) validateName(formats strfmt.Registry) error {

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

func (m *PowerOutletTemplate) validatePowerPort(formats strfmt.Registry) error {
	if swag.IsZero(m.PowerPort) { // not required
		return nil
	}

	if m.PowerPort != nil {
		if err := m.PowerPort.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("power_port")
			}
			return err
		}
	}

	return nil
}

func (m *PowerOutletTemplate) validateType(formats strfmt.Registry) error {
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

func (m *PowerOutletTemplate) validateURL(formats strfmt.Registry) error {
	if swag.IsZero(m.URL) { // not required
		return nil
	}

	if err := validate.FormatOf("url", "body", "uri", m.URL.String(), formats); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this power outlet template based on the context it is used
func (m *PowerOutletTemplate) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateDeviceType(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateDisplay(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateFeedLeg(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateID(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidatePowerPort(ctx, formats); err != nil {
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

func (m *PowerOutletTemplate) contextValidateDeviceType(ctx context.Context, formats strfmt.Registry) error {

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

func (m *PowerOutletTemplate) contextValidateDisplay(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "display", "body", string(m.Display)); err != nil {
		return err
	}

	return nil
}

func (m *PowerOutletTemplate) contextValidateFeedLeg(ctx context.Context, formats strfmt.Registry) error {

	if m.FeedLeg != nil {
		if err := m.FeedLeg.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("feed_leg")
			}
			return err
		}
	}

	return nil
}

func (m *PowerOutletTemplate) contextValidateID(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "id", "body", strfmt.UUID(m.ID)); err != nil {
		return err
	}

	return nil
}

func (m *PowerOutletTemplate) contextValidatePowerPort(ctx context.Context, formats strfmt.Registry) error {

	if m.PowerPort != nil {
		if err := m.PowerPort.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("power_port")
			}
			return err
		}
	}

	return nil
}

func (m *PowerOutletTemplate) contextValidateType(ctx context.Context, formats strfmt.Registry) error {

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

func (m *PowerOutletTemplate) contextValidateURL(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "url", "body", strfmt.URI(m.URL)); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *PowerOutletTemplate) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *PowerOutletTemplate) UnmarshalBinary(b []byte) error {
	var res PowerOutletTemplate
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// PowerOutletTemplateFeedLeg Feed leg
//
// swagger:model PowerOutletTemplateFeedLeg
type PowerOutletTemplateFeedLeg struct {

	// label
	// Required: true
	// Enum: [A B C]
	Label *string `json:"label"`

	// value
	// Required: true
	// Enum: [A B C]
	Value *string `json:"value"`
}

// Validate validates this power outlet template feed leg
func (m *PowerOutletTemplateFeedLeg) Validate(formats strfmt.Registry) error {
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

var powerOutletTemplateFeedLegTypeLabelPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["A","B","C"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		powerOutletTemplateFeedLegTypeLabelPropEnum = append(powerOutletTemplateFeedLegTypeLabelPropEnum, v)
	}
}

const (

	// PowerOutletTemplateFeedLegLabelA captures enum value "A"
	PowerOutletTemplateFeedLegLabelA string = "A"

	// PowerOutletTemplateFeedLegLabelB captures enum value "B"
	PowerOutletTemplateFeedLegLabelB string = "B"

	// PowerOutletTemplateFeedLegLabelC captures enum value "C"
	PowerOutletTemplateFeedLegLabelC string = "C"
)

// prop value enum
func (m *PowerOutletTemplateFeedLeg) validateLabelEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, powerOutletTemplateFeedLegTypeLabelPropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *PowerOutletTemplateFeedLeg) validateLabel(formats strfmt.Registry) error {

	if err := validate.Required("feed_leg"+"."+"label", "body", m.Label); err != nil {
		return err
	}

	// value enum
	if err := m.validateLabelEnum("feed_leg"+"."+"label", "body", *m.Label); err != nil {
		return err
	}

	return nil
}

var powerOutletTemplateFeedLegTypeValuePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["A","B","C"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		powerOutletTemplateFeedLegTypeValuePropEnum = append(powerOutletTemplateFeedLegTypeValuePropEnum, v)
	}
}

const (

	// PowerOutletTemplateFeedLegValueA captures enum value "A"
	PowerOutletTemplateFeedLegValueA string = "A"

	// PowerOutletTemplateFeedLegValueB captures enum value "B"
	PowerOutletTemplateFeedLegValueB string = "B"

	// PowerOutletTemplateFeedLegValueC captures enum value "C"
	PowerOutletTemplateFeedLegValueC string = "C"
)

// prop value enum
func (m *PowerOutletTemplateFeedLeg) validateValueEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, powerOutletTemplateFeedLegTypeValuePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *PowerOutletTemplateFeedLeg) validateValue(formats strfmt.Registry) error {

	if err := validate.Required("feed_leg"+"."+"value", "body", m.Value); err != nil {
		return err
	}

	// value enum
	if err := m.validateValueEnum("feed_leg"+"."+"value", "body", *m.Value); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this power outlet template feed leg based on context it is used
func (m *PowerOutletTemplateFeedLeg) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *PowerOutletTemplateFeedLeg) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *PowerOutletTemplateFeedLeg) UnmarshalBinary(b []byte) error {
	var res PowerOutletTemplateFeedLeg
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// PowerOutletTemplateType Type
//
// swagger:model PowerOutletTemplateType
type PowerOutletTemplateType struct {

	// label
	// Required: true
	// Enum: [C5 C7 C13 C15 C19 P+N+E 4H P+N+E 6H P+N+E 9H 2P+E 4H 2P+E 6H 2P+E 9H 3P+E 4H 3P+E 6H 3P+E 9H 3P+N+E 4H 3P+N+E 6H 3P+N+E 9H NEMA 1-15R NEMA 5-15R NEMA 5-20R NEMA 5-30R NEMA 5-50R NEMA 6-15R NEMA 6-20R NEMA 6-30R NEMA 6-50R NEMA 10-30R NEMA 10-50R NEMA 14-20R NEMA 14-30R NEMA 14-50R NEMA 14-60R NEMA 15-15R NEMA 15-20R NEMA 15-30R NEMA 15-50R NEMA 15-60R NEMA L1-15R NEMA L5-15R NEMA L5-20R NEMA L5-30R NEMA L5-50R NEMA L6-15R NEMA L6-20R NEMA L6-30R NEMA L6-50R NEMA L10-30R NEMA L14-20R NEMA L14-30R NEMA L14-50R NEMA L14-60R NEMA L15-20R NEMA L15-30R NEMA L15-50R NEMA L15-60R NEMA L21-20R NEMA L21-30R CS6360C CS6364C CS8164C CS8264C CS8364C CS8464C ITA Type E (CEE7/5) ITA Type F (CEE7/3) ITA Type G (BS 1363) ITA Type H ITA Type I ITA Type J ITA Type K ITA Type L (CEI 23-50) ITA Type M (BS 546) ITA Type N ITA Type O USB Type A USB Micro B USB Type C HDOT Cx]
	Label *string `json:"label"`

	// value
	// Required: true
	// Enum: [iec-60320-c5 iec-60320-c7 iec-60320-c13 iec-60320-c15 iec-60320-c19 iec-60309-p-n-e-4h iec-60309-p-n-e-6h iec-60309-p-n-e-9h iec-60309-2p-e-4h iec-60309-2p-e-6h iec-60309-2p-e-9h iec-60309-3p-e-4h iec-60309-3p-e-6h iec-60309-3p-e-9h iec-60309-3p-n-e-4h iec-60309-3p-n-e-6h iec-60309-3p-n-e-9h nema-1-15r nema-5-15r nema-5-20r nema-5-30r nema-5-50r nema-6-15r nema-6-20r nema-6-30r nema-6-50r nema-10-30r nema-10-50r nema-14-20r nema-14-30r nema-14-50r nema-14-60r nema-15-15r nema-15-20r nema-15-30r nema-15-50r nema-15-60r nema-l1-15r nema-l5-15r nema-l5-20r nema-l5-30r nema-l5-50r nema-l6-15r nema-l6-20r nema-l6-30r nema-l6-50r nema-l10-30r nema-l14-20r nema-l14-30r nema-l14-50r nema-l14-60r nema-l15-20r nema-l15-30r nema-l15-50r nema-l15-60r nema-l21-20r nema-l21-30r CS6360C CS6364C CS8164C CS8264C CS8364C CS8464C ita-e ita-f ita-g ita-h ita-i ita-j ita-k ita-l ita-m ita-n ita-o usb-a usb-micro-b usb-c hdot-cx]
	Value *string `json:"value"`
}

// Validate validates this power outlet template type
func (m *PowerOutletTemplateType) Validate(formats strfmt.Registry) error {
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

var powerOutletTemplateTypeTypeLabelPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["C5","C7","C13","C15","C19","P+N+E 4H","P+N+E 6H","P+N+E 9H","2P+E 4H","2P+E 6H","2P+E 9H","3P+E 4H","3P+E 6H","3P+E 9H","3P+N+E 4H","3P+N+E 6H","3P+N+E 9H","NEMA 1-15R","NEMA 5-15R","NEMA 5-20R","NEMA 5-30R","NEMA 5-50R","NEMA 6-15R","NEMA 6-20R","NEMA 6-30R","NEMA 6-50R","NEMA 10-30R","NEMA 10-50R","NEMA 14-20R","NEMA 14-30R","NEMA 14-50R","NEMA 14-60R","NEMA 15-15R","NEMA 15-20R","NEMA 15-30R","NEMA 15-50R","NEMA 15-60R","NEMA L1-15R","NEMA L5-15R","NEMA L5-20R","NEMA L5-30R","NEMA L5-50R","NEMA L6-15R","NEMA L6-20R","NEMA L6-30R","NEMA L6-50R","NEMA L10-30R","NEMA L14-20R","NEMA L14-30R","NEMA L14-50R","NEMA L14-60R","NEMA L15-20R","NEMA L15-30R","NEMA L15-50R","NEMA L15-60R","NEMA L21-20R","NEMA L21-30R","CS6360C","CS6364C","CS8164C","CS8264C","CS8364C","CS8464C","ITA Type E (CEE7/5)","ITA Type F (CEE7/3)","ITA Type G (BS 1363)","ITA Type H","ITA Type I","ITA Type J","ITA Type K","ITA Type L (CEI 23-50)","ITA Type M (BS 546)","ITA Type N","ITA Type O","USB Type A","USB Micro B","USB Type C","HDOT Cx"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		powerOutletTemplateTypeTypeLabelPropEnum = append(powerOutletTemplateTypeTypeLabelPropEnum, v)
	}
}

const (

	// PowerOutletTemplateTypeLabelC5 captures enum value "C5"
	PowerOutletTemplateTypeLabelC5 string = "C5"

	// PowerOutletTemplateTypeLabelC7 captures enum value "C7"
	PowerOutletTemplateTypeLabelC7 string = "C7"

	// PowerOutletTemplateTypeLabelC13 captures enum value "C13"
	PowerOutletTemplateTypeLabelC13 string = "C13"

	// PowerOutletTemplateTypeLabelC15 captures enum value "C15"
	PowerOutletTemplateTypeLabelC15 string = "C15"

	// PowerOutletTemplateTypeLabelC19 captures enum value "C19"
	PowerOutletTemplateTypeLabelC19 string = "C19"

	// PowerOutletTemplateTypeLabelPPlusNPlusE4H captures enum value "P+N+E 4H"
	PowerOutletTemplateTypeLabelPPlusNPlusE4H string = "P+N+E 4H"

	// PowerOutletTemplateTypeLabelPPlusNPlusE6H captures enum value "P+N+E 6H"
	PowerOutletTemplateTypeLabelPPlusNPlusE6H string = "P+N+E 6H"

	// PowerOutletTemplateTypeLabelPPlusNPlusE9H captures enum value "P+N+E 9H"
	PowerOutletTemplateTypeLabelPPlusNPlusE9H string = "P+N+E 9H"

	// PowerOutletTemplateTypeLabelNr2PPlusE4H captures enum value "2P+E 4H"
	PowerOutletTemplateTypeLabelNr2PPlusE4H string = "2P+E 4H"

	// PowerOutletTemplateTypeLabelNr2PPlusE6H captures enum value "2P+E 6H"
	PowerOutletTemplateTypeLabelNr2PPlusE6H string = "2P+E 6H"

	// PowerOutletTemplateTypeLabelNr2PPlusE9H captures enum value "2P+E 9H"
	PowerOutletTemplateTypeLabelNr2PPlusE9H string = "2P+E 9H"

	// PowerOutletTemplateTypeLabelNr3PPlusE4H captures enum value "3P+E 4H"
	PowerOutletTemplateTypeLabelNr3PPlusE4H string = "3P+E 4H"

	// PowerOutletTemplateTypeLabelNr3PPlusE6H captures enum value "3P+E 6H"
	PowerOutletTemplateTypeLabelNr3PPlusE6H string = "3P+E 6H"

	// PowerOutletTemplateTypeLabelNr3PPlusE9H captures enum value "3P+E 9H"
	PowerOutletTemplateTypeLabelNr3PPlusE9H string = "3P+E 9H"

	// PowerOutletTemplateTypeLabelNr3PPlusNPlusE4H captures enum value "3P+N+E 4H"
	PowerOutletTemplateTypeLabelNr3PPlusNPlusE4H string = "3P+N+E 4H"

	// PowerOutletTemplateTypeLabelNr3PPlusNPlusE6H captures enum value "3P+N+E 6H"
	PowerOutletTemplateTypeLabelNr3PPlusNPlusE6H string = "3P+N+E 6H"

	// PowerOutletTemplateTypeLabelNr3PPlusNPlusE9H captures enum value "3P+N+E 9H"
	PowerOutletTemplateTypeLabelNr3PPlusNPlusE9H string = "3P+N+E 9H"

	// PowerOutletTemplateTypeLabelNEMA1Dash15R captures enum value "NEMA 1-15R"
	PowerOutletTemplateTypeLabelNEMA1Dash15R string = "NEMA 1-15R"

	// PowerOutletTemplateTypeLabelNEMA5Dash15R captures enum value "NEMA 5-15R"
	PowerOutletTemplateTypeLabelNEMA5Dash15R string = "NEMA 5-15R"

	// PowerOutletTemplateTypeLabelNEMA5Dash20R captures enum value "NEMA 5-20R"
	PowerOutletTemplateTypeLabelNEMA5Dash20R string = "NEMA 5-20R"

	// PowerOutletTemplateTypeLabelNEMA5Dash30R captures enum value "NEMA 5-30R"
	PowerOutletTemplateTypeLabelNEMA5Dash30R string = "NEMA 5-30R"

	// PowerOutletTemplateTypeLabelNEMA5Dash50R captures enum value "NEMA 5-50R"
	PowerOutletTemplateTypeLabelNEMA5Dash50R string = "NEMA 5-50R"

	// PowerOutletTemplateTypeLabelNEMA6Dash15R captures enum value "NEMA 6-15R"
	PowerOutletTemplateTypeLabelNEMA6Dash15R string = "NEMA 6-15R"

	// PowerOutletTemplateTypeLabelNEMA6Dash20R captures enum value "NEMA 6-20R"
	PowerOutletTemplateTypeLabelNEMA6Dash20R string = "NEMA 6-20R"

	// PowerOutletTemplateTypeLabelNEMA6Dash30R captures enum value "NEMA 6-30R"
	PowerOutletTemplateTypeLabelNEMA6Dash30R string = "NEMA 6-30R"

	// PowerOutletTemplateTypeLabelNEMA6Dash50R captures enum value "NEMA 6-50R"
	PowerOutletTemplateTypeLabelNEMA6Dash50R string = "NEMA 6-50R"

	// PowerOutletTemplateTypeLabelNEMA10Dash30R captures enum value "NEMA 10-30R"
	PowerOutletTemplateTypeLabelNEMA10Dash30R string = "NEMA 10-30R"

	// PowerOutletTemplateTypeLabelNEMA10Dash50R captures enum value "NEMA 10-50R"
	PowerOutletTemplateTypeLabelNEMA10Dash50R string = "NEMA 10-50R"

	// PowerOutletTemplateTypeLabelNEMA14Dash20R captures enum value "NEMA 14-20R"
	PowerOutletTemplateTypeLabelNEMA14Dash20R string = "NEMA 14-20R"

	// PowerOutletTemplateTypeLabelNEMA14Dash30R captures enum value "NEMA 14-30R"
	PowerOutletTemplateTypeLabelNEMA14Dash30R string = "NEMA 14-30R"

	// PowerOutletTemplateTypeLabelNEMA14Dash50R captures enum value "NEMA 14-50R"
	PowerOutletTemplateTypeLabelNEMA14Dash50R string = "NEMA 14-50R"

	// PowerOutletTemplateTypeLabelNEMA14Dash60R captures enum value "NEMA 14-60R"
	PowerOutletTemplateTypeLabelNEMA14Dash60R string = "NEMA 14-60R"

	// PowerOutletTemplateTypeLabelNEMA15Dash15R captures enum value "NEMA 15-15R"
	PowerOutletTemplateTypeLabelNEMA15Dash15R string = "NEMA 15-15R"

	// PowerOutletTemplateTypeLabelNEMA15Dash20R captures enum value "NEMA 15-20R"
	PowerOutletTemplateTypeLabelNEMA15Dash20R string = "NEMA 15-20R"

	// PowerOutletTemplateTypeLabelNEMA15Dash30R captures enum value "NEMA 15-30R"
	PowerOutletTemplateTypeLabelNEMA15Dash30R string = "NEMA 15-30R"

	// PowerOutletTemplateTypeLabelNEMA15Dash50R captures enum value "NEMA 15-50R"
	PowerOutletTemplateTypeLabelNEMA15Dash50R string = "NEMA 15-50R"

	// PowerOutletTemplateTypeLabelNEMA15Dash60R captures enum value "NEMA 15-60R"
	PowerOutletTemplateTypeLabelNEMA15Dash60R string = "NEMA 15-60R"

	// PowerOutletTemplateTypeLabelNEMAL1Dash15R captures enum value "NEMA L1-15R"
	PowerOutletTemplateTypeLabelNEMAL1Dash15R string = "NEMA L1-15R"

	// PowerOutletTemplateTypeLabelNEMAL5Dash15R captures enum value "NEMA L5-15R"
	PowerOutletTemplateTypeLabelNEMAL5Dash15R string = "NEMA L5-15R"

	// PowerOutletTemplateTypeLabelNEMAL5Dash20R captures enum value "NEMA L5-20R"
	PowerOutletTemplateTypeLabelNEMAL5Dash20R string = "NEMA L5-20R"

	// PowerOutletTemplateTypeLabelNEMAL5Dash30R captures enum value "NEMA L5-30R"
	PowerOutletTemplateTypeLabelNEMAL5Dash30R string = "NEMA L5-30R"

	// PowerOutletTemplateTypeLabelNEMAL5Dash50R captures enum value "NEMA L5-50R"
	PowerOutletTemplateTypeLabelNEMAL5Dash50R string = "NEMA L5-50R"

	// PowerOutletTemplateTypeLabelNEMAL6Dash15R captures enum value "NEMA L6-15R"
	PowerOutletTemplateTypeLabelNEMAL6Dash15R string = "NEMA L6-15R"

	// PowerOutletTemplateTypeLabelNEMAL6Dash20R captures enum value "NEMA L6-20R"
	PowerOutletTemplateTypeLabelNEMAL6Dash20R string = "NEMA L6-20R"

	// PowerOutletTemplateTypeLabelNEMAL6Dash30R captures enum value "NEMA L6-30R"
	PowerOutletTemplateTypeLabelNEMAL6Dash30R string = "NEMA L6-30R"

	// PowerOutletTemplateTypeLabelNEMAL6Dash50R captures enum value "NEMA L6-50R"
	PowerOutletTemplateTypeLabelNEMAL6Dash50R string = "NEMA L6-50R"

	// PowerOutletTemplateTypeLabelNEMAL10Dash30R captures enum value "NEMA L10-30R"
	PowerOutletTemplateTypeLabelNEMAL10Dash30R string = "NEMA L10-30R"

	// PowerOutletTemplateTypeLabelNEMAL14Dash20R captures enum value "NEMA L14-20R"
	PowerOutletTemplateTypeLabelNEMAL14Dash20R string = "NEMA L14-20R"

	// PowerOutletTemplateTypeLabelNEMAL14Dash30R captures enum value "NEMA L14-30R"
	PowerOutletTemplateTypeLabelNEMAL14Dash30R string = "NEMA L14-30R"

	// PowerOutletTemplateTypeLabelNEMAL14Dash50R captures enum value "NEMA L14-50R"
	PowerOutletTemplateTypeLabelNEMAL14Dash50R string = "NEMA L14-50R"

	// PowerOutletTemplateTypeLabelNEMAL14Dash60R captures enum value "NEMA L14-60R"
	PowerOutletTemplateTypeLabelNEMAL14Dash60R string = "NEMA L14-60R"

	// PowerOutletTemplateTypeLabelNEMAL15Dash20R captures enum value "NEMA L15-20R"
	PowerOutletTemplateTypeLabelNEMAL15Dash20R string = "NEMA L15-20R"

	// PowerOutletTemplateTypeLabelNEMAL15Dash30R captures enum value "NEMA L15-30R"
	PowerOutletTemplateTypeLabelNEMAL15Dash30R string = "NEMA L15-30R"

	// PowerOutletTemplateTypeLabelNEMAL15Dash50R captures enum value "NEMA L15-50R"
	PowerOutletTemplateTypeLabelNEMAL15Dash50R string = "NEMA L15-50R"

	// PowerOutletTemplateTypeLabelNEMAL15Dash60R captures enum value "NEMA L15-60R"
	PowerOutletTemplateTypeLabelNEMAL15Dash60R string = "NEMA L15-60R"

	// PowerOutletTemplateTypeLabelNEMAL21Dash20R captures enum value "NEMA L21-20R"
	PowerOutletTemplateTypeLabelNEMAL21Dash20R string = "NEMA L21-20R"

	// PowerOutletTemplateTypeLabelNEMAL21Dash30R captures enum value "NEMA L21-30R"
	PowerOutletTemplateTypeLabelNEMAL21Dash30R string = "NEMA L21-30R"

	// PowerOutletTemplateTypeLabelCS6360C captures enum value "CS6360C"
	PowerOutletTemplateTypeLabelCS6360C string = "CS6360C"

	// PowerOutletTemplateTypeLabelCS6364C captures enum value "CS6364C"
	PowerOutletTemplateTypeLabelCS6364C string = "CS6364C"

	// PowerOutletTemplateTypeLabelCS8164C captures enum value "CS8164C"
	PowerOutletTemplateTypeLabelCS8164C string = "CS8164C"

	// PowerOutletTemplateTypeLabelCS8264C captures enum value "CS8264C"
	PowerOutletTemplateTypeLabelCS8264C string = "CS8264C"

	// PowerOutletTemplateTypeLabelCS8364C captures enum value "CS8364C"
	PowerOutletTemplateTypeLabelCS8364C string = "CS8364C"

	// PowerOutletTemplateTypeLabelCS8464C captures enum value "CS8464C"
	PowerOutletTemplateTypeLabelCS8464C string = "CS8464C"

	// PowerOutletTemplateTypeLabelITATypeECEE75 captures enum value "ITA Type E (CEE7/5)"
	PowerOutletTemplateTypeLabelITATypeECEE75 string = "ITA Type E (CEE7/5)"

	// PowerOutletTemplateTypeLabelITATypeFCEE73 captures enum value "ITA Type F (CEE7/3)"
	PowerOutletTemplateTypeLabelITATypeFCEE73 string = "ITA Type F (CEE7/3)"

	// PowerOutletTemplateTypeLabelITATypeGBS1363 captures enum value "ITA Type G (BS 1363)"
	PowerOutletTemplateTypeLabelITATypeGBS1363 string = "ITA Type G (BS 1363)"

	// PowerOutletTemplateTypeLabelITATypeH captures enum value "ITA Type H"
	PowerOutletTemplateTypeLabelITATypeH string = "ITA Type H"

	// PowerOutletTemplateTypeLabelITATypeI captures enum value "ITA Type I"
	PowerOutletTemplateTypeLabelITATypeI string = "ITA Type I"

	// PowerOutletTemplateTypeLabelITATypeJ captures enum value "ITA Type J"
	PowerOutletTemplateTypeLabelITATypeJ string = "ITA Type J"

	// PowerOutletTemplateTypeLabelITATypeK captures enum value "ITA Type K"
	PowerOutletTemplateTypeLabelITATypeK string = "ITA Type K"

	// PowerOutletTemplateTypeLabelITATypeLCEI23Dash50 captures enum value "ITA Type L (CEI 23-50)"
	PowerOutletTemplateTypeLabelITATypeLCEI23Dash50 string = "ITA Type L (CEI 23-50)"

	// PowerOutletTemplateTypeLabelITATypeMBS546 captures enum value "ITA Type M (BS 546)"
	PowerOutletTemplateTypeLabelITATypeMBS546 string = "ITA Type M (BS 546)"

	// PowerOutletTemplateTypeLabelITATypeN captures enum value "ITA Type N"
	PowerOutletTemplateTypeLabelITATypeN string = "ITA Type N"

	// PowerOutletTemplateTypeLabelITATypeO captures enum value "ITA Type O"
	PowerOutletTemplateTypeLabelITATypeO string = "ITA Type O"

	// PowerOutletTemplateTypeLabelUSBTypeA captures enum value "USB Type A"
	PowerOutletTemplateTypeLabelUSBTypeA string = "USB Type A"

	// PowerOutletTemplateTypeLabelUSBMicroB captures enum value "USB Micro B"
	PowerOutletTemplateTypeLabelUSBMicroB string = "USB Micro B"

	// PowerOutletTemplateTypeLabelUSBTypeC captures enum value "USB Type C"
	PowerOutletTemplateTypeLabelUSBTypeC string = "USB Type C"

	// PowerOutletTemplateTypeLabelHDOTCx captures enum value "HDOT Cx"
	PowerOutletTemplateTypeLabelHDOTCx string = "HDOT Cx"
)

// prop value enum
func (m *PowerOutletTemplateType) validateLabelEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, powerOutletTemplateTypeTypeLabelPropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *PowerOutletTemplateType) validateLabel(formats strfmt.Registry) error {

	if err := validate.Required("type"+"."+"label", "body", m.Label); err != nil {
		return err
	}

	// value enum
	if err := m.validateLabelEnum("type"+"."+"label", "body", *m.Label); err != nil {
		return err
	}

	return nil
}

var powerOutletTemplateTypeTypeValuePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["iec-60320-c5","iec-60320-c7","iec-60320-c13","iec-60320-c15","iec-60320-c19","iec-60309-p-n-e-4h","iec-60309-p-n-e-6h","iec-60309-p-n-e-9h","iec-60309-2p-e-4h","iec-60309-2p-e-6h","iec-60309-2p-e-9h","iec-60309-3p-e-4h","iec-60309-3p-e-6h","iec-60309-3p-e-9h","iec-60309-3p-n-e-4h","iec-60309-3p-n-e-6h","iec-60309-3p-n-e-9h","nema-1-15r","nema-5-15r","nema-5-20r","nema-5-30r","nema-5-50r","nema-6-15r","nema-6-20r","nema-6-30r","nema-6-50r","nema-10-30r","nema-10-50r","nema-14-20r","nema-14-30r","nema-14-50r","nema-14-60r","nema-15-15r","nema-15-20r","nema-15-30r","nema-15-50r","nema-15-60r","nema-l1-15r","nema-l5-15r","nema-l5-20r","nema-l5-30r","nema-l5-50r","nema-l6-15r","nema-l6-20r","nema-l6-30r","nema-l6-50r","nema-l10-30r","nema-l14-20r","nema-l14-30r","nema-l14-50r","nema-l14-60r","nema-l15-20r","nema-l15-30r","nema-l15-50r","nema-l15-60r","nema-l21-20r","nema-l21-30r","CS6360C","CS6364C","CS8164C","CS8264C","CS8364C","CS8464C","ita-e","ita-f","ita-g","ita-h","ita-i","ita-j","ita-k","ita-l","ita-m","ita-n","ita-o","usb-a","usb-micro-b","usb-c","hdot-cx"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		powerOutletTemplateTypeTypeValuePropEnum = append(powerOutletTemplateTypeTypeValuePropEnum, v)
	}
}

const (

	// PowerOutletTemplateTypeValueIecDash60320DashC5 captures enum value "iec-60320-c5"
	PowerOutletTemplateTypeValueIecDash60320DashC5 string = "iec-60320-c5"

	// PowerOutletTemplateTypeValueIecDash60320DashC7 captures enum value "iec-60320-c7"
	PowerOutletTemplateTypeValueIecDash60320DashC7 string = "iec-60320-c7"

	// PowerOutletTemplateTypeValueIecDash60320DashC13 captures enum value "iec-60320-c13"
	PowerOutletTemplateTypeValueIecDash60320DashC13 string = "iec-60320-c13"

	// PowerOutletTemplateTypeValueIecDash60320DashC15 captures enum value "iec-60320-c15"
	PowerOutletTemplateTypeValueIecDash60320DashC15 string = "iec-60320-c15"

	// PowerOutletTemplateTypeValueIecDash60320DashC19 captures enum value "iec-60320-c19"
	PowerOutletTemplateTypeValueIecDash60320DashC19 string = "iec-60320-c19"

	// PowerOutletTemplateTypeValueIecDash60309DashpDashnDasheDash4h captures enum value "iec-60309-p-n-e-4h"
	PowerOutletTemplateTypeValueIecDash60309DashpDashnDasheDash4h string = "iec-60309-p-n-e-4h"

	// PowerOutletTemplateTypeValueIecDash60309DashpDashnDasheDash6h captures enum value "iec-60309-p-n-e-6h"
	PowerOutletTemplateTypeValueIecDash60309DashpDashnDasheDash6h string = "iec-60309-p-n-e-6h"

	// PowerOutletTemplateTypeValueIecDash60309DashpDashnDasheDash9h captures enum value "iec-60309-p-n-e-9h"
	PowerOutletTemplateTypeValueIecDash60309DashpDashnDasheDash9h string = "iec-60309-p-n-e-9h"

	// PowerOutletTemplateTypeValueIecDash60309Dash2pDasheDash4h captures enum value "iec-60309-2p-e-4h"
	PowerOutletTemplateTypeValueIecDash60309Dash2pDasheDash4h string = "iec-60309-2p-e-4h"

	// PowerOutletTemplateTypeValueIecDash60309Dash2pDasheDash6h captures enum value "iec-60309-2p-e-6h"
	PowerOutletTemplateTypeValueIecDash60309Dash2pDasheDash6h string = "iec-60309-2p-e-6h"

	// PowerOutletTemplateTypeValueIecDash60309Dash2pDasheDash9h captures enum value "iec-60309-2p-e-9h"
	PowerOutletTemplateTypeValueIecDash60309Dash2pDasheDash9h string = "iec-60309-2p-e-9h"

	// PowerOutletTemplateTypeValueIecDash60309Dash3pDasheDash4h captures enum value "iec-60309-3p-e-4h"
	PowerOutletTemplateTypeValueIecDash60309Dash3pDasheDash4h string = "iec-60309-3p-e-4h"

	// PowerOutletTemplateTypeValueIecDash60309Dash3pDasheDash6h captures enum value "iec-60309-3p-e-6h"
	PowerOutletTemplateTypeValueIecDash60309Dash3pDasheDash6h string = "iec-60309-3p-e-6h"

	// PowerOutletTemplateTypeValueIecDash60309Dash3pDasheDash9h captures enum value "iec-60309-3p-e-9h"
	PowerOutletTemplateTypeValueIecDash60309Dash3pDasheDash9h string = "iec-60309-3p-e-9h"

	// PowerOutletTemplateTypeValueIecDash60309Dash3pDashnDasheDash4h captures enum value "iec-60309-3p-n-e-4h"
	PowerOutletTemplateTypeValueIecDash60309Dash3pDashnDasheDash4h string = "iec-60309-3p-n-e-4h"

	// PowerOutletTemplateTypeValueIecDash60309Dash3pDashnDasheDash6h captures enum value "iec-60309-3p-n-e-6h"
	PowerOutletTemplateTypeValueIecDash60309Dash3pDashnDasheDash6h string = "iec-60309-3p-n-e-6h"

	// PowerOutletTemplateTypeValueIecDash60309Dash3pDashnDasheDash9h captures enum value "iec-60309-3p-n-e-9h"
	PowerOutletTemplateTypeValueIecDash60309Dash3pDashnDasheDash9h string = "iec-60309-3p-n-e-9h"

	// PowerOutletTemplateTypeValueNemaDash1Dash15r captures enum value "nema-1-15r"
	PowerOutletTemplateTypeValueNemaDash1Dash15r string = "nema-1-15r"

	// PowerOutletTemplateTypeValueNemaDash5Dash15r captures enum value "nema-5-15r"
	PowerOutletTemplateTypeValueNemaDash5Dash15r string = "nema-5-15r"

	// PowerOutletTemplateTypeValueNemaDash5Dash20r captures enum value "nema-5-20r"
	PowerOutletTemplateTypeValueNemaDash5Dash20r string = "nema-5-20r"

	// PowerOutletTemplateTypeValueNemaDash5Dash30r captures enum value "nema-5-30r"
	PowerOutletTemplateTypeValueNemaDash5Dash30r string = "nema-5-30r"

	// PowerOutletTemplateTypeValueNemaDash5Dash50r captures enum value "nema-5-50r"
	PowerOutletTemplateTypeValueNemaDash5Dash50r string = "nema-5-50r"

	// PowerOutletTemplateTypeValueNemaDash6Dash15r captures enum value "nema-6-15r"
	PowerOutletTemplateTypeValueNemaDash6Dash15r string = "nema-6-15r"

	// PowerOutletTemplateTypeValueNemaDash6Dash20r captures enum value "nema-6-20r"
	PowerOutletTemplateTypeValueNemaDash6Dash20r string = "nema-6-20r"

	// PowerOutletTemplateTypeValueNemaDash6Dash30r captures enum value "nema-6-30r"
	PowerOutletTemplateTypeValueNemaDash6Dash30r string = "nema-6-30r"

	// PowerOutletTemplateTypeValueNemaDash6Dash50r captures enum value "nema-6-50r"
	PowerOutletTemplateTypeValueNemaDash6Dash50r string = "nema-6-50r"

	// PowerOutletTemplateTypeValueNemaDash10Dash30r captures enum value "nema-10-30r"
	PowerOutletTemplateTypeValueNemaDash10Dash30r string = "nema-10-30r"

	// PowerOutletTemplateTypeValueNemaDash10Dash50r captures enum value "nema-10-50r"
	PowerOutletTemplateTypeValueNemaDash10Dash50r string = "nema-10-50r"

	// PowerOutletTemplateTypeValueNemaDash14Dash20r captures enum value "nema-14-20r"
	PowerOutletTemplateTypeValueNemaDash14Dash20r string = "nema-14-20r"

	// PowerOutletTemplateTypeValueNemaDash14Dash30r captures enum value "nema-14-30r"
	PowerOutletTemplateTypeValueNemaDash14Dash30r string = "nema-14-30r"

	// PowerOutletTemplateTypeValueNemaDash14Dash50r captures enum value "nema-14-50r"
	PowerOutletTemplateTypeValueNemaDash14Dash50r string = "nema-14-50r"

	// PowerOutletTemplateTypeValueNemaDash14Dash60r captures enum value "nema-14-60r"
	PowerOutletTemplateTypeValueNemaDash14Dash60r string = "nema-14-60r"

	// PowerOutletTemplateTypeValueNemaDash15Dash15r captures enum value "nema-15-15r"
	PowerOutletTemplateTypeValueNemaDash15Dash15r string = "nema-15-15r"

	// PowerOutletTemplateTypeValueNemaDash15Dash20r captures enum value "nema-15-20r"
	PowerOutletTemplateTypeValueNemaDash15Dash20r string = "nema-15-20r"

	// PowerOutletTemplateTypeValueNemaDash15Dash30r captures enum value "nema-15-30r"
	PowerOutletTemplateTypeValueNemaDash15Dash30r string = "nema-15-30r"

	// PowerOutletTemplateTypeValueNemaDash15Dash50r captures enum value "nema-15-50r"
	PowerOutletTemplateTypeValueNemaDash15Dash50r string = "nema-15-50r"

	// PowerOutletTemplateTypeValueNemaDash15Dash60r captures enum value "nema-15-60r"
	PowerOutletTemplateTypeValueNemaDash15Dash60r string = "nema-15-60r"

	// PowerOutletTemplateTypeValueNemaDashL1Dash15r captures enum value "nema-l1-15r"
	PowerOutletTemplateTypeValueNemaDashL1Dash15r string = "nema-l1-15r"

	// PowerOutletTemplateTypeValueNemaDashL5Dash15r captures enum value "nema-l5-15r"
	PowerOutletTemplateTypeValueNemaDashL5Dash15r string = "nema-l5-15r"

	// PowerOutletTemplateTypeValueNemaDashL5Dash20r captures enum value "nema-l5-20r"
	PowerOutletTemplateTypeValueNemaDashL5Dash20r string = "nema-l5-20r"

	// PowerOutletTemplateTypeValueNemaDashL5Dash30r captures enum value "nema-l5-30r"
	PowerOutletTemplateTypeValueNemaDashL5Dash30r string = "nema-l5-30r"

	// PowerOutletTemplateTypeValueNemaDashL5Dash50r captures enum value "nema-l5-50r"
	PowerOutletTemplateTypeValueNemaDashL5Dash50r string = "nema-l5-50r"

	// PowerOutletTemplateTypeValueNemaDashL6Dash15r captures enum value "nema-l6-15r"
	PowerOutletTemplateTypeValueNemaDashL6Dash15r string = "nema-l6-15r"

	// PowerOutletTemplateTypeValueNemaDashL6Dash20r captures enum value "nema-l6-20r"
	PowerOutletTemplateTypeValueNemaDashL6Dash20r string = "nema-l6-20r"

	// PowerOutletTemplateTypeValueNemaDashL6Dash30r captures enum value "nema-l6-30r"
	PowerOutletTemplateTypeValueNemaDashL6Dash30r string = "nema-l6-30r"

	// PowerOutletTemplateTypeValueNemaDashL6Dash50r captures enum value "nema-l6-50r"
	PowerOutletTemplateTypeValueNemaDashL6Dash50r string = "nema-l6-50r"

	// PowerOutletTemplateTypeValueNemaDashL10Dash30r captures enum value "nema-l10-30r"
	PowerOutletTemplateTypeValueNemaDashL10Dash30r string = "nema-l10-30r"

	// PowerOutletTemplateTypeValueNemaDashL14Dash20r captures enum value "nema-l14-20r"
	PowerOutletTemplateTypeValueNemaDashL14Dash20r string = "nema-l14-20r"

	// PowerOutletTemplateTypeValueNemaDashL14Dash30r captures enum value "nema-l14-30r"
	PowerOutletTemplateTypeValueNemaDashL14Dash30r string = "nema-l14-30r"

	// PowerOutletTemplateTypeValueNemaDashL14Dash50r captures enum value "nema-l14-50r"
	PowerOutletTemplateTypeValueNemaDashL14Dash50r string = "nema-l14-50r"

	// PowerOutletTemplateTypeValueNemaDashL14Dash60r captures enum value "nema-l14-60r"
	PowerOutletTemplateTypeValueNemaDashL14Dash60r string = "nema-l14-60r"

	// PowerOutletTemplateTypeValueNemaDashL15Dash20r captures enum value "nema-l15-20r"
	PowerOutletTemplateTypeValueNemaDashL15Dash20r string = "nema-l15-20r"

	// PowerOutletTemplateTypeValueNemaDashL15Dash30r captures enum value "nema-l15-30r"
	PowerOutletTemplateTypeValueNemaDashL15Dash30r string = "nema-l15-30r"

	// PowerOutletTemplateTypeValueNemaDashL15Dash50r captures enum value "nema-l15-50r"
	PowerOutletTemplateTypeValueNemaDashL15Dash50r string = "nema-l15-50r"

	// PowerOutletTemplateTypeValueNemaDashL15Dash60r captures enum value "nema-l15-60r"
	PowerOutletTemplateTypeValueNemaDashL15Dash60r string = "nema-l15-60r"

	// PowerOutletTemplateTypeValueNemaDashL21Dash20r captures enum value "nema-l21-20r"
	PowerOutletTemplateTypeValueNemaDashL21Dash20r string = "nema-l21-20r"

	// PowerOutletTemplateTypeValueNemaDashL21Dash30r captures enum value "nema-l21-30r"
	PowerOutletTemplateTypeValueNemaDashL21Dash30r string = "nema-l21-30r"

	// PowerOutletTemplateTypeValueCS6360C captures enum value "CS6360C"
	PowerOutletTemplateTypeValueCS6360C string = "CS6360C"

	// PowerOutletTemplateTypeValueCS6364C captures enum value "CS6364C"
	PowerOutletTemplateTypeValueCS6364C string = "CS6364C"

	// PowerOutletTemplateTypeValueCS8164C captures enum value "CS8164C"
	PowerOutletTemplateTypeValueCS8164C string = "CS8164C"

	// PowerOutletTemplateTypeValueCS8264C captures enum value "CS8264C"
	PowerOutletTemplateTypeValueCS8264C string = "CS8264C"

	// PowerOutletTemplateTypeValueCS8364C captures enum value "CS8364C"
	PowerOutletTemplateTypeValueCS8364C string = "CS8364C"

	// PowerOutletTemplateTypeValueCS8464C captures enum value "CS8464C"
	PowerOutletTemplateTypeValueCS8464C string = "CS8464C"

	// PowerOutletTemplateTypeValueItaDashe captures enum value "ita-e"
	PowerOutletTemplateTypeValueItaDashe string = "ita-e"

	// PowerOutletTemplateTypeValueItaDashf captures enum value "ita-f"
	PowerOutletTemplateTypeValueItaDashf string = "ita-f"

	// PowerOutletTemplateTypeValueItaDashg captures enum value "ita-g"
	PowerOutletTemplateTypeValueItaDashg string = "ita-g"

	// PowerOutletTemplateTypeValueItaDashh captures enum value "ita-h"
	PowerOutletTemplateTypeValueItaDashh string = "ita-h"

	// PowerOutletTemplateTypeValueItaDashi captures enum value "ita-i"
	PowerOutletTemplateTypeValueItaDashi string = "ita-i"

	// PowerOutletTemplateTypeValueItaDashj captures enum value "ita-j"
	PowerOutletTemplateTypeValueItaDashj string = "ita-j"

	// PowerOutletTemplateTypeValueItaDashk captures enum value "ita-k"
	PowerOutletTemplateTypeValueItaDashk string = "ita-k"

	// PowerOutletTemplateTypeValueItaDashl captures enum value "ita-l"
	PowerOutletTemplateTypeValueItaDashl string = "ita-l"

	// PowerOutletTemplateTypeValueItaDashm captures enum value "ita-m"
	PowerOutletTemplateTypeValueItaDashm string = "ita-m"

	// PowerOutletTemplateTypeValueItaDashn captures enum value "ita-n"
	PowerOutletTemplateTypeValueItaDashn string = "ita-n"

	// PowerOutletTemplateTypeValueItaDasho captures enum value "ita-o"
	PowerOutletTemplateTypeValueItaDasho string = "ita-o"

	// PowerOutletTemplateTypeValueUsbDasha captures enum value "usb-a"
	PowerOutletTemplateTypeValueUsbDasha string = "usb-a"

	// PowerOutletTemplateTypeValueUsbDashMicroDashb captures enum value "usb-micro-b"
	PowerOutletTemplateTypeValueUsbDashMicroDashb string = "usb-micro-b"

	// PowerOutletTemplateTypeValueUsbDashc captures enum value "usb-c"
	PowerOutletTemplateTypeValueUsbDashc string = "usb-c"

	// PowerOutletTemplateTypeValueHdotDashCx captures enum value "hdot-cx"
	PowerOutletTemplateTypeValueHdotDashCx string = "hdot-cx"
)

// prop value enum
func (m *PowerOutletTemplateType) validateValueEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, powerOutletTemplateTypeTypeValuePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *PowerOutletTemplateType) validateValue(formats strfmt.Registry) error {

	if err := validate.Required("type"+"."+"value", "body", m.Value); err != nil {
		return err
	}

	// value enum
	if err := m.validateValueEnum("type"+"."+"value", "body", *m.Value); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this power outlet template type based on context it is used
func (m *PowerOutletTemplateType) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *PowerOutletTemplateType) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *PowerOutletTemplateType) UnmarshalBinary(b []byte) error {
	var res PowerOutletTemplateType
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
