package main

import (
	"fmt"
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
var url = ""

func getUrl(w http.ResponseWriter, r *http.Request) {
	url = r.URL.Path
	fmt.Fprint(w, r.URL.Path[1:7])
	fmt.Fprint(w, " - ")
	fmt.Fprint(w, r.URL.Path[8:13])
	fmt.Fprint(w, " - ")
	fmt.Fprint(w, url)
}

func main() {
	port := os.Getenv("PORT")
	http.HandleFunc("/", getUrl)

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