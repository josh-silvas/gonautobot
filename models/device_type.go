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

// DeviceType device type
//
// swagger:model DeviceType
type DeviceType struct {

	// Comments
	Comments string `json:"comments,omitempty"`

	// Created
	// Read Only: true
	// Format: date
	Created strfmt.Date `json:"created,omitempty"`

	// Custom fields
	CustomFields interface{} `json:"custom_fields,omitempty"`

	// Device count
	// Read Only: true
	DeviceCount int64 `json:"device_count,omitempty"`

	// Display
	//
	// Human friendly display value
	// Read Only: true
	// Min Length: 1
	Display string `json:"display,omitempty"`

	// Front image
	// Read Only: true
	// Format: uri
	FrontImage strfmt.URI `json:"front_image,omitempty"`

	// Id
	// Read Only: true
	// Format: uuid
	ID strfmt.UUID `json:"id,omitempty"`

	// Is full depth
	//
	// Device consumes both front and rear rack faces
	IsFullDepth bool `json:"is_full_depth,omitempty"`

	// Last updated
	// Read Only: true
	// Format: date-time
	LastUpdated strfmt.DateTime `json:"last_updated,omitempty"`

	// manufacturer
	// Required: true
	Manufacturer *NestedManufacturer `json:"manufacturer"`

	// Model
	// Required: true
	// Max Length: 100
	// Min Length: 1
	Model *string `json:"model"`

	// Part number
	//
	// Discrete part number (optional)
	// Max Length: 50
	PartNumber string `json:"part_number,omitempty"`

	// Rear image
	// Read Only: true
	// Format: uri
	RearImage strfmt.URI `json:"rear_image,omitempty"`

	// Slug
	// Required: true
	// Max Length: 100
	// Min Length: 1
	// Pattern: ^[-a-zA-Z0-9_]+$
	Slug *string `json:"slug"`

	// subdevice role
	SubdeviceRole *DeviceTypeSubdeviceRole `json:"subdevice_role,omitempty"`

	// tags
	Tags []*NestedTag `json:"tags"`

	// Height (U)
	// Maximum: 32767
	// Minimum: 0
	UHeight *int64 `json:"u_height,omitempty"`

	// Url
	// Read Only: true
	// Format: uri
	URL strfmt.URI `json:"url,omitempty"`
}

