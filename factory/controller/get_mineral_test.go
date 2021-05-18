package controller

import (
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/manjdk/Carbon-Based-Life-Forms/domain"

	"github.com/manjdk/Carbon-Based-Life-Forms/domain/usecase"

	"github.com/gorilla/mux"
	"github.com/manjdk/Carbon-Based-Life-Forms/mocks"
	"github.com/stretchr/testify/mock"
)

func TestGetMineralFactory(t *testing.T) {
	type args struct {
		getUC   usecase.GetMineralUC
		urlVars map[string]string
	}
	type want struct {
		statusCode int
	}
	tests := []struct {
		name string
		args args
		want want
	}{
		{
			name: "success",
			args: args{
				getUC: func() usecase.GetMineralUC {
					getMock := &mocks.MineralGetByIDIFace{}
					getMock.On("Get", mock.AnythingOfType("string")).Return(&domain.Mineral{ClientID: "testID"}, nil)
					return usecase.NewGetMineralUC(getMock)
				}(),
				urlVars: map[string]string{"mineralId": "testID"},
			},
			want: want{
				statusCode: http.StatusOK,
			},
		},
		{
			name: "id not set",
			args: args{
				getUC: func() usecase.GetMineralUC {
					getMock := &mocks.MineralGetByIDIFace{}
					return usecase.NewGetMineralUC(getMock)
				}(),
			},
			want: want{
				statusCode: http.StatusBadRequest,
			},
		},
		{
			name: "error on get",
			args: args{
				getUC: func() usecase.GetMineralUC {
					getMock := &mocks.MineralGetByIDIFace{}
					getMock.On("Get", mock.AnythingOfType("string")).Return(nil, errors.New("some err"))
					return usecase.NewGetMineralUC(getMock)
				}(),
				urlVars: map[string]string{"mineralId": "testID"},
			},
			want: want{
				statusCode: http.StatusInternalServerError,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			rr := httptest.NewRecorder()
			req, err := http.NewRequest(http.MethodGet, "/minerals/testID", nil)
			if err != nil {
				t.Fatal(err)
			}

			if tt.args.urlVars != nil {
				req = mux.SetURLVars(req, tt.args.urlVars)
			}

			got := GetMineralFactory(tt.args.getUC)
			got.ServeHTTP(rr, req)

			if status := rr.Code; status != tt.want.statusCode {
				t.Errorf("handler returned wrong status code: got %v want %v",
					status, tt.want.statusCode)
			}
		})
	}
}
