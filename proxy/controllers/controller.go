package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/google/uuid"
	"github.com/tonitienda/go-proxy-mermaid/models"
)

var shortID string

func init() {
	// Short ID is based on the first 8 chars of the UUID
	shortID = uuid.NewString()[0:8]
}

func getEnv(key, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return fallback
}

var myClient = &http.Client{Timeout: 10 * time.Second}

func getJson(url string, target interface{}) error {

	req, err := http.NewRequest("GET", url, nil)

	if err != nil {
		fmt.Println("Error creating HTTP request:", err)
		return err
	}

	req.Header.Add("Content-Type", "application/json")

	// Send the HTTP request
	resp, err := myClient.Do(req)
	if err != nil {
		fmt.Println("Error sending HTTP request:", err)
		return err
	}

	defer resp.Body.Close()

	return json.NewDecoder(resp.Body).Decode(target)
}

func getRemoteServiceData() ([]models.Node, error) {
	nextSvc := getEnv("NEXT_SERVICE", "")
	fmt.Println("NEXT_SERVICE: " + nextSvc)
	if len(nextSvc) == 0 {
		return []models.Node{}, nil
	}

	urls := strings.Split(nextSvc, ",")
	allServicesData := []models.Node{}

	for _, url := range urls {
		serviceData := []models.Node{}
		error := getJson(url+"?requestorID="+shortID, &serviceData)

		if error != nil {
			allServicesData = append(allServicesData, models.Node{
				ID:          uuid.NewString()[0:8],
				Name:        url,
				RequestorID: shortID,
				HasError:    true,
				Url:         url,
			})

		} else {
			allServicesData = append(allServicesData, serviceData...)
		}
	}

	return allServicesData, nil

}

func GetCurrentServiceData(requestorID string, serviceURL string) models.Node {
	return models.Node{
		ID:          shortID,
		Name:        getEnv("SERVICE_NAME", "Service"),
		RequestorID: requestorID,
		Url:         serviceURL,
	}
}

func GetServiceData(requestorID string, serviceUrl string) []models.Node {
	currentService := GetCurrentServiceData(requestorID, serviceUrl)
	fmt.Println("Current service", currentService)
	services := []models.Node{}

	remoteServiceData, err := getRemoteServiceData()
	fmt.Println("Remote service data", remoteServiceData)

	services = append(services, currentService)
	if err == nil || len(remoteServiceData) > 0 {
		services = append(services, remoteServiceData...)
	}
	fmt.Println("Services", services)
	return services
}
