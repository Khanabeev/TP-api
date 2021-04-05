package dto

import "database/sql"

type GetMasterByIdResponse struct {
	ID                int            `json:"id"`
	Name              string         `json:"name"`
	Email             string         `json:"email"`
	Status            string         `json:"status"`
	Photo             sql.NullString `json:"photo"`
	Gender            string         `json:"gender"`
	Description       sql.NullString `json:"description"`
	ExperienceFrom    sql.NullTime   `json:"experience_from"`
	IsWorkingWithMen  bool           `json:"is_working_with_men"`
	HasSingleUseItems bool           `json:"has_single_use_items"`
	Instagram         sql.NullString `json:"instagram"`
	Education         sql.NullString `json:"education"`
	PhoneNumber       string         `json:"phone_number"`
	HasWhatsapp       bool           `json:"has_whatsapp"`
	Materials         sql.NullString `json:"materials"`
	CreatedAt         sql.NullTime   `json:"created_at"`
	UpdatedAt         sql.NullTime   `json:"updated_at"`
}
