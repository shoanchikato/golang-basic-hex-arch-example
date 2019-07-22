package toy

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

type ToyHandler interface {
	Get(w http.ResponseWriter, r *http.Request)
	GetByID(w http.ResponseWriter, r *http.Request)
	Create(w http.ResponseWriter, r *http.Request)
}

func NewToyHandler(toyService Service) ToyHandler {
	return &toyHandler{
		toyService,
	}
}

type toyHandler struct {
	toyService Service
}

func response(w http.ResponseWriter, data interface{}) {
	response, _ := json.Marshal(data)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	_, _ = w.Write(response)
}

func (h *toyHandler) Get(w http.ResponseWriter, r *http.Request) {
	toys, err := h.toyService.FindAllToys()
	if err != nil {
		fmt.Printf("error when getting toys %v\n", err)
	}
	response(w, toys)
}

func (h *toyHandler) GetByID(w http.ResponseWriter, r *http.Request) {
	param := mux.Vars(r)
	id := param["id"]
	toy, err := h.toyService.FindToyByID(id)
	if err != nil {
		fmt.Printf("error when getting one toy %v\n", err)
	}
	response(w, toy)
}

func (h *toyHandler) Create(w http.ResponseWriter, r *http.Request) {
	var toy Toy
	decoder := json.NewDecoder(r.Body)
	_ = decoder.Decode(&toy)
	err := h.toyService.CreateToy(&toy)
	if err != nil {
		fmt.Printf("error when creating toy %v\n", err)
	}
	response(w, toy)
}
