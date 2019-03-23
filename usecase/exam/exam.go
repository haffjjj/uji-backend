package exam

import (
	"github.com/haffjjj/uji-backend/models"
	"github.com/haffjjj/uji-backend/repository/exam"
)

type examUsecase struct {
	eRepository exam.Repository
}

//NewExamUsecase represent initializatin courseUsecase
func NewExamUsecase(eR exam.Repository) Usecase {
	return &examUsecase{eR}
}

func (eU *examUsecase) FetchG(filter models.Filter) ([]*models.ExamG, error) {
	examsGs, err := eU.eRepository.FetchG(filter)

	if err != nil {
		return nil, err
	}

	return examsGs, nil
}