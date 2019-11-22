package persistence

import "github.com/fadlikadn/go-api-tutorial/api/models"

type DatabaseHandler interface {
	AddEvent(models.Event) ([]byte, error)
	FindEvent([]byte) (models.Event, error)
	FindEventByName(string) (models.Event, error)
	FindAllAvailableEvents() ([]models.Event, error)
}
