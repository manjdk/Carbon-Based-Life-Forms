package usecase

import (
	"github.com/manjdk/Carbon-Based-Life-Forms/domain"
	"github.com/manjdk/Carbon-Based-Life-Forms/repository"
)

type GetMineralUC struct {
	GetOP repository.MineralGetByIDIFace
}

func NewGetMineralUC(getOP repository.MineralGetByIDIFace) GetMineralUC {
	return GetMineralUC{
		GetOP: getOP,
	}
}

func (m *GetMineralUC) GetByID(mineralID string) (*domain.Mineral, error) {
	return m.GetOP.Get(mineralID)
}
