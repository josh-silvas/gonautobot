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

// PowerFeed power feed
//
// swagger:model PowerFeed
type PowerFeed struct {

	// Amperage
	// Maximum: 32767
	// Minimum: 1
	Amperage int64 `json:"amperage,omitempty"`

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

	// Comments
	Comments string `json:"comments,omitempty"`

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

	// Created
	// Read Only: true
	// Format: date
	Created strfmt.Date `json:"created,omitempty"`

	// Custom fields
	CustomFields interface{} `json:"custom_fields,omitempty"`

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

	// Last updated
	// Read Only: true
	// Format: date-time
	LastUpdated strfmt.DateTime `json:"last_updated,omitempty"`

	// Max utilization
	//
	// Maximum permissible draw (percentage)
	// Maximum: 100
	// Minimum: 1
	MaxUtilization int64 `json:"max_utilization,omitempty"`

	// Name
	// Required: true
	// Max Length: 100
	// Min Length: 1
	Name *string `json:"name"`

	// phase
	Phase *PowerFeedPhase `json:"phase,omitempty"`

	// power panel
	// Required: true
	PowerPanel *NestedPowerPanel `json:"power_panel"`

	// rack
	Rack *NestedRack `json:"rack,omitempty"`

	// Status
	// Pattern: ^[-a-zA-Z0-9_]+$
	// Enum: [active failed offline planned]
	Status string `json:"status,omitempty"`

	// supply
	Supply *PowerFeedSupply `json:"supply,omitempty"`

	// tags
	Tags []*NestedTag `json:"tags"`

	// type
	Type *PowerFeedType `json:"type,omitempty"`

	// Url
	// Read Only: true
	// Format: uri
	URL strfmt.URI `json:"url,omitempty"`

	// Voltage
	// Maximum: 32767
	// Minimum: -32768
	Voltage *int64 `json:"voltage,omitempty"`
}

