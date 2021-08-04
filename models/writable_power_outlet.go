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

// WritablePowerOutlet writable power outlet
//
// swagger:model WritablePowerOutlet
type WritablePowerOutlet struct {

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

	// Feed leg
	//
	// Phase (for three-phase feeds)
	// Enum: [A B C]
	FeedLeg string `json:"feed_leg,omitempty"`

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

	// Power port
	// Format: uuid
	PowerPort *strfmt.UUID `json:"power_port,omitempty"`

	// tags
	Tags []*NestedTag `json:"tags"`

	// Type
	//
	// Physical port type
	// Enum: [iec-60320-c5 iec-60320-c7 iec-60320-c13 iec-60320-c15 iec-60320-c19 iec-60309-p-n-e-4h iec-60309-p-n-e-6h iec-60309-p-n-e-9h iec-60309-2p-e-4h iec-60309-2p-e-6h iec-60309-2p-e-9h iec-60309-3p-e-4h iec-60309-3p-e-6h iec-60309-3p-e-9h iec-60309-3p-n-e-4h iec-60309-3p-n-e-6h iec-60309-3p-n-e-9h nema-1-15r nema-5-15r nema-5-20r nema-5-30r nema-5-50r nema-6-15r nema-6-20r nema-6-30r nema-6-50r nema-10-30r nema-10-50r nema-14-20r nema-14-30r nema-14-50r nema-14-60r nema-15-15r nema-15-20r nema-15-30r nema-15-50r nema-15-60r nema-l1-15r nema-l5-15r nema-l5-20r nema-l5-30r nema-l5-50r nema-l6-15r nema-l6-20r nema-l6-30r nema-l6-50r nema-l10-30r nema-l14-20r nema-l14-30r nema-l14-50r nema-l14-60r nema-l15-20r nema-l15-30r nema-l15-50r nema-l15-60r nema-l21-20r nema-l21-30r CS6360C CS6364C CS8164C CS8264C CS8364C CS8464C ita-e ita-f ita-g ita-h ita-i ita-j ita-k ita-l ita-m ita-n ita-o usb-a usb-micro-b usb-c hdot-cx]
	Type string `json:"type,omitempty"`

	// Url
	// Read Only: true
	// Format: uri
	URL strfmt.URI `json:"url,omitempty"`
}

