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

	"github.com/manjdk/Carbon-Based-Life-Forms/queue"
)

func TestUpdateMineral(t *testing.T) {
	type args struct {
		publisher queue.Publisher
		body      io.Reader
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
				publisher: func() queue.Publisher {
					pubMock := &mocks.Publisher{}
					pubMock.On("Publish", mock.AnythingOfType("*domain.QueueMessage")).
						Return(nil)
					return pubMock
				}(),
				body: bytes.NewReader([]byte(`{"mineralId": "testID", "action":"melt"}`)),
			},
			want: want{
				statusCode: http.StatusOK,
			},
		},
		{
			name: "decode error",
			args: args{
				body: bytes.NewReader([]byte(`{"mineralId": "testID", "action":"melt`)),
			},
			want: want{
				statusCode: http.StatusBadRequest,
			},
		},
		{
			name: "failed publish",
			args: args{
				publisher: func() queue.Publisher {
					pubMock := &mocks.Publisher{}
					pubMock.On("Publish", mock.AnythingOfType("*domain.QueueMessage")).
						Return(errors.New("some err"))
					return pubMock
				}(),
				body: bytes.NewReader([]byte(`{"mineralId": "testID", "action":"melt"}`)),
			},
			want: want{
				statusCode: http.StatusInternalServerError,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			rr := httptest.NewRecorder()
			req, err := http.NewRequest(http.MethodPut, "/minerals", tt.args.body)
			if err != nil {
				t.Fatal(err)
			}

			got := UpdateMineral(tt.args.publisher)
			got.ServeHTTP(rr, req)

			if status := rr.Code; status != tt.want.statusCode {
				t.Errorf("handler returned wrong status code: got %v want %v",
					status, tt.want.statusCode)
			}
		})
	}
}
