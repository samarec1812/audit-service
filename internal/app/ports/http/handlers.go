package http

import (
	"net/http"

	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/render"
	"golang.org/x/exp/slog"

	"github.com/samarec1812/audit-service/internal/app/service"
	"github.com/samarec1812/audit-service/internal/pkg/status"
)

func saveEvent(log *slog.Logger, a service.App) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		const op = "handlers.analytics.saveEvent"

		log = log.With(
			slog.String("op", op),
			slog.String("request_id", middleware.GetReqID(r.Context())),
		)

		var reqBody map[string]any
		if err := render.DecodeJSON(r.Body, &reqBody); err != nil {
			log.Error("failed to decode request body", err)
			render.JSON(w, r, EventErrorResponse(status.GetStringStatusByCode(http.StatusBadRequest), err))

			return
		}

		headers := make(map[string]any)
		for key, value := range r.Header {
			headers[key] = value
		}

		err := a.SaveEvent(r.Context(), headers, reqBody)
		if err != nil {
			log.Error("error with save", err)
			render.JSON(w, r, EventErrorResponse(status.GetStringStatusByCode(http.StatusBadRequest), err))

			return
		}

		render.JSON(w, r, status.GetStringStatusByCode(http.StatusAccepted))
	}
}
