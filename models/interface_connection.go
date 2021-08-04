package models

import (
	"context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// InterfaceConnection interface connection
//
// swagger:model InterfaceConnection
type InterfaceConnection struct {

	// Connected endpoint reachable
	// Read Only: true
	ConnectedEndpointReachable *bool `json:"connected_endpoint_reachable,omitempty"`

	// Display
	//
	// Human friendly display value
	// Read Only: true
	// Min Length: 1
	Display string `json:"display,omitempty"`

	// interface a
	Interfacea *NestedInterface `json:"interface_a,omitempty"`

	// interface b
	// Required: true
	Interfaceb *NestedInterface `json:"interface_b"`
}

// Validate validates this interface connection
func (m *InterfaceConnection) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateDisplay(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateInterfacea(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateInterfaceb(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *InterfaceConnection) validateDisplay(formats strfmt.Registry) error {
	if swag.IsZero(m.Display) { // not required
		return nil
	}

	if err := validate.MinLength("display", "body", m.Display, 1); err != nil {
		return err
	}

	return nil
}

func (m *InterfaceConnection) validateInterfacea(formats strfmt.Registry) error {
	if swag.IsZero(m.Interfacea) { // not required
		return nil
	}

	if m.Interfacea != nil {
		if err := m.Interfacea.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("interface_a")
			}
			return err
		}
	}

	return nil
}

func (m *InterfaceConnection) validateInterfaceb(formats strfmt.Registry) error {

	if err := validate.Required("interface_b", "body", m.Interfaceb); err != nil {
		return err
	}

	if m.Interfaceb != nil {
		if err := m.Interfaceb.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("interface_b")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this interface connection based on the context it is used
func (m *InterfaceConnection) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateConnectedEndpointReachable(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateDisplay(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateInterfacea(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateInterfaceb(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *InterfaceConnection) contextValidateConnectedEndpointReachable(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "connected_endpoint_reachable", "body", m.ConnectedEndpointReachable); err != nil {
		return err
	}

	return nil
}

func (m *InterfaceConnection) contextValidateDisplay(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "display", "body", string(m.Display)); err != nil {
		return err
	}

	return nil
}

func (m *InterfaceConnection) contextValidateInterfacea(ctx context.Context, formats strfmt.Registry) error {

	if m.Interfacea != nil {
		if err := m.Interfacea.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("interface_a")
			}
			return err
		}
	}

	return nil
}

func (m *InterfaceConnection) contextValidateInterfaceb(ctx context.Context, formats strfmt.Registry) error {

	if m.Interfaceb != nil {
		if err := m.Interfaceb.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("interface_b")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *InterfaceConnection) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *InterfaceConnection) UnmarshalBinary(b []byte) error {
	var res InterfaceConnection
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
