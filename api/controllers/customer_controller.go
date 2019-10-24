package controllers

import (
	"encoding/json"
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

func (server *Server) GetCustomers(w http.ResponseWriter, r *http.Request) {
	customer := models.Customer{}

	customers, err := customer.FindAllCustomer(server.DB)
	if err != nil {
		responses.ERROR(w, http.StatusInternalServerError, err)
		return
	}
	responses.JSON(w, http.StatusOK, customers)
}

func (server *Server) GetCustomer(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	uid, err := strconv.ParseUint(vars["id"], 10, 32)
	if err != nil {
		responses.ERROR(w, http.StatusBadRequest, err)
		return
	}
	customer := models.Customer{}
	customerGotten, err := customer.FindCustomerByID(server.DB, uint32(uid))
	if err != nil {
		responses.ERROR(w, http.StatusBadRequest, err)
		return
	}
	responses.JSON(w, http.StatusOK, customerGotten)
}

func (server *Server) CreateCustomer(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}
	customer := models.Customer{}
	err = json.Unmarshal(body, &customer)
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}

	customer.Prepare()
	err = customer.Validate("")
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}
	customerCreated, err := customer.SaveCustomer(server.DB)

	if err != nil {
		formattedError := formateerror.FormatError(err.Error())
		responses.ERROR(w, http.StatusInternalServerError, formattedError)
		return
	}
	w.Header().Set("Location", fmt.Sprintf("%s%s/%d", r.Host, r.RequestURI, customerCreated.ID))
	responses.JSON(w, http.StatusCreated, customerCreated)

}

func (server *Server) UpdateCustomer(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	uid, err := strconv.ParseUint(vars["id"], 10, 32)
	if err != nil {
		responses.ERROR(w, http.StatusBadRequest, err)
		return
	}
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}
	customer := models.Customer{}
	err = json.Unmarshal(body, &customer)
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}
	customer.Prepare()
	err = customer.Validate("update")

	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}
	updatedCustomer, err := customer.UpdateCustomer(server.DB, uint32(uid))
	if err != nil {
		formattedError := formateerror.FormatError(err.Error())
		responses.ERROR(w, http.StatusInternalServerError, formattedError)
		return
	}
	responses.JSON(w, http.StatusOK, updatedCustomer)
}

func (server *Server) DeleteCustomer(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	customer := models.Customer{}

	uid, err := strconv.ParseUint(vars["id"], 10, 32)
	if err != nil {
		responses.ERROR(w, http.StatusBadRequest, err)
		return
	}

	_, err = customer.DeleteCustomer(server.DB, uint32(uid))
	if err != nil {
		responses.ERROR(w, http.StatusInternalServerError, err)
		return
	}
	w.Header().Set("Entity", fmt.Sprintf("%d", uid))
	responses.JSON(w, http.StatusNoContent, "")
}

func (server *Server) ManageCustomerWeb(w http.ResponseWriter, r *http.Request) {
	customersTemplate := append(mainTemplateString, path.Join("views", "customers.html"))
	var tmpl = template.Must(template.ParseFiles(customersTemplate...))

	var data = M{
		"title": baseTitle + "Customer Management",
		"sidebar": "customer",
	}

	err := tmpl.ExecuteTemplate(w, "customers", data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
