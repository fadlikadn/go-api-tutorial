package controllers

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/fadlikadn/go-api-tutorial/api/models"
	"github.com/fadlikadn/go-api-tutorial/api/responses"
	"github.com/fadlikadn/go-api-tutorial/api/utils/email"
	"github.com/fadlikadn/go-api-tutorial/api/utils/formateerror"
	"github.com/gorilla/mux"
	"github.com/tidwall/gjson"
	"github.com/unidoc/unipdf/v3/creator"
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

func (server *Server) ActionCreateNewAdditionalItem(additionalItemValue gjson.Result, serviceTransactionId int) (*models.AdditionalItem, error) {
	fmt.Println(additionalItemValue)
	additionalItem := models.AdditionalItem{}
	err := additionalItem.UnmarshalJSON([]byte(additionalItemValue.String()), serviceTransactionId)
	if err != nil {
		fmt.Println("error on unmarshal json additional item")
		fmt.Println(err)
	}

	fmt.Println(additionalItem)

	additionalItem.Prepare()
	err = additionalItem.Validate()
	if err != nil {
		return nil, err
	}

	additionalItemCreated, err := additionalItem.SaveAdditionalItem(server.DB)
	if err != nil {
		return nil, err
	}
	return additionalItemCreated, nil
}

func (server *Server) ActionCreateNewCustomer(customerValue gjson.Result) (*models.Customer, error) {
	customer := models.Customer{}
	err := json.Unmarshal([]byte(customerValue.String()), &customer)
	if err != nil {
		return nil, err
	}

	customer.Prepare()
	err = customer.Validate("")
	if err != nil {
		return nil, err
	}
	customerCreated, err := customer.SaveCustomer(server.DB)

	if err != nil {
		return nil, err
	}

	//fmt.Println(customerCreated)
	return customerCreated, nil
}

func (server *Server) ActionCreateNewTransaction(serviceTransactionValue gjson.Result, customerId int) (*models.ServiceTransaction, error) {
	serviceTransaction := models.ServiceTransaction{}
	err := serviceTransaction.UnmarshalJSON([]byte(serviceTransactionValue.String()), customerId)
	if err != nil {
		fmt.Println("error again")
		fmt.Println(err)
	}

	fmt.Println("data service transaction")
	//fmt.Println(serviceTransaction)

	serviceTransaction.Prepare()
	err = serviceTransaction.Validate()
	if err != nil {
		return nil, err
	}

	serviceTransactionCreated, err := serviceTransaction.SaveServiceTransaction(server.DB)
	if err != nil {
		return nil, err
	}

	return serviceTransactionCreated, nil
}

func (server *Server) CreateComplexServiceTransaction(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}
	customerValue := gjson.Get(string(body), "customer")
	serviceTransactionValue := gjson.Get(string(body), "serviceTransaction")
	additionalItemsValue := gjson.Get(string(body), "additionalItems")
	fmt.Println(additionalItemsValue)

	//fmt.Println(gjson.Get(customerValue.String(), "id"))
	//fmt.Println(gjson.Get(customerValue.String(), "id").Type)
	if gjson.Get(customerValue.String(), "id").Type != gjson.Null {
		// existing customer
		customerInt := int(gjson.Get(customerValue.String(), "id").Int())
		serviceTransactionCreated, err := server.ActionCreateNewTransaction(serviceTransactionValue, customerInt)
		if err != nil {
			formattedError := formateerror.FormatError(err.Error())
			responses.ERROR(w, http.StatusInternalServerError, formattedError)
		}

		additionalItemCreated := &models.AdditionalItem{}
		additionalItemsValue.ForEach(func(key, value gjson.Result) bool {
			additionalItemCreated, err = server.ActionCreateNewAdditionalItem(value, int(serviceTransactionCreated.ID))
			return true
		})

		serviceTransaction := models.ServiceTransaction{}
		serviceTransactionCreated, err = serviceTransaction.FindServiceTransactionByID(server.DB, serviceTransactionCreated.ID)
		if err != nil {
			formattedError := formateerror.FormatError(err.Error())
			responses.ERROR(w, http.StatusInternalServerError, formattedError)
		}

		w.Header().Set("Location", fmt.Sprintf("%s%s/%d", r.Host, r.URL.Path, serviceTransactionCreated.ID))
		responses.JSON(w, http.StatusCreated, serviceTransactionCreated)
	} else {
		// create new customer
		fmt.Println("create new customer first")

		customerCreated, err := server.ActionCreateNewCustomer(customerValue)
		if err != nil {
			formattedError := formateerror.FormatError(err.Error())
			responses.ERROR(w, http.StatusInternalServerError, formattedError)
		}

		customerInt := int(customerCreated.ID)
		serviceTransactionCreated, err := server.ActionCreateNewTransaction(serviceTransactionValue, customerInt)
		if err != nil {
			formattedError := formateerror.FormatError(err.Error())
			responses.ERROR(w, http.StatusInternalServerError, formattedError)
		}

		additionalItemCreated := &models.AdditionalItem{}
		additionalItemsValue.ForEach(func(key, value gjson.Result) bool {
			additionalItemCreated, err = server.ActionCreateNewAdditionalItem(value, int(serviceTransactionCreated.ID))
			return true
		})

		serviceTransaction := models.ServiceTransaction{}
		serviceTransactionCreated, err = serviceTransaction.FindServiceTransactionByID(server.DB, serviceTransactionCreated.ID)
		if err != nil {
			formattedError := formateerror.FormatError(err.Error())
			responses.ERROR(w, http.StatusInternalServerError, formattedError)
		}

		w.Header().Set("Location", fmt.Sprintf("%s%s/%d", r.Host, r.URL.Path, serviceTransactionCreated.ID))
		responses.JSON(w, http.StatusCreated, serviceTransactionCreated)
	}
}

