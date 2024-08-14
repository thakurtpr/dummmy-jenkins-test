package main

import (
    "fmt"
    "log"
    "net/http"

    "github.com/gorilla/mux"
)

func main() {
    // Create a new router
    r := mux.NewRouter()

    // Define a single endpoint
    r.HandleFunc("/jenkins", JenkinsEndpointHandler).Methods("GET")

    // Start the server on port 8080
    log.Println("Starting server on :7070...")
    log.Fatal(http.ListenAndServe(":7070", r))
}

// JenkinsEndpointHandler handles the /jenkins-endpoint route
func JenkinsEndpointHandler(w http.ResponseWriter, r *http.Request) {
    // Respond with a simple message
    fmt.Fprintf(w, "Jenkins Endpoint Hit!")
}
