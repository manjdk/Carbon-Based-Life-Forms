package usecase

import (
	"fmt"

	"github.com/manjdk/Carbon-Based-Life-Forms/api/domain"
	"github.com/manjdk/Carbon-Based-Life-Forms/repository"
)

type MeltMineralUC struct {
	GetOP    repository.MineralGetByIDIFace
	UpdateOP repository.MineralUpdateStateIFace
}

func NewMeltMineralUC(
	updateOP repository.MineralUpdateStateIFace,
	getOP repository.MineralGetByIDIFace,
) MeltMineralUC {
	return MeltMineralUC{
		UpdateOP: updateOP,
		GetOP:    getOP,
	}
}

func (m *MeltMineralUC) Melt(mineralID string) error {
	existingMineral, err := m.GetOP.Get(mineralID)
	if err != nil {
		return err
	}

	if existingMineral.State == domain.Liquid {
		return fmt.Errorf("mineral state is already liquid. MineralID: %s", existingMineral.ID)
	}

	existingMineral.SetState(domain.Liquid)
	existingMineral.SetFractureNumber(0)
	return m.UpdateOP.Update(existingMineral)
}
