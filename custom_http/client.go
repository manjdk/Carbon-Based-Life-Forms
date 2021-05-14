package custom_http

import (
	"bytes"
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"

	"github.com/manjdk/Carbon-Based-Life-Forms/log"
)

const (
	headerKeyContentType  = "Content-Type"
	headerApplicationJSON = "application/json"
	HeaderXTraceID        = "x-trace-id"
)

type CallClientIFace interface {
	Call(traceID, method, url string, body []byte) ([]byte, int, error)
}

type HttpClient struct {
	client *http.Client
}

func NewHttpClient(client *http.Client) *HttpClient {
	return &HttpClient{
		client: client,
	}
}

func (h *HttpClient) Call(traceID, method, url string, body []byte) ([]byte, int, error) {
	req, err := http.NewRequest(method, url, bytes.NewReader(body))
	if err != nil {
		log.ErrorZ(traceID, err).
			Str("url", url).
			Str("body", string(body)).
			Msg("Failed to create a request")
		return nil, 0, err
	}

	addTraceIDToRequest(req, traceID)
	response, err := h.client.Do(req)
	if err != nil {
		log.ErrorZ(traceID, err).Msg("Failed to do a call")
		return nil, 0, err
	}
	defer response.Body.Close()

	if response == nil {
		err := errors.New("response is nil")
		log.ErrorZ(traceID, err).Msg("Response is nil")
		return nil, 0, err
	}

	responseBytes, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.ErrorZ(traceID, err).Msg("Failed to read response body")
		return nil, 0, err
	}

	return responseBytes, response.StatusCode, nil
}

func addTraceIDToRequest(req *http.Request, traceID string) {
	req.Header.Add(HeaderXTraceID, traceID)
}

func NewResponse(w http.ResponseWriter, status int, payload interface{}) {
	w.Header().Set(headerKeyContentType, headerApplicationJSON)
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(payload)
}