// Validate validates this writable power outlet
func (m *WritablePowerOutlet) Validate(formats strfmt.Registry) error {
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

func (m *WritablePowerOutlet) validateCable(formats strfmt.Registry) error {
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

func (m *WritablePowerOutlet) validateDescription(formats strfmt.Registry) error {
	if swag.IsZero(m.Description) { // not required
		return nil
	}

	if err := validate.MaxLength("description", "body", m.Description, 200); err != nil {
		return err
	}

	return nil
}

func (m *WritablePowerOutlet) validateDevice(formats strfmt.Registry) error {

	if err := validate.Required("device", "body", m.Device); err != nil {
		return err
	}

	if err := validate.FormatOf("device", "body", "uuid", m.Device.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *WritablePowerOutlet) validateDisplay(formats strfmt.Registry) error {
	if swag.IsZero(m.Display) { // not required
		return nil
	}

	if err := validate.MinLength("display", "body", m.Display, 1); err != nil {
		return err
	}

	return nil
}

var writablePowerOutletTypeFeedLegPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["A","B","C"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		writablePowerOutletTypeFeedLegPropEnum = append(writablePowerOutletTypeFeedLegPropEnum, v)
	}
}

const (

	// WritablePowerOutletFeedLegA captures enum value "A"
	WritablePowerOutletFeedLegA string = "A"

	// WritablePowerOutletFeedLegB captures enum value "B"
	WritablePowerOutletFeedLegB string = "B"

	// WritablePowerOutletFeedLegC captures enum value "C"
	WritablePowerOutletFeedLegC string = "C"
)

// prop value enum
func (m *WritablePowerOutlet) validateFeedLegEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, writablePowerOutletTypeFeedLegPropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *WritablePowerOutlet) validateFeedLeg(formats strfmt.Registry) error {
	if swag.IsZero(m.FeedLeg) { // not required
		return nil
	}

	// value enum
	if err := m.validateFeedLegEnum("feed_leg", "body", m.FeedLeg); err != nil {
		return err
	}

	return nil
}

func (m *WritablePowerOutlet) validateID(formats strfmt.Registry) error {
	if swag.IsZero(m.ID) { // not required
		return nil
	}

	if err := validate.FormatOf("id", "body", "uuid", m.ID.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *WritablePowerOutlet) validateLabel(formats strfmt.Registry) error {
	if swag.IsZero(m.Label) { // not required
		return nil
	}

	if err := validate.MaxLength("label", "body", m.Label, 64); err != nil {
		return err
	}

	return nil
}

func (m *WritablePowerOutlet) validateName(formats strfmt.Registry) error {

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

func (m *WritablePowerOutlet) validatePowerPort(formats strfmt.Registry) error {
	if swag.IsZero(m.PowerPort) { // not required
		return nil
	}

	if err := validate.FormatOf("power_port", "body", "uuid", m.PowerPort.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *WritablePowerOutlet) validateTags(formats strfmt.Registry) error {
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

var writablePowerOutletTypeTypePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["iec-60320-c5","iec-60320-c7","iec-60320-c13","iec-60320-c15","iec-60320-c19","iec-60309-p-n-e-4h","iec-60309-p-n-e-6h","iec-60309-p-n-e-9h","iec-60309-2p-e-4h","iec-60309-2p-e-6h","iec-60309-2p-e-9h","iec-60309-3p-e-4h","iec-60309-3p-e-6h","iec-60309-3p-e-9h","iec-60309-3p-n-e-4h","iec-60309-3p-n-e-6h","iec-60309-3p-n-e-9h","nema-1-15r","nema-5-15r","nema-5-20r","nema-5-30r","nema-5-50r","nema-6-15r","nema-6-20r","nema-6-30r","nema-6-50r","nema-10-30r","nema-10-50r","nema-14-20r","nema-14-30r","nema-14-50r","nema-14-60r","nema-15-15r","nema-15-20r","nema-15-30r","nema-15-50r","nema-15-60r","nema-l1-15r","nema-l5-15r","nema-l5-20r","nema-l5-30r","nema-l5-50r","nema-l6-15r","nema-l6-20r","nema-l6-30r","nema-l6-50r","nema-l10-30r","nema-l14-20r","nema-l14-30r","nema-l14-50r","nema-l14-60r","nema-l15-20r","nema-l15-30r","nema-l15-50r","nema-l15-60r","nema-l21-20r","nema-l21-30r","CS6360C","CS6364C","CS8164C","CS8264C","CS8364C","CS8464C","ita-e","ita-f","ita-g","ita-h","ita-i","ita-j","ita-k","ita-l","ita-m","ita-n","ita-o","usb-a","usb-micro-b","usb-c","hdot-cx"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		writablePowerOutletTypeTypePropEnum = append(writablePowerOutletTypeTypePropEnum, v)
	}
}

const (

	// WritablePowerOutletTypeIecDash60320DashC5 captures enum value "iec-60320-c5"
	WritablePowerOutletTypeIecDash60320DashC5 string = "iec-60320-c5"

	// WritablePowerOutletTypeIecDash60320DashC7 captures enum value "iec-60320-c7"
	WritablePowerOutletTypeIecDash60320DashC7 string = "iec-60320-c7"

	// WritablePowerOutletTypeIecDash60320DashC13 captures enum value "iec-60320-c13"
	WritablePowerOutletTypeIecDash60320DashC13 string = "iec-60320-c13"

	// WritablePowerOutletTypeIecDash60320DashC15 captures enum value "iec-60320-c15"
	WritablePowerOutletTypeIecDash60320DashC15 string = "iec-60320-c15"

	// WritablePowerOutletTypeIecDash60320DashC19 captures enum value "iec-60320-c19"
	WritablePowerOutletTypeIecDash60320DashC19 string = "iec-60320-c19"

	// WritablePowerOutletTypeIecDash60309DashpDashnDasheDash4h captures enum value "iec-60309-p-n-e-4h"
	WritablePowerOutletTypeIecDash60309DashpDashnDasheDash4h string = "iec-60309-p-n-e-4h"

	// WritablePowerOutletTypeIecDash60309DashpDashnDasheDash6h captures enum value "iec-60309-p-n-e-6h"
	WritablePowerOutletTypeIecDash60309DashpDashnDasheDash6h string = "iec-60309-p-n-e-6h"

	// WritablePowerOutletTypeIecDash60309DashpDashnDasheDash9h captures enum value "iec-60309-p-n-e-9h"
	WritablePowerOutletTypeIecDash60309DashpDashnDasheDash9h string = "iec-60309-p-n-e-9h"

	// WritablePowerOutletTypeIecDash60309Dash2pDasheDash4h captures enum value "iec-60309-2p-e-4h"
	WritablePowerOutletTypeIecDash60309Dash2pDasheDash4h string = "iec-60309-2p-e-4h"

	// WritablePowerOutletTypeIecDash60309Dash2pDasheDash6h captures enum value "iec-60309-2p-e-6h"
	WritablePowerOutletTypeIecDash60309Dash2pDasheDash6h string = "iec-60309-2p-e-6h"

	// WritablePowerOutletTypeIecDash60309Dash2pDasheDash9h captures enum value "iec-60309-2p-e-9h"
	WritablePowerOutletTypeIecDash60309Dash2pDasheDash9h string = "iec-60309-2p-e-9h"

	// WritablePowerOutletTypeIecDash60309Dash3pDasheDash4h captures enum value "iec-60309-3p-e-4h"
	WritablePowerOutletTypeIecDash60309Dash3pDasheDash4h string = "iec-60309-3p-e-4h"

	// WritablePowerOutletTypeIecDash60309Dash3pDasheDash6h captures enum value "iec-60309-3p-e-6h"
	WritablePowerOutletTypeIecDash60309Dash3pDasheDash6h string = "iec-60309-3p-e-6h"

	// WritablePowerOutletTypeIecDash60309Dash3pDasheDash9h captures enum value "iec-60309-3p-e-9h"
	WritablePowerOutletTypeIecDash60309Dash3pDasheDash9h string = "iec-60309-3p-e-9h"

	// WritablePowerOutletTypeIecDash60309Dash3pDashnDasheDash4h captures enum value "iec-60309-3p-n-e-4h"
	WritablePowerOutletTypeIecDash60309Dash3pDashnDasheDash4h string = "iec-60309-3p-n-e-4h"

	// WritablePowerOutletTypeIecDash60309Dash3pDashnDasheDash6h captures enum value "iec-60309-3p-n-e-6h"
	WritablePowerOutletTypeIecDash60309Dash3pDashnDasheDash6h string = "iec-60309-3p-n-e-6h"

	// WritablePowerOutletTypeIecDash60309Dash3pDashnDasheDash9h captures enum value "iec-60309-3p-n-e-9h"
	WritablePowerOutletTypeIecDash60309Dash3pDashnDasheDash9h string = "iec-60309-3p-n-e-9h"

	// WritablePowerOutletTypeNemaDash1Dash15r captures enum value "nema-1-15r"
	WritablePowerOutletTypeNemaDash1Dash15r string = "nema-1-15r"

	// WritablePowerOutletTypeNemaDash5Dash15r captures enum value "nema-5-15r"
	WritablePowerOutletTypeNemaDash5Dash15r string = "nema-5-15r"

	// WritablePowerOutletTypeNemaDash5Dash20r captures enum value "nema-5-20r"
	WritablePowerOutletTypeNemaDash5Dash20r string = "nema-5-20r"

	// WritablePowerOutletTypeNemaDash5Dash30r captures enum value "nema-5-30r"
	WritablePowerOutletTypeNemaDash5Dash30r string = "nema-5-30r"

	// WritablePowerOutletTypeNemaDash5Dash50r captures enum value "nema-5-50r"
	WritablePowerOutletTypeNemaDash5Dash50r string = "nema-5-50r"

	// WritablePowerOutletTypeNemaDash6Dash15r captures enum value "nema-6-15r"
	WritablePowerOutletTypeNemaDash6Dash15r string = "nema-6-15r"

	// WritablePowerOutletTypeNemaDash6Dash20r captures enum value "nema-6-20r"
	WritablePowerOutletTypeNemaDash6Dash20r string = "nema-6-20r"

	// WritablePowerOutletTypeNemaDash6Dash30r captures enum value "nema-6-30r"
	WritablePowerOutletTypeNemaDash6Dash30r string = "nema-6-30r"

	// WritablePowerOutletTypeNemaDash6Dash50r captures enum value "nema-6-50r"
	WritablePowerOutletTypeNemaDash6Dash50r string = "nema-6-50r"

	// WritablePowerOutletTypeNemaDash10Dash30r captures enum value "nema-10-30r"
	WritablePowerOutletTypeNemaDash10Dash30r string = "nema-10-30r"

	// WritablePowerOutletTypeNemaDash10Dash50r captures enum value "nema-10-50r"
	WritablePowerOutletTypeNemaDash10Dash50r string = "nema-10-50r"

	// WritablePowerOutletTypeNemaDash14Dash20r captures enum value "nema-14-20r"
	WritablePowerOutletTypeNemaDash14Dash20r string = "nema-14-20r"

	// WritablePowerOutletTypeNemaDash14Dash30r captures enum value "nema-14-30r"
	WritablePowerOutletTypeNemaDash14Dash30r string = "nema-14-30r"

	// WritablePowerOutletTypeNemaDash14Dash50r captures enum value "nema-14-50r"
	WritablePowerOutletTypeNemaDash14Dash50r string = "nema-14-50r"

	// WritablePowerOutletTypeNemaDash14Dash60r captures enum value "nema-14-60r"
	WritablePowerOutletTypeNemaDash14Dash60r string = "nema-14-60r"

	// WritablePowerOutletTypeNemaDash15Dash15r captures enum value "nema-15-15r"
	WritablePowerOutletTypeNemaDash15Dash15r string = "nema-15-15r"

	// WritablePowerOutletTypeNemaDash15Dash20r captures enum value "nema-15-20r"
	WritablePowerOutletTypeNemaDash15Dash20r string = "nema-15-20r"

	// WritablePowerOutletTypeNemaDash15Dash30r captures enum value "nema-15-30r"
	WritablePowerOutletTypeNemaDash15Dash30r string = "nema-15-30r"

	// WritablePowerOutletTypeNemaDash15Dash50r captures enum value "nema-15-50r"
	WritablePowerOutletTypeNemaDash15Dash50r string = "nema-15-50r"

	// WritablePowerOutletTypeNemaDash15Dash60r captures enum value "nema-15-60r"
	WritablePowerOutletTypeNemaDash15Dash60r string = "nema-15-60r"

	// WritablePowerOutletTypeNemaDashL1Dash15r captures enum value "nema-l1-15r"
	WritablePowerOutletTypeNemaDashL1Dash15r string = "nema-l1-15r"

	// WritablePowerOutletTypeNemaDashL5Dash15r captures enum value "nema-l5-15r"
	WritablePowerOutletTypeNemaDashL5Dash15r string = "nema-l5-15r"

	// WritablePowerOutletTypeNemaDashL5Dash20r captures enum value "nema-l5-20r"
	WritablePowerOutletTypeNemaDashL5Dash20r string = "nema-l5-20r"

	// WritablePowerOutletTypeNemaDashL5Dash30r captures enum value "nema-l5-30r"
	WritablePowerOutletTypeNemaDashL5Dash30r string = "nema-l5-30r"

	// WritablePowerOutletTypeNemaDashL5Dash50r captures enum value "nema-l5-50r"
	WritablePowerOutletTypeNemaDashL5Dash50r string = "nema-l5-50r"

	// WritablePowerOutletTypeNemaDashL6Dash15r captures enum value "nema-l6-15r"
	WritablePowerOutletTypeNemaDashL6Dash15r string = "nema-l6-15r"

	// WritablePowerOutletTypeNemaDashL6Dash20r captures enum value "nema-l6-20r"
	WritablePowerOutletTypeNemaDashL6Dash20r string = "nema-l6-20r"

	// WritablePowerOutletTypeNemaDashL6Dash30r captures enum value "nema-l6-30r"
	WritablePowerOutletTypeNemaDashL6Dash30r string = "nema-l6-30r"

	// WritablePowerOutletTypeNemaDashL6Dash50r captures enum value "nema-l6-50r"
	WritablePowerOutletTypeNemaDashL6Dash50r string = "nema-l6-50r"

	// WritablePowerOutletTypeNemaDashL10Dash30r captures enum value "nema-l10-30r"
	WritablePowerOutletTypeNemaDashL10Dash30r string = "nema-l10-30r"

	// WritablePowerOutletTypeNemaDashL14Dash20r captures enum value "nema-l14-20r"
	WritablePowerOutletTypeNemaDashL14Dash20r string = "nema-l14-20r"

	// WritablePowerOutletTypeNemaDashL14Dash30r captures enum value "nema-l14-30r"
	WritablePowerOutletTypeNemaDashL14Dash30r string = "nema-l14-30r"

	// WritablePowerOutletTypeNemaDashL14Dash50r captures enum value "nema-l14-50r"
	WritablePowerOutletTypeNemaDashL14Dash50r string = "nema-l14-50r"

	// WritablePowerOutletTypeNemaDashL14Dash60r captures enum value "nema-l14-60r"
	WritablePowerOutletTypeNemaDashL14Dash60r string = "nema-l14-60r"

	// WritablePowerOutletTypeNemaDashL15Dash20r captures enum value "nema-l15-20r"
	WritablePowerOutletTypeNemaDashL15Dash20r string = "nema-l15-20r"

	// WritablePowerOutletTypeNemaDashL15Dash30r captures enum value "nema-l15-30r"
	WritablePowerOutletTypeNemaDashL15Dash30r string = "nema-l15-30r"

	// WritablePowerOutletTypeNemaDashL15Dash50r captures enum value "nema-l15-50r"
	WritablePowerOutletTypeNemaDashL15Dash50r string = "nema-l15-50r"

	// WritablePowerOutletTypeNemaDashL15Dash60r captures enum value "nema-l15-60r"
	WritablePowerOutletTypeNemaDashL15Dash60r string = "nema-l15-60r"

	// WritablePowerOutletTypeNemaDashL21Dash20r captures enum value "nema-l21-20r"
	WritablePowerOutletTypeNemaDashL21Dash20r string = "nema-l21-20r"

	// WritablePowerOutletTypeNemaDashL21Dash30r captures enum value "nema-l21-30r"
	WritablePowerOutletTypeNemaDashL21Dash30r string = "nema-l21-30r"

	// WritablePowerOutletTypeCS6360C captures enum value "CS6360C"
	WritablePowerOutletTypeCS6360C string = "CS6360C"

	// WritablePowerOutletTypeCS6364C captures enum value "CS6364C"
	WritablePowerOutletTypeCS6364C string = "CS6364C"

	// WritablePowerOutletTypeCS8164C captures enum value "CS8164C"
	WritablePowerOutletTypeCS8164C string = "CS8164C"

	// WritablePowerOutletTypeCS8264C captures enum value "CS8264C"
	WritablePowerOutletTypeCS8264C string = "CS8264C"

	// WritablePowerOutletTypeCS8364C captures enum value "CS8364C"
	WritablePowerOutletTypeCS8364C string = "CS8364C"

	// WritablePowerOutletTypeCS8464C captures enum value "CS8464C"
	WritablePowerOutletTypeCS8464C string = "CS8464C"

	// WritablePowerOutletTypeItaDashe captures enum value "ita-e"
	WritablePowerOutletTypeItaDashe string = "ita-e"

	// WritablePowerOutletTypeItaDashf captures enum value "ita-f"
	WritablePowerOutletTypeItaDashf string = "ita-f"

	// WritablePowerOutletTypeItaDashg captures enum value "ita-g"
	WritablePowerOutletTypeItaDashg string = "ita-g"

	// WritablePowerOutletTypeItaDashh captures enum value "ita-h"
	WritablePowerOutletTypeItaDashh string = "ita-h"

	// WritablePowerOutletTypeItaDashi captures enum value "ita-i"
	WritablePowerOutletTypeItaDashi string = "ita-i"

	// WritablePowerOutletTypeItaDashj captures enum value "ita-j"
	WritablePowerOutletTypeItaDashj string = "ita-j"

	// WritablePowerOutletTypeItaDashk captures enum value "ita-k"
	WritablePowerOutletTypeItaDashk string = "ita-k"

	// WritablePowerOutletTypeItaDashl captures enum value "ita-l"
	WritablePowerOutletTypeItaDashl string = "ita-l"

	// WritablePowerOutletTypeItaDashm captures enum value "ita-m"
	WritablePowerOutletTypeItaDashm string = "ita-m"

	// WritablePowerOutletTypeItaDashn captures enum value "ita-n"
	WritablePowerOutletTypeItaDashn string = "ita-n"

	// WritablePowerOutletTypeItaDasho captures enum value "ita-o"
	WritablePowerOutletTypeItaDasho string = "ita-o"

	// WritablePowerOutletTypeUsbDasha captures enum value "usb-a"
	WritablePowerOutletTypeUsbDasha string = "usb-a"

	// WritablePowerOutletTypeUsbDashMicroDashb captures enum value "usb-micro-b"
	WritablePowerOutletTypeUsbDashMicroDashb string = "usb-micro-b"

	// WritablePowerOutletTypeUsbDashc captures enum value "usb-c"
	WritablePowerOutletTypeUsbDashc string = "usb-c"

	// WritablePowerOutletTypeHdotDashCx captures enum value "hdot-cx"
	WritablePowerOutletTypeHdotDashCx string = "hdot-cx"
)

// prop value enum
func (m *WritablePowerOutlet) validateTypeEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, writablePowerOutletTypeTypePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *WritablePowerOutlet) validateType(formats strfmt.Registry) error {
	if swag.IsZero(m.Type) { // not required
		return nil
	}

	// value enum
	if err := m.validateTypeEnum("type", "body", m.Type); err != nil {
		return err
	}

	return nil
}

func (m *WritablePowerOutlet) validateURL(formats strfmt.Registry) error {
	if swag.IsZero(m.URL) { // not required
		return nil
	}

	if err := validate.FormatOf("url", "body", "uri", m.URL.String(), formats); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this writable power outlet based on the context it is used
func (m *WritablePowerOutlet) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
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

	if err := m.contextValidateConnectedEndpoint(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateConnectedEndpointReachable(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateConnectedEndpointType(ctx, formats); err != nil {
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

func (m *WritablePowerOutlet) contextValidateCable(ctx context.Context, formats strfmt.Registry) error {

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

func (m *WritablePowerOutlet) contextValidateCablePeer(ctx context.Context, formats strfmt.Registry) error {

	return nil
}

func (m *WritablePowerOutlet) contextValidateCablePeerType(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "cable_peer_type", "body", string(m.CablePeerType)); err != nil {
		return err
	}

	return nil
}

func (m *WritablePowerOutlet) contextValidateComputedFields(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "computed_fields", "body", string(m.ComputedFields)); err != nil {
		return err
	}

	return nil
}

func (m *WritablePowerOutlet) contextValidateConnectedEndpoint(ctx context.Context, formats strfmt.Registry) error {

	return nil
}

func (m *WritablePowerOutlet) contextValidateConnectedEndpointReachable(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "connected_endpoint_reachable", "body", m.ConnectedEndpointReachable); err != nil {
		return err
	}

	return nil
}

func (m *WritablePowerOutlet) contextValidateConnectedEndpointType(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "connected_endpoint_type", "body", string(m.ConnectedEndpointType)); err != nil {
		return err
	}

	return nil
}

func (m *WritablePowerOutlet) contextValidateDisplay(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "display", "body", string(m.Display)); err != nil {
		return err
	}

	return nil
}

func (m *WritablePowerOutlet) contextValidateID(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "id", "body", strfmt.UUID(m.ID)); err != nil {
		return err
	}

	return nil
}

func (m *WritablePowerOutlet) contextValidateTags(ctx context.Context, formats strfmt.Registry) error {

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

func (m *WritablePowerOutlet) contextValidateURL(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "url", "body", strfmt.URI(m.URL)); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *WritablePowerOutlet) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *WritablePowerOutlet) UnmarshalBinary(b []byte) error {
	var res WritablePowerOutlet
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