func (server *Server) ActionUpdateTransaction(serviceTransactionValue gjson.Result, customerId int, pid int) (*models.ServiceTransaction, error) {
	serviceTransaction := models.ServiceTransaction{}
	err := serviceTransaction.UnmarshalJSON([]byte(serviceTransactionValue.String()), customerId)
	if err != nil {
		fmt.Println("error again")
		fmt.Println(err)
	}

	serviceTransaction.Prepare()
	err = serviceTransaction.Validate()
	if err != nil {
		fmt.Println("error on validate")
		return nil, err
	}

	updatedServiceTransaction, err := serviceTransaction.UpdateServiceTransaction(server.DB, uint32(pid))
	if err != nil {
		fmt.Println("error on update service transaction")
		return nil, err
	}

	return updatedServiceTransaction, nil
}

func (server *Server) UpdateComplexServiceTransaction(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
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

	serviceTransactionValue := gjson.Get(string(body), "serviceTransaction")
	additionalItemsValue := gjson.Get(string(body), "additionalItems")
	fmt.Println(additionalItemsValue)

	customerId := int(gjson.Get(serviceTransactionValue.String(), "customer_id").Int())
	serviceTransactionUpdated, err := server.ActionUpdateTransaction(serviceTransactionValue, customerId, int(pid))
	if err != nil {
		formattedError := formateerror.FormatError(err.Error())
		responses.ERROR(w, http.StatusInternalServerError, formattedError)
	}

	// Delete existing additional item
	additionaltem := models.AdditionalItem{}
	_, err = additionaltem.DeleteAllAdditionalItemBySTid(server.DB, serviceTransactionUpdated.ID)
	//if deletedAdditionalItems == 0{
	//	formattedError := formateerror.FormatError(err.Error())
	//	responses.ERROR(w, http.StatusInternalServerError, formattedError)
	//}

	additionalItemCreated := &models.AdditionalItem{}
	additionalItemsValue.ForEach(func(key, value gjson.Result) bool {
		additionalItemCreated, err = server.ActionCreateNewAdditionalItem(value, int(serviceTransactionUpdated.ID))
		return true
	})

	serviceTransactionUpdated, err = serviceTransaction.FindServiceTransactionByID(server.DB, serviceTransactionUpdated.ID)
	if err != nil {
		formattedError := formateerror.FormatError(err.Error())
		responses.ERROR(w, http.StatusInternalServerError, formattedError)
	}

	w.Header().Set("Location", fmt.Sprintf("%s%s/%d", r.Host, r.URL.Path, serviceTransactionUpdated.ID))
	responses.JSON(w, http.StatusCreated, serviceTransactionUpdated)
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

func (server *Server) AddServiceTransactionWeb(w http.ResponseWriter, r *http.Request) {
	servicesTemplate := append(mainTemplateString, path.Join("views", "service-transactions-add.html"))
	var tmpl = template.Must(template.ParseFiles(servicesTemplate...))

	var data = M{
		"title": baseTitle + "Add Service Transaction",
		"sidebar": "service-transaction",
	}

	err := tmpl.ExecuteTemplate(w, "service-transactions-add", data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
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

func (server *Server) SendTransactionStatusEmail(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	pid, err := strconv.ParseUint(vars["id"], 10, 32)
	if err != nil {
		responses.ERROR(w, http.StatusBadRequest, err)
		return
	}

	serviceTransaction := models.ServiceTransaction{}
	serviceTransactionGotten, err := serviceTransaction.FindServiceTransactionByID(server.DB, uint32(pid))
	if err != nil {
		responses.ERROR(w, http.StatusBadRequest, err)
		return
	}

	_, err = email.SendStatusEmail(serviceTransactionGotten.Customer.Email)
	if err != nil {
		responses.ERROR(w, http.StatusBadRequest, err)
		return
	}

	w.Header().Set("Location", fmt.Sprintf("%s%s/%d", r.Host, r.URL.Path, serviceTransactionGotten.ID))
	responses.JSON(w, http.StatusCreated, serviceTransactionGotten)
}

func (server *Server) CreateInvoiceServiceTransaction(w http.ResponseWriter, r *http.Request) {
	// Instantiate new PDF creator
	c := creator.New()

	// Create a new PDF page and select it for editing
	c.NewPage()

	// Create new invoice and populate it with date
	serviceTransaction := models.ServiceTransaction{}
	invoice, err := serviceTransaction.CreateInvoice(c, "")
	if err != nil {
		responses.ERROR(w, http.StatusBadRequest, err)
		return
	}

	// Write invoice to page
	err = c.Draw(invoice)
	if err != nil {
		responses.ERROR(w, http.StatusBadRequest, err)
		return
	}

	// Write output file
	// Alternative is writing to a Writer interface by using c.Write
	//err = c.WriteToFile("simple_invoice.pdf")
	err = c.Write(w)
	if err != nil {
		fmt.Println("error on write PDF File")
		responses.ERROR(w, http.StatusBadRequest, err)
		return
	}
	w.Header().Set("Content-type", "application/pdf")
}
