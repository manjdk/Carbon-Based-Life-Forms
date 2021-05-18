package controller

import (
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/manjdk/Carbon-Based-Life-Forms/domain/usecase"

	"github.com/gorilla/mux"
	"github.com/manjdk/Carbon-Based-Life-Forms/mocks"
	"github.com/stretchr/testify/mock"
)

func TestDeleteMineralFactory(t *testing.T) {
	type args struct {
		deleteUC usecase.DeleteMineralUC
		urlVars  map[string]string
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
				deleteUC: func() usecase.DeleteMineralUC {
					deleteMock := &mocks.MineralDeleteIFace{}
					deleteMock.On("Delete", mock.AnythingOfType("string")).Return(nil)
					return usecase.NewDeleteMineralUC(deleteMock)
				}(),
				urlVars: map[string]string{"mineralId": "testID"},
			},
			want: want{
				statusCode: http.StatusNoContent,
			},
		},
		{
			name: "id not set",
			args: args{
				deleteUC: func() usecase.DeleteMineralUC {
					deleteMock := &mocks.MineralDeleteIFace{}
					return usecase.NewDeleteMineralUC(deleteMock)
				}(),
			},
			want: want{
				statusCode: http.StatusBadRequest,
			},
		},
		{
			name: "error on delete",
			args: args{
				deleteUC: func() usecase.DeleteMineralUC {
					deleteMock := &mocks.MineralDeleteIFace{}
					deleteMock.On("Delete", mock.AnythingOfType("string")).Return(errors.New("some err"))
					return usecase.NewDeleteMineralUC(deleteMock)
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
			req, err := http.NewRequest(http.MethodDelete, "/minerals/testID", nil)
			if err != nil {
				t.Fatal(err)
			}

			if tt.args.urlVars != nil {
				req = mux.SetURLVars(req, tt.args.urlVars)
			}

			got := DeleteMineralFactory(tt.args.deleteUC)
			got.ServeHTTP(rr, req)

			if status := rr.Code; status != tt.want.statusCode {
				t.Errorf("handler returned wrong status code: got %v want %v",
					status, tt.want.statusCode)
			}
		})
	}
}
