package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"github.com/google/uuid"
	"github.com/gorilla/mux"
)

type Origin struct {
	ShopName     string `json:"shopName"`
	ShopLocation string `json:"shopLocation"`
}

type Tea struct {
	Id          string `json:"id"`
	Origin        Origin `json:"origin"`
	Temperature     int `json:"temperature"`
	PortionWeight   int `json:"portionWeight"`
	ContainerWeight int `json:"containerWeight"`
	InitialWeight   int `json:"initialWeight"`
	BrewingDuration int `json:"brewingDuration"`
}

var Teas []Tea

func generateTeas() {
	Teas = []Tea{
		Tea{
			Id: uuid.NewString(),
			Origin: Origin{
				ShopName:     "AC Perchs",
				ShopLocation: "Copenhagen",
			},
			Temperature:     80,
			PortionWeight:   10,
			ContainerWeight: 95,
			InitialWeight:   100,
			BrewingDuration: 120,
		},
		Tea{
			Id: uuid.NewString(),
			Origin: Origin{
				ShopName:     "Edeka Zurheide",
				ShopLocation: "Düsseldorf",
			},
			Temperature:     120,
			PortionWeight:   50,
			ContainerWeight: 35,
			InitialWeight:   120,
			BrewingDuration: 1000,
		},
	}
}

func homePage(w http.ResponseWriter, r *http.Request){
	fmt.Fprintf(w, "Welcome to the HomePage!")
}

func getTeas(w http.ResponseWriter, r *http.Request){
	json.NewEncoder(w).Encode(Teas)
}

func returnSingleTea(w http.ResponseWriter, r *http.Request) {
	routeVariables := mux.Vars(r)
	key := routeVariables["id"]

	for _, tea := range Teas {
		if tea.Id == key {
			json.NewEncoder(w).Encode(tea)
		}
	}
}

func createNewTea(w http.ResponseWriter, r *http.Request) {
	reqBody, _ := ioutil.ReadAll(r.Body)

	var tea Tea
	tea.Id = uuid.NewString()
	json.Unmarshal(reqBody, &tea)

	Teas = append(Teas, tea)
	json.NewEncoder(w).Encode(tea)
}

func updateTea(w http.ResponseWriter, r *http.Request) {
	routeVariables := mux.Vars(r)
	id := routeVariables["id"]

	var updatedTea Tea

	reqBody, _ := ioutil.ReadAll(r.Body)
	json.Unmarshal(reqBody, &updatedTea)
	for i, tea := range Teas {
		if tea.Id == id {

			tea.Temperature = updatedTea.Temperature
			tea.PortionWeight = updatedTea.PortionWeight
			tea.InitialWeight = updatedTea.InitialWeight
			tea.ContainerWeight = updatedTea.ContainerWeight
			tea.BrewingDuration = updatedTea.BrewingDuration
			tea.Origin.ShopLocation = updatedTea.Origin.ShopLocation
			tea.Origin.ShopName = updatedTea.Origin.ShopName

			Teas[i] = tea

			json.NewEncoder(w).Encode(tea)
		}
	}
}

func deleteTea(w http.ResponseWriter, r *http.Request) {
	routeVariables := mux.Vars(r)
	id := routeVariables["id"]

	for index, tea := range Teas {
		if tea.Id == id {
			Teas = append(Teas[:index], Teas[index+1:]...)
		}
	}

	json.NewEncoder(w).Encode(Teas)
}

func handleRequests() {
	myRouter := mux.NewRouter().StrictSlash(true)
	myRouter.HandleFunc("/", homePage)
	myRouter.HandleFunc("/teas", getTeas)
	// NOTE: Ordering is important here! This has to be defined before
	// the other `/tea` endpoint.
	myRouter.HandleFunc("/tea", createNewTea).Methods("POST")
	myRouter.HandleFunc("/tea/{id}", deleteTea).Methods("DELETE")
	myRouter.HandleFunc("/tea/{id}", updateTea).Methods("PUT")
	myRouter.HandleFunc("/tea/{id}", returnSingleTea)
	log.Fatal(http.ListenAndServe(":10000", myRouter))
}

func main() {
	fmt.Println("Rest API v2.0 - Mux Routers")
	generateTeas()
	handleRequests()
}