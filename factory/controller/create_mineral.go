package controller

import (
	"encoding/json"
	"net/http"

	"github.com/manjdk/Carbon-Based-Life-Forms/api/domain"
	"github.com/manjdk/Carbon-Based-Life-Forms/api/domain/usecase"
	"github.com/manjdk/Carbon-Based-Life-Forms/custom_http"
	"github.com/manjdk/Carbon-Based-Life-Forms/error"
	"github.com/manjdk/Carbon-Based-Life-Forms/log"
)

func CreateMineralFactory(createUC usecase.CreateMineralUC) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		traceID := r.Header.Get(custom_http.HeaderXTraceID)

		mineral := new(domain.Mineral)
		if err := json.NewDecoder(r.Body).Decode(mineral); err != nil {
			log.ErrorZ(traceID, err).Msg("Failed to decode mineral create request in factory")
			custom_http.NewResponse(w, http.StatusBadRequest, error.NewErrorMessage(err))
			return
		}

		if err := createUC.Create(mineral); err != nil {
			log.ErrorZ(traceID, err).Msg("Failed to create mineral")
			custom_http.NewResponse(w, http.StatusInternalServerError, error.NewErrorMessage(err))
			return
		}

		custom_http.NewResponse(w, http.StatusCreated, nil)
	}
}
