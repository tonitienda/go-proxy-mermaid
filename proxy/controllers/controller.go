package controllers

import (
	"encoding/json"
	"net/http"
	"os"
	"time"

	"github.com/google/uuid"
	"github.com/tonitienda/go-proxy-mermaid/models"
)

func init() {
	// Short ID is based on the first 8 chars of the UUID
	shortID := uuid.NewString()[0:8]
	os.Setenv("SERVICE_ID", shortID)
}

func getEnv(key, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return fallback
}

var myClient = &http.Client{Timeout: 10 * time.Second}

func getJson(url string, target interface{}) error {
	r, err := myClient.Get(url)
	if err != nil {
		return err
	}
	defer r.Body.Close()

	return json.NewDecoder(r.Body).Decode(target)
}

func getRemoteServiceData() ([]models.Node, error) {
	nextSvc := getEnv("NEXT_SERVICE", "")

	if len(nextSvc) == 0 {
		return []models.Node{}, nil
	}

	serviceData := []models.Node{}
	getJson(nextSvc, &serviceData)

	return serviceData, nil

}

func GetCurrentServiceData() models.Node {
	return models.Node{
		ID:   getEnv("SERVICE_ID", "<no id>"),
		Name: getEnv("SERVICE_NAME", "Service"),
	}
}

func GetServiceData() []models.Node {

	remoteServiceData, err := getRemoteServiceData()

	services := []models.Node{}

	services = append(services, GetCurrentServiceData())

	if err == nil || len(remoteServiceData) > 0 {
		for _, remoteSvc := range remoteServiceData {
			services = append(services, remoteSvc)

		}
	}

	return services
}
