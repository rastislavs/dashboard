// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// ApplicationDefinitionListItemSpec ApplicationDefinitionListItemSpec defines the desired state of ApplicationDefinitionListItemSpec.
//
// swagger:model ApplicationDefinitionListItemSpec
type ApplicationDefinitionListItemSpec struct {

	// Description of the application. what is its purpose
	Description string `json:"description,omitempty"`

	// Labels can contain metadata about the application, such as the owner who manages it.
	Labels map[string]string `json:"labels,omitempty"`
}

// Validate validates this application definition list item spec
func (m *ApplicationDefinitionListItemSpec) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this application definition list item spec based on context it is used
func (m *ApplicationDefinitionListItemSpec) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *ApplicationDefinitionListItemSpec) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ApplicationDefinitionListItemSpec) UnmarshalBinary(b []byte) error {
	var res ApplicationDefinitionListItemSpec
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
