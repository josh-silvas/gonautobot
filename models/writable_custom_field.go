package models

import (
	"context"
	"encoding/json"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// WritableCustomField writable custom field
//
// swagger:model WritableCustomField
type WritableCustomField struct {

	// content types
	// Required: true
	// Unique: true
	ContentTypes []string `json:"content_types"`

	// Default
	//
	// Default value for the field (must be a JSON value). Encapsulate strings with double quotes (e.g. "Foo").
	Default *string `json:"default,omitempty"`

	// Description
	// Max Length: 200
	Description string `json:"description,omitempty"`

	// Display
	//
	// Human friendly display value
	// Read Only: true
	// Min Length: 1
	Display string `json:"display,omitempty"`

	// Filter logic
	//
	// Loose matches any instance of a given string; exact matches the entire field.
	// Enum: [disabled loose exact]
	FilterLogic string `json:"filter_logic,omitempty"`

	// Id
	// Read Only: true
	// Format: uuid
	ID strfmt.UUID `json:"id,omitempty"`

	// Label
	//
	// Name of the field as displayed to users (if not provided, the field's name will be used)
	// Max Length: 50
	Label string `json:"label,omitempty"`

	// Name
	//
	// Internal field name
	// Required: true
	// Max Length: 50
	// Min Length: 1
	Name *string `json:"name"`

	// Required
	//
	// If true, this field is required when creating new objects or editing an existing object.
	Required bool `json:"required,omitempty"`

	// Type
	// Enum: [text integer boolean date url select multi-select]
	Type string `json:"type,omitempty"`

	// Url
	// Read Only: true
	// Format: uri
	URL strfmt.URI `json:"url,omitempty"`

	// Maximum value
	//
	// Maximum allowed value (for numeric fields)
	// Maximum: 9.223372036854776e+18
	// Minimum: -9.223372036854776e+18
	ValidationMaximum *int64 `json:"validation_maximum,omitempty"`

	// Minimum value
	//
	// Minimum allowed value (for numeric fields)
	// Maximum: 9.223372036854776e+18
	// Minimum: -9.223372036854776e+18
	ValidationMinimum *int64 `json:"validation_minimum,omitempty"`

	// Validation regex
	//
	// Regular expression to enforce on text field values. Use ^ and $ to force matching of entire string. For example, <code>^[A-Z]{3}$</code> will limit values to exactly three uppercase letters.
	// Max Length: 500
	ValidationRegex string `json:"validation_regex,omitempty"`

	// Weight
	//
	// Fields with higher weights appear lower in a form.
	// Maximum: 32767
	// Minimum: 0
	Weight *int64 `json:"weight,omitempty"`
}

// Validate validates this writable custom field
func (m *WritableCustomField) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateContentTypes(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDescription(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDisplay(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateFilterLogic(formats); err != nil {
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

	if err := m.validateValidationMaximum(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateValidationMinimum(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateValidationRegex(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateWeight(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *WritableCustomField) validateContentTypes(formats strfmt.Registry) error {

	if err := validate.Required("content_types", "body", m.ContentTypes); err != nil {
		return err
	}

	if err := validate.UniqueItems("content_types", "body", m.ContentTypes); err != nil {
		return err
	}

	return nil
}

func (m *WritableCustomField) validateDescription(formats strfmt.Registry) error {
	if swag.IsZero(m.Description) { // not required
		return nil
	}

	if err := validate.MaxLength("description", "body", m.Description, 200); err != nil {
		return err
	}

	return nil
}

func (m *WritableCustomField) validateDisplay(formats strfmt.Registry) error {
	if swag.IsZero(m.Display) { // not required
		return nil
	}

	if err := validate.MinLength("display", "body", m.Display, 1); err != nil {
		return err
	}

	return nil
}

var writableCustomFieldTypeFilterLogicPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["disabled","loose","exact"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		writableCustomFieldTypeFilterLogicPropEnum = append(writableCustomFieldTypeFilterLogicPropEnum, v)
	}
}

const (

	// WritableCustomFieldFilterLogicDisabled captures enum value "disabled"
	WritableCustomFieldFilterLogicDisabled string = "disabled"

	// WritableCustomFieldFilterLogicLoose captures enum value "loose"
	WritableCustomFieldFilterLogicLoose string = "loose"

	// WritableCustomFieldFilterLogicExact captures enum value "exact"
	WritableCustomFieldFilterLogicExact string = "exact"
)

// prop value enum
func (m *WritableCustomField) validateFilterLogicEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, writableCustomFieldTypeFilterLogicPropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *WritableCustomField) validateFilterLogic(formats strfmt.Registry) error {
	if swag.IsZero(m.FilterLogic) { // not required
		return nil
	}

	// value enum
	if err := m.validateFilterLogicEnum("filter_logic", "body", m.FilterLogic); err != nil {
		return err
	}

	return nil
}

func (m *WritableCustomField) validateID(formats strfmt.Registry) error {
	if swag.IsZero(m.ID) { // not required
		return nil
	}

	if err := validate.FormatOf("id", "body", "uuid", m.ID.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *WritableCustomField) validateLabel(formats strfmt.Registry) error {
	if swag.IsZero(m.Label) { // not required
		return nil
	}

	if err := validate.MaxLength("label", "body", m.Label, 50); err != nil {
		return err
	}

	return nil
}

func (m *WritableCustomField) validateName(formats strfmt.Registry) error {

	if err := validate.Required("name", "body", m.Name); err != nil {
		return err
	}

	if err := validate.MinLength("name", "body", *m.Name, 1); err != nil {
		return err
	}

	if err := validate.MaxLength("name", "body", *m.Name, 50); err != nil {
		return err
	}

	return nil
}

var writableCustomFieldTypeTypePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["text","integer","boolean","date","url","select","multi-select"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		writableCustomFieldTypeTypePropEnum = append(writableCustomFieldTypeTypePropEnum, v)
	}
}

const (

	// WritableCustomFieldTypeText captures enum value "text"
	WritableCustomFieldTypeText string = "text"

	// WritableCustomFieldTypeInteger captures enum value "integer"
	WritableCustomFieldTypeInteger string = "integer"

	// WritableCustomFieldTypeBoolean captures enum value "boolean"
	WritableCustomFieldTypeBoolean string = "boolean"

	// WritableCustomFieldTypeDate captures enum value "date"
	WritableCustomFieldTypeDate string = "date"

	// WritableCustomFieldTypeURL captures enum value "url"
	WritableCustomFieldTypeURL string = "url"

	// WritableCustomFieldTypeSelect captures enum value "select"
	WritableCustomFieldTypeSelect string = "select"

	// WritableCustomFieldTypeMultiDashSelect captures enum value "multi-select"
	WritableCustomFieldTypeMultiDashSelect string = "multi-select"
)

// prop value enum
func (m *WritableCustomField) validateTypeEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, writableCustomFieldTypeTypePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *WritableCustomField) validateType(formats strfmt.Registry) error {
	if swag.IsZero(m.Type) { // not required
		return nil
	}

	// value enum
	if err := m.validateTypeEnum("type", "body", m.Type); err != nil {
		return err
	}

	return nil
}

func (m *WritableCustomField) validateURL(formats strfmt.Registry) error {
	if swag.IsZero(m.URL) { // not required
		return nil
	}

	if err := validate.FormatOf("url", "body", "uri", m.URL.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *WritableCustomField) validateValidationMaximum(formats strfmt.Registry) error {
	if swag.IsZero(m.ValidationMaximum) { // not required
		return nil
	}

	if err := validate.MinimumInt("validation_maximum", "body", *m.ValidationMaximum, -9.223372036854776e+18, false); err != nil {
		return err
	}

	if err := validate.MaximumInt("validation_maximum", "body", *m.ValidationMaximum, 9.223372036854776e+18, false); err != nil {
		return err
	}

	return nil
}

func (m *WritableCustomField) validateValidationMinimum(formats strfmt.Registry) error {
	if swag.IsZero(m.ValidationMinimum) { // not required
		return nil
	}

	if err := validate.MinimumInt("validation_minimum", "body", *m.ValidationMinimum, -9.223372036854776e+18, false); err != nil {
		return err
	}

	if err := validate.MaximumInt("validation_minimum", "body", *m.ValidationMinimum, 9.223372036854776e+18, false); err != nil {
		return err
	}

	return nil
}

func (m *WritableCustomField) validateValidationRegex(formats strfmt.Registry) error {
	if swag.IsZero(m.ValidationRegex) { // not required
		return nil
	}

	if err := validate.MaxLength("validation_regex", "body", m.ValidationRegex, 500); err != nil {
		return err
	}

	return nil
}

func (m *WritableCustomField) validateWeight(formats strfmt.Registry) error {
	if swag.IsZero(m.Weight) { // not required
		return nil
	}

	if err := validate.MinimumInt("weight", "body", *m.Weight, 0, false); err != nil {
		return err
	}

	if err := validate.MaximumInt("weight", "body", *m.Weight, 32767, false); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this writable custom field based on the context it is used
func (m *WritableCustomField) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

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

func (m *WritableCustomField) contextValidateDisplay(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "display", "body", string(m.Display)); err != nil {
		return err
	}

	return nil
}

func (m *WritableCustomField) contextValidateID(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "id", "body", strfmt.UUID(m.ID)); err != nil {
		return err
	}

	return nil
}

func (m *WritableCustomField) contextValidateURL(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "url", "body", strfmt.URI(m.URL)); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *WritableCustomField) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *WritableCustomField) UnmarshalBinary(b []byte) error {
	var res WritableCustomField
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
