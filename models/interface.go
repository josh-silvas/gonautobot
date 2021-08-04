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

// Interface interface
//
// swagger:model Interface
type Interface struct {

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

	// Count ipaddresses
	// Read Only: true
	CountIpaddresses int64 `json:"count_ipaddresses,omitempty"`

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

	// Enabled
	Enabled bool `json:"enabled,omitempty"`

	// Id
	// Read Only: true
	// Format: uuid
	ID strfmt.UUID `json:"id,omitempty"`

	// Label
	//
	// Physical label
	// Max Length: 64
	Label string `json:"label,omitempty"`

	// lag
	Lag *NestedInterface `json:"lag,omitempty"`

	// MAC Address
	// Max Length: 18
	MacAddress *string `json:"mac_address,omitempty"`

	// Management only
	//
	// This interface is used only for out-of-band management
	MgmtOnly bool `json:"mgmt_only,omitempty"`

	// mode
	Mode *InterfaceMode `json:"mode,omitempty"`

	// MTU
	// Maximum: 65536
	// Minimum: 1
	Mtu *int64 `json:"mtu,omitempty"`

	// Name
	// Required: true
	// Max Length: 64
	// Min Length: 1
	Name *string `json:"name"`

	// tagged vlans
	// Unique: true
	TaggedVlans []strfmt.UUID `json:"tagged_vlans"`

	// tags
	Tags []*NestedTag `json:"tags"`

	// type
	// Required: true
	Type *InterfaceType `json:"type"`

	// untagged vlan
	UntaggedVlan *NestedVLAN `json:"untagged_vlan,omitempty"`

	// Url
	// Read Only: true
	// Format: uri
	URL strfmt.URI `json:"url,omitempty"`
}

