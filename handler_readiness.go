package main

import "net/http"

/*
that serves as an HTTP handler for readiness checks. This function is typically used in the context of a web server to handle HTTP requests at a specific route.*/
func handlerReadiness(w http.ResponseWriter, r *http.Request){
	respondWithJSON(w, http.StatusOK, struct{}{})
}

/* 
`w http.ResponseWriter` is an interface that allows the server to write an HTTP response. You can use it to set HTTP headers, write the body of the response, set the HTTP status code, etc.
`r *http.Request` is a pointer to an http.Request object. This object contains all the information about the HTTP request that was received by the server, such as the HTTP method (GET, POST, etc.), the URL, headers, and body data.
 */