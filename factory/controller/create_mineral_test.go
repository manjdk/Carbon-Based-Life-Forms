package controller

import (
	"bytes"
	"errors"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/manjdk/Carbon-Based-Life-Forms/domain/usecase"
	"github.com/manjdk/Carbon-Based-Life-Forms/mocks"
	"github.com/stretchr/testify/mock"
)

func TestCreateMineralFactory(t *testing.T) {
	type args struct {
		createUC usecase.CreateMineralUC
		body     io.Reader
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
				createUC: func() usecase.CreateMineralUC {
					createMock := &mocks.MineralCreateIFace{}
					createMock.On("Create", mock.AnythingOfType("*domain.Mineral")).Return(nil)
					return usecase.NewCreateMineralUC(createMock)
				}(),
				body: bytes.NewReader([]byte(`{"clientId": "testClient"}`)),
			},
			want: want{
				statusCode: http.StatusCreated,
			},
		},
		{
			name: "error on create",
			args: args{
				createUC: func() usecase.CreateMineralUC {
					createMock := &mocks.MineralCreateIFace{}
					createMock.On("Create", mock.AnythingOfType("*domain.Mineral")).Return(errors.New("some error"))
					return usecase.NewCreateMineralUC(createMock)
				}(),
				body: bytes.NewReader([]byte(`{"clientId": "testClient"}`)),
			},
			want: want{
				statusCode: http.StatusInternalServerError,
			},
		},
		{
			name: "decode failure",
			args: args{
				createUC: func() usecase.CreateMineralUC {
					createMock := &mocks.MineralCreateIFace{}
					return usecase.NewCreateMineralUC(createMock)
				}(),
				body: bytes.NewReader([]byte(`{"clientId": "testClien}`)),
			},
			want: want{
				statusCode: http.StatusBadRequest,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			rr := httptest.NewRecorder()
			req, err := http.NewRequest(http.MethodPost, "/minerals", tt.args.body)
			if err != nil {
				t.Fatal(err)
			}

			got := CreateMineralFactory(tt.args.createUC)
			got.ServeHTTP(rr, req)

			if status := rr.Code; status != tt.want.statusCode {
				t.Errorf("handler returned wrong status code: got %v want %v",
					status, tt.want.statusCode)
			}
		})
	}
}
