package main

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"github.com/marni/goigc"
	"log"
	"net/http"
	"os"
	"path"
	"strconv"
	"strings"
	"time"
)

type Metadata struct {
	Uptime string `json:"uptime,omitempty"`
	Desc string `json:"desc,omitempty"`
	Version string `json:"version,omitempty"`
}

type Track struct {
	ID int `json:"id,omitempty"`
	URL string `json:"url,omitempty"`
}

type TrackInfo struct {
	FDate time.Time `json:"fdate,omitempty"`
	Pilot string `json:"pilot,omitempty"`
	Glider string `json:"glider,omitempty"`
	GliderId string `json:"glider_id,omitempty"`
	TrackLength int `json:"track_length,omitempty"`
}

type IDList struct {
	ID int `json:"id,omitempty"`
}

var idlist []IDList
var tracks []Track
var lastTrack = 0

func main() {
	router := mux.NewRouter()
	port := os.Getenv("PORT")

	router.HandleFunc("/igcinfo/api", getMetadata).Methods("GET")
	router.HandleFunc("/igcinfo/api/igc", registerTrack).Methods("POST")
	router.HandleFunc("/igcinfo/api/igc", getIDs).Methods("GET")
	router.HandleFunc("/igcinfo/api/igc/{id}", getTrackMeta).Methods("GET")
	router.HandleFunc("/igcinfo/api/igc/{id}/{field}", getTrackMetaField).Methods("GET")

	http.ListenAndServe(":"+port, router)
}

func getMetadata(w http.ResponseWriter, r *http.Request) {
	metadata := Metadata{"Yes", "Service for IGC tracks", "v0.8"}
	json.NewEncoder(w).Encode(metadata)
}

	// Reads a URL as a parameter, makes a new track for it in memory, and writes out the new id in json format
func registerTrack(w http.ResponseWriter, r *http.Request) {
	url, err := r.URL.Query()["url"]
	if !err || len(url[0]) < 1 {
		log.Println("URL parameter is missing")
	} else {	// If a URL is sent
		var track Track
		var id IDList
		_ = json.NewDecoder(r.Body).Decode(&track)
		track.URL = string(url[0])
		lastTrack += 1
		track.ID = lastTrack
		id.ID = lastTrack
		tracks = append(tracks, track)
		idlist = append(idlist, id)
		jsonConverter := fmt.Sprintf(`"{"id":%d}"`, track.ID)
		output := []byte(jsonConverter)
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write(output)
	}
}

	// Writes all the registered IDs
func getIDs(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(idlist)
}

	// Writes information about a specific track registered in the memory
func getTrackMeta(w http.ResponseWriter, r *http.Request) {
	url := r.URL.String()
	_, input := path.Split(url)

	in, err := strconv.Atoi(input)
	if err != nil {
		log.Fatal(err)
	}

	if in <= lastTrack {
		t, e := igc.ParseLocation(tracks[in - 1].URL)
		if e != nil {
			log.Fatal(e)
		}

		info := TrackInfo{t.Date, t.Pilot, t.GliderType, t.GliderID, 9}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(info)

	} else {
		w.WriteHeader(http.StatusNotFound)
	}
}

	// Writes a specific piece of information about a specific track
func getTrackMetaField(w http.ResponseWriter, r *http.Request) {
	url := r.URL.String()
	temp := strings.Split(url, "/")
	f := temp[5]
	t := temp[4]

	fmt.Fprintln(w, f)
	fmt.Fprintln(w, t)


	/*
	in, err := strconv.Atoi(field)
	if err != nil {
		log.Fatal(err)
	}*/

}