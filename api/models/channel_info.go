// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// ChannelInfo channel info
// swagger:model ChannelInfo
type ChannelInfo struct {
	CEWithoutDataWithKeptncontext

	// data
	Data *ChannelInfoAO1Data `json:"data,omitempty"`
}

// UnmarshalJSON unmarshals this object from a JSON structure
func (m *ChannelInfo) UnmarshalJSON(raw []byte) error {
	// AO0
	var aO0 CEWithoutDataWithKeptncontext
	if err := swag.ReadJSON(raw, &aO0); err != nil {
		return err
	}
	m.CEWithoutDataWithKeptncontext = aO0

	// AO1
	var dataAO1 struct {
		Data *ChannelInfoAO1Data `json:"data,omitempty"`
	}
	if err := swag.ReadJSON(raw, &dataAO1); err != nil {
		return err
	}

	m.Data = dataAO1.Data

	return nil
}

// MarshalJSON marshals this object to a JSON structure
func (m ChannelInfo) MarshalJSON() ([]byte, error) {
	_parts := make([][]byte, 0, 2)

	aO0, err := swag.WriteJSON(m.CEWithoutDataWithKeptncontext)
	if err != nil {
		return nil, err
	}
	_parts = append(_parts, aO0)

	var dataAO1 struct {
		Data *ChannelInfoAO1Data `json:"data,omitempty"`
	}

	dataAO1.Data = m.Data

	jsonDataAO1, errAO1 := swag.WriteJSON(dataAO1)
	if errAO1 != nil {
		return nil, errAO1
	}
	_parts = append(_parts, jsonDataAO1)

	return swag.ConcatJSON(_parts...), nil
}

// Validate validates this channel info
func (m *ChannelInfo) Validate(formats strfmt.Registry) error {
	var res []error

	// validation for a type composition with CEWithoutDataWithKeptncontext
	if err := m.CEWithoutDataWithKeptncontext.Validate(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateData(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ChannelInfo) validateData(formats strfmt.Registry) error {

	if swag.IsZero(m.Data) { // not required
		return nil
	}

	if m.Data != nil {
		if err := m.Data.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("data")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *ChannelInfo) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ChannelInfo) UnmarshalBinary(b []byte) error {
	var res ChannelInfo
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// ChannelInfoAO1Data channel info a o1 data
// swagger:model ChannelInfoAO1Data
type ChannelInfoAO1Data struct {

	// channel info
	ChannelInfo *ChannelInfoAO1DataChannelInfo `json:"channelInfo,omitempty"`
}

// Validate validates this channel info a o1 data
func (m *ChannelInfoAO1Data) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateChannelInfo(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ChannelInfoAO1Data) validateChannelInfo(formats strfmt.Registry) error {

	if swag.IsZero(m.ChannelInfo) { // not required
		return nil
	}

	if m.ChannelInfo != nil {
		if err := m.ChannelInfo.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("data" + "." + "channelInfo")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *ChannelInfoAO1Data) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ChannelInfoAO1Data) UnmarshalBinary(b []byte) error {
	var res ChannelInfoAO1Data
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// ChannelInfoAO1DataChannelInfo channel info a o1 data channel info
// swagger:model ChannelInfoAO1DataChannelInfo
type ChannelInfoAO1DataChannelInfo struct {

	// channel ID
	// Required: true
	ChannelID *string `json:"channelID"`

	// token
	// Required: true
	Token *string `json:"token"`
}

// Validate validates this channel info a o1 data channel info
func (m *ChannelInfoAO1DataChannelInfo) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateChannelID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateToken(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ChannelInfoAO1DataChannelInfo) validateChannelID(formats strfmt.Registry) error {

	if err := validate.Required("data"+"."+"channelInfo"+"."+"channelID", "body", m.ChannelID); err != nil {
		return err
	}

	return nil
}

func (m *ChannelInfoAO1DataChannelInfo) validateToken(formats strfmt.Registry) error {

	if err := validate.Required("data"+"."+"channelInfo"+"."+"token", "body", m.Token); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *ChannelInfoAO1DataChannelInfo) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ChannelInfoAO1DataChannelInfo) UnmarshalBinary(b []byte) error {
	var res ChannelInfoAO1DataChannelInfo
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}