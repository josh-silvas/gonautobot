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

// WritableDeviceWithConfigContext writable device with config context
//
// swagger:model WritableDeviceWithConfigContext
type WritableDeviceWithConfigContext struct {

	// Asset tag
	//
	// A unique tag used to identify this device
	// Max Length: 50
	AssetTag *string `json:"asset_tag,omitempty"`

	// Cluster
	// Format: uuid
	Cluster *strfmt.UUID `json:"cluster,omitempty"`

	// Comments
	Comments string `json:"comments,omitempty"`

	// Computed fields
	// Read Only: true
	ComputedFields string `json:"computed_fields,omitempty"`

	// Config context
	// Read Only: true
	ConfigContext map[string]*string `json:"config_context,omitempty"`

	// Created
	// Read Only: true
	// Format: date
	Created strfmt.Date `json:"created,omitempty"`

	// Custom fields
	CustomFields interface{} `json:"custom_fields,omitempty"`

	// Device role
	// Required: true
	// Format: uuid
	DeviceRole *strfmt.UUID `json:"device_role"`

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

	// Rack face
	// Enum: [front rear]
	Face string `json:"face,omitempty"`

	// Id
	// Read Only: true
	// Format: uuid
	ID strfmt.UUID `json:"id,omitempty"`

	// Last updated
	// Read Only: true
	// Format: date-time
	LastUpdated strfmt.DateTime `json:"last_updated,omitempty"`

	// Local context data
	LocalContextData *string `json:"local_context_data,omitempty"`

	// Local context schema
	//
	// Optional schema to validate the structure of the data
	// Format: uuid
	LocalContextSchema *strfmt.UUID `json:"local_context_schema,omitempty"`

	// Name
	// Max Length: 64
	Name *string `json:"name,omitempty"`

	// parent device
	ParentDevice *NestedDevice `json:"parent_device,omitempty"`

	// Platform
	// Format: uuid
	Platform *strfmt.UUID `json:"platform,omitempty"`

	// Position (U)
	//
	// The lowest-numbered unit occupied by the device
	// Maximum: 32767
	// Minimum: 1
	Position *int64 `json:"position,omitempty"`

	// Primary ip
	// Read Only: true
	PrimaryIP string `json:"primary_ip,omitempty"`

	// Primary IPv4
	// Format: uuid
	PrimaryIp4 *strfmt.UUID `json:"primary_ip4,omitempty"`

	// Primary IPv6
	// Format: uuid
	PrimaryIp6 *strfmt.UUID `json:"primary_ip6,omitempty"`

	// Rack
	// Format: uuid
	Rack *strfmt.UUID `json:"rack,omitempty"`

	// Serial number
	// Max Length: 50
	Serial string `json:"serial,omitempty"`

	// Site
	// Required: true
	// Format: uuid
	Site *strfmt.UUID `json:"site"`

	// Status
	// Pattern: ^[-a-zA-Z0-9_]+$
	// Enum: [active decommissioning failed inventory offline planned staged]
	Status string `json:"status,omitempty"`

	// tags
	Tags []*NestedTag `json:"tags"`

	// Tenant
	// Format: uuid
	Tenant *strfmt.UUID `json:"tenant,omitempty"`

	// Url
	// Read Only: true
	// Format: uri
	URL strfmt.URI `json:"url,omitempty"`

	// Vc position
	// Maximum: 255
	// Minimum: 0
	VcPosition *int64 `json:"vc_position,omitempty"`

	// Vc priority
	// Maximum: 255
	// Minimum: 0
	VcPriority *int64 `json:"vc_priority,omitempty"`

	// Virtual chassis
	// Format: uuid
	VirtualChassis *strfmt.UUID `json:"virtual_chassis,omitempty"`
}

