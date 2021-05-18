package controller

import (
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/manjdk/Carbon-Based-Life-Forms/custom_http"
	"github.com/manjdk/Carbon-Based-Life-Forms/error"
	"github.com/manjdk/Carbon-Based-Life-Forms/log"
)

func CreateMineralManager(httpClient custom_http.CallClientIFace, factoryURL string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		traceID := r.Header.Get(custom_http.HeaderXTraceID)

		requestBytes, err := ioutil.ReadAll(r.Body)
		if err != nil {
			log.ErrorZ(traceID, err).Msg("Failed to read crate mineral request body in manager")
			custom_http.NewResponse(w, http.StatusBadRequest, error.NewErrorMessage(err))
			return
		}

		responseBytes, statusCode, err := httpClient.Call(traceID, http.MethodPost, factoryURL, nil, requestBytes)
		if err != nil {
			log.ErrorZ(traceID, err).Msg("Request to create mineral failed in manager")
			custom_http.NewResponse(w, http.StatusFailedDependency, error.NewErrorMessage(err))
			return
		}

		if statusCode != http.StatusCreated {
			err := fmt.Errorf("wrong response code: %d", statusCode)
			log.ErrorZ(traceID, err).Int("code", statusCode).Msg("Wrong create mineral response code")
			custom_http.NewResponse(w, http.StatusFailedDependency, error.NewErrorMessage(err))
			return
		}

		custom_http.NewResponse(w, http.StatusCreated, responseBytes)
	}
}
