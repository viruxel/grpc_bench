package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/validate"
)

// Event event
// swagger:model Event
type Event struct {

	// event sport
	// Required: true
	Category *string `json:"category"`

	// Unique identifier of event
	// Required: true
	ID *int32 `json:"id"`

	// event sport
	// Required: true
	Round *string `json:"round"`

	// event sport
	// Required: true
	Score *string `json:"score"`

	// event sport
	// Required: true
	Sport *string `json:"sport"`

	// event sport
	// Required: true
	StartTimestamp *float64 `json:"start_timestamp"`

	// event sport
	// Required: true
	State *string `json:"state"`

	// event sport
	// Required: true
	TeamA *string `json:"team_a"`

	// event sport
	// Required: true
	TeamB *string `json:"team_b"`

	// event sport
	// Required: true
	Tournament *string `json:"tournament"`
}

// Validate validates this event
func (m *Event) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCategory(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateID(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateRound(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateScore(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateSport(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateStartTimestamp(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateState(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateTeamA(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateTeamB(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateTournament(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Event) validateCategory(formats strfmt.Registry) error {

	if err := validate.Required("category", "body", m.Category); err != nil {
		return err
	}

	return nil
}

func (m *Event) validateID(formats strfmt.Registry) error {

	if err := validate.Required("id", "body", m.ID); err != nil {
		return err
	}

	return nil
}

func (m *Event) validateRound(formats strfmt.Registry) error {

	if err := validate.Required("round", "body", m.Round); err != nil {
		return err
	}

	return nil
}

func (m *Event) validateScore(formats strfmt.Registry) error {

	if err := validate.Required("score", "body", m.Score); err != nil {
		return err
	}

	return nil
}

func (m *Event) validateSport(formats strfmt.Registry) error {

	if err := validate.Required("sport", "body", m.Sport); err != nil {
		return err
	}

	return nil
}

func (m *Event) validateStartTimestamp(formats strfmt.Registry) error {

	if err := validate.Required("start_timestamp", "body", m.StartTimestamp); err != nil {
		return err
	}

	return nil
}

func (m *Event) validateState(formats strfmt.Registry) error {

	if err := validate.Required("state", "body", m.State); err != nil {
		return err
	}

	return nil
}

func (m *Event) validateTeamA(formats strfmt.Registry) error {

	if err := validate.Required("team_a", "body", m.TeamA); err != nil {
		return err
	}

	return nil
}

func (m *Event) validateTeamB(formats strfmt.Registry) error {

	if err := validate.Required("team_b", "body", m.TeamB); err != nil {
		return err
	}

	return nil
}

func (m *Event) validateTournament(formats strfmt.Registry) error {

	if err := validate.Required("tournament", "body", m.Tournament); err != nil {
		return err
	}

	return nil
}
