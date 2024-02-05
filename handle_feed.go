package main

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/ahmed-khlifi/go-rss-aggregator/internal/database"
	"github.com/google/uuid"
)

func (apiCfg *apiConfig) handleCreatFeed(w http.ResponseWriter, r *http.Request, user database.User) {

	type parameters struct {
		Name string `json:"name"`
		URL string `json:"url"`

	}

	decoder := json.NewDecoder(r.Body)
	paramas := parameters{}
	err := decoder.Decode(&paramas)
	if err != nil {
		responseWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}

	feed, err := apiCfg.DB.CreateFeed(r.Context(), database.CreateFeedParams{
		ID: uuid.New(),
		CreatedAt:  time.Now().UTC(),
		UpdatedAt : time.Now().UTC(),
		Name: paramas.Name,
		Url: paramas.URL,
		UserID: user.ID,
	})
	if err != nil {
		responseWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	respondWithJSON(w, http.StatusCreated, feed)
}

 
func (apiCfg *apiConfig) handleGetFeeds(w http.ResponseWriter, r *http.Request) {
 

	feeds, err := apiCfg.DB.GetFeeds(r.Context())
	if err != nil {
		responseWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	respondWithJSON(w, http.StatusCreated, feeds)
}