// Validate validates this power feed
func (m *PowerFeed) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAmperage(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCable(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCreated(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDisplay(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateLastUpdated(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateMaxUtilization(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateName(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePhase(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePowerPanel(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateRack(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateStatus(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSupply(formats); err != nil {
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

	if err := m.validateVoltage(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *PowerFeed) validateAmperage(formats strfmt.Registry) error {
	if swag.IsZero(m.Amperage) { // not required
		return nil
	}

	if err := validate.MinimumInt("amperage", "body", m.Amperage, 1, false); err != nil {
		return err
	}

	if err := validate.MaximumInt("amperage", "body", m.Amperage, 32767, false); err != nil {
		return err
	}

	return nil
}

func (m *PowerFeed) validateCable(formats strfmt.Registry) error {
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

func (m *PowerFeed) validateCreated(formats strfmt.Registry) error {
	if swag.IsZero(m.Created) { // not required
		return nil
	}

	if err := validate.FormatOf("created", "body", "date", m.Created.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *PowerFeed) validateDisplay(formats strfmt.Registry) error {
	if swag.IsZero(m.Display) { // not required
		return nil
	}

	if err := validate.MinLength("display", "body", m.Display, 1); err != nil {
		return err
	}

	return nil
}

func (m *PowerFeed) validateID(formats strfmt.Registry) error {
	if swag.IsZero(m.ID) { // not required
		return nil
	}

	if err := validate.FormatOf("id", "body", "uuid", m.ID.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *PowerFeed) validateLastUpdated(formats strfmt.Registry) error {
	if swag.IsZero(m.LastUpdated) { // not required
		return nil
	}

	if err := validate.FormatOf("last_updated", "body", "date-time", m.LastUpdated.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *PowerFeed) validateMaxUtilization(formats strfmt.Registry) error {
	if swag.IsZero(m.MaxUtilization) { // not required
		return nil
	}

	if err := validate.MinimumInt("max_utilization", "body", m.MaxUtilization, 1, false); err != nil {
		return err
	}

	if err := validate.MaximumInt("max_utilization", "body", m.MaxUtilization, 100, false); err != nil {
		return err
	}

	return nil
}

func (m *PowerFeed) validateName(formats strfmt.Registry) error {

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

func (m *PowerFeed) validatePhase(formats strfmt.Registry) error {
	if swag.IsZero(m.Phase) { // not required
		return nil
	}

	if m.Phase != nil {
		if err := m.Phase.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("phase")
			}
			return err
		}
	}

	return nil
}

func (m *PowerFeed) validatePowerPanel(formats strfmt.Registry) error {

	if err := validate.Required("power_panel", "body", m.PowerPanel); err != nil {
		return err
	}

	if m.PowerPanel != nil {
		if err := m.PowerPanel.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("power_panel")
			}
			return err
		}
	}

	return nil
}

func (m *PowerFeed) validateRack(formats strfmt.Registry) error {
	if swag.IsZero(m.Rack) { // not required
		return nil
	}

	if m.Rack != nil {
		if err := m.Rack.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("rack")
			}
			return err
		}
	}

	return nil
}

var powerFeedTypeStatusPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["active","failed","offline","planned"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		powerFeedTypeStatusPropEnum = append(powerFeedTypeStatusPropEnum, v)
	}
}

const (

	// PowerFeedStatusActive captures enum value "active"
	PowerFeedStatusActive string = "active"

	// PowerFeedStatusFailed captures enum value "failed"
	PowerFeedStatusFailed string = "failed"

	// PowerFeedStatusOffline captures enum value "offline"
	PowerFeedStatusOffline string = "offline"

	// PowerFeedStatusPlanned captures enum value "planned"
	PowerFeedStatusPlanned string = "planned"
)

// prop value enum
func (m *PowerFeed) validateStatusEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, powerFeedTypeStatusPropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *PowerFeed) validateStatus(formats strfmt.Registry) error {
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

func (m *PowerFeed) validateSupply(formats strfmt.Registry) error {
	if swag.IsZero(m.Supply) { // not required
		return nil
	}

	if m.Supply != nil {
		if err := m.Supply.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("supply")
			}
			return err
		}
	}

	return nil
}

func (m *PowerFeed) validateTags(formats strfmt.Registry) error {
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

func (m *PowerFeed) validateType(formats strfmt.Registry) error {
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

func (m *PowerFeed) validateURL(formats strfmt.Registry) error {
	if swag.IsZero(m.URL) { // not required
		return nil
	}

	if err := validate.FormatOf("url", "body", "uri", m.URL.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *PowerFeed) validateVoltage(formats strfmt.Registry) error {
	if swag.IsZero(m.Voltage) { // not required
		return nil
	}

	if err := validate.MinimumInt("voltage", "body", *m.Voltage, -32768, false); err != nil {
		return err
	}

	if err := validate.MaximumInt("voltage", "body", *m.Voltage, 32767, false); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this power feed based on the context it is used
func (m *PowerFeed) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
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

	if err := m.contextValidatePhase(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidatePowerPanel(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateRack(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateSupply(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateTags(ctx, formats); err != nil {
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

func (m *PowerFeed) contextValidateCable(ctx context.Context, formats strfmt.Registry) error {

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

func (m *PowerFeed) contextValidateCablePeer(ctx context.Context, formats strfmt.Registry) error {

	return nil
}

func (m *PowerFeed) contextValidateCablePeerType(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "cable_peer_type", "body", string(m.CablePeerType)); err != nil {
		return err
	}

	return nil
}

func (m *PowerFeed) contextValidateConnectedEndpoint(ctx context.Context, formats strfmt.Registry) error {

	return nil
}

func (m *PowerFeed) contextValidateConnectedEndpointReachable(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "connected_endpoint_reachable", "body", m.ConnectedEndpointReachable); err != nil {
		return err
	}

	return nil
}

func (m *PowerFeed) contextValidateConnectedEndpointType(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "connected_endpoint_type", "body", string(m.ConnectedEndpointType)); err != nil {
		return err
	}

	return nil
}

func (m *PowerFeed) contextValidateCreated(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "created", "body", strfmt.Date(m.Created)); err != nil {
		return err
	}

	return nil
}

func (m *PowerFeed) contextValidateDisplay(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "display", "body", string(m.Display)); err != nil {
		return err
	}

	return nil
}

func (m *PowerFeed) contextValidateID(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "id", "body", strfmt.UUID(m.ID)); err != nil {
		return err
	}

	return nil
}

func (m *PowerFeed) contextValidateLastUpdated(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "last_updated", "body", strfmt.DateTime(m.LastUpdated)); err != nil {
		return err
	}

	return nil
}

func (m *PowerFeed) contextValidatePhase(ctx context.Context, formats strfmt.Registry) error {

	if m.Phase != nil {
		if err := m.Phase.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("phase")
			}
			return err
		}
	}

	return nil
}

func (m *PowerFeed) contextValidatePowerPanel(ctx context.Context, formats strfmt.Registry) error {

	if m.PowerPanel != nil {
		if err := m.PowerPanel.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("power_panel")
			}
			return err
		}
	}

	return nil
}

func (m *PowerFeed) contextValidateRack(ctx context.Context, formats strfmt.Registry) error {

	if m.Rack != nil {
		if err := m.Rack.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("rack")
			}
			return err
		}
	}

	return nil
}

func (m *PowerFeed) contextValidateSupply(ctx context.Context, formats strfmt.Registry) error {

	if m.Supply != nil {
		if err := m.Supply.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("supply")
			}
			return err
		}
	}

	return nil
}

func (m *PowerFeed) contextValidateTags(ctx context.Context, formats strfmt.Registry) error {

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

func (m *PowerFeed) contextValidateType(ctx context.Context, formats strfmt.Registry) error {

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

func (m *PowerFeed) contextValidateURL(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "url", "body", strfmt.URI(m.URL)); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *PowerFeed) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *PowerFeed) UnmarshalBinary(b []byte) error {
	var res PowerFeed
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// PowerFeedPhase Phase
//
// swagger:model PowerFeedPhase
type PowerFeedPhase struct {

	// label
	// Required: true
	// Enum: [Single phase Three-phase]
	Label *string `json:"label"`

	// value
	// Required: true
	// Enum: [single-phase three-phase]
	Value *string `json:"value"`
}

func (m *PowerFeedPhase) UnmarshalJSON(b []byte) error {
	type PowerFeedPhaseAlias PowerFeedPhase
	var t PowerFeedPhaseAlias
	if err := json.Unmarshal([]byte("{\"label\":\"Single phase\",\"value\":\"single-phase\"}"), &t); err != nil {
		return err
	}
	if err := json.Unmarshal(b, &t); err != nil {
		return err
	}
	*m = PowerFeedPhase(t)
	return nil
}

// Validate validates this power feed phase
func (m *PowerFeedPhase) Validate(formats strfmt.Registry) error {
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

var powerFeedPhaseTypeLabelPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["Single phase","Three-phase"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		powerFeedPhaseTypeLabelPropEnum = append(powerFeedPhaseTypeLabelPropEnum, v)
	}
}

const (

	// PowerFeedPhaseLabelSinglePhase captures enum value "Single phase"
	PowerFeedPhaseLabelSinglePhase string = "Single phase"

	// PowerFeedPhaseLabelThreeDashPhase captures enum value "Three-phase"
	PowerFeedPhaseLabelThreeDashPhase string = "Three-phase"
)

// prop value enum
func (m *PowerFeedPhase) validateLabelEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, powerFeedPhaseTypeLabelPropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *PowerFeedPhase) validateLabel(formats strfmt.Registry) error {

	if err := validate.Required("phase"+"."+"label", "body", m.Label); err != nil {
		return err
	}

	// value enum
	if err := m.validateLabelEnum("phase"+"."+"label", "body", *m.Label); err != nil {
		return err
	}

	return nil
}

var powerFeedPhaseTypeValuePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["single-phase","three-phase"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		powerFeedPhaseTypeValuePropEnum = append(powerFeedPhaseTypeValuePropEnum, v)
	}
}

const (

	// PowerFeedPhaseValueSingleDashPhase captures enum value "single-phase"
	PowerFeedPhaseValueSingleDashPhase string = "single-phase"

	// PowerFeedPhaseValueThreeDashPhase captures enum value "three-phase"
	PowerFeedPhaseValueThreeDashPhase string = "three-phase"
)

// prop value enum
func (m *PowerFeedPhase) validateValueEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, powerFeedPhaseTypeValuePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *PowerFeedPhase) validateValue(formats strfmt.Registry) error {

	if err := validate.Required("phase"+"."+"value", "body", m.Value); err != nil {
		return err
	}

	// value enum
	if err := m.validateValueEnum("phase"+"."+"value", "body", *m.Value); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this power feed phase based on context it is used
func (m *PowerFeedPhase) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *PowerFeedPhase) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *PowerFeedPhase) UnmarshalBinary(b []byte) error {
	var res PowerFeedPhase
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// PowerFeedSupply Supply
//
// swagger:model PowerFeedSupply
type PowerFeedSupply struct {

	// label
	// Required: true
	// Enum: [AC DC]
	Label *string `json:"label"`

	// value
	// Required: true
	// Enum: [ac dc]
	Value *string `json:"value"`
}

func (m *PowerFeedSupply) UnmarshalJSON(b []byte) error {
	type PowerFeedSupplyAlias PowerFeedSupply
	var t PowerFeedSupplyAlias
	if err := json.Unmarshal([]byte("{\"label\":\"AC\",\"value\":\"ac\"}"), &t); err != nil {
		return err
	}
	if err := json.Unmarshal(b, &t); err != nil {
		return err
	}
	*m = PowerFeedSupply(t)
	return nil
}

// Validate validates this power feed supply
func (m *PowerFeedSupply) Validate(formats strfmt.Registry) error {
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

var powerFeedSupplyTypeLabelPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["AC","DC"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		powerFeedSupplyTypeLabelPropEnum = append(powerFeedSupplyTypeLabelPropEnum, v)
	}
}

const (

	// PowerFeedSupplyLabelAC captures enum value "AC"
	PowerFeedSupplyLabelAC string = "AC"

	// PowerFeedSupplyLabelDC captures enum value "DC"
	PowerFeedSupplyLabelDC string = "DC"
)

// prop value enum
func (m *PowerFeedSupply) validateLabelEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, powerFeedSupplyTypeLabelPropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *PowerFeedSupply) validateLabel(formats strfmt.Registry) error {

	if err := validate.Required("supply"+"."+"label", "body", m.Label); err != nil {
		return err
	}

	// value enum
	if err := m.validateLabelEnum("supply"+"."+"label", "body", *m.Label); err != nil {
		return err
	}

	return nil
}

var powerFeedSupplyTypeValuePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["ac","dc"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		powerFeedSupplyTypeValuePropEnum = append(powerFeedSupplyTypeValuePropEnum, v)
	}
}

const (

	// PowerFeedSupplyValueAc captures enum value "ac"
	PowerFeedSupplyValueAc string = "ac"

	// PowerFeedSupplyValueDc captures enum value "dc"
	PowerFeedSupplyValueDc string = "dc"
)

// prop value enum
func (m *PowerFeedSupply) validateValueEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, powerFeedSupplyTypeValuePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *PowerFeedSupply) validateValue(formats strfmt.Registry) error {

	if err := validate.Required("supply"+"."+"value", "body", m.Value); err != nil {
		return err
	}

	// value enum
	if err := m.validateValueEnum("supply"+"."+"value", "body", *m.Value); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this power feed supply based on context it is used
func (m *PowerFeedSupply) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *PowerFeedSupply) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *PowerFeedSupply) UnmarshalBinary(b []byte) error {
	var res PowerFeedSupply
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// PowerFeedType Type
//
// swagger:model PowerFeedType
type PowerFeedType struct {

	// label
	// Required: true
	// Enum: [Primary Redundant]
	Label *string `json:"label"`

	// value
	// Required: true
	// Enum: [primary redundant]
	Value *string `json:"value"`
}

func (m *PowerFeedType) UnmarshalJSON(b []byte) error {
	type PowerFeedTypeAlias PowerFeedType
	var t PowerFeedTypeAlias
	if err := json.Unmarshal([]byte("{\"label\":\"Primary\",\"value\":\"primary\"}"), &t); err != nil {
		return err
	}
	if err := json.Unmarshal(b, &t); err != nil {
		return err
	}
	*m = PowerFeedType(t)
	return nil
}

// Validate validates this power feed type
func (m *PowerFeedType) Validate(formats strfmt.Registry) error {
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

var powerFeedTypeTypeLabelPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["Primary","Redundant"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		powerFeedTypeTypeLabelPropEnum = append(powerFeedTypeTypeLabelPropEnum, v)
	}
}

const (

	// PowerFeedTypeLabelPrimary captures enum value "Primary"
	PowerFeedTypeLabelPrimary string = "Primary"

	// PowerFeedTypeLabelRedundant captures enum value "Redundant"
	PowerFeedTypeLabelRedundant string = "Redundant"
)

// prop value enum
func (m *PowerFeedType) validateLabelEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, powerFeedTypeTypeLabelPropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *PowerFeedType) validateLabel(formats strfmt.Registry) error {

	if err := validate.Required("type"+"."+"label", "body", m.Label); err != nil {
		return err
	}

	// value enum
	if err := m.validateLabelEnum("type"+"."+"label", "body", *m.Label); err != nil {
		return err
	}

	return nil
}

var powerFeedTypeTypeValuePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["primary","redundant"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		powerFeedTypeTypeValuePropEnum = append(powerFeedTypeTypeValuePropEnum, v)
	}
}

const (

	// PowerFeedTypeValuePrimary captures enum value "primary"
	PowerFeedTypeValuePrimary string = "primary"

	// PowerFeedTypeValueRedundant captures enum value "redundant"
	PowerFeedTypeValueRedundant string = "redundant"
)

// prop value enum
func (m *PowerFeedType) validateValueEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, powerFeedTypeTypeValuePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *PowerFeedType) validateValue(formats strfmt.Registry) error {

	if err := validate.Required("type"+"."+"value", "body", m.Value); err != nil {
		return err
	}

	// value enum
	if err := m.validateValueEnum("type"+"."+"value", "body", *m.Value); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this power feed type based on context it is used
func (m *PowerFeedType) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *PowerFeedType) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *PowerFeedType) UnmarshalBinary(b []byte) error {
	var res PowerFeedType
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
