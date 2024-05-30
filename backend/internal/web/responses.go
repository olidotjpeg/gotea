package web

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

// var connString string = "postgresql://docker:postgres123@postgres:5432/gotea?sslmode=disable"

func getTeas(w http.ResponseWriter, r *http.Request) {
	// Execute the query
	rows, err := pool.Query(context.Background(), "SELECT * FROM teas")
	if err != nil {
		log.Fatalf("Query failed: %v\n", err)
	}
	defer rows.Close()

	// // Iterate over the rows and print each tea
	var teas []Tea
	for rows.Next() {
		var tea Tea

		err := rows.Scan(&tea.Id, &tea.TeaName, &tea.TeaType, &tea.Origin.ShopName, &tea.Origin.ShopLocation, &tea.Temperature, &tea.PortionWeight, &tea.ContainerWeight, &tea.InitialWeight, &tea.BrewingDuration, &tea.Color, &tea.Size, &tea.InUse, &tea.BlendDescription)
		if err != nil {
			log.Fatalf("Unable to scan row: %v\n", err)
		}

		teas = append(teas, tea)
	}

	if rows.Err() != nil {
		log.Fatalf("Row iteration error: %v\n", rows.Err())
	}

	json.NewEncoder(w).Encode(teas)
}

func returnSingleTea(w http.ResponseWriter, r *http.Request) {
	routeVariables := mux.Vars(r)
	key := routeVariables["id"]

	sqlQuery := `SELECT * FROM teas WHERE id = $1`
	row := pool.QueryRow(context.Background(), sqlQuery, key)

	var tea Tea
	err := row.Scan(&tea.Id, &tea.TeaName, &tea.TeaType, &tea.Origin.ShopName, &tea.Origin.ShopLocation, &tea.Temperature, &tea.PortionWeight, &tea.ContainerWeight, &tea.InitialWeight, &tea.BrewingDuration, &tea.Color, &tea.Size, &tea.InUse, &tea.BlendDescription)
	if err != nil {
		http.Error(w, fmt.Sprintf("Unable to scan row: %v", err), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(tea)
}

func createNewTea(w http.ResponseWriter, r *http.Request) {
	var tea Tea
	err := json.NewDecoder(r.Body).Decode(&tea)
	if err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	_, err = pool.Exec(context.Background(), `
			INSERT INTO teas (
				teaName, teaType, shopName, shopLocation, temperature, portionWeight,
				containerWeight, initialWeight, brewingDuration, color, size, inUse, blendDescription
			) VALUES (
				$1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13
			)`,
		tea.TeaName, tea.TeaType, tea.Origin.ShopName, tea.Origin.ShopLocation, tea.Temperature, tea.PortionWeight,
		tea.ContainerWeight, tea.InitialWeight, tea.BrewingDuration, tea.Color, tea.Size, tea.InUse, tea.BlendDescription,
	)

	if err != nil {
		json.NewEncoder(w).Encode(err)
	}

	json.NewEncoder(w).Encode(tea)
}

func updateTea(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPut {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		return
	}

	routeVariables := mux.Vars(r)
	teaID := routeVariables["id"]

	var tea Tea
	err := json.NewDecoder(r.Body).Decode(&tea)
	if err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	query := `UPDATE teas SET
        teaName = $1,
        teaType = $2,
        shopName = $3,
        shopLocation = $4,
        temperature = $5,
        portionWeight = $6,
        containerWeight = $7,
        initialWeight = $8,
        brewingDuration = $9,
        color = $10,
        size = $11,
        inUse = $12,
        blendDescription = $13
        WHERE id = $14`

	_, err = pool.Exec(context.Background(), query,
		tea.TeaName, tea.TeaType, tea.Origin.ShopName, tea.Origin.ShopLocation, tea.Temperature,
		tea.PortionWeight, tea.ContainerWeight, tea.InitialWeight, tea.BrewingDuration,
		tea.Color, tea.Size, tea.InUse, tea.BlendDescription, teaID)

	if err != nil {
		http.Error(w, fmt.Sprintf("Failed to update tea: %v", err), http.StatusInternalServerError)
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(tea)
}

func deleteTea(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodDelete {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		return
	}

	routeVariables := mux.Vars(r)
	teaID := routeVariables["id"]
	if teaID == "" {
		http.Error(w, "Tea ID is required", http.StatusBadRequest)
		return
	}

	// Convert tea ID to int64
	id, err := strconv.ParseInt(teaID, 10, 64)
	if err != nil {
		http.Error(w, "Invalid tea ID", http.StatusBadRequest)
		return
	}

	query := "DELETE FROM teas WHERE id = $1"
	_, err = pool.Exec(context.Background(), query, id)
	if err != nil {
		http.Error(w, fmt.Sprintf("Failed to delete tea: %v", err), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode("Teas Deleted")
}

func getTeaStatus(w http.ResponseWriter, r *http.Request) {
	rows, err := Database.Query("SELECT * FROM teas")
	if err != nil {
		log.Fatal(err)
	}

	defer rows.Close()

	var teas []Tea
	var teaStatus TeaStatus

	for rows.Next() {
		var tea Tea
		err := rows.Scan(&tea.TeaName, &tea.Origin.ShopName, &tea.Origin.ShopLocation, &tea.Temperature, &tea.PortionWeight, &tea.ContainerWeight, &tea.InitialWeight, &tea.BrewingDuration, &tea.Id, &tea.TeaType, &tea.Color, &tea.InUse, &tea.Size, &tea.BlendDescription)
		if err != nil {
			log.Fatal(err)
			return
		}
		teas = append(teas, tea)
	}

	for _, tea := range teas {
		if tea.InUse != 0 {
			teaStatus.Active = append(teaStatus.Active, tea)
		}
		teaStatus.Inactive = append(teaStatus.Inactive, tea)
	}

	json.NewEncoder(w).Encode(teaStatus)
}
