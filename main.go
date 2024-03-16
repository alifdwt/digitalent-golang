package main

import (
	"encoding/json"
	"fmt"
	"html/template"
	"log"
	"math/rand"
	"net/http"
	"os"
	"time"
)

type Status struct {
	Status struct {
		Water int `json:"water"`
		Wind  int `json:"wind"`
	} `json:"status"`
}

type StatusWithCondition struct {
	Water struct {
		Value     int    `json:"value"`
		Condition string `json:"condition"`
	} `json:"water"`
	Wind struct {
		Value     int    `json:"value"`
		Condition string `json:"condition"`
	} `json:"wind"`
}

func main() {
	filePath := "status.json"

	go func() {
		for {
			updateStatus(filePath)
			statusLogger(statusToStatusWithCondition(filePath))
			time.Sleep(15 * time.Second)
		}
	}()

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		statusWithCondition := statusToStatusWithCondition(filePath)

		templateHTML, err := template.ParseFiles("views/index.html")
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		err = templateHTML.Execute(w, statusWithCondition)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	})

	fmt.Println("Server started at http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func updateStatus(filePath string) Status {
	water := rand.Intn(100) + 1
	wind := rand.Intn(100) + 1

	status := Status{}
	status.Status.Water = water
	status.Status.Wind = wind

	statusJSON, err := json.MarshalIndent(status, "", "  ")
	if err != nil {
		fmt.Println("Error while marshalling JSON: ", err)
		return Status{}
	}

	err = os.WriteFile(filePath, statusJSON, 0644)
	if err != nil {
		fmt.Println("Error while writing to JSON file: ", err)
		return Status{}
	}

	return status
}

func statusToStatusWithCondition(filePath string) StatusWithCondition {
	file, err := os.ReadFile(filePath)
	if err != nil {
		log.Fatal("Error while reading JSON file: ", err)
	}

	var status Status
	err = json.Unmarshal(file, &status)
	if err != nil {
		log.Fatal("Error while unmarshalling JSON: ", err)
	}

	var waterCondition string
	if status.Status.Water < 5 {
		waterCondition = "aman"
	} else if status.Status.Water < 9 {
		waterCondition = "siaga"
	} else {
		waterCondition = "bahaya"
	}

	var windCondition string
	if status.Status.Wind < 6 {
		windCondition = "aman"
	} else if status.Status.Wind < 16 {
		windCondition = "siaga"
	} else {
		windCondition = "bahaya"
	}

	statusWithCondition := StatusWithCondition{
		Water: struct {
			Value     int    `json:"value"`
			Condition string `json:"condition"`
		}{
			Value:     status.Status.Water,
			Condition: waterCondition,
		},
		Wind: struct {
			Value     int    `json:"value"`
			Condition string `json:"condition"`
		}{
			Value:     status.Status.Wind,
			Condition: windCondition,
		},
	}

	return statusWithCondition
}

func statusLogger(status StatusWithCondition) {
	fmt.Println("Status updated successfully!")
	fmt.Printf("|| Water: %dm 			|| Wind: %dm/s 			||\n", status.Water.Value, status.Wind.Value)
	fmt.Printf("|| Water condition: %s 	|| Wind condition: %s 	||\n\n", status.Water.Condition, status.Wind.Condition)
}
