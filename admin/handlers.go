package admin

import (
	"fmt"
	"log"
	"net/http"
	"os/exec"
)

func lsHome(w http.ResponseWriter, r *http.Request) {
	fmt.Println("ls /data")
	out, err := exec.Command("ls", "/").Output()
	if err != nil {
		log.Printf("ERROR: %s\n", err)
	}

	_, err = fmt.Fprintf(w, "ls / :\n %s", out)
	if err != nil {
		log.Printf("ERROR: %s\n", err)
	}
}

func lsData(w http.ResponseWriter, r *http.Request) {
	fmt.Println("ls /data")
	out, err := exec.Command("ls", "/data").Output()
	if err != nil {
		log.Printf("ERROR: %s\n", err)
	}

	_, err = fmt.Fprintf(w, "ls /data :\n %s", out)
	if err != nil {
		log.Printf("ERROR: %s\n", err)
	}
}

func lsOutput(w http.ResponseWriter, r *http.Request) {
	fmt.Println(r)
	fmt.Println("ls /output")
	out, err := exec.Command("ls", "/output").Output()
	if err != nil {
		log.Printf("ERROR: %s\n", err)
	}

	_, err = fmt.Fprintf(w, "ls /output :\n %s", out)
	if err != nil {
		log.Printf("ERROR: %s\n", err)
	}
}
func hello(rw http.ResponseWriter, r *http.Request) {
	u := r.Context().Value("user")
	user := "unset"
	if u != nil {
		user = u.(string)
	}

	switch r.Method {
	case http.MethodGet:
		if _, err := rw.Write([]byte("Hello " + user + "\n")); err != nil {
			fmt.Println("error when writing response for /hello request")
			rw.WriteHeader(http.StatusInternalServerError)
		}
	case http.MethodPost:
		if err := r.ParseForm(); err != nil { // Parses the request body
			fmt.Println("error when parsing form")
			rw.WriteHeader(http.StatusInternalServerError)
			_, _ = rw.Write([]byte("The POST parameters were not valid"))
		}
		x := r.Form.Get("user") // x will be "" if parameter is not set
		fmt.Println(x)

		if _, err := rw.Write([]byte("Thanks for posting to me, " + x + "\n")); err != nil {
			fmt.Println("error when writing response for /hello request")
			rw.WriteHeader(http.StatusInternalServerError)
		}
	}
}

func ping(rw http.ResponseWriter, req *http.Request) {
	switch req.Method {
	case http.MethodGet:
		if _, err := rw.Write([]byte("pong\n")); err != nil {
			fmt.Println("error when writing response for /ping request")
			rw.WriteHeader(http.StatusInternalServerError)
		}
	}
}

func SetupRoutes(mux *http.ServeMux) {

	mux.HandleFunc("/hello", hello)
	mux.HandleFunc("/ping", ping)
	mux.HandleFunc("/data", lsData)
	mux.HandleFunc("/output", lsOutput)
	mux.HandleFunc("/ls", lsHome)

}
