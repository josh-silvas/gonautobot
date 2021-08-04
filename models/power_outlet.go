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

// PowerOutlet power outlet
//
// swagger:model PowerOutlet
type PowerOutlet struct {

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

	// Connected endpoint
	//
	//
	// Return the appropriate serializer for the type of connected object.
	//
	// Read Only: true
	ConnectedEndpoint map[string]*string `json:"connected_endpoint,omitempty"`

	// Connected endpoint reachable
	// Read Only: true
	ConnectedEndpointReachable *bool `json:"connected_endpoint_reachable,omitempty"`

	// Connected endpoint type
	// Read Only: true
	ConnectedEndpointType string `json:"connected_endpoint_type,omitempty"`

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

	// feed leg
	FeedLeg *PowerOutletFeedLeg `json:"feed_leg,omitempty"`

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
	PowerPort *NestedPowerPort `json:"power_port,omitempty"`

	// tags
	Tags []*NestedTag `json:"tags"`

	// type
	Type *PowerOutletType `json:"type,omitempty"`

	// Url
	// Read Only: true
	// Format: uri
	URL strfmt.URI `json:"url,omitempty"`
}

// Validate validates this power outlet
func (m *PowerOutlet) Validate(formats strfmt.Registry) error {
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

func (m *PowerOutlet) validateCable(formats strfmt.Registry) error {
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

func (m *PowerOutlet) validateDescription(formats strfmt.Registry) error {
	if swag.IsZero(m.Description) { // not required
		return nil
	}

	if err := validate.MaxLength("description", "body", m.Description, 200); err != nil {
		return err
	}

	return nil
}

func (m *PowerOutlet) validateDevice(formats strfmt.Registry) error {

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

func (m *PowerOutlet) validateDisplay(formats strfmt.Registry) error {
	if swag.IsZero(m.Display) { // not required
		return nil
	}

	if err := validate.MinLength("display", "body", m.Display, 1); err != nil {
		return err
	}

	return nil
}

func (m *PowerOutlet) validateFeedLeg(formats strfmt.Registry) error {
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

func (m *PowerOutlet) validateID(formats strfmt.Registry) error {
	if swag.IsZero(m.ID) { // not required
		return nil
	}

	if err := validate.FormatOf("id", "body", "uuid", m.ID.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *PowerOutlet) validateLabel(formats strfmt.Registry) error {
	if swag.IsZero(m.Label) { // not required
		return nil
	}

	if err := validate.MaxLength("label", "body", m.Label, 64); err != nil {
		return err
	}

	return nil
}

func (m *PowerOutlet) validateName(formats strfmt.Registry) error {

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

func (m *PowerOutlet) validatePowerPort(formats strfmt.Registry) error {
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

func (m *PowerOutlet) validateTags(formats strfmt.Registry) error {
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

func (m *PowerOutlet) validateType(formats strfmt.Registry) error {
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

func (m *PowerOutlet) validateURL(formats strfmt.Registry) error {
	if swag.IsZero(m.URL) { // not required
		return nil
	}

	if err := validate.FormatOf("url", "body", "uri", m.URL.String(), formats); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this power outlet based on the context it is used
func (m *PowerOutlet) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
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

	if err := m.contextValidateConnectedEndpoint(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateConnectedEndpointReachable(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateConnectedEndpointType(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateDevice(ctx, formats); err != nil {
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

func (m *PowerOutlet) contextValidateCable(ctx context.Context, formats strfmt.Registry) error {

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

func (m *PowerOutlet) contextValidateCablePeer(ctx context.Context, formats strfmt.Registry) error {

	return nil
}

func (m *PowerOutlet) contextValidateCablePeerType(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "cable_peer_type", "body", string(m.CablePeerType)); err != nil {
		return err
	}

	return nil
}

func (m *PowerOutlet) contextValidateConnectedEndpoint(ctx context.Context, formats strfmt.Registry) error {

	return nil
}

func (m *PowerOutlet) contextValidateConnectedEndpointReachable(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "connected_endpoint_reachable", "body", m.ConnectedEndpointReachable); err != nil {
		return err
	}

	return nil
}

func (m *PowerOutlet) contextValidateConnectedEndpointType(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "connected_endpoint_type", "body", string(m.ConnectedEndpointType)); err != nil {
		return err
	}

	return nil
}

func (m *PowerOutlet) contextValidateDevice(ctx context.Context, formats strfmt.Registry) error {

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

func (m *PowerOutlet) contextValidateDisplay(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "display", "body", string(m.Display)); err != nil {
		return err
	}

	return nil
}

func (m *PowerOutlet) contextValidateFeedLeg(ctx context.Context, formats strfmt.Registry) error {

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

func (m *PowerOutlet) contextValidateID(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "id", "body", strfmt.UUID(m.ID)); err != nil {
		return err
	}

	return nil
}

func (m *PowerOutlet) contextValidatePowerPort(ctx context.Context, formats strfmt.Registry) error {

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

func (m *PowerOutlet) contextValidateTags(ctx context.Context, formats strfmt.Registry) error {

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

func (m *PowerOutlet) contextValidateType(ctx context.Context, formats strfmt.Registry) error {

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

func (m *PowerOutlet) contextValidateURL(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "url", "body", strfmt.URI(m.URL)); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *PowerOutlet) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *PowerOutlet) UnmarshalBinary(b []byte) error {
	var res PowerOutlet
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// PowerOutletFeedLeg Feed leg
//
// swagger:model PowerOutletFeedLeg
type PowerOutletFeedLeg struct {

	// label
	// Required: true
	// Enum: [A B C]
	Label *string `json:"label"`

	// value
	// Required: true
	// Enum: [A B C]
	Value *string `json:"value"`
}

// Validate validates this power outlet feed leg
func (m *PowerOutletFeedLeg) Validate(formats strfmt.Registry) error {
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

var powerOutletFeedLegTypeLabelPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["A","B","C"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		powerOutletFeedLegTypeLabelPropEnum = append(powerOutletFeedLegTypeLabelPropEnum, v)
	}
}

const (

	// PowerOutletFeedLegLabelA captures enum value "A"
	PowerOutletFeedLegLabelA string = "A"

	// PowerOutletFeedLegLabelB captures enum value "B"
	PowerOutletFeedLegLabelB string = "B"

	// PowerOutletFeedLegLabelC captures enum value "C"
	PowerOutletFeedLegLabelC string = "C"
)

// prop value enum
func (m *PowerOutletFeedLeg) validateLabelEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, powerOutletFeedLegTypeLabelPropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *PowerOutletFeedLeg) validateLabel(formats strfmt.Registry) error {

	if err := validate.Required("feed_leg"+"."+"label", "body", m.Label); err != nil {
		return err
	}

	// value enum
	if err := m.validateLabelEnum("feed_leg"+"."+"label", "body", *m.Label); err != nil {
		return err
	}

	return nil
}

var powerOutletFeedLegTypeValuePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["A","B","C"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		powerOutletFeedLegTypeValuePropEnum = append(powerOutletFeedLegTypeValuePropEnum, v)
	}
}

const (

	// PowerOutletFeedLegValueA captures enum value "A"
	PowerOutletFeedLegValueA string = "A"

	// PowerOutletFeedLegValueB captures enum value "B"
	PowerOutletFeedLegValueB string = "B"

	// PowerOutletFeedLegValueC captures enum value "C"
	PowerOutletFeedLegValueC string = "C"
)

// prop value enum
func (m *PowerOutletFeedLeg) validateValueEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, powerOutletFeedLegTypeValuePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *PowerOutletFeedLeg) validateValue(formats strfmt.Registry) error {

	if err := validate.Required("feed_leg"+"."+"value", "body", m.Value); err != nil {
		return err
	}

	// value enum
	if err := m.validateValueEnum("feed_leg"+"."+"value", "body", *m.Value); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this power outlet feed leg based on context it is used
func (m *PowerOutletFeedLeg) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *PowerOutletFeedLeg) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *PowerOutletFeedLeg) UnmarshalBinary(b []byte) error {
	var res PowerOutletFeedLeg
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// PowerOutletType Type
//
// swagger:model PowerOutletType
type PowerOutletType struct {

	// label
	// Required: true
	// Enum: [C5 C7 C13 C15 C19 P+N+E 4H P+N+E 6H P+N+E 9H 2P+E 4H 2P+E 6H 2P+E 9H 3P+E 4H 3P+E 6H 3P+E 9H 3P+N+E 4H 3P+N+E 6H 3P+N+E 9H NEMA 1-15R NEMA 5-15R NEMA 5-20R NEMA 5-30R NEMA 5-50R NEMA 6-15R NEMA 6-20R NEMA 6-30R NEMA 6-50R NEMA 10-30R NEMA 10-50R NEMA 14-20R NEMA 14-30R NEMA 14-50R NEMA 14-60R NEMA 15-15R NEMA 15-20R NEMA 15-30R NEMA 15-50R NEMA 15-60R NEMA L1-15R NEMA L5-15R NEMA L5-20R NEMA L5-30R NEMA L5-50R NEMA L6-15R NEMA L6-20R NEMA L6-30R NEMA L6-50R NEMA L10-30R NEMA L14-20R NEMA L14-30R NEMA L14-50R NEMA L14-60R NEMA L15-20R NEMA L15-30R NEMA L15-50R NEMA L15-60R NEMA L21-20R NEMA L21-30R CS6360C CS6364C CS8164C CS8264C CS8364C CS8464C ITA Type E (CEE7/5) ITA Type F (CEE7/3) ITA Type G (BS 1363) ITA Type H ITA Type I ITA Type J ITA Type K ITA Type L (CEI 23-50) ITA Type M (BS 546) ITA Type N ITA Type O USB Type A USB Micro B USB Type C HDOT Cx]
	Label *string `json:"label"`

	// value
	// Required: true
	// Enum: [iec-60320-c5 iec-60320-c7 iec-60320-c13 iec-60320-c15 iec-60320-c19 iec-60309-p-n-e-4h iec-60309-p-n-e-6h iec-60309-p-n-e-9h iec-60309-2p-e-4h iec-60309-2p-e-6h iec-60309-2p-e-9h iec-60309-3p-e-4h iec-60309-3p-e-6h iec-60309-3p-e-9h iec-60309-3p-n-e-4h iec-60309-3p-n-e-6h iec-60309-3p-n-e-9h nema-1-15r nema-5-15r nema-5-20r nema-5-30r nema-5-50r nema-6-15r nema-6-20r nema-6-30r nema-6-50r nema-10-30r nema-10-50r nema-14-20r nema-14-30r nema-14-50r nema-14-60r nema-15-15r nema-15-20r nema-15-30r nema-15-50r nema-15-60r nema-l1-15r nema-l5-15r nema-l5-20r nema-l5-30r nema-l5-50r nema-l6-15r nema-l6-20r nema-l6-30r nema-l6-50r nema-l10-30r nema-l14-20r nema-l14-30r nema-l14-50r nema-l14-60r nema-l15-20r nema-l15-30r nema-l15-50r nema-l15-60r nema-l21-20r nema-l21-30r CS6360C CS6364C CS8164C CS8264C CS8364C CS8464C ita-e ita-f ita-g ita-h ita-i ita-j ita-k ita-l ita-m ita-n ita-o usb-a usb-micro-b usb-c hdot-cx]
	Value *string `json:"value"`
}

// Validate validates this power outlet type
func (m *PowerOutletType) Validate(formats strfmt.Registry) error {
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

var powerOutletTypeTypeLabelPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["C5","C7","C13","C15","C19","P+N+E 4H","P+N+E 6H","P+N+E 9H","2P+E 4H","2P+E 6H","2P+E 9H","3P+E 4H","3P+E 6H","3P+E 9H","3P+N+E 4H","3P+N+E 6H","3P+N+E 9H","NEMA 1-15R","NEMA 5-15R","NEMA 5-20R","NEMA 5-30R","NEMA 5-50R","NEMA 6-15R","NEMA 6-20R","NEMA 6-30R","NEMA 6-50R","NEMA 10-30R","NEMA 10-50R","NEMA 14-20R","NEMA 14-30R","NEMA 14-50R","NEMA 14-60R","NEMA 15-15R","NEMA 15-20R","NEMA 15-30R","NEMA 15-50R","NEMA 15-60R","NEMA L1-15R","NEMA L5-15R","NEMA L5-20R","NEMA L5-30R","NEMA L5-50R","NEMA L6-15R","NEMA L6-20R","NEMA L6-30R","NEMA L6-50R","NEMA L10-30R","NEMA L14-20R","NEMA L14-30R","NEMA L14-50R","NEMA L14-60R","NEMA L15-20R","NEMA L15-30R","NEMA L15-50R","NEMA L15-60R","NEMA L21-20R","NEMA L21-30R","CS6360C","CS6364C","CS8164C","CS8264C","CS8364C","CS8464C","ITA Type E (CEE7/5)","ITA Type F (CEE7/3)","ITA Type G (BS 1363)","ITA Type H","ITA Type I","ITA Type J","ITA Type K","ITA Type L (CEI 23-50)","ITA Type M (BS 546)","ITA Type N","ITA Type O","USB Type A","USB Micro B","USB Type C","HDOT Cx"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		powerOutletTypeTypeLabelPropEnum = append(powerOutletTypeTypeLabelPropEnum, v)
	}
}

const (

	// PowerOutletTypeLabelC5 captures enum value "C5"
	PowerOutletTypeLabelC5 string = "C5"

	// PowerOutletTypeLabelC7 captures enum value "C7"
	PowerOutletTypeLabelC7 string = "C7"

	// PowerOutletTypeLabelC13 captures enum value "C13"
	PowerOutletTypeLabelC13 string = "C13"

	// PowerOutletTypeLabelC15 captures enum value "C15"
	PowerOutletTypeLabelC15 string = "C15"

	// PowerOutletTypeLabelC19 captures enum value "C19"
	PowerOutletTypeLabelC19 string = "C19"

	// PowerOutletTypeLabelPPlusNPlusE4H captures enum value "P+N+E 4H"
	PowerOutletTypeLabelPPlusNPlusE4H string = "P+N+E 4H"

	// PowerOutletTypeLabelPPlusNPlusE6H captures enum value "P+N+E 6H"
	PowerOutletTypeLabelPPlusNPlusE6H string = "P+N+E 6H"

	// PowerOutletTypeLabelPPlusNPlusE9H captures enum value "P+N+E 9H"
	PowerOutletTypeLabelPPlusNPlusE9H string = "P+N+E 9H"

	// PowerOutletTypeLabelNr2PPlusE4H captures enum value "2P+E 4H"
	PowerOutletTypeLabelNr2PPlusE4H string = "2P+E 4H"

	// PowerOutletTypeLabelNr2PPlusE6H captures enum value "2P+E 6H"
	PowerOutletTypeLabelNr2PPlusE6H string = "2P+E 6H"

	// PowerOutletTypeLabelNr2PPlusE9H captures enum value "2P+E 9H"
	PowerOutletTypeLabelNr2PPlusE9H string = "2P+E 9H"

	// PowerOutletTypeLabelNr3PPlusE4H captures enum value "3P+E 4H"
	PowerOutletTypeLabelNr3PPlusE4H string = "3P+E 4H"

	// PowerOutletTypeLabelNr3PPlusE6H captures enum value "3P+E 6H"
	PowerOutletTypeLabelNr3PPlusE6H string = "3P+E 6H"

	// PowerOutletTypeLabelNr3PPlusE9H captures enum value "3P+E 9H"
	PowerOutletTypeLabelNr3PPlusE9H string = "3P+E 9H"

	// PowerOutletTypeLabelNr3PPlusNPlusE4H captures enum value "3P+N+E 4H"
	PowerOutletTypeLabelNr3PPlusNPlusE4H string = "3P+N+E 4H"

	// PowerOutletTypeLabelNr3PPlusNPlusE6H captures enum value "3P+N+E 6H"
	PowerOutletTypeLabelNr3PPlusNPlusE6H string = "3P+N+E 6H"

	// PowerOutletTypeLabelNr3PPlusNPlusE9H captures enum value "3P+N+E 9H"
	PowerOutletTypeLabelNr3PPlusNPlusE9H string = "3P+N+E 9H"

	// PowerOutletTypeLabelNEMA1Dash15R captures enum value "NEMA 1-15R"
	PowerOutletTypeLabelNEMA1Dash15R string = "NEMA 1-15R"

	// PowerOutletTypeLabelNEMA5Dash15R captures enum value "NEMA 5-15R"
	PowerOutletTypeLabelNEMA5Dash15R string = "NEMA 5-15R"

	// PowerOutletTypeLabelNEMA5Dash20R captures enum value "NEMA 5-20R"
	PowerOutletTypeLabelNEMA5Dash20R string = "NEMA 5-20R"

	// PowerOutletTypeLabelNEMA5Dash30R captures enum value "NEMA 5-30R"
	PowerOutletTypeLabelNEMA5Dash30R string = "NEMA 5-30R"

	// PowerOutletTypeLabelNEMA5Dash50R captures enum value "NEMA 5-50R"
	PowerOutletTypeLabelNEMA5Dash50R string = "NEMA 5-50R"

	// PowerOutletTypeLabelNEMA6Dash15R captures enum value "NEMA 6-15R"
	PowerOutletTypeLabelNEMA6Dash15R string = "NEMA 6-15R"

	// PowerOutletTypeLabelNEMA6Dash20R captures enum value "NEMA 6-20R"
	PowerOutletTypeLabelNEMA6Dash20R string = "NEMA 6-20R"

	// PowerOutletTypeLabelNEMA6Dash30R captures enum value "NEMA 6-30R"
	PowerOutletTypeLabelNEMA6Dash30R string = "NEMA 6-30R"

	// PowerOutletTypeLabelNEMA6Dash50R captures enum value "NEMA 6-50R"
	PowerOutletTypeLabelNEMA6Dash50R string = "NEMA 6-50R"

	// PowerOutletTypeLabelNEMA10Dash30R captures enum value "NEMA 10-30R"
	PowerOutletTypeLabelNEMA10Dash30R string = "NEMA 10-30R"

	// PowerOutletTypeLabelNEMA10Dash50R captures enum value "NEMA 10-50R"
	PowerOutletTypeLabelNEMA10Dash50R string = "NEMA 10-50R"

	// PowerOutletTypeLabelNEMA14Dash20R captures enum value "NEMA 14-20R"
	PowerOutletTypeLabelNEMA14Dash20R string = "NEMA 14-20R"

	// PowerOutletTypeLabelNEMA14Dash30R captures enum value "NEMA 14-30R"
	PowerOutletTypeLabelNEMA14Dash30R string = "NEMA 14-30R"

	// PowerOutletTypeLabelNEMA14Dash50R captures enum value "NEMA 14-50R"
	PowerOutletTypeLabelNEMA14Dash50R string = "NEMA 14-50R"

	// PowerOutletTypeLabelNEMA14Dash60R captures enum value "NEMA 14-60R"
	PowerOutletTypeLabelNEMA14Dash60R string = "NEMA 14-60R"

	// PowerOutletTypeLabelNEMA15Dash15R captures enum value "NEMA 15-15R"
	PowerOutletTypeLabelNEMA15Dash15R string = "NEMA 15-15R"

	// PowerOutletTypeLabelNEMA15Dash20R captures enum value "NEMA 15-20R"
	PowerOutletTypeLabelNEMA15Dash20R string = "NEMA 15-20R"

	// PowerOutletTypeLabelNEMA15Dash30R captures enum value "NEMA 15-30R"
	PowerOutletTypeLabelNEMA15Dash30R string = "NEMA 15-30R"

	// PowerOutletTypeLabelNEMA15Dash50R captures enum value "NEMA 15-50R"
	PowerOutletTypeLabelNEMA15Dash50R string = "NEMA 15-50R"

	// PowerOutletTypeLabelNEMA15Dash60R captures enum value "NEMA 15-60R"
	PowerOutletTypeLabelNEMA15Dash60R string = "NEMA 15-60R"

	// PowerOutletTypeLabelNEMAL1Dash15R captures enum value "NEMA L1-15R"
	PowerOutletTypeLabelNEMAL1Dash15R string = "NEMA L1-15R"

	// PowerOutletTypeLabelNEMAL5Dash15R captures enum value "NEMA L5-15R"
	PowerOutletTypeLabelNEMAL5Dash15R string = "NEMA L5-15R"

	// PowerOutletTypeLabelNEMAL5Dash20R captures enum value "NEMA L5-20R"
	PowerOutletTypeLabelNEMAL5Dash20R string = "NEMA L5-20R"

	// PowerOutletTypeLabelNEMAL5Dash30R captures enum value "NEMA L5-30R"
	PowerOutletTypeLabelNEMAL5Dash30R string = "NEMA L5-30R"

	// PowerOutletTypeLabelNEMAL5Dash50R captures enum value "NEMA L5-50R"
	PowerOutletTypeLabelNEMAL5Dash50R string = "NEMA L5-50R"

	// PowerOutletTypeLabelNEMAL6Dash15R captures enum value "NEMA L6-15R"
	PowerOutletTypeLabelNEMAL6Dash15R string = "NEMA L6-15R"

	// PowerOutletTypeLabelNEMAL6Dash20R captures enum value "NEMA L6-20R"
	PowerOutletTypeLabelNEMAL6Dash20R string = "NEMA L6-20R"

	// PowerOutletTypeLabelNEMAL6Dash30R captures enum value "NEMA L6-30R"
	PowerOutletTypeLabelNEMAL6Dash30R string = "NEMA L6-30R"

	// PowerOutletTypeLabelNEMAL6Dash50R captures enum value "NEMA L6-50R"
	PowerOutletTypeLabelNEMAL6Dash50R string = "NEMA L6-50R"

	// PowerOutletTypeLabelNEMAL10Dash30R captures enum value "NEMA L10-30R"
	PowerOutletTypeLabelNEMAL10Dash30R string = "NEMA L10-30R"

	// PowerOutletTypeLabelNEMAL14Dash20R captures enum value "NEMA L14-20R"
	PowerOutletTypeLabelNEMAL14Dash20R string = "NEMA L14-20R"

	// PowerOutletTypeLabelNEMAL14Dash30R captures enum value "NEMA L14-30R"
	PowerOutletTypeLabelNEMAL14Dash30R string = "NEMA L14-30R"

	// PowerOutletTypeLabelNEMAL14Dash50R captures enum value "NEMA L14-50R"
	PowerOutletTypeLabelNEMAL14Dash50R string = "NEMA L14-50R"

	// PowerOutletTypeLabelNEMAL14Dash60R captures enum value "NEMA L14-60R"
	PowerOutletTypeLabelNEMAL14Dash60R string = "NEMA L14-60R"

	// PowerOutletTypeLabelNEMAL15Dash20R captures enum value "NEMA L15-20R"
	PowerOutletTypeLabelNEMAL15Dash20R string = "NEMA L15-20R"

	// PowerOutletTypeLabelNEMAL15Dash30R captures enum value "NEMA L15-30R"
	PowerOutletTypeLabelNEMAL15Dash30R string = "NEMA L15-30R"

	// PowerOutletTypeLabelNEMAL15Dash50R captures enum value "NEMA L15-50R"
	PowerOutletTypeLabelNEMAL15Dash50R string = "NEMA L15-50R"

	// PowerOutletTypeLabelNEMAL15Dash60R captures enum value "NEMA L15-60R"
	PowerOutletTypeLabelNEMAL15Dash60R string = "NEMA L15-60R"

	// PowerOutletTypeLabelNEMAL21Dash20R captures enum value "NEMA L21-20R"
	PowerOutletTypeLabelNEMAL21Dash20R string = "NEMA L21-20R"

	// PowerOutletTypeLabelNEMAL21Dash30R captures enum value "NEMA L21-30R"
	PowerOutletTypeLabelNEMAL21Dash30R string = "NEMA L21-30R"

	// PowerOutletTypeLabelCS6360C captures enum value "CS6360C"
	PowerOutletTypeLabelCS6360C string = "CS6360C"

	// PowerOutletTypeLabelCS6364C captures enum value "CS6364C"
	PowerOutletTypeLabelCS6364C string = "CS6364C"

	// PowerOutletTypeLabelCS8164C captures enum value "CS8164C"
	PowerOutletTypeLabelCS8164C string = "CS8164C"

	// PowerOutletTypeLabelCS8264C captures enum value "CS8264C"
	PowerOutletTypeLabelCS8264C string = "CS8264C"

	// PowerOutletTypeLabelCS8364C captures enum value "CS8364C"
	PowerOutletTypeLabelCS8364C string = "CS8364C"

	// PowerOutletTypeLabelCS8464C captures enum value "CS8464C"
	PowerOutletTypeLabelCS8464C string = "CS8464C"

	// PowerOutletTypeLabelITATypeECEE75 captures enum value "ITA Type E (CEE7/5)"
	PowerOutletTypeLabelITATypeECEE75 string = "ITA Type E (CEE7/5)"

	// PowerOutletTypeLabelITATypeFCEE73 captures enum value "ITA Type F (CEE7/3)"
	PowerOutletTypeLabelITATypeFCEE73 string = "ITA Type F (CEE7/3)"

	// PowerOutletTypeLabelITATypeGBS1363 captures enum value "ITA Type G (BS 1363)"
	PowerOutletTypeLabelITATypeGBS1363 string = "ITA Type G (BS 1363)"

	// PowerOutletTypeLabelITATypeH captures enum value "ITA Type H"
	PowerOutletTypeLabelITATypeH string = "ITA Type H"

	// PowerOutletTypeLabelITATypeI captures enum value "ITA Type I"
	PowerOutletTypeLabelITATypeI string = "ITA Type I"

	// PowerOutletTypeLabelITATypeJ captures enum value "ITA Type J"
	PowerOutletTypeLabelITATypeJ string = "ITA Type J"

	// PowerOutletTypeLabelITATypeK captures enum value "ITA Type K"
	PowerOutletTypeLabelITATypeK string = "ITA Type K"

	// PowerOutletTypeLabelITATypeLCEI23Dash50 captures enum value "ITA Type L (CEI 23-50)"
	PowerOutletTypeLabelITATypeLCEI23Dash50 string = "ITA Type L (CEI 23-50)"

	// PowerOutletTypeLabelITATypeMBS546 captures enum value "ITA Type M (BS 546)"
	PowerOutletTypeLabelITATypeMBS546 string = "ITA Type M (BS 546)"

	// PowerOutletTypeLabelITATypeN captures enum value "ITA Type N"
	PowerOutletTypeLabelITATypeN string = "ITA Type N"

	// PowerOutletTypeLabelITATypeO captures enum value "ITA Type O"
	PowerOutletTypeLabelITATypeO string = "ITA Type O"

	// PowerOutletTypeLabelUSBTypeA captures enum value "USB Type A"
	PowerOutletTypeLabelUSBTypeA string = "USB Type A"

	// PowerOutletTypeLabelUSBMicroB captures enum value "USB Micro B"
	PowerOutletTypeLabelUSBMicroB string = "USB Micro B"

	// PowerOutletTypeLabelUSBTypeC captures enum value "USB Type C"
	PowerOutletTypeLabelUSBTypeC string = "USB Type C"

	// PowerOutletTypeLabelHDOTCx captures enum value "HDOT Cx"
	PowerOutletTypeLabelHDOTCx string = "HDOT Cx"
)

// prop value enum
func (m *PowerOutletType) validateLabelEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, powerOutletTypeTypeLabelPropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *PowerOutletType) validateLabel(formats strfmt.Registry) error {

	if err := validate.Required("type"+"."+"label", "body", m.Label); err != nil {
		return err
	}

	// value enum
	if err := m.validateLabelEnum("type"+"."+"label", "body", *m.Label); err != nil {
		return err
	}

	return nil
}

var powerOutletTypeTypeValuePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["iec-60320-c5","iec-60320-c7","iec-60320-c13","iec-60320-c15","iec-60320-c19","iec-60309-p-n-e-4h","iec-60309-p-n-e-6h","iec-60309-p-n-e-9h","iec-60309-2p-e-4h","iec-60309-2p-e-6h","iec-60309-2p-e-9h","iec-60309-3p-e-4h","iec-60309-3p-e-6h","iec-60309-3p-e-9h","iec-60309-3p-n-e-4h","iec-60309-3p-n-e-6h","iec-60309-3p-n-e-9h","nema-1-15r","nema-5-15r","nema-5-20r","nema-5-30r","nema-5-50r","nema-6-15r","nema-6-20r","nema-6-30r","nema-6-50r","nema-10-30r","nema-10-50r","nema-14-20r","nema-14-30r","nema-14-50r","nema-14-60r","nema-15-15r","nema-15-20r","nema-15-30r","nema-15-50r","nema-15-60r","nema-l1-15r","nema-l5-15r","nema-l5-20r","nema-l5-30r","nema-l5-50r","nema-l6-15r","nema-l6-20r","nema-l6-30r","nema-l6-50r","nema-l10-30r","nema-l14-20r","nema-l14-30r","nema-l14-50r","nema-l14-60r","nema-l15-20r","nema-l15-30r","nema-l15-50r","nema-l15-60r","nema-l21-20r","nema-l21-30r","CS6360C","CS6364C","CS8164C","CS8264C","CS8364C","CS8464C","ita-e","ita-f","ita-g","ita-h","ita-i","ita-j","ita-k","ita-l","ita-m","ita-n","ita-o","usb-a","usb-micro-b","usb-c","hdot-cx"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		powerOutletTypeTypeValuePropEnum = append(powerOutletTypeTypeValuePropEnum, v)
	}
}

const (

	// PowerOutletTypeValueIecDash60320DashC5 captures enum value "iec-60320-c5"
	PowerOutletTypeValueIecDash60320DashC5 string = "iec-60320-c5"

	// PowerOutletTypeValueIecDash60320DashC7 captures enum value "iec-60320-c7"
	PowerOutletTypeValueIecDash60320DashC7 string = "iec-60320-c7"

	// PowerOutletTypeValueIecDash60320DashC13 captures enum value "iec-60320-c13"
	PowerOutletTypeValueIecDash60320DashC13 string = "iec-60320-c13"

	// PowerOutletTypeValueIecDash60320DashC15 captures enum value "iec-60320-c15"
	PowerOutletTypeValueIecDash60320DashC15 string = "iec-60320-c15"

	// PowerOutletTypeValueIecDash60320DashC19 captures enum value "iec-60320-c19"
	PowerOutletTypeValueIecDash60320DashC19 string = "iec-60320-c19"

	// PowerOutletTypeValueIecDash60309DashpDashnDasheDash4h captures enum value "iec-60309-p-n-e-4h"
	PowerOutletTypeValueIecDash60309DashpDashnDasheDash4h string = "iec-60309-p-n-e-4h"

	// PowerOutletTypeValueIecDash60309DashpDashnDasheDash6h captures enum value "iec-60309-p-n-e-6h"
	PowerOutletTypeValueIecDash60309DashpDashnDasheDash6h string = "iec-60309-p-n-e-6h"

	// PowerOutletTypeValueIecDash60309DashpDashnDasheDash9h captures enum value "iec-60309-p-n-e-9h"
	PowerOutletTypeValueIecDash60309DashpDashnDasheDash9h string = "iec-60309-p-n-e-9h"

	// PowerOutletTypeValueIecDash60309Dash2pDasheDash4h captures enum value "iec-60309-2p-e-4h"
	PowerOutletTypeValueIecDash60309Dash2pDasheDash4h string = "iec-60309-2p-e-4h"

	// PowerOutletTypeValueIecDash60309Dash2pDasheDash6h captures enum value "iec-60309-2p-e-6h"
	PowerOutletTypeValueIecDash60309Dash2pDasheDash6h string = "iec-60309-2p-e-6h"

	// PowerOutletTypeValueIecDash60309Dash2pDasheDash9h captures enum value "iec-60309-2p-e-9h"
	PowerOutletTypeValueIecDash60309Dash2pDasheDash9h string = "iec-60309-2p-e-9h"

	// PowerOutletTypeValueIecDash60309Dash3pDasheDash4h captures enum value "iec-60309-3p-e-4h"
	PowerOutletTypeValueIecDash60309Dash3pDasheDash4h string = "iec-60309-3p-e-4h"

	// PowerOutletTypeValueIecDash60309Dash3pDasheDash6h captures enum value "iec-60309-3p-e-6h"
	PowerOutletTypeValueIecDash60309Dash3pDasheDash6h string = "iec-60309-3p-e-6h"

	// PowerOutletTypeValueIecDash60309Dash3pDasheDash9h captures enum value "iec-60309-3p-e-9h"
	PowerOutletTypeValueIecDash60309Dash3pDasheDash9h string = "iec-60309-3p-e-9h"

	// PowerOutletTypeValueIecDash60309Dash3pDashnDasheDash4h captures enum value "iec-60309-3p-n-e-4h"
	PowerOutletTypeValueIecDash60309Dash3pDashnDasheDash4h string = "iec-60309-3p-n-e-4h"

	// PowerOutletTypeValueIecDash60309Dash3pDashnDasheDash6h captures enum value "iec-60309-3p-n-e-6h"
	PowerOutletTypeValueIecDash60309Dash3pDashnDasheDash6h string = "iec-60309-3p-n-e-6h"

	// PowerOutletTypeValueIecDash60309Dash3pDashnDasheDash9h captures enum value "iec-60309-3p-n-e-9h"
	PowerOutletTypeValueIecDash60309Dash3pDashnDasheDash9h string = "iec-60309-3p-n-e-9h"

	// PowerOutletTypeValueNemaDash1Dash15r captures enum value "nema-1-15r"
	PowerOutletTypeValueNemaDash1Dash15r string = "nema-1-15r"

	// PowerOutletTypeValueNemaDash5Dash15r captures enum value "nema-5-15r"
	PowerOutletTypeValueNemaDash5Dash15r string = "nema-5-15r"

	// PowerOutletTypeValueNemaDash5Dash20r captures enum value "nema-5-20r"
	PowerOutletTypeValueNemaDash5Dash20r string = "nema-5-20r"

	// PowerOutletTypeValueNemaDash5Dash30r captures enum value "nema-5-30r"
	PowerOutletTypeValueNemaDash5Dash30r string = "nema-5-30r"

	// PowerOutletTypeValueNemaDash5Dash50r captures enum value "nema-5-50r"
	PowerOutletTypeValueNemaDash5Dash50r string = "nema-5-50r"

	// PowerOutletTypeValueNemaDash6Dash15r captures enum value "nema-6-15r"
	PowerOutletTypeValueNemaDash6Dash15r string = "nema-6-15r"

	// PowerOutletTypeValueNemaDash6Dash20r captures enum value "nema-6-20r"
	PowerOutletTypeValueNemaDash6Dash20r string = "nema-6-20r"

	// PowerOutletTypeValueNemaDash6Dash30r captures enum value "nema-6-30r"
	PowerOutletTypeValueNemaDash6Dash30r string = "nema-6-30r"

	// PowerOutletTypeValueNemaDash6Dash50r captures enum value "nema-6-50r"
	PowerOutletTypeValueNemaDash6Dash50r string = "nema-6-50r"

	// PowerOutletTypeValueNemaDash10Dash30r captures enum value "nema-10-30r"
	PowerOutletTypeValueNemaDash10Dash30r string = "nema-10-30r"

	// PowerOutletTypeValueNemaDash10Dash50r captures enum value "nema-10-50r"
	PowerOutletTypeValueNemaDash10Dash50r string = "nema-10-50r"

	// PowerOutletTypeValueNemaDash14Dash20r captures enum value "nema-14-20r"
	PowerOutletTypeValueNemaDash14Dash20r string = "nema-14-20r"

	// PowerOutletTypeValueNemaDash14Dash30r captures enum value "nema-14-30r"
	PowerOutletTypeValueNemaDash14Dash30r string = "nema-14-30r"

	// PowerOutletTypeValueNemaDash14Dash50r captures enum value "nema-14-50r"
	PowerOutletTypeValueNemaDash14Dash50r string = "nema-14-50r"

	// PowerOutletTypeValueNemaDash14Dash60r captures enum value "nema-14-60r"
	PowerOutletTypeValueNemaDash14Dash60r string = "nema-14-60r"

	// PowerOutletTypeValueNemaDash15Dash15r captures enum value "nema-15-15r"
	PowerOutletTypeValueNemaDash15Dash15r string = "nema-15-15r"

	// PowerOutletTypeValueNemaDash15Dash20r captures enum value "nema-15-20r"
	PowerOutletTypeValueNemaDash15Dash20r string = "nema-15-20r"

	// PowerOutletTypeValueNemaDash15Dash30r captures enum value "nema-15-30r"
	PowerOutletTypeValueNemaDash15Dash30r string = "nema-15-30r"

	// PowerOutletTypeValueNemaDash15Dash50r captures enum value "nema-15-50r"
	PowerOutletTypeValueNemaDash15Dash50r string = "nema-15-50r"

	// PowerOutletTypeValueNemaDash15Dash60r captures enum value "nema-15-60r"
	PowerOutletTypeValueNemaDash15Dash60r string = "nema-15-60r"

	// PowerOutletTypeValueNemaDashL1Dash15r captures enum value "nema-l1-15r"
	PowerOutletTypeValueNemaDashL1Dash15r string = "nema-l1-15r"

	// PowerOutletTypeValueNemaDashL5Dash15r captures enum value "nema-l5-15r"
	PowerOutletTypeValueNemaDashL5Dash15r string = "nema-l5-15r"

	// PowerOutletTypeValueNemaDashL5Dash20r captures enum value "nema-l5-20r"
	PowerOutletTypeValueNemaDashL5Dash20r string = "nema-l5-20r"

	// PowerOutletTypeValueNemaDashL5Dash30r captures enum value "nema-l5-30r"
	PowerOutletTypeValueNemaDashL5Dash30r string = "nema-l5-30r"

	// PowerOutletTypeValueNemaDashL5Dash50r captures enum value "nema-l5-50r"
	PowerOutletTypeValueNemaDashL5Dash50r string = "nema-l5-50r"

	// PowerOutletTypeValueNemaDashL6Dash15r captures enum value "nema-l6-15r"
	PowerOutletTypeValueNemaDashL6Dash15r string = "nema-l6-15r"

	// PowerOutletTypeValueNemaDashL6Dash20r captures enum value "nema-l6-20r"
	PowerOutletTypeValueNemaDashL6Dash20r string = "nema-l6-20r"

	// PowerOutletTypeValueNemaDashL6Dash30r captures enum value "nema-l6-30r"
	PowerOutletTypeValueNemaDashL6Dash30r string = "nema-l6-30r"

	// PowerOutletTypeValueNemaDashL6Dash50r captures enum value "nema-l6-50r"
	PowerOutletTypeValueNemaDashL6Dash50r string = "nema-l6-50r"

	// PowerOutletTypeValueNemaDashL10Dash30r captures enum value "nema-l10-30r"
	PowerOutletTypeValueNemaDashL10Dash30r string = "nema-l10-30r"

	// PowerOutletTypeValueNemaDashL14Dash20r captures enum value "nema-l14-20r"
	PowerOutletTypeValueNemaDashL14Dash20r string = "nema-l14-20r"

	// PowerOutletTypeValueNemaDashL14Dash30r captures enum value "nema-l14-30r"
	PowerOutletTypeValueNemaDashL14Dash30r string = "nema-l14-30r"

	// PowerOutletTypeValueNemaDashL14Dash50r captures enum value "nema-l14-50r"
	PowerOutletTypeValueNemaDashL14Dash50r string = "nema-l14-50r"

	// PowerOutletTypeValueNemaDashL14Dash60r captures enum value "nema-l14-60r"
	PowerOutletTypeValueNemaDashL14Dash60r string = "nema-l14-60r"

	// PowerOutletTypeValueNemaDashL15Dash20r captures enum value "nema-l15-20r"
	PowerOutletTypeValueNemaDashL15Dash20r string = "nema-l15-20r"

	// PowerOutletTypeValueNemaDashL15Dash30r captures enum value "nema-l15-30r"
	PowerOutletTypeValueNemaDashL15Dash30r string = "nema-l15-30r"

	// PowerOutletTypeValueNemaDashL15Dash50r captures enum value "nema-l15-50r"
	PowerOutletTypeValueNemaDashL15Dash50r string = "nema-l15-50r"

	// PowerOutletTypeValueNemaDashL15Dash60r captures enum value "nema-l15-60r"
	PowerOutletTypeValueNemaDashL15Dash60r string = "nema-l15-60r"

	// PowerOutletTypeValueNemaDashL21Dash20r captures enum value "nema-l21-20r"
	PowerOutletTypeValueNemaDashL21Dash20r string = "nema-l21-20r"

	// PowerOutletTypeValueNemaDashL21Dash30r captures enum value "nema-l21-30r"
	PowerOutletTypeValueNemaDashL21Dash30r string = "nema-l21-30r"

	// PowerOutletTypeValueCS6360C captures enum value "CS6360C"
	PowerOutletTypeValueCS6360C string = "CS6360C"

	// PowerOutletTypeValueCS6364C captures enum value "CS6364C"
	PowerOutletTypeValueCS6364C string = "CS6364C"

	// PowerOutletTypeValueCS8164C captures enum value "CS8164C"
	PowerOutletTypeValueCS8164C string = "CS8164C"

	// PowerOutletTypeValueCS8264C captures enum value "CS8264C"
	PowerOutletTypeValueCS8264C string = "CS8264C"

	// PowerOutletTypeValueCS8364C captures enum value "CS8364C"
	PowerOutletTypeValueCS8364C string = "CS8364C"

	// PowerOutletTypeValueCS8464C captures enum value "CS8464C"
	PowerOutletTypeValueCS8464C string = "CS8464C"

	// PowerOutletTypeValueItaDashe captures enum value "ita-e"
	PowerOutletTypeValueItaDashe string = "ita-e"

	// PowerOutletTypeValueItaDashf captures enum value "ita-f"
	PowerOutletTypeValueItaDashf string = "ita-f"

	// PowerOutletTypeValueItaDashg captures enum value "ita-g"
	PowerOutletTypeValueItaDashg string = "ita-g"

	// PowerOutletTypeValueItaDashh captures enum value "ita-h"
	PowerOutletTypeValueItaDashh string = "ita-h"

	// PowerOutletTypeValueItaDashi captures enum value "ita-i"
	PowerOutletTypeValueItaDashi string = "ita-i"

	// PowerOutletTypeValueItaDashj captures enum value "ita-j"
	PowerOutletTypeValueItaDashj string = "ita-j"

	// PowerOutletTypeValueItaDashk captures enum value "ita-k"
	PowerOutletTypeValueItaDashk string = "ita-k"

	// PowerOutletTypeValueItaDashl captures enum value "ita-l"
	PowerOutletTypeValueItaDashl string = "ita-l"

	// PowerOutletTypeValueItaDashm captures enum value "ita-m"
	PowerOutletTypeValueItaDashm string = "ita-m"

	// PowerOutletTypeValueItaDashn captures enum value "ita-n"
	PowerOutletTypeValueItaDashn string = "ita-n"

	// PowerOutletTypeValueItaDasho captures enum value "ita-o"
	PowerOutletTypeValueItaDasho string = "ita-o"

	// PowerOutletTypeValueUsbDasha captures enum value "usb-a"
	PowerOutletTypeValueUsbDasha string = "usb-a"

	// PowerOutletTypeValueUsbDashMicroDashb captures enum value "usb-micro-b"
	PowerOutletTypeValueUsbDashMicroDashb string = "usb-micro-b"

	// PowerOutletTypeValueUsbDashc captures enum value "usb-c"
	PowerOutletTypeValueUsbDashc string = "usb-c"

	// PowerOutletTypeValueHdotDashCx captures enum value "hdot-cx"
	PowerOutletTypeValueHdotDashCx string = "hdot-cx"
)

// prop value enum
func (m *PowerOutletType) validateValueEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, powerOutletTypeTypeValuePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *PowerOutletType) validateValue(formats strfmt.Registry) error {

	if err := validate.Required("type"+"."+"value", "body", m.Value); err != nil {
		return err
	}

	// value enum
	if err := m.validateValueEnum("type"+"."+"value", "body", *m.Value); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this power outlet type based on context it is used
func (m *PowerOutletType) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *PowerOutletType) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *PowerOutletType) UnmarshalBinary(b []byte) error {
	var res PowerOutletType
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
