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

// Site site
//
// swagger:model Site
type Site struct {

	// ASN
	//
	// 32-bit autonomous system number
	// Maximum: 4.294967295e+09
	// Minimum: 1
	Asn *int64 `json:"asn,omitempty"`

	// Circuit count
	// Read Only: true
	CircuitCount int64 `json:"circuit_count,omitempty"`

	// Comments
	Comments string `json:"comments,omitempty"`

	// Contact E-mail
	// Max Length: 254
	// Format: email
	ContactEmail strfmt.Email `json:"contact_email,omitempty"`

	// Contact name
	// Max Length: 50
	ContactName string `json:"contact_name,omitempty"`

	// Contact phone
	// Max Length: 20
	ContactPhone string `json:"contact_phone,omitempty"`

	// Created
	// Read Only: true
	// Format: date
	Created strfmt.Date `json:"created,omitempty"`

	// Custom fields
	CustomFields interface{} `json:"custom_fields,omitempty"`

	// Description
	// Max Length: 200
	Description string `json:"description,omitempty"`

	// Device count
	// Read Only: true
	DeviceCount int64 `json:"device_count,omitempty"`

	// Display
	//
	// Human friendly display value
	// Read Only: true
	// Min Length: 1
	Display string `json:"display,omitempty"`

	// Facility
	//
	// Local facility ID or description
	// Max Length: 50
	Facility string `json:"facility,omitempty"`

	// Id
	// Read Only: true
	// Format: uuid
	ID strfmt.UUID `json:"id,omitempty"`

	// Last updated
	// Read Only: true
	// Format: date-time
	LastUpdated strfmt.DateTime `json:"last_updated,omitempty"`

	// Latitude
	//
	// GPS coordinate (latitude)
	Latitude *string `json:"latitude,omitempty"`

	// Longitude
	//
	// GPS coordinate (longitude)
	Longitude *string `json:"longitude,omitempty"`

	// Name
	// Required: true
	// Max Length: 100
	// Min Length: 1
	Name *string `json:"name"`

	// Physical address
	// Max Length: 200
	PhysicalAddress string `json:"physical_address,omitempty"`

	// Prefix count
	// Read Only: true
	PrefixCount int64 `json:"prefix_count,omitempty"`

	// Rack count
	// Read Only: true
	RackCount int64 `json:"rack_count,omitempty"`

	// region
	Region *NestedRegion `json:"region,omitempty"`

	// Shipping address
	// Max Length: 200
	ShippingAddress string `json:"shipping_address,omitempty"`

	// Slug
	// Required: true
	// Max Length: 100
	// Min Length: 1
	// Pattern: ^[-a-zA-Z0-9_]+$
	Slug *string `json:"slug"`

	// Status
	// Pattern: ^[-a-zA-Z0-9_]+$
	// Enum: [active decommissioning planned retired staging]
	Status string `json:"status,omitempty"`

	// tags
	Tags []*NestedTag `json:"tags"`

	// tenant
	Tenant *NestedTenant `json:"tenant,omitempty"`

	// Time zone
	TimeZone *string `json:"time_zone,omitempty"`

	// Url
	// Read Only: true
	// Format: uri
	URL strfmt.URI `json:"url,omitempty"`

	// Virtualmachine count
	// Read Only: true
	VirtualmachineCount int64 `json:"virtualmachine_count,omitempty"`

	// Vlan count
	// Read Only: true
	VlanCount int64 `json:"vlan_count,omitempty"`
}