// Validate validates this interface
func (m *Interface) Validate(formats strfmt.Registry) error {
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

	if err := m.validateLag(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateMacAddress(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateMode(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateMtu(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateName(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTaggedVlans(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTags(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateType(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateUntaggedVlan(formats); err != nil {
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

func (m *Interface) validateCable(formats strfmt.Registry) error {
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

func (m *Interface) validateDescription(formats strfmt.Registry) error {
	if swag.IsZero(m.Description) { // not required
		return nil
	}

	if err := validate.MaxLength("description", "body", m.Description, 200); err != nil {
		return err
	}

	return nil
}

func (m *Interface) validateDevice(formats strfmt.Registry) error {

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

func (m *Interface) validateDisplay(formats strfmt.Registry) error {
	if swag.IsZero(m.Display) { // not required
		return nil
	}

	if err := validate.MinLength("display", "body", m.Display, 1); err != nil {
		return err
	}

	return nil
}

func (m *Interface) validateID(formats strfmt.Registry) error {
	if swag.IsZero(m.ID) { // not required
		return nil
	}

	if err := validate.FormatOf("id", "body", "uuid", m.ID.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *Interface) validateLabel(formats strfmt.Registry) error {
	if swag.IsZero(m.Label) { // not required
		return nil
	}

	if err := validate.MaxLength("label", "body", m.Label, 64); err != nil {
		return err
	}

	return nil
}

func (m *Interface) validateLag(formats strfmt.Registry) error {
	if swag.IsZero(m.Lag) { // not required
		return nil
	}

	if m.Lag != nil {
		if err := m.Lag.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("lag")
			}
			return err
		}
	}

	return nil
}

func (m *Interface) validateMacAddress(formats strfmt.Registry) error {
	if swag.IsZero(m.MacAddress) { // not required
		return nil
	}

	if err := validate.MaxLength("mac_address", "body", *m.MacAddress, 18); err != nil {
		return err
	}

	return nil
}

func (m *Interface) validateMode(formats strfmt.Registry) error {
	if swag.IsZero(m.Mode) { // not required
		return nil
	}

	if m.Mode != nil {
		if err := m.Mode.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("mode")
			}
			return err
		}
	}

	return nil
}

func (m *Interface) validateMtu(formats strfmt.Registry) error {
	if swag.IsZero(m.Mtu) { // not required
		return nil
	}

	if err := validate.MinimumInt("mtu", "body", *m.Mtu, 1, false); err != nil {
		return err
	}

	if err := validate.MaximumInt("mtu", "body", *m.Mtu, 65536, false); err != nil {
		return err
	}

	return nil
}

func (m *Interface) validateName(formats strfmt.Registry) error {

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

func (m *Interface) validateTaggedVlans(formats strfmt.Registry) error {
	if swag.IsZero(m.TaggedVlans) { // not required
		return nil
	}

	if err := validate.UniqueItems("tagged_vlans", "body", m.TaggedVlans); err != nil {
		return err
	}

	for i := 0; i < len(m.TaggedVlans); i++ {

		if err := validate.FormatOf("tagged_vlans"+"."+strconv.Itoa(i), "body", "uuid", m.TaggedVlans[i].String(), formats); err != nil {
			return err
		}

	}

	return nil
}

func (m *Interface) validateTags(formats strfmt.Registry) error {
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

func (m *Interface) validateType(formats strfmt.Registry) error {

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

func (m *Interface) validateUntaggedVlan(formats strfmt.Registry) error {
	if swag.IsZero(m.UntaggedVlan) { // not required
		return nil
	}

	if m.UntaggedVlan != nil {
		if err := m.UntaggedVlan.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("untagged_vlan")
			}
			return err
		}
	}

	return nil
}

func (m *Interface) validateURL(formats strfmt.Registry) error {
	if swag.IsZero(m.URL) { // not required
		return nil
	}

	if err := validate.FormatOf("url", "body", "uri", m.URL.String(), formats); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this interface based on the context it is used
func (m *Interface) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
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

	if err := m.contextValidateCountIpaddresses(ctx, formats); err != nil {
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

	if err := m.contextValidateLag(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateMode(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateTags(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateType(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateUntaggedVlan(ctx, formats); err != nil {
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

func (m *Interface) contextValidateCable(ctx context.Context, formats strfmt.Registry) error {

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

func (m *Interface) contextValidateCablePeer(ctx context.Context, formats strfmt.Registry) error {

	return nil
}

func (m *Interface) contextValidateCablePeerType(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "cable_peer_type", "body", string(m.CablePeerType)); err != nil {
		return err
	}

	return nil
}

func (m *Interface) contextValidateConnectedEndpoint(ctx context.Context, formats strfmt.Registry) error {

	return nil
}

func (m *Interface) contextValidateConnectedEndpointReachable(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "connected_endpoint_reachable", "body", m.ConnectedEndpointReachable); err != nil {
		return err
	}

	return nil
}

func (m *Interface) contextValidateConnectedEndpointType(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "connected_endpoint_type", "body", string(m.ConnectedEndpointType)); err != nil {
		return err
	}

	return nil
}

func (m *Interface) contextValidateCountIpaddresses(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "count_ipaddresses", "body", int64(m.CountIpaddresses)); err != nil {
		return err
	}

	return nil
}

func (m *Interface) contextValidateDevice(ctx context.Context, formats strfmt.Registry) error {

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

func (m *Interface) contextValidateDisplay(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "display", "body", string(m.Display)); err != nil {
		return err
	}

	return nil
}

func (m *Interface) contextValidateID(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "id", "body", strfmt.UUID(m.ID)); err != nil {
		return err
	}

	return nil
}

func (m *Interface) contextValidateLag(ctx context.Context, formats strfmt.Registry) error {

	if m.Lag != nil {
		if err := m.Lag.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("lag")
			}
			return err
		}
	}

	return nil
}

func (m *Interface) contextValidateMode(ctx context.Context, formats strfmt.Registry) error {

	if m.Mode != nil {
		if err := m.Mode.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("mode")
			}
			return err
		}
	}

	return nil
}

func (m *Interface) contextValidateTags(ctx context.Context, formats strfmt.Registry) error {

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

func (m *Interface) contextValidateType(ctx context.Context, formats strfmt.Registry) error {

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

func (m *Interface) contextValidateUntaggedVlan(ctx context.Context, formats strfmt.Registry) error {

	if m.UntaggedVlan != nil {
		if err := m.UntaggedVlan.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("untagged_vlan")
			}
			return err
		}
	}

	return nil
}

func (m *Interface) contextValidateURL(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "url", "body", strfmt.URI(m.URL)); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *Interface) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Interface) UnmarshalBinary(b []byte) error {
	var res Interface
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// InterfaceMode Mode
//
// swagger:model InterfaceMode
type InterfaceMode struct {

	// label
	// Required: true
	// Enum: [Access Tagged Tagged (All)]
	Label *string `json:"label"`

	// value
	// Required: true
	// Enum: [access tagged tagged-all]
	Value *string `json:"value"`
}

// Validate validates this interface mode
func (m *InterfaceMode) Validate(formats strfmt.Registry) error {
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

var interfaceModeTypeLabelPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["Access","Tagged","Tagged (All)"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		interfaceModeTypeLabelPropEnum = append(interfaceModeTypeLabelPropEnum, v)
	}
}

const (

	// InterfaceModeLabelAccess captures enum value "Access"
	InterfaceModeLabelAccess string = "Access"

	// InterfaceModeLabelTagged captures enum value "Tagged"
	InterfaceModeLabelTagged string = "Tagged"

	// InterfaceModeLabelTaggedAll captures enum value "Tagged (All)"
	InterfaceModeLabelTaggedAll string = "Tagged (All)"
)

// prop value enum
func (m *InterfaceMode) validateLabelEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, interfaceModeTypeLabelPropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *InterfaceMode) validateLabel(formats strfmt.Registry) error {

	if err := validate.Required("mode"+"."+"label", "body", m.Label); err != nil {
		return err
	}

	// value enum
	if err := m.validateLabelEnum("mode"+"."+"label", "body", *m.Label); err != nil {
		return err
	}

	return nil
}

var interfaceModeTypeValuePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["access","tagged","tagged-all"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		interfaceModeTypeValuePropEnum = append(interfaceModeTypeValuePropEnum, v)
	}
}

const (

	// InterfaceModeValueAccess captures enum value "access"
	InterfaceModeValueAccess string = "access"

	// InterfaceModeValueTagged captures enum value "tagged"
	InterfaceModeValueTagged string = "tagged"

	// InterfaceModeValueTaggedDashAll captures enum value "tagged-all"
	InterfaceModeValueTaggedDashAll string = "tagged-all"
)

// prop value enum
func (m *InterfaceMode) validateValueEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, interfaceModeTypeValuePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *InterfaceMode) validateValue(formats strfmt.Registry) error {

	if err := validate.Required("mode"+"."+"value", "body", m.Value); err != nil {
		return err
	}

	// value enum
	if err := m.validateValueEnum("mode"+"."+"value", "body", *m.Value); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this interface mode based on context it is used
func (m *InterfaceMode) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *InterfaceMode) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *InterfaceMode) UnmarshalBinary(b []byte) error {
	var res InterfaceMode
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// InterfaceType Type
//
// swagger:model InterfaceType
type InterfaceType struct {

	// label
	// Required: true
	// Enum: [Virtual Link Aggregation Group (LAG) 100BASE-TX (10/100ME) 1000BASE-T (1GE) 2.5GBASE-T (2.5GE) 5GBASE-T (5GE) 10GBASE-T (10GE) 10GBASE-CX4 (10GE) GBIC (1GE) SFP (1GE) SFP+ (10GE) XFP (10GE) XENPAK (10GE) X2 (10GE) SFP28 (25GE) QSFP+ (40GE) QSFP28 (50GE) CFP (100GE) CFP2 (100GE) CFP2 (200GE) CFP4 (100GE) Cisco CPAK (100GE) QSFP28 (100GE) QSFP56 (200GE) QSFP-DD (400GE) OSFP (400GE) IEEE 802.11a IEEE 802.11b/g IEEE 802.11n IEEE 802.11ac IEEE 802.11ad IEEE 802.11ax GSM CDMA LTE OC-3/STM-1 OC-12/STM-4 OC-48/STM-16 OC-192/STM-64 OC-768/STM-256 OC-1920/STM-640 OC-3840/STM-1234 SFP (1GFC) SFP (2GFC) SFP (4GFC) SFP+ (8GFC) SFP+ (16GFC) SFP28 (32GFC) QSFP+ (64GFC) QSFP28 (128GFC) SDR (2 Gbps) DDR (4 Gbps) QDR (8 Gbps) FDR10 (10 Gbps) FDR (13.5 Gbps) EDR (25 Gbps) HDR (50 Gbps) NDR (100 Gbps) XDR (250 Gbps) T1 (1.544 Mbps) E1 (2.048 Mbps) T3 (45 Mbps) E3 (34 Mbps) Cisco StackWise Cisco StackWise Plus Cisco FlexStack Cisco FlexStack Plus Juniper VCP Extreme SummitStack Extreme SummitStack-128 Extreme SummitStack-256 Extreme SummitStack-512 Other]
	Label *string `json:"label"`

	// value
	// Required: true
	// Enum: [virtual lag 100base-tx 1000base-t 2.5gbase-t 5gbase-t 10gbase-t 10gbase-cx4 1000base-x-gbic 1000base-x-sfp 10gbase-x-sfpp 10gbase-x-xfp 10gbase-x-xenpak 10gbase-x-x2 25gbase-x-sfp28 40gbase-x-qsfpp 50gbase-x-sfp28 100gbase-x-cfp 100gbase-x-cfp2 200gbase-x-cfp2 100gbase-x-cfp4 100gbase-x-cpak 100gbase-x-qsfp28 200gbase-x-qsfp56 400gbase-x-qsfpdd 400gbase-x-osfp ieee802.11a ieee802.11g ieee802.11n ieee802.11ac ieee802.11ad ieee802.11ax gsm cdma lte sonet-oc3 sonet-oc12 sonet-oc48 sonet-oc192 sonet-oc768 sonet-oc1920 sonet-oc3840 1gfc-sfp 2gfc-sfp 4gfc-sfp 8gfc-sfpp 16gfc-sfpp 32gfc-sfp28 64gfc-qsfpp 128gfc-sfp28 infiniband-sdr infiniband-ddr infiniband-qdr infiniband-fdr10 infiniband-fdr infiniband-edr infiniband-hdr infiniband-ndr infiniband-xdr t1 e1 t3 e3 cisco-stackwise cisco-stackwise-plus cisco-flexstack cisco-flexstack-plus juniper-vcp extreme-summitstack extreme-summitstack-128 extreme-summitstack-256 extreme-summitstack-512 other]
	Value *string `json:"value"`
}

// Validate validates this interface type
func (m *InterfaceType) Validate(formats strfmt.Registry) error {
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

var interfaceTypeTypeLabelPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["Virtual","Link Aggregation Group (LAG)","100BASE-TX (10/100ME)","1000BASE-T (1GE)","2.5GBASE-T (2.5GE)","5GBASE-T (5GE)","10GBASE-T (10GE)","10GBASE-CX4 (10GE)","GBIC (1GE)","SFP (1GE)","SFP+ (10GE)","XFP (10GE)","XENPAK (10GE)","X2 (10GE)","SFP28 (25GE)","QSFP+ (40GE)","QSFP28 (50GE)","CFP (100GE)","CFP2 (100GE)","CFP2 (200GE)","CFP4 (100GE)","Cisco CPAK (100GE)","QSFP28 (100GE)","QSFP56 (200GE)","QSFP-DD (400GE)","OSFP (400GE)","IEEE 802.11a","IEEE 802.11b/g","IEEE 802.11n","IEEE 802.11ac","IEEE 802.11ad","IEEE 802.11ax","GSM","CDMA","LTE","OC-3/STM-1","OC-12/STM-4","OC-48/STM-16","OC-192/STM-64","OC-768/STM-256","OC-1920/STM-640","OC-3840/STM-1234","SFP (1GFC)","SFP (2GFC)","SFP (4GFC)","SFP+ (8GFC)","SFP+ (16GFC)","SFP28 (32GFC)","QSFP+ (64GFC)","QSFP28 (128GFC)","SDR (2 Gbps)","DDR (4 Gbps)","QDR (8 Gbps)","FDR10 (10 Gbps)","FDR (13.5 Gbps)","EDR (25 Gbps)","HDR (50 Gbps)","NDR (100 Gbps)","XDR (250 Gbps)","T1 (1.544 Mbps)","E1 (2.048 Mbps)","T3 (45 Mbps)","E3 (34 Mbps)","Cisco StackWise","Cisco StackWise Plus","Cisco FlexStack","Cisco FlexStack Plus","Juniper VCP","Extreme SummitStack","Extreme SummitStack-128","Extreme SummitStack-256","Extreme SummitStack-512","Other"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		interfaceTypeTypeLabelPropEnum = append(interfaceTypeTypeLabelPropEnum, v)
	}
}

const (

	// InterfaceTypeLabelVirtual captures enum value "Virtual"
	InterfaceTypeLabelVirtual string = "Virtual"

	// InterfaceTypeLabelLinkAggregationGroupLAG captures enum value "Link Aggregation Group (LAG)"
	InterfaceTypeLabelLinkAggregationGroupLAG string = "Link Aggregation Group (LAG)"

	// InterfaceTypeLabelNr100BASEDashTX10100ME captures enum value "100BASE-TX (10/100ME)"
	InterfaceTypeLabelNr100BASEDashTX10100ME string = "100BASE-TX (10/100ME)"

	// InterfaceTypeLabelNr1000BASEDashT1GE captures enum value "1000BASE-T (1GE)"
	InterfaceTypeLabelNr1000BASEDashT1GE string = "1000BASE-T (1GE)"

	// InterfaceTypeLabelNr2Dot5GBASEDashT2Dot5GE captures enum value "2.5GBASE-T (2.5GE)"
	InterfaceTypeLabelNr2Dot5GBASEDashT2Dot5GE string = "2.5GBASE-T (2.5GE)"

	// InterfaceTypeLabelNr5GBASEDashT5GE captures enum value "5GBASE-T (5GE)"
	InterfaceTypeLabelNr5GBASEDashT5GE string = "5GBASE-T (5GE)"

	// InterfaceTypeLabelNr10GBASEDashT10GE captures enum value "10GBASE-T (10GE)"
	InterfaceTypeLabelNr10GBASEDashT10GE string = "10GBASE-T (10GE)"

	// InterfaceTypeLabelNr10GBASEDashCX410GE captures enum value "10GBASE-CX4 (10GE)"
	InterfaceTypeLabelNr10GBASEDashCX410GE string = "10GBASE-CX4 (10GE)"

	// InterfaceTypeLabelGBIC1GE captures enum value "GBIC (1GE)"
	InterfaceTypeLabelGBIC1GE string = "GBIC (1GE)"

	// InterfaceTypeLabelSFP1GE captures enum value "SFP (1GE)"
	InterfaceTypeLabelSFP1GE string = "SFP (1GE)"

	// InterfaceTypeLabelSFPPlus10GE captures enum value "SFP+ (10GE)"
	InterfaceTypeLabelSFPPlus10GE string = "SFP+ (10GE)"

	// InterfaceTypeLabelXFP10GE captures enum value "XFP (10GE)"
	InterfaceTypeLabelXFP10GE string = "XFP (10GE)"

	// InterfaceTypeLabelXENPAK10GE captures enum value "XENPAK (10GE)"
	InterfaceTypeLabelXENPAK10GE string = "XENPAK (10GE)"

	// InterfaceTypeLabelX210GE captures enum value "X2 (10GE)"
	InterfaceTypeLabelX210GE string = "X2 (10GE)"

	// InterfaceTypeLabelSFP2825GE captures enum value "SFP28 (25GE)"
	InterfaceTypeLabelSFP2825GE string = "SFP28 (25GE)"

	// InterfaceTypeLabelQSFPPlus40GE captures enum value "QSFP+ (40GE)"
	InterfaceTypeLabelQSFPPlus40GE string = "QSFP+ (40GE)"

	// InterfaceTypeLabelQSFP2850GE captures enum value "QSFP28 (50GE)"
	InterfaceTypeLabelQSFP2850GE string = "QSFP28 (50GE)"

	// InterfaceTypeLabelCFP100GE captures enum value "CFP (100GE)"
	InterfaceTypeLabelCFP100GE string = "CFP (100GE)"

	// InterfaceTypeLabelCFP2100GE captures enum value "CFP2 (100GE)"
	InterfaceTypeLabelCFP2100GE string = "CFP2 (100GE)"

	// InterfaceTypeLabelCFP2200GE captures enum value "CFP2 (200GE)"
	InterfaceTypeLabelCFP2200GE string = "CFP2 (200GE)"

	// InterfaceTypeLabelCFP4100GE captures enum value "CFP4 (100GE)"
	InterfaceTypeLabelCFP4100GE string = "CFP4 (100GE)"

	// InterfaceTypeLabelCiscoCPAK100GE captures enum value "Cisco CPAK (100GE)"
	InterfaceTypeLabelCiscoCPAK100GE string = "Cisco CPAK (100GE)"

	// InterfaceTypeLabelQSFP28100GE captures enum value "QSFP28 (100GE)"
	InterfaceTypeLabelQSFP28100GE string = "QSFP28 (100GE)"

	// InterfaceTypeLabelQSFP56200GE captures enum value "QSFP56 (200GE)"
	InterfaceTypeLabelQSFP56200GE string = "QSFP56 (200GE)"

	// InterfaceTypeLabelQSFPDashDD400GE captures enum value "QSFP-DD (400GE)"
	InterfaceTypeLabelQSFPDashDD400GE string = "QSFP-DD (400GE)"

	// InterfaceTypeLabelOSFP400GE captures enum value "OSFP (400GE)"
	InterfaceTypeLabelOSFP400GE string = "OSFP (400GE)"

	// InterfaceTypeLabelIEEE802Dot11a captures enum value "IEEE 802.11a"
	InterfaceTypeLabelIEEE802Dot11a string = "IEEE 802.11a"

	// InterfaceTypeLabelIEEE802Dot11bg captures enum value "IEEE 802.11b/g"
	InterfaceTypeLabelIEEE802Dot11bg string = "IEEE 802.11b/g"

	// InterfaceTypeLabelIEEE802Dot11n captures enum value "IEEE 802.11n"
	InterfaceTypeLabelIEEE802Dot11n string = "IEEE 802.11n"

	// InterfaceTypeLabelIEEE802Dot11ac captures enum value "IEEE 802.11ac"
	InterfaceTypeLabelIEEE802Dot11ac string = "IEEE 802.11ac"

	// InterfaceTypeLabelIEEE802Dot11ad captures enum value "IEEE 802.11ad"
	InterfaceTypeLabelIEEE802Dot11ad string = "IEEE 802.11ad"

	// InterfaceTypeLabelIEEE802Dot11ax captures enum value "IEEE 802.11ax"
	InterfaceTypeLabelIEEE802Dot11ax string = "IEEE 802.11ax"

	// InterfaceTypeLabelGSM captures enum value "GSM"
	InterfaceTypeLabelGSM string = "GSM"

	// InterfaceTypeLabelCDMA captures enum value "CDMA"
	InterfaceTypeLabelCDMA string = "CDMA"

	// InterfaceTypeLabelLTE captures enum value "LTE"
	InterfaceTypeLabelLTE string = "LTE"

	// InterfaceTypeLabelOCDash3STMDash1 captures enum value "OC-3/STM-1"
	InterfaceTypeLabelOCDash3STMDash1 string = "OC-3/STM-1"

	// InterfaceTypeLabelOCDash12STMDash4 captures enum value "OC-12/STM-4"
	InterfaceTypeLabelOCDash12STMDash4 string = "OC-12/STM-4"

	// InterfaceTypeLabelOCDash48STMDash16 captures enum value "OC-48/STM-16"
	InterfaceTypeLabelOCDash48STMDash16 string = "OC-48/STM-16"

	// InterfaceTypeLabelOCDash192STMDash64 captures enum value "OC-192/STM-64"
	InterfaceTypeLabelOCDash192STMDash64 string = "OC-192/STM-64"

	// InterfaceTypeLabelOCDash768STMDash256 captures enum value "OC-768/STM-256"
	InterfaceTypeLabelOCDash768STMDash256 string = "OC-768/STM-256"

	// InterfaceTypeLabelOCDash1920STMDash640 captures enum value "OC-1920/STM-640"
	InterfaceTypeLabelOCDash1920STMDash640 string = "OC-1920/STM-640"

	// InterfaceTypeLabelOCDash3840STMDash1234 captures enum value "OC-3840/STM-1234"
	InterfaceTypeLabelOCDash3840STMDash1234 string = "OC-3840/STM-1234"

	// InterfaceTypeLabelSFP1GFC captures enum value "SFP (1GFC)"
	InterfaceTypeLabelSFP1GFC string = "SFP (1GFC)"

	// InterfaceTypeLabelSFP2GFC captures enum value "SFP (2GFC)"
	InterfaceTypeLabelSFP2GFC string = "SFP (2GFC)"

	// InterfaceTypeLabelSFP4GFC captures enum value "SFP (4GFC)"
	InterfaceTypeLabelSFP4GFC string = "SFP (4GFC)"

	// InterfaceTypeLabelSFPPlus8GFC captures enum value "SFP+ (8GFC)"
	InterfaceTypeLabelSFPPlus8GFC string = "SFP+ (8GFC)"

	// InterfaceTypeLabelSFPPlus16GFC captures enum value "SFP+ (16GFC)"
	InterfaceTypeLabelSFPPlus16GFC string = "SFP+ (16GFC)"

	// InterfaceTypeLabelSFP2832GFC captures enum value "SFP28 (32GFC)"
	InterfaceTypeLabelSFP2832GFC string = "SFP28 (32GFC)"

	// InterfaceTypeLabelQSFPPlus64GFC captures enum value "QSFP+ (64GFC)"
	InterfaceTypeLabelQSFPPlus64GFC string = "QSFP+ (64GFC)"

	// InterfaceTypeLabelQSFP28128GFC captures enum value "QSFP28 (128GFC)"
	InterfaceTypeLabelQSFP28128GFC string = "QSFP28 (128GFC)"

	// InterfaceTypeLabelSDR2Gbps captures enum value "SDR (2 Gbps)"
	InterfaceTypeLabelSDR2Gbps string = "SDR (2 Gbps)"

	// InterfaceTypeLabelDDR4Gbps captures enum value "DDR (4 Gbps)"
	InterfaceTypeLabelDDR4Gbps string = "DDR (4 Gbps)"

	// InterfaceTypeLabelQDR8Gbps captures enum value "QDR (8 Gbps)"
	InterfaceTypeLabelQDR8Gbps string = "QDR (8 Gbps)"

	// InterfaceTypeLabelFDR1010Gbps captures enum value "FDR10 (10 Gbps)"
	InterfaceTypeLabelFDR1010Gbps string = "FDR10 (10 Gbps)"

	// InterfaceTypeLabelFDR13Dot5Gbps captures enum value "FDR (13.5 Gbps)"
	InterfaceTypeLabelFDR13Dot5Gbps string = "FDR (13.5 Gbps)"

	// InterfaceTypeLabelEDR25Gbps captures enum value "EDR (25 Gbps)"
	InterfaceTypeLabelEDR25Gbps string = "EDR (25 Gbps)"

	// InterfaceTypeLabelHDR50Gbps captures enum value "HDR (50 Gbps)"
	InterfaceTypeLabelHDR50Gbps string = "HDR (50 Gbps)"

	// InterfaceTypeLabelNDR100Gbps captures enum value "NDR (100 Gbps)"
	InterfaceTypeLabelNDR100Gbps string = "NDR (100 Gbps)"

	// InterfaceTypeLabelXDR250Gbps captures enum value "XDR (250 Gbps)"
	InterfaceTypeLabelXDR250Gbps string = "XDR (250 Gbps)"

	// InterfaceTypeLabelT11Dot544Mbps captures enum value "T1 (1.544 Mbps)"
	InterfaceTypeLabelT11Dot544Mbps string = "T1 (1.544 Mbps)"

	// InterfaceTypeLabelE12Dot048Mbps captures enum value "E1 (2.048 Mbps)"
	InterfaceTypeLabelE12Dot048Mbps string = "E1 (2.048 Mbps)"

	// InterfaceTypeLabelT345Mbps captures enum value "T3 (45 Mbps)"
	InterfaceTypeLabelT345Mbps string = "T3 (45 Mbps)"

	// InterfaceTypeLabelE334Mbps captures enum value "E3 (34 Mbps)"
	InterfaceTypeLabelE334Mbps string = "E3 (34 Mbps)"

	// InterfaceTypeLabelCiscoStackWise captures enum value "Cisco StackWise"
	InterfaceTypeLabelCiscoStackWise string = "Cisco StackWise"

	// InterfaceTypeLabelCiscoStackWisePlus captures enum value "Cisco StackWise Plus"
	InterfaceTypeLabelCiscoStackWisePlus string = "Cisco StackWise Plus"

	// InterfaceTypeLabelCiscoFlexStack captures enum value "Cisco FlexStack"
	InterfaceTypeLabelCiscoFlexStack string = "Cisco FlexStack"

	// InterfaceTypeLabelCiscoFlexStackPlus captures enum value "Cisco FlexStack Plus"
	InterfaceTypeLabelCiscoFlexStackPlus string = "Cisco FlexStack Plus"

	// InterfaceTypeLabelJuniperVCP captures enum value "Juniper VCP"
	InterfaceTypeLabelJuniperVCP string = "Juniper VCP"

	// InterfaceTypeLabelExtremeSummitStack captures enum value "Extreme SummitStack"
	InterfaceTypeLabelExtremeSummitStack string = "Extreme SummitStack"

	// InterfaceTypeLabelExtremeSummitStackDash128 captures enum value "Extreme SummitStack-128"
	InterfaceTypeLabelExtremeSummitStackDash128 string = "Extreme SummitStack-128"

	// InterfaceTypeLabelExtremeSummitStackDash256 captures enum value "Extreme SummitStack-256"
	InterfaceTypeLabelExtremeSummitStackDash256 string = "Extreme SummitStack-256"

	// InterfaceTypeLabelExtremeSummitStackDash512 captures enum value "Extreme SummitStack-512"
	InterfaceTypeLabelExtremeSummitStackDash512 string = "Extreme SummitStack-512"

	// InterfaceTypeLabelOther captures enum value "Other"
	InterfaceTypeLabelOther string = "Other"
)

// prop value enum
func (m *InterfaceType) validateLabelEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, interfaceTypeTypeLabelPropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *InterfaceType) validateLabel(formats strfmt.Registry) error {

	if err := validate.Required("type"+"."+"label", "body", m.Label); err != nil {
		return err
	}

	// value enum
	if err := m.validateLabelEnum("type"+"."+"label", "body", *m.Label); err != nil {
		return err
	}

	return nil
}

var interfaceTypeTypeValuePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["virtual","lag","100base-tx","1000base-t","2.5gbase-t","5gbase-t","10gbase-t","10gbase-cx4","1000base-x-gbic","1000base-x-sfp","10gbase-x-sfpp","10gbase-x-xfp","10gbase-x-xenpak","10gbase-x-x2","25gbase-x-sfp28","40gbase-x-qsfpp","50gbase-x-sfp28","100gbase-x-cfp","100gbase-x-cfp2","200gbase-x-cfp2","100gbase-x-cfp4","100gbase-x-cpak","100gbase-x-qsfp28","200gbase-x-qsfp56","400gbase-x-qsfpdd","400gbase-x-osfp","ieee802.11a","ieee802.11g","ieee802.11n","ieee802.11ac","ieee802.11ad","ieee802.11ax","gsm","cdma","lte","sonet-oc3","sonet-oc12","sonet-oc48","sonet-oc192","sonet-oc768","sonet-oc1920","sonet-oc3840","1gfc-sfp","2gfc-sfp","4gfc-sfp","8gfc-sfpp","16gfc-sfpp","32gfc-sfp28","64gfc-qsfpp","128gfc-sfp28","infiniband-sdr","infiniband-ddr","infiniband-qdr","infiniband-fdr10","infiniband-fdr","infiniband-edr","infiniband-hdr","infiniband-ndr","infiniband-xdr","t1","e1","t3","e3","cisco-stackwise","cisco-stackwise-plus","cisco-flexstack","cisco-flexstack-plus","juniper-vcp","extreme-summitstack","extreme-summitstack-128","extreme-summitstack-256","extreme-summitstack-512","other"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		interfaceTypeTypeValuePropEnum = append(interfaceTypeTypeValuePropEnum, v)
	}
}

const (

	// InterfaceTypeValueVirtual captures enum value "virtual"
	InterfaceTypeValueVirtual string = "virtual"

	// InterfaceTypeValueLag captures enum value "lag"
	InterfaceTypeValueLag string = "lag"

	// InterfaceTypeValueNr100baseDashTx captures enum value "100base-tx"
	InterfaceTypeValueNr100baseDashTx string = "100base-tx"

	// InterfaceTypeValueNr1000baseDasht captures enum value "1000base-t"
	InterfaceTypeValueNr1000baseDasht string = "1000base-t"

	// InterfaceTypeValueNr2Dot5gbaseDasht captures enum value "2.5gbase-t"
	InterfaceTypeValueNr2Dot5gbaseDasht string = "2.5gbase-t"

	// InterfaceTypeValueNr5gbaseDasht captures enum value "5gbase-t"
	InterfaceTypeValueNr5gbaseDasht string = "5gbase-t"

	// InterfaceTypeValueNr10gbaseDasht captures enum value "10gbase-t"
	InterfaceTypeValueNr10gbaseDasht string = "10gbase-t"

	// InterfaceTypeValueNr10gbaseDashCx4 captures enum value "10gbase-cx4"
	InterfaceTypeValueNr10gbaseDashCx4 string = "10gbase-cx4"

	// InterfaceTypeValueNr1000baseDashxDashGbic captures enum value "1000base-x-gbic"
	InterfaceTypeValueNr1000baseDashxDashGbic string = "1000base-x-gbic"

	// InterfaceTypeValueNr1000baseDashxDashSfp captures enum value "1000base-x-sfp"
	InterfaceTypeValueNr1000baseDashxDashSfp string = "1000base-x-sfp"

	// InterfaceTypeValueNr10gbaseDashxDashSfpp captures enum value "10gbase-x-sfpp"
	InterfaceTypeValueNr10gbaseDashxDashSfpp string = "10gbase-x-sfpp"

	// InterfaceTypeValueNr10gbaseDashxDashXfp captures enum value "10gbase-x-xfp"
	InterfaceTypeValueNr10gbaseDashxDashXfp string = "10gbase-x-xfp"

	// InterfaceTypeValueNr10gbaseDashxDashXenpak captures enum value "10gbase-x-xenpak"
	InterfaceTypeValueNr10gbaseDashxDashXenpak string = "10gbase-x-xenpak"

	// InterfaceTypeValueNr10gbaseDashxDashX2 captures enum value "10gbase-x-x2"
	InterfaceTypeValueNr10gbaseDashxDashX2 string = "10gbase-x-x2"

	// InterfaceTypeValueNr25gbaseDashxDashSfp28 captures enum value "25gbase-x-sfp28"
	InterfaceTypeValueNr25gbaseDashxDashSfp28 string = "25gbase-x-sfp28"

	// InterfaceTypeValueNr40gbaseDashxDashQsfpp captures enum value "40gbase-x-qsfpp"
	InterfaceTypeValueNr40gbaseDashxDashQsfpp string = "40gbase-x-qsfpp"

	// InterfaceTypeValueNr50gbaseDashxDashSfp28 captures enum value "50gbase-x-sfp28"
	InterfaceTypeValueNr50gbaseDashxDashSfp28 string = "50gbase-x-sfp28"

	// InterfaceTypeValueNr100gbaseDashxDashCfp captures enum value "100gbase-x-cfp"
	InterfaceTypeValueNr100gbaseDashxDashCfp string = "100gbase-x-cfp"

	// InterfaceTypeValueNr100gbaseDashxDashCfp2 captures enum value "100gbase-x-cfp2"
	InterfaceTypeValueNr100gbaseDashxDashCfp2 string = "100gbase-x-cfp2"

	// InterfaceTypeValueNr200gbaseDashxDashCfp2 captures enum value "200gbase-x-cfp2"
	InterfaceTypeValueNr200gbaseDashxDashCfp2 string = "200gbase-x-cfp2"

	// InterfaceTypeValueNr100gbaseDashxDashCfp4 captures enum value "100gbase-x-cfp4"
	InterfaceTypeValueNr100gbaseDashxDashCfp4 string = "100gbase-x-cfp4"

	// InterfaceTypeValueNr100gbaseDashxDashCpak captures enum value "100gbase-x-cpak"
	InterfaceTypeValueNr100gbaseDashxDashCpak string = "100gbase-x-cpak"

	// InterfaceTypeValueNr100gbaseDashxDashQsfp28 captures enum value "100gbase-x-qsfp28"
	InterfaceTypeValueNr100gbaseDashxDashQsfp28 string = "100gbase-x-qsfp28"

	// InterfaceTypeValueNr200gbaseDashxDashQsfp56 captures enum value "200gbase-x-qsfp56"
	InterfaceTypeValueNr200gbaseDashxDashQsfp56 string = "200gbase-x-qsfp56"

	// InterfaceTypeValueNr400gbaseDashxDashQsfpdd captures enum value "400gbase-x-qsfpdd"
	InterfaceTypeValueNr400gbaseDashxDashQsfpdd string = "400gbase-x-qsfpdd"

	// InterfaceTypeValueNr400gbaseDashxDashOsfp captures enum value "400gbase-x-osfp"
	InterfaceTypeValueNr400gbaseDashxDashOsfp string = "400gbase-x-osfp"

	// InterfaceTypeValueIeee802Dot11a captures enum value "ieee802.11a"
	InterfaceTypeValueIeee802Dot11a string = "ieee802.11a"

	// InterfaceTypeValueIeee802Dot11g captures enum value "ieee802.11g"
	InterfaceTypeValueIeee802Dot11g string = "ieee802.11g"

	// InterfaceTypeValueIeee802Dot11n captures enum value "ieee802.11n"
	InterfaceTypeValueIeee802Dot11n string = "ieee802.11n"

	// InterfaceTypeValueIeee802Dot11ac captures enum value "ieee802.11ac"
	InterfaceTypeValueIeee802Dot11ac string = "ieee802.11ac"

	// InterfaceTypeValueIeee802Dot11ad captures enum value "ieee802.11ad"
	InterfaceTypeValueIeee802Dot11ad string = "ieee802.11ad"

	// InterfaceTypeValueIeee802Dot11ax captures enum value "ieee802.11ax"
	InterfaceTypeValueIeee802Dot11ax string = "ieee802.11ax"

	// InterfaceTypeValueGsm captures enum value "gsm"
	InterfaceTypeValueGsm string = "gsm"

	// InterfaceTypeValueCdma captures enum value "cdma"
	InterfaceTypeValueCdma string = "cdma"

	// InterfaceTypeValueLte captures enum value "lte"
	InterfaceTypeValueLte string = "lte"

	// InterfaceTypeValueSonetDashOc3 captures enum value "sonet-oc3"
	InterfaceTypeValueSonetDashOc3 string = "sonet-oc3"

	// InterfaceTypeValueSonetDashOc12 captures enum value "sonet-oc12"
	InterfaceTypeValueSonetDashOc12 string = "sonet-oc12"

	// InterfaceTypeValueSonetDashOc48 captures enum value "sonet-oc48"
	InterfaceTypeValueSonetDashOc48 string = "sonet-oc48"

	// InterfaceTypeValueSonetDashOc192 captures enum value "sonet-oc192"
	InterfaceTypeValueSonetDashOc192 string = "sonet-oc192"

	// InterfaceTypeValueSonetDashOc768 captures enum value "sonet-oc768"
	InterfaceTypeValueSonetDashOc768 string = "sonet-oc768"

	// InterfaceTypeValueSonetDashOc1920 captures enum value "sonet-oc1920"
	InterfaceTypeValueSonetDashOc1920 string = "sonet-oc1920"

	// InterfaceTypeValueSonetDashOc3840 captures enum value "sonet-oc3840"
	InterfaceTypeValueSonetDashOc3840 string = "sonet-oc3840"

	// InterfaceTypeValueNr1gfcDashSfp captures enum value "1gfc-sfp"
	InterfaceTypeValueNr1gfcDashSfp string = "1gfc-sfp"

	// InterfaceTypeValueNr2gfcDashSfp captures enum value "2gfc-sfp"
	InterfaceTypeValueNr2gfcDashSfp string = "2gfc-sfp"

	// InterfaceTypeValueNr4gfcDashSfp captures enum value "4gfc-sfp"
	InterfaceTypeValueNr4gfcDashSfp string = "4gfc-sfp"

	// InterfaceTypeValueNr8gfcDashSfpp captures enum value "8gfc-sfpp"
	InterfaceTypeValueNr8gfcDashSfpp string = "8gfc-sfpp"

	// InterfaceTypeValueNr16gfcDashSfpp captures enum value "16gfc-sfpp"
	InterfaceTypeValueNr16gfcDashSfpp string = "16gfc-sfpp"

	// InterfaceTypeValueNr32gfcDashSfp28 captures enum value "32gfc-sfp28"
	InterfaceTypeValueNr32gfcDashSfp28 string = "32gfc-sfp28"

	// InterfaceTypeValueNr64gfcDashQsfpp captures enum value "64gfc-qsfpp"
	InterfaceTypeValueNr64gfcDashQsfpp string = "64gfc-qsfpp"

	// InterfaceTypeValueNr128gfcDashSfp28 captures enum value "128gfc-sfp28"
	InterfaceTypeValueNr128gfcDashSfp28 string = "128gfc-sfp28"

	// InterfaceTypeValueInfinibandDashSdr captures enum value "infiniband-sdr"
	InterfaceTypeValueInfinibandDashSdr string = "infiniband-sdr"

	// InterfaceTypeValueInfinibandDashDdr captures enum value "infiniband-ddr"
	InterfaceTypeValueInfinibandDashDdr string = "infiniband-ddr"

	// InterfaceTypeValueInfinibandDashQdr captures enum value "infiniband-qdr"
	InterfaceTypeValueInfinibandDashQdr string = "infiniband-qdr"

	// InterfaceTypeValueInfinibandDashFdr10 captures enum value "infiniband-fdr10"
	InterfaceTypeValueInfinibandDashFdr10 string = "infiniband-fdr10"

	// InterfaceTypeValueInfinibandDashFdr captures enum value "infiniband-fdr"
	InterfaceTypeValueInfinibandDashFdr string = "infiniband-fdr"

	// InterfaceTypeValueInfinibandDashEdr captures enum value "infiniband-edr"
	InterfaceTypeValueInfinibandDashEdr string = "infiniband-edr"

	// InterfaceTypeValueInfinibandDashHdr captures enum value "infiniband-hdr"
	InterfaceTypeValueInfinibandDashHdr string = "infiniband-hdr"

	// InterfaceTypeValueInfinibandDashNdr captures enum value "infiniband-ndr"
	InterfaceTypeValueInfinibandDashNdr string = "infiniband-ndr"

	// InterfaceTypeValueInfinibandDashXdr captures enum value "infiniband-xdr"
	InterfaceTypeValueInfinibandDashXdr string = "infiniband-xdr"

	// InterfaceTypeValueT1 captures enum value "t1"
	InterfaceTypeValueT1 string = "t1"

	// InterfaceTypeValueE1 captures enum value "e1"
	InterfaceTypeValueE1 string = "e1"

	// InterfaceTypeValueT3 captures enum value "t3"
	InterfaceTypeValueT3 string = "t3"

	// InterfaceTypeValueE3 captures enum value "e3"
	InterfaceTypeValueE3 string = "e3"

	// InterfaceTypeValueCiscoDashStackwise captures enum value "cisco-stackwise"
	InterfaceTypeValueCiscoDashStackwise string = "cisco-stackwise"

	// InterfaceTypeValueCiscoDashStackwiseDashPlus captures enum value "cisco-stackwise-plus"
	InterfaceTypeValueCiscoDashStackwiseDashPlus string = "cisco-stackwise-plus"

	// InterfaceTypeValueCiscoDashFlexstack captures enum value "cisco-flexstack"
	InterfaceTypeValueCiscoDashFlexstack string = "cisco-flexstack"

	// InterfaceTypeValueCiscoDashFlexstackDashPlus captures enum value "cisco-flexstack-plus"
	InterfaceTypeValueCiscoDashFlexstackDashPlus string = "cisco-flexstack-plus"

	// InterfaceTypeValueJuniperDashVcp captures enum value "juniper-vcp"
	InterfaceTypeValueJuniperDashVcp string = "juniper-vcp"

	// InterfaceTypeValueExtremeDashSummitstack captures enum value "extreme-summitstack"
	InterfaceTypeValueExtremeDashSummitstack string = "extreme-summitstack"

	// InterfaceTypeValueExtremeDashSummitstackDash128 captures enum value "extreme-summitstack-128"
	InterfaceTypeValueExtremeDashSummitstackDash128 string = "extreme-summitstack-128"

	// InterfaceTypeValueExtremeDashSummitstackDash256 captures enum value "extreme-summitstack-256"
	InterfaceTypeValueExtremeDashSummitstackDash256 string = "extreme-summitstack-256"

	// InterfaceTypeValueExtremeDashSummitstackDash512 captures enum value "extreme-summitstack-512"
	InterfaceTypeValueExtremeDashSummitstackDash512 string = "extreme-summitstack-512"

	// InterfaceTypeValueOther captures enum value "other"
	InterfaceTypeValueOther string = "other"
)

// prop value enum
func (m *InterfaceType) validateValueEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, interfaceTypeTypeValuePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *InterfaceType) validateValue(formats strfmt.Registry) error {

	if err := validate.Required("type"+"."+"value", "body", m.Value); err != nil {
		return err
	}

	// value enum
	if err := m.validateValueEnum("type"+"."+"value", "body", *m.Value); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this interface type based on context it is used
func (m *InterfaceType) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *InterfaceType) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *InterfaceType) UnmarshalBinary(b []byte) error {
	var res InterfaceType
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
