package handler

import (
	"fmt"
	"net/http"
	"my-go-webserver/internal/service"
)

type Handler struct {
	service service.Service
}

func NewHandler(s service.Service) *Handler {
	return &Handler{service: s}
}

func (h *Handler) Home(w http.ResponseWriter, r *http.Request) {
	message := h.service.GetWelcomeMessage()
	fmt.Fprintf(w, message)
}
