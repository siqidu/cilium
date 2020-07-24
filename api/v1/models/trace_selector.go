// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// TraceSelector Context describing a pair of source and destination identity
//
// swagger:model TraceSelector
type TraceSelector struct {

	// from
	From *TraceFrom `json:"from,omitempty"`

	// to
	To *TraceTo `json:"to,omitempty"`

	// Enable verbose tracing.
	//
	Verbose bool `json:"verbose,omitempty"`
}

// Validate validates this trace selector
func (m *TraceSelector) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateFrom(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTo(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *TraceSelector) validateFrom(formats strfmt.Registry) error {

	if swag.IsZero(m.From) { // not required
		return nil
	}

	if m.From != nil {
		if err := m.From.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("from")
			}
			return err
		}
	}

	return nil
}

func (m *TraceSelector) validateTo(formats strfmt.Registry) error {

	if swag.IsZero(m.To) { // not required
		return nil
	}

	if m.To != nil {
		if err := m.To.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("to")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *TraceSelector) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *TraceSelector) UnmarshalBinary(b []byte) error {
	var res TraceSelector
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
