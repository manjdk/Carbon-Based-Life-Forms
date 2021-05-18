package usecase

import (
	"github.com/manjdk/Carbon-Based-Life-Forms/domain"
	"github.com/manjdk/Carbon-Based-Life-Forms/repository"
)

type GetAllMineralUC struct {
	GetAllOP        repository.MineralsGetIFace
	GetByClientIDOP repository.MineralsGetByClientIDIFace
}

func NewGetAllMineralUC(
	getAllOP repository.MineralsGetIFace,
	getByClientIDOP repository.MineralsGetByClientIDIFace,
) GetAllMineralUC {
	return GetAllMineralUC{
		GetAllOP:        getAllOP,
		GetByClientIDOP: getByClientIDOP,
	}
}

func (m *GetAllMineralUC) GetAll(clientID string) ([]domain.Mineral, error) {
	switch clientID {
	case "":
		return m.GetAllOP.GetAll()
	default:
		return m.GetByClientIDOP.GetByClientID(clientID)
	}
}
