package api

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/jaimeiherrera/schmo_login_go/src/entity"
)

type Handlers struct {
	Components *Components
}

func NewHandlers(components Components) *Handlers {
	return &Handlers{
		Components: &components,
	}
}

func (h Handlers) Ping(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(
		`{"message": "pong"}`,
	))
}

func (h Handlers) CreateUser(w http.ResponseWriter, r *http.Request) {
	user := entity.User{}

	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(
			`{"message": "Invalid request"}`,
		))
		return
	}

	userCreated, err := h.Components.UserRepository.Create(user)
	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(
			fmt.Sprint(`{"message": "Error creating user: `, err, `"}`),
		))
		return
	}

	resp, err := json.Marshal(userCreated)
	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(
			fmt.Sprint(`{"message": "Error creating user: `, err, `"}`),
		))
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	w.Write(resp)
}
