package models

import (
	"context"
	"encoding/json"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// WritablePowerPortTemplate writable power port template
//
// swagger:model WritablePowerPortTemplate
type WritablePowerPortTemplate struct {

	// Allocated draw
	//
	// Allocated power draw (watts)
	// Maximum: 32767
	// Minimum: 1
	AllocatedDraw *int64 `json:"allocated_draw,omitempty"`

	// Computed fields
	// Read Only: true
	ComputedFields string `json:"computed_fields,omitempty"`

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

	// Type
	// Enum: [iec-60320-c6 iec-60320-c8 iec-60320-c14 iec-60320-c16 iec-60320-c20 iec-60309-p-n-e-4h iec-60309-p-n-e-6h iec-60309-p-n-e-9h iec-60309-2p-e-4h iec-60309-2p-e-6h iec-60309-2p-e-9h iec-60309-3p-e-4h iec-60309-3p-e-6h iec-60309-3p-e-9h iec-60309-3p-n-e-4h iec-60309-3p-n-e-6h iec-60309-3p-n-e-9h nema-1-15p nema-5-15p nema-5-20p nema-5-30p nema-5-50p nema-6-15p nema-6-20p nema-6-30p nema-6-50p nema-10-30p nema-10-50p nema-14-20p nema-14-30p nema-14-50p nema-14-60p nema-15-15p nema-15-20p nema-15-30p nema-15-50p nema-15-60p nema-l1-15p nema-l5-15p nema-l5-20p nema-l5-30p nema-l5-50p nema-l6-15p nema-l6-20p nema-l6-30p nema-l6-50p nema-l10-30p nema-l14-20p nema-l14-30p nema-l14-50p nema-l14-60p nema-l15-20p nema-l15-30p nema-l15-50p nema-l15-60p nema-l21-20p nema-l21-30p cs6361c cs6365c cs8165c cs8265c cs8365c cs8465c ita-e ita-f ita-ef ita-g ita-h ita-i ita-j ita-k ita-l ita-m ita-n ita-o usb-a usb-b usb-c usb-mini-a usb-mini-b usb-micro-a usb-micro-b usb-3-b usb-3-micro-b]
	Type string `json:"type,omitempty"`

	// Url
	// Read Only: true
	// Format: uri
	URL strfmt.URI `json:"url,omitempty"`
}

// Validate validates this writable power port template
func (m *WritablePowerPortTemplate) Validate(formats strfmt.Registry) error {
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

func (m *WritablePowerPortTemplate) validateAllocatedDraw(formats strfmt.Registry) error {
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

func (m *WritablePowerPortTemplate) validateDescription(formats strfmt.Registry) error {
	if swag.IsZero(m.Description) { // not required
		return nil
	}

	if err := validate.MaxLength("description", "body", m.Description, 200); err != nil {
		return err
	}

	return nil
}

func (m *WritablePowerPortTemplate) validateDeviceType(formats strfmt.Registry) error {

	if err := validate.Required("device_type", "body", m.DeviceType); err != nil {
		return err
	}

	if err := validate.FormatOf("device_type", "body", "uuid", m.DeviceType.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *WritablePowerPortTemplate) validateDisplay(formats strfmt.Registry) error {
	if swag.IsZero(m.Display) { // not required
		return nil
	}

	if err := validate.MinLength("display", "body", m.Display, 1); err != nil {
		return err
	}

	return nil
}

func (m *WritablePowerPortTemplate) validateID(formats strfmt.Registry) error {
	if swag.IsZero(m.ID) { // not required
		return nil
	}

	if err := validate.FormatOf("id", "body", "uuid", m.ID.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *WritablePowerPortTemplate) validateLabel(formats strfmt.Registry) error {
	if swag.IsZero(m.Label) { // not required
		return nil
	}

	if err := validate.MaxLength("label", "body", m.Label, 64); err != nil {
		return err
	}

	return nil
}

func (m *WritablePowerPortTemplate) validateMaximumDraw(formats strfmt.Registry) error {
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

func (m *WritablePowerPortTemplate) validateName(formats strfmt.Registry) error {

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

var writablePowerPortTemplateTypeTypePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["iec-60320-c6","iec-60320-c8","iec-60320-c14","iec-60320-c16","iec-60320-c20","iec-60309-p-n-e-4h","iec-60309-p-n-e-6h","iec-60309-p-n-e-9h","iec-60309-2p-e-4h","iec-60309-2p-e-6h","iec-60309-2p-e-9h","iec-60309-3p-e-4h","iec-60309-3p-e-6h","iec-60309-3p-e-9h","iec-60309-3p-n-e-4h","iec-60309-3p-n-e-6h","iec-60309-3p-n-e-9h","nema-1-15p","nema-5-15p","nema-5-20p","nema-5-30p","nema-5-50p","nema-6-15p","nema-6-20p","nema-6-30p","nema-6-50p","nema-10-30p","nema-10-50p","nema-14-20p","nema-14-30p","nema-14-50p","nema-14-60p","nema-15-15p","nema-15-20p","nema-15-30p","nema-15-50p","nema-15-60p","nema-l1-15p","nema-l5-15p","nema-l5-20p","nema-l5-30p","nema-l5-50p","nema-l6-15p","nema-l6-20p","nema-l6-30p","nema-l6-50p","nema-l10-30p","nema-l14-20p","nema-l14-30p","nema-l14-50p","nema-l14-60p","nema-l15-20p","nema-l15-30p","nema-l15-50p","nema-l15-60p","nema-l21-20p","nema-l21-30p","cs6361c","cs6365c","cs8165c","cs8265c","cs8365c","cs8465c","ita-e","ita-f","ita-ef","ita-g","ita-h","ita-i","ita-j","ita-k","ita-l","ita-m","ita-n","ita-o","usb-a","usb-b","usb-c","usb-mini-a","usb-mini-b","usb-micro-a","usb-micro-b","usb-3-b","usb-3-micro-b"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		writablePowerPortTemplateTypeTypePropEnum = append(writablePowerPortTemplateTypeTypePropEnum, v)
	}
}

const (

	// WritablePowerPortTemplateTypeIecDash60320DashC6 captures enum value "iec-60320-c6"
	WritablePowerPortTemplateTypeIecDash60320DashC6 string = "iec-60320-c6"

	// WritablePowerPortTemplateTypeIecDash60320DashC8 captures enum value "iec-60320-c8"
	WritablePowerPortTemplateTypeIecDash60320DashC8 string = "iec-60320-c8"

	// WritablePowerPortTemplateTypeIecDash60320DashC14 captures enum value "iec-60320-c14"
	WritablePowerPortTemplateTypeIecDash60320DashC14 string = "iec-60320-c14"

	// WritablePowerPortTemplateTypeIecDash60320DashC16 captures enum value "iec-60320-c16"
	WritablePowerPortTemplateTypeIecDash60320DashC16 string = "iec-60320-c16"

	// WritablePowerPortTemplateTypeIecDash60320DashC20 captures enum value "iec-60320-c20"
	WritablePowerPortTemplateTypeIecDash60320DashC20 string = "iec-60320-c20"

	// WritablePowerPortTemplateTypeIecDash60309DashpDashnDasheDash4h captures enum value "iec-60309-p-n-e-4h"
	WritablePowerPortTemplateTypeIecDash60309DashpDashnDasheDash4h string = "iec-60309-p-n-e-4h"

	// WritablePowerPortTemplateTypeIecDash60309DashpDashnDasheDash6h captures enum value "iec-60309-p-n-e-6h"
	WritablePowerPortTemplateTypeIecDash60309DashpDashnDasheDash6h string = "iec-60309-p-n-e-6h"

	// WritablePowerPortTemplateTypeIecDash60309DashpDashnDasheDash9h captures enum value "iec-60309-p-n-e-9h"
	WritablePowerPortTemplateTypeIecDash60309DashpDashnDasheDash9h string = "iec-60309-p-n-e-9h"

	// WritablePowerPortTemplateTypeIecDash60309Dash2pDasheDash4h captures enum value "iec-60309-2p-e-4h"
	WritablePowerPortTemplateTypeIecDash60309Dash2pDasheDash4h string = "iec-60309-2p-e-4h"

	// WritablePowerPortTemplateTypeIecDash60309Dash2pDasheDash6h captures enum value "iec-60309-2p-e-6h"
	WritablePowerPortTemplateTypeIecDash60309Dash2pDasheDash6h string = "iec-60309-2p-e-6h"

	// WritablePowerPortTemplateTypeIecDash60309Dash2pDasheDash9h captures enum value "iec-60309-2p-e-9h"
	WritablePowerPortTemplateTypeIecDash60309Dash2pDasheDash9h string = "iec-60309-2p-e-9h"

	// WritablePowerPortTemplateTypeIecDash60309Dash3pDasheDash4h captures enum value "iec-60309-3p-e-4h"
	WritablePowerPortTemplateTypeIecDash60309Dash3pDasheDash4h string = "iec-60309-3p-e-4h"

	// WritablePowerPortTemplateTypeIecDash60309Dash3pDasheDash6h captures enum value "iec-60309-3p-e-6h"
	WritablePowerPortTemplateTypeIecDash60309Dash3pDasheDash6h string = "iec-60309-3p-e-6h"

	// WritablePowerPortTemplateTypeIecDash60309Dash3pDasheDash9h captures enum value "iec-60309-3p-e-9h"
	WritablePowerPortTemplateTypeIecDash60309Dash3pDasheDash9h string = "iec-60309-3p-e-9h"

	// WritablePowerPortTemplateTypeIecDash60309Dash3pDashnDasheDash4h captures enum value "iec-60309-3p-n-e-4h"
	WritablePowerPortTemplateTypeIecDash60309Dash3pDashnDasheDash4h string = "iec-60309-3p-n-e-4h"

	// WritablePowerPortTemplateTypeIecDash60309Dash3pDashnDasheDash6h captures enum value "iec-60309-3p-n-e-6h"
	WritablePowerPortTemplateTypeIecDash60309Dash3pDashnDasheDash6h string = "iec-60309-3p-n-e-6h"

	// WritablePowerPortTemplateTypeIecDash60309Dash3pDashnDasheDash9h captures enum value "iec-60309-3p-n-e-9h"
	WritablePowerPortTemplateTypeIecDash60309Dash3pDashnDasheDash9h string = "iec-60309-3p-n-e-9h"

	// WritablePowerPortTemplateTypeNemaDash1Dash15p captures enum value "nema-1-15p"
	WritablePowerPortTemplateTypeNemaDash1Dash15p string = "nema-1-15p"

	// WritablePowerPortTemplateTypeNemaDash5Dash15p captures enum value "nema-5-15p"
	WritablePowerPortTemplateTypeNemaDash5Dash15p string = "nema-5-15p"

	// WritablePowerPortTemplateTypeNemaDash5Dash20p captures enum value "nema-5-20p"
	WritablePowerPortTemplateTypeNemaDash5Dash20p string = "nema-5-20p"

	// WritablePowerPortTemplateTypeNemaDash5Dash30p captures enum value "nema-5-30p"
	WritablePowerPortTemplateTypeNemaDash5Dash30p string = "nema-5-30p"

	// WritablePowerPortTemplateTypeNemaDash5Dash50p captures enum value "nema-5-50p"
	WritablePowerPortTemplateTypeNemaDash5Dash50p string = "nema-5-50p"

	// WritablePowerPortTemplateTypeNemaDash6Dash15p captures enum value "nema-6-15p"
	WritablePowerPortTemplateTypeNemaDash6Dash15p string = "nema-6-15p"

	// WritablePowerPortTemplateTypeNemaDash6Dash20p captures enum value "nema-6-20p"
	WritablePowerPortTemplateTypeNemaDash6Dash20p string = "nema-6-20p"

	// WritablePowerPortTemplateTypeNemaDash6Dash30p captures enum value "nema-6-30p"
	WritablePowerPortTemplateTypeNemaDash6Dash30p string = "nema-6-30p"

	// WritablePowerPortTemplateTypeNemaDash6Dash50p captures enum value "nema-6-50p"
	WritablePowerPortTemplateTypeNemaDash6Dash50p string = "nema-6-50p"

	// WritablePowerPortTemplateTypeNemaDash10Dash30p captures enum value "nema-10-30p"
	WritablePowerPortTemplateTypeNemaDash10Dash30p string = "nema-10-30p"

	// WritablePowerPortTemplateTypeNemaDash10Dash50p captures enum value "nema-10-50p"
	WritablePowerPortTemplateTypeNemaDash10Dash50p string = "nema-10-50p"

	// WritablePowerPortTemplateTypeNemaDash14Dash20p captures enum value "nema-14-20p"
	WritablePowerPortTemplateTypeNemaDash14Dash20p string = "nema-14-20p"

	// WritablePowerPortTemplateTypeNemaDash14Dash30p captures enum value "nema-14-30p"
	WritablePowerPortTemplateTypeNemaDash14Dash30p string = "nema-14-30p"

	// WritablePowerPortTemplateTypeNemaDash14Dash50p captures enum value "nema-14-50p"
	WritablePowerPortTemplateTypeNemaDash14Dash50p string = "nema-14-50p"

	// WritablePowerPortTemplateTypeNemaDash14Dash60p captures enum value "nema-14-60p"
	WritablePowerPortTemplateTypeNemaDash14Dash60p string = "nema-14-60p"

	// WritablePowerPortTemplateTypeNemaDash15Dash15p captures enum value "nema-15-15p"
	WritablePowerPortTemplateTypeNemaDash15Dash15p string = "nema-15-15p"

	// WritablePowerPortTemplateTypeNemaDash15Dash20p captures enum value "nema-15-20p"
	WritablePowerPortTemplateTypeNemaDash15Dash20p string = "nema-15-20p"

	// WritablePowerPortTemplateTypeNemaDash15Dash30p captures enum value "nema-15-30p"
	WritablePowerPortTemplateTypeNemaDash15Dash30p string = "nema-15-30p"

	// WritablePowerPortTemplateTypeNemaDash15Dash50p captures enum value "nema-15-50p"
	WritablePowerPortTemplateTypeNemaDash15Dash50p string = "nema-15-50p"

	// WritablePowerPortTemplateTypeNemaDash15Dash60p captures enum value "nema-15-60p"
	WritablePowerPortTemplateTypeNemaDash15Dash60p string = "nema-15-60p"

	// WritablePowerPortTemplateTypeNemaDashL1Dash15p captures enum value "nema-l1-15p"
	WritablePowerPortTemplateTypeNemaDashL1Dash15p string = "nema-l1-15p"

	// WritablePowerPortTemplateTypeNemaDashL5Dash15p captures enum value "nema-l5-15p"
	WritablePowerPortTemplateTypeNemaDashL5Dash15p string = "nema-l5-15p"

	// WritablePowerPortTemplateTypeNemaDashL5Dash20p captures enum value "nema-l5-20p"
	WritablePowerPortTemplateTypeNemaDashL5Dash20p string = "nema-l5-20p"

	// WritablePowerPortTemplateTypeNemaDashL5Dash30p captures enum value "nema-l5-30p"
	WritablePowerPortTemplateTypeNemaDashL5Dash30p string = "nema-l5-30p"

	// WritablePowerPortTemplateTypeNemaDashL5Dash50p captures enum value "nema-l5-50p"
	WritablePowerPortTemplateTypeNemaDashL5Dash50p string = "nema-l5-50p"

	// WritablePowerPortTemplateTypeNemaDashL6Dash15p captures enum value "nema-l6-15p"
	WritablePowerPortTemplateTypeNemaDashL6Dash15p string = "nema-l6-15p"

	// WritablePowerPortTemplateTypeNemaDashL6Dash20p captures enum value "nema-l6-20p"
	WritablePowerPortTemplateTypeNemaDashL6Dash20p string = "nema-l6-20p"

	// WritablePowerPortTemplateTypeNemaDashL6Dash30p captures enum value "nema-l6-30p"
	WritablePowerPortTemplateTypeNemaDashL6Dash30p string = "nema-l6-30p"

	// WritablePowerPortTemplateTypeNemaDashL6Dash50p captures enum value "nema-l6-50p"
	WritablePowerPortTemplateTypeNemaDashL6Dash50p string = "nema-l6-50p"

	// WritablePowerPortTemplateTypeNemaDashL10Dash30p captures enum value "nema-l10-30p"
	WritablePowerPortTemplateTypeNemaDashL10Dash30p string = "nema-l10-30p"

	// WritablePowerPortTemplateTypeNemaDashL14Dash20p captures enum value "nema-l14-20p"
	WritablePowerPortTemplateTypeNemaDashL14Dash20p string = "nema-l14-20p"

	// WritablePowerPortTemplateTypeNemaDashL14Dash30p captures enum value "nema-l14-30p"
	WritablePowerPortTemplateTypeNemaDashL14Dash30p string = "nema-l14-30p"

	// WritablePowerPortTemplateTypeNemaDashL14Dash50p captures enum value "nema-l14-50p"
	WritablePowerPortTemplateTypeNemaDashL14Dash50p string = "nema-l14-50p"

	// WritablePowerPortTemplateTypeNemaDashL14Dash60p captures enum value "nema-l14-60p"
	WritablePowerPortTemplateTypeNemaDashL14Dash60p string = "nema-l14-60p"

	// WritablePowerPortTemplateTypeNemaDashL15Dash20p captures enum value "nema-l15-20p"
	WritablePowerPortTemplateTypeNemaDashL15Dash20p string = "nema-l15-20p"

	// WritablePowerPortTemplateTypeNemaDashL15Dash30p captures enum value "nema-l15-30p"
	WritablePowerPortTemplateTypeNemaDashL15Dash30p string = "nema-l15-30p"

	// WritablePowerPortTemplateTypeNemaDashL15Dash50p captures enum value "nema-l15-50p"
	WritablePowerPortTemplateTypeNemaDashL15Dash50p string = "nema-l15-50p"

	// WritablePowerPortTemplateTypeNemaDashL15Dash60p captures enum value "nema-l15-60p"
	WritablePowerPortTemplateTypeNemaDashL15Dash60p string = "nema-l15-60p"

	// WritablePowerPortTemplateTypeNemaDashL21Dash20p captures enum value "nema-l21-20p"
	WritablePowerPortTemplateTypeNemaDashL21Dash20p string = "nema-l21-20p"

	// WritablePowerPortTemplateTypeNemaDashL21Dash30p captures enum value "nema-l21-30p"
	WritablePowerPortTemplateTypeNemaDashL21Dash30p string = "nema-l21-30p"

	// WritablePowerPortTemplateTypeCs6361c captures enum value "cs6361c"
	WritablePowerPortTemplateTypeCs6361c string = "cs6361c"

	// WritablePowerPortTemplateTypeCs6365c captures enum value "cs6365c"
	WritablePowerPortTemplateTypeCs6365c string = "cs6365c"

	// WritablePowerPortTemplateTypeCs8165c captures enum value "cs8165c"
	WritablePowerPortTemplateTypeCs8165c string = "cs8165c"

	// WritablePowerPortTemplateTypeCs8265c captures enum value "cs8265c"
	WritablePowerPortTemplateTypeCs8265c string = "cs8265c"

	// WritablePowerPortTemplateTypeCs8365c captures enum value "cs8365c"
	WritablePowerPortTemplateTypeCs8365c string = "cs8365c"

	// WritablePowerPortTemplateTypeCs8465c captures enum value "cs8465c"
	WritablePowerPortTemplateTypeCs8465c string = "cs8465c"

	// WritablePowerPortTemplateTypeItaDashe captures enum value "ita-e"
	WritablePowerPortTemplateTypeItaDashe string = "ita-e"

	// WritablePowerPortTemplateTypeItaDashf captures enum value "ita-f"
	WritablePowerPortTemplateTypeItaDashf string = "ita-f"

	// WritablePowerPortTemplateTypeItaDashEf captures enum value "ita-ef"
	WritablePowerPortTemplateTypeItaDashEf string = "ita-ef"

	// WritablePowerPortTemplateTypeItaDashg captures enum value "ita-g"
	WritablePowerPortTemplateTypeItaDashg string = "ita-g"

	// WritablePowerPortTemplateTypeItaDashh captures enum value "ita-h"
	WritablePowerPortTemplateTypeItaDashh string = "ita-h"

	// WritablePowerPortTemplateTypeItaDashi captures enum value "ita-i"
	WritablePowerPortTemplateTypeItaDashi string = "ita-i"

	// WritablePowerPortTemplateTypeItaDashj captures enum value "ita-j"
	WritablePowerPortTemplateTypeItaDashj string = "ita-j"

	// WritablePowerPortTemplateTypeItaDashk captures enum value "ita-k"
	WritablePowerPortTemplateTypeItaDashk string = "ita-k"

	// WritablePowerPortTemplateTypeItaDashl captures enum value "ita-l"
	WritablePowerPortTemplateTypeItaDashl string = "ita-l"

	// WritablePowerPortTemplateTypeItaDashm captures enum value "ita-m"
	WritablePowerPortTemplateTypeItaDashm string = "ita-m"

	// WritablePowerPortTemplateTypeItaDashn captures enum value "ita-n"
	WritablePowerPortTemplateTypeItaDashn string = "ita-n"

	// WritablePowerPortTemplateTypeItaDasho captures enum value "ita-o"
	WritablePowerPortTemplateTypeItaDasho string = "ita-o"

	// WritablePowerPortTemplateTypeUsbDasha captures enum value "usb-a"
	WritablePowerPortTemplateTypeUsbDasha string = "usb-a"

	// WritablePowerPortTemplateTypeUsbDashb captures enum value "usb-b"
	WritablePowerPortTemplateTypeUsbDashb string = "usb-b"

	// WritablePowerPortTemplateTypeUsbDashc captures enum value "usb-c"
	WritablePowerPortTemplateTypeUsbDashc string = "usb-c"

	// WritablePowerPortTemplateTypeUsbDashMiniDasha captures enum value "usb-mini-a"
	WritablePowerPortTemplateTypeUsbDashMiniDasha string = "usb-mini-a"

	// WritablePowerPortTemplateTypeUsbDashMiniDashb captures enum value "usb-mini-b"
	WritablePowerPortTemplateTypeUsbDashMiniDashb string = "usb-mini-b"

	// WritablePowerPortTemplateTypeUsbDashMicroDasha captures enum value "usb-micro-a"
	WritablePowerPortTemplateTypeUsbDashMicroDasha string = "usb-micro-a"

	// WritablePowerPortTemplateTypeUsbDashMicroDashb captures enum value "usb-micro-b"
	WritablePowerPortTemplateTypeUsbDashMicroDashb string = "usb-micro-b"

	// WritablePowerPortTemplateTypeUsbDash3Dashb captures enum value "usb-3-b"
	WritablePowerPortTemplateTypeUsbDash3Dashb string = "usb-3-b"

	// WritablePowerPortTemplateTypeUsbDash3DashMicroDashb captures enum value "usb-3-micro-b"
	WritablePowerPortTemplateTypeUsbDash3DashMicroDashb string = "usb-3-micro-b"
)

// prop value enum
func (m *WritablePowerPortTemplate) validateTypeEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, writablePowerPortTemplateTypeTypePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *WritablePowerPortTemplate) validateType(formats strfmt.Registry) error {
	if swag.IsZero(m.Type) { // not required
		return nil
	}

	// value enum
	if err := m.validateTypeEnum("type", "body", m.Type); err != nil {
		return err
	}

	return nil
}

func (m *WritablePowerPortTemplate) validateURL(formats strfmt.Registry) error {
	if swag.IsZero(m.URL) { // not required
		return nil
	}

	if err := validate.FormatOf("url", "body", "uri", m.URL.String(), formats); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this writable power port template based on the context it is used
func (m *WritablePowerPortTemplate) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateComputedFields(ctx, formats); err != nil {
		res = append(res, err)
	}

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

func (m *WritablePowerPortTemplate) contextValidateComputedFields(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "computed_fields", "body", string(m.ComputedFields)); err != nil {
		return err
	}

	return nil
}

func (m *WritablePowerPortTemplate) contextValidateDisplay(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "display", "body", string(m.Display)); err != nil {
		return err
	}

	return nil
}

func (m *WritablePowerPortTemplate) contextValidateID(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "id", "body", strfmt.UUID(m.ID)); err != nil {
		return err
	}

	return nil
}

func (m *WritablePowerPortTemplate) contextValidateURL(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "url", "body", strfmt.URI(m.URL)); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *WritablePowerPortTemplate) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *WritablePowerPortTemplate) UnmarshalBinary(b []byte) error {
	var res WritablePowerPortTemplate
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