// Validate validates this device type
func (m *DeviceType) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCreated(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDisplay(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateFrontImage(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateLastUpdated(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateManufacturer(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateModel(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePartNumber(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateRearImage(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSlug(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSubdeviceRole(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTags(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateUHeight(formats); err != nil {
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

func (m *DeviceType) validateCreated(formats strfmt.Registry) error {
	if swag.IsZero(m.Created) { // not required
		return nil
	}

	if err := validate.FormatOf("created", "body", "date", m.Created.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *DeviceType) validateDisplay(formats strfmt.Registry) error {
	if swag.IsZero(m.Display) { // not required
		return nil
	}

	if err := validate.MinLength("display", "body", m.Display, 1); err != nil {
		return err
	}

	return nil
}

func (m *DeviceType) validateFrontImage(formats strfmt.Registry) error {
	if swag.IsZero(m.FrontImage) { // not required
		return nil
	}

	if err := validate.FormatOf("front_image", "body", "uri", m.FrontImage.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *DeviceType) validateID(formats strfmt.Registry) error {
	if swag.IsZero(m.ID) { // not required
		return nil
	}

	if err := validate.FormatOf("id", "body", "uuid", m.ID.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *DeviceType) validateLastUpdated(formats strfmt.Registry) error {
	if swag.IsZero(m.LastUpdated) { // not required
		return nil
	}

	if err := validate.FormatOf("last_updated", "body", "date-time", m.LastUpdated.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *DeviceType) validateManufacturer(formats strfmt.Registry) error {

	if err := validate.Required("manufacturer", "body", m.Manufacturer); err != nil {
		return err
	}

	if m.Manufacturer != nil {
		if err := m.Manufacturer.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("manufacturer")
			}
			return err
		}
	}

	return nil
}

func (m *DeviceType) validateModel(formats strfmt.Registry) error {

	if err := validate.Required("model", "body", m.Model); err != nil {
		return err
	}

	if err := validate.MinLength("model", "body", *m.Model, 1); err != nil {
		return err
	}

	if err := validate.MaxLength("model", "body", *m.Model, 100); err != nil {
		return err
	}

	return nil
}

func (m *DeviceType) validatePartNumber(formats strfmt.Registry) error {
	if swag.IsZero(m.PartNumber) { // not required
		return nil
	}

	if err := validate.MaxLength("part_number", "body", m.PartNumber, 50); err != nil {
		return err
	}

	return nil
}

func (m *DeviceType) validateRearImage(formats strfmt.Registry) error {
	if swag.IsZero(m.RearImage) { // not required
		return nil
	}

	if err := validate.FormatOf("rear_image", "body", "uri", m.RearImage.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *DeviceType) validateSlug(formats strfmt.Registry) error {

	if err := validate.Required("slug", "body", m.Slug); err != nil {
		return err
	}

	if err := validate.MinLength("slug", "body", *m.Slug, 1); err != nil {
		return err
	}

	if err := validate.MaxLength("slug", "body", *m.Slug, 100); err != nil {
		return err
	}

	if err := validate.Pattern("slug", "body", *m.Slug, `^[-a-zA-Z0-9_]+$`); err != nil {
		return err
	}

	return nil
}

func (m *DeviceType) validateSubdeviceRole(formats strfmt.Registry) error {
	if swag.IsZero(m.SubdeviceRole) { // not required
		return nil
	}

	if m.SubdeviceRole != nil {
		if err := m.SubdeviceRole.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("subdevice_role")
			}
			return err
		}
	}

	return nil
}

func (m *DeviceType) validateTags(formats strfmt.Registry) error {
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

func (m *DeviceType) validateUHeight(formats strfmt.Registry) error {
	if swag.IsZero(m.UHeight) { // not required
		return nil
	}

	if err := validate.MinimumInt("u_height", "body", *m.UHeight, 0, false); err != nil {
		return err
	}

	if err := validate.MaximumInt("u_height", "body", *m.UHeight, 32767, false); err != nil {
		return err
	}

	return nil
}

func (m *DeviceType) validateURL(formats strfmt.Registry) error {
	if swag.IsZero(m.URL) { // not required
		return nil
	}

	if err := validate.FormatOf("url", "body", "uri", m.URL.String(), formats); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this device type based on the context it is used
func (m *DeviceType) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
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

	if err := m.contextValidateFrontImage(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateID(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateLastUpdated(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateManufacturer(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateRearImage(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateSubdeviceRole(ctx, formats); err != nil {
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

func (m *DeviceType) contextValidateCreated(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "created", "body", strfmt.Date(m.Created)); err != nil {
		return err
	}

	return nil
}

func (m *DeviceType) contextValidateDeviceCount(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "device_count", "body", int64(m.DeviceCount)); err != nil {
		return err
	}

	return nil
}

func (m *DeviceType) contextValidateDisplay(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "display", "body", string(m.Display)); err != nil {
		return err
	}

	return nil
}

func (m *DeviceType) contextValidateFrontImage(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "front_image", "body", strfmt.URI(m.FrontImage)); err != nil {
		return err
	}

	return nil
}

func (m *DeviceType) contextValidateID(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "id", "body", strfmt.UUID(m.ID)); err != nil {
		return err
	}

	return nil
}

func (m *DeviceType) contextValidateLastUpdated(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "last_updated", "body", strfmt.DateTime(m.LastUpdated)); err != nil {
		return err
	}

	return nil
}

func (m *DeviceType) contextValidateManufacturer(ctx context.Context, formats strfmt.Registry) error {

	if m.Manufacturer != nil {
		if err := m.Manufacturer.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("manufacturer")
			}
			return err
		}
	}

	return nil
}

func (m *DeviceType) contextValidateRearImage(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "rear_image", "body", strfmt.URI(m.RearImage)); err != nil {
		return err
	}

	return nil
}

func (m *DeviceType) contextValidateSubdeviceRole(ctx context.Context, formats strfmt.Registry) error {

	if m.SubdeviceRole != nil {
		if err := m.SubdeviceRole.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("subdevice_role")
			}
			return err
		}
	}

	return nil
}

func (m *DeviceType) contextValidateTags(ctx context.Context, formats strfmt.Registry) error {

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

func (m *DeviceType) contextValidateURL(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "url", "body", strfmt.URI(m.URL)); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *DeviceType) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *DeviceType) UnmarshalBinary(b []byte) error {
	var res DeviceType
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// DeviceTypeSubdeviceRole Subdevice role
//
// swagger:model DeviceTypeSubdeviceRole
type DeviceTypeSubdeviceRole struct {

	// label
	// Required: true
	// Enum: [Parent Child]
	Label *string `json:"label"`

	// value
	// Required: true
	// Enum: [parent child]
	Value *string `json:"value"`
}

// Validate validates this device type subdevice role
func (m *DeviceTypeSubdeviceRole) Validate(formats strfmt.Registry) error {
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

var deviceTypeSubdeviceRoleTypeLabelPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["Parent","Child"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		deviceTypeSubdeviceRoleTypeLabelPropEnum = append(deviceTypeSubdeviceRoleTypeLabelPropEnum, v)
	}
}

const (

	// DeviceTypeSubdeviceRoleLabelParent captures enum value "Parent"
	DeviceTypeSubdeviceRoleLabelParent string = "Parent"

	// DeviceTypeSubdeviceRoleLabelChild captures enum value "Child"
	DeviceTypeSubdeviceRoleLabelChild string = "Child"
)

// prop value enum
func (m *DeviceTypeSubdeviceRole) validateLabelEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, deviceTypeSubdeviceRoleTypeLabelPropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *DeviceTypeSubdeviceRole) validateLabel(formats strfmt.Registry) error {

	if err := validate.Required("subdevice_role"+"."+"label", "body", m.Label); err != nil {
		return err
	}

	// value enum
	if err := m.validateLabelEnum("subdevice_role"+"."+"label", "body", *m.Label); err != nil {
		return err
	}

	return nil
}

var deviceTypeSubdeviceRoleTypeValuePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["parent","child"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		deviceTypeSubdeviceRoleTypeValuePropEnum = append(deviceTypeSubdeviceRoleTypeValuePropEnum, v)
	}
}

const (

	// DeviceTypeSubdeviceRoleValueParent captures enum value "parent"
	DeviceTypeSubdeviceRoleValueParent string = "parent"

	// DeviceTypeSubdeviceRoleValueChild captures enum value "child"
	DeviceTypeSubdeviceRoleValueChild string = "child"
)

// prop value enum
func (m *DeviceTypeSubdeviceRole) validateValueEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, deviceTypeSubdeviceRoleTypeValuePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *DeviceTypeSubdeviceRole) validateValue(formats strfmt.Registry) error {

	if err := validate.Required("subdevice_role"+"."+"value", "body", m.Value); err != nil {
		return err
	}

	// value enum
	if err := m.validateValueEnum("subdevice_role"+"."+"value", "body", *m.Value); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this device type subdevice role based on context it is used
func (m *DeviceTypeSubdeviceRole) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *DeviceTypeSubdeviceRole) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *DeviceTypeSubdeviceRole) UnmarshalBinary(b []byte) error {
	var res DeviceTypeSubdeviceRole
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
