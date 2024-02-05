package main

import (
	"fmt"
	"net/http"

	"github.com/ahmed-khlifi/go-rss-aggregator/internal/auth"
	"github.com/ahmed-khlifi/go-rss-aggregator/internal/database"
)

/*
The provided Go code is a middleware function for authentication in an HTTP server.
Middleware is a design pattern used in web server frameworks for injecting behavior or preprocessing requests before they reach the actual route handlers.
This particular middleware function is used for authenticating users based on an API key.

Let's break down the code:

The authedHandler type is a function type that takes an http.ResponseWriter, a pointer to an http.Request, and a database.User as parameters. This type is used to define the handlers that will be wrapped by the middlewareAuth function.

The middlewareAuth function is a method of the apiConfig type. It takes an authedHandler as a parameter and returns an http.HandlerFunc. The returned function is a closure that has access to the authedHandler parameter.

Inside the returned function, it first tries to get the API key from the request header by calling auth.GetAPIKey(r.Header). If there's an error (which likely means that the API key is missing or invalid), it responds with an HTTP 401 Unauthorized status code and returns.

Next, it tries to get the user associated with the API key by calling cfg.DB.GetUserByAPIKey(r.Context(), apiKey). If there's an error (which could mean that the API key is not associated with any user), it responds with an HTTP 401 Unauthorized status code and returns.

If both steps are successful, it calls the authedHandler function with the http.ResponseWriter, the http.Request, and the database.User as parameters. This allows the authedHandler to handle the request with the knowledge that the user has been authenticated.

In summary, this middleware function ensures that all requests handled by the authedHandler are authenticated. If a request is not authenticated, it immediately responds with an HTTP 401 Unauthorized status code and does not call the authedHandler.
*/

type authedHandler func(http.ResponseWriter, *http.Request, database.User)

func (cfg *apiConfig) middlewareAuth(handler authedHandler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		apiKey, err := auth.GetAPIKey(r.Header)
	if err != nil {
		responseWithError(w, http.StatusUnauthorized, "Unauthorized")
		return
	}
	
	user, err := cfg.DB.GetUserByAPIKey(r.Context(), apiKey)
	if err != nil {
		responseWithError(w, http.StatusUnauthorized, fmt.Sprintf("Failed to get user: %v", err.Error()))
		return
	}

	handler(w, r, user)

	}
}