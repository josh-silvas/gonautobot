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

// Rack rack
//
// swagger:model Rack
type Rack struct {

	// Asset tag
	//
	// A unique tag used to identify this rack
	// Max Length: 50
	AssetTag *string `json:"asset_tag,omitempty"`

	// Comments
	Comments string `json:"comments,omitempty"`

	// Created
	// Read Only: true
	// Format: date
	Created strfmt.Date `json:"created,omitempty"`

	// Custom fields
	CustomFields interface{} `json:"custom_fields,omitempty"`

	// Descending units
	//
	// Units are numbered top-to-bottom
	DescUnits bool `json:"desc_units,omitempty"`

	// Device count
	// Read Only: true
	DeviceCount int64 `json:"device_count,omitempty"`

	// Display
	//
	// Human friendly display value
	// Read Only: true
	// Min Length: 1
	Display string `json:"display,omitempty"`

	// Facility ID
	//
	// Locally-assigned identifier
	// Max Length: 50
	FacilityID *string `json:"facility_id,omitempty"`

	// group
	Group *NestedRackGroup `json:"group,omitempty"`

	// Id
	// Read Only: true
	// Format: uuid
	ID strfmt.UUID `json:"id,omitempty"`

	// Last updated
	// Read Only: true
	// Format: date-time
	LastUpdated strfmt.DateTime `json:"last_updated,omitempty"`

	// Name
	// Required: true
	// Max Length: 100
	// Min Length: 1
	Name *string `json:"name"`

	// Outer depth
	//
	// Outer dimension of rack (depth)
	// Maximum: 32767
	// Minimum: 0
	OuterDepth *int64 `json:"outer_depth,omitempty"`

	// outer unit
	OuterUnit *RackOuterUnit `json:"outer_unit,omitempty"`

	// Outer width
	//
	// Outer dimension of rack (width)
	// Maximum: 32767
	// Minimum: 0
	OuterWidth *int64 `json:"outer_width,omitempty"`

	// Powerfeed count
	// Read Only: true
	PowerfeedCount int64 `json:"powerfeed_count,omitempty"`

	// role
	Role *NestedRackRole `json:"role,omitempty"`

	// Serial number
	// Max Length: 50
	Serial string `json:"serial,omitempty"`

	// site
	// Required: true
	Site *NestedSite `json:"site"`

	// Status
	// Pattern: ^[-a-zA-Z0-9_]+$
	// Enum: [active available deprecated planned reserved]
	Status string `json:"status,omitempty"`

	// tags
	Tags []*NestedTag `json:"tags"`

	// tenant
	Tenant *NestedTenant `json:"tenant,omitempty"`

	// type
	Type *RackType `json:"type,omitempty"`

	// Height (U)
	//
	// Height in rack units
	// Maximum: 100
	// Minimum: 1
	UHeight int64 `json:"u_height,omitempty"`

	// Url
	// Read Only: true
	// Format: uri
	URL strfmt.URI `json:"url,omitempty"`

	// width
	Width *RackWidth `json:"width,omitempty"`
}

