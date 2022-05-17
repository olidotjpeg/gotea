package main

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"io/ioutil"
	"log"
	"net/http"
)

type Article struct {
	Id      string `json:"Id"`
	Title   string `json:"Title"`
	Desc    string `json:"desc"`
	Content string `json:"content"`
}

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
			Id: "1",
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
			Id: "2",
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

	// Loop over all of our Articles
	// if the article.Id equals the key we pass in
	// return the article encoded as JSON
	for _, tea := range Teas {
		if tea.Id == key {
			json.NewEncoder(w).Encode(tea)
		}
	}
}

func createNewTea(w http.ResponseWriter, r *http.Request) {
	reqBody, _ := ioutil.ReadAll(r.Body)
	var tea Tea
	json.Unmarshal(reqBody, &tea)

	Teas = append(Teas, tea)

	json.NewEncoder(w).Encode(tea)
}

func updateTea(w http.ResponseWriter, r *http.Request) {
	routeVariables := mux.Vars(r)
	id := routeVariables["id"]
	var updatedEvent Tea
	reqBody, _ := ioutil.ReadAll(r.Body)
	json.Unmarshal(reqBody, &updatedEvent)
	for i, tea := range Teas {
		if tea.Id == id {

			tea.Temperature = updatedEvent.Temperature
			tea.PortionWeight = updatedEvent.PortionWeight
			tea.InitialWeight = updatedEvent.InitialWeight
			tea.ContainerWeight = updatedEvent.ContainerWeight
			tea.BrewingDuration = updatedEvent.BrewingDuration
			tea.Origin.ShopLocation = updatedEvent.Origin.ShopLocation
			tea.Origin.ShopName = updatedEvent.Origin.ShopName

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