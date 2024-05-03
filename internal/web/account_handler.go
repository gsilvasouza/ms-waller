package web

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gsilvasouza/ms-waller/internal/usercase/create_account"
)

type WebAccountHandler struct {
	CreateAccountUseCase create_account.CreateAccountUseCase
}

func NewWebAccountHandler(createAccountUseCase create_account.CreateAccountUseCase) *WebAccountHandler {
	return &WebAccountHandler{
		CreateAccountUseCase: createAccountUseCase,
	}
}

func (h *WebAccountHandler) CreateAccount(w http.ResponseWriter, r *http.Request) {
	var dto create_account.CreateAccountInputDTO
	err := json.NewDecoder(r.Body).Decode(&dto)
	if err != nil {
		fmt.Println(err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	fmt.Println(dto)
	output, err := h.CreateAccountUseCase.Execute(dto)
	if err != nil {
		fmt.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(output)
	if err != nil {
		fmt.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusCreated)

}
