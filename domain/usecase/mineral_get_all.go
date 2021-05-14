package usecase

import (
	"github.com/manjdk/Carbon-Based-Life-Forms/domain"
	"github.com/manjdk/Carbon-Based-Life-Forms/repository"
)

type GetAllMineralUC struct {
	GetAllOP repository.MineralsGetIFace
}

func NewGetAllMineralUC(getAllOP repository.MineralsGetIFace) GetAllMineralUC {
	return GetAllMineralUC{
		GetAllOP: getAllOP,
	}
}

func (m *GetAllMineralUC) GetAll() ([]domain.Mineral, error) {
	return m.GetAllOP.GetAll()
}
