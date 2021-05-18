package usecase

import (
	"errors"
	"testing"

	"github.com/manjdk/Carbon-Based-Life-Forms/domain"

	"github.com/manjdk/Carbon-Based-Life-Forms/mocks"
	"github.com/stretchr/testify/mock"

	"github.com/manjdk/Carbon-Based-Life-Forms/repository"
)

func TestCondenseMineralUC_Condense(t *testing.T) {
	type fields struct {
		GetOP    repository.MineralGetByIDIFace
		UpdateOP repository.MineralUpdateStateIFace
	}
	type args struct {
		mineralID string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{
			name: "success",
			fields: fields{
				GetOP: func() repository.MineralGetByIDIFace {
					getMock := &mocks.MineralGetByIDIFace{}
					getMock.On("Get", mock.AnythingOfType("string")).Return(
						&domain.Mineral{State: domain.Liquid}, nil)
					return getMock
				}(),
				UpdateOP: func() repository.MineralUpdateStateIFace {
					updateMock := &mocks.MineralUpdateStateIFace{}
					updateMock.On("Update", mock.AnythingOfType("*domain.Mineral")).Return(nil)
					return updateMock
				}(),
			},
			wantErr: false,
		},
		{
			name: "get failure",
			fields: fields{
				GetOP: func() repository.MineralGetByIDIFace {
					getMock := &mocks.MineralGetByIDIFace{}
					getMock.On("Get", mock.AnythingOfType("string")).Return(
						nil, errors.New("some err"))
					return getMock
				}(),
			},
			wantErr: true,
		},
		{
			name: "already same solid state",
			fields: fields{
				GetOP: func() repository.MineralGetByIDIFace {
					getMock := &mocks.MineralGetByIDIFace{}
					getMock.On("Get", mock.AnythingOfType("string")).Return(
						&domain.Mineral{State: domain.Solid}, nil)
					return getMock
				}(),
			},
			wantErr: true,
		},
		{
			name: "update fails",
			fields: fields{
				GetOP: func() repository.MineralGetByIDIFace {
					getMock := &mocks.MineralGetByIDIFace{}
					getMock.On("Get", mock.AnythingOfType("string")).Return(
						&domain.Mineral{State: domain.Liquid}, nil)
					return getMock
				}(),
				UpdateOP: func() repository.MineralUpdateStateIFace {
					updateMock := &mocks.MineralUpdateStateIFace{}
					updateMock.On("Update", mock.AnythingOfType("*domain.Mineral")).Return(errors.New("some err"))
					return updateMock
				}(),
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := &CondenseMineralUC{
				GetOP:    tt.fields.GetOP,
				UpdateOP: tt.fields.UpdateOP,
			}
			if err := m.Condense(tt.args.mineralID); (err != nil) != tt.wantErr {
				t.Errorf("Condense() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
