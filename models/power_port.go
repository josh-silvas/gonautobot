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

// PowerPort power port
//
// swagger:model PowerPort
type PowerPort struct {

	// Allocated draw
	//
	// Allocated power draw (watts)
	// Maximum: 32767
	// Minimum: 1
	AllocatedDraw *int64 `json:"allocated_draw,omitempty"`

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

	// tags
	Tags []*NestedTag `json:"tags"`

	// type
	Type *PowerPortType `json:"type,omitempty"`

	// Url
	// Read Only: true
	// Format: uri
	URL strfmt.URI `json:"url,omitempty"`
}

// Validate validates this power port
func (m *PowerPort) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAllocatedDraw(formats); err != nil {
		res = append(res, err)
	}

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

	if err := m.validateMaximumDraw(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateName(formats); err != nil {
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

func (m *PowerPort) validateAllocatedDraw(formats strfmt.Registry) error {
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

func (m *PowerPort) validateCable(formats strfmt.Registry) error {
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

func (m *PowerPort) validateDescription(formats strfmt.Registry) error {
	if swag.IsZero(m.Description) { // not required
		return nil
	}

	if err := validate.MaxLength("description", "body", m.Description, 200); err != nil {
		return err
	}

	return nil
}

func (m *PowerPort) validateDevice(formats strfmt.Registry) error {

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

func (m *PowerPort) validateDisplay(formats strfmt.Registry) error {
	if swag.IsZero(m.Display) { // not required
		return nil
	}

	if err := validate.MinLength("display", "body", m.Display, 1); err != nil {
		return err
	}

	return nil
}

func (m *PowerPort) validateID(formats strfmt.Registry) error {
	if swag.IsZero(m.ID) { // not required
		return nil
	}

	if err := validate.FormatOf("id", "body", "uuid", m.ID.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *PowerPort) validateLabel(formats strfmt.Registry) error {
	if swag.IsZero(m.Label) { // not required
		return nil
	}

	if err := validate.MaxLength("label", "body", m.Label, 64); err != nil {
		return err
	}

	return nil
}

func (m *PowerPort) validateMaximumDraw(formats strfmt.Registry) error {
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

func (m *PowerPort) validateName(formats strfmt.Registry) error {

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

func (m *PowerPort) validateTags(formats strfmt.Registry) error {
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

func (m *PowerPort) validateType(formats strfmt.Registry) error {
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

func (m *PowerPort) validateURL(formats strfmt.Registry) error {
	if swag.IsZero(m.URL) { // not required
		return nil
	}

	if err := validate.FormatOf("url", "body", "uri", m.URL.String(), formats); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this power port based on the context it is used
func (m *PowerPort) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
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

func (m *PowerPort) contextValidateCable(ctx context.Context, formats strfmt.Registry) error {

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

func (m *PowerPort) contextValidateCablePeer(ctx context.Context, formats strfmt.Registry) error {

	return nil
}

func (m *PowerPort) contextValidateCablePeerType(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "cable_peer_type", "body", string(m.CablePeerType)); err != nil {
		return err
	}

	return nil
}

func (m *PowerPort) contextValidateConnectedEndpoint(ctx context.Context, formats strfmt.Registry) error {

	return nil
}

func (m *PowerPort) contextValidateConnectedEndpointReachable(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "connected_endpoint_reachable", "body", m.ConnectedEndpointReachable); err != nil {
		return err
	}

	return nil
}

func (m *PowerPort) contextValidateConnectedEndpointType(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "connected_endpoint_type", "body", string(m.ConnectedEndpointType)); err != nil {
		return err
	}

	return nil
}

func (m *PowerPort) contextValidateDevice(ctx context.Context, formats strfmt.Registry) error {

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

func (m *PowerPort) contextValidateDisplay(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "display", "body", string(m.Display)); err != nil {
		return err
	}

	return nil
}

func (m *PowerPort) contextValidateID(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "id", "body", strfmt.UUID(m.ID)); err != nil {
		return err
	}

	return nil
}

func (m *PowerPort) contextValidateTags(ctx context.Context, formats strfmt.Registry) error {

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

func (m *PowerPort) contextValidateType(ctx context.Context, formats strfmt.Registry) error {

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

func (m *PowerPort) contextValidateURL(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "url", "body", strfmt.URI(m.URL)); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *PowerPort) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *PowerPort) UnmarshalBinary(b []byte) error {
	var res PowerPort
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// PowerPortType Type
//
// swagger:model PowerPortType
type PowerPortType struct {

	// label
	// Required: true
	// Enum: [C6 C8 C14 C16 C20 P+N+E 4H P+N+E 6H P+N+E 9H 2P+E 4H 2P+E 6H 2P+E 9H 3P+E 4H 3P+E 6H 3P+E 9H 3P+N+E 4H 3P+N+E 6H 3P+N+E 9H NEMA 1-15P NEMA 5-15P NEMA 5-20P NEMA 5-30P NEMA 5-50P NEMA 6-15P NEMA 6-20P NEMA 6-30P NEMA 6-50P NEMA 10-30P NEMA 10-50P NEMA 14-20P NEMA 14-30P NEMA 14-50P NEMA 14-60P NEMA 15-15P NEMA 15-20P NEMA 15-30P NEMA 15-50P NEMA 15-60P NEMA L1-15P NEMA L5-15P NEMA L5-20P NEMA L5-30P NEMA L5-50P NEMA L6-15P NEMA L6-20P NEMA L6-30P NEMA L6-50P NEMA L10-30P NEMA L14-20P NEMA L14-30P NEMA L14-50P NEMA L14-60P NEMA L15-20P NEMA L15-30P NEMA L15-50P NEMA L15-60P NEMA L21-20P NEMA L21-30P CS6361C CS6365C CS8165C CS8265C CS8365C CS8465C ITA Type E (CEE 7/5) ITA Type F (CEE 7/4) ITA Type E/F (CEE 7/7) ITA Type G (BS 1363) ITA Type H ITA Type I ITA Type J ITA Type K ITA Type L (CEI 23-50) ITA Type M (BS 546) ITA Type N ITA Type O USB Type A USB Type B USB Type C USB Mini A USB Mini B USB Micro A USB Micro B USB 3.0 Type B USB 3.0 Micro B]
	Label *string `json:"label"`

	// value
	// Required: true
	// Enum: [iec-60320-c6 iec-60320-c8 iec-60320-c14 iec-60320-c16 iec-60320-c20 iec-60309-p-n-e-4h iec-60309-p-n-e-6h iec-60309-p-n-e-9h iec-60309-2p-e-4h iec-60309-2p-e-6h iec-60309-2p-e-9h iec-60309-3p-e-4h iec-60309-3p-e-6h iec-60309-3p-e-9h iec-60309-3p-n-e-4h iec-60309-3p-n-e-6h iec-60309-3p-n-e-9h nema-1-15p nema-5-15p nema-5-20p nema-5-30p nema-5-50p nema-6-15p nema-6-20p nema-6-30p nema-6-50p nema-10-30p nema-10-50p nema-14-20p nema-14-30p nema-14-50p nema-14-60p nema-15-15p nema-15-20p nema-15-30p nema-15-50p nema-15-60p nema-l1-15p nema-l5-15p nema-l5-20p nema-l5-30p nema-l5-50p nema-l6-15p nema-l6-20p nema-l6-30p nema-l6-50p nema-l10-30p nema-l14-20p nema-l14-30p nema-l14-50p nema-l14-60p nema-l15-20p nema-l15-30p nema-l15-50p nema-l15-60p nema-l21-20p nema-l21-30p cs6361c cs6365c cs8165c cs8265c cs8365c cs8465c ita-e ita-f ita-ef ita-g ita-h ita-i ita-j ita-k ita-l ita-m ita-n ita-o usb-a usb-b usb-c usb-mini-a usb-mini-b usb-micro-a usb-micro-b usb-3-b usb-3-micro-b]
	Value *string `json:"value"`
}

// Validate validates this power port type
func (m *PowerPortType) Validate(formats strfmt.Registry) error {
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

var powerPortTypeTypeLabelPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["C6","C8","C14","C16","C20","P+N+E 4H","P+N+E 6H","P+N+E 9H","2P+E 4H","2P+E 6H","2P+E 9H","3P+E 4H","3P+E 6H","3P+E 9H","3P+N+E 4H","3P+N+E 6H","3P+N+E 9H","NEMA 1-15P","NEMA 5-15P","NEMA 5-20P","NEMA 5-30P","NEMA 5-50P","NEMA 6-15P","NEMA 6-20P","NEMA 6-30P","NEMA 6-50P","NEMA 10-30P","NEMA 10-50P","NEMA 14-20P","NEMA 14-30P","NEMA 14-50P","NEMA 14-60P","NEMA 15-15P","NEMA 15-20P","NEMA 15-30P","NEMA 15-50P","NEMA 15-60P","NEMA L1-15P","NEMA L5-15P","NEMA L5-20P","NEMA L5-30P","NEMA L5-50P","NEMA L6-15P","NEMA L6-20P","NEMA L6-30P","NEMA L6-50P","NEMA L10-30P","NEMA L14-20P","NEMA L14-30P","NEMA L14-50P","NEMA L14-60P","NEMA L15-20P","NEMA L15-30P","NEMA L15-50P","NEMA L15-60P","NEMA L21-20P","NEMA L21-30P","CS6361C","CS6365C","CS8165C","CS8265C","CS8365C","CS8465C","ITA Type E (CEE 7/5)","ITA Type F (CEE 7/4)","ITA Type E/F (CEE 7/7)","ITA Type G (BS 1363)","ITA Type H","ITA Type I","ITA Type J","ITA Type K","ITA Type L (CEI 23-50)","ITA Type M (BS 546)","ITA Type N","ITA Type O","USB Type A","USB Type B","USB Type C","USB Mini A","USB Mini B","USB Micro A","USB Micro B","USB 3.0 Type B","USB 3.0 Micro B"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		powerPortTypeTypeLabelPropEnum = append(powerPortTypeTypeLabelPropEnum, v)
	}
}

const (

	// PowerPortTypeLabelC6 captures enum value "C6"
	PowerPortTypeLabelC6 string = "C6"

	// PowerPortTypeLabelC8 captures enum value "C8"
	PowerPortTypeLabelC8 string = "C8"

	// PowerPortTypeLabelC14 captures enum value "C14"
	PowerPortTypeLabelC14 string = "C14"

	// PowerPortTypeLabelC16 captures enum value "C16"
	PowerPortTypeLabelC16 string = "C16"

	// PowerPortTypeLabelC20 captures enum value "C20"
	PowerPortTypeLabelC20 string = "C20"

	// PowerPortTypeLabelPPlusNPlusE4H captures enum value "P+N+E 4H"
	PowerPortTypeLabelPPlusNPlusE4H string = "P+N+E 4H"

	// PowerPortTypeLabelPPlusNPlusE6H captures enum value "P+N+E 6H"
	PowerPortTypeLabelPPlusNPlusE6H string = "P+N+E 6H"

	// PowerPortTypeLabelPPlusNPlusE9H captures enum value "P+N+E 9H"
	PowerPortTypeLabelPPlusNPlusE9H string = "P+N+E 9H"

	// PowerPortTypeLabelNr2PPlusE4H captures enum value "2P+E 4H"
	PowerPortTypeLabelNr2PPlusE4H string = "2P+E 4H"

	// PowerPortTypeLabelNr2PPlusE6H captures enum value "2P+E 6H"
	PowerPortTypeLabelNr2PPlusE6H string = "2P+E 6H"

	// PowerPortTypeLabelNr2PPlusE9H captures enum value "2P+E 9H"
	PowerPortTypeLabelNr2PPlusE9H string = "2P+E 9H"

	// PowerPortTypeLabelNr3PPlusE4H captures enum value "3P+E 4H"
	PowerPortTypeLabelNr3PPlusE4H string = "3P+E 4H"

	// PowerPortTypeLabelNr3PPlusE6H captures enum value "3P+E 6H"
	PowerPortTypeLabelNr3PPlusE6H string = "3P+E 6H"

	// PowerPortTypeLabelNr3PPlusE9H captures enum value "3P+E 9H"
	PowerPortTypeLabelNr3PPlusE9H string = "3P+E 9H"

	// PowerPortTypeLabelNr3PPlusNPlusE4H captures enum value "3P+N+E 4H"
	PowerPortTypeLabelNr3PPlusNPlusE4H string = "3P+N+E 4H"

	// PowerPortTypeLabelNr3PPlusNPlusE6H captures enum value "3P+N+E 6H"
	PowerPortTypeLabelNr3PPlusNPlusE6H string = "3P+N+E 6H"

	// PowerPortTypeLabelNr3PPlusNPlusE9H captures enum value "3P+N+E 9H"
	PowerPortTypeLabelNr3PPlusNPlusE9H string = "3P+N+E 9H"

	// PowerPortTypeLabelNEMA1Dash15P captures enum value "NEMA 1-15P"
	PowerPortTypeLabelNEMA1Dash15P string = "NEMA 1-15P"

	// PowerPortTypeLabelNEMA5Dash15P captures enum value "NEMA 5-15P"
	PowerPortTypeLabelNEMA5Dash15P string = "NEMA 5-15P"

	// PowerPortTypeLabelNEMA5Dash20P captures enum value "NEMA 5-20P"
	PowerPortTypeLabelNEMA5Dash20P string = "NEMA 5-20P"

	// PowerPortTypeLabelNEMA5Dash30P captures enum value "NEMA 5-30P"
	PowerPortTypeLabelNEMA5Dash30P string = "NEMA 5-30P"

	// PowerPortTypeLabelNEMA5Dash50P captures enum value "NEMA 5-50P"
	PowerPortTypeLabelNEMA5Dash50P string = "NEMA 5-50P"

	// PowerPortTypeLabelNEMA6Dash15P captures enum value "NEMA 6-15P"
	PowerPortTypeLabelNEMA6Dash15P string = "NEMA 6-15P"

	// PowerPortTypeLabelNEMA6Dash20P captures enum value "NEMA 6-20P"
	PowerPortTypeLabelNEMA6Dash20P string = "NEMA 6-20P"

	// PowerPortTypeLabelNEMA6Dash30P captures enum value "NEMA 6-30P"
	PowerPortTypeLabelNEMA6Dash30P string = "NEMA 6-30P"

	// PowerPortTypeLabelNEMA6Dash50P captures enum value "NEMA 6-50P"
	PowerPortTypeLabelNEMA6Dash50P string = "NEMA 6-50P"

	// PowerPortTypeLabelNEMA10Dash30P captures enum value "NEMA 10-30P"
	PowerPortTypeLabelNEMA10Dash30P string = "NEMA 10-30P"

	// PowerPortTypeLabelNEMA10Dash50P captures enum value "NEMA 10-50P"
	PowerPortTypeLabelNEMA10Dash50P string = "NEMA 10-50P"

	// PowerPortTypeLabelNEMA14Dash20P captures enum value "NEMA 14-20P"
	PowerPortTypeLabelNEMA14Dash20P string = "NEMA 14-20P"

	// PowerPortTypeLabelNEMA14Dash30P captures enum value "NEMA 14-30P"
	PowerPortTypeLabelNEMA14Dash30P string = "NEMA 14-30P"

	// PowerPortTypeLabelNEMA14Dash50P captures enum value "NEMA 14-50P"
	PowerPortTypeLabelNEMA14Dash50P string = "NEMA 14-50P"

	// PowerPortTypeLabelNEMA14Dash60P captures enum value "NEMA 14-60P"
	PowerPortTypeLabelNEMA14Dash60P string = "NEMA 14-60P"

	// PowerPortTypeLabelNEMA15Dash15P captures enum value "NEMA 15-15P"
	PowerPortTypeLabelNEMA15Dash15P string = "NEMA 15-15P"

	// PowerPortTypeLabelNEMA15Dash20P captures enum value "NEMA 15-20P"
	PowerPortTypeLabelNEMA15Dash20P string = "NEMA 15-20P"

	// PowerPortTypeLabelNEMA15Dash30P captures enum value "NEMA 15-30P"
	PowerPortTypeLabelNEMA15Dash30P string = "NEMA 15-30P"

	// PowerPortTypeLabelNEMA15Dash50P captures enum value "NEMA 15-50P"
	PowerPortTypeLabelNEMA15Dash50P string = "NEMA 15-50P"

	// PowerPortTypeLabelNEMA15Dash60P captures enum value "NEMA 15-60P"
	PowerPortTypeLabelNEMA15Dash60P string = "NEMA 15-60P"

	// PowerPortTypeLabelNEMAL1Dash15P captures enum value "NEMA L1-15P"
	PowerPortTypeLabelNEMAL1Dash15P string = "NEMA L1-15P"

	// PowerPortTypeLabelNEMAL5Dash15P captures enum value "NEMA L5-15P"
	PowerPortTypeLabelNEMAL5Dash15P string = "NEMA L5-15P"

	// PowerPortTypeLabelNEMAL5Dash20P captures enum value "NEMA L5-20P"
	PowerPortTypeLabelNEMAL5Dash20P string = "NEMA L5-20P"

	// PowerPortTypeLabelNEMAL5Dash30P captures enum value "NEMA L5-30P"
	PowerPortTypeLabelNEMAL5Dash30P string = "NEMA L5-30P"

	// PowerPortTypeLabelNEMAL5Dash50P captures enum value "NEMA L5-50P"
	PowerPortTypeLabelNEMAL5Dash50P string = "NEMA L5-50P"

	// PowerPortTypeLabelNEMAL6Dash15P captures enum value "NEMA L6-15P"
	PowerPortTypeLabelNEMAL6Dash15P string = "NEMA L6-15P"

	// PowerPortTypeLabelNEMAL6Dash20P captures enum value "NEMA L6-20P"
	PowerPortTypeLabelNEMAL6Dash20P string = "NEMA L6-20P"

	// PowerPortTypeLabelNEMAL6Dash30P captures enum value "NEMA L6-30P"
	PowerPortTypeLabelNEMAL6Dash30P string = "NEMA L6-30P"

	// PowerPortTypeLabelNEMAL6Dash50P captures enum value "NEMA L6-50P"
	PowerPortTypeLabelNEMAL6Dash50P string = "NEMA L6-50P"

	// PowerPortTypeLabelNEMAL10Dash30P captures enum value "NEMA L10-30P"
	PowerPortTypeLabelNEMAL10Dash30P string = "NEMA L10-30P"

	// PowerPortTypeLabelNEMAL14Dash20P captures enum value "NEMA L14-20P"
	PowerPortTypeLabelNEMAL14Dash20P string = "NEMA L14-20P"

	// PowerPortTypeLabelNEMAL14Dash30P captures enum value "NEMA L14-30P"
	PowerPortTypeLabelNEMAL14Dash30P string = "NEMA L14-30P"

	// PowerPortTypeLabelNEMAL14Dash50P captures enum value "NEMA L14-50P"
	PowerPortTypeLabelNEMAL14Dash50P string = "NEMA L14-50P"

	// PowerPortTypeLabelNEMAL14Dash60P captures enum value "NEMA L14-60P"
	PowerPortTypeLabelNEMAL14Dash60P string = "NEMA L14-60P"

	// PowerPortTypeLabelNEMAL15Dash20P captures enum value "NEMA L15-20P"
	PowerPortTypeLabelNEMAL15Dash20P string = "NEMA L15-20P"

	// PowerPortTypeLabelNEMAL15Dash30P captures enum value "NEMA L15-30P"
	PowerPortTypeLabelNEMAL15Dash30P string = "NEMA L15-30P"

	// PowerPortTypeLabelNEMAL15Dash50P captures enum value "NEMA L15-50P"
	PowerPortTypeLabelNEMAL15Dash50P string = "NEMA L15-50P"

	// PowerPortTypeLabelNEMAL15Dash60P captures enum value "NEMA L15-60P"
	PowerPortTypeLabelNEMAL15Dash60P string = "NEMA L15-60P"

	// PowerPortTypeLabelNEMAL21Dash20P captures enum value "NEMA L21-20P"
	PowerPortTypeLabelNEMAL21Dash20P string = "NEMA L21-20P"

	// PowerPortTypeLabelNEMAL21Dash30P captures enum value "NEMA L21-30P"
	PowerPortTypeLabelNEMAL21Dash30P string = "NEMA L21-30P"

	// PowerPortTypeLabelCS6361C captures enum value "CS6361C"
	PowerPortTypeLabelCS6361C string = "CS6361C"

	// PowerPortTypeLabelCS6365C captures enum value "CS6365C"
	PowerPortTypeLabelCS6365C string = "CS6365C"

	// PowerPortTypeLabelCS8165C captures enum value "CS8165C"
	PowerPortTypeLabelCS8165C string = "CS8165C"

	// PowerPortTypeLabelCS8265C captures enum value "CS8265C"
	PowerPortTypeLabelCS8265C string = "CS8265C"

	// PowerPortTypeLabelCS8365C captures enum value "CS8365C"
	PowerPortTypeLabelCS8365C string = "CS8365C"

	// PowerPortTypeLabelCS8465C captures enum value "CS8465C"
	PowerPortTypeLabelCS8465C string = "CS8465C"

	// PowerPortTypeLabelITATypeECEE75 captures enum value "ITA Type E (CEE 7/5)"
	PowerPortTypeLabelITATypeECEE75 string = "ITA Type E (CEE 7/5)"

	// PowerPortTypeLabelITATypeFCEE74 captures enum value "ITA Type F (CEE 7/4)"
	PowerPortTypeLabelITATypeFCEE74 string = "ITA Type F (CEE 7/4)"

	// PowerPortTypeLabelITATypeEFCEE77 captures enum value "ITA Type E/F (CEE 7/7)"
	PowerPortTypeLabelITATypeEFCEE77 string = "ITA Type E/F (CEE 7/7)"

	// PowerPortTypeLabelITATypeGBS1363 captures enum value "ITA Type G (BS 1363)"
	PowerPortTypeLabelITATypeGBS1363 string = "ITA Type G (BS 1363)"

	// PowerPortTypeLabelITATypeH captures enum value "ITA Type H"
	PowerPortTypeLabelITATypeH string = "ITA Type H"

	// PowerPortTypeLabelITATypeI captures enum value "ITA Type I"
	PowerPortTypeLabelITATypeI string = "ITA Type I"

	// PowerPortTypeLabelITATypeJ captures enum value "ITA Type J"
	PowerPortTypeLabelITATypeJ string = "ITA Type J"

	// PowerPortTypeLabelITATypeK captures enum value "ITA Type K"
	PowerPortTypeLabelITATypeK string = "ITA Type K"

	// PowerPortTypeLabelITATypeLCEI23Dash50 captures enum value "ITA Type L (CEI 23-50)"
	PowerPortTypeLabelITATypeLCEI23Dash50 string = "ITA Type L (CEI 23-50)"

	// PowerPortTypeLabelITATypeMBS546 captures enum value "ITA Type M (BS 546)"
	PowerPortTypeLabelITATypeMBS546 string = "ITA Type M (BS 546)"

	// PowerPortTypeLabelITATypeN captures enum value "ITA Type N"
	PowerPortTypeLabelITATypeN string = "ITA Type N"

	// PowerPortTypeLabelITATypeO captures enum value "ITA Type O"
	PowerPortTypeLabelITATypeO string = "ITA Type O"

	// PowerPortTypeLabelUSBTypeA captures enum value "USB Type A"
	PowerPortTypeLabelUSBTypeA string = "USB Type A"

	// PowerPortTypeLabelUSBTypeB captures enum value "USB Type B"
	PowerPortTypeLabelUSBTypeB string = "USB Type B"

	// PowerPortTypeLabelUSBTypeC captures enum value "USB Type C"
	PowerPortTypeLabelUSBTypeC string = "USB Type C"

	// PowerPortTypeLabelUSBMiniA captures enum value "USB Mini A"
	PowerPortTypeLabelUSBMiniA string = "USB Mini A"

	// PowerPortTypeLabelUSBMiniB captures enum value "USB Mini B"
	PowerPortTypeLabelUSBMiniB string = "USB Mini B"

	// PowerPortTypeLabelUSBMicroA captures enum value "USB Micro A"
	PowerPortTypeLabelUSBMicroA string = "USB Micro A"

	// PowerPortTypeLabelUSBMicroB captures enum value "USB Micro B"
	PowerPortTypeLabelUSBMicroB string = "USB Micro B"

	// PowerPortTypeLabelUSB3Dot0TypeB captures enum value "USB 3.0 Type B"
	PowerPortTypeLabelUSB3Dot0TypeB string = "USB 3.0 Type B"

	// PowerPortTypeLabelUSB3Dot0MicroB captures enum value "USB 3.0 Micro B"
	PowerPortTypeLabelUSB3Dot0MicroB string = "USB 3.0 Micro B"
)

// prop value enum
func (m *PowerPortType) validateLabelEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, powerPortTypeTypeLabelPropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *PowerPortType) validateLabel(formats strfmt.Registry) error {

	if err := validate.Required("type"+"."+"label", "body", m.Label); err != nil {
		return err
	}

	// value enum
	if err := m.validateLabelEnum("type"+"."+"label", "body", *m.Label); err != nil {
		return err
	}

	return nil
}

var powerPortTypeTypeValuePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["iec-60320-c6","iec-60320-c8","iec-60320-c14","iec-60320-c16","iec-60320-c20","iec-60309-p-n-e-4h","iec-60309-p-n-e-6h","iec-60309-p-n-e-9h","iec-60309-2p-e-4h","iec-60309-2p-e-6h","iec-60309-2p-e-9h","iec-60309-3p-e-4h","iec-60309-3p-e-6h","iec-60309-3p-e-9h","iec-60309-3p-n-e-4h","iec-60309-3p-n-e-6h","iec-60309-3p-n-e-9h","nema-1-15p","nema-5-15p","nema-5-20p","nema-5-30p","nema-5-50p","nema-6-15p","nema-6-20p","nema-6-30p","nema-6-50p","nema-10-30p","nema-10-50p","nema-14-20p","nema-14-30p","nema-14-50p","nema-14-60p","nema-15-15p","nema-15-20p","nema-15-30p","nema-15-50p","nema-15-60p","nema-l1-15p","nema-l5-15p","nema-l5-20p","nema-l5-30p","nema-l5-50p","nema-l6-15p","nema-l6-20p","nema-l6-30p","nema-l6-50p","nema-l10-30p","nema-l14-20p","nema-l14-30p","nema-l14-50p","nema-l14-60p","nema-l15-20p","nema-l15-30p","nema-l15-50p","nema-l15-60p","nema-l21-20p","nema-l21-30p","cs6361c","cs6365c","cs8165c","cs8265c","cs8365c","cs8465c","ita-e","ita-f","ita-ef","ita-g","ita-h","ita-i","ita-j","ita-k","ita-l","ita-m","ita-n","ita-o","usb-a","usb-b","usb-c","usb-mini-a","usb-mini-b","usb-micro-a","usb-micro-b","usb-3-b","usb-3-micro-b"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		powerPortTypeTypeValuePropEnum = append(powerPortTypeTypeValuePropEnum, v)
	}
}

const (

	// PowerPortTypeValueIecDash60320DashC6 captures enum value "iec-60320-c6"
	PowerPortTypeValueIecDash60320DashC6 string = "iec-60320-c6"

	// PowerPortTypeValueIecDash60320DashC8 captures enum value "iec-60320-c8"
	PowerPortTypeValueIecDash60320DashC8 string = "iec-60320-c8"

	// PowerPortTypeValueIecDash60320DashC14 captures enum value "iec-60320-c14"
	PowerPortTypeValueIecDash60320DashC14 string = "iec-60320-c14"

	// PowerPortTypeValueIecDash60320DashC16 captures enum value "iec-60320-c16"
	PowerPortTypeValueIecDash60320DashC16 string = "iec-60320-c16"

	// PowerPortTypeValueIecDash60320DashC20 captures enum value "iec-60320-c20"
	PowerPortTypeValueIecDash60320DashC20 string = "iec-60320-c20"

	// PowerPortTypeValueIecDash60309DashpDashnDasheDash4h captures enum value "iec-60309-p-n-e-4h"
	PowerPortTypeValueIecDash60309DashpDashnDasheDash4h string = "iec-60309-p-n-e-4h"

	// PowerPortTypeValueIecDash60309DashpDashnDasheDash6h captures enum value "iec-60309-p-n-e-6h"
	PowerPortTypeValueIecDash60309DashpDashnDasheDash6h string = "iec-60309-p-n-e-6h"

	// PowerPortTypeValueIecDash60309DashpDashnDasheDash9h captures enum value "iec-60309-p-n-e-9h"
	PowerPortTypeValueIecDash60309DashpDashnDasheDash9h string = "iec-60309-p-n-e-9h"

	// PowerPortTypeValueIecDash60309Dash2pDasheDash4h captures enum value "iec-60309-2p-e-4h"
	PowerPortTypeValueIecDash60309Dash2pDasheDash4h string = "iec-60309-2p-e-4h"

	// PowerPortTypeValueIecDash60309Dash2pDasheDash6h captures enum value "iec-60309-2p-e-6h"
	PowerPortTypeValueIecDash60309Dash2pDasheDash6h string = "iec-60309-2p-e-6h"

	// PowerPortTypeValueIecDash60309Dash2pDasheDash9h captures enum value "iec-60309-2p-e-9h"
	PowerPortTypeValueIecDash60309Dash2pDasheDash9h string = "iec-60309-2p-e-9h"

	// PowerPortTypeValueIecDash60309Dash3pDasheDash4h captures enum value "iec-60309-3p-e-4h"
	PowerPortTypeValueIecDash60309Dash3pDasheDash4h string = "iec-60309-3p-e-4h"

	// PowerPortTypeValueIecDash60309Dash3pDasheDash6h captures enum value "iec-60309-3p-e-6h"
	PowerPortTypeValueIecDash60309Dash3pDasheDash6h string = "iec-60309-3p-e-6h"

	// PowerPortTypeValueIecDash60309Dash3pDasheDash9h captures enum value "iec-60309-3p-e-9h"
	PowerPortTypeValueIecDash60309Dash3pDasheDash9h string = "iec-60309-3p-e-9h"

	// PowerPortTypeValueIecDash60309Dash3pDashnDasheDash4h captures enum value "iec-60309-3p-n-e-4h"
	PowerPortTypeValueIecDash60309Dash3pDashnDasheDash4h string = "iec-60309-3p-n-e-4h"

	// PowerPortTypeValueIecDash60309Dash3pDashnDasheDash6h captures enum value "iec-60309-3p-n-e-6h"
	PowerPortTypeValueIecDash60309Dash3pDashnDasheDash6h string = "iec-60309-3p-n-e-6h"

	// PowerPortTypeValueIecDash60309Dash3pDashnDasheDash9h captures enum value "iec-60309-3p-n-e-9h"
	PowerPortTypeValueIecDash60309Dash3pDashnDasheDash9h string = "iec-60309-3p-n-e-9h"

	// PowerPortTypeValueNemaDash1Dash15p captures enum value "nema-1-15p"
	PowerPortTypeValueNemaDash1Dash15p string = "nema-1-15p"

	// PowerPortTypeValueNemaDash5Dash15p captures enum value "nema-5-15p"
	PowerPortTypeValueNemaDash5Dash15p string = "nema-5-15p"

	// PowerPortTypeValueNemaDash5Dash20p captures enum value "nema-5-20p"
	PowerPortTypeValueNemaDash5Dash20p string = "nema-5-20p"

	// PowerPortTypeValueNemaDash5Dash30p captures enum value "nema-5-30p"
	PowerPortTypeValueNemaDash5Dash30p string = "nema-5-30p"

	// PowerPortTypeValueNemaDash5Dash50p captures enum value "nema-5-50p"
	PowerPortTypeValueNemaDash5Dash50p string = "nema-5-50p"

	// PowerPortTypeValueNemaDash6Dash15p captures enum value "nema-6-15p"
	PowerPortTypeValueNemaDash6Dash15p string = "nema-6-15p"

	// PowerPortTypeValueNemaDash6Dash20p captures enum value "nema-6-20p"
	PowerPortTypeValueNemaDash6Dash20p string = "nema-6-20p"

	// PowerPortTypeValueNemaDash6Dash30p captures enum value "nema-6-30p"
	PowerPortTypeValueNemaDash6Dash30p string = "nema-6-30p"

	// PowerPortTypeValueNemaDash6Dash50p captures enum value "nema-6-50p"
	PowerPortTypeValueNemaDash6Dash50p string = "nema-6-50p"

	// PowerPortTypeValueNemaDash10Dash30p captures enum value "nema-10-30p"
	PowerPortTypeValueNemaDash10Dash30p string = "nema-10-30p"

	// PowerPortTypeValueNemaDash10Dash50p captures enum value "nema-10-50p"
	PowerPortTypeValueNemaDash10Dash50p string = "nema-10-50p"

	// PowerPortTypeValueNemaDash14Dash20p captures enum value "nema-14-20p"
	PowerPortTypeValueNemaDash14Dash20p string = "nema-14-20p"

	// PowerPortTypeValueNemaDash14Dash30p captures enum value "nema-14-30p"
	PowerPortTypeValueNemaDash14Dash30p string = "nema-14-30p"

	// PowerPortTypeValueNemaDash14Dash50p captures enum value "nema-14-50p"
	PowerPortTypeValueNemaDash14Dash50p string = "nema-14-50p"

	// PowerPortTypeValueNemaDash14Dash60p captures enum value "nema-14-60p"
	PowerPortTypeValueNemaDash14Dash60p string = "nema-14-60p"

	// PowerPortTypeValueNemaDash15Dash15p captures enum value "nema-15-15p"
	PowerPortTypeValueNemaDash15Dash15p string = "nema-15-15p"

	// PowerPortTypeValueNemaDash15Dash20p captures enum value "nema-15-20p"
	PowerPortTypeValueNemaDash15Dash20p string = "nema-15-20p"

	// PowerPortTypeValueNemaDash15Dash30p captures enum value "nema-15-30p"
	PowerPortTypeValueNemaDash15Dash30p string = "nema-15-30p"

	// PowerPortTypeValueNemaDash15Dash50p captures enum value "nema-15-50p"
	PowerPortTypeValueNemaDash15Dash50p string = "nema-15-50p"

	// PowerPortTypeValueNemaDash15Dash60p captures enum value "nema-15-60p"
	PowerPortTypeValueNemaDash15Dash60p string = "nema-15-60p"

	// PowerPortTypeValueNemaDashL1Dash15p captures enum value "nema-l1-15p"
	PowerPortTypeValueNemaDashL1Dash15p string = "nema-l1-15p"

	// PowerPortTypeValueNemaDashL5Dash15p captures enum value "nema-l5-15p"
	PowerPortTypeValueNemaDashL5Dash15p string = "nema-l5-15p"

	// PowerPortTypeValueNemaDashL5Dash20p captures enum value "nema-l5-20p"
	PowerPortTypeValueNemaDashL5Dash20p string = "nema-l5-20p"

	// PowerPortTypeValueNemaDashL5Dash30p captures enum value "nema-l5-30p"
	PowerPortTypeValueNemaDashL5Dash30p string = "nema-l5-30p"

	// PowerPortTypeValueNemaDashL5Dash50p captures enum value "nema-l5-50p"
	PowerPortTypeValueNemaDashL5Dash50p string = "nema-l5-50p"

	// PowerPortTypeValueNemaDashL6Dash15p captures enum value "nema-l6-15p"
	PowerPortTypeValueNemaDashL6Dash15p string = "nema-l6-15p"

	// PowerPortTypeValueNemaDashL6Dash20p captures enum value "nema-l6-20p"
	PowerPortTypeValueNemaDashL6Dash20p string = "nema-l6-20p"

	// PowerPortTypeValueNemaDashL6Dash30p captures enum value "nema-l6-30p"
	PowerPortTypeValueNemaDashL6Dash30p string = "nema-l6-30p"

	// PowerPortTypeValueNemaDashL6Dash50p captures enum value "nema-l6-50p"
	PowerPortTypeValueNemaDashL6Dash50p string = "nema-l6-50p"

	// PowerPortTypeValueNemaDashL10Dash30p captures enum value "nema-l10-30p"
	PowerPortTypeValueNemaDashL10Dash30p string = "nema-l10-30p"

	// PowerPortTypeValueNemaDashL14Dash20p captures enum value "nema-l14-20p"
	PowerPortTypeValueNemaDashL14Dash20p string = "nema-l14-20p"

	// PowerPortTypeValueNemaDashL14Dash30p captures enum value "nema-l14-30p"
	PowerPortTypeValueNemaDashL14Dash30p string = "nema-l14-30p"

	// PowerPortTypeValueNemaDashL14Dash50p captures enum value "nema-l14-50p"
	PowerPortTypeValueNemaDashL14Dash50p string = "nema-l14-50p"

	// PowerPortTypeValueNemaDashL14Dash60p captures enum value "nema-l14-60p"
	PowerPortTypeValueNemaDashL14Dash60p string = "nema-l14-60p"

	// PowerPortTypeValueNemaDashL15Dash20p captures enum value "nema-l15-20p"
	PowerPortTypeValueNemaDashL15Dash20p string = "nema-l15-20p"

	// PowerPortTypeValueNemaDashL15Dash30p captures enum value "nema-l15-30p"
	PowerPortTypeValueNemaDashL15Dash30p string = "nema-l15-30p"

	// PowerPortTypeValueNemaDashL15Dash50p captures enum value "nema-l15-50p"
	PowerPortTypeValueNemaDashL15Dash50p string = "nema-l15-50p"

	// PowerPortTypeValueNemaDashL15Dash60p captures enum value "nema-l15-60p"
	PowerPortTypeValueNemaDashL15Dash60p string = "nema-l15-60p"

	// PowerPortTypeValueNemaDashL21Dash20p captures enum value "nema-l21-20p"
	PowerPortTypeValueNemaDashL21Dash20p string = "nema-l21-20p"

	// PowerPortTypeValueNemaDashL21Dash30p captures enum value "nema-l21-30p"
	PowerPortTypeValueNemaDashL21Dash30p string = "nema-l21-30p"

	// PowerPortTypeValueCs6361c captures enum value "cs6361c"
	PowerPortTypeValueCs6361c string = "cs6361c"

	// PowerPortTypeValueCs6365c captures enum value "cs6365c"
	PowerPortTypeValueCs6365c string = "cs6365c"

	// PowerPortTypeValueCs8165c captures enum value "cs8165c"
	PowerPortTypeValueCs8165c string = "cs8165c"

	// PowerPortTypeValueCs8265c captures enum value "cs8265c"
	PowerPortTypeValueCs8265c string = "cs8265c"

	// PowerPortTypeValueCs8365c captures enum value "cs8365c"
	PowerPortTypeValueCs8365c string = "cs8365c"

	// PowerPortTypeValueCs8465c captures enum value "cs8465c"
	PowerPortTypeValueCs8465c string = "cs8465c"

	// PowerPortTypeValueItaDashe captures enum value "ita-e"
	PowerPortTypeValueItaDashe string = "ita-e"

	// PowerPortTypeValueItaDashf captures enum value "ita-f"
	PowerPortTypeValueItaDashf string = "ita-f"

	// PowerPortTypeValueItaDashEf captures enum value "ita-ef"
	PowerPortTypeValueItaDashEf string = "ita-ef"

	// PowerPortTypeValueItaDashg captures enum value "ita-g"
	PowerPortTypeValueItaDashg string = "ita-g"

	// PowerPortTypeValueItaDashh captures enum value "ita-h"
	PowerPortTypeValueItaDashh string = "ita-h"

	// PowerPortTypeValueItaDashi captures enum value "ita-i"
	PowerPortTypeValueItaDashi string = "ita-i"

	// PowerPortTypeValueItaDashj captures enum value "ita-j"
	PowerPortTypeValueItaDashj string = "ita-j"

	// PowerPortTypeValueItaDashk captures enum value "ita-k"
	PowerPortTypeValueItaDashk string = "ita-k"

	// PowerPortTypeValueItaDashl captures enum value "ita-l"
	PowerPortTypeValueItaDashl string = "ita-l"

	// PowerPortTypeValueItaDashm captures enum value "ita-m"
	PowerPortTypeValueItaDashm string = "ita-m"

	// PowerPortTypeValueItaDashn captures enum value "ita-n"
	PowerPortTypeValueItaDashn string = "ita-n"

	// PowerPortTypeValueItaDasho captures enum value "ita-o"
	PowerPortTypeValueItaDasho string = "ita-o"

	// PowerPortTypeValueUsbDasha captures enum value "usb-a"
	PowerPortTypeValueUsbDasha string = "usb-a"

	// PowerPortTypeValueUsbDashb captures enum value "usb-b"
	PowerPortTypeValueUsbDashb string = "usb-b"

	// PowerPortTypeValueUsbDashc captures enum value "usb-c"
	PowerPortTypeValueUsbDashc string = "usb-c"

	// PowerPortTypeValueUsbDashMiniDasha captures enum value "usb-mini-a"
	PowerPortTypeValueUsbDashMiniDasha string = "usb-mini-a"

	// PowerPortTypeValueUsbDashMiniDashb captures enum value "usb-mini-b"
	PowerPortTypeValueUsbDashMiniDashb string = "usb-mini-b"

	// PowerPortTypeValueUsbDashMicroDasha captures enum value "usb-micro-a"
	PowerPortTypeValueUsbDashMicroDasha string = "usb-micro-a"

	// PowerPortTypeValueUsbDashMicroDashb captures enum value "usb-micro-b"
	PowerPortTypeValueUsbDashMicroDashb string = "usb-micro-b"

	// PowerPortTypeValueUsbDash3Dashb captures enum value "usb-3-b"
	PowerPortTypeValueUsbDash3Dashb string = "usb-3-b"

	// PowerPortTypeValueUsbDash3DashMicroDashb captures enum value "usb-3-micro-b"
	PowerPortTypeValueUsbDash3DashMicroDashb string = "usb-3-micro-b"
)

// prop value enum
func (m *PowerPortType) validateValueEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, powerPortTypeTypeValuePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *PowerPortType) validateValue(formats strfmt.Registry) error {

	if err := validate.Required("type"+"."+"value", "body", m.Value); err != nil {
		return err
	}

	// value enum
	if err := m.validateValueEnum("type"+"."+"value", "body", *m.Value); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this power port type based on context it is used
func (m *PowerPortType) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *PowerPortType) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *PowerPortType) UnmarshalBinary(b []byte) error {
	var res PowerPortType
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
