package usecase

import (
	"errors"
	"testing"

	"github.com/manjdk/Carbon-Based-Life-Forms/domain"
	"github.com/manjdk/Carbon-Based-Life-Forms/mocks"
	"github.com/manjdk/Carbon-Based-Life-Forms/repository"
	"github.com/stretchr/testify/mock"
)

func TestCreateMineralUC_Create(t *testing.T) {
	type fields struct {
		CreateOP repository.MineralCreateIFace
	}
	tests := []struct {
		name    string
		fields  fields
		wantErr bool
	}{
		{
			name: "success",
			fields: fields{
				CreateOP: func() repository.MineralCreateIFace {
					createMock := &mocks.MineralCreateIFace{}
					createMock.On("Create", mock.AnythingOfType("*domain.Mineral")).Return(nil)
					return createMock
				}(),
			},
			wantErr: false,
		},
		{
			name: "error",
			fields: fields{
				CreateOP: func() repository.MineralCreateIFace {
					createMock := &mocks.MineralCreateIFace{}
					createMock.On("Create", mock.AnythingOfType("*domain.Mineral")).Return(errors.New("some err"))
					return createMock
				}(),
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := &CreateMineralUC{
				CreateOP: tt.fields.CreateOP,
			}
			if err := m.Create(&domain.Mineral{}); (err != nil) != tt.wantErr {
				t.Errorf("Create() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
