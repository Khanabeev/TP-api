package service

import (
	"github.com/Khanabeev/TP-api/internal/app/domain"
	"github.com/Khanabeev/TP-api/internal/app/dto"
	"github.com/Khanabeev/TP-api/pkg/errs"
)

type MasterService interface {
	GetMasterById(string) (*dto.GetMasterByIdResponse, *errs.AppError)
}

type DefaultMasterService struct {
	repository domain.MasterRepository
}

func NewMasterService(repository domain.MasterRepository) DefaultMasterService {
	return DefaultMasterService{
		repository: repository,
	}
}

func (s DefaultMasterService) GetMasterById(id string) (*dto.GetMasterByIdResponse, *errs.AppError) {

	master, err := s.repository.FindById(id)
	if err != nil {
		return nil, err
	}

	response := master.ToGetMasterByIdResponseDto()

	return &response, nil
}
