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

// WritableInterface writable interface
//
// swagger:model WritableInterface
type WritableInterface struct {

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

	// Count ipaddresses
	// Read Only: true
	CountIpaddresses int64 `json:"count_ipaddresses,omitempty"`

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

	// Parent LAG
	// Format: uuid
	Lag *strfmt.UUID `json:"lag,omitempty"`

	// MAC Address
	// Max Length: 18
	MacAddress *string `json:"mac_address,omitempty"`

	// Management only
	//
	// This interface is used only for out-of-band management
	MgmtOnly bool `json:"mgmt_only,omitempty"`

	// Mode
	// Enum: [access tagged tagged-all]
	Mode string `json:"mode,omitempty"`

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

	// Type
	// Required: true
	// Enum: [virtual lag 100base-tx 1000base-t 2.5gbase-t 5gbase-t 10gbase-t 10gbase-cx4 1000base-x-gbic 1000base-x-sfp 10gbase-x-sfpp 10gbase-x-xfp 10gbase-x-xenpak 10gbase-x-x2 25gbase-x-sfp28 40gbase-x-qsfpp 50gbase-x-sfp28 100gbase-x-cfp 100gbase-x-cfp2 200gbase-x-cfp2 100gbase-x-cfp4 100gbase-x-cpak 100gbase-x-qsfp28 200gbase-x-qsfp56 400gbase-x-qsfpdd 400gbase-x-osfp ieee802.11a ieee802.11g ieee802.11n ieee802.11ac ieee802.11ad ieee802.11ax gsm cdma lte sonet-oc3 sonet-oc12 sonet-oc48 sonet-oc192 sonet-oc768 sonet-oc1920 sonet-oc3840 1gfc-sfp 2gfc-sfp 4gfc-sfp 8gfc-sfpp 16gfc-sfpp 32gfc-sfp28 64gfc-qsfpp 128gfc-sfp28 infiniband-sdr infiniband-ddr infiniband-qdr infiniband-fdr10 infiniband-fdr infiniband-edr infiniband-hdr infiniband-ndr infiniband-xdr t1 e1 t3 e3 cisco-stackwise cisco-stackwise-plus cisco-flexstack cisco-flexstack-plus juniper-vcp extreme-summitstack extreme-summitstack-128 extreme-summitstack-256 extreme-summitstack-512 other]
	Type *string `json:"type"`

	// Untagged VLAN
	// Format: uuid
	UntaggedVlan *strfmt.UUID `json:"untagged_vlan,omitempty"`

	// Url
	// Read Only: true
	// Format: uri
	URL strfmt.URI `json:"url,omitempty"`
}

