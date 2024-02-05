package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

// The function "respondWithJSON" is used to send a JSON response with a specified HTTP status code and
// payload.
func respondWithJSON(w http.ResponseWriter, code int, payload interface{}){
	// convert the payload to a JSON String and return it in Bytes format
	data, err := json.Marshal(payload)
	if err != nil {
		fmt.Printf("Failed tomarshal JSON response payload %v", payload )
		w.WriteHeader(http.StatusInternalServerError) // 500
	}

	// The code `w.Header().Add(("Content-Type"), "application/json")` is setting the "Content-Type" header
	// of the HTTP response to "application/json". This tells the client that the response body will be in
	// JSON format.
	w.Header().Add(("Content-Type"), "application/json")
	w.WriteHeader(code) 
	// `w.Write(data)` is writing the JSON data to the response body. It takes the `data` variable, which
	// contains the JSON payload in bytes format, and writes it to the `http.ResponseWriter` object `w`.
	// This sends the JSON response to the client.
	w.Write(data)
}


func responseWithError(w http.ResponseWriter, code int, message string){
	if code > 499 {
		fmt.Println("Server Error [5XX]: ", message)
	}

	type errResponse struct {
		// The `Error string `json:"error"`` is a struct tag in Go. It is used to specify the JSON key for
		// the `Error` field in the JSON response.
		// In our case the key will be "error"
		Error string `json:"error"`
	}	

	// `errResponse{Error: message}` is creating a new instance of the
	// `errResponse` struct and setting the value of the `Error` field to the
	// `message` parameter. This is done using a struct literal syntax in Go.
	
	
	respondWithJSON(w, code, errResponse{Error: message})
}