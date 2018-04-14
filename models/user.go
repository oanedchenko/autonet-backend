// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"strconv"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// User user
// swagger:model User
type User struct {

	// avatar
	Avatar strfmt.UUID `json:"avatar,omitempty"`

	// contacts
	Contacts []*Contact `json:"contacts"`

	// department Id
	DepartmentID strfmt.UUID `json:"departmentId,omitempty"`

	// department name
	DepartmentName string `json:"departmentName,omitempty"`

	// first name
	FirstName string `json:"firstName,omitempty"`

	// id
	ID strfmt.UUID `json:"id,omitempty"`

	// job title
	JobTitle string `json:"jobTitle,omitempty"`

	// last name
	LastName string `json:"lastName,omitempty"`

	// team Id
	TeamID strfmt.UUID `json:"teamId,omitempty"`

	// team name
	TeamName string `json:"teamName,omitempty"`
}

// Validate validates this user
func (m *User) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAvatar(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateContacts(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateDepartmentID(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateID(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateTeamID(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *User) validateAvatar(formats strfmt.Registry) error {

	if swag.IsZero(m.Avatar) { // not required
		return nil
	}

	if err := validate.FormatOf("avatar", "body", "uuid", m.Avatar.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *User) validateContacts(formats strfmt.Registry) error {

	if swag.IsZero(m.Contacts) { // not required
		return nil
	}

	for i := 0; i < len(m.Contacts); i++ {

		if swag.IsZero(m.Contacts[i]) { // not required
			continue
		}

		if m.Contacts[i] != nil {

			if err := m.Contacts[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("contacts" + "." + strconv.Itoa(i))
				}
				return err
			}

		}

	}

	return nil
}

func (m *User) validateDepartmentID(formats strfmt.Registry) error {

	if swag.IsZero(m.DepartmentID) { // not required
		return nil
	}

	if err := validate.FormatOf("departmentId", "body", "uuid", m.DepartmentID.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *User) validateID(formats strfmt.Registry) error {

	if swag.IsZero(m.ID) { // not required
		return nil
	}

	if err := validate.FormatOf("id", "body", "uuid", m.ID.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *User) validateTeamID(formats strfmt.Registry) error {

	if swag.IsZero(m.TeamID) { // not required
		return nil
	}

	if err := validate.FormatOf("teamId", "body", "uuid", m.TeamID.String(), formats); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *User) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *User) UnmarshalBinary(b []byte) error {
	var res User
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
