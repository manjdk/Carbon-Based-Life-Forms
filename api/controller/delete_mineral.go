package controller

import (
	"errors"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/manjdk/Carbon-Based-Life-Forms/custom_http"
	"github.com/manjdk/Carbon-Based-Life-Forms/error"
	"github.com/manjdk/Carbon-Based-Life-Forms/log"
)

func DeleteMineral(httpClient custom_http.CallClientIFace, managerURL string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		traceID := log.NewTraceID()

		mineralID, exists := mux.Vars(r)["mineralId"]
		if !exists || mineralID == "" {
			err := errors.New("url parameter not found: mineralId")
			log.ErrorZ(traceID, err).Msg("URL parameter mineralId not set")
			custom_http.NewResponse(w, http.StatusBadRequest, error.NewErrorMessage(err))
			return
		}

		url := fmt.Sprintf("%s/%s", managerURL, mineralID)
		_, statusCode, err := httpClient.Call(traceID, http.MethodDelete, url, nil)
		if err != nil {
			log.ErrorZ(traceID, err).
				Str("url", url).
				Str("mineralID", mineralID).
				Msg("Error when calling manager on delete mineral")
			custom_http.NewResponse(w, http.StatusFailedDependency, error.NewErrorMessage(err))
			return
		}

		if statusCode != http.StatusNoContent {
			err := fmt.Errorf("wrong response code: %d", statusCode)
			log.ErrorZ(traceID, err).
				Int("code", statusCode).
				Str("mineralID", mineralID).
				Msg("Wrong delete mineral response code")
			custom_http.NewResponse(w, http.StatusFailedDependency, error.NewErrorMessage(err))
			return
		}

		custom_http.NewResponse(w, http.StatusNoContent, nil)
	}
}
