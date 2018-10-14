package main

import (
	"fmt"
	"net/http"
	"os"
	"path"
	"time"
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
var url = ""
func readURL(w http.ResponseWriter, r *http.Request) {
	start := time.Now()
	url = r.URL.String()
	dir, file := path.Split(url)
	fmt.Fprintln(w, dir)
	fmt.Fprintln(w, file)
	if dir == "/igcinfo/" && file == "api" {
		uptime := time.Now().Sub(start)
		fmt.Fprintf(w, "Uptime: %s\n", uptime)
		fmt.Fprintln(w, "Service for tracking igc files")
		fmt.Fprintln(w, "Version 0.8")
	}
}

func main() {
	port := os.Getenv("PORT")
	http.HandleFunc("/", readURL)
	http.ListenAndServe(":"+port, nil)
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