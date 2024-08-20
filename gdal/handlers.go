package gdal

import (
	"fmt"
	"log"
	"net/http"
	"os/exec"
)

func gdalHome(w http.ResponseWriter, r *http.Request) {
	_, err := fmt.Fprintf(w, "GDAL Home")
	if err != nil {
		log.Println(err)
	}
}

func gdalHelp(w http.ResponseWriter, r *http.Request) {
	fmt.Println("GDAL Help")

	out, err := exec.Command("gdalinfo", "--help").Output()
	if err != nil {
		log.Printf("ERROR: %s\n", err)
	}

	_, err = fmt.Fprintf(w, "GDAL Help\n %s", out)
	if err != nil {
		log.Printf("ERROR: %s\n", err)
	}
}

func SetupRoutes(mux *http.ServeMux) {

	mux.HandleFunc("/", gdalHome)
	mux.HandleFunc("/help", gdalHelp)

}
