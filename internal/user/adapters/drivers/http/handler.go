package http

import (
	"encoding/json"
	"net/http"

	driver_ports "github.com/agustin-carnevale/hexagonal-go/internal/user/ports/drivers"
)

// Handler exposes the HTTP interface to the outside world.
type Handler struct {
	service driver_ports.UserInputPort
}

// New creates a new HTTP handler.
func NewHandler(service driver_ports.UserInputPort) *Handler {
	return &Handler{service}
}

// RegisterRoutes attaches routes to a given mux.
func (h *Handler) RegisterRoutes(mux *http.ServeMux) {
	mux.HandleFunc("/user", h.getUser)
}

// getUser handles GET /user?id=123
func (h *Handler) getUser(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")
	if id == "" {
		http.Error(w, "missing id", http.StatusBadRequest)
		return
	}

	user, err := h.service.GetUser(id)
	if err != nil {
		http.Error(w, "user not found", http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(user)
}
