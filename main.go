package main

import (
	"encoding/json"
	"fmt"
	"github.com/google/uuid"
	"github.com/gorilla/mux"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"time"
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
	TeaType string `json:"teaType"`
	TeaName string `json:"teaName"`
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
			TeaType: "Green Tea",
			TeaName: "Organic Sencha Arata",
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
			TeaType: "Green Tea",
			TeaName: "AC Afternoon Tea",
		},
	}
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
			tea.TeaName = updatedTea.TeaName

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

// spaHandler implements the http.Handler interface, so we can use it
// to respond to HTTP requests. The path to the static directory and
// path to the index file within that static directory are used to
// serve the SPA in the given static directory.
type spaHandler struct {
	staticPath string
	indexPath  string
}

// ServeHTTP inspects the URL path to locate a file within the static dir
// on the SPA handler. If a file is found, it will be served. If not, the
// file located at the index path on the SPA handler will be served. This
// is suitable behavior for serving an SPA (single page application).
func (h spaHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	// get the absolute path to prevent directory traversal
	path, err := filepath.Abs(r.URL.Path)
	if err != nil {
		// if we failed to get the absolute path respond with a 400 bad request
		// and stop
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// prepend the path with the path to the static directory
	path = filepath.Join(h.staticPath, path)

	// check whether a file exists at the given path
	_, err = os.Stat(path)
	if os.IsNotExist(err) {
		// file does not exist, serve index.html
		http.ServeFile(w, r, filepath.Join(h.staticPath, h.indexPath))
		return
	} else if err != nil {
		// if we got an error (that wasn't that the file doesn't exist) stating the
		// file, return a 500 internal server error and stop
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// otherwise, use http.FileServer to serve the static dir
	http.FileServer(http.Dir(h.staticPath)).ServeHTTP(w, r)
}

func handleRequests() {
	myRouter := mux.NewRouter().StrictSlash(true)
	myRouter.HandleFunc("/teas", getTeas)
	// NOTE: Ordering is important here! This has to be defined before
	// the other `/tea` endpoint.
	myRouter.HandleFunc("/tea", createNewTea).Methods("POST")
	myRouter.HandleFunc("/tea/{id}", deleteTea).Methods("DELETE")
	myRouter.HandleFunc("/tea/{id}", updateTea).Methods("PUT")
	myRouter.HandleFunc("/tea/{id}", returnSingleTea)


	spa := spaHandler{staticPath: "static", indexPath: "index.html"}
	myRouter.PathPrefix("/").Handler(spa)

	srv := &http.Server{
		Handler:      myRouter,
		Addr:         "0.0.0.0:8000",
		// Good practice: enforce timeouts for servers you create!
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}


	log.Fatal(srv.ListenAndServe())
}

func main() {
	fmt.Println("Server is live on port 8000")
	generateTeas()
	handleRequests()
}