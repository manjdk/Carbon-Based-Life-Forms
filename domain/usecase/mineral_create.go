package usecase

import (
	"github.com/manjdk/Carbon-Based-Life-Forms/domain"
	"github.com/manjdk/Carbon-Based-Life-Forms/repository"
)

type CreateMineralUC struct {
	CreateOP repository.MineralCreateIFace
}

func NewCreateMineralUC(createOP repository.MineralCreateIFace) CreateMineralUC {
	return CreateMineralUC{
		CreateOP: createOP,
	}
}

func (m *CreateMineralUC) Create(mineral *domain.Mineral) error {
	return m.CreateOP.Create(mineral)
}
