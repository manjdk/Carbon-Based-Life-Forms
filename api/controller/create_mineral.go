package controller

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gofrs/uuid"
	"github.com/manjdk/Carbon-Based-Life-Forms/api/domain"
	"github.com/manjdk/Carbon-Based-Life-Forms/custom_http"
	"github.com/manjdk/Carbon-Based-Life-Forms/error"
	"github.com/manjdk/Carbon-Based-Life-Forms/log"
)

func CreateMineral(httpClient custom_http.CallClientIFace, managerURL string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		traceID := log.NewTraceID()

		mineral := new(mineral)
		if err := json.NewDecoder(r.Body).Decode(mineral); err != nil {
			log.ErrorZ(traceID, err).Msg("Failed to decode create mineral request")
			custom_http.NewResponse(w, http.StatusBadRequest, error.NewErrorMessage(err))
			return
		}

		mineral.ID = uuid.Must(uuid.NewV4()).String()

		switch domain.NewMineralState(mineral.State) {
		case domain.Liquid, domain.Solid:
			mineral.Fractures = 0
		case domain.Fractured:
			if mineral.Fractures < 1 {
				mineral.Fractures = 1
			}
		default:
			err := fmt.Errorf("unsupported mineral state: %s", mineral.State)
			log.ErrorZ(traceID, err).
				Str("state", mineral.State).
				Msg("Unsupported mineral state")
			custom_http.NewResponse(w, http.StatusBadRequest, error.NewErrorMessage(err))
			return
		}

		requestBytes, err := json.Marshal(mineral.toDomain())
		if err != nil {
			log.ErrorZ(traceID, err).Msg("Failed to marshal mineral create request")
			custom_http.NewResponse(w, http.StatusBadRequest, error.NewErrorMessage(err))
			return
		}

		_, statusCode, err := httpClient.Call(traceID, http.MethodPost, managerURL, requestBytes)
		if err != nil {
			log.ErrorZ(traceID, err).Msg("Request to create mineral failed")
			custom_http.NewResponse(w, http.StatusFailedDependency, error.NewErrorMessage(err))
			return
		}

		if statusCode != http.StatusCreated {
			err := fmt.Errorf("wrong response code: %d", statusCode)
			log.ErrorZ(traceID, err).Int("code", statusCode).Msg("Wrong create mineral response code")
			custom_http.NewResponse(w, http.StatusFailedDependency, error.NewErrorMessage(err))
			return
		}

		custom_http.NewResponse(w, http.StatusCreated, mineral)
	}
}
