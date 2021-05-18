package usecase

import (
	"fmt"

	"github.com/manjdk/Carbon-Based-Life-Forms/domain"
	"github.com/manjdk/Carbon-Based-Life-Forms/repository"
)

type FractureMineralUC struct {
	GetOP    repository.MineralGetByIDIFace
	UpdateOP repository.MineralUpdateStateIFace
}

func NewFractureMineralUC(
	updateOP repository.MineralUpdateStateIFace,
	getOP repository.MineralGetByIDIFace,
) FractureMineralUC {
	return FractureMineralUC{
		UpdateOP: updateOP,
		GetOP:    getOP,
	}
}

func (m *FractureMineralUC) Fracture(mineralID string) error {
	existingMineral, err := m.GetOP.Get(mineralID)
	if err != nil {
		return err
	}

	if existingMineral.State == domain.Liquid {
		return fmt.Errorf("mineral can not be converted from liquid to fractures. MineralID: %s", existingMineral.ID)
	}

	switch existingMineral.State {
	case domain.Fractured:
		if existingMineral.Fractures < 1 {
			existingMineral.SetFractureNumber(1)
		}

		existingMineral.SetFractureNumber(existingMineral.Fractures * 2)
	case domain.Solid:
		existingMineral.SetFractureNumber(2)
	}

	existingMineral.SetState(domain.Fractured)
	return m.UpdateOP.Update(existingMineral)
}
