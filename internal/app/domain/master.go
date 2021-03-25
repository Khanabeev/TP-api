package domain

import (
	"database/sql"
)

type Master struct {
	ID                int            `db:"id"`
	Name              string         `db:"name"`
	Email             string         `db:"email"`
	Password          string         `db:"password"`
	Status            string         `db:"status"`
	Photo             sql.NullString `db:"photo"`
	Gender            string         `db:"gender"`
	Description       sql.NullString `db:"description"`
	ExperienceFrom    sql.NullTime   `db:"experience_from"`
	IsWorkingWithMen  bool           `db:"is_working_with_men"`
	HasSingleUseItems bool           `db:"has_single_use_items"`
	Instagram         sql.NullString `db:"instagram"`
	Education         sql.NullString `db:"education"`
	PhoneNumber       string         `db:"phone_number"`
	HasWhatsapp       bool           `db:"has_whatsapp"`
	Materials         sql.NullString `db:"materials"`
	CreatedAt         sql.NullTime   `db:"created_at"`
	UpdatedAt         sql.NullTime   `db:"updated_at"`
}