// Validate validates this rack
func (m *Rack) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAssetTag(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCreated(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDisplay(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateFacilityID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateGroup(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateLastUpdated(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateName(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateOuterDepth(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateOuterUnit(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateOuterWidth(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateRole(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSerial(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSite(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateStatus(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTags(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTenant(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateType(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateUHeight(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateURL(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateWidth(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Rack) validateAssetTag(formats strfmt.Registry) error {
	if swag.IsZero(m.AssetTag) { // not required
		return nil
	}

	if err := validate.MaxLength("asset_tag", "body", *m.AssetTag, 50); err != nil {
		return err
	}

	return nil
}

func (m *Rack) validateCreated(formats strfmt.Registry) error {
	if swag.IsZero(m.Created) { // not required
		return nil
	}

	if err := validate.FormatOf("created", "body", "date", m.Created.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *Rack) validateDisplay(formats strfmt.Registry) error {
	if swag.IsZero(m.Display) { // not required
		return nil
	}

	if err := validate.MinLength("display", "body", m.Display, 1); err != nil {
		return err
	}

	return nil
}

func (m *Rack) validateFacilityID(formats strfmt.Registry) error {
	if swag.IsZero(m.FacilityID) { // not required
		return nil
	}

	if err := validate.MaxLength("facility_id", "body", *m.FacilityID, 50); err != nil {
		return err
	}

	return nil
}

func (m *Rack) validateGroup(formats strfmt.Registry) error {
	if swag.IsZero(m.Group) { // not required
		return nil
	}

	if m.Group != nil {
		if err := m.Group.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("group")
			}
			return err
		}
	}

	return nil
}

func (m *Rack) validateID(formats strfmt.Registry) error {
	if swag.IsZero(m.ID) { // not required
		return nil
	}

	if err := validate.FormatOf("id", "body", "uuid", m.ID.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *Rack) validateLastUpdated(formats strfmt.Registry) error {
	if swag.IsZero(m.LastUpdated) { // not required
		return nil
	}

	if err := validate.FormatOf("last_updated", "body", "date-time", m.LastUpdated.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *Rack) validateName(formats strfmt.Registry) error {

	if err := validate.Required("name", "body", m.Name); err != nil {
		return err
	}

	if err := validate.MinLength("name", "body", *m.Name, 1); err != nil {
		return err
	}

	if err := validate.MaxLength("name", "body", *m.Name, 100); err != nil {
		return err
	}

	return nil
}

func (m *Rack) validateOuterDepth(formats strfmt.Registry) error {
	if swag.IsZero(m.OuterDepth) { // not required
		return nil
	}

	if err := validate.MinimumInt("outer_depth", "body", *m.OuterDepth, 0, false); err != nil {
		return err
	}

	if err := validate.MaximumInt("outer_depth", "body", *m.OuterDepth, 32767, false); err != nil {
		return err
	}

	return nil
}

func (m *Rack) validateOuterUnit(formats strfmt.Registry) error {
	if swag.IsZero(m.OuterUnit) { // not required
		return nil
	}

	if m.OuterUnit != nil {
		if err := m.OuterUnit.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("outer_unit")
			}
			return err
		}
	}

	return nil
}

func (m *Rack) validateOuterWidth(formats strfmt.Registry) error {
	if swag.IsZero(m.OuterWidth) { // not required
		return nil
	}

	if err := validate.MinimumInt("outer_width", "body", *m.OuterWidth, 0, false); err != nil {
		return err
	}

	if err := validate.MaximumInt("outer_width", "body", *m.OuterWidth, 32767, false); err != nil {
		return err
	}

	return nil
}

func (m *Rack) validateRole(formats strfmt.Registry) error {
	if swag.IsZero(m.Role) { // not required
		return nil
	}

	if m.Role != nil {
		if err := m.Role.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("role")
			}
			return err
		}
	}

	return nil
}

func (m *Rack) validateSerial(formats strfmt.Registry) error {
	if swag.IsZero(m.Serial) { // not required
		return nil
	}

	if err := validate.MaxLength("serial", "body", m.Serial, 50); err != nil {
		return err
	}

	return nil
}

func (m *Rack) validateSite(formats strfmt.Registry) error {

	if err := validate.Required("site", "body", m.Site); err != nil {
		return err
	}

	if m.Site != nil {
		if err := m.Site.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("site")
			}
			return err
		}
	}

	return nil
}

var rackTypeStatusPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["active","available","deprecated","planned","reserved"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		rackTypeStatusPropEnum = append(rackTypeStatusPropEnum, v)
	}
}

const (

	// RackStatusActive captures enum value "active"
	RackStatusActive string = "active"

	// RackStatusAvailable captures enum value "available"
	RackStatusAvailable string = "available"

	// RackStatusDeprecated captures enum value "deprecated"
	RackStatusDeprecated string = "deprecated"

	// RackStatusPlanned captures enum value "planned"
	RackStatusPlanned string = "planned"

	// RackStatusReserved captures enum value "reserved"
	RackStatusReserved string = "reserved"
)

// prop value enum
func (m *Rack) validateStatusEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, rackTypeStatusPropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *Rack) validateStatus(formats strfmt.Registry) error {
	if swag.IsZero(m.Status) { // not required
		return nil
	}

	if err := validate.Pattern("status", "body", m.Status, `^[-a-zA-Z0-9_]+$`); err != nil {
		return err
	}

	// value enum
	if err := m.validateStatusEnum("status", "body", m.Status); err != nil {
		return err
	}

	return nil
}

func (m *Rack) validateTags(formats strfmt.Registry) error {
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

func (m *Rack) validateTenant(formats strfmt.Registry) error {
	if swag.IsZero(m.Tenant) { // not required
		return nil
	}

	if m.Tenant != nil {
		if err := m.Tenant.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("tenant")
			}
			return err
		}
	}

	return nil
}

func (m *Rack) validateType(formats strfmt.Registry) error {
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

func (m *Rack) validateUHeight(formats strfmt.Registry) error {
	if swag.IsZero(m.UHeight) { // not required
		return nil
	}

	if err := validate.MinimumInt("u_height", "body", m.UHeight, 1, false); err != nil {
		return err
	}

	if err := validate.MaximumInt("u_height", "body", m.UHeight, 100, false); err != nil {
		return err
	}

	return nil
}

func (m *Rack) validateURL(formats strfmt.Registry) error {
	if swag.IsZero(m.URL) { // not required
		return nil
	}

	if err := validate.FormatOf("url", "body", "uri", m.URL.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *Rack) validateWidth(formats strfmt.Registry) error {
	if swag.IsZero(m.Width) { // not required
		return nil
	}

	if m.Width != nil {
		if err := m.Width.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("width")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this rack based on the context it is used
func (m *Rack) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateCreated(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateDeviceCount(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateDisplay(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateGroup(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateID(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateLastUpdated(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateOuterUnit(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidatePowerfeedCount(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateRole(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateSite(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateTags(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateTenant(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateType(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateURL(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateWidth(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Rack) contextValidateCreated(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "created", "body", strfmt.Date(m.Created)); err != nil {
		return err
	}

	return nil
}

func (m *Rack) contextValidateDeviceCount(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "device_count", "body", int64(m.DeviceCount)); err != nil {
		return err
	}

	return nil
}

func (m *Rack) contextValidateDisplay(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "display", "body", string(m.Display)); err != nil {
		return err
	}

	return nil
}

func (m *Rack) contextValidateGroup(ctx context.Context, formats strfmt.Registry) error {

	if m.Group != nil {
		if err := m.Group.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("group")
			}
			return err
		}
	}

	return nil
}

func (m *Rack) contextValidateID(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "id", "body", strfmt.UUID(m.ID)); err != nil {
		return err
	}

	return nil
}

func (m *Rack) contextValidateLastUpdated(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "last_updated", "body", strfmt.DateTime(m.LastUpdated)); err != nil {
		return err
	}

	return nil
}

func (m *Rack) contextValidateOuterUnit(ctx context.Context, formats strfmt.Registry) error {

	if m.OuterUnit != nil {
		if err := m.OuterUnit.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("outer_unit")
			}
			return err
		}
	}

	return nil
}

func (m *Rack) contextValidatePowerfeedCount(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "powerfeed_count", "body", int64(m.PowerfeedCount)); err != nil {
		return err
	}

	return nil
}

func (m *Rack) contextValidateRole(ctx context.Context, formats strfmt.Registry) error {

	if m.Role != nil {
		if err := m.Role.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("role")
			}
			return err
		}
	}

	return nil
}

func (m *Rack) contextValidateSite(ctx context.Context, formats strfmt.Registry) error {

	if m.Site != nil {
		if err := m.Site.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("site")
			}
			return err
		}
	}

	return nil
}

func (m *Rack) contextValidateTags(ctx context.Context, formats strfmt.Registry) error {

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

func (m *Rack) contextValidateTenant(ctx context.Context, formats strfmt.Registry) error {

	if m.Tenant != nil {
		if err := m.Tenant.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("tenant")
			}
			return err
		}
	}

	return nil
}

func (m *Rack) contextValidateType(ctx context.Context, formats strfmt.Registry) error {

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

func (m *Rack) contextValidateURL(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "url", "body", strfmt.URI(m.URL)); err != nil {
		return err
	}

	return nil
}

func (m *Rack) contextValidateWidth(ctx context.Context, formats strfmt.Registry) error {

	if m.Width != nil {
		if err := m.Width.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("width")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *Rack) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Rack) UnmarshalBinary(b []byte) error {
	var res Rack
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// RackOuterUnit Outer unit
//
// swagger:model RackOuterUnit
type RackOuterUnit struct {

	// label
	// Required: true
	// Enum: [Millimeters Inches]
	Label *string `json:"label"`

	// value
	// Required: true
	// Enum: [mm in]
	Value *string `json:"value"`
}

// Validate validates this rack outer unit
func (m *RackOuterUnit) Validate(formats strfmt.Registry) error {
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

var rackOuterUnitTypeLabelPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["Millimeters","Inches"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		rackOuterUnitTypeLabelPropEnum = append(rackOuterUnitTypeLabelPropEnum, v)
	}
}

const (

	// RackOuterUnitLabelMillimeters captures enum value "Millimeters"
	RackOuterUnitLabelMillimeters string = "Millimeters"

	// RackOuterUnitLabelInches captures enum value "Inches"
	RackOuterUnitLabelInches string = "Inches"
)

// prop value enum
func (m *RackOuterUnit) validateLabelEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, rackOuterUnitTypeLabelPropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *RackOuterUnit) validateLabel(formats strfmt.Registry) error {

	if err := validate.Required("outer_unit"+"."+"label", "body", m.Label); err != nil {
		return err
	}

	// value enum
	if err := m.validateLabelEnum("outer_unit"+"."+"label", "body", *m.Label); err != nil {
		return err
	}

	return nil
}

var rackOuterUnitTypeValuePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["mm","in"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		rackOuterUnitTypeValuePropEnum = append(rackOuterUnitTypeValuePropEnum, v)
	}
}

const (

	// RackOuterUnitValueMm captures enum value "mm"
	RackOuterUnitValueMm string = "mm"

	// RackOuterUnitValueIn captures enum value "in"
	RackOuterUnitValueIn string = "in"
)

// prop value enum
func (m *RackOuterUnit) validateValueEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, rackOuterUnitTypeValuePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *RackOuterUnit) validateValue(formats strfmt.Registry) error {

	if err := validate.Required("outer_unit"+"."+"value", "body", m.Value); err != nil {
		return err
	}

	// value enum
	if err := m.validateValueEnum("outer_unit"+"."+"value", "body", *m.Value); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this rack outer unit based on context it is used
func (m *RackOuterUnit) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *RackOuterUnit) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *RackOuterUnit) UnmarshalBinary(b []byte) error {
	var res RackOuterUnit
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// RackType Type
//
// swagger:model RackType
type RackType struct {

	// label
	// Required: true
	// Enum: [2-post frame 4-post frame 4-post cabinet Wall-mounted frame Wall-mounted cabinet]
	Label *string `json:"label"`

	// value
	// Required: true
	// Enum: [2-post-frame 4-post-frame 4-post-cabinet wall-frame wall-cabinet]
	Value *string `json:"value"`
}

// Validate validates this rack type
func (m *RackType) Validate(formats strfmt.Registry) error {
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

var rackTypeTypeLabelPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["2-post frame","4-post frame","4-post cabinet","Wall-mounted frame","Wall-mounted cabinet"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		rackTypeTypeLabelPropEnum = append(rackTypeTypeLabelPropEnum, v)
	}
}

const (

	// RackTypeLabelNr2DashPostFrame captures enum value "2-post frame"
	RackTypeLabelNr2DashPostFrame string = "2-post frame"

	// RackTypeLabelNr4DashPostFrame captures enum value "4-post frame"
	RackTypeLabelNr4DashPostFrame string = "4-post frame"

	// RackTypeLabelNr4DashPostCabinet captures enum value "4-post cabinet"
	RackTypeLabelNr4DashPostCabinet string = "4-post cabinet"

	// RackTypeLabelWallDashMountedFrame captures enum value "Wall-mounted frame"
	RackTypeLabelWallDashMountedFrame string = "Wall-mounted frame"

	// RackTypeLabelWallDashMountedCabinet captures enum value "Wall-mounted cabinet"
	RackTypeLabelWallDashMountedCabinet string = "Wall-mounted cabinet"
)

// prop value enum
func (m *RackType) validateLabelEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, rackTypeTypeLabelPropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *RackType) validateLabel(formats strfmt.Registry) error {

	if err := validate.Required("type"+"."+"label", "body", m.Label); err != nil {
		return err
	}

	// value enum
	if err := m.validateLabelEnum("type"+"."+"label", "body", *m.Label); err != nil {
		return err
	}

	return nil
}

var rackTypeTypeValuePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["2-post-frame","4-post-frame","4-post-cabinet","wall-frame","wall-cabinet"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		rackTypeTypeValuePropEnum = append(rackTypeTypeValuePropEnum, v)
	}
}

const (

	// RackTypeValueNr2DashPostDashFrame captures enum value "2-post-frame"
	RackTypeValueNr2DashPostDashFrame string = "2-post-frame"

	// RackTypeValueNr4DashPostDashFrame captures enum value "4-post-frame"
	RackTypeValueNr4DashPostDashFrame string = "4-post-frame"

	// RackTypeValueNr4DashPostDashCabinet captures enum value "4-post-cabinet"
	RackTypeValueNr4DashPostDashCabinet string = "4-post-cabinet"

	// RackTypeValueWallDashFrame captures enum value "wall-frame"
	RackTypeValueWallDashFrame string = "wall-frame"

	// RackTypeValueWallDashCabinet captures enum value "wall-cabinet"
	RackTypeValueWallDashCabinet string = "wall-cabinet"
)

// prop value enum
func (m *RackType) validateValueEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, rackTypeTypeValuePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *RackType) validateValue(formats strfmt.Registry) error {

	if err := validate.Required("type"+"."+"value", "body", m.Value); err != nil {
		return err
	}

	// value enum
	if err := m.validateValueEnum("type"+"."+"value", "body", *m.Value); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this rack type based on context it is used
func (m *RackType) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *RackType) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *RackType) UnmarshalBinary(b []byte) error {
	var res RackType
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// RackWidth Width
//
// swagger:model RackWidth
type RackWidth struct {

	// label
	// Required: true
	// Enum: [10 inches 19 inches 21 inches 23 inches]
	Label *string `json:"label"`

	// value
	// Required: true
	// Enum: [10 19 21 23]
	Value *int64 `json:"value"`
}

// Validate validates this rack width
func (m *RackWidth) Validate(formats strfmt.Registry) error {
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

var rackWidthTypeLabelPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["10 inches","19 inches","21 inches","23 inches"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		rackWidthTypeLabelPropEnum = append(rackWidthTypeLabelPropEnum, v)
	}
}

const (

	// RackWidthLabelNr10Inches captures enum value "10 inches"
	RackWidthLabelNr10Inches string = "10 inches"

	// RackWidthLabelNr19Inches captures enum value "19 inches"
	RackWidthLabelNr19Inches string = "19 inches"

	// RackWidthLabelNr21Inches captures enum value "21 inches"
	RackWidthLabelNr21Inches string = "21 inches"

	// RackWidthLabelNr23Inches captures enum value "23 inches"
	RackWidthLabelNr23Inches string = "23 inches"
)

// prop value enum
func (m *RackWidth) validateLabelEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, rackWidthTypeLabelPropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *RackWidth) validateLabel(formats strfmt.Registry) error {

	if err := validate.Required("width"+"."+"label", "body", m.Label); err != nil {
		return err
	}

	// value enum
	if err := m.validateLabelEnum("width"+"."+"label", "body", *m.Label); err != nil {
		return err
	}

	return nil
}

var rackWidthTypeValuePropEnum []interface{}

func init() {
	var res []int64
	if err := json.Unmarshal([]byte(`[10,19,21,23]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		rackWidthTypeValuePropEnum = append(rackWidthTypeValuePropEnum, v)
	}
}

// prop value enum
func (m *RackWidth) validateValueEnum(path, location string, value int64) error {
	if err := validate.EnumCase(path, location, value, rackWidthTypeValuePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *RackWidth) validateValue(formats strfmt.Registry) error {

	if err := validate.Required("width"+"."+"value", "body", m.Value); err != nil {
		return err
	}

	// value enum
	if err := m.validateValueEnum("width"+"."+"value", "body", *m.Value); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this rack width based on context it is used
func (m *RackWidth) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *RackWidth) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *RackWidth) UnmarshalBinary(b []byte) error {
	var res RackWidth
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
