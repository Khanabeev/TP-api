package domain

import (
	"database/sql"

	"github.com/Khanabeev/TP-api/internal/app/dto"
	"github.com/Khanabeev/TP-api/pkg/errs"
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

type MasterRepository interface {
	All() ([]*Master, *errs.AppError)
	FindById(id string) (*Master, *errs.AppError)
}

func (m Master) ToGetMasterByIdResponseDto() dto.GetMasterByIdResponse {
	return dto.GetMasterByIdResponse{
		ID:                m.ID,
		Name:              m.Name,
		Email:             m.Email,
		Status:            m.Status,
		Photo:             m.Photo,
		Gender:            m.Gender,
		Description:       m.Description,
		ExperienceFrom:    m.ExperienceFrom,
		IsWorkingWithMen:  m.IsWorkingWithMen,
		HasSingleUseItems: m.HasSingleUseItems,
		Instagram:         m.Instagram,
		Education:         m.Education,
		PhoneNumber:       m.PhoneNumber,
		HasWhatsapp:       m.HasWhatsapp,
		Materials:         m.Materials,
		CreatedAt:         m.CreatedAt,
		UpdatedAt:         m.UpdatedAt,
	}
}
