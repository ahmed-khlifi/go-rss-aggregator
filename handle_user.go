package main

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/ahmed-khlifi/go-rss-aggregator/internal/database"
	"github.com/google/uuid"
)

func (apiCfg *apiConfig) handleCreatUser(w http.ResponseWriter, r *http.Request) {

	type parameters struct {
		Name string `json:"name"`
	}

	decoder := json.NewDecoder(r.Body)
	paramas := parameters{}
	err := decoder.Decode(&paramas)
	if err != nil {
		responseWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}

	user, err := apiCfg.DB.CreateUser(r.Context(), database.CreateUserParams{
		ID: uuid.New(),
		CreatedAt:  time.Now().UTC(),
		UpdatedAt : time.Now().UTC(),
		Name: paramas.Name,
	})
	if err != nil {
		responseWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	respondWithJSON(w, http.StatusCreated, user)
}



func (apiCfg *apiConfig) handlerGetUser(w http.ResponseWriter, r *http.Request, user database.User) {
	respondWithJSON(w, http.StatusOK, user)
}


 
func (apiCfg *apiConfig) handlerGetPostsForUser(w http.ResponseWriter, r *http.Request, user database.User) {
	posts, err := apiCfg.DB.GetPostsForUser(r.Context(), database.GetPostsForUserParams{
		UserID: user.ID,
		Limit: 10,
	})
	if err != nil {
		responseWithError(w, http.StatusNotFound, "No posts found for this user")
		return
	}
	respondWithJSON(w, http.StatusOK, posts)

}

