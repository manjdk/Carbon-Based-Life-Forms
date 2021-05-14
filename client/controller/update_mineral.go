package controller

import (
	"encoding/json"
	"net/http"

	"github.com/manjdk/Carbon-Based-Life-Forms/custom_http"
	"github.com/manjdk/Carbon-Based-Life-Forms/error"
	"github.com/manjdk/Carbon-Based-Life-Forms/log"
	"github.com/manjdk/Carbon-Based-Life-Forms/queue"
)

func UpdateMineral(publisher queue.Publisher) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		traceID := log.NewTraceID()

		updateReq := new(mineralStateUpdateRequest)
		if err := json.NewDecoder(r.Body).Decode(updateReq); err != nil {
			log.ErrorZ(traceID, err).Msg("Failed to decode update state request")
			custom_http.NewResponse(w, http.StatusBadRequest, error.NewErrorMessage(err))
			return
		}

		if err := publisher.Publish(updateReq.toQueueMessage(traceID)); err != nil {
			log.ErrorZ(traceID, err).
				Str("mineralID", updateReq.MineralID).
				Msg("Failed to publish action message to manager")
			custom_http.NewResponse(w, http.StatusInternalServerError, nil)
			return
		}

		custom_http.NewResponse(w, http.StatusOK, updateReq)
	}
}
