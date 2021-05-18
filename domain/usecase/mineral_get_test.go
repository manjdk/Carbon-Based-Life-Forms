package usecase

import (
	"errors"
	"reflect"
	"testing"

	"github.com/manjdk/Carbon-Based-Life-Forms/mocks"
	"github.com/stretchr/testify/mock"

	"github.com/manjdk/Carbon-Based-Life-Forms/domain"
	"github.com/manjdk/Carbon-Based-Life-Forms/repository"
)

func TestGetMineralUC_GetByID(t *testing.T) {
	type fields struct {
		GetOP repository.MineralGetByIDIFace
	}
	tests := []struct {
		name    string
		fields  fields
		want    *domain.Mineral
		wantErr bool
	}{
		{
			name: "success",
			fields: fields{
				GetOP: func() repository.MineralGetByIDIFace {
					getMock := &mocks.MineralGetByIDIFace{}
					getMock.On("Get", mock.AnythingOfType("string")).
						Return(&domain.Mineral{ClientID: "clientID"}, nil)
					return getMock
				}(),
			},
			want:    &domain.Mineral{ClientID: "clientID"},
			wantErr: false,
		},
		{
			name: "error",
			fields: fields{
				GetOP: func() repository.MineralGetByIDIFace {
					getMock := &mocks.MineralGetByIDIFace{}
					getMock.On("Get", mock.AnythingOfType("string")).
						Return(nil, errors.New("some err"))
					return getMock
				}(),
			},
			want:    nil,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := &GetMineralUC{
				GetOP: tt.fields.GetOP,
			}
			got, err := m.GetByID("mineralId")
			if (err != nil) != tt.wantErr {
				t.Errorf("GetByID() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetByID() got = %v, want %v", got, tt.want)
			}
		})
	}
}
