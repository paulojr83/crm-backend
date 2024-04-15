package web

import (
	"encoding/json"
	"github.com/paulojr83/crm-backend/internal/entity"
	"github.com/paulojr83/crm-backend/internal/usecase"
	"net/http"
)

type WebCustomerHandler struct {
	CustomerRepository entity.CustomerRepositoryInterface
}

func NewWebCustomerHandler(
	CustomerRepository entity.CustomerRepositoryInterface,
) *WebCustomerHandler {
	return &WebCustomerHandler{
		CustomerRepository: CustomerRepository,
	}
}

func (h *WebCustomerHandler) AddCustomer(w http.ResponseWriter, r *http.Request) {
	var dto usecase.CustomerInputDTO
	err := json.NewDecoder(r.Body).Decode(&dto)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	createOrder := usecase.NewCreateCustomerUseCase(h.CustomerRepository)
	output, err := createOrder.Execute(dto)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	err = json.NewEncoder(w).Encode(output)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func (h *WebCustomerHandler) GetCustomers(w http.ResponseWriter, r *http.Request) {
	listOrderUseCase := usecase.NewListCustomerUseCase(h.CustomerRepository)
	output, err := listOrderUseCase.Execute()

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	jsonResponse, err := json.Marshal(output)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(jsonResponse)
}

func (h *WebCustomerHandler) GetCustomer(w http.ResponseWriter, r *http.Request) {

	id := r.PathValue("id")
	if id == "" {
		http.Error(w, "Id not found!", http.StatusNotFound)
		return
	}
	dto := usecase.FindCustomerInputDTO{ID: id}
	customerUseCase := usecase.NewFindCustomerUseCase(h.CustomerRepository)

	output, err := customerUseCase.Execute(dto)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	jsonResponse, err := json.Marshal(output)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(jsonResponse)
}

func (h *WebCustomerHandler) UpdateCustomer(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")
	if id == "" {
		http.Error(w, "Id not found!", http.StatusNotFound)
		return
	}

	customerUseCase := usecase.NewUpdateCustomerUseCase(h.CustomerRepository)
	var dto usecase.UpdateCustomerInputDTO
	err := json.NewDecoder(r.Body).Decode(&dto)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	dto.ID = id
	customer, err := customerUseCase.Execute(dto)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	jsonResponse, err := json.Marshal(customer)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(jsonResponse)
}

type DeleteCustomerOutPut struct {
	Message string `json:"message"`
}

func (h *WebCustomerHandler) DeleteCustomer(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")
	if id == "" {
		http.Error(w, "Id not found!", http.StatusNotFound)
		return
	}
	dto := usecase.DeleteCustomerInputDTO{ID: id}
	customerUseCase := usecase.NewDeleteCustomerUseCase(h.CustomerRepository)

	err := customerUseCase.Execute(dto)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	jsonResponse, err := json.Marshal(DeleteCustomerOutPut{Message: "Customer deleted successfully!"})
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(jsonResponse)
}