// Validate validates this writable interface
func (m *WritableInterface) Validate(formats strfmt.Registry) error {
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

func (m *WritableInterface) validateCable(formats strfmt.Registry) error {
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

func (m *WritableInterface) validateDescription(formats strfmt.Registry) error {
	if swag.IsZero(m.Description) { // not required
		return nil
	}

	if err := validate.MaxLength("description", "body", m.Description, 200); err != nil {
		return err
	}

	return nil
}

func (m *WritableInterface) validateDevice(formats strfmt.Registry) error {

	if err := validate.Required("device", "body", m.Device); err != nil {
		return err
	}

	if err := validate.FormatOf("device", "body", "uuid", m.Device.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *WritableInterface) validateDisplay(formats strfmt.Registry) error {
	if swag.IsZero(m.Display) { // not required
		return nil
	}

	if err := validate.MinLength("display", "body", m.Display, 1); err != nil {
		return err
	}

	return nil
}

func (m *WritableInterface) validateID(formats strfmt.Registry) error {
	if swag.IsZero(m.ID) { // not required
		return nil
	}

	if err := validate.FormatOf("id", "body", "uuid", m.ID.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *WritableInterface) validateLabel(formats strfmt.Registry) error {
	if swag.IsZero(m.Label) { // not required
		return nil
	}

	if err := validate.MaxLength("label", "body", m.Label, 64); err != nil {
		return err
	}

	return nil
}

func (m *WritableInterface) validateLag(formats strfmt.Registry) error {
	if swag.IsZero(m.Lag) { // not required
		return nil
	}

	if err := validate.FormatOf("lag", "body", "uuid", m.Lag.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *WritableInterface) validateMacAddress(formats strfmt.Registry) error {
	if swag.IsZero(m.MacAddress) { // not required
		return nil
	}

	if err := validate.MaxLength("mac_address", "body", *m.MacAddress, 18); err != nil {
		return err
	}

	return nil
}

var writableInterfaceTypeModePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["access","tagged","tagged-all"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		writableInterfaceTypeModePropEnum = append(writableInterfaceTypeModePropEnum, v)
	}
}

const (

	// WritableInterfaceModeAccess captures enum value "access"
	WritableInterfaceModeAccess string = "access"

	// WritableInterfaceModeTagged captures enum value "tagged"
	WritableInterfaceModeTagged string = "tagged"

	// WritableInterfaceModeTaggedDashAll captures enum value "tagged-all"
	WritableInterfaceModeTaggedDashAll string = "tagged-all"
)

// prop value enum
func (m *WritableInterface) validateModeEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, writableInterfaceTypeModePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *WritableInterface) validateMode(formats strfmt.Registry) error {
	if swag.IsZero(m.Mode) { // not required
		return nil
	}

	// value enum
	if err := m.validateModeEnum("mode", "body", m.Mode); err != nil {
		return err
	}

	return nil
}

func (m *WritableInterface) validateMtu(formats strfmt.Registry) error {
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

func (m *WritableInterface) validateName(formats strfmt.Registry) error {

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

func (m *WritableInterface) validateTaggedVlans(formats strfmt.Registry) error {
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

func (m *WritableInterface) validateTags(formats strfmt.Registry) error {
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

var writableInterfaceTypeTypePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["virtual","lag","100base-tx","1000base-t","2.5gbase-t","5gbase-t","10gbase-t","10gbase-cx4","1000base-x-gbic","1000base-x-sfp","10gbase-x-sfpp","10gbase-x-xfp","10gbase-x-xenpak","10gbase-x-x2","25gbase-x-sfp28","40gbase-x-qsfpp","50gbase-x-sfp28","100gbase-x-cfp","100gbase-x-cfp2","200gbase-x-cfp2","100gbase-x-cfp4","100gbase-x-cpak","100gbase-x-qsfp28","200gbase-x-qsfp56","400gbase-x-qsfpdd","400gbase-x-osfp","ieee802.11a","ieee802.11g","ieee802.11n","ieee802.11ac","ieee802.11ad","ieee802.11ax","gsm","cdma","lte","sonet-oc3","sonet-oc12","sonet-oc48","sonet-oc192","sonet-oc768","sonet-oc1920","sonet-oc3840","1gfc-sfp","2gfc-sfp","4gfc-sfp","8gfc-sfpp","16gfc-sfpp","32gfc-sfp28","64gfc-qsfpp","128gfc-sfp28","infiniband-sdr","infiniband-ddr","infiniband-qdr","infiniband-fdr10","infiniband-fdr","infiniband-edr","infiniband-hdr","infiniband-ndr","infiniband-xdr","t1","e1","t3","e3","cisco-stackwise","cisco-stackwise-plus","cisco-flexstack","cisco-flexstack-plus","juniper-vcp","extreme-summitstack","extreme-summitstack-128","extreme-summitstack-256","extreme-summitstack-512","other"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		writableInterfaceTypeTypePropEnum = append(writableInterfaceTypeTypePropEnum, v)
	}
}

const (

	// WritableInterfaceTypeVirtual captures enum value "virtual"
	WritableInterfaceTypeVirtual string = "virtual"

	// WritableInterfaceTypeLag captures enum value "lag"
	WritableInterfaceTypeLag string = "lag"

	// WritableInterfaceTypeNr100baseDashTx captures enum value "100base-tx"
	WritableInterfaceTypeNr100baseDashTx string = "100base-tx"

	// WritableInterfaceTypeNr1000baseDasht captures enum value "1000base-t"
	WritableInterfaceTypeNr1000baseDasht string = "1000base-t"

	// WritableInterfaceTypeNr2Dot5gbaseDasht captures enum value "2.5gbase-t"
	WritableInterfaceTypeNr2Dot5gbaseDasht string = "2.5gbase-t"

	// WritableInterfaceTypeNr5gbaseDasht captures enum value "5gbase-t"
	WritableInterfaceTypeNr5gbaseDasht string = "5gbase-t"

	// WritableInterfaceTypeNr10gbaseDasht captures enum value "10gbase-t"
	WritableInterfaceTypeNr10gbaseDasht string = "10gbase-t"

	// WritableInterfaceTypeNr10gbaseDashCx4 captures enum value "10gbase-cx4"
	WritableInterfaceTypeNr10gbaseDashCx4 string = "10gbase-cx4"

	// WritableInterfaceTypeNr1000baseDashxDashGbic captures enum value "1000base-x-gbic"
	WritableInterfaceTypeNr1000baseDashxDashGbic string = "1000base-x-gbic"

	// WritableInterfaceTypeNr1000baseDashxDashSfp captures enum value "1000base-x-sfp"
	WritableInterfaceTypeNr1000baseDashxDashSfp string = "1000base-x-sfp"

	// WritableInterfaceTypeNr10gbaseDashxDashSfpp captures enum value "10gbase-x-sfpp"
	WritableInterfaceTypeNr10gbaseDashxDashSfpp string = "10gbase-x-sfpp"

	// WritableInterfaceTypeNr10gbaseDashxDashXfp captures enum value "10gbase-x-xfp"
	WritableInterfaceTypeNr10gbaseDashxDashXfp string = "10gbase-x-xfp"

	// WritableInterfaceTypeNr10gbaseDashxDashXenpak captures enum value "10gbase-x-xenpak"
	WritableInterfaceTypeNr10gbaseDashxDashXenpak string = "10gbase-x-xenpak"

	// WritableInterfaceTypeNr10gbaseDashxDashX2 captures enum value "10gbase-x-x2"
	WritableInterfaceTypeNr10gbaseDashxDashX2 string = "10gbase-x-x2"

	// WritableInterfaceTypeNr25gbaseDashxDashSfp28 captures enum value "25gbase-x-sfp28"
	WritableInterfaceTypeNr25gbaseDashxDashSfp28 string = "25gbase-x-sfp28"

	// WritableInterfaceTypeNr40gbaseDashxDashQsfpp captures enum value "40gbase-x-qsfpp"
	WritableInterfaceTypeNr40gbaseDashxDashQsfpp string = "40gbase-x-qsfpp"

	// WritableInterfaceTypeNr50gbaseDashxDashSfp28 captures enum value "50gbase-x-sfp28"
	WritableInterfaceTypeNr50gbaseDashxDashSfp28 string = "50gbase-x-sfp28"

	// WritableInterfaceTypeNr100gbaseDashxDashCfp captures enum value "100gbase-x-cfp"
	WritableInterfaceTypeNr100gbaseDashxDashCfp string = "100gbase-x-cfp"

	// WritableInterfaceTypeNr100gbaseDashxDashCfp2 captures enum value "100gbase-x-cfp2"
	WritableInterfaceTypeNr100gbaseDashxDashCfp2 string = "100gbase-x-cfp2"

	// WritableInterfaceTypeNr200gbaseDashxDashCfp2 captures enum value "200gbase-x-cfp2"
	WritableInterfaceTypeNr200gbaseDashxDashCfp2 string = "200gbase-x-cfp2"

	// WritableInterfaceTypeNr100gbaseDashxDashCfp4 captures enum value "100gbase-x-cfp4"
	WritableInterfaceTypeNr100gbaseDashxDashCfp4 string = "100gbase-x-cfp4"

	// WritableInterfaceTypeNr100gbaseDashxDashCpak captures enum value "100gbase-x-cpak"
	WritableInterfaceTypeNr100gbaseDashxDashCpak string = "100gbase-x-cpak"

	// WritableInterfaceTypeNr100gbaseDashxDashQsfp28 captures enum value "100gbase-x-qsfp28"
	WritableInterfaceTypeNr100gbaseDashxDashQsfp28 string = "100gbase-x-qsfp28"

	// WritableInterfaceTypeNr200gbaseDashxDashQsfp56 captures enum value "200gbase-x-qsfp56"
	WritableInterfaceTypeNr200gbaseDashxDashQsfp56 string = "200gbase-x-qsfp56"

	// WritableInterfaceTypeNr400gbaseDashxDashQsfpdd captures enum value "400gbase-x-qsfpdd"
	WritableInterfaceTypeNr400gbaseDashxDashQsfpdd string = "400gbase-x-qsfpdd"

	// WritableInterfaceTypeNr400gbaseDashxDashOsfp captures enum value "400gbase-x-osfp"
	WritableInterfaceTypeNr400gbaseDashxDashOsfp string = "400gbase-x-osfp"

	// WritableInterfaceTypeIeee802Dot11a captures enum value "ieee802.11a"
	WritableInterfaceTypeIeee802Dot11a string = "ieee802.11a"

	// WritableInterfaceTypeIeee802Dot11g captures enum value "ieee802.11g"
	WritableInterfaceTypeIeee802Dot11g string = "ieee802.11g"

	// WritableInterfaceTypeIeee802Dot11n captures enum value "ieee802.11n"
	WritableInterfaceTypeIeee802Dot11n string = "ieee802.11n"

	// WritableInterfaceTypeIeee802Dot11ac captures enum value "ieee802.11ac"
	WritableInterfaceTypeIeee802Dot11ac string = "ieee802.11ac"

	// WritableInterfaceTypeIeee802Dot11ad captures enum value "ieee802.11ad"
	WritableInterfaceTypeIeee802Dot11ad string = "ieee802.11ad"

	// WritableInterfaceTypeIeee802Dot11ax captures enum value "ieee802.11ax"
	WritableInterfaceTypeIeee802Dot11ax string = "ieee802.11ax"

	// WritableInterfaceTypeGsm captures enum value "gsm"
	WritableInterfaceTypeGsm string = "gsm"

	// WritableInterfaceTypeCdma captures enum value "cdma"
	WritableInterfaceTypeCdma string = "cdma"

	// WritableInterfaceTypeLte captures enum value "lte"
	WritableInterfaceTypeLte string = "lte"

	// WritableInterfaceTypeSonetDashOc3 captures enum value "sonet-oc3"
	WritableInterfaceTypeSonetDashOc3 string = "sonet-oc3"

	// WritableInterfaceTypeSonetDashOc12 captures enum value "sonet-oc12"
	WritableInterfaceTypeSonetDashOc12 string = "sonet-oc12"

	// WritableInterfaceTypeSonetDashOc48 captures enum value "sonet-oc48"
	WritableInterfaceTypeSonetDashOc48 string = "sonet-oc48"

	// WritableInterfaceTypeSonetDashOc192 captures enum value "sonet-oc192"
	WritableInterfaceTypeSonetDashOc192 string = "sonet-oc192"

	// WritableInterfaceTypeSonetDashOc768 captures enum value "sonet-oc768"
	WritableInterfaceTypeSonetDashOc768 string = "sonet-oc768"

	// WritableInterfaceTypeSonetDashOc1920 captures enum value "sonet-oc1920"
	WritableInterfaceTypeSonetDashOc1920 string = "sonet-oc1920"

	// WritableInterfaceTypeSonetDashOc3840 captures enum value "sonet-oc3840"
	WritableInterfaceTypeSonetDashOc3840 string = "sonet-oc3840"

	// WritableInterfaceTypeNr1gfcDashSfp captures enum value "1gfc-sfp"
	WritableInterfaceTypeNr1gfcDashSfp string = "1gfc-sfp"

	// WritableInterfaceTypeNr2gfcDashSfp captures enum value "2gfc-sfp"
	WritableInterfaceTypeNr2gfcDashSfp string = "2gfc-sfp"

	// WritableInterfaceTypeNr4gfcDashSfp captures enum value "4gfc-sfp"
	WritableInterfaceTypeNr4gfcDashSfp string = "4gfc-sfp"

	// WritableInterfaceTypeNr8gfcDashSfpp captures enum value "8gfc-sfpp"
	WritableInterfaceTypeNr8gfcDashSfpp string = "8gfc-sfpp"

	// WritableInterfaceTypeNr16gfcDashSfpp captures enum value "16gfc-sfpp"
	WritableInterfaceTypeNr16gfcDashSfpp string = "16gfc-sfpp"

	// WritableInterfaceTypeNr32gfcDashSfp28 captures enum value "32gfc-sfp28"
	WritableInterfaceTypeNr32gfcDashSfp28 string = "32gfc-sfp28"

	// WritableInterfaceTypeNr64gfcDashQsfpp captures enum value "64gfc-qsfpp"
	WritableInterfaceTypeNr64gfcDashQsfpp string = "64gfc-qsfpp"

	// WritableInterfaceTypeNr128gfcDashSfp28 captures enum value "128gfc-sfp28"
	WritableInterfaceTypeNr128gfcDashSfp28 string = "128gfc-sfp28"

	// WritableInterfaceTypeInfinibandDashSdr captures enum value "infiniband-sdr"
	WritableInterfaceTypeInfinibandDashSdr string = "infiniband-sdr"

	// WritableInterfaceTypeInfinibandDashDdr captures enum value "infiniband-ddr"
	WritableInterfaceTypeInfinibandDashDdr string = "infiniband-ddr"

	// WritableInterfaceTypeInfinibandDashQdr captures enum value "infiniband-qdr"
	WritableInterfaceTypeInfinibandDashQdr string = "infiniband-qdr"

	// WritableInterfaceTypeInfinibandDashFdr10 captures enum value "infiniband-fdr10"
	WritableInterfaceTypeInfinibandDashFdr10 string = "infiniband-fdr10"

	// WritableInterfaceTypeInfinibandDashFdr captures enum value "infiniband-fdr"
	WritableInterfaceTypeInfinibandDashFdr string = "infiniband-fdr"

	// WritableInterfaceTypeInfinibandDashEdr captures enum value "infiniband-edr"
	WritableInterfaceTypeInfinibandDashEdr string = "infiniband-edr"

	// WritableInterfaceTypeInfinibandDashHdr captures enum value "infiniband-hdr"
	WritableInterfaceTypeInfinibandDashHdr string = "infiniband-hdr"

	// WritableInterfaceTypeInfinibandDashNdr captures enum value "infiniband-ndr"
	WritableInterfaceTypeInfinibandDashNdr string = "infiniband-ndr"

	// WritableInterfaceTypeInfinibandDashXdr captures enum value "infiniband-xdr"
	WritableInterfaceTypeInfinibandDashXdr string = "infiniband-xdr"

	// WritableInterfaceTypeT1 captures enum value "t1"
	WritableInterfaceTypeT1 string = "t1"

	// WritableInterfaceTypeE1 captures enum value "e1"
	WritableInterfaceTypeE1 string = "e1"

	// WritableInterfaceTypeT3 captures enum value "t3"
	WritableInterfaceTypeT3 string = "t3"

	// WritableInterfaceTypeE3 captures enum value "e3"
	WritableInterfaceTypeE3 string = "e3"

	// WritableInterfaceTypeCiscoDashStackwise captures enum value "cisco-stackwise"
	WritableInterfaceTypeCiscoDashStackwise string = "cisco-stackwise"

	// WritableInterfaceTypeCiscoDashStackwiseDashPlus captures enum value "cisco-stackwise-plus"
	WritableInterfaceTypeCiscoDashStackwiseDashPlus string = "cisco-stackwise-plus"

	// WritableInterfaceTypeCiscoDashFlexstack captures enum value "cisco-flexstack"
	WritableInterfaceTypeCiscoDashFlexstack string = "cisco-flexstack"

	// WritableInterfaceTypeCiscoDashFlexstackDashPlus captures enum value "cisco-flexstack-plus"
	WritableInterfaceTypeCiscoDashFlexstackDashPlus string = "cisco-flexstack-plus"

	// WritableInterfaceTypeJuniperDashVcp captures enum value "juniper-vcp"
	WritableInterfaceTypeJuniperDashVcp string = "juniper-vcp"

	// WritableInterfaceTypeExtremeDashSummitstack captures enum value "extreme-summitstack"
	WritableInterfaceTypeExtremeDashSummitstack string = "extreme-summitstack"

	// WritableInterfaceTypeExtremeDashSummitstackDash128 captures enum value "extreme-summitstack-128"
	WritableInterfaceTypeExtremeDashSummitstackDash128 string = "extreme-summitstack-128"

	// WritableInterfaceTypeExtremeDashSummitstackDash256 captures enum value "extreme-summitstack-256"
	WritableInterfaceTypeExtremeDashSummitstackDash256 string = "extreme-summitstack-256"

	// WritableInterfaceTypeExtremeDashSummitstackDash512 captures enum value "extreme-summitstack-512"
	WritableInterfaceTypeExtremeDashSummitstackDash512 string = "extreme-summitstack-512"

	// WritableInterfaceTypeOther captures enum value "other"
	WritableInterfaceTypeOther string = "other"
)

// prop value enum
func (m *WritableInterface) validateTypeEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, writableInterfaceTypeTypePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *WritableInterface) validateType(formats strfmt.Registry) error {

	if err := validate.Required("type", "body", m.Type); err != nil {
		return err
	}

	// value enum
	if err := m.validateTypeEnum("type", "body", *m.Type); err != nil {
		return err
	}

	return nil
}

func (m *WritableInterface) validateUntaggedVlan(formats strfmt.Registry) error {
	if swag.IsZero(m.UntaggedVlan) { // not required
		return nil
	}

	if err := validate.FormatOf("untagged_vlan", "body", "uuid", m.UntaggedVlan.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *WritableInterface) validateURL(formats strfmt.Registry) error {
	if swag.IsZero(m.URL) { // not required
		return nil
	}

	if err := validate.FormatOf("url", "body", "uri", m.URL.String(), formats); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this writable interface based on the context it is used
func (m *WritableInterface) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
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

	if err := m.contextValidateCountIpaddresses(ctx, formats); err != nil {
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

func (m *WritableInterface) contextValidateCable(ctx context.Context, formats strfmt.Registry) error {

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

func (m *WritableInterface) contextValidateCablePeer(ctx context.Context, formats strfmt.Registry) error {

	return nil
}

func (m *WritableInterface) contextValidateCablePeerType(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "cable_peer_type", "body", string(m.CablePeerType)); err != nil {
		return err
	}

	return nil
}

func (m *WritableInterface) contextValidateComputedFields(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "computed_fields", "body", string(m.ComputedFields)); err != nil {
		return err
	}

	return nil
}

func (m *WritableInterface) contextValidateConnectedEndpoint(ctx context.Context, formats strfmt.Registry) error {

	return nil
}

func (m *WritableInterface) contextValidateConnectedEndpointReachable(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "connected_endpoint_reachable", "body", m.ConnectedEndpointReachable); err != nil {
		return err
	}

	return nil
}

func (m *WritableInterface) contextValidateConnectedEndpointType(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "connected_endpoint_type", "body", string(m.ConnectedEndpointType)); err != nil {
		return err
	}

	return nil
}

func (m *WritableInterface) contextValidateCountIpaddresses(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "count_ipaddresses", "body", int64(m.CountIpaddresses)); err != nil {
		return err
	}

	return nil
}

func (m *WritableInterface) contextValidateDisplay(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "display", "body", string(m.Display)); err != nil {
		return err
	}

	return nil
}

func (m *WritableInterface) contextValidateID(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "id", "body", strfmt.UUID(m.ID)); err != nil {
		return err
	}

	return nil
}

func (m *WritableInterface) contextValidateTags(ctx context.Context, formats strfmt.Registry) error {

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

func (m *WritableInterface) contextValidateURL(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "url", "body", strfmt.URI(m.URL)); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *WritableInterface) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *WritableInterface) UnmarshalBinary(b []byte) error {
	var res WritableInterface
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
