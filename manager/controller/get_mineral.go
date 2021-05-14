package controller

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/manjdk/Carbon-Based-Life-Forms/custom_http"
	"github.com/manjdk/Carbon-Based-Life-Forms/error"
	"github.com/manjdk/Carbon-Based-Life-Forms/log"
)

func GetMineralManager(httpClient custom_http.CallClientIFace, factoryURL string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		traceID := r.Header.Get(custom_http.HeaderXTraceID)

		mineralID, exists := mux.Vars(r)["mineralId"]
		if !exists || mineralID == "" {
			err := errors.New("url parameter not found: mineralId")
			log.ErrorZ(traceID, err).Msg("URL parameter mineralId not set in manager")
			custom_http.NewResponse(w, http.StatusBadRequest, error.NewErrorMessage(err))
			return
		}

		url := fmt.Sprintf("%s/%s", factoryURL, mineralID)
		responseBytes, statusCode, err := httpClient.Call(traceID, http.MethodGet, url, nil)
		if err != nil {
			log.ErrorZ(traceID, err).
				Str("url", url).
				Str("mineralID", mineralID).
				Msg("Error when calling factory on get mineral")
			custom_http.NewResponse(w, http.StatusFailedDependency, error.NewErrorMessage(err))
			return
		}

		if statusCode != http.StatusOK {
			err := fmt.Errorf("wrong response code: %d", statusCode)
			log.ErrorZ(traceID, err).
				Int("code", statusCode).
				Str("mineralID", mineralID).
				Msg("Wrong get mineral response code")
			custom_http.NewResponse(w, http.StatusFailedDependency, error.NewErrorMessage(err))
			return
		}

		mineral := new(mineral)
		if err := json.Unmarshal(responseBytes, mineral); err != nil {
			log.ErrorZ(traceID, err).
				Str("response", string(responseBytes)).
				Msg("Failed to unmarshal mineral response in manager")
			custom_http.NewResponse(w, http.StatusFailedDependency, error.NewErrorMessage(err))
			return
		}

		custom_http.NewResponse(w, http.StatusOK, mineral)
	}
}
