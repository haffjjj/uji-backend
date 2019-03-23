package examgroup

import (
	"github.com/haffjjj/uji-backend/models"
	"github.com/haffjjj/uji-backend/repository/examgroup"
)

type examGroupUsecase struct {
	cRepository examgroup.Repository
}

//NewExamGroupUsecase represent initializatin courseUsecase
func NewExamGroupUsecase(cR examgroup.Repository) Usecase {
	return &examGroupUsecase{cR}
}

func (c *examGroupUsecase) FetchG(filter models.Filter) ([]*models.ExamGroupG, error) {
	examGroupGs, err := c.cRepository.FetchG(filter)

	if err != nil {
		return nil, err
	}

	return examGroupGs, nil
}