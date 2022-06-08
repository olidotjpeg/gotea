package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"github.com/google/uuid"
	"github.com/gorilla/mux"
	_ "github.com/mattn/go-sqlite3"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"time"
)

type Origin struct {
	ShopName     string `json:"shopName" sql:"ShopName"`
	ShopLocation string `json:"shopLocation" sql:"ShopLocation"`
}

type Tea struct {
	Id              string `json:"id" sql:"Id"`
	Origin          Origin `json:"origin" sql:"Origin"`
	Temperature     int    `json:"temperature" sql:"Temperature"`
	PortionWeight   int    `json:"portionWeight" sql:"PortionWeight"`
	ContainerWeight int    `json:"containerWeight" sql:"ContainerWeight"`
	InitialWeight   int    `json:"initialWeight" sql:"InitialWeight"`
	BrewingDuration int    `json:"brewingDuration" sql:"BrewingDuration"`
	TeaType         string `json:"teaType" sql:"TeaType"`
	TeaName         string `json:"teaName" sql:"TeaName"`
}

var database *sql.DB

func getTeas(w http.ResponseWriter, r *http.Request) {
	rows, err := database.Query("SELECT * FROM teas")
	if err != nil {
		log.Fatal(err)
	}

	defer rows.Close()

	var teas []Tea

	for rows.Next() {
		var tea Tea
		err := rows.Scan(&tea.TeaName, &tea.Origin.ShopName, &tea.Origin.ShopLocation, &tea.Temperature, &tea.PortionWeight, &tea.ContainerWeight, &tea.InitialWeight, &tea.BrewingDuration, &tea.Id, &tea.TeaType)
		if err != nil {
			return
		}
		teas = append(teas, tea)
	}

	json.NewEncoder(w).Encode(teas)
}

func returnSingleTea(w http.ResponseWriter, r *http.Request) {
	routeVariables := mux.Vars(r)
	key := routeVariables["id"]

	sqlQuery := fmt.Sprintf(`SELECT id, teaName FROM teas WHERE id = '%s'`, key)

	row := database.QueryRow(sqlQuery)
	var id string
	var teaName string
	err := row.Scan(&id, &teaName)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(id, teaName)
}

func createNewTea(w http.ResponseWriter, r *http.Request) {
	var tea Tea
	tea.Id = uuid.NewString()

	decoder := json.NewDecoder(r.Body)
	decodeError := decoder.Decode(&tea)
	if decodeError != nil {
		panic(decodeError)
	}

	sqlQuery := fmt.Sprintf(`INSERT INTO teas(id, teaName, shopName, shopLocation, teaType, temperature, portionWeight, containerWeight, initialWeight, brewingDuration) VALUES('%s', '%s', '%s', '%s', '%s', %d, %d, %d, %d, %d)`, tea.Id, tea.TeaName, tea.Origin.ShopName, tea.Origin.ShopLocation, tea.TeaType, tea.Temperature, tea.PortionWeight, tea.ContainerWeight, tea.InitialWeight, tea.BrewingDuration)

	var err error

	_, err = database.Exec(sqlQuery)
	if err != nil {
		log.Fatal(err)
	}

	json.NewEncoder(w).Encode(tea)
}

func updateTea(w http.ResponseWriter, r *http.Request) {
	// @TODO turn this into a call to the sql database
	//routeVariables := mux.Vars(r)
	//id := routeVariables["id"]
	//
	//var updatedTea Tea
	//
	//reqBody, _ := ioutil.ReadAll(r.Body)
	//json.Unmarshal(reqBody, &updatedTea)
	//for i, tea := range Teas {
	//	if tea.Id == id {
	//
	//		tea.Temperature = updatedTea.Temperature
	//		tea.PortionWeight = updatedTea.PortionWeight
	//		tea.InitialWeight = updatedTea.InitialWeight
	//		tea.ContainerWeight = updatedTea.ContainerWeight
	//		tea.BrewingDuration = updatedTea.BrewingDuration
	//		tea.Origin.ShopLocation = updatedTea.Origin.ShopLocation
	//		tea.Origin.ShopName = updatedTea.Origin.ShopName
	//		tea.TeaName = updatedTea.TeaName
	//
	//		Teas[i] = tea
	//
	//		json.NewEncoder(w).Encode(tea)
	//	}
	//}
}

func deleteTea(w http.ResponseWriter, r *http.Request) {
	// @TODO turn this into a call to the sql database
	//routeVariables := mux.Vars(r)
	//id := routeVariables["id"]
	//
	//for index, tea := range Teas {
	//	if tea.Id == id {
	//		Teas = append(Teas[:index], Teas[index+1:]...)
	//	}
	//}
	//
	//json.NewEncoder(w).Encode(Teas)
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
		Handler: myRouter,
		Addr:    "0.0.0.0:8000",
		// Good practice: enforce timeouts for servers you create!
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	log.Fatal(srv.ListenAndServe())
}

func main() {
	fmt.Println("Server is live on port 8000")

	var err error

	database, err = sql.Open("sqlite3", "./gotea.sqlite")
	if err != nil {
		log.Fatal(err)
	}
	defer database.Close()

	handleRequests()
}
