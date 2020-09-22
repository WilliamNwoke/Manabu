package main

import (
	"log"
	"net/http"
)

func apiResponse(w http.ResponseWriter, r *http.Request) {
	// Set the return content-Type as JSON like before
	w.Header().Set("Content-Type", "application/json")

	// Change the response depending on the method being requested
	switch r.Method {
	case "GET":
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"message":"GET method requested"}`))
	case "POST":
		w.WriteHeader(http.StatusCreated)
		w.Write([]byte(`{"message":"POST method requested"}`))
	case "VIEW":
		w.WriteHeader(http.StatusCreated)
		w.Write([]byte(`{"message":"View method requested"}`))
	case "DELETE":
		w.WriteHeader(http.StatusCreated)
		w.Write([]byte(`{"message":"Delete method requested"}`))
	case "PUT":
		w.WriteHeader(http.StatusCreated)
		w.Write([]byte(`{"message":"Put method requested"}`))
	case "COPY":
		w.WriteHeader(http.StatusCreated)
		w.Write([]byte(`{"message":"Copy method requested"}`))

	default:
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte(`{"message":"Can't find method requested"}`))
	}
}

func main() {
	http.HandleFunc("/", apiResponse)
	log.Fatal(http.ListenAndServe(":8080", nil))

}
