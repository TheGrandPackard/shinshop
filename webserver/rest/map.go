package rest

import (
	"encoding/csv"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"strings"

	"github.com/Xackery/eqemuconfig"
	"github.com/Xackery/shinshop/database"
	"github.com/Xackery/shinshop/database/spawn"
)

type Line struct {
	X1 float64 `json:"x1"`
	Y1 float64 `json:"y1"`
	X2 float64 `json:"x2"`
	Y2 float64 `json:"y2"`
}

type SpawnPoint struct {
	SpawnGroupID int     `json:"spawn_group_id"`
	X            float64 `json:"x"`
	Y            float64 `json:"y"`
	Z            float64 `json:"z"`
}

func MapGetByShortname(w http.ResponseWriter, r *http.Request) {

	log.Println("Serving map request")
	var err error

	type Index struct {
		*Site
		Lines       []Line       `json:"lines"`
		SpawnPoints []SpawnPoint `json:"spawn_points"`
	}

	resp := Index{
		Site: SiteInit(),
	}

	if r.Method != http.MethodPost {
		log.Println("Only POST methods are accepted")
		return
	}

	name := r.FormValue("name")
	if len(name) == 0 {
		log.Println("Post value of name required")
		return
	}

	bMap, err := Asset(fmt.Sprintf("rest/map/%s_1.txt", name))
	if err != nil {
		log.Println("Error finding map (" + name + "): " + err.Error())
		return
	}

	reader := csv.NewReader(strings.NewReader(string(bMap)))
	records, err := reader.ReadAll()
	if err != nil {
		log.Println("Error reading map file", err.Error())
		return
	}

	for _, record := range records {
		entries := strings.Split(record[0], " ")
		drawType := entries[0]
		if drawType == "L" {
			line := Line{}
			line.X1, _ = strconv.ParseFloat(strings.TrimSpace(entries[1]), 64)
			line.Y1, _ = strconv.ParseFloat(strings.TrimSpace(record[1]), 64)
			line.X2, _ = strconv.ParseFloat(strings.TrimSpace(record[3]), 64)
			line.Y2, _ = strconv.ParseFloat(strings.TrimSpace(record[4]), 64)
			resp.Lines = append(resp.Lines, line)
		}
	}

	config, err := eqemuconfig.GetConfig()
	if err != nil {
		log.Println("Error getting config", err.Error())
		return
	}

	db, err := database.Connect(config)
	if err != nil {
		log.Println("Error connecting to DB:", err.Error())
		return
	}

	spawnPoints, err := spawn.GetSpawnsByZone(db, name)
	if err != nil {
		log.Println("Error getting spawn points:", err.Error())
		return
	}

	for _, spawnPoint := range spawnPoints {
		spawn := SpawnPoint{
			SpawnGroupID: spawnPoint.Spawngroupid,
			X:            -spawnPoint.X,
			Y:            -spawnPoint.Y,
		}
		resp.SpawnPoints = append(resp.SpawnPoints, spawn)
	}

	resp.Status = 1
	resp.Message = "Here's the map"
	err = json.NewEncoder(w).Encode(resp)
	if err != nil {
		log.Println("Error requesting RestIndex:", err.Error())
	}
}
