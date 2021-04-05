package domain

import (
	"database/sql"

	"github.com/Khanabeev/TP-api/pkg/errs"
	"github.com/Khanabeev/TP-api/pkg/logger"
)

type MasterRepositoryDB struct {
	dbClient *sql.DB
}

func NewMasterRepository(dbClient *sql.DB) MasterRepositoryDB {
	return MasterRepositoryDB{
		dbClient: dbClient,
	}
}

func (r MasterRepositoryDB) All() ([]*Master, *errs.AppError) {
	return nil, nil
}

func (r MasterRepositoryDB) FindById(id string) (*Master, *errs.AppError) {
	stmt := `SELECT 
				id,
				name,
				email,
				status,
				photo,
				gender,
				description,
				experience_from,
				is_working_with_men,
				has_single_use_items,
				instagram,
				education,
				phone_number,
				has_whatsapp,
				materials,
				created_at,
				updated_at
			FROM public.masters
			WHERE id = $1`
	row := r.dbClient.QueryRow(stmt, id)
	if row.Err() != nil {
		logger.Error("Error while getting master: " + row.Err().Error())
		return nil, errs.NewUnexpectedError("Unexpected error from database")
	}
	var m Master

	err := row.Scan(
		&m.ID,
		&m.Name,
		&m.Email,
		&m.Status,
		&m.Photo,
		&m.Gender,
		&m.Description,
		&m.ExperienceFrom,
		&m.IsWorkingWithMen,
		&m.HasSingleUseItems,
		&m.Instagram,
		&m.Education,
		&m.PhoneNumber,
		&m.HasWhatsapp,
		&m.Materials,
		&m.CreatedAt,
		&m.UpdatedAt,
	)

	if err == sql.ErrNoRows {
		return nil, errs.NewNotFoundError("Master Not Found")
	}

	if err != nil {
		logger.Error("Error while scanning row from database: " + err.Error())
		return nil, errs.NewUnexpectedError("Unexpected database error")
	}

	return &m, nil
}
