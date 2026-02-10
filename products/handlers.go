package products

import (
	"encoding/json"
	"gocrud/core"
	"net/http"

	"github.com/go-chi/chi"
)

func ProductListHandler(w http.ResponseWriter, r *http.Request, service *ProductService) {
	var productResponseDTOS []ProductResponseDTO = []ProductResponseDTO{}
	var err error = nil

	productResponseDTOS, err = service.GetAllProduct()

	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(core.DetailResponseErrorDTO{
			Detail: err.Error(),
		})
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(productResponseDTOS)
}

func ProductCreateHandler(w http.ResponseWriter, r *http.Request, service *ProductService) {
	var err error = nil
	var ValidationErrors core.ValidationErrors = core.ValidationErrors{}
	var productRequestDTO ProductRequestDTO = ProductRequestDTO{}
	var productResponseDTO ProductResponseDTO = ProductResponseDTO{}

	json.NewDecoder(r.Body).Decode(&productRequestDTO)

	ValidationErrors, err = productRequestDTO.Validate()

	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(ValidationErrors)
		return
	}

	productResponseDTO, err = service.CreateProduct(productRequestDTO)

	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(core.DetailResponseErrorDTO{Detail: err.Error()})
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(productResponseDTO)
}

func ProductDetailHandler(w http.ResponseWriter, r *http.Request, service *ProductService) {
	var err error = nil
	var ID string = chi.URLParam(r, "id")
	var productResponseDTO ProductResponseDTO = ProductResponseDTO{}

	productResponseDTO, err = service.GetProduct(ID)
	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode(core.DetailResponseErrorDTO{Detail: err.Error()})
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(productResponseDTO)
}

func ProductUpdateHandler(w http.ResponseWriter, r *http.Request, service *ProductService) {
	var err error = nil
	var ID string = chi.URLParam(r, "id")
	var validationErrors core.ValidationErrors = core.ValidationErrors{}
	var productRequestDTO ProductRequestDTO = ProductRequestDTO{}
	var productResponseDTO ProductResponseDTO = ProductResponseDTO{}

	json.NewDecoder(r.Body).Decode(&productRequestDTO)

	validationErrors, err = productRequestDTO.Validate()
	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(validationErrors)
		return
	}

	productResponseDTO, err = service.UpdateProduct(ID, productRequestDTO)
	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(core.DetailResponseErrorDTO{Detail: err.Error()})
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(productResponseDTO)
}

func ProductDeleteHandler(w http.ResponseWriter, r *http.Request, service *ProductService) {
	var err error = nil
	var ID string = chi.URLParam(r, "id")

	err = service.DeleteProduct(ID)
	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode(core.DetailResponseErrorDTO{Detail: err.Error()})
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusNoContent)
	json.NewEncoder(w).Encode(nil)
}
