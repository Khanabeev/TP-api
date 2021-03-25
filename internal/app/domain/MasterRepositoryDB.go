package domain

import (
	"database/sql"
	"github.com/Khanabeev/TP-api/pkg/errs"
)

type MasterRepositoryDB struct {
	client *sql.DB
}

func (r MasterRepositoryDB) All() ([]*Master, *errs.AppError) {
	return nil, nil
}

func (r MasterRepositoryDB) FindById(masterId int) (*Master, *errs.AppError) {
	return nil, nil
}
