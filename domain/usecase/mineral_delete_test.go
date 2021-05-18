package usecase

import (
	"errors"
	"testing"

	"github.com/manjdk/Carbon-Based-Life-Forms/mocks"
	"github.com/manjdk/Carbon-Based-Life-Forms/repository"
	"github.com/stretchr/testify/mock"
)

func TestDeleteMineralUC_Delete(t *testing.T) {
	type fields struct {
		DeleteOP repository.MineralDeleteIFace
	}
	tests := []struct {
		name    string
		fields  fields
		wantErr bool
	}{
		{
			name: "success",
			fields: fields{
				DeleteOP: func() repository.MineralDeleteIFace {
					deleteMock := &mocks.MineralDeleteIFace{}
					deleteMock.On("Delete", mock.AnythingOfType("string")).Return(nil)
					return deleteMock
				}(),
			},
			wantErr: false,
		},
		{
			name: "error",
			fields: fields{
				DeleteOP: func() repository.MineralDeleteIFace {
					deleteMock := &mocks.MineralDeleteIFace{}
					deleteMock.On("Delete", mock.AnythingOfType("string")).Return(errors.New("some err"))
					return deleteMock
				}(),
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := &DeleteMineralUC{
				DeleteOP: tt.fields.DeleteOP,
			}
			if err := m.Delete("test"); (err != nil) != tt.wantErr {
				t.Errorf("Delete() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
