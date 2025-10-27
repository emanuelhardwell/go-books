package transport

import (
	"encoding/json"
	"go-books/internal/model"
	"go-books/internal/service"
	"net/http"
	"strconv"
	"strings"
)

type BookHandler struct {
	service *service.Service
}

func New(s *service.Service) *BookHandler {
	return &BookHandler{service: s}
}

func (h *BookHandler) HandleBooksGetPost(w http.ResponseWriter, r *http.Request) {

	switch r.Method {
	case http.MethodGet:
		books, err := h.service.GetAllCtrl()
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-type", "application/json")
		json.NewEncoder(w).Encode(books)

	case http.MethodPost:
		var book *model.Book
		err := json.NewDecoder(r.Body).Decode(&book)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		res, err := h.service.CreateCtrl(book)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusCreated)
		w.Header().Set("Content-type", "application/json")
		json.NewEncoder(w).Encode(res)

	default:
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		return
	}
}

// TODO: get y delete
func (h *BookHandler) HandlerBookGetDeleteById(w http.ResponseWriter, r *http.Request) {
	idString := strings.TrimPrefix(r.URL.Path, "/books/")

	id, err := strconv.Atoi(idString)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	switch r.Method {
	case http.MethodGet:
		resp, err := h.service.GetByIdCtrl(id)
		if err != nil {
			http.Error(w, err.Error(), http.StatusNotFound)
			return
		}

		w.Header().Set("Content-type", "application/json")
		json.NewEncoder(w).Encode(resp)

	case http.MethodPut:
		var book *model.Book

		if err := json.NewDecoder(r.Body).Decode(&book); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		resp, err := h.service.UpdateCtrl(id, book)
		if err != nil {
			http.Error(w, err.Error(), http.StatusNotFound)
			return
		}

		w.Header().Set("Content-type", "application/json")
		json.NewEncoder(w).Encode(resp)

	case http.MethodDelete:
		err := h.service.DeleteCtrl(id)
		if err != nil {
			http.Error(w, err.Error(), http.StatusNotFound)
			return
		}

		w.WriteHeader(http.StatusNoContent)

	default:
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		return
	}
}
