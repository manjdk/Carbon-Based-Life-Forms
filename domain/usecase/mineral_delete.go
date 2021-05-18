package usecase

import (
	"github.com/manjdk/Carbon-Based-Life-Forms/repository"
)

type DeleteMineralUC struct {
	DeleteOP repository.MineralDeleteIFace
}

func NewDeleteMineralUC(
	deleteOP repository.MineralDeleteIFace,
) DeleteMineralUC {
	return DeleteMineralUC{
		DeleteOP: deleteOP,
	}
}

func (m *DeleteMineralUC) Delete(mineralID string) error {
	return m.DeleteOP.Delete(mineralID)
}
