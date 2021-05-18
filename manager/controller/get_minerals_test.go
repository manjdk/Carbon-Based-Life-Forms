package controller

import (
	"errors"
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"

	"github.com/manjdk/Carbon-Based-Life-Forms/custom_http"
	"github.com/manjdk/Carbon-Based-Life-Forms/mocks"
	"github.com/stretchr/testify/mock"
)

func TestGetMineralsManager(t *testing.T) {
	type args struct {
		httpClient custom_http.CallClientIFace
		urlValues  url.Values
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
				httpClient: func() custom_http.CallClientIFace {
					m := &mocks.CallClientIFace{}
					m.On(
						"Call",
						mock.AnythingOfType("string"),
						http.MethodGet,
						mock.AnythingOfType("string"),
						mock.AnythingOfType("map[string]string"),
						[]byte(nil),
					).Return([]byte(`[{"name": "tesName"}]`), http.StatusOK, nil)

					return m
				}(),
			},
			want: want{
				statusCode: http.StatusOK,
			},
		},
		{
			name: "success clientID given",
			args: args{
				httpClient: func() custom_http.CallClientIFace {
					m := &mocks.CallClientIFace{}
					m.On(
						"Call",
						mock.AnythingOfType("string"),
						http.MethodGet,
						mock.AnythingOfType("string"),
						mock.AnythingOfType("map[string]string"),
						[]byte(nil),
					).Return([]byte(`[{"name": "tesName"}]`), http.StatusOK, nil)

					return m
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
				httpClient: func() custom_http.CallClientIFace {
					m := &mocks.CallClientIFace{}
					m.On(
						"Call",
						mock.AnythingOfType("string"),
						http.MethodGet,
						mock.AnythingOfType("string"),
						mock.AnythingOfType("map[string]string"),
						[]byte(nil),
					).Return(nil, http.StatusFailedDependency, errors.New("some err"))

					return m
				}(),
			},
			want: want{
				statusCode: http.StatusFailedDependency,
			},
		},
		{
			name: "failed with clientID given",
			args: args{
				httpClient: func() custom_http.CallClientIFace {
					m := &mocks.CallClientIFace{}
					m.On(
						"Call",
						mock.AnythingOfType("string"),
						http.MethodGet,
						mock.AnythingOfType("string"),
						mock.AnythingOfType("map[string]string"),
						[]byte(nil),
					).Return(nil, http.StatusFailedDependency, errors.New("some err"))

					return m
				}(),
				urlValues: func() url.Values {
					vals := url.Values{}
					vals.Add("clientId", "clientId")
					return vals
				}(),
			},
			want: want{
				statusCode: http.StatusFailedDependency,
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

			got := GetMineralsManager(tt.args.httpClient, "testID")
			got.ServeHTTP(rr, req)

			if status := rr.Code; status != tt.want.statusCode {
				t.Errorf("handler returned wrong status code: got %v want %v",
					status, tt.want.statusCode)
			}
		})
	}
}
