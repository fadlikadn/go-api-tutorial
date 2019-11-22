package controllers

import (
	"encoding/hex"
	"encoding/json"
	"errors"
	"github.com/fadlikadn/go-api-tutorial/api/models"
	"github.com/fadlikadn/go-api-tutorial/api/responses"
	"github.com/fadlikadn/go-api-tutorial/persistence"
	"github.com/gorilla/mux"
	"net/http"
	"strings"
)

type EventServiceHandler struct {
	dbhandler persistence.DatabaseHandler
}

func NewEventHandler(databasehandler persistence.DatabaseHandler) *EventServiceHandler {
	return &EventServiceHandler{dbhandler: databasehandler}
}

func (eh *EventServiceHandler) FindEventHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	criteria, ok := vars["SearchCriteria"]
	if !ok {
		responses.ERROR(w, http.StatusUnprocessableEntity, errors.New("no search keys found, you can either search by id via /id/4, to search by name via /name/coldplayconcern"))
		return
	}
	searchkey, ok := vars["search"]
	if !ok {
		responses.ERROR(w, http.StatusUnprocessableEntity, errors.New("no search keys found, you can either search by id via /id/4, to search by name via /name/coldplayconcern"))
		return
	}
	var event models.Event
	var err error
	switch strings.ToLower(criteria) {
	case "name" :
		event, err = eh.dbhandler.FindEventByName(searchkey)
	case "id" :
		id, err := hex.DecodeString(searchkey)
		if err == nil {
			event, err = eh.dbhandler.FindEvent(id)
		}
	}

	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}
	responses.JSON(w, http.StatusOK, event)
}

func (eh *EventServiceHandler) AllEventHandler(w http.ResponseWriter, r *http.Request) {
	events, err := eh.dbhandler.FindAllAvailableEvents()
	if err != nil {
		responses.ERROR(w, http.StatusBadGateway, errors.New("error occured while trying to find available events"))
		return
	}
	responses.JSON(w, http.StatusOK, events)
}

func (eh *EventServiceHandler) NewEventHandler(w http.ResponseWriter, r *http.Request) {
	event := models.Event{}
	err := json.NewDecoder(r.Body).Decode(&event)
	if err != nil {
		responses.ERROR(w, http.StatusBadGateway, errors.New("error occured while decoding event data"))
		return
	}
	id, err := eh.dbhandler.AddEvent(event)
	if err != nil {
		responses.ERROR(w, http.StatusBadGateway, errors.New("error occurred while persisting event"))
		return
	}
	responses.JSON(w, http.StatusCreated, id)
}

