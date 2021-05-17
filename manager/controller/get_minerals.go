package controller

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/manjdk/Carbon-Based-Life-Forms/custom_http"
	"github.com/manjdk/Carbon-Based-Life-Forms/error"
	"github.com/manjdk/Carbon-Based-Life-Forms/log"
)

const (
	queryClientID = "clientId"
)

func GetMineralsManager(httpClient custom_http.CallClientIFace, factoryURL string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		traceID := r.Header.Get(custom_http.HeaderXTraceID)

		queryParams := make(map[string]string)
		if clientID := r.URL.Query().Get(queryClientID); clientID != "" {
			queryParams[queryClientID] = clientID
		}

		responseBytes, statusCode, err := httpClient.Call(traceID, http.MethodGet, factoryURL, queryParams, nil)
		if err != nil {
			log.ErrorZ(traceID, err).Msg("Error when calling factory on get all minerals")
			custom_http.NewResponse(w, http.StatusFailedDependency, error.NewErrorMessage(err))
			return
		}

		if statusCode != http.StatusOK {
			err := fmt.Errorf("wrong response code: %d", statusCode)
			log.ErrorZ(traceID, err).
				Int("code", statusCode).
				Msg("Wrong response status code")
			custom_http.NewResponse(w, http.StatusFailedDependency, error.NewErrorMessage(err))
			return
		}

		minerals := make([]mineral, 0)
		if err := json.Unmarshal(responseBytes, &minerals); err != nil {
			log.ErrorZ(traceID, err).
				Str("response", string(responseBytes)).
				Msg("Failed to unmarshal get minerals response in manager")
			custom_http.NewResponse(w, http.StatusFailedDependency, error.NewErrorMessage(err))
			return
		}

		custom_http.NewResponse(w, http.StatusOK, minerals)
	}
}
