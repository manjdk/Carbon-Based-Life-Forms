package controller

import (
	"errors"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/manjdk/Carbon-Based-Life-Forms/api/domain/usecase"
	"github.com/manjdk/Carbon-Based-Life-Forms/custom_http"
	"github.com/manjdk/Carbon-Based-Life-Forms/error"
	"github.com/manjdk/Carbon-Based-Life-Forms/log"
)

func GetMineralFactory(getUC usecase.GetMineralUC) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		traceID := r.Header.Get(custom_http.HeaderXTraceID)

		mineralID, exists := mux.Vars(r)["mineralId"]
		if !exists || mineralID == "" {
			err := errors.New("url parameter not found: mineralId")
			log.ErrorZ(traceID, err).Msg("URL parameter mineralId not set in factory")
			custom_http.NewResponse(w, http.StatusBadRequest, error.NewErrorMessage(err))
			return
		}

		mineral, err := getUC.GetByID(mineralID)
		if err != nil {
			log.ErrorZ(traceID, err).
				Str("mineralID", mineralID).
				Msg("Failed to get mineral by ID")
			custom_http.NewResponse(w, http.StatusInternalServerError, error.NewErrorMessage(err))
			return
		}

		custom_http.NewResponse(w, http.StatusOK, mineral)
	}
}
