package controllers

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/fadlikadn/go-api-tutorial/api/models"
	"github.com/fadlikadn/go-api-tutorial/api/responses"
	"github.com/fadlikadn/go-api-tutorial/api/utils/formateerror"
	"github.com/gorilla/mux"
	"html/template"
	"io/ioutil"
	"net/http"
	"path"
	"strconv"
)

func (server *Server) CreateServiceTransaction(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}
	serviceTransaction := models.ServiceTransaction{}
	err = json.Unmarshal(body, &serviceTransaction)
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}
	serviceTransaction.Prepare()
	err = serviceTransaction.Validate()
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}

	serviceTransactionCreated, err := serviceTransaction.SaveServiceTransaction(server.DB)
	if err != nil {
		formattedError := formateerror.FormatError(err.Error())
		responses.ERROR(w, http.StatusInternalServerError, formattedError)
		return
	}
	w.Header().Set("Location", fmt.Sprintf("%s%s/%d", r.Host, r.URL.Path, serviceTransactionCreated.ID))
	responses.JSON(w, http.StatusCreated, serviceTransactionCreated)
}

func (server *Server) GetServiceTransactions(w http.ResponseWriter, r *http.Request) {
	serviceTransaction := models.ServiceTransaction{}

	serviceTransactions, err := serviceTransaction.FindAllServiceTransactions(server.DB)
	if err != nil {
		responses.ERROR(w, http.StatusInternalServerError, err)
		return
	}
	responses.JSON(w, http.StatusOK, serviceTransactions)
}

func (server *Server) GetServiceTransaction(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	pid, err := strconv.ParseUint(vars["id"], 10, 32)
	if err != nil {
		responses.ERROR(w, http.StatusBadRequest, err)
		return
	}
	serviceTransaction := models.ServiceTransaction{}

	serviceTransactionReceived, err := serviceTransaction.FindServiceTransactionByID(server.DB, uint32(pid))
	if err != nil {
		responses.ERROR(w, http.StatusInternalServerError, err)
		return
	}

	responses.JSON(w, http.StatusOK, serviceTransactionReceived)
}

func (server *Server) UpdateServiceTransaction(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	// check if the serviceTransaction id is valid
	pid, err := strconv.ParseUint(vars["id"], 10, 32)
	if err != nil {
		responses.ERROR(w, http.StatusBadRequest, err)
		return
	}

	// Check if the service transaction exist
	serviceTransaction := models.ServiceTransaction{}
	err = server.DB.Debug().Model(models.ServiceTransaction{}).Where("id = ?", pid).Take(&serviceTransaction).Error
	if err != nil {
		responses.ERROR(w, http.StatusNotFound, errors.New("Service Transaction not found"))
		return
	}

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}

	// Start processing the request data
	serviceTransactionUpdate := models.ServiceTransaction{}
	err = json.Unmarshal(body, &serviceTransactionUpdate)
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}

	serviceTransactionUpdate.Prepare()
	err = serviceTransactionUpdate.Validate()
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}
}

func (server *Server) DeleteServiceTransaction(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	pid, err := strconv.ParseUint(vars["id"], 10, 32)
	if err != nil {
		responses.ERROR(w, http.StatusBadRequest, err)
		return
	}

	serviceTransaction := models.ServiceTransaction{}
	err = server.DB.Debug().Model(models.ServiceTransaction{}).Where("id = ?", pid).Take(&serviceTransaction).Error
	if err != nil {
		responses.ERROR(w, http.StatusNotFound, errors.New("Unauthorized"))
		return
	}

	_, err = serviceTransaction.DeleteServiceTransaction(server.DB, uint32(pid))
	if err != nil {
		responses.ERROR(w, http.StatusBadRequest, err)
		return
	}

	w.Header().Set("Entity", fmt.Sprintf("%d", pid))
	responses.JSON(w, http.StatusNoContent, "")
}

func (server *Server) ManageServiceTransactionWeb(w http.ResponseWriter, r *http.Request) {
	servicesTemplate := append(mainTemplateString, path.Join("views", "service-transactions.html"))
	var tmpl = template.Must(template.ParseFiles(servicesTemplate...))

	var data = M{
		"title": baseTitle + "Service Transaction Management",
		"sidebar": "service-transaction",
	}

	err := tmpl.ExecuteTemplate(w, "service-transactions", data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
