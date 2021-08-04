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

// VMInterface VM interface
//
// swagger:model VMInterface
type VMInterface struct {

	// Description
	// Max Length: 200
	Description string `json:"description,omitempty"`

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

	// MAC Address
	// Max Length: 18
	MacAddress *string `json:"mac_address,omitempty"`

	// mode
	Mode *VMInterfaceMode `json:"mode,omitempty"`

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

	// untagged vlan
	UntaggedVlan *NestedVLAN `json:"untagged_vlan,omitempty"`

	// Url
	// Read Only: true
	// Format: uri
	URL strfmt.URI `json:"url,omitempty"`

	// virtual machine
	// Required: true
	VirtualMachine *NestedVirtualMachine `json:"virtual_machine"`
}

// Validate validates this VM interface
func (m *VMInterface) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateDescription(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDisplay(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateID(formats); err != nil {
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

	if err := m.validateUntaggedVlan(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateURL(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateVirtualMachine(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *VMInterface) validateDescription(formats strfmt.Registry) error {
	if swag.IsZero(m.Description) { // not required
		return nil
	}

	if err := validate.MaxLength("description", "body", m.Description, 200); err != nil {
		return err
	}

	return nil
}

func (m *VMInterface) validateDisplay(formats strfmt.Registry) error {
	if swag.IsZero(m.Display) { // not required
		return nil
	}

	if err := validate.MinLength("display", "body", m.Display, 1); err != nil {
		return err
	}

	return nil
}

func (m *VMInterface) validateID(formats strfmt.Registry) error {
	if swag.IsZero(m.ID) { // not required
		return nil
	}

	if err := validate.FormatOf("id", "body", "uuid", m.ID.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *VMInterface) validateMacAddress(formats strfmt.Registry) error {
	if swag.IsZero(m.MacAddress) { // not required
		return nil
	}

	if err := validate.MaxLength("mac_address", "body", *m.MacAddress, 18); err != nil {
		return err
	}

	return nil
}

func (m *VMInterface) validateMode(formats strfmt.Registry) error {
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

func (m *VMInterface) validateMtu(formats strfmt.Registry) error {
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

func (m *VMInterface) validateName(formats strfmt.Registry) error {

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

func (m *VMInterface) validateTaggedVlans(formats strfmt.Registry) error {
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

func (m *VMInterface) validateTags(formats strfmt.Registry) error {
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

func (m *VMInterface) validateUntaggedVlan(formats strfmt.Registry) error {
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

func (m *VMInterface) validateURL(formats strfmt.Registry) error {
	if swag.IsZero(m.URL) { // not required
		return nil
	}

	if err := validate.FormatOf("url", "body", "uri", m.URL.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *VMInterface) validateVirtualMachine(formats strfmt.Registry) error {

	if err := validate.Required("virtual_machine", "body", m.VirtualMachine); err != nil {
		return err
	}

	if m.VirtualMachine != nil {
		if err := m.VirtualMachine.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("virtual_machine")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this VM interface based on the context it is used
func (m *VMInterface) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateDisplay(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateID(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateMode(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateTags(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateUntaggedVlan(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateURL(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateVirtualMachine(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *VMInterface) contextValidateDisplay(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "display", "body", string(m.Display)); err != nil {
		return err
	}

	return nil
}

func (m *VMInterface) contextValidateID(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "id", "body", strfmt.UUID(m.ID)); err != nil {
		return err
	}

	return nil
}

func (m *VMInterface) contextValidateMode(ctx context.Context, formats strfmt.Registry) error {

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

func (m *VMInterface) contextValidateTags(ctx context.Context, formats strfmt.Registry) error {

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

func (m *VMInterface) contextValidateUntaggedVlan(ctx context.Context, formats strfmt.Registry) error {

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

func (m *VMInterface) contextValidateURL(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "url", "body", strfmt.URI(m.URL)); err != nil {
		return err
	}

	return nil
}

func (m *VMInterface) contextValidateVirtualMachine(ctx context.Context, formats strfmt.Registry) error {

	if m.VirtualMachine != nil {
		if err := m.VirtualMachine.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("virtual_machine")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *VMInterface) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *VMInterface) UnmarshalBinary(b []byte) error {
	var res VMInterface
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// VMInterfaceMode Mode
//
// swagger:model VMInterfaceMode
type VMInterfaceMode struct {

	// label
	// Required: true
	// Enum: [Access Tagged Tagged (All)]
	Label *string `json:"label"`

	// value
	// Required: true
	// Enum: [access tagged tagged-all]
	Value *string `json:"value"`
}

// Validate validates this VM interface mode
func (m *VMInterfaceMode) Validate(formats strfmt.Registry) error {
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

var vmInterfaceModeTypeLabelPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["Access","Tagged","Tagged (All)"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		vmInterfaceModeTypeLabelPropEnum = append(vmInterfaceModeTypeLabelPropEnum, v)
	}
}

const (

	// VMInterfaceModeLabelAccess captures enum value "Access"
	VMInterfaceModeLabelAccess string = "Access"

	// VMInterfaceModeLabelTagged captures enum value "Tagged"
	VMInterfaceModeLabelTagged string = "Tagged"

	// VMInterfaceModeLabelTaggedAll captures enum value "Tagged (All)"
	VMInterfaceModeLabelTaggedAll string = "Tagged (All)"
)

// prop value enum
func (m *VMInterfaceMode) validateLabelEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, vmInterfaceModeTypeLabelPropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *VMInterfaceMode) validateLabel(formats strfmt.Registry) error {

	if err := validate.Required("mode"+"."+"label", "body", m.Label); err != nil {
		return err
	}

	// value enum
	if err := m.validateLabelEnum("mode"+"."+"label", "body", *m.Label); err != nil {
		return err
	}

	return nil
}

var vmInterfaceModeTypeValuePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["access","tagged","tagged-all"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		vmInterfaceModeTypeValuePropEnum = append(vmInterfaceModeTypeValuePropEnum, v)
	}
}

const (

	// VMInterfaceModeValueAccess captures enum value "access"
	VMInterfaceModeValueAccess string = "access"

	// VMInterfaceModeValueTagged captures enum value "tagged"
	VMInterfaceModeValueTagged string = "tagged"

	// VMInterfaceModeValueTaggedDashAll captures enum value "tagged-all"
	VMInterfaceModeValueTaggedDashAll string = "tagged-all"
)

// prop value enum
func (m *VMInterfaceMode) validateValueEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, vmInterfaceModeTypeValuePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *VMInterfaceMode) validateValue(formats strfmt.Registry) error {

	if err := validate.Required("mode"+"."+"value", "body", m.Value); err != nil {
		return err
	}

	// value enum
	if err := m.validateValueEnum("mode"+"."+"value", "body", *m.Value); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this VM interface mode based on context it is used
func (m *VMInterfaceMode) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *VMInterfaceMode) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *VMInterfaceMode) UnmarshalBinary(b []byte) error {
	var res VMInterfaceMode
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
