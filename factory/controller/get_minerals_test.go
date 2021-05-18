package controller

import (
	"errors"
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"

	"github.com/stretchr/testify/mock"

	"github.com/manjdk/Carbon-Based-Life-Forms/domain"

	"github.com/manjdk/Carbon-Based-Life-Forms/domain/usecase"

	"github.com/manjdk/Carbon-Based-Life-Forms/mocks"
)

func TestGetMineralsFactory(t *testing.T) {
	type args struct {
		getUC     usecase.GetAllMineralUC
		urlValues url.Values
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
			name: "success empty clientID",
			args: args{
				getUC: func() usecase.GetAllMineralUC {
					getMock := &mocks.MineralsGetIFace{}
					getByClientMock := &mocks.MineralsGetByClientIDIFace{}
					getMock.On("GetAll").Return([]domain.Mineral{}, nil)
					return usecase.NewGetAllMineralUC(getMock, getByClientMock)
				}(),
			},
			want: want{
				statusCode: http.StatusOK,
			},
		},
		{
			name: "success clientID given",
			args: args{
				getUC: func() usecase.GetAllMineralUC {
					getMock := &mocks.MineralsGetIFace{}
					getByClientMock := &mocks.MineralsGetByClientIDIFace{}
					getByClientMock.On("GetByClientID", mock.AnythingOfType("string")).Return([]domain.Mineral{}, nil)
					return usecase.NewGetAllMineralUC(getMock, getByClientMock)
				}(),
				urlValues: func() url.Values {
					vals := url.Values{}
					vals.Add("clientId", "clientId")
					return vals
				}(),
			},
			want: want{
				statusCode: http.StatusOK,
			},
		},
		{
			name: "failed with empty clientID",
			args: args{
				getUC: func() usecase.GetAllMineralUC {
					getMock := &mocks.MineralsGetIFace{}
					getByClientMock := &mocks.MineralsGetByClientIDIFace{}
					getMock.On("GetAll").Return(nil, errors.New("some err"))
					return usecase.NewGetAllMineralUC(getMock, getByClientMock)
				}(),
			},
			want: want{
				statusCode: http.StatusInternalServerError,
			},
		},
		{
			name: "failed with clientID given",
			args: args{
				getUC: func() usecase.GetAllMineralUC {
					getMock := &mocks.MineralsGetIFace{}
					getByClientMock := &mocks.MineralsGetByClientIDIFace{}
					getByClientMock.On("GetByClientID", mock.AnythingOfType("string")).Return(nil, errors.New("some err"))
					return usecase.NewGetAllMineralUC(getMock, getByClientMock)
				}(),
				urlValues: func() url.Values {
					vals := url.Values{}
					vals.Add("clientId", "clientId")
					return vals
				}(),
			},
			want: want{
				statusCode: http.StatusInternalServerError,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			rr := httptest.NewRecorder()
			req, err := http.NewRequest(http.MethodGet, "/minerals", nil)
			if err != nil {
				t.Fatal(err)
			}

			if tt.args.urlValues != nil {
				req.URL.RawQuery = tt.args.urlValues.Encode()
			}

			got := GetMineralsFactory(tt.args.getUC)
			got.ServeHTTP(rr, req)

			if status := rr.Code; status != tt.want.statusCode {
				t.Errorf("handler returned wrong status code: got %v want %v",
					status, tt.want.statusCode)
			}
		})
	}
}
