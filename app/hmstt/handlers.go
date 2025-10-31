package hmstt

import (
	"net/http"

	"github.com/nurhudajoantama/stthmauto/app/server"
	"github.com/nurhudajoantama/stthmauto/internal/response"
)

type hmsttHandler struct {
	service *hmsttService
}

func RegisterHandlers(s *server.Server, svc *hmsttService) {
	h := &hmsttHandler{
		service: svc,
	}
	srv := s.GetRouter()

	// hmsttGroup := srv.PathPrefix("/hmstt/").Subrouter()
	// hmsttGroup.HandleFunc("/getstate", h.GetState)

	hmsttApiGroup := srv.PathPrefix("/api/hmstt/").Subrouter()
	hmsttApiGroup.HandleFunc("/getstate", h.ApiGetState).Methods("GET")
}

func (h *hmsttHandler) ApiGetState(w http.ResponseWriter, r *http.Request) {
	states := r.URL.Query()["states"]

	result, err := h.service.GetState(r.Context(), states...)
	if err != nil {
		response.ErrorResponse(w, http.StatusInternalServerError, "failed to get state", err)
		return
	}
	// Return the result as JSON
	response.SuccessResponse(w, result)
}
