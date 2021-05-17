package controller

import (
	"net/http"

	"github.com/manjdk/Carbon-Based-Life-Forms/api/domain/usecase"
	"github.com/manjdk/Carbon-Based-Life-Forms/custom_http"
	"github.com/manjdk/Carbon-Based-Life-Forms/error"
	"github.com/manjdk/Carbon-Based-Life-Forms/log"
)

const (
	queryClientID = "clientId"
)

func GetMineralsFactory(getAllUC usecase.GetAllMineralUC) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		traceID := r.Header.Get(custom_http.HeaderXTraceID)

		minerals, err := getAllUC.GetAll(r.URL.Query().Get(queryClientID))
		if err != nil {
			log.ErrorZ(traceID, err).Msg("Failed to get minerals in factory")
			custom_http.NewResponse(w, http.StatusInternalServerError, error.NewErrorMessage(err))
			return
		}

		custom_http.NewResponse(w, http.StatusOK, minerals)
	}
}
