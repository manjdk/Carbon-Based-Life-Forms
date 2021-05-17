package controller

import (
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/manjdk/Carbon-Based-Life-Forms/api/domain"
	"github.com/manjdk/Carbon-Based-Life-Forms/custom_http"
	"github.com/manjdk/Carbon-Based-Life-Forms/mocks"
	"github.com/stretchr/testify/mock"
)

func TestGetMineralsManager(t *testing.T) {
	wantBytes, _ := json.Marshal([]domain.Mineral{{Name: "tesName"}})

	type args struct {
		httpClient custom_http.CallClientIFace
		factoryURL string
	}
	type want struct {
		statusCode int
		body       []byte
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
						http.MethodGet,
						"testURL",
						make(map[string]string),
						[]byte(nil),
					).Return(wantBytes, http.StatusOK, nil)

					return m
				}(),
				factoryURL: "test",
			},
			want: want{
				statusCode: http.StatusOK,
				body:       wantBytes,
			},
		},
		{
			name: "wrong response status",
			args: args{
				httpClient: func() custom_http.CallClientIFace {
					m := &mocks.CallClientIFace{}
					m.On(
						"Call",
						mock.AnythingOfType("string"),
						http.MethodGet,
						"testURL",
						make(map[string]string),
						[]byte(nil),
					).Return(wantBytes, http.StatusFailedDependency, nil)

					return m
				}(),
				factoryURL: "test",
			},
			want: want{
				statusCode: http.StatusFailedDependency,
				body:       wantBytes,
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
						http.MethodGet,
						"testURL",
						make(map[string]string),
						[]byte(nil),
					).Return(nil, http.StatusFailedDependency, errors.New("some error"))

					return m
				}(),
				factoryURL: "test",
			},
			want: want{
				statusCode: http.StatusFailedDependency,
				body:       wantBytes,
			},
		},
		{
			name: "wrong response body",
			args: args{
				httpClient: func() custom_http.CallClientIFace {
					m := &mocks.CallClientIFace{}
					m.On(
						"Call",
						mock.AnythingOfType("string"),
						http.MethodGet,
						"testURL",
						make(map[string]string),
						[]byte(nil),
					).Return([]byte(`[{"id":"","clientId":"","name":"tesName","state":"","fractures":0]`), http.StatusOK, nil)

					return m
				}(),
				factoryURL: "test",
			},
			want: want{
				statusCode: http.StatusFailedDependency,
				body:       nil,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			req, err := http.NewRequest(http.MethodGet, "/minerals", nil)
			if err != nil {
				t.Fatal(err)
			}
			rr := httptest.NewRecorder()

			got := GetMineralsManager(tt.args.httpClient, "testURL")
			got.ServeHTTP(rr, req)

			if status := rr.Code; status != tt.want.statusCode {
				t.Errorf("handler returned wrong status code: got %v want %v",
					status, tt.want.statusCode)
			}
		})
	}
}
