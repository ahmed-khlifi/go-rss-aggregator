package main

import "net/http"

func handleError(w http.ResponseWriter, r *http.Request){
	responseWithError(w, 400, "Something wen wrong")
}