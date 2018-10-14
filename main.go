package main

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"os"
)

/*func main() {
	s := "http://skypolaris.org/wp-content/uploads/IGS%20Files/Madrid%20to%20Jerez.igc"
	track, err := igc.ParseLocation(s)
	if err != nil {
		fmt.Errorf("Problem reading the track", err)
	}

	fmt.Printf("Pilot: %s, gilderType: %s, date: %s",
		track.Pilot, track.GliderType, track.Date.String())

	if len(os.Args) != 2 {
		os.Exit(1)
	}
	response, err := http.Get(os.Args[1])
	if err != nil {
		log.Fatal(err)
	} else {
		fmt.Print(response)
	}


}*/

/*
start := time.Now()
	url = r.URL.String()
	dir, file := path.Split(url)
	fmt.Fprintln(w, dir)
	fmt.Fprintln(w, file)
	if dir == "/igcinfo/" && file == "api" {
		uptime := time.Now().Sub(start)
		fmt.Fprintf(w, "Uptime: %s\n", uptime)
		fmt.Fprintln(w, "Service for IGC tracks")
		fmt.Fprintln(w, "Version 0.8")
	} else if dir == "/igcinfo/api/" && file == "igc" {

	} else {
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprintln(w,"404 - Page not found")
	}

 */
type Metadata struct {
	Uptime string `json:"uptime,omitempty"`
	Desc string `json:"desc,omitempty"`
	Version string `json:"version,omitempty"`
}

type Track struct {
	ID int `json:"id,omitempty"`
	URL string `json:"url,omitempty"`
}

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

func registerTrack(w http.ResponseWriter, r *http.Request) {
	url, err := r.URL.Query()["url"]
	if !err || len(url[0]) < 1 {
		log.Println("URL parameter is missing")
	} else {
		var track Track
		_ = json.NewDecoder(r.Body).Decode(&track)
		track.URL = string(url[0])
		lastTrack += 1
		track.ID = lastTrack
		tracks = append(tracks, track)
		// json.NewEncoder(w).Encode(track.ID)
		jsonConverter := fmt.Sprintf(`"{"id":%d}"`, track.ID)
		output := []byte(jsonConverter)
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write(output)
	}
}

func getIDs(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(tracks)
}

func getTrackMeta(w http.ResponseWriter, r *http.Request) {

}

func getTrackMetaField(w http.ResponseWriter, r *http.Request) {

}

/*
var url = ""
func readURL(w http.ResponseWriter, r *http.Request) {
	url = r.URL.Path[1:]
}

func pageNotFound(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusNotFound)
	w.Write([]byte("404 - Page not found!"))
	fmt.Fprint(w, "404 - Page not found!")
}

func main() {
	port := os.Getenv("PORT")
	http.HandleFunc("/", readURL)
	http.ListenAndServe(":"+port, nil)

	http.HandleFunc("/", pageNotFound)
}
*/