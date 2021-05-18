package controller

import (
	"bytes"
	"errors"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/manjdk/Carbon-Based-Life-Forms/mocks"
	"github.com/stretchr/testify/mock"

	"github.com/manjdk/Carbon-Based-Life-Forms/custom_http"
)

func TestCreateMineral(t *testing.T) {
	type args struct {
		httpClient custom_http.CallClientIFace
		body       io.Reader
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
						http.MethodPost,
						mock.AnythingOfType("string"),
						mock.AnythingOfType("map[string]string"),
						mock.AnythingOfType("[]uint8"),
					).Return([]byte(`{"name": "testName"}`), http.StatusCreated, nil)

					return m
				}(),
				body: bytes.NewReader([]byte(`{"clientId": "client", "state":"solid"}`)),
			},
			want: want{
				statusCode: http.StatusCreated,
			},
		},
		{
			name: "request body decode failure",
			args: args{
				body: bytes.NewReader([]byte(`{"clientId": "client", "state":"solid`)),
			},
			want: want{
				statusCode: http.StatusBadRequest,
			},
		},
		{
			name: "unsupported mineral state",
			args: args{
				body: bytes.NewReader([]byte(`{"clientId": "client", "state":"wrongState"}`)),
			},
			want: want{
				statusCode: http.StatusBadRequest,
			},
		},
		{
			name: "error from manager",
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
					).Return(nil, http.StatusFailedDependency, errors.New("some er"))

					return m
				}(),
				body: bytes.NewReader([]byte(`{"clientId": "client", "state":"solid"}`)),
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
						http.MethodPost,
						mock.AnythingOfType("string"),
						mock.AnythingOfType("map[string]string"),
						mock.AnythingOfType("[]uint8"),
					).Return([]byte(`{"name": "testName"}`), http.StatusOK, nil)

					return m
				}(),
				body: bytes.NewReader([]byte(`{"clientId": "client", "state":"solid"}`)),
			},
			want: want{
				statusCode: http.StatusFailedDependency,
			},
		},
		{
			name: "default fractures for liquid",
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
				body: bytes.NewReader([]byte(`{"clientId": "client", "state":"liquid"}`)),
			},
			want: want{
				statusCode: http.StatusCreated,
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

			got := CreateMineral(tt.args.httpClient, "testURL")
			got.ServeHTTP(rr, req)

			if status := rr.Code; status != tt.want.statusCode {
				t.Errorf("handler returned wrong status code: got %v want %v",
					status, tt.want.statusCode)
				return
			}
		})
	}
}
