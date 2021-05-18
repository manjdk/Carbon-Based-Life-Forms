package controller

import (
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gorilla/mux"
	"github.com/manjdk/Carbon-Based-Life-Forms/custom_http"
	"github.com/manjdk/Carbon-Based-Life-Forms/mocks"
	"github.com/stretchr/testify/mock"
)

func TestDeleteMineralManager(t *testing.T) {
	type args struct {
		httpClient custom_http.CallClientIFace
		factoryURL string
		urlVars    map[string]string
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
				httpClient: func() custom_http.CallClientIFace {
					m := &mocks.CallClientIFace{}
					m.On(
						"Call",
						mock.AnythingOfType("string"),
						http.MethodDelete,
						mock.AnythingOfType("string"),
						mock.AnythingOfType("map[string]string"),
						[]byte(nil),
					).Return(nil, http.StatusNoContent, nil)

					return m
				}(),
				factoryURL: "test",
				urlVars:    map[string]string{"mineralId": "testID"},
			},
			want: want{
				statusCode: http.StatusNoContent,
			},
		},
		{
			name: "no mineral ID passed",
			args: args{
				httpClient: func() custom_http.CallClientIFace {
					return &mocks.CallClientIFace{}
				}(),
				factoryURL: "test",
			},
			want: want{
				statusCode: http.StatusBadRequest,
			},
		},
		{
			name: "error response",
			args: args{
				httpClient: func() custom_http.CallClientIFace {
					m := &mocks.CallClientIFace{}
					m.On(
						"Call",
						mock.AnythingOfType("string"),
						http.MethodDelete,
						mock.AnythingOfType("string"),
						mock.AnythingOfType("map[string]string"),
						[]byte(nil),
					).Return([]byte(nil), http.StatusTeapot, errors.New("some error"))

					return m
				}(),
				factoryURL: "test",
				urlVars:    map[string]string{"mineralId": "testID"},
			},
			want: want{
				statusCode: http.StatusFailedDependency,
			},
		},
		{
			name: "wrong response code",
			args: args{
				httpClient: func() custom_http.CallClientIFace {
					m := &mocks.CallClientIFace{}
					m.On(
						"Call",
						mock.AnythingOfType("string"),
						http.MethodDelete,
						mock.AnythingOfType("string"),
						mock.AnythingOfType("map[string]string"),
						[]byte(nil),
					).Return([]byte(nil), http.StatusTeapot, nil)

					return m
				}(),
				factoryURL: "test",
				urlVars:    map[string]string{"mineralId": "testID"},
			},
			want: want{
				statusCode: http.StatusFailedDependency,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			req, err := http.NewRequest(http.MethodDelete, "/minerals", nil)
			if err != nil {
				t.Fatal(err)
			}

			if tt.args.urlVars != nil {
				req = mux.SetURLVars(req, tt.args.urlVars)
			}

			rr := httptest.NewRecorder()

			got := DeleteMineralManager(tt.args.httpClient, "testURL")
			got.ServeHTTP(rr, req)

			if status := rr.Code; status != tt.want.statusCode {
				t.Errorf("handler returned wrong status code: got %v want %v",
					status, tt.want.statusCode)
			}
		})
	}
}
