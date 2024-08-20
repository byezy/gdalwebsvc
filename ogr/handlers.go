package ogr

import (
	"fmt"
	"log"
	"net/http"
	"os/exec"
)

func ogrHome(w http.ResponseWriter, r *http.Request) {
	_, err := fmt.Fprintf(w, "OGR Home")
	if err != nil {
		log.Println(err)
	}
}

func ogrHelp(w http.ResponseWriter, r *http.Request) {
	fmt.Println("OGR Help")

	out, err := exec.Command("ogr2ogr", "--help").Output()
	if err != nil {
		fmt.Printf("ERROR: %s\n", err)
	}

	_, err = fmt.Fprintf(w, "OGR Help\n %s", out)
	if err != nil {
		fmt.Printf("ERROR: %s\n", err)
	}
}

func SetupRoutes(mux *http.ServeMux) {

	mux.HandleFunc("/", ogrHome)
	mux.HandleFunc("/help", ogrHelp)

}
