package web

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/google/uuid"
	"github.com/gorilla/mux"
)

func getTeas(w http.ResponseWriter, r *http.Request) {
	rows, err := Database.Query("SELECT * FROM teas")
	if err != nil {
		fmt.Println("No such table HMMM")
	}

	defer rows.Close()

	var teas []Tea

	for rows.Next() {
		var tea Tea
		err := rows.Scan(&tea.TeaName, &tea.Origin.ShopName, &tea.Origin.ShopLocation, &tea.Temperature, &tea.PortionWeight, &tea.ContainerWeight, &tea.InitialWeight, &tea.BrewingDuration, &tea.Id, &tea.TeaType, &tea.Color, &tea.InUse, &tea.Size, &tea.BlendDescription)
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

	sqlQuery := fmt.Sprintf(`SELECT * FROM teas WHERE id = '%s'`, key)

	row := Database.QueryRow(sqlQuery)
	var tea Tea
	err := row.Scan(&tea.TeaName, &tea.Origin.ShopName, &tea.Origin.ShopLocation, &tea.Temperature, &tea.PortionWeight, &tea.ContainerWeight, &tea.InitialWeight, &tea.BrewingDuration, &tea.Id, &tea.TeaType, &tea.Color, &tea.InUse, &tea.Size, &tea.BlendDescription)
	if err != nil {
		log.Fatal(err)
	}

	json.NewEncoder(w).Encode(tea)
}

func createNewTea(w http.ResponseWriter, r *http.Request) {
	var tea Tea
	tea.Id = uuid.NewString()

	decoder := json.NewDecoder(r.Body)
	decodeError := decoder.Decode(&tea)
	if decodeError != nil {
		panic(decodeError)
	}

	sqlQuery := fmt.Sprintf(`INSERT INTO teas(id, teaName, shopName, shopLocation, temperature, portionWeight, containerWeight, initialWeight, brewingDuration, teaType, color, inUse, size, blendDescription) VALUES('%s', '%s', '%s', '%s', %d, %d, %d, %d, %d, '%s', '%s', %d, '%s', '%s')`, tea.Id, tea.TeaName, tea.Origin.ShopName, tea.Origin.ShopLocation, tea.Temperature, tea.PortionWeight, tea.ContainerWeight, tea.InitialWeight, tea.BrewingDuration, tea.TeaType, tea.Color, tea.InUse, tea.Size, tea.BlendDescription)

	var err error

	_, err = Database.Exec(sqlQuery)
	if err != nil {
		log.Fatal(err)
	}

	json.NewEncoder(w).Encode(tea)
}

func updateTea(w http.ResponseWriter, r *http.Request) {
	routeVariables := mux.Vars(r)
	id := routeVariables["id"]

	var updatedTea Tea

	reqBody, _ := ioutil.ReadAll(r.Body)
	json.Unmarshal(reqBody, &updatedTea)

	sqlQuery := fmt.Sprintf(
		`UPDATE teas SET teaName = '%s', teaType = '%s', shopName = '%s', shopLocation = '%s', temperature = '%d', portionWeight = '%d', containerWeight = '%d', initialWeight = '%d', brewingDuration = '%d', teaType = '%s', color = '%s', inUse = '%d', size = '%s', blendDescription = '%s' WHERE id = '%s'`,
		updatedTea.TeaName, updatedTea.TeaType, updatedTea.Origin.ShopName, updatedTea.Origin.ShopLocation, updatedTea.Temperature, updatedTea.PortionWeight, updatedTea.ContainerWeight, updatedTea.InitialWeight, updatedTea.BrewingDuration, updatedTea.TeaType, updatedTea.Color, updatedTea.InUse, updatedTea.Size, updatedTea.BlendDescription, id)

	Database.Exec(sqlQuery)

	json.NewEncoder(w).Encode(updatedTea)
}

func deleteTea(w http.ResponseWriter, r *http.Request) {
	routeVariables := mux.Vars(r)
	id := routeVariables["id"]

	Database.Exec("DELETE FROM teas WHERE id = ?", id)

	json.NewEncoder(w).Encode("Teas Deleted")
}
