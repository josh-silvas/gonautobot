package models

import (
	"context"
	"encoding/json"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// InterfaceTemplate interface template
//
// swagger:model InterfaceTemplate
type InterfaceTemplate struct {

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

	// Management only
	MgmtOnly bool `json:"mgmt_only,omitempty"`

	// Name
	// Required: true
	// Max Length: 64
	// Min Length: 1
	Name *string `json:"name"`

	// type
	// Required: true
	Type *InterfaceTemplateType `json:"type"`

	// Url
	// Read Only: true
	// Format: uri
	URL strfmt.URI `json:"url,omitempty"`
}

// Validate validates this interface template
func (m *InterfaceTemplate) Validate(formats strfmt.Registry) error {
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

func (m *InterfaceTemplate) validateDescription(formats strfmt.Registry) error {
	if swag.IsZero(m.Description) { // not required
		return nil
	}

	if err := validate.MaxLength("description", "body", m.Description, 200); err != nil {
		return err
	}

	return nil
}

func (m *InterfaceTemplate) validateDeviceType(formats strfmt.Registry) error {

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

func (m *InterfaceTemplate) validateDisplay(formats strfmt.Registry) error {
	if swag.IsZero(m.Display) { // not required
		return nil
	}

	if err := validate.MinLength("display", "body", m.Display, 1); err != nil {
		return err
	}

	return nil
}

func (m *InterfaceTemplate) validateID(formats strfmt.Registry) error {
	if swag.IsZero(m.ID) { // not required
		return nil
	}

	if err := validate.FormatOf("id", "body", "uuid", m.ID.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *InterfaceTemplate) validateLabel(formats strfmt.Registry) error {
	if swag.IsZero(m.Label) { // not required
		return nil
	}

	if err := validate.MaxLength("label", "body", m.Label, 64); err != nil {
		return err
	}

	return nil
}

func (m *InterfaceTemplate) validateName(formats strfmt.Registry) error {

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

func (m *InterfaceTemplate) validateType(formats strfmt.Registry) error {

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

func (m *InterfaceTemplate) validateURL(formats strfmt.Registry) error {
	if swag.IsZero(m.URL) { // not required
		return nil
	}

	if err := validate.FormatOf("url", "body", "uri", m.URL.String(), formats); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this interface template based on the context it is used
func (m *InterfaceTemplate) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
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

func (m *InterfaceTemplate) contextValidateDeviceType(ctx context.Context, formats strfmt.Registry) error {

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

func (m *InterfaceTemplate) contextValidateDisplay(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "display", "body", string(m.Display)); err != nil {
		return err
	}

	return nil
}

func (m *InterfaceTemplate) contextValidateID(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "id", "body", strfmt.UUID(m.ID)); err != nil {
		return err
	}

	return nil
}

func (m *InterfaceTemplate) contextValidateType(ctx context.Context, formats strfmt.Registry) error {

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

func (m *InterfaceTemplate) contextValidateURL(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "url", "body", strfmt.URI(m.URL)); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *InterfaceTemplate) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *InterfaceTemplate) UnmarshalBinary(b []byte) error {
	var res InterfaceTemplate
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// InterfaceTemplateType Type
//
// swagger:model InterfaceTemplateType
type InterfaceTemplateType struct {

	// label
	// Required: true
	// Enum: [Virtual Link Aggregation Group (LAG) 100BASE-TX (10/100ME) 1000BASE-T (1GE) 2.5GBASE-T (2.5GE) 5GBASE-T (5GE) 10GBASE-T (10GE) 10GBASE-CX4 (10GE) GBIC (1GE) SFP (1GE) SFP+ (10GE) XFP (10GE) XENPAK (10GE) X2 (10GE) SFP28 (25GE) QSFP+ (40GE) QSFP28 (50GE) CFP (100GE) CFP2 (100GE) CFP2 (200GE) CFP4 (100GE) Cisco CPAK (100GE) QSFP28 (100GE) QSFP56 (200GE) QSFP-DD (400GE) OSFP (400GE) IEEE 802.11a IEEE 802.11b/g IEEE 802.11n IEEE 802.11ac IEEE 802.11ad IEEE 802.11ax GSM CDMA LTE OC-3/STM-1 OC-12/STM-4 OC-48/STM-16 OC-192/STM-64 OC-768/STM-256 OC-1920/STM-640 OC-3840/STM-1234 SFP (1GFC) SFP (2GFC) SFP (4GFC) SFP+ (8GFC) SFP+ (16GFC) SFP28 (32GFC) QSFP+ (64GFC) QSFP28 (128GFC) SDR (2 Gbps) DDR (4 Gbps) QDR (8 Gbps) FDR10 (10 Gbps) FDR (13.5 Gbps) EDR (25 Gbps) HDR (50 Gbps) NDR (100 Gbps) XDR (250 Gbps) T1 (1.544 Mbps) E1 (2.048 Mbps) T3 (45 Mbps) E3 (34 Mbps) Cisco StackWise Cisco StackWise Plus Cisco FlexStack Cisco FlexStack Plus Juniper VCP Extreme SummitStack Extreme SummitStack-128 Extreme SummitStack-256 Extreme SummitStack-512 Other]
	Label *string `json:"label"`

	// value
	// Required: true
	// Enum: [virtual lag 100base-tx 1000base-t 2.5gbase-t 5gbase-t 10gbase-t 10gbase-cx4 1000base-x-gbic 1000base-x-sfp 10gbase-x-sfpp 10gbase-x-xfp 10gbase-x-xenpak 10gbase-x-x2 25gbase-x-sfp28 40gbase-x-qsfpp 50gbase-x-sfp28 100gbase-x-cfp 100gbase-x-cfp2 200gbase-x-cfp2 100gbase-x-cfp4 100gbase-x-cpak 100gbase-x-qsfp28 200gbase-x-qsfp56 400gbase-x-qsfpdd 400gbase-x-osfp ieee802.11a ieee802.11g ieee802.11n ieee802.11ac ieee802.11ad ieee802.11ax gsm cdma lte sonet-oc3 sonet-oc12 sonet-oc48 sonet-oc192 sonet-oc768 sonet-oc1920 sonet-oc3840 1gfc-sfp 2gfc-sfp 4gfc-sfp 8gfc-sfpp 16gfc-sfpp 32gfc-sfp28 64gfc-qsfpp 128gfc-sfp28 infiniband-sdr infiniband-ddr infiniband-qdr infiniband-fdr10 infiniband-fdr infiniband-edr infiniband-hdr infiniband-ndr infiniband-xdr t1 e1 t3 e3 cisco-stackwise cisco-stackwise-plus cisco-flexstack cisco-flexstack-plus juniper-vcp extreme-summitstack extreme-summitstack-128 extreme-summitstack-256 extreme-summitstack-512 other]
	Value *string `json:"value"`
}

// Validate validates this interface template type
func (m *InterfaceTemplateType) Validate(formats strfmt.Registry) error {
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

var interfaceTemplateTypeTypeLabelPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["Virtual","Link Aggregation Group (LAG)","100BASE-TX (10/100ME)","1000BASE-T (1GE)","2.5GBASE-T (2.5GE)","5GBASE-T (5GE)","10GBASE-T (10GE)","10GBASE-CX4 (10GE)","GBIC (1GE)","SFP (1GE)","SFP+ (10GE)","XFP (10GE)","XENPAK (10GE)","X2 (10GE)","SFP28 (25GE)","QSFP+ (40GE)","QSFP28 (50GE)","CFP (100GE)","CFP2 (100GE)","CFP2 (200GE)","CFP4 (100GE)","Cisco CPAK (100GE)","QSFP28 (100GE)","QSFP56 (200GE)","QSFP-DD (400GE)","OSFP (400GE)","IEEE 802.11a","IEEE 802.11b/g","IEEE 802.11n","IEEE 802.11ac","IEEE 802.11ad","IEEE 802.11ax","GSM","CDMA","LTE","OC-3/STM-1","OC-12/STM-4","OC-48/STM-16","OC-192/STM-64","OC-768/STM-256","OC-1920/STM-640","OC-3840/STM-1234","SFP (1GFC)","SFP (2GFC)","SFP (4GFC)","SFP+ (8GFC)","SFP+ (16GFC)","SFP28 (32GFC)","QSFP+ (64GFC)","QSFP28 (128GFC)","SDR (2 Gbps)","DDR (4 Gbps)","QDR (8 Gbps)","FDR10 (10 Gbps)","FDR (13.5 Gbps)","EDR (25 Gbps)","HDR (50 Gbps)","NDR (100 Gbps)","XDR (250 Gbps)","T1 (1.544 Mbps)","E1 (2.048 Mbps)","T3 (45 Mbps)","E3 (34 Mbps)","Cisco StackWise","Cisco StackWise Plus","Cisco FlexStack","Cisco FlexStack Plus","Juniper VCP","Extreme SummitStack","Extreme SummitStack-128","Extreme SummitStack-256","Extreme SummitStack-512","Other"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		interfaceTemplateTypeTypeLabelPropEnum = append(interfaceTemplateTypeTypeLabelPropEnum, v)
	}
}

const (

	// InterfaceTemplateTypeLabelVirtual captures enum value "Virtual"
	InterfaceTemplateTypeLabelVirtual string = "Virtual"

	// InterfaceTemplateTypeLabelLinkAggregationGroupLAG captures enum value "Link Aggregation Group (LAG)"
	InterfaceTemplateTypeLabelLinkAggregationGroupLAG string = "Link Aggregation Group (LAG)"

	// InterfaceTemplateTypeLabelNr100BASEDashTX10100ME captures enum value "100BASE-TX (10/100ME)"
	InterfaceTemplateTypeLabelNr100BASEDashTX10100ME string = "100BASE-TX (10/100ME)"

	// InterfaceTemplateTypeLabelNr1000BASEDashT1GE captures enum value "1000BASE-T (1GE)"
	InterfaceTemplateTypeLabelNr1000BASEDashT1GE string = "1000BASE-T (1GE)"

	// InterfaceTemplateTypeLabelNr2Dot5GBASEDashT2Dot5GE captures enum value "2.5GBASE-T (2.5GE)"
	InterfaceTemplateTypeLabelNr2Dot5GBASEDashT2Dot5GE string = "2.5GBASE-T (2.5GE)"

	// InterfaceTemplateTypeLabelNr5GBASEDashT5GE captures enum value "5GBASE-T (5GE)"
	InterfaceTemplateTypeLabelNr5GBASEDashT5GE string = "5GBASE-T (5GE)"

	// InterfaceTemplateTypeLabelNr10GBASEDashT10GE captures enum value "10GBASE-T (10GE)"
	InterfaceTemplateTypeLabelNr10GBASEDashT10GE string = "10GBASE-T (10GE)"

	// InterfaceTemplateTypeLabelNr10GBASEDashCX410GE captures enum value "10GBASE-CX4 (10GE)"
	InterfaceTemplateTypeLabelNr10GBASEDashCX410GE string = "10GBASE-CX4 (10GE)"

	// InterfaceTemplateTypeLabelGBIC1GE captures enum value "GBIC (1GE)"
	InterfaceTemplateTypeLabelGBIC1GE string = "GBIC (1GE)"

	// InterfaceTemplateTypeLabelSFP1GE captures enum value "SFP (1GE)"
	InterfaceTemplateTypeLabelSFP1GE string = "SFP (1GE)"

	// InterfaceTemplateTypeLabelSFPPlus10GE captures enum value "SFP+ (10GE)"
	InterfaceTemplateTypeLabelSFPPlus10GE string = "SFP+ (10GE)"

	// InterfaceTemplateTypeLabelXFP10GE captures enum value "XFP (10GE)"
	InterfaceTemplateTypeLabelXFP10GE string = "XFP (10GE)"

	// InterfaceTemplateTypeLabelXENPAK10GE captures enum value "XENPAK (10GE)"
	InterfaceTemplateTypeLabelXENPAK10GE string = "XENPAK (10GE)"

	// InterfaceTemplateTypeLabelX210GE captures enum value "X2 (10GE)"
	InterfaceTemplateTypeLabelX210GE string = "X2 (10GE)"

	// InterfaceTemplateTypeLabelSFP2825GE captures enum value "SFP28 (25GE)"
	InterfaceTemplateTypeLabelSFP2825GE string = "SFP28 (25GE)"

	// InterfaceTemplateTypeLabelQSFPPlus40GE captures enum value "QSFP+ (40GE)"
	InterfaceTemplateTypeLabelQSFPPlus40GE string = "QSFP+ (40GE)"

	// InterfaceTemplateTypeLabelQSFP2850GE captures enum value "QSFP28 (50GE)"
	InterfaceTemplateTypeLabelQSFP2850GE string = "QSFP28 (50GE)"

	// InterfaceTemplateTypeLabelCFP100GE captures enum value "CFP (100GE)"
	InterfaceTemplateTypeLabelCFP100GE string = "CFP (100GE)"

	// InterfaceTemplateTypeLabelCFP2100GE captures enum value "CFP2 (100GE)"
	InterfaceTemplateTypeLabelCFP2100GE string = "CFP2 (100GE)"

	// InterfaceTemplateTypeLabelCFP2200GE captures enum value "CFP2 (200GE)"
	InterfaceTemplateTypeLabelCFP2200GE string = "CFP2 (200GE)"

	// InterfaceTemplateTypeLabelCFP4100GE captures enum value "CFP4 (100GE)"
	InterfaceTemplateTypeLabelCFP4100GE string = "CFP4 (100GE)"

	// InterfaceTemplateTypeLabelCiscoCPAK100GE captures enum value "Cisco CPAK (100GE)"
	InterfaceTemplateTypeLabelCiscoCPAK100GE string = "Cisco CPAK (100GE)"

	// InterfaceTemplateTypeLabelQSFP28100GE captures enum value "QSFP28 (100GE)"
	InterfaceTemplateTypeLabelQSFP28100GE string = "QSFP28 (100GE)"

	// InterfaceTemplateTypeLabelQSFP56200GE captures enum value "QSFP56 (200GE)"
	InterfaceTemplateTypeLabelQSFP56200GE string = "QSFP56 (200GE)"

	// InterfaceTemplateTypeLabelQSFPDashDD400GE captures enum value "QSFP-DD (400GE)"
	InterfaceTemplateTypeLabelQSFPDashDD400GE string = "QSFP-DD (400GE)"

	// InterfaceTemplateTypeLabelOSFP400GE captures enum value "OSFP (400GE)"
	InterfaceTemplateTypeLabelOSFP400GE string = "OSFP (400GE)"

	// InterfaceTemplateTypeLabelIEEE802Dot11a captures enum value "IEEE 802.11a"
	InterfaceTemplateTypeLabelIEEE802Dot11a string = "IEEE 802.11a"

	// InterfaceTemplateTypeLabelIEEE802Dot11bg captures enum value "IEEE 802.11b/g"
	InterfaceTemplateTypeLabelIEEE802Dot11bg string = "IEEE 802.11b/g"

	// InterfaceTemplateTypeLabelIEEE802Dot11n captures enum value "IEEE 802.11n"
	InterfaceTemplateTypeLabelIEEE802Dot11n string = "IEEE 802.11n"

	// InterfaceTemplateTypeLabelIEEE802Dot11ac captures enum value "IEEE 802.11ac"
	InterfaceTemplateTypeLabelIEEE802Dot11ac string = "IEEE 802.11ac"

	// InterfaceTemplateTypeLabelIEEE802Dot11ad captures enum value "IEEE 802.11ad"
	InterfaceTemplateTypeLabelIEEE802Dot11ad string = "IEEE 802.11ad"

	// InterfaceTemplateTypeLabelIEEE802Dot11ax captures enum value "IEEE 802.11ax"
	InterfaceTemplateTypeLabelIEEE802Dot11ax string = "IEEE 802.11ax"

	// InterfaceTemplateTypeLabelGSM captures enum value "GSM"
	InterfaceTemplateTypeLabelGSM string = "GSM"

	// InterfaceTemplateTypeLabelCDMA captures enum value "CDMA"
	InterfaceTemplateTypeLabelCDMA string = "CDMA"

	// InterfaceTemplateTypeLabelLTE captures enum value "LTE"
	InterfaceTemplateTypeLabelLTE string = "LTE"

	// InterfaceTemplateTypeLabelOCDash3STMDash1 captures enum value "OC-3/STM-1"
	InterfaceTemplateTypeLabelOCDash3STMDash1 string = "OC-3/STM-1"

	// InterfaceTemplateTypeLabelOCDash12STMDash4 captures enum value "OC-12/STM-4"
	InterfaceTemplateTypeLabelOCDash12STMDash4 string = "OC-12/STM-4"

	// InterfaceTemplateTypeLabelOCDash48STMDash16 captures enum value "OC-48/STM-16"
	InterfaceTemplateTypeLabelOCDash48STMDash16 string = "OC-48/STM-16"

	// InterfaceTemplateTypeLabelOCDash192STMDash64 captures enum value "OC-192/STM-64"
	InterfaceTemplateTypeLabelOCDash192STMDash64 string = "OC-192/STM-64"

	// InterfaceTemplateTypeLabelOCDash768STMDash256 captures enum value "OC-768/STM-256"
	InterfaceTemplateTypeLabelOCDash768STMDash256 string = "OC-768/STM-256"

	// InterfaceTemplateTypeLabelOCDash1920STMDash640 captures enum value "OC-1920/STM-640"
	InterfaceTemplateTypeLabelOCDash1920STMDash640 string = "OC-1920/STM-640"

	// InterfaceTemplateTypeLabelOCDash3840STMDash1234 captures enum value "OC-3840/STM-1234"
	InterfaceTemplateTypeLabelOCDash3840STMDash1234 string = "OC-3840/STM-1234"

	// InterfaceTemplateTypeLabelSFP1GFC captures enum value "SFP (1GFC)"
	InterfaceTemplateTypeLabelSFP1GFC string = "SFP (1GFC)"

	// InterfaceTemplateTypeLabelSFP2GFC captures enum value "SFP (2GFC)"
	InterfaceTemplateTypeLabelSFP2GFC string = "SFP (2GFC)"

	// InterfaceTemplateTypeLabelSFP4GFC captures enum value "SFP (4GFC)"
	InterfaceTemplateTypeLabelSFP4GFC string = "SFP (4GFC)"

	// InterfaceTemplateTypeLabelSFPPlus8GFC captures enum value "SFP+ (8GFC)"
	InterfaceTemplateTypeLabelSFPPlus8GFC string = "SFP+ (8GFC)"

	// InterfaceTemplateTypeLabelSFPPlus16GFC captures enum value "SFP+ (16GFC)"
	InterfaceTemplateTypeLabelSFPPlus16GFC string = "SFP+ (16GFC)"

	// InterfaceTemplateTypeLabelSFP2832GFC captures enum value "SFP28 (32GFC)"
	InterfaceTemplateTypeLabelSFP2832GFC string = "SFP28 (32GFC)"

	// InterfaceTemplateTypeLabelQSFPPlus64GFC captures enum value "QSFP+ (64GFC)"
	InterfaceTemplateTypeLabelQSFPPlus64GFC string = "QSFP+ (64GFC)"

	// InterfaceTemplateTypeLabelQSFP28128GFC captures enum value "QSFP28 (128GFC)"
	InterfaceTemplateTypeLabelQSFP28128GFC string = "QSFP28 (128GFC)"

	// InterfaceTemplateTypeLabelSDR2Gbps captures enum value "SDR (2 Gbps)"
	InterfaceTemplateTypeLabelSDR2Gbps string = "SDR (2 Gbps)"

	// InterfaceTemplateTypeLabelDDR4Gbps captures enum value "DDR (4 Gbps)"
	InterfaceTemplateTypeLabelDDR4Gbps string = "DDR (4 Gbps)"

	// InterfaceTemplateTypeLabelQDR8Gbps captures enum value "QDR (8 Gbps)"
	InterfaceTemplateTypeLabelQDR8Gbps string = "QDR (8 Gbps)"

	// InterfaceTemplateTypeLabelFDR1010Gbps captures enum value "FDR10 (10 Gbps)"
	InterfaceTemplateTypeLabelFDR1010Gbps string = "FDR10 (10 Gbps)"

	// InterfaceTemplateTypeLabelFDR13Dot5Gbps captures enum value "FDR (13.5 Gbps)"
	InterfaceTemplateTypeLabelFDR13Dot5Gbps string = "FDR (13.5 Gbps)"

	// InterfaceTemplateTypeLabelEDR25Gbps captures enum value "EDR (25 Gbps)"
	InterfaceTemplateTypeLabelEDR25Gbps string = "EDR (25 Gbps)"

	// InterfaceTemplateTypeLabelHDR50Gbps captures enum value "HDR (50 Gbps)"
	InterfaceTemplateTypeLabelHDR50Gbps string = "HDR (50 Gbps)"

	// InterfaceTemplateTypeLabelNDR100Gbps captures enum value "NDR (100 Gbps)"
	InterfaceTemplateTypeLabelNDR100Gbps string = "NDR (100 Gbps)"

	// InterfaceTemplateTypeLabelXDR250Gbps captures enum value "XDR (250 Gbps)"
	InterfaceTemplateTypeLabelXDR250Gbps string = "XDR (250 Gbps)"

	// InterfaceTemplateTypeLabelT11Dot544Mbps captures enum value "T1 (1.544 Mbps)"
	InterfaceTemplateTypeLabelT11Dot544Mbps string = "T1 (1.544 Mbps)"

	// InterfaceTemplateTypeLabelE12Dot048Mbps captures enum value "E1 (2.048 Mbps)"
	InterfaceTemplateTypeLabelE12Dot048Mbps string = "E1 (2.048 Mbps)"

	// InterfaceTemplateTypeLabelT345Mbps captures enum value "T3 (45 Mbps)"
	InterfaceTemplateTypeLabelT345Mbps string = "T3 (45 Mbps)"

	// InterfaceTemplateTypeLabelE334Mbps captures enum value "E3 (34 Mbps)"
	InterfaceTemplateTypeLabelE334Mbps string = "E3 (34 Mbps)"

	// InterfaceTemplateTypeLabelCiscoStackWise captures enum value "Cisco StackWise"
	InterfaceTemplateTypeLabelCiscoStackWise string = "Cisco StackWise"

	// InterfaceTemplateTypeLabelCiscoStackWisePlus captures enum value "Cisco StackWise Plus"
	InterfaceTemplateTypeLabelCiscoStackWisePlus string = "Cisco StackWise Plus"

	// InterfaceTemplateTypeLabelCiscoFlexStack captures enum value "Cisco FlexStack"
	InterfaceTemplateTypeLabelCiscoFlexStack string = "Cisco FlexStack"

	// InterfaceTemplateTypeLabelCiscoFlexStackPlus captures enum value "Cisco FlexStack Plus"
	InterfaceTemplateTypeLabelCiscoFlexStackPlus string = "Cisco FlexStack Plus"

	// InterfaceTemplateTypeLabelJuniperVCP captures enum value "Juniper VCP"
	InterfaceTemplateTypeLabelJuniperVCP string = "Juniper VCP"

	// InterfaceTemplateTypeLabelExtremeSummitStack captures enum value "Extreme SummitStack"
	InterfaceTemplateTypeLabelExtremeSummitStack string = "Extreme SummitStack"

	// InterfaceTemplateTypeLabelExtremeSummitStackDash128 captures enum value "Extreme SummitStack-128"
	InterfaceTemplateTypeLabelExtremeSummitStackDash128 string = "Extreme SummitStack-128"

	// InterfaceTemplateTypeLabelExtremeSummitStackDash256 captures enum value "Extreme SummitStack-256"
	InterfaceTemplateTypeLabelExtremeSummitStackDash256 string = "Extreme SummitStack-256"

	// InterfaceTemplateTypeLabelExtremeSummitStackDash512 captures enum value "Extreme SummitStack-512"
	InterfaceTemplateTypeLabelExtremeSummitStackDash512 string = "Extreme SummitStack-512"

	// InterfaceTemplateTypeLabelOther captures enum value "Other"
	InterfaceTemplateTypeLabelOther string = "Other"
)

// prop value enum
func (m *InterfaceTemplateType) validateLabelEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, interfaceTemplateTypeTypeLabelPropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *InterfaceTemplateType) validateLabel(formats strfmt.Registry) error {

	if err := validate.Required("type"+"."+"label", "body", m.Label); err != nil {
		return err
	}

	// value enum
	if err := m.validateLabelEnum("type"+"."+"label", "body", *m.Label); err != nil {
		return err
	}

	return nil
}

var interfaceTemplateTypeTypeValuePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["virtual","lag","100base-tx","1000base-t","2.5gbase-t","5gbase-t","10gbase-t","10gbase-cx4","1000base-x-gbic","1000base-x-sfp","10gbase-x-sfpp","10gbase-x-xfp","10gbase-x-xenpak","10gbase-x-x2","25gbase-x-sfp28","40gbase-x-qsfpp","50gbase-x-sfp28","100gbase-x-cfp","100gbase-x-cfp2","200gbase-x-cfp2","100gbase-x-cfp4","100gbase-x-cpak","100gbase-x-qsfp28","200gbase-x-qsfp56","400gbase-x-qsfpdd","400gbase-x-osfp","ieee802.11a","ieee802.11g","ieee802.11n","ieee802.11ac","ieee802.11ad","ieee802.11ax","gsm","cdma","lte","sonet-oc3","sonet-oc12","sonet-oc48","sonet-oc192","sonet-oc768","sonet-oc1920","sonet-oc3840","1gfc-sfp","2gfc-sfp","4gfc-sfp","8gfc-sfpp","16gfc-sfpp","32gfc-sfp28","64gfc-qsfpp","128gfc-sfp28","infiniband-sdr","infiniband-ddr","infiniband-qdr","infiniband-fdr10","infiniband-fdr","infiniband-edr","infiniband-hdr","infiniband-ndr","infiniband-xdr","t1","e1","t3","e3","cisco-stackwise","cisco-stackwise-plus","cisco-flexstack","cisco-flexstack-plus","juniper-vcp","extreme-summitstack","extreme-summitstack-128","extreme-summitstack-256","extreme-summitstack-512","other"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		interfaceTemplateTypeTypeValuePropEnum = append(interfaceTemplateTypeTypeValuePropEnum, v)
	}
}

const (

	// InterfaceTemplateTypeValueVirtual captures enum value "virtual"
	InterfaceTemplateTypeValueVirtual string = "virtual"

	// InterfaceTemplateTypeValueLag captures enum value "lag"
	InterfaceTemplateTypeValueLag string = "lag"

	// InterfaceTemplateTypeValueNr100baseDashTx captures enum value "100base-tx"
	InterfaceTemplateTypeValueNr100baseDashTx string = "100base-tx"

	// InterfaceTemplateTypeValueNr1000baseDasht captures enum value "1000base-t"
	InterfaceTemplateTypeValueNr1000baseDasht string = "1000base-t"

	// InterfaceTemplateTypeValueNr2Dot5gbaseDasht captures enum value "2.5gbase-t"
	InterfaceTemplateTypeValueNr2Dot5gbaseDasht string = "2.5gbase-t"

	// InterfaceTemplateTypeValueNr5gbaseDasht captures enum value "5gbase-t"
	InterfaceTemplateTypeValueNr5gbaseDasht string = "5gbase-t"

	// InterfaceTemplateTypeValueNr10gbaseDasht captures enum value "10gbase-t"
	InterfaceTemplateTypeValueNr10gbaseDasht string = "10gbase-t"

	// InterfaceTemplateTypeValueNr10gbaseDashCx4 captures enum value "10gbase-cx4"
	InterfaceTemplateTypeValueNr10gbaseDashCx4 string = "10gbase-cx4"

	// InterfaceTemplateTypeValueNr1000baseDashxDashGbic captures enum value "1000base-x-gbic"
	InterfaceTemplateTypeValueNr1000baseDashxDashGbic string = "1000base-x-gbic"

	// InterfaceTemplateTypeValueNr1000baseDashxDashSfp captures enum value "1000base-x-sfp"
	InterfaceTemplateTypeValueNr1000baseDashxDashSfp string = "1000base-x-sfp"

	// InterfaceTemplateTypeValueNr10gbaseDashxDashSfpp captures enum value "10gbase-x-sfpp"
	InterfaceTemplateTypeValueNr10gbaseDashxDashSfpp string = "10gbase-x-sfpp"

	// InterfaceTemplateTypeValueNr10gbaseDashxDashXfp captures enum value "10gbase-x-xfp"
	InterfaceTemplateTypeValueNr10gbaseDashxDashXfp string = "10gbase-x-xfp"

	// InterfaceTemplateTypeValueNr10gbaseDashxDashXenpak captures enum value "10gbase-x-xenpak"
	InterfaceTemplateTypeValueNr10gbaseDashxDashXenpak string = "10gbase-x-xenpak"

	// InterfaceTemplateTypeValueNr10gbaseDashxDashX2 captures enum value "10gbase-x-x2"
	InterfaceTemplateTypeValueNr10gbaseDashxDashX2 string = "10gbase-x-x2"

	// InterfaceTemplateTypeValueNr25gbaseDashxDashSfp28 captures enum value "25gbase-x-sfp28"
	InterfaceTemplateTypeValueNr25gbaseDashxDashSfp28 string = "25gbase-x-sfp28"

	// InterfaceTemplateTypeValueNr40gbaseDashxDashQsfpp captures enum value "40gbase-x-qsfpp"
	InterfaceTemplateTypeValueNr40gbaseDashxDashQsfpp string = "40gbase-x-qsfpp"

	// InterfaceTemplateTypeValueNr50gbaseDashxDashSfp28 captures enum value "50gbase-x-sfp28"
	InterfaceTemplateTypeValueNr50gbaseDashxDashSfp28 string = "50gbase-x-sfp28"

	// InterfaceTemplateTypeValueNr100gbaseDashxDashCfp captures enum value "100gbase-x-cfp"
	InterfaceTemplateTypeValueNr100gbaseDashxDashCfp string = "100gbase-x-cfp"

	// InterfaceTemplateTypeValueNr100gbaseDashxDashCfp2 captures enum value "100gbase-x-cfp2"
	InterfaceTemplateTypeValueNr100gbaseDashxDashCfp2 string = "100gbase-x-cfp2"

	// InterfaceTemplateTypeValueNr200gbaseDashxDashCfp2 captures enum value "200gbase-x-cfp2"
	InterfaceTemplateTypeValueNr200gbaseDashxDashCfp2 string = "200gbase-x-cfp2"

	// InterfaceTemplateTypeValueNr100gbaseDashxDashCfp4 captures enum value "100gbase-x-cfp4"
	InterfaceTemplateTypeValueNr100gbaseDashxDashCfp4 string = "100gbase-x-cfp4"

	// InterfaceTemplateTypeValueNr100gbaseDashxDashCpak captures enum value "100gbase-x-cpak"
	InterfaceTemplateTypeValueNr100gbaseDashxDashCpak string = "100gbase-x-cpak"

	// InterfaceTemplateTypeValueNr100gbaseDashxDashQsfp28 captures enum value "100gbase-x-qsfp28"
	InterfaceTemplateTypeValueNr100gbaseDashxDashQsfp28 string = "100gbase-x-qsfp28"

	// InterfaceTemplateTypeValueNr200gbaseDashxDashQsfp56 captures enum value "200gbase-x-qsfp56"
	InterfaceTemplateTypeValueNr200gbaseDashxDashQsfp56 string = "200gbase-x-qsfp56"

	// InterfaceTemplateTypeValueNr400gbaseDashxDashQsfpdd captures enum value "400gbase-x-qsfpdd"
	InterfaceTemplateTypeValueNr400gbaseDashxDashQsfpdd string = "400gbase-x-qsfpdd"

	// InterfaceTemplateTypeValueNr400gbaseDashxDashOsfp captures enum value "400gbase-x-osfp"
	InterfaceTemplateTypeValueNr400gbaseDashxDashOsfp string = "400gbase-x-osfp"

	// InterfaceTemplateTypeValueIeee802Dot11a captures enum value "ieee802.11a"
	InterfaceTemplateTypeValueIeee802Dot11a string = "ieee802.11a"

	// InterfaceTemplateTypeValueIeee802Dot11g captures enum value "ieee802.11g"
	InterfaceTemplateTypeValueIeee802Dot11g string = "ieee802.11g"

	// InterfaceTemplateTypeValueIeee802Dot11n captures enum value "ieee802.11n"
	InterfaceTemplateTypeValueIeee802Dot11n string = "ieee802.11n"

	// InterfaceTemplateTypeValueIeee802Dot11ac captures enum value "ieee802.11ac"
	InterfaceTemplateTypeValueIeee802Dot11ac string = "ieee802.11ac"

	// InterfaceTemplateTypeValueIeee802Dot11ad captures enum value "ieee802.11ad"
	InterfaceTemplateTypeValueIeee802Dot11ad string = "ieee802.11ad"

	// InterfaceTemplateTypeValueIeee802Dot11ax captures enum value "ieee802.11ax"
	InterfaceTemplateTypeValueIeee802Dot11ax string = "ieee802.11ax"

	// InterfaceTemplateTypeValueGsm captures enum value "gsm"
	InterfaceTemplateTypeValueGsm string = "gsm"

	// InterfaceTemplateTypeValueCdma captures enum value "cdma"
	InterfaceTemplateTypeValueCdma string = "cdma"

	// InterfaceTemplateTypeValueLte captures enum value "lte"
	InterfaceTemplateTypeValueLte string = "lte"

	// InterfaceTemplateTypeValueSonetDashOc3 captures enum value "sonet-oc3"
	InterfaceTemplateTypeValueSonetDashOc3 string = "sonet-oc3"

	// InterfaceTemplateTypeValueSonetDashOc12 captures enum value "sonet-oc12"
	InterfaceTemplateTypeValueSonetDashOc12 string = "sonet-oc12"

	// InterfaceTemplateTypeValueSonetDashOc48 captures enum value "sonet-oc48"
	InterfaceTemplateTypeValueSonetDashOc48 string = "sonet-oc48"

	// InterfaceTemplateTypeValueSonetDashOc192 captures enum value "sonet-oc192"
	InterfaceTemplateTypeValueSonetDashOc192 string = "sonet-oc192"

	// InterfaceTemplateTypeValueSonetDashOc768 captures enum value "sonet-oc768"
	InterfaceTemplateTypeValueSonetDashOc768 string = "sonet-oc768"

	// InterfaceTemplateTypeValueSonetDashOc1920 captures enum value "sonet-oc1920"
	InterfaceTemplateTypeValueSonetDashOc1920 string = "sonet-oc1920"

	// InterfaceTemplateTypeValueSonetDashOc3840 captures enum value "sonet-oc3840"
	InterfaceTemplateTypeValueSonetDashOc3840 string = "sonet-oc3840"

	// InterfaceTemplateTypeValueNr1gfcDashSfp captures enum value "1gfc-sfp"
	InterfaceTemplateTypeValueNr1gfcDashSfp string = "1gfc-sfp"

	// InterfaceTemplateTypeValueNr2gfcDashSfp captures enum value "2gfc-sfp"
	InterfaceTemplateTypeValueNr2gfcDashSfp string = "2gfc-sfp"

	// InterfaceTemplateTypeValueNr4gfcDashSfp captures enum value "4gfc-sfp"
	InterfaceTemplateTypeValueNr4gfcDashSfp string = "4gfc-sfp"

	// InterfaceTemplateTypeValueNr8gfcDashSfpp captures enum value "8gfc-sfpp"
	InterfaceTemplateTypeValueNr8gfcDashSfpp string = "8gfc-sfpp"

	// InterfaceTemplateTypeValueNr16gfcDashSfpp captures enum value "16gfc-sfpp"
	InterfaceTemplateTypeValueNr16gfcDashSfpp string = "16gfc-sfpp"

	// InterfaceTemplateTypeValueNr32gfcDashSfp28 captures enum value "32gfc-sfp28"
	InterfaceTemplateTypeValueNr32gfcDashSfp28 string = "32gfc-sfp28"

	// InterfaceTemplateTypeValueNr64gfcDashQsfpp captures enum value "64gfc-qsfpp"
	InterfaceTemplateTypeValueNr64gfcDashQsfpp string = "64gfc-qsfpp"

	// InterfaceTemplateTypeValueNr128gfcDashSfp28 captures enum value "128gfc-sfp28"
	InterfaceTemplateTypeValueNr128gfcDashSfp28 string = "128gfc-sfp28"

	// InterfaceTemplateTypeValueInfinibandDashSdr captures enum value "infiniband-sdr"
	InterfaceTemplateTypeValueInfinibandDashSdr string = "infiniband-sdr"

	// InterfaceTemplateTypeValueInfinibandDashDdr captures enum value "infiniband-ddr"
	InterfaceTemplateTypeValueInfinibandDashDdr string = "infiniband-ddr"

	// InterfaceTemplateTypeValueInfinibandDashQdr captures enum value "infiniband-qdr"
	InterfaceTemplateTypeValueInfinibandDashQdr string = "infiniband-qdr"

	// InterfaceTemplateTypeValueInfinibandDashFdr10 captures enum value "infiniband-fdr10"
	InterfaceTemplateTypeValueInfinibandDashFdr10 string = "infiniband-fdr10"

	// InterfaceTemplateTypeValueInfinibandDashFdr captures enum value "infiniband-fdr"
	InterfaceTemplateTypeValueInfinibandDashFdr string = "infiniband-fdr"

	// InterfaceTemplateTypeValueInfinibandDashEdr captures enum value "infiniband-edr"
	InterfaceTemplateTypeValueInfinibandDashEdr string = "infiniband-edr"

	// InterfaceTemplateTypeValueInfinibandDashHdr captures enum value "infiniband-hdr"
	InterfaceTemplateTypeValueInfinibandDashHdr string = "infiniband-hdr"

	// InterfaceTemplateTypeValueInfinibandDashNdr captures enum value "infiniband-ndr"
	InterfaceTemplateTypeValueInfinibandDashNdr string = "infiniband-ndr"

	// InterfaceTemplateTypeValueInfinibandDashXdr captures enum value "infiniband-xdr"
	InterfaceTemplateTypeValueInfinibandDashXdr string = "infiniband-xdr"

	// InterfaceTemplateTypeValueT1 captures enum value "t1"
	InterfaceTemplateTypeValueT1 string = "t1"

	// InterfaceTemplateTypeValueE1 captures enum value "e1"
	InterfaceTemplateTypeValueE1 string = "e1"

	// InterfaceTemplateTypeValueT3 captures enum value "t3"
	InterfaceTemplateTypeValueT3 string = "t3"

	// InterfaceTemplateTypeValueE3 captures enum value "e3"
	InterfaceTemplateTypeValueE3 string = "e3"

	// InterfaceTemplateTypeValueCiscoDashStackwise captures enum value "cisco-stackwise"
	InterfaceTemplateTypeValueCiscoDashStackwise string = "cisco-stackwise"

	// InterfaceTemplateTypeValueCiscoDashStackwiseDashPlus captures enum value "cisco-stackwise-plus"
	InterfaceTemplateTypeValueCiscoDashStackwiseDashPlus string = "cisco-stackwise-plus"

	// InterfaceTemplateTypeValueCiscoDashFlexstack captures enum value "cisco-flexstack"
	InterfaceTemplateTypeValueCiscoDashFlexstack string = "cisco-flexstack"

	// InterfaceTemplateTypeValueCiscoDashFlexstackDashPlus captures enum value "cisco-flexstack-plus"
	InterfaceTemplateTypeValueCiscoDashFlexstackDashPlus string = "cisco-flexstack-plus"

	// InterfaceTemplateTypeValueJuniperDashVcp captures enum value "juniper-vcp"
	InterfaceTemplateTypeValueJuniperDashVcp string = "juniper-vcp"

	// InterfaceTemplateTypeValueExtremeDashSummitstack captures enum value "extreme-summitstack"
	InterfaceTemplateTypeValueExtremeDashSummitstack string = "extreme-summitstack"

	// InterfaceTemplateTypeValueExtremeDashSummitstackDash128 captures enum value "extreme-summitstack-128"
	InterfaceTemplateTypeValueExtremeDashSummitstackDash128 string = "extreme-summitstack-128"

	// InterfaceTemplateTypeValueExtremeDashSummitstackDash256 captures enum value "extreme-summitstack-256"
	InterfaceTemplateTypeValueExtremeDashSummitstackDash256 string = "extreme-summitstack-256"

	// InterfaceTemplateTypeValueExtremeDashSummitstackDash512 captures enum value "extreme-summitstack-512"
	InterfaceTemplateTypeValueExtremeDashSummitstackDash512 string = "extreme-summitstack-512"

	// InterfaceTemplateTypeValueOther captures enum value "other"
	InterfaceTemplateTypeValueOther string = "other"
)

// prop value enum
func (m *InterfaceTemplateType) validateValueEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, interfaceTemplateTypeTypeValuePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *InterfaceTemplateType) validateValue(formats strfmt.Registry) error {

	if err := validate.Required("type"+"."+"value", "body", m.Value); err != nil {
		return err
	}

	// value enum
	if err := m.validateValueEnum("type"+"."+"value", "body", *m.Value); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this interface template type based on context it is used
func (m *InterfaceTemplateType) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *InterfaceTemplateType) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *InterfaceTemplateType) UnmarshalBinary(b []byte) error {
	var res InterfaceTemplateType
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