// Validate validates this writable device with config context
func (m *WritableDeviceWithConfigContext) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAssetTag(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCluster(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCreated(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDeviceRole(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDeviceType(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDisplay(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateFace(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateLastUpdated(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateLocalContextSchema(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateName(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateParentDevice(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePlatform(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePosition(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePrimaryIp4(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePrimaryIp6(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateRack(formats); err != nil {
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

	if err := m.validateURL(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateVcPosition(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateVcPriority(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateVirtualChassis(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *WritableDeviceWithConfigContext) validateAssetTag(formats strfmt.Registry) error {
	if swag.IsZero(m.AssetTag) { // not required
		return nil
	}

	if err := validate.MaxLength("asset_tag", "body", *m.AssetTag, 50); err != nil {
		return err
	}

	return nil
}

func (m *WritableDeviceWithConfigContext) validateCluster(formats strfmt.Registry) error {
	if swag.IsZero(m.Cluster) { // not required
		return nil
	}

	if err := validate.FormatOf("cluster", "body", "uuid", m.Cluster.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *WritableDeviceWithConfigContext) validateCreated(formats strfmt.Registry) error {
	if swag.IsZero(m.Created) { // not required
		return nil
	}

	if err := validate.FormatOf("created", "body", "date", m.Created.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *WritableDeviceWithConfigContext) validateDeviceRole(formats strfmt.Registry) error {

	if err := validate.Required("device_role", "body", m.DeviceRole); err != nil {
		return err
	}

	if err := validate.FormatOf("device_role", "body", "uuid", m.DeviceRole.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *WritableDeviceWithConfigContext) validateDeviceType(formats strfmt.Registry) error {

	if err := validate.Required("device_type", "body", m.DeviceType); err != nil {
		return err
	}

	if err := validate.FormatOf("device_type", "body", "uuid", m.DeviceType.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *WritableDeviceWithConfigContext) validateDisplay(formats strfmt.Registry) error {
	if swag.IsZero(m.Display) { // not required
		return nil
	}

	if err := validate.MinLength("display", "body", m.Display, 1); err != nil {
		return err
	}

	return nil
}

var writableDeviceWithConfigContextTypeFacePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["front","rear"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		writableDeviceWithConfigContextTypeFacePropEnum = append(writableDeviceWithConfigContextTypeFacePropEnum, v)
	}
}

const (

	// WritableDeviceWithConfigContextFaceFront captures enum value "front"
	WritableDeviceWithConfigContextFaceFront string = "front"

	// WritableDeviceWithConfigContextFaceRear captures enum value "rear"
	WritableDeviceWithConfigContextFaceRear string = "rear"
)

// prop value enum
func (m *WritableDeviceWithConfigContext) validateFaceEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, writableDeviceWithConfigContextTypeFacePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *WritableDeviceWithConfigContext) validateFace(formats strfmt.Registry) error {
	if swag.IsZero(m.Face) { // not required
		return nil
	}

	// value enum
	if err := m.validateFaceEnum("face", "body", m.Face); err != nil {
		return err
	}

	return nil
}

func (m *WritableDeviceWithConfigContext) validateID(formats strfmt.Registry) error {
	if swag.IsZero(m.ID) { // not required
		return nil
	}

	if err := validate.FormatOf("id", "body", "uuid", m.ID.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *WritableDeviceWithConfigContext) validateLastUpdated(formats strfmt.Registry) error {
	if swag.IsZero(m.LastUpdated) { // not required
		return nil
	}

	if err := validate.FormatOf("last_updated", "body", "date-time", m.LastUpdated.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *WritableDeviceWithConfigContext) validateLocalContextSchema(formats strfmt.Registry) error {
	if swag.IsZero(m.LocalContextSchema) { // not required
		return nil
	}

	if err := validate.FormatOf("local_context_schema", "body", "uuid", m.LocalContextSchema.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *WritableDeviceWithConfigContext) validateName(formats strfmt.Registry) error {
	if swag.IsZero(m.Name) { // not required
		return nil
	}

	if err := validate.MaxLength("name", "body", *m.Name, 64); err != nil {
		return err
	}

	return nil
}

func (m *WritableDeviceWithConfigContext) validateParentDevice(formats strfmt.Registry) error {
	if swag.IsZero(m.ParentDevice) { // not required
		return nil
	}

	if m.ParentDevice != nil {
		if err := m.ParentDevice.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("parent_device")
			}
			return err
		}
	}

	return nil
}

func (m *WritableDeviceWithConfigContext) validatePlatform(formats strfmt.Registry) error {
	if swag.IsZero(m.Platform) { // not required
		return nil
	}

	if err := validate.FormatOf("platform", "body", "uuid", m.Platform.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *WritableDeviceWithConfigContext) validatePosition(formats strfmt.Registry) error {
	if swag.IsZero(m.Position) { // not required
		return nil
	}

	if err := validate.MinimumInt("position", "body", *m.Position, 1, false); err != nil {
		return err
	}

	if err := validate.MaximumInt("position", "body", *m.Position, 32767, false); err != nil {
		return err
	}

	return nil
}

func (m *WritableDeviceWithConfigContext) validatePrimaryIp4(formats strfmt.Registry) error {
	if swag.IsZero(m.PrimaryIp4) { // not required
		return nil
	}

	if err := validate.FormatOf("primary_ip4", "body", "uuid", m.PrimaryIp4.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *WritableDeviceWithConfigContext) validatePrimaryIp6(formats strfmt.Registry) error {
	if swag.IsZero(m.PrimaryIp6) { // not required
		return nil
	}

	if err := validate.FormatOf("primary_ip6", "body", "uuid", m.PrimaryIp6.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *WritableDeviceWithConfigContext) validateRack(formats strfmt.Registry) error {
	if swag.IsZero(m.Rack) { // not required
		return nil
	}

	if err := validate.FormatOf("rack", "body", "uuid", m.Rack.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *WritableDeviceWithConfigContext) validateSerial(formats strfmt.Registry) error {
	if swag.IsZero(m.Serial) { // not required
		return nil
	}

	if err := validate.MaxLength("serial", "body", m.Serial, 50); err != nil {
		return err
	}

	return nil
}

func (m *WritableDeviceWithConfigContext) validateSite(formats strfmt.Registry) error {

	if err := validate.Required("site", "body", m.Site); err != nil {
		return err
	}

	if err := validate.FormatOf("site", "body", "uuid", m.Site.String(), formats); err != nil {
		return err
	}

	return nil
}

var writableDeviceWithConfigContextTypeStatusPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["active","decommissioning","failed","inventory","offline","planned","staged"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		writableDeviceWithConfigContextTypeStatusPropEnum = append(writableDeviceWithConfigContextTypeStatusPropEnum, v)
	}
}

const (

	// WritableDeviceWithConfigContextStatusActive captures enum value "active"
	WritableDeviceWithConfigContextStatusActive string = "active"

	// WritableDeviceWithConfigContextStatusDecommissioning captures enum value "decommissioning"
	WritableDeviceWithConfigContextStatusDecommissioning string = "decommissioning"

	// WritableDeviceWithConfigContextStatusFailed captures enum value "failed"
	WritableDeviceWithConfigContextStatusFailed string = "failed"

	// WritableDeviceWithConfigContextStatusInventory captures enum value "inventory"
	WritableDeviceWithConfigContextStatusInventory string = "inventory"

	// WritableDeviceWithConfigContextStatusOffline captures enum value "offline"
	WritableDeviceWithConfigContextStatusOffline string = "offline"

	// WritableDeviceWithConfigContextStatusPlanned captures enum value "planned"
	WritableDeviceWithConfigContextStatusPlanned string = "planned"

	// WritableDeviceWithConfigContextStatusStaged captures enum value "staged"
	WritableDeviceWithConfigContextStatusStaged string = "staged"
)

// prop value enum
func (m *WritableDeviceWithConfigContext) validateStatusEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, writableDeviceWithConfigContextTypeStatusPropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *WritableDeviceWithConfigContext) validateStatus(formats strfmt.Registry) error {
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

func (m *WritableDeviceWithConfigContext) validateTags(formats strfmt.Registry) error {
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

func (m *WritableDeviceWithConfigContext) validateTenant(formats strfmt.Registry) error {
	if swag.IsZero(m.Tenant) { // not required
		return nil
	}

	if err := validate.FormatOf("tenant", "body", "uuid", m.Tenant.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *WritableDeviceWithConfigContext) validateURL(formats strfmt.Registry) error {
	if swag.IsZero(m.URL) { // not required
		return nil
	}

	if err := validate.FormatOf("url", "body", "uri", m.URL.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *WritableDeviceWithConfigContext) validateVcPosition(formats strfmt.Registry) error {
	if swag.IsZero(m.VcPosition) { // not required
		return nil
	}

	if err := validate.MinimumInt("vc_position", "body", *m.VcPosition, 0, false); err != nil {
		return err
	}

	if err := validate.MaximumInt("vc_position", "body", *m.VcPosition, 255, false); err != nil {
		return err
	}

	return nil
}

func (m *WritableDeviceWithConfigContext) validateVcPriority(formats strfmt.Registry) error {
	if swag.IsZero(m.VcPriority) { // not required
		return nil
	}

	if err := validate.MinimumInt("vc_priority", "body", *m.VcPriority, 0, false); err != nil {
		return err
	}

	if err := validate.MaximumInt("vc_priority", "body", *m.VcPriority, 255, false); err != nil {
		return err
	}

	return nil
}

func (m *WritableDeviceWithConfigContext) validateVirtualChassis(formats strfmt.Registry) error {
	if swag.IsZero(m.VirtualChassis) { // not required
		return nil
	}

	if err := validate.FormatOf("virtual_chassis", "body", "uuid", m.VirtualChassis.String(), formats); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this writable device with config context based on the context it is used
func (m *WritableDeviceWithConfigContext) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateComputedFields(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateConfigContext(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateCreated(ctx, formats); err != nil {
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

	if err := m.contextValidateParentDevice(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidatePrimaryIP(ctx, formats); err != nil {
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

func (m *WritableDeviceWithConfigContext) contextValidateComputedFields(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "computed_fields", "body", string(m.ComputedFields)); err != nil {
		return err
	}

	return nil
}

func (m *WritableDeviceWithConfigContext) contextValidateConfigContext(ctx context.Context, formats strfmt.Registry) error {

	return nil
}

func (m *WritableDeviceWithConfigContext) contextValidateCreated(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "created", "body", strfmt.Date(m.Created)); err != nil {
		return err
	}

	return nil
}

func (m *WritableDeviceWithConfigContext) contextValidateDisplay(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "display", "body", string(m.Display)); err != nil {
		return err
	}

	return nil
}

func (m *WritableDeviceWithConfigContext) contextValidateID(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "id", "body", strfmt.UUID(m.ID)); err != nil {
		return err
	}

	return nil
}

func (m *WritableDeviceWithConfigContext) contextValidateLastUpdated(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "last_updated", "body", strfmt.DateTime(m.LastUpdated)); err != nil {
		return err
	}

	return nil
}

func (m *WritableDeviceWithConfigContext) contextValidateParentDevice(ctx context.Context, formats strfmt.Registry) error {

	if m.ParentDevice != nil {
		if err := m.ParentDevice.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("parent_device")
			}
			return err
		}
	}

	return nil
}

func (m *WritableDeviceWithConfigContext) contextValidatePrimaryIP(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "primary_ip", "body", string(m.PrimaryIP)); err != nil {
		return err
	}

	return nil
}

func (m *WritableDeviceWithConfigContext) contextValidateTags(ctx context.Context, formats strfmt.Registry) error {

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

func (m *WritableDeviceWithConfigContext) contextValidateURL(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "url", "body", strfmt.URI(m.URL)); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *WritableDeviceWithConfigContext) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *WritableDeviceWithConfigContext) UnmarshalBinary(b []byte) error {
	var res WritableDeviceWithConfigContext
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