// Validate validates this site
func (m *Site) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAsn(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateContactEmail(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateContactName(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateContactPhone(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCreated(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDescription(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDisplay(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateFacility(formats); err != nil {
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

	if err := m.validatePhysicalAddress(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateRegion(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateShippingAddress(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSlug(formats); err != nil {
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

	if err := m.validateURL(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Site) validateAsn(formats strfmt.Registry) error {
	if swag.IsZero(m.Asn) { // not required
		return nil
	}

	if err := validate.MinimumInt("asn", "body", *m.Asn, 1, false); err != nil {
		return err
	}

	if err := validate.MaximumInt("asn", "body", *m.Asn, 4.294967295e+09, false); err != nil {
		return err
	}

	return nil
}

func (m *Site) validateContactEmail(formats strfmt.Registry) error {
	if swag.IsZero(m.ContactEmail) { // not required
		return nil
	}

	if err := validate.MaxLength("contact_email", "body", m.ContactEmail.String(), 254); err != nil {
		return err
	}

	if err := validate.FormatOf("contact_email", "body", "email", m.ContactEmail.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *Site) validateContactName(formats strfmt.Registry) error {
	if swag.IsZero(m.ContactName) { // not required
		return nil
	}

	if err := validate.MaxLength("contact_name", "body", m.ContactName, 50); err != nil {
		return err
	}

	return nil
}

func (m *Site) validateContactPhone(formats strfmt.Registry) error {
	if swag.IsZero(m.ContactPhone) { // not required
		return nil
	}

	if err := validate.MaxLength("contact_phone", "body", m.ContactPhone, 20); err != nil {
		return err
	}

	return nil
}

func (m *Site) validateCreated(formats strfmt.Registry) error {
	if swag.IsZero(m.Created) { // not required
		return nil
	}

	if err := validate.FormatOf("created", "body", "date", m.Created.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *Site) validateDescription(formats strfmt.Registry) error {
	if swag.IsZero(m.Description) { // not required
		return nil
	}

	if err := validate.MaxLength("description", "body", m.Description, 200); err != nil {
		return err
	}

	return nil
}

func (m *Site) validateDisplay(formats strfmt.Registry) error {
	if swag.IsZero(m.Display) { // not required
		return nil
	}

	if err := validate.MinLength("display", "body", m.Display, 1); err != nil {
		return err
	}

	return nil
}

func (m *Site) validateFacility(formats strfmt.Registry) error {
	if swag.IsZero(m.Facility) { // not required
		return nil
	}

	if err := validate.MaxLength("facility", "body", m.Facility, 50); err != nil {
		return err
	}

	return nil
}

func (m *Site) validateID(formats strfmt.Registry) error {
	if swag.IsZero(m.ID) { // not required
		return nil
	}

	if err := validate.FormatOf("id", "body", "uuid", m.ID.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *Site) validateLastUpdated(formats strfmt.Registry) error {
	if swag.IsZero(m.LastUpdated) { // not required
		return nil
	}

	if err := validate.FormatOf("last_updated", "body", "date-time", m.LastUpdated.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *Site) validateName(formats strfmt.Registry) error {

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

func (m *Site) validatePhysicalAddress(formats strfmt.Registry) error {
	if swag.IsZero(m.PhysicalAddress) { // not required
		return nil
	}

	if err := validate.MaxLength("physical_address", "body", m.PhysicalAddress, 200); err != nil {
		return err
	}

	return nil
}

func (m *Site) validateRegion(formats strfmt.Registry) error {
	if swag.IsZero(m.Region) { // not required
		return nil
	}

	if m.Region != nil {
		if err := m.Region.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("region")
			}
			return err
		}
	}

	return nil
}

func (m *Site) validateShippingAddress(formats strfmt.Registry) error {
	if swag.IsZero(m.ShippingAddress) { // not required
		return nil
	}

	if err := validate.MaxLength("shipping_address", "body", m.ShippingAddress, 200); err != nil {
		return err
	}

	return nil
}

func (m *Site) validateSlug(formats strfmt.Registry) error {

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

var siteTypeStatusPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["active","decommissioning","planned","retired","staging"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		siteTypeStatusPropEnum = append(siteTypeStatusPropEnum, v)
	}
}

const (

	// SiteStatusActive captures enum value "active"
	SiteStatusActive string = "active"

	// SiteStatusDecommissioning captures enum value "decommissioning"
	SiteStatusDecommissioning string = "decommissioning"

	// SiteStatusPlanned captures enum value "planned"
	SiteStatusPlanned string = "planned"

	// SiteStatusRetired captures enum value "retired"
	SiteStatusRetired string = "retired"

	// SiteStatusStaging captures enum value "staging"
	SiteStatusStaging string = "staging"
)

// prop value enum
func (m *Site) validateStatusEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, siteTypeStatusPropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *Site) validateStatus(formats strfmt.Registry) error {
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

func (m *Site) validateTags(formats strfmt.Registry) error {
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

func (m *Site) validateTenant(formats strfmt.Registry) error {
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

func (m *Site) validateURL(formats strfmt.Registry) error {
	if swag.IsZero(m.URL) { // not required
		return nil
	}

	if err := validate.FormatOf("url", "body", "uri", m.URL.String(), formats); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this site based on the context it is used
func (m *Site) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateCircuitCount(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateCreated(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateDeviceCount(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateDisplay(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateID(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateLastUpdated(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidatePrefixCount(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateRackCount(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateRegion(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateTags(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateTenant(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateURL(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateVirtualmachineCount(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateVlanCount(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Site) contextValidateCircuitCount(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "circuit_count", "body", int64(m.CircuitCount)); err != nil {
		return err
	}

	return nil
}

func (m *Site) contextValidateCreated(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "created", "body", strfmt.Date(m.Created)); err != nil {
		return err
	}

	return nil
}

func (m *Site) contextValidateDeviceCount(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "device_count", "body", int64(m.DeviceCount)); err != nil {
		return err
	}

	return nil
}

func (m *Site) contextValidateDisplay(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "display", "body", string(m.Display)); err != nil {
		return err
	}

	return nil
}

func (m *Site) contextValidateID(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "id", "body", strfmt.UUID(m.ID)); err != nil {
		return err
	}

	return nil
}

func (m *Site) contextValidateLastUpdated(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "last_updated", "body", strfmt.DateTime(m.LastUpdated)); err != nil {
		return err
	}

	return nil
}

func (m *Site) contextValidatePrefixCount(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "prefix_count", "body", int64(m.PrefixCount)); err != nil {
		return err
	}

	return nil
}

func (m *Site) contextValidateRackCount(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "rack_count", "body", int64(m.RackCount)); err != nil {
		return err
	}

	return nil
}

func (m *Site) contextValidateRegion(ctx context.Context, formats strfmt.Registry) error {

	if m.Region != nil {
		if err := m.Region.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("region")
			}
			return err
		}
	}

	return nil
}

func (m *Site) contextValidateTags(ctx context.Context, formats strfmt.Registry) error {

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

func (m *Site) contextValidateTenant(ctx context.Context, formats strfmt.Registry) error {

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

func (m *Site) contextValidateURL(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "url", "body", strfmt.URI(m.URL)); err != nil {
		return err
	}

	return nil
}

func (m *Site) contextValidateVirtualmachineCount(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "virtualmachine_count", "body", int64(m.VirtualmachineCount)); err != nil {
		return err
	}

	return nil
}

func (m *Site) contextValidateVlanCount(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "vlan_count", "body", int64(m.VlanCount)); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *Site) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Site) UnmarshalBinary(b []byte) error {
	var res Site
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
