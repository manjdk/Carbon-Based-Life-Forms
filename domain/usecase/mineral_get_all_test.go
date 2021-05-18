package usecase

import (
	"errors"
	"reflect"
	"testing"

	"github.com/stretchr/testify/mock"

	"github.com/manjdk/Carbon-Based-Life-Forms/mocks"

	"github.com/manjdk/Carbon-Based-Life-Forms/domain"
	"github.com/manjdk/Carbon-Based-Life-Forms/repository"
)

func TestGetAllMineralUC_GetAll(t *testing.T) {
	type fields struct {
		GetAllOP        repository.MineralsGetIFace
		GetByClientIDOP repository.MineralsGetByClientIDIFace
	}
	type args struct {
		clientID string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    int
		wantErr bool
	}{
		{
			name: "success all",
			fields: fields{
				GetAllOP: func() repository.MineralsGetIFace {
					getAllOP := &mocks.MineralsGetIFace{}
					getAllOP.On("GetAll").
						Return([]domain.Mineral{{ClientID: "client1"}, {ClientID: "client2"}}, nil)
					return getAllOP
				}(),
			},
			want:    2,
			wantErr: false,
		},
		{
			name: "success by client ID",
			fields: fields{
				GetByClientIDOP: func() repository.MineralsGetByClientIDIFace {
					getAllOP := &mocks.MineralsGetByClientIDIFace{}
					getAllOP.On("GetByClientID", mock.AnythingOfType("string")).
						Return([]domain.Mineral{{ClientID: "client2"}, {ClientID: "client2"}}, nil)
					return getAllOP
				}(),
			},
			args:    args{clientID: "test"},
			want:    2,
			wantErr: false,
		},
		{
			name: "error all",
			fields: fields{
				GetAllOP: func() repository.MineralsGetIFace {
					getAllOP := &mocks.MineralsGetIFace{}
					getAllOP.On("GetAll").
						Return(nil, errors.New("some err"))
					return getAllOP
				}(),
			},
			want:    0,
			wantErr: true,
		},
		{
			name: "error by client ID",
			fields: fields{
				GetByClientIDOP: func() repository.MineralsGetByClientIDIFace {
					getAllOP := &mocks.MineralsGetByClientIDIFace{}
					getAllOP.On("GetByClientID", mock.AnythingOfType("string")).
						Return(nil, errors.New("some err"))
					return getAllOP
				}(),
			},
			args:    args{clientID: "test"},
			want:    0,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := &GetAllMineralUC{
				GetAllOP:        tt.fields.GetAllOP,
				GetByClientIDOP: tt.fields.GetByClientIDOP,
			}
			got, err := m.GetAll(tt.args.clientID)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetAll() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(len(got), tt.want) {
				t.Errorf("GetAll() got = %v, want %v", len(got), tt.want)
			}
		})
	}
}
