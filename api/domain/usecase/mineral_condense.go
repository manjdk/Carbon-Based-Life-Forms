package usecase

import (
	"fmt"

	"github.com/manjdk/Carbon-Based-Life-Forms/api/domain"
	"github.com/manjdk/Carbon-Based-Life-Forms/repository"
)

type CondenseMineralUC struct {
	GetOP    repository.MineralGetByIDIFace
	UpdateOP repository.MineralUpdateStateIFace
}

func NewCondenseMineralUC(
	updateOP repository.MineralUpdateStateIFace,
	getOP repository.MineralGetByIDIFace,
) CondenseMineralUC {
	return CondenseMineralUC{
		UpdateOP: updateOP,
		GetOP:    getOP,
	}
}

func (m *CondenseMineralUC) Condense(mineralID string) error {
	existingMineral, err := m.GetOP.Get(mineralID)
	if err != nil {
		return err
	}

	if existingMineral.State == domain.Solid {
		return fmt.Errorf("mineral is already in solid state: %s", existingMineral.ID)
	}

	existingMineral.SetState(domain.Solid)
	existingMineral.SetFractureNumber(0)
	return m.UpdateOP.Update(existingMineral)
}
