package controller

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/manjdk/Carbon-Based-Life-Forms/custom_http"
	"github.com/manjdk/Carbon-Based-Life-Forms/error"
	"github.com/manjdk/Carbon-Based-Life-Forms/log"
)

func GetMinerals(httpClient custom_http.CallClientIFace, managerURL string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		traceID := log.NewTraceID()

		responseBytes, statusCode, err := httpClient.Call(traceID, http.MethodGet, managerURL, nil)
		if err != nil {
			log.ErrorZ(traceID, err).Msg("Error when calling manager on get all minerals")
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
				Msg("Failed to unmarshal get minerals response")
			custom_http.NewResponse(w, http.StatusFailedDependency, error.NewErrorMessage(err))
			return
		}

		custom_http.NewResponse(w, statusCode, minerals)
	}
}
