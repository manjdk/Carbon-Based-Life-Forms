package controller

import (
	"bytes"
	"errors"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/manjdk/Carbon-Based-Life-Forms/custom_http"
	"github.com/manjdk/Carbon-Based-Life-Forms/mocks"
	"github.com/stretchr/testify/mock"
)

func TestCreateMineralManager(t *testing.T) {
	type args struct {
		httpClient custom_http.CallClientIFace
		factoryURL string
		body       io.Reader
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
						http.MethodPost,
						mock.AnythingOfType("string"),
						mock.AnythingOfType("map[string]string"),
						mock.AnythingOfType("[]uint8"),
					).Return([]byte(`{"name": "testName"}`), http.StatusCreated, nil)

					return m
				}(),
				factoryURL: "test",
				body:       bytes.NewReader([]byte(`{"name": "testName"}`)),
			},
			want: want{
				statusCode: http.StatusCreated,
				body:       []byte(`{"name":"testName"}`),
			},
		},
		{
			name: "error from response",
			args: args{
				httpClient: func() custom_http.CallClientIFace {
					m := &mocks.CallClientIFace{}
					m.On(
						"Call",
						mock.AnythingOfType("string"),
						http.MethodPost,
						mock.AnythingOfType("string"),
						mock.AnythingOfType("map[string]string"),
						mock.AnythingOfType("[]uint8"),
					).Return(nil, http.StatusCreated, errors.New("some error"))

					return m
				}(),
				factoryURL: "test",
				body:       bytes.NewReader([]byte(`{"name": "testName"}`)),
			},
			want: want{
				statusCode: http.StatusFailedDependency,
			},
		},
		{
			name: "wrong status",
			args: args{
				httpClient: func() custom_http.CallClientIFace {
					m := &mocks.CallClientIFace{}
					m.On(
						"Call",
						mock.AnythingOfType("string"),
						http.MethodPost,
						mock.AnythingOfType("string"),
						mock.AnythingOfType("map[string]string"),
						mock.AnythingOfType("[]uint8"),
					).Return([]byte(`{"name": "testName"}`), http.StatusOK, nil)

					return m
				}(),
				factoryURL: "test",
				body:       bytes.NewReader([]byte(`{"name": "testName"}`)),
			},
			want: want{
				statusCode: http.StatusFailedDependency,
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

			got := CreateMineralManager(tt.args.httpClient, "testURL")
			got.ServeHTTP(rr, req)

			if status := rr.Code; status != tt.want.statusCode {
				t.Errorf("handler returned wrong status code: got %v want %v",
					status, tt.want.statusCode)
			}
		})
	}
}
